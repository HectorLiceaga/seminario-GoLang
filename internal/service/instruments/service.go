package instruments

import (
	"seminario-GoLang/internal/config"

	"github.com/jmoiron/sqlx"
)

// Instrument ...
type Instrument struct {
	ID          int64
	Name        string
	Description string
	Price       int32
}

// InstrumentService ...
type InstrumentService interface {
	AddInstrument(Instrument) error
	FindByID(int) *Instrument
	FindAll() []*Instrument
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (InstrumentService, error) {
	return service{db, c}, nil
}

func (s service) AddInstrument(i Instrument) error {
	return nil
}

func (s service) FindByID(ID int) *Instrument {
	return nil
}

func (s service) FindAll() []*Instrument {
	var list []*Instrument
	if err := s.db.Select(&list, "SELECT * FROM instruments"); err != nil {
		panic(err)
	}
	return list
}
