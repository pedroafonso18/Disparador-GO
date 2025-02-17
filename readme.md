# Go-Disparador

## Visão Geral
Sistema de disparo automatizado de mensagens que integra com múltiplas APIs de mensageria. O projeto busca dados de campanhas em uma base PostgreSQL e realiza o disparo através das APIs do Wuzapi e Evolution.

## Funcionalidades
- **Controle de Horário:** 
  - Serviço opera apenas entre 08:00 e 17:00
  - Verificação automática do horário atual
  - Prevenção de disparos fora do horário comercial

- **Gerenciamento de Base de Dados:**
  - Conexão com PostgreSQL
  - Busca de campanhas ativas
  - Controle de instâncias de disparo
  - Tracking de mensagens enviadas

- **APIs Integradas:**
  - Evolution API:
    - Envio de mensagens de texto
    - Suporte a múltiplas instâncias
  - Wuzapi:
    - Envio de mensagens via endpoint específico
    - Autenticação via token

## Configuração
O projeto utiliza variáveis de ambiente através de um arquivo `.env`:
- `DB_URL`: URL de conexão com o PostgreSQL
- `EVO_URL`: Endpoint base da API Evolution
- `EVO_TOKEN`: Token de autenticação Evolution
- `WUZ_URL`: URL base da API Wuzapi

## Estrutura do Projeto
```
internal/
  ├── api/         # Integrações com Evolution e Wuzapi
  ├── config/      # Configurações e variáveis de ambiente
  ├── database/    # Operações com PostgreSQL
  └── services/    # Lógica de negócio e controle de tempo
```

## Como Executar
1. Configure o arquivo `.env` com as credenciais necessárias
2. Execute o comando:
   ```bash
   go run cmd/main.go
   ```

## Dependências Principais
- `github.com/jackc/pgx/v5`: Driver PostgreSQL
- `github.com/joho/godotenv`: Gerenciamento de variáveis de ambiente

