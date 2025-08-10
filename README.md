# Serviço de Favoritos

Microserviço responsável por gerenciar produtos favoritos dos clientes, com alta disponibilidade e performance.

## Funcionalidades

- Adicionar produtos aos favoritos
- Listar produtos favoritos com detalhes (ID, título, imagem, preço e review)
- Remover produtos dos favoritos
- Cache distribuído para alta performance
- Métricas e monitoramento

## Tecnologias

- Go 1.24
- PostgreSQL
- Redis
- Docker & Docker Compose
- Prometheus & Grafana

## Requisitos

- Go 1.24+
- Docker & Docker Compose
- Make

## Como executar

### Desenvolvimento local

```bash
# Iniciar todos os serviços com Docker Compose
make docker-compose

# Executar a aplicação localmente (fora do Docker)
make run

# Rodar testes
make test
```

### Produção

```bash
# Fazer build e push da imagem Docker
make deploy
```

## Estrutura do Projeto

```
.
├── cmd/           # Pontos de entrada da aplicação
├── config/        # Configurações
├── docs/          # Documentação (Swagger, etc)
├── internal/      # Código interno da aplicação
│   ├── domain/    # Modelos de domínio
│   ├── infrastructure/  # Implementações concretas
│   ├── interfaces/      # Adaptadores HTTP
│   ├── routes/          # Rotas da API
│   └── usecases/        # Casos de uso da aplicação
├── middlewares/   # Middlewares HTTP
└── services/      # Serviços auxiliares
```

## APIs

A documentação completa das APIs está disponível via Swagger em `/docs/swagger.json` ou acessando `/swagger/index.html` quando a aplicação estiver em execução.

Rotas principais:

- `GET /customer/{id}/favorites` - Lista produtos favoritos
- `POST /customer/{id}/favorites` - Adiciona produto aos favoritos
- `DELETE /customer/{id}/favorites/{productId}` - Remove produto dos favoritos

## Monitoramento

- Métricas: `/metrics` (formato Prometheus)
- Health check: `/health`
- Grafana: `http://localhost:3000` (usuário: admin, senha: admin)

## Escalabilidade e Alta Disponibilidade

Este serviço foi projetado para:

- Escalar horizontalmente (múltiplas instâncias)
- Utilizar cache distribuído (Redis)
- Implementar circuit breaker para API externa
- Monitoramento em tempo real
- Health checks para auto-recuperação

## Contribuindo

1. Faça fork do projeto
2. Crie sua feature branch (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request
