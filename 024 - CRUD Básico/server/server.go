package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	Id    uint16 `json:"id"`
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyReq, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if error = json.Unmarshal(bodyReq, &usuario); error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao converter o usuário para struct!"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	//Prepare statement
	statement, error := db.Prepare("INSERT INTO usuario (nome, senha, email) VALUES (?, ?, ?)")
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao criar o statement!"))
		return
	}
	defer statement.Close()

	insert, error := statement.Exec(usuario.Nome, usuario.Senha, usuario.Email)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao executar o statement!"))
		return
	}

	idInserted, error := insert.LastInsertId()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao resgatar o id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso. ID: %d", idInserted)))
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connection()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	lines, error := db.Query("SELECT * FROM usuario")
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao buscar usuários!"))
		return
	}
	defer lines.Close()

	var usuarios []usuario
	for lines.Next() {
		var usuario usuario

		if error := lines.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha, &usuario.Email); error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Falha ao escanear os usuários"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(usuarios); error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao converter os usuários para json"))
		return
	}
}

func MostrarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, error := strconv.ParseUint(params["id"], 10, 32)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao converter os parâmetro para inteiro"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	line, error := db.Query("SELECT * FROM usuario WHERE id = ?", ID)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao buscar usuário!"))
		return
	}
	defer line.Close()

	var usuario usuario
	if line.Next() {
		if error := line.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha, &usuario.Email); error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Falha ao escanear os usuários"))
			return
		}
	}

	if usuario.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(usuario); error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao converter o usuário para json"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao converter o parâmetro para inteiro"))
		return
	}

	bodyReq, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if error = json.Unmarshal(bodyReq, &usuario); error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao converter o usuário para struct!"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("UPDATE usuario SET nome = ?, senha = ?, email = ? WHERE id = ?")
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(usuario.Nome, usuario.Senha, usuario.Email, ID); error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao executar o statement!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao converter o parâmetro para inteiro"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Falha ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("DELETE FROM usuario WHERE id = ?")
	if error != nil {
		w.Write([]byte("Falha ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		w.Write([]byte("Falha ao deletar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
