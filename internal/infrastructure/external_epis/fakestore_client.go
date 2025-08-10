package external_epis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	ErrProductNotFound = errors.New("produto não encontrado na API externa")
)

type ExternalProduct struct {
	ID    uint    `json:"id"`
	Title string  `json:"title"`
	Image string  `json:"image"`
	Price float32 `json:"price"`
}

type ProductClient interface {
	GetProduct(ctx context.Context, id uint) (ExternalProduct, error)
}

// ProductCache implementa um cache em memória para produtos
type ProductCache struct {
	cache map[uint]ExternalProduct
	mutex sync.RWMutex
	ttl   time.Duration
}

func NewProductCache() *ProductCache {
	return &ProductCache{
		cache: make(map[uint]ExternalProduct),
		ttl:   15 * time.Minute, // 15 minutos de TTL
	}
}

func (c *ProductCache) Get(id uint) (ExternalProduct, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	p, ok := c.cache[id]
	return p, ok
}

func (c *ProductCache) Set(id uint, product ExternalProduct) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[id] = product
}

// Circuit breaker state
type CircuitBreaker struct {
	FailureThreshold   int
	FailureCount       int
	OpenUntil          time.Time
	CooldownPeriod     time.Duration
	HalfOpenMaxRetries int
	HalfOpenRetries    int
	mutex              sync.Mutex
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{
		FailureThreshold:   5,           // 5 falhas consecutivas
		CooldownPeriod:     time.Minute, // 1 minuto no estado aberto
		HalfOpenMaxRetries: 2,           // Tenta 2 vezes no modo semi-aberto
	}
}

func (cb *CircuitBreaker) IsOpen() bool {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	// Se o tempo de esfriamento passou, move para half-open
	if cb.OpenUntil.Before(time.Now()) {
		if cb.FailureCount >= cb.FailureThreshold {
			cb.HalfOpenRetries = 0 // Reset para tentar alguns requests
		}
		return false
	}
	return cb.FailureCount >= cb.FailureThreshold
}

func (cb *CircuitBreaker) Success() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.FailureCount = 0
	cb.HalfOpenRetries = 0
}

func (cb *CircuitBreaker) Failure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	cb.FailureCount++

	// Se está em half-open e excedeu as tentativas, abre o circuito novamente
	if cb.OpenUntil.Before(time.Now()) && cb.FailureCount > cb.FailureThreshold {
		cb.HalfOpenRetries++
		if cb.HalfOpenRetries >= cb.HalfOpenMaxRetries {
			cb.OpenUntil = time.Now().Add(cb.CooldownPeriod)
		}
	}

	// Se atingiu o limite, abre o circuito
	if cb.FailureCount >= cb.FailureThreshold && cb.OpenUntil.IsZero() {
		cb.OpenUntil = time.Now().Add(cb.CooldownPeriod)
	}
}

type FakeStoreClient struct {
	baseURL string
	http    *http.Client
	cache   *ProductCache
	cb      *CircuitBreaker
}

func NewFakeStoreClient() *FakeStoreClient {
	return &FakeStoreClient{
		baseURL: "https://fakestoreapi.com",
		http: &http.Client{
			Timeout: 3 * time.Second, // Reduzido para 3 segundos
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          200, // Aumentado para 200
				MaxConnsPerHost:       20,  // Limita conexões por host
				MaxIdleConnsPerHost:   10,  // Limita conexões ociosas por host
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   3 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
		cache: NewProductCache(),
		cb:    NewCircuitBreaker(),
	}
}

func (c *FakeStoreClient) GetProduct(ctx context.Context, id uint) (ExternalProduct, error) {
	// Verifica se está no cache
	if product, found := c.cache.Get(id); found {
		return product, nil
	}

	// Verifica circuit breaker
	if c.cb.IsOpen() {
		return ExternalProduct{}, fmt.Errorf("circuit breaker aberto: serviço temporariamente indisponível")
	}

	// Prepara request com timeout e retry
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/products/%d", c.baseURL, id), nil)
	if err != nil {
		return ExternalProduct{}, err
	}

	// Adiciona headers essenciais
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "aiqfome-favorites-service/1.0")

	// Realiza request com retry (máx 2 tentativas)
	var resp *http.Response
	var lastErr error

	for attempt := 0; attempt < 2; attempt++ {
		resp, err = c.http.Do(req)
		if err == nil {
			break
		}

		lastErr = err

		// Se for o último retry, não espera
		if attempt < 1 {
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Se todas as tentativas falharam
	if err != nil {
		c.cb.Failure()
		return ExternalProduct{}, fmt.Errorf("falha após retry: %w", lastErr)
	}

	defer resp.Body.Close()

	// Verifica status code
	if resp.StatusCode == http.StatusNotFound {
		c.cb.Success() // Not Found é esperado, não indica falha do serviço
		return ExternalProduct{}, ErrProductNotFound
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		c.cb.Failure()
		return ExternalProduct{}, fmt.Errorf("erro ao consultar API externa: status %d", resp.StatusCode)
	}

	// Lê corpo com limite (evita ataques)
	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024)) // Máx 1MB
	if err != nil {
		c.cb.Failure()
		return ExternalProduct{}, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var p ExternalProduct
	if err := json.Unmarshal(bodyBytes, &p); err != nil {
		c.cb.Failure()
		return ExternalProduct{}, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	if p.ID == 0 {
		c.cb.Success() // Resposta com formato correto
		return ExternalProduct{}, ErrProductNotFound
	}

	// Cache o produto para futuras requisições
	c.cache.Set(id, p)

	// Marca sucesso no circuit breaker
	c.cb.Success()

	return p, nil
}
