# Arquitetura do Microserviço de Favoritos

## Visão Geral

Este microserviço gerencia os produtos favoritos dos clientes, integrando-se com APIs externas para obter detalhes de produtos.

## Componentes

- **API Gateway**: Entrada única para o serviço, gerencia autenticação e rate limiting
- **Instâncias de Serviço**: Múltiplas réplicas do serviço para balanceamento de carga
- **Cache**: Redis para caching de produtos e favoritos frequentemente acessados
- **Banco de Dados**: PostgreSQL para armazenamento de dados, com réplicas de leitura
- **Sistema de Mensageria**: Kafka para processamento assíncrono e integração com outros serviços

## Escalabilidade

- Escala horizontal via contêineres (Kubernetes)
- Separação entre operações de leitura e escrita no banco de dados
- Cache distribuído para reduzir chamadas à API externa e ao banco de dados

## Alta Disponibilidade

- Múltiplas zonas de disponibilidade
- Failover automático para réplicas de banco de dados
- Circuit breaker para proteção contra falhas na API externa
- Health checks e auto-recuperação

## Monitoramento

- Prometheus para métricas
- Jaeger para tracing distribuído
- Loki para logs centralizados
- Alertas automatizados
