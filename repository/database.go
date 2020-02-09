package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type Database struct {
	// ha pointers til din database klient eller pool her
	data map[int64]Pet
}

// Definer typer av feil som tydelig forteller årsaken til feilen

var ErrNotFound = errors.New("not found")

// Definer typer som databasen kan jobbe med

type Pet struct {
	ID   int64
	Name string
}

// Implementer en funksjon som kan initialisere en ny instans av repoet.
// Typisk vil man opprette connection utenfor (i main) og injisere inn hit.

func New(d map[int64]Pet) (*Database, error) {
	db := Database{data: d}
	return &db, nil
}

// Implementer funksjoner til å lagre og lese data.
// Vi bruker pointer receiver for å sikre at den underliggende
// klienten er den samme som vi opprinnelig opprettet connection med.

func (d *Database) Store(ctx context.Context, p Pet) (Pet, error) {
	p.ID = time.Now().Unix()
	d.data[p.ID] = p
	return p, nil // her kunne vi returnere en feil hvis vår database feiler
}

func (d *Database) Get(ID int64) (Pet, error) {
	p, ok := d.data[ID]
	if !ok {
		return Pet{}, ErrNotFound
	}

	return p, nil
}

func (d *Database) GetAll(ctx context.Context) ([]Pet, error) {
	all := make([]Pet, 0, len(d.data))
	for _, p := range d.data {
		all = append(all, p)
	}
	return all, nil
}
