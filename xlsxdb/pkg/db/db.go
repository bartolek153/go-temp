package db

import (
	"database/sql"
	mssql "github.com/denisenkom/go-mssqldb"
)

type DbConnection struct[] {

}

func openConection(connStr string) {
	db, err := sql.Open("mssql", connStr)
	
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
		return false
	}
}

func testConnection(connStr string) bool {
	
	// defer db.Close()
	
	err = db.Ping()
	
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
		return false
	}

	return true
}


