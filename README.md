# TechTask

A empresa "TechTask" deseja desenvolver um sistema de gerenciamento de tarefas para ajudar suas equipes a acompanhar e gerenciar suas atividades diárias de forma eficiente. O sistema deve permitir que os usuários criem, atualizem, removam e visualizem tarefas.

**Título do Projeto:** Sistema de Gerenciamento de Tarefas

**Descrição:** A empresa "TechTask" deseja desenvolver um sistema de gerenciamento de tarefas para ajudar suas equipes a acompanhar e gerenciar suas atividades diárias de forma eficiente. O sistema deve permitir que os usuários criem, atualizem, removam e visualizem tarefas.

**Requisitos Essenciais:**

1. Autenticação de Usuários: Os usuários devem ser capazes de se autenticar na API para acessar as funcionalidades do sistema.
2. CRUD de Tarefas: Os usuários devem ser capazes de criar, ler, atualizar e excluir tarefas.
Validação de Dados: A API deve validar os dados recebidos para garantir que estão corretos e consistentes.
3. Persistência de Dados: As informações das tarefas devem ser armazenadas em um banco de dados para garantir a persistência dos dados.
4. Endpoints RESTful: A API deve seguir os princípios RESTful para interação com os clientes.

**Estrutura de Dados:**

**Usuário:**
* ID (int)
* Nome (string)
* Email (string)
* Senha (string)

**Tarefa:**
* ID (int)
* UserID (int) - Referência ao ID do usuário que criou a tarefa
* Título (string)
* Descrição (string)
* Data de Criação (timestamp)
* Data de Conclusão (timestamp)
* Status (string: pendente, em andamento, concluída)

**Endpoints:**

1. Autenticação de Usuário:
    * POST /api/login: Autentica um usuário e retorna um token de acesso.
    * POST /api/register: Registra um novo usuário na plataforma.
2. CRUD de Tarefas:
    * GET /api/tasks: Retorna todas as tarefas do usuário autenticado.
    * GET /api/tasks/{id}: Retorna os detalhes de uma tarefa específica.
    * POST /api/tasks: Cria uma nova tarefa.
    * PUT /api/tasks/{id}: Atualiza uma tarefa existente.
    * DELETE /api/tasks/{id}: Remove uma tarefa existente.

## Como Rodar o Projeto

A maior parte das configurações do projeto estão no arquivo ``.env``. Para rodar o projeto bastar subir o mysql e o server.

Subir mysql com docker:
```cmd
docker run --name my-mysql -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_DATABASE=db_techtask -d mysql:latest
```

Subir o mysql com docker-compose:
```
docker-compose up
```

Subir o server:
```cmd
go run main.go
```

OBS: Para acessar o container do mysql rodar o seguinte comando:
```
docker exec -it my-mysql bash
```
