package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario - Cria um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao se conectar ao banco de dados"))
		return
	}

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if erro != nil {
		w.Write([]byte("Falha ao criar o statement"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	//STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso! ID: %d", idInserido)))

}

// BuscarUsuarios traz todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao se conectar ao banco de dados"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario
		// SCAN escaneia na linha cada umdos campos da tabela 1 JOÃO JOÃO.PEDRO@GMAIL.COM, então preciso passar o endereço de memória de cada campo para serem populados
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil { //json.newencoder(w) - cria um enconder para o write e já envia os dados
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
	w.WriteHeader(http.StatusOK)

}

// BuscarUsuarios traz um usuário especifico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o ID para uint32"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao se conectar ao banco de dados"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)

	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	defer linha.Close()

	var usuario usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	} else {
		w.Write([]byte("Nenhum usuário encontrado para o ID especificado"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil { //json.newencoder(w) - cria um enconder para o write e já envia os dados
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}
}

// AtualizarUsuario altera os dados de um usuário no bando de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32) // base 10 - 32 bits

	if erro != nil {
		w.Write([]byte("Erro ao covnerter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o JSON para struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()
	// não interessa o número de linhas afetadas, apenas se deu erro ou não
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeletarUsuario deleta os dados de um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao conveter o ID para int"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
