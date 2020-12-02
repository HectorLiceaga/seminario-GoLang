package main

import (
	"flag"
	"fmt"
	"os"
	"seminario-GoLang/internal/config"
	"seminario-GoLang/internal/database"
	"seminario-GoLang/internal/service/instruments"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDataBase(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	createSchema(db)
	service, _ := instruments.New(db, cfg)
	httpService := instruments.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
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
