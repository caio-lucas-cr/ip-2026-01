package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"servidorHTTP/app/utils"
)

// ListPatientsHandler exibe a lista de todos os pacientes
func ListPatientsHandler(response http.ResponseWriter, request *http.Request) {
	patients, err := utils.GetAllPatients()
	if err != nil {
		http.Error(response, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/patients.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, patients)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}

// PatientDetailHandler exibe os detalhes de um paciente específico e seus laudos
func PatientDetailHandler(response http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	if idStr == "" {
		http.Error(response, "ID do paciente não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID do paciente inválido", http.StatusBadRequest)
		return
	}

	patient, err := utils.GetPatientByID(id)
	if err != nil {
		http.Error(response, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	reports, err := utils.GetMedicalReportsByPatientID(id)
	if err != nil {
		http.Error(response, "Erro ao buscar laudos", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Patient": patient,
		"Reports": reports,
	}

	tmpl, err := template.ParseFiles("static/patientDetail.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, data)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}

// CreatePatientHandler processa o formulário de criação de paciente
func CreatePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		// Exibe o formulário
		tmpl, err := template.ParseFiles("static/forms/createPatient.html")
		if err != nil {
			http.Error(response, "Erro ao carregar o formulário", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(response, nil)
		return
	}

	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os dados do formulário
	name := request.FormValue("name")
	email := request.FormValue("email")
	cpf := request.FormValue("cpf")
	bornDate := request.FormValue("bornDate")
	phone := request.FormValue("phone")
	address := request.FormValue("address")

	// Valida os dados
	if name == "" || email == "" || cpf == "" || bornDate == "" || phone == "" || address == "" {
		http.Error(response, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	// Insere o paciente no banco de dados
	_, err := utils.InsertPatient(name, email, cpf, bornDate, phone, address)
	if err != nil {
		http.Error(response, "Erro ao salvar o paciente", http.StatusInternalServerError)
		return
	}

	// Redireciona para a lista de pacientes
	http.Redirect(response, request, "/patients", http.StatusSeeOther)
}

// UpdatePatientHandler processa a atualização de um paciente
func UpdatePatientHandler(response http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	if idStr == "" {
		http.Error(response, "ID do paciente não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID do paciente inválido", http.StatusBadRequest)
		return
	}

	if request.Method == http.MethodGet {
		// Exibe o formulário com dados do paciente
		patient, err := utils.GetPatientByID(id)
		if err != nil {
			http.Error(response, "Paciente não encontrado", http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("static/forms/editPatient.html")
		if err != nil {
			http.Error(response, "Erro ao carregar o formulário", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(response, patient)
		return
	}

	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os dados do formulário
	name := request.FormValue("name")
	email := request.FormValue("email")
	cpf := request.FormValue("cpf")
	bornDate := request.FormValue("bornDate")
	phone := request.FormValue("phone")
	address := request.FormValue("address")

	// Atualiza o paciente no banco de dados
	err = utils.UpdatePatient(id, name, email, cpf, bornDate, phone, address)
	if err != nil {
		http.Error(response, "Erro ao atualizar o paciente", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página de detalhes do paciente
	http.Redirect(response, request, fmt.Sprintf("/patientDetail?id=%d", id), http.StatusSeeOther)
}

// DeletePatientHandler deleta um paciente
func DeletePatientHandler(response http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	if idStr == "" {
		http.Error(response, "ID do paciente não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID do paciente inválido", http.StatusBadRequest)
		return
	}

	err = utils.DeletePatient(id)
	if err != nil {
		http.Error(response, "Erro ao deletar o paciente", http.StatusInternalServerError)
		return
	}

	// Redireciona para a lista de pacientes
	http.Redirect(response, request, "/patients", http.StatusSeeOther)
}
