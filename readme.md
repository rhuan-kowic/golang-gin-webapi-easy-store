# **Easy Store - Sistema de Gestão de Vendas e Estoque**

**Easy Store** é um sistema simples para gerenciar a venda de produtos e o controle de estoque. Ele permite cadastrar produtos, registrar vendas, controlar o estoque, gerar relatórios de vendas e receber alertas quando o estoque de um produto estiver baixo. Ideal para pequenos empresários que buscam uma solução prática e eficiente.

## **Tecnologias Utilizadas**

- **Backend**: Go (Golang)
- **Banco de Dados**: Postgres
- **Framework para API**: [Gin](https://github.com/gin-gonic/gin)

## **Funcionalidades**

1. **Cadastro de Produtos**:
   - Cadastrar, editar e excluir produtos.
   - Informações do produto: nome, descrição, preço, quantidade em estoque.
2. **Controle de Vendas**:
   - Registrar vendas de produtos.
   - Atualizar automaticamente a quantidade de estoque ao registrar uma venda.
   - Exibir total de vendas por dia, semana ou mês.
3. **Controle de Estoque**:
   - Visualizar a quantidade de cada produto em estoque.
   - Receber alertas quando o estoque de um produto estiver abaixo do limite mínimo.
4. **Relatórios**:
   - Gerar relatórios de vendas por período (dia, semana, mês).
   - Exibir produtos vendidos e estoque em falta.
5. **Cadastro de Clientes**:
   - Registrar informações básicas dos clientes (nome e contato).
6. **Interface Simples e Intuitiva**:
   - Tela inicial com login.
   - Página de gerenciamento de produtos, vendas e estoque.
   - Geração de relatórios simples de vendas e estoque.

## **Como Rodar o Projeto**

## **Rotas da API**

Aqui estão algumas rotas principais da API:

- **POST** `/login` - Realiza o login de um usuário (não implementado em profundidade no início).
- **GET** `/products` - Lista todos os produtos cadastrados.
- **POST** `/products` - Cadastra um novo produto.
- **PUT** `/products/:id` - Edita um produto existente.
- **DELETE** `/products/:id` - Exclui um produto.
- **GET** `/sales` - Exibe todas as vendas realizadas.
- **POST** `/sales` - Registra uma nova venda, atualizando o estoque automaticamente.
- **GET** `/stock` - Exibe o status do estoque de todos os produtos.
- **GET** `/reports/sales` - Gera relatórios de vendas (por dia, semana ou mês).
- **GET** `/reports/stock` - Exibe os produtos com estoque baixo.

## **Licença**

Este projeto está licenciado sob a MIT License - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
