-- Criação da tabela de usuários
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    born_date DATE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Criação da tabela de pacientes
CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    born_date DATE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Criação da tabela de laudos médicos
CREATE TABLE IF NOT EXISTS medical_reports (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    doctor VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Criação de índices para melhorar performance
CREATE INDEX IF NOT EXISTS idx_patients_email ON patients(email);
CREATE INDEX IF NOT EXISTS idx_patients_cpf ON patients(cpf);
CREATE INDEX IF NOT EXISTS idx_medical_reports_patient_id ON medical_reports(patient_id);

-- Comentários sobre as tabelas
COMMENT ON TABLE patients IS 'Tabela que armazena informações dos pacientes';
COMMENT ON TABLE medical_reports IS 'Tabela que armazena os laudos médicos dos pacientes';
COMMENT ON COLUMN patients.id IS 'Identificador único do paciente';
COMMENT ON COLUMN patients.name IS 'Nome completo do paciente';
COMMENT ON COLUMN patients.email IS 'E-mail do paciente';
COMMENT ON COLUMN patients.cpf IS 'CPF do paciente (formato: XXX.XXX.XXX-XX)';
COMMENT ON COLUMN patients.born_date IS 'Data de nascimento do paciente';
COMMENT ON COLUMN patients.phone IS 'Telefone do paciente';
COMMENT ON COLUMN patients.address IS 'Endereço completo do paciente';
COMMENT ON COLUMN medical_reports.id IS 'Identificador único do laudo';
COMMENT ON COLUMN medical_reports.patient_id IS 'Referência ao paciente proprietário do laudo';
COMMENT ON COLUMN medical_reports.doctor IS 'Nome do médico que realizou o laudo';
COMMENT ON COLUMN medical_reports.title IS 'Título/tipo do laudo médico';
COMMENT ON COLUMN medical_reports.content IS 'Conteúdo detalhado do laudo médico';
