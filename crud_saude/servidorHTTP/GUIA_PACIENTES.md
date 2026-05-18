# Guia de Uso - Sistema de Gestão de Pacientes e Laudos Médicos

## 📋 Índice
1. [Visão Geral](#visão-geral)
2. [Instalação e Configuração](#instalação-e-configuração)
3. [Como Usar](#como-usar)
4. [Fluxo de Trabalho](#fluxo-de-trabalho)
5. [Estrutura do Banco de Dados](#estrutura-do-banco-de-dados)
6. [Troubleshooting](#troubleshooting)

---

## 🏥 Visão Geral

Este sistema oferece um conjunto completo de funcionalidades para gerenciar pacientes e seus laudos médicos:

- **Cadastro de Pacientes**: Adicione novos pacientes com informações pessoais e de contato
- **Edição de Pacientes**: Atualize dados de pacientes existentes
- **Exclusão de Pacientes**: Remova pacientes do sistema
- **Gestão de Laudos**: Crie, edite e delete laudos médicos para cada paciente
- **Visualização**: Consulte todas as informações de forma organizada e intuitiva

---

## 🛠️ Instalação e Configuração

### Pré-requisitos
- Go 1.24.1 ou superior
- PostgreSQL 12 ou superior
- Docker e Docker Compose (opcional)

### Passo 1: Clonar o Repositório
```bash
git clone <URL_DO_REPOSITORIO>
cd servidorHTTP
```

### Passo 2: Configurar Variáveis de Ambiente
Crie um arquivo `.env` na raiz do projeto:
```env
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_banco
DB_HOST=localhost
DB_PORT=5432
```

### Passo 3: Iniciar o Banco de Dados
Opção A - Com Docker:
```bash
docker compose up -d
```

Opção B - PostgreSQL Local:
Certifique-se de que o PostgreSQL está em execução.

### Passo 4: Criar as Tabelas
Execute o script SQL:
```bash
psql -U seu_usuario -d nome_do_banco -f init.sql
```

### Passo 5: Instalar Dependências
```bash
go mod tidy
```

### Passo 6: Iniciar o Servidor
```bash
go run app/main.go
```

O servidor estará disponível em `http://localhost:3000`

---

## 📚 Como Usar

### Acessar o Sistema

1. Acesse `http://localhost:3000` no seu navegador
2. Clique em "Pacientes" na navegação principal

### Listar Pacientes

1. Na página principal, clique em "Pacientes"
2. Você verá uma tabela com todos os pacientes cadastrados
3. Para cada paciente, você pode:
   - **Ver**: Clique em "Ver" para acessar os detalhes
   - **Editar**: Clique em "Editar" para modificar os dados
   - **Deletar**: Clique em "Deletar" para remover (com confirmação)

### Criar Novo Paciente

1. Na página de pacientes, clique em "+ Novo Paciente"
2. Preencha os campos obrigatórios:
   - **Nome Completo**: Ex: João Silva
   - **Email**: Ex: joao@example.com
   - **CPF**: Ex: 123.456.789-00
   - **Data de Nascimento**: Selecione a data
   - **Telefone**: Ex: (11) 98765-4321
   - **Endereço**: Rua das Flores, 123...
3. Clique em "Criar Paciente"

### Editar Paciente

1. Na lista de pacientes, clique em "Editar" para o paciente desejado
2. Modifique os campos necessários
3. Clique em "Salvar Alterações"

### Deletar Paciente

1. Na lista de pacientes ou na página de detalhes, clique em "Deletar"
2. Confirme a exclusão quando solicitado
3. **Atenção**: Deletar um paciente também deleta todos os seus laudos

### Gerenciar Laudos de um Paciente

1. Na lista de pacientes, clique em "Ver" para abrir os detalhes
2. Role até a seção "📋 Laudos Médicos"
3. Clique em "+ Novo Laudo" para adicionar um novo laudo

### Criar Novo Laudo

1. Clique em "+ Novo Laudo" na página de detalhes do paciente
2. Preencha os campos:
   - **Médico Responsável**: Nome do médico
   - **Título do Laudo**: Ex: Consulta de Rotina
   - **Conteúdo do Laudo**: Descrição detalhada do diagnóstico e recomendações
3. Clique em "Criar Laudo"

### Editar Laudo

1. Na seção de laudos, clique em "Editar" para o laudo desejado
2. Modifique as informações
3. Clique em "Salvar Alterações"

### Deletar Laudo

1. Na seção de laudos, clique em "Deletar"
2. Confirme a exclusão
3. Você será redirecionado para a página de detalhes do paciente

---

## 🔄 Fluxo de Trabalho Típico

### Cenário 1: Cadastrar um Novo Paciente e Adicionar seu Primeiro Laudo

```
1. Acesar /patients
2. Clicar em "+ Novo Paciente"
3. Preencher formulário e criar
4. Clique em "Ver" no paciente recém-criado
5. Clique em "+ Novo Laudo"
6. Preencher dados do laudo
7. Laudo é criado e exibido na página
```

### Cenário 2: Atualizar Informações de um Paciente

```
1. Acesar /patients
2. Clicar em "Editar" para o paciente
3. Modificar campos necessários
4. Clicar em "Salvar Alterações"
5. Redirecionado para detalhes do paciente
```

### Cenário 3: Remover Paciente e Seus Laudos

```
1. Acesar /patients
2. Clicar em "Deletar" para o paciente
3. Confirmar exclusão
4. Paciente e todos seus laudos são removidos
```

---

## 🗄️ Estrutura do Banco de Dados

### Tabela: `patients`
Armazena informações dos pacientes.

| Campo | Tipo | Descrição |
|-------|------|-----------|
| id | SERIAL | Identificador único |
| name | VARCHAR(255) | Nome completo |
| email | VARCHAR(255) | E-mail (único) |
| cpf | VARCHAR(14) | CPF (único) |
| born_date | DATE | Data de nascimento |
| phone | VARCHAR(20) | Telefone |
| address | TEXT | Endereço completo |
| created_at | TIMESTAMP | Data de criação |
| updated_at | TIMESTAMP | Data da última atualização |

### Tabela: `medical_reports`
Armazena laudos médicos dos pacientes.

| Campo | Tipo | Descrição |
|-------|------|-----------|
| id | SERIAL | Identificador único |
| patient_id | INTEGER | ID do paciente (FK) |
| doctor | VARCHAR(255) | Nome do médico |
| title | VARCHAR(255) | Título do laudo |
| content | TEXT | Conteúdo do laudo |
| created_at | TIMESTAMP | Data de criação |
| updated_at | TIMESTAMP | Data da última atualização |

### Relacionamentos
- Um paciente pode ter múltiplos laudos (1:N)
- Ao deletar um paciente, todos seus laudos são deletados automaticamente (CASCADE)

---

## 🔍 Endpoints da API

### Pacientes
- `GET /patients` - Lista todos os pacientes
- `GET /patientDetail?id=<id>` - Detalhes de um paciente específico
- `GET /createPatient` - Formulário de criação
- `POST /createPatient` - Criar novo paciente
- `GET /updatePatient?id=<id>` - Formulário de edição
- `POST /updatePatient?id=<id>` - Atualizar paciente
- `GET /deletePatient?id=<id>` - Deletar paciente

### Laudos Médicos
- `GET /createMedicalReport?patientID=<id>` - Formulário de novo laudo
- `POST /createMedicalReport?patientID=<id>` - Criar laudo
- `GET /editMedicalReport?reportID=<id>` - Formulário de edição
- `POST /editMedicalReport?reportID=<id>` - Atualizar laudo
- `GET /deleteMedicalReport?reportID=<id>&patientID=<id>` - Deletar laudo

---

## 🐛 Troubleshooting

### "Erro ao conectar ao banco de dados"
- Verifique se o PostgreSQL está em execução
- Confirme as variáveis de ambiente no arquivo `.env`
- Verifique a conectividade de rede

### "Tabelas não encontradas"
- Execute o script `init.sql` para criar as tabelas
- Verifique se está usando o banco de dados correto

### "Email ou CPF já existem"
- O sistema verifica unicidade para email e CPF
- Use valores diferentes ou verifique registros existentes

### "Erro ao deletar paciente"
- Certifique-se de que o paciente existe
- Verifique permissões no banco de dados

### "Página não carrega"
- Verifique se o servidor está em execução
- Confirme a porta (padrão: 3000)
- Limpe cache do navegador (Ctrl+Shift+Delete)

---

## 📝 Exemplo de Dados de Teste

```sql
-- Inserir paciente de teste
INSERT INTO patients (name, email, cpf, born_date, phone, address)
VALUES ('Maria Silva', 'maria.silva@example.com', '123.456.789-00', '1990-05-15', '(11) 98765-4321', 'Rua das Flores, 123, São Paulo - SP');

-- Inserir laudo de teste
INSERT INTO medical_reports (patient_id, doctor, title, content)
VALUES (1, 'Dr. Carlos Santos', 'Consulta de Rotina', 'Paciente apresenta bom estado geral de saúde. Sem queixas. Pressão arterial e frequência cardíaca normais.');
```

---

## 💡 Dicas Úteis

1. **Backup Regular**: Realize backups periódicos do banco de dados
2. **Validação**: O sistema valida todos os campos obrigatórios
3. **Segurança**: Senhas são criptografadas antes de serem armazenadas
4. **Performance**: O banco tem índices nas colunas mais consultadas
5. **Auditoria**: Cada registro tem timestamps de criação e atualização

---

## 📞 Suporte

Para mais informações ou relatar problemas, entre em contato com o desenvolvedor.

---

**Versão**: 1.0  
**Última Atualização**: 2026-05-18
