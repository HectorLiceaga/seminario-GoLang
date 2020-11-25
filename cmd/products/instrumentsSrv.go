package main

import (
	"flag"
	"fmt"
	"os"
	"seminario-GoLang/internal/config"
	"seminario-GoLang/internal/database"
	"seminario-GoLang/internal/service/instruments"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDataBase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := instruments.New(db, cfg)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is how you should config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS instruments (
		id integer primary key autoincrement,
		name text,
		description text,
		price integer);`

	// excecute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// this is another way using MustExec, which panics on error
	insertInstrument := `INSERT INTO instruments (name, description, price) VALUES (?,?,?)`
	i1 := fmt.Sprintf("Instrument %v", "GenericName")
	i2 := "bla bla bla"
	i3 := 0
	db.MustExec(insertInstrument, i1, i2, i3)
	return nil
}
