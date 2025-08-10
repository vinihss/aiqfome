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
		FailureThreshold:   5,
		CooldownPeriod:     time.Minute,
		HalfOpenMaxRetries: 2,
	}
}

func (cb *CircuitBreaker) IsOpen() bool {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.OpenUntil.Before(time.Now()) {
		if cb.FailureCount >= cb.FailureThreshold {
			cb.HalfOpenRetries = 0
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

	if cb.OpenUntil.Before(time.Now()) && cb.FailureCount > cb.FailureThreshold {
		cb.HalfOpenRetries++
		if cb.HalfOpenRetries >= cb.HalfOpenMaxRetries {
			cb.OpenUntil = time.Now().Add(cb.CooldownPeriod)
		}
	}

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
			Timeout: 3 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          200,
				MaxConnsPerHost:       20,
				MaxIdleConnsPerHost:   10,
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
	if product, found := c.cache.Get(id); found {
		return product, nil
	}

	if c.cb.IsOpen() {
		return ExternalProduct{}, fmt.Errorf("circuit breaker aberto: serviço temporariamente indisponível")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/products/%d", c.baseURL, id), nil)
	if err != nil {
		return ExternalProduct{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "aiqfome-favorites-service/1.0")

	var resp *http.Response
	var lastErr error

	for attempt := 0; attempt < 2; attempt++ {
		resp, err = c.http.Do(req)
		if err == nil {
			break
		}

		lastErr = err

		if attempt < 1 {
			time.Sleep(100 * time.Millisecond)
		}
	}

	if err != nil {
		c.cb.Failure()
		return ExternalProduct{}, fmt.Errorf("falha após retry: %w", lastErr)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		c.cb.Success()
		return ExternalProduct{}, ErrProductNotFound
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		c.cb.Failure()
		return ExternalProduct{}, fmt.Errorf("erro ao consultar API externa: status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024))
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
		c.cb.Success()
		return ExternalProduct{}, ErrProductNotFound
	}

	c.cache.Set(id, p)
	c.cb.Success()

	return p, nil
}
