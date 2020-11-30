package instruments

import (
	"database/sql"
	"seminario-GoLang/internal/config"

	"github.com/jmoiron/sqlx"
)

// Instrument ...
type Instrument struct {
	ID          int    `json:"id,omitempty" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int32  `json:"price" binding:"required"`
}

// Service ...
type Service interface {
	AddInstrument(*Instrument) error
	FindByID(int) *Instrument
	FindAll() []*Instrument
	Delete(int) *sql.Result
	Edit(*Instrument) *sql.Result
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// AddInstrument ...
func (s service) AddInstrument(i *Instrument) error {
	//fmt.Println(i)
	insertInstrument := `INSERT INTO instruments (name, description, price) VALUES (?,?,?)`
	s.db.MustExec(insertInstrument, i.Name, i.Description, i.Price)
	return nil
}

// FindByID ...
func (s service) FindByID(ID int) *Instrument {
	i := Instrument{}
	getInstrumentByID := `SELECT * FROM instruments WHERE id=(?)`
	err := s.db.Get(&i, getInstrumentByID, ID)
	if err != nil {
		panic(err)
	}
	return &i
}

// FindAll ...
func (s service) FindAll() []*Instrument {
	var list []*Instrument
	if err := s.db.Select(&list, "SELECT * FROM instruments"); err != nil {
		panic(err)
	}
	return list
}

// Delete ...
func (s service) Delete(ID int) *sql.Result {
	q := `DELETE FROM instruments WHERE id=(?)`
	result := s.db.MustExec(q, ID)
	return &result
}

// Edit ...
func (s service) Edit(i *Instrument) *sql.Result {
	//fmt.Println(i)
	insertInstrument := `UPDATE instruments SET name=(?), description=(?), price=(?) WHERE id=(?)`
	result := s.db.MustExec(insertInstrument, i.Name, i.Description, i.Price, i.ID)
	return &result
}
