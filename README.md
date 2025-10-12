# Client-Server API

> Projeto Go para demonstração de arquitetura client-server com organização modular, controle de usuários, e integração com banco de dados SQL.

---

## Sobre o Projeto

Este projeto implementa uma API de exemplo em Go, utilizando arquitetura dividida entre cliente e servidor, com organização interna por camadas (handlers, models, services etc). O objetivo é servir como referência para boas práticas de estruturação de projetos Go e uso de banco de dados relacional. Projeto em GO será utilizado apenas para demandas que exigem maior complexidade e necessidade de melhor performace.

---

## Estrutura de Pastas

```
CLIENT-SERVER-API/
├── cmd/
│   ├── client/
│   │   └── main.go
│   └── server/
│       └── main.go -- Execução de api
├── database/
│   ├── diagram.sql
│   └── users.sql
├── internal/ -- pastas responsáveis por gerir a API
│   ├── config/
│   ├── db/
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── requests/
│   ├── response/
│   ├── router/
│   └── service/
├── .env
├── go.mod
└── go.sum