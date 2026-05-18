package handlers

// Importa os pacotes necessários para o funcionamento do handler
import (
	"net/http" // Usado para lidar com requisições e respostas HTTP
	"net/url"  // Usado para escapar parâmetros de consulta

	"servidorHTTP/app/utils" // Importa funções utilitárias, como validação de usuário e criptografia
)

// LoginHandler é responsável por processar os dados enviados pelo formulário de login
func LoginHandler(response http.ResponseWriter, request *http.Request) {

	// Verifica se o método da requisição é POST
	if request.Method != http.MethodPost {
		// Retorna um erro caso o método não seja suportado
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os valores enviados pelo formulário
	email := request.FormValue("email")       // E-mail do usuário
	password := request.FormValue("password") // Senha do usuário

	// Criptografa a senha para compará-la com a armazenada no banco de dados
	encryptedPassword := utils.Encrypt(password)

	// Verifica se o usuário existe no banco de dados
	isValidUser, err := utils.ValidateUser(email, encryptedPassword)
	if err != nil {
		// Retorna um erro caso ocorra falha ao validar o usuário
		http.Error(response, "Erro ao validar o usuário", http.StatusInternalServerError)
		return
	}

	// Verifica se as credenciais são inválidas
	if !isValidUser {
		// Retorna um erro caso as credenciais estejam incorretas
		http.Error(response, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Busca as informações do usuário no banco de dados para garantir que o email existe
	_, err = utils.GetUserByEmail(email)
	if err != nil {
		http.Error(response, "Erro ao buscar informações do usuário", http.StatusInternalServerError)
		return
	}

	// Redireciona o usuário para a página de perfil após o sucesso
	http.Redirect(response, request, "/profile.html?email="+url.QueryEscape(email), http.StatusSeeOther)
}
