package instruments

import "seminario-GoLang/internal/config"

// Instrument ...
type Instrument struct {
	ID          int16
	name        string
	description string
	price       int32
}

// InstrumentService ...
type InstrumentService interface {
	AddInstrument(Instrument) error
	FindByID(int) *Instrument
	FindAll() []*Instrument
}

type service struct {
	conf *config.Config
}

// New ...
func New(c *config.Config) (InstrumentService, error) {
	return service{c}, nil
}

func (s service) AddInstrument(i Instrument) error {
	return nil
}

func (s service) FindByID(ID int) *Instrument {
	return nil
}

func (s service) FindAll() []*Instrument {
	var list []*Instrument
	list = append(list, &Instrument{0, "Bajo", "Bajo eléctrico de 4 cuerdas", 4500})
	return list
}
