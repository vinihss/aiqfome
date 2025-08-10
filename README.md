# Serviço de Favoritos

Microserviço responsável por gerenciar produtos favoritos dos clientes, com alta disponibilidade e performance.

## Funcionalidades
- Adicionar produtos aos favoritos
- Adicionar produtos aos favoritos
- Listar produtos favoritos com detalhes (ID, título, imagem, preço e review)
- Remover produtos dos favoritos
- Cache distribuído para alta performance

## Tecnologias

- Go 1.24
- PostgreSQL
- Redis
- Docker & Docker Compose

## Requisitos

- Go 1.24+
- Docker & Docker Compose
- Make

## Como Utilizar
- Subir os containers da aplicação
```
docker-composer up -d
```
- A documentação estará disponível em: http://localhost:8080/swagger/index.html
- Realizar autenticão no endpoint authenticate. Os parametros são apenas demostrativos, e irá retornar um token de deve ser utilizado nos demais endpoins.
```
{
  "email": "string",
  "name": "string"
}
```
- Utilizar o token no header da seguinte maneira:
```
Bearer <token>
```
## Estrutura do Projeto

```
.
├── cmd/           # Pontos de entrada da aplicação
├── config/        # Configurações
├── docs/          # Documentação (Swagger)
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


## Escalabilidade e Alta Disponibilidade

Este serviço foi projetado para:

- Escalar horizontalmente (múltiplas instâncias)
- Utilizar cache distribuído (Redis)
- Implementar circuit breaker para API externa
