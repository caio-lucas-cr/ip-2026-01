package utils

import (
	"log"
)

type MedicalReport struct {
	ID        int    `json:"id"`
	PatientID int    `json:"patient_id"`
	Doctor    string `json:"doctor"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// Insere um novo laudo médico no banco de dados
func InsertMedicalReport(patientID int, doctor, title, content string) (int, error) {
	query := `INSERT INTO medical_reports (patient_id, doctor, title, content) 
			 VALUES ($1, $2, $3, $4) RETURNING id`
	var reportID int
	err := DB.QueryRow(query, patientID, doctor, title, content).Scan(&reportID)
	if err != nil {
		log.Printf("Erro ao inserir laudo médico no banco de dados: %v", err)
		return 0, err
	}
	log.Println("Laudo médico inserido com sucesso!")
	return reportID, nil
}

// Busca todos os laudos de um paciente específico
func GetMedicalReportsByPatientID(patientID int) ([]MedicalReport, error) {
	query := `SELECT id, patient_id, doctor, title, content, created_at FROM medical_reports WHERE patient_id = $1 ORDER BY created_at DESC`
	rows, err := DB.Query(query, patientID)
	if err != nil {
		log.Printf("Erro ao buscar laudos do paciente no banco de dados: %v", err)
		return nil, err
	}
	defer rows.Close()

	var reports []MedicalReport
	for rows.Next() {
		var report MedicalReport
		err := rows.Scan(&report.ID, &report.PatientID, &report.Doctor, &report.Title, &report.Content, &report.CreatedAt)
		if err != nil {
			log.Printf("Erro ao escanear laudo: %v", err)
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

// Busca um laudo específico pelo ID
func GetMedicalReportByID(id int) (*MedicalReport, error) {
	query := `SELECT id, patient_id, doctor, title, content, created_at FROM medical_reports WHERE id = $1`
	var report MedicalReport
	err := DB.QueryRow(query, id).Scan(&report.ID, &report.PatientID, &report.Doctor, &report.Title, &report.Content, &report.CreatedAt)
	if err != nil {
		log.Printf("Erro ao buscar laudo no banco de dados: %v", err)
		return nil, err
	}
	return &report, nil
}

// Atualiza um laudo médico
func UpdateMedicalReport(id int, doctor, title, content string) error {
	query := `UPDATE medical_reports SET doctor = $1, title = $2, content = $3 WHERE id = $4`
	_, err := DB.Exec(query, doctor, title, content, id)
	if err != nil {
		log.Printf("Erro ao atualizar laudo no banco de dados: %v", err)
		return err
	}
	log.Println("Laudo médico atualizado com sucesso!")
	return nil
}

// Deleta um laudo médico
func DeleteMedicalReport(id int) error {
	query := `DELETE FROM medical_reports WHERE id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao deletar laudo no banco de dados: %v", err)
		return err
	}
	log.Println("Laudo médico deletado com sucesso!")
	return nil
}
