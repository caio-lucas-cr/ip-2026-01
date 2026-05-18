# Servidor HTTP com GO - Sistema de Gestão de Saúde

## Visão Geral
Este projeto é um servidor HTTP desenvolvido em GoLang que permite a criação, autenticação, atualização e exclusão de contas de usuários, além de um **sistema completo de gestão de pacientes e laudos médicos**. Ele utiliza PostgreSQL como banco de dados e fornece uma interface web moderna para interação com os usuários.

---

## Estrutura do Projeto
### Diretórios Principais:
- **`app/`**: Contém a lógica do servidor, incluindo handlers e utilitários.
  - **`handlers/`**: Lida com as requisições HTTP e processa os dados.
  - **`utils/`**: Contém funções auxiliares, como conexão ao banco de dados e manipulação de dados.
- **`static/`**: Contém os arquivos estáticos (HTML e CSS) que compõem o front-end da aplicação.
  - **`forms/`**: Formulários HTML para criar, atualizar, excluir contas e fazer login.
  - **`styles/`**: Arquivos CSS para estilização das páginas.

---

## Configuração do Ambiente

### Pré-requisitos
1. **GoLang**: Certifique-se de que o Go está instalado na sua máquina.
2. **PostgreSQL**: Banco de dados utilizado pelo projeto.
3. **Docker** (opcional): Para facilitar a configuração do banco de dados.

### Passos para Configuração
1. **Clone o repositório**:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd servidorHTTP
   ```

2. **Configure o arquivo `.env`**:
   Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```
   DB_USER=<seu_usuario>
   DB_PASSWORD=<sua_senha>
   DB_NAME=<nome_do_banco>
   DB_HOST=<host_do_banco>
   DB_PORT=<porta_do_banco>
   ```

3. **Configuração do Banco de Dados**:
   - Se estiver usando Docker, utilize o arquivo `docker-compose.yml` para subir o banco de dados:
     ```bash
     docker compose up
     ```
   - Execute o script `init.sql` para criar todas as tabelas necessárias:
     ```bash
     psql -U <seu_usuario> -d <nome_do_banco> -f init.sql
     ```
   
   Alternativamente, você pode criar manualmente as tabelas executando os comandos SQL no arquivo `init.sql`. Este arquivo cria:
   - **Tabela `users`**: Para gerenciar contas de usuários
   - **Tabela `patients`**: Para armazenar informações dos pacientes
   - **Tabela `medical_reports`**: Para armazenar os laudos médicos dos pacientes

4. **Instale as dependências**:
   Execute o comando abaixo para instalar as dependências do projeto:
   ```bash
   go mod tidy
   ```

---

## Executando o Projeto
1. **Inicie o servidor**:
   ```bash
   go run app/main.go
   ```

2. **Acesse a aplicação**:
   O servidor estará disponível no endereço exibido no terminal, geralmente algo como:
   ```
   http://127.0.0.1:3000/
   ```

---

## Funcionalidades

### 👥 Gerenciamento de Usuários
- **`/`**: Página inicial
- **`/form`**: Criar nova conta
- **`/login`**: Login de usuário
- **`/updateAccount`**: Atualizar dados da conta
- **`/deleteAccount`**: Deletar conta

### 🏥 CRUD de Pacientes
- **`/patients`**: Listar todos os pacientes
- **`/patientDetail?id=<id>`**: Ver detalhes de um paciente específico
- **`/createPatient`**: Criar novo paciente (GET: exibe formulário, POST: salva dados)
- **`/updatePatient?id=<id>`**: Editar dados do paciente (GET: exibe formulário, POST: salva alterações)
- **`/deletePatient?id=<id>`**: Deletar um paciente

### 📋 Gestão de Laudos Médicos
- **`/createMedicalReport?patientID=<id>`**: Criar novo laudo para um paciente (GET: formulário, POST: salva)
- **`/editMedicalReport?reportID=<id>`**: Editar um laudo médico existente (GET: formulário, POST: salva)
- **`/deleteMedicalReport?reportID=<id>&patientID=<id>`**: Deletar um laudo médico

### Handlers Disponíveis
- **User Management**: `FormHandler`, `LoginHandler`, `UpdateAccountHandler`, `DeleteAccountHandler`
- **Patient Management**: `ListPatientsHandler`, `PatientDetailHandler`, `CreatePatientHandler`, `UpdatePatientHandler`, `DeletePatientHandler`
- **Medical Reports**: `CreateMedicalReportHandler`, `EditMedicalReportHandler`, `DeleteMedicalReportHandler`

---

## Estrutura de Banco de Dados

### Tabela `patients`
Armazena informações dos pacientes:
```sql
- id (Serial, Primary Key)
- name (VARCHAR 255)
- email (VARCHAR 255, Unique)
- cpf (VARCHAR 14, Unique)
- born_date (DATE)
- phone (VARCHAR 20)
- address (TEXT)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

### Tabela `medical_reports`
Armazena laudos médicos associados aos pacientes:
```sql
- id (Serial, Primary Key)
- patient_id (Integer, Foreign Key referenciando patients.id)
- doctor (VARCHAR 255)
- title (VARCHAR 255)
- content (TEXT)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

---

## Estrutura de Pastas
```
.env
.gitignore
docker-compose.yml
go.mod
init.sql
Readme.md
app/
  main.go
  handlers/
    helloHandler.go
    formHandler.go
    loginHandler.go
    updateAccountHandler.go
    deleteAccountHandler.go
    patientHandler.go
    medicalReportHandler.go
  utils/
    connectToDB.go
    DB.go
    encrypt.go
    validateUser.go
    getUserByEmail.go
    updateUser.go
    deleteUser.go
    patient.go
    medicalReport.go
static/
  index.html
  profile.html
  patients.html
  patientDetail.html
  forms/
    createAccount.html
    deleteAccount.html
    login.html
    updateAccount.html
    createPatient.html
    editPatient.html
    createMedicalReport.html
    editMedicalReport.html
  styles/
    index.style.css
    createAccount.style.css
    deleteAccount.style.css
    login.style.css
    profile.style.css
    updateAccount.style.css
    patients.style.css
    patientDetail.style.css
    createPatient.style.css
```

---

## Observações
- Certifique-se de que o banco de dados está configurado corretamente antes de iniciar o servidor.
- O projeto utiliza o driver `github.com/lib/pq` para conexão com o PostgreSQL.
- Para mais informações, consulte os comentários no código ou entre em contato com o desenvolvedor.
