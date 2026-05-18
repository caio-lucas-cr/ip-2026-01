package utils

import (
	"log"
)

type Patient struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	CPF          string `json:"cpf"`
	BornDate     string `json:"born_date"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	CreatedAt    string `json:"created_at"`
}

// Insere um novo paciente no banco de dados
func InsertPatient(name, email, cpf, bornDate, phone, address string) (int, error) {
	query := `INSERT INTO patients (name, email, cpf, born_date, phone, address) 
			 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var patientID int
	err := DB.QueryRow(query, name, email, cpf, bornDate, phone, address).Scan(&patientID)
	if err != nil {
		log.Printf("Erro ao inserir paciente no banco de dados: %v", err)
		return 0, err
	}
	log.Println("Paciente inserido com sucesso!")
	return patientID, nil
}

// Busca todos os pacientes no banco de dados
func GetAllPatients() ([]Patient, error) {
	query := `SELECT id, name, email, cpf, born_date, phone, address, created_at FROM patients ORDER BY created_at DESC`
	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("Erro ao buscar pacientes no banco de dados: %v", err)
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var patient Patient
		err := rows.Scan(&patient.ID, &patient.Name, &patient.Email, &patient.CPF, &patient.BornDate, &patient.Phone, &patient.Address, &patient.CreatedAt)
		if err != nil {
			log.Printf("Erro ao escanear paciente: %v", err)
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

// Busca um paciente específico pelo ID
func GetPatientByID(id int) (*Patient, error) {
	query := `SELECT id, name, email, cpf, born_date, phone, address, created_at FROM patients WHERE id = $1`
	var patient Patient
	err := DB.QueryRow(query, id).Scan(&patient.ID, &patient.Name, &patient.Email, &patient.CPF, &patient.BornDate, &patient.Phone, &patient.Address, &patient.CreatedAt)
	if err != nil {
		log.Printf("Erro ao buscar paciente no banco de dados: %v", err)
		return nil, err
	}
	return &patient, nil
}

// Atualiza os dados de um paciente
func UpdatePatient(id int, name, email, cpf, bornDate, phone, address string) error {
	query := `UPDATE patients SET name = $1, email = $2, cpf = $3, born_date = $4, phone = $5, address = $6 WHERE id = $7`
	_, err := DB.Exec(query, name, email, cpf, bornDate, phone, address, id)
	if err != nil {
		log.Printf("Erro ao atualizar paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente atualizado com sucesso!")
	return nil
}

// Deleta um paciente e seus laudos associados
func DeletePatient(id int) error {
	// Primeiro deleta todos os laudos associados ao paciente
	deleteReportsQuery := `DELETE FROM medical_reports WHERE patient_id = $1`
	_, err := DB.Exec(deleteReportsQuery, id)
	if err != nil {
		log.Printf("Erro ao deletar laudos do paciente: %v", err)
		return err
	}

	// Depois deleta o paciente
	query := `DELETE FROM patients WHERE id = $1`
	_, err = DB.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao deletar paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente deletado com sucesso!")
	return nil
}
