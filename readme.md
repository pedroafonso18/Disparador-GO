# Documentação do Projeto Go-Disparador

## Visão Geral
Este projeto é responsável por buscar dados de uma base de dados já alimentada por outro painel, e posteriormente utilizar essas informações para disparar mensagens via APIs externas, como Wuzapi e Evolution.

## Funcionalidades Atuais
- **Verificação de Horário:** 
  - O serviço verifica se o horário atual está entre um intervalo configurado (08:00 e 17:00).
- **Consulta na Base de Dados:**
  - Executa queries para buscar conexões e campanhas a partir da base de dados PostgreSQL.
- **Integração com APIs Externas:**
  - Como próximo passo, integrará com APIs de disparo de mensagens (Wuzapi e Evolution).

## Utilização
- Execute o serviço utilizando o comando apropriado (detalhes sobre execução podem ser incluídos conforme a necessidade).
- As configurações, incluindo a URL do Banco de Dados, devem ser definidas em um arquivo `.env`.

## Dependências
- PostgreSQL com driver pgx.
- Módulo godotenv para carregamento das variáveis de ambiente.

## Futuras Melhorias
- Integração completa com as APIs Wuzapi e Evolution para o disparo de mensagens.
- Ajustes e melhorias na lógica de verificação de horário e tratamento de erros.

