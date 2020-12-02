package instruments

import (
	"seminario-GoLang/internal/config"

	"github.com/jmoiron/sqlx"
)

// Instrument ...
type Instrument struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int32  `json:"price" binding:"required"`
}

// Service ...
type Service interface {
	AddInstrument(*Instrument) (int64, error)
	FindByID(int64) (*Instrument, error)
	FindAll() ([]*Instrument, error)
	Delete(int64) error
	Edit(*Instrument) error
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
func (s service) AddInstrument(i *Instrument) (int64, error) {
	insertInstrument := `INSERT INTO instruments (name, description, price) VALUES (?,?,?)`
	r := s.db.MustExec(insertInstrument, i.Name, i.Description, i.Price)
	ID, err := r.LastInsertId()
	return ID, err
}

// FindByID ...
func (s service) FindByID(ID int64) (*Instrument, error) {
	i := Instrument{}
	getInstrumentByID := `SELECT * FROM instruments WHERE id=(?)`
	err := s.db.Get(&i, getInstrumentByID, ID)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

// FindAll ...
func (s service) FindAll() ([]*Instrument, error) {
	var list []*Instrument
	if err := s.db.Select(&list, "SELECT * FROM instruments"); err != nil {
		return nil, err
	}
	return list, nil
}

// Delete ...
func (s service) Delete(ID int64) error {
	q := `DELETE FROM instruments WHERE id=(?)`
	result := s.db.MustExec(q, ID)
	_, err := result.RowsAffected()
	return err
}

// Edit ...
func (s service) Edit(i *Instrument) error {
	insertInstrument := `UPDATE instruments SET name=(?), description=(?), price=(?) WHERE id=(?)`
	result := s.db.MustExec(insertInstrument, i.Name, i.Description, i.Price, i.ID)
	_, err := result.RowsAffected()
	return err
}
