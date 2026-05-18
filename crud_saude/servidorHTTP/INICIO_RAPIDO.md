# ⚡ INÍCIO RÁPIDO - 5 Minutos

## 🚀 Setup em 5 Passos

### 1️⃣ Configure o Banco
```bash
psql -U seu_usuario -d seu_banco -f init.sql
```

### 2️⃣ Inicie o Servidor
```bash
go run app/main.go
```

### 3️⃣ Acesse no Navegador
```
http://localhost:3000/patients
```

### 4️⃣ Crie um Paciente
- Clique em "+ Novo Paciente"
- Preencha os 6 campos
- Clique em "Criar Paciente"

### 5️⃣ Adicione um Laudo
- Clique em "Ver" no paciente
- Clique em "+ Novo Laudo"
- Preencha os dados
- Clique em "Criar Laudo"

---

## 📍 Rotas Principais

```
/patients                          → Lista de pacientes
/createPatient                     → Novo paciente
/patientDetail?id=1                → Detalhes + laudos
/updatePatient?id=1                → Editar paciente
/deletePatient?id=1                → Deletar paciente

/createMedicalReport?patientID=1   → Novo laudo
/editMedicalReport?reportID=1      → Editar laudo
/deleteMedicalReport?reportID=1    → Deletar laudo
```

---

## 🗄️ Banco de Dados

### Tabelas Criadas
- `patients` - Informações dos pacientes
- `medical_reports` - Laudos médicos

### Campos Principais
**Pacientes**: Nome, Email, CPF, Data Nasc., Telefone, Endereço  
**Laudos**: Médico, Título, Conteúdo, Data

---

## 📁 O Que Mudou

### ✨ 16 Arquivos Novos
- 2 Go handlers (pacientes + laudos)
- 2 Go utils (funções do BD)
- 6 páginas HTML (listagem + formulários)
- 3 arquivos CSS (estilos)
- 3 documentos de guia

### 📝 2 Arquivos Modificados
- `app/main.go` - Adicionadas 9 rotas
- `Readme.md` - Documentação atualizada

---

## 💡 Funcionalidades

| Ação | Status |
|------|--------|
| Listar Pacientes | ✅ |
| Criar Paciente | ✅ |
| Editar Paciente | ✅ |
| Deletar Paciente | ✅ |
| Ver Detalhes | ✅ |
| Criar Laudo | ✅ |
| Editar Laudo | ✅ |
| Deletar Laudo | ✅ |
| Validação | ✅ |
| Responsividade | ✅ |

---

## 📚 Documentação

```
Leia também:
├── SUMARIO_EXECUTIVO.md      - Visão geral
├── GUIA_PACIENTES.md          - Manual completo
├── ARQUITETURA.md             - Estrutura técnica
├── RESUMO_IMPLEMENTACOES.md   - Lista de mudanças
└── CHECKLIST_VERIFICACAO.md   - Verificação
```

---

## 🐛 Problemas Comuns

**"Erro ao conectar ao BD"**
→ Verifique `.env` e se PostgreSQL está rodando

**"Tabelas não encontradas"**
→ Execute `psql -U user -d db -f init.sql`

**"Erro: Email/CPF já existe"**
→ Use dados diferentes para teste

---

## ✅ Pronto!

Seu sistema de gestão de saúde está completo e funcionando. 🎉

**Próximo passo**: Acesse http://localhost:3000/patients

---

*Versão: 1.0 | Data: 18/05/2026*
