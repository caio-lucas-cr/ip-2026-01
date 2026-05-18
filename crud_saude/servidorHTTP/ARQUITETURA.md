# Arquitetura do Sistema - CRUD de Pacientes e Laudos Médicos

## 🏗️ Diagrama de Arquitetura

```
┌─────────────────────────────────────────────────────────────────┐
│                         FRONTEND (HTML/CSS)                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐           │
│  │   Patients   │  │ Patient      │  │Create/Edit  │           │
│  │   List       │→ │   Detail     │→ │  Patient    │           │
│  │   (GET)      │  │   (GET)      │  │  (POST)     │           │
│  └──────────────┘  └──────────────┘  └──────────────┘           │
│                            │                                     │
│                            ↓                                     │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐           │
│  │Create/Edit  │  │Delete Medical│  │Medical       │           │
│  │  Medical    │→ │   Report     │→ │  Reports     │           │
│  │  Report     │  │   (POST)     │  │  List        │           │
│  │  (POST)     │  └──────────────┘  │  (GET)       │           │
│  └──────────────┘                    └──────────────┘           │
│                                                                   │
└────────────────┬──────────────────────────────────────────────┬─┘
                 │                                              │
                 ↓                                              ↓
        ┌────────────────┐                            ┌────────────────┐
        │  HTTP HANDLERS │                            │ GO BACKEND API │
        ├────────────────┤                            ├────────────────┤
        │                │                            │                │
        │ patientHandler │                            │ Patient Ops:   │
        │ ├─ List        │←──────────────────────────→│ ├─ Get All     │
        │ ├─ Detail      │                            │ ├─ Get By ID   │
        │ ├─ Create      │                            │ ├─ Create      │
        │ ├─ Update      │                            │ ├─ Update      │
        │ └─ Delete      │                            │ └─ Delete      │
        │                │                            │                │
        │ medicalReport  │                            │ Medical Ops:   │
        │ Handler        │                            │ ├─ Create      │
        │ ├─ Create      │←──────────────────────────→│ ├─ Get List    │
        │ ├─ Edit        │                            │ ├─ Get By ID   │
        │ └─ Delete      │                            │ ├─ Update      │
        │                │                            │ └─ Delete      │
        └────────────────┘                            └────────────────┘
                 │                                              │
                 └──────────────────────┬───────────────────────┘
                                        ↓
                          ┌─────────────────────────┐
                          │  PostgreSQL Database    │
                          ├─────────────────────────┤
                          │                         │
                          │  ┌──────────────────┐   │
                          │  │    patients      │   │
                          │  │  ├─ id (PK)      │   │
                          │  │  ├─ name         │   │
                          │  │  ├─ email (UQ)   │   │
                          │  │  ├─ cpf (UQ)     │   │
                          │  │  ├─ born_date    │   │
                          │  │  ├─ phone        │   │
                          │  │  ├─ address      │   │
                          │  │  ├─ created_at   │   │
                          │  │  └─ updated_at   │   │
                          │  └──────────────────┘   │
                          │                         │
                          │  ┌──────────────────┐   │
                          │  │ medical_reports  │   │
                          │  │  ├─ id (PK)      │   │
                          │  │  ├─ patient_id   │   │
                          │  │  │  (FK→patients)│   │
                          │  │  ├─ doctor       │   │
                          │  │  ├─ title        │   │
                          │  │  ├─ content      │   │
                          │  │  ├─ created_at   │   │
                          │  │  └─ updated_at   │   │
                          │  └──────────────────┘   │
                          │                         │
                          │  Relacionamento: 1:N    │
                          │  (1 Paciente : N Laudos)│
                          └─────────────────────────┘
```

## 📁 Estrutura de Arquivos

```
servidorHTTP/
│
├── app/
│   ├── main.go                          # Arquivo principal, configuração de rotas
│   │
│   ├── handlers/
│   │   ├── helloHandler.go              # Handler para /hello (original)
│   │   ├── formHandler.go               # Handler para criar usuário (original)
│   │   ├── loginHandler.go              # Handler para login (original)
│   │   ├── updateAccountHandler.go      # Handler para atualizar usuário (original)
│   │   ├── deleteAccountHandler.go      # Handler para deletar usuário (original)
│   │   ├── patientHandler.go            # ✨ NOVO: Handlers para pacientes
│   │   └── medicalReportHandler.go      # ✨ NOVO: Handlers para laudos médicos
│   │
│   └── utils/
│       ├── connectToDB.go               # Conexão ao banco de dados
│       ├── DB.go                        # Funções de usuários
│       ├── encrypt.go                   # Criptografia de senhas
│       ├── validateUser.go              # Validação de usuários
│       ├── getUserByEmail.go            # Query de usuários
│       ├── updateUser.go                # Atualização de usuários
│       ├── deleteUser.go                # Deleção de usuários
│       ├── patient.go                   # ✨ NOVO: Funções para pacientes
│       └── medicalReport.go             # ✨ NOVO: Funções para laudos
│
├── static/
│   ├── index.html                       # Página inicial (original)
│   ├── profile.html                     # Perfil de usuário (original)
│   ├── patients.html                    # ✨ NOVO: Lista de pacientes
│   ├── patientDetail.html               # ✨ NOVO: Detalhes do paciente
│   │
│   ├── forms/
│   │   ├── createAccount.html           # Criar conta (original)
│   │   ├── deleteAccount.html           # Deletar conta (original)
│   │   ├── login.html                   # Login (original)
│   │   ├── updateAccount.html           # Atualizar conta (original)
│   │   ├── createPatient.html           # ✨ NOVO: Criar paciente
│   │   ├── editPatient.html             # ✨ NOVO: Editar paciente
│   │   ├── createMedicalReport.html     # ✨ NOVO: Criar laudo
│   │   └── editMedicalReport.html       # ✨ NOVO: Editar laudo
│   │
│   └── styles/
│       ├── index.style.css              # Estilo inicial (original)
│       ├── createAccount.style.css      # Estilo formulário (original)
│       ├── deleteAccount.style.css      # Estilo formulário (original)
│       ├── login.style.css              # Estilo login (original)
│       ├── profile.style.css            # Estilo perfil (original)
│       ├── updateAccount.style.css      # Estilo formulário (original)
│       ├── patients.style.css           # ✨ NOVO: Estilo lista pacientes
│       ├── patientDetail.style.css      # ✨ NOVO: Estilo detalhes
│       └── createPatient.style.css      # ✨ NOVO: Estilo formulários
│
├── .env                                 # Variáveis de ambiente
├── .gitignore                           # Arquivos ignorados pelo git
├── docker-compose.yml                   # Configuração Docker
├── docker-compose.yml.example           # Exemplo Docker
├── go.mod                               # Dependências Go
├── Readme.md                            # Documentação principal
├── GUIA_PACIENTES.md                    # ✨ NOVO: Guia de uso do sistema
└── init.sql                             # ✨ NOVO: Script para criar tabelas
```

## 🔄 Fluxo de Dados

### Criar Paciente
```
Usuário preenche formulário em /createPatient
    ↓
form.html (POST) → CreatePatientHandler
    ↓
Valida dados
    ↓
utils.InsertPatient() → PostgreSQL
    ↓
Paciente armazenado em patients
    ↓
Redireciona para /patients (lista atualizada)
```

### Visualizar Detalhes do Paciente
```
Usuário clica "Ver" na lista
    ↓
/patientDetail?id=X → PatientDetailHandler
    ↓
GetPatientByID(X) → PostgreSQL
GetMedicalReportsByPatientID(X) → PostgreSQL
    ↓
Dados reunidos em map
    ↓
Renderiza patientDetail.html com dados
    ↓
Exibe paciente + seus laudos
```

### Criar Laudo
```
Usuário em detalhes do paciente clica "+ Novo Laudo"
    ↓
/createMedicalReport?patientID=X → form exibido
    ↓
Usuário preenche e submete
    ↓
CreateMedicalReportHandler (POST)
    ↓
InsertMedicalReport(patientID, doctor, title, content)
    ↓
Laudo inserido em medical_reports
    ↓
Redireciona para /patientDetail?id=X (com novo laudo)
```

## 🛠️ Tecnologias Utilizadas

| Camada | Tecnologia | Função |
|--------|-----------|--------|
| Frontend | HTML5, CSS3 | Interface do usuário |
| Backend | Go 1.24.1 | Servidor HTTP e lógica |
| Banco de Dados | PostgreSQL 12+ | Armazenamento de dados |
| Driver BD | github.com/lib/pq | Conexão PostgreSQL |
| Env | github.com/joho/godotenv | Variáveis de ambiente |
| HTTP | net/http (stdlib) | Framework HTTP nativo |
| Templates | text/template (stdlib) | Renderização de HTML |

## 📊 Relacionamento entre Tabelas

```
PATIENTS (1)
    ↓
    ├── id (PK)
    ├── name
    ├── email (UNIQUE)
    ├── cpf (UNIQUE)
    ├── born_date
    ├── phone
    ├── address
    ├── created_at
    └── updated_at
    
    ↕ (1:N)
    
MEDICAL_REPORTS (N)
    ├── id (PK)
    ├── patient_id (FK) ← referencia patients.id
    ├── doctor
    ├── title
    ├── content
    ├── created_at
    └── updated_at

Constraint: ON DELETE CASCADE
(Deletar um paciente deleta automaticamente seus laudos)
```

## 🚀 Fluxo de Requisição Completo

```
1. Navegador do Usuário
   └─ GET http://localhost:3000/patients

2. Go Server recebe requisição
   └─ Match com handler → ListPatientsHandler

3. Handler executa
   ├─ Chama utils.GetAllPatients()
   ├─ Conecta ao PostgreSQL
   ├─ Executa SELECT * FROM patients
   ├─ Retorna []Patient
   └─ Recebe lista de pacientes

4. Renderizar Template
   ├─ Carrega patients.html
   ├─ Itera sobre patients []
   ├─ Gera HTML dinamicamente
   └─ Aplica CSS de patients.style.css

5. Resposta ao Navegador
   └─ HTTP 200 OK com HTML renderizado

6. Navegador Renderiza
   ├─ Exibe navbar
   ├─ Exibe tabela com pacientes
   ├─ Aplica estilos CSS
   └─ Exibe ao usuário
```

---

**Diagrama atualizado em**: 2026-05-18
