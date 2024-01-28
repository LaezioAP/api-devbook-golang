package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	database, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = database.Ping(); erro != nil {
		database.Close()
		fmt.Println("Acesso negado ao banco!")
		return nil, erro
	}

	return database, nil
}
