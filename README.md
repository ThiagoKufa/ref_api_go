# **Documentação de Referência para API-Rest em Go: CRUD e WebSocket**

## **Visão Geral**

Este projeto foi desenvolvido para servir como uma referência abrangente para a implementação de uma API WEB. Ele inclui a elaboração de um sistema CRUD (Create, Read, Update, Delete) e a integração de uma conexão WebSocket. As tecnologias utilizadas neste projeto são:

- **Go (Golang)**: Uma linguagem de programação de alto desempenho.
- **PostgreSQL**: Um sistema de gerenciamento de banco de dados relacional robusto.
- **Docker Compose**: Uma ferramenta para definir e gerenciar aplicações multi-container Docker.

---

## **Lista de Tarefas para Implementação da API**

- [ ]  **Criar Produto**: Método POST no endpoint **`/product`**.
- [ ]  **Buscar Todos os Produtos**: Método GET no endpoint **`/product`**.
- [ ]  **Buscar Produto por ID**: Método GET no endpoint **`/product/{ID}`**.
- [ ]  **Atualizar Produto**: Método PUT no endpoint **`/product/{ID}`**.
- [ ]  **Deletar Produto**: Método DELETE no endpoint **`/product/{ID}`**.
- [ ]  **Monitoramento de Produtos em Tempo Real**: Implementar funcionalidade para busca de produtos em tempo real (possivelmente utilizando WebSocket).