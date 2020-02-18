package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var HitsDatbase Context

type Context struct {
	DB   *sql.DB
	Stmt map[string]*sql.Stmt
}

func createHitsContext(stop chan int) {
	db, err := createSQLite()
	defer db.Close()
	if err != nil {
		panic("Fatal error database: " + err.Error())
	}
	HitsDatbase = Context{
		DB:   db,
		Stmt: map[string]*sql.Stmt{},
	}
	if err := createStmt(); err != nil {
		panic("Fatal error database: " + err.Error())
	}
	stop <- 1
	<-stop
}

func createSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	db.Exec(
		`
		CREATE TABLE IF NOT EXISTS IP(
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
			IPAddress TEXT NOT NULL,
			Chapter_ID INTEGER NOT NULL
		  );
		  
		  CREATE TABLE IF NOT EXISTS Novel(
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
			UpdateTime DATETIME,
			Name TEXT NOT NULL UNIQUE
		  );
		
		  CREATE TABLE IF NOT EXISTS Novel_Chapter(
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			CreateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
			UpdateTime DATETIME,
			Novel_ID INTEGER NOT NULL,
			Episode INTEGER NOT NULL,
			Chapter INTEGER NOT NULL,
			Name TEXT NOT NULL UNIQUE
		  );

		`,
	)
	return db, nil
}
