package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"servidorHTTP/app/utils"
)

// CreateMedicalReportHandler processa a criação de um novo laudo médico
func CreateMedicalReportHandler(response http.ResponseWriter, request *http.Request) {
	patientIDStr := request.URL.Query().Get("patientID")
	if patientIDStr == "" {
		http.Error(response, "ID do paciente não fornecido", http.StatusBadRequest)
		return
	}

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		http.Error(response, "ID do paciente inválido", http.StatusBadRequest)
		return
	}

	if request.Method == http.MethodGet {
		// Exibe o formulário
		data := map[string]interface{}{
			"PatientID": patientID,
		}
		tmpl, err := template.ParseFiles("static/forms/createMedicalReport.html")
		if err != nil {
			http.Error(response, "Erro ao carregar o formulário", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(response, data)
		return
	}

	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os dados do formulário
	doctor := request.FormValue("doctor")
	title := request.FormValue("title")
	content := request.FormValue("content")

	// Valida os dados
	if doctor == "" || title == "" || content == "" {
		http.Error(response, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	// Insere o laudo no banco de dados
	_, err = utils.InsertMedicalReport(patientID, doctor, title, content)
	if err != nil {
		http.Error(response, "Erro ao salvar o laudo", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página de detalhes do paciente
	http.Redirect(response, request, fmt.Sprintf("/patientDetail?id=%d", patientID), http.StatusSeeOther)
}

// EditMedicalReportHandler processa a edição de um laudo médico
func EditMedicalReportHandler(response http.ResponseWriter, request *http.Request) {
	reportIDStr := request.URL.Query().Get("reportID")
	if reportIDStr == "" {
		http.Error(response, "ID do laudo não fornecido", http.StatusBadRequest)
		return
	}

	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		http.Error(response, "ID do laudo inválido", http.StatusBadRequest)
		return
	}

	if request.Method == http.MethodGet {
		// Exibe o formulário com dados do laudo
		report, err := utils.GetMedicalReportByID(reportID)
		if err != nil {
			http.Error(response, "Laudo não encontrado", http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("static/forms/editMedicalReport.html")
		if err != nil {
			http.Error(response, "Erro ao carregar o formulário", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(response, report)
		return
	}

	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os dados do formulário
	doctor := request.FormValue("doctor")
	title := request.FormValue("title")
	content := request.FormValue("content")

	// Atualiza o laudo no banco de dados
	err = utils.UpdateMedicalReport(reportID, doctor, title, content)
	if err != nil {
		http.Error(response, "Erro ao atualizar o laudo", http.StatusInternalServerError)
		return
	}

	// Obtém o paciente para redirecionar corretamente
	report, err := utils.GetMedicalReportByID(reportID)
	if err != nil {
		http.Error(response, "Erro ao buscar laudo", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página de detalhes do paciente
	http.Redirect(response, request, fmt.Sprintf("/patientDetail?id=%d", report.PatientID), http.StatusSeeOther)
}

// DeleteMedicalReportHandler deleta um laudo médico
func DeleteMedicalReportHandler(response http.ResponseWriter, request *http.Request) {
	reportIDStr := request.URL.Query().Get("reportID")
	patientIDStr := request.URL.Query().Get("patientID")

	if reportIDStr == "" || patientIDStr == "" {
		http.Error(response, "IDs não fornecidos", http.StatusBadRequest)
		return
	}

	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		http.Error(response, "ID do laudo inválido", http.StatusBadRequest)
		return
	}

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		http.Error(response, "ID do paciente inválido", http.StatusBadRequest)
		return
	}

	err = utils.DeleteMedicalReport(reportID)
	if err != nil {
		http.Error(response, "Erro ao deletar o laudo", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página de detalhes do paciente
	http.Redirect(response, request, fmt.Sprintf("/patientDetail?id=%d", patientID), http.StatusSeeOther)
}
