package handlers

import (
	"net/http"
	"net/url"
	"text/template"

	"servidorHTTP/app/utils"
)

// ProfileHandler exibe a página de perfil do usuário usando o email fornecido na query string.
func ProfileHandler(response http.ResponseWriter, request *http.Request) {
	email := request.URL.Query().Get("email")
	if email == "" {
		http.Redirect(response, request, "/index.html", http.StatusSeeOther)
		return
	}

	decodedEmail, err := url.QueryUnescape(email)
	if err != nil {
		http.Redirect(response, request, "/index.html", http.StatusSeeOther)
		return
	}

	user, err := utils.GetUserByEmail(decodedEmail)
	if err != nil {
		http.Error(response, "Erro ao buscar informações do usuário", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, user)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}
