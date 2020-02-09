package pets

// Denne pakken representerer forretningslogikken til servicen.
// Den har sin egen representasjon av data, gjør nødvendige operasjoner
// på data og har veldefinerte typer som kan returneres.

import (
	"context"
	"log"
	"strconv"

	"github.com/apparatno/sample-webservice/repository"
)

// Definer typer som representerer forretningslogikken vår.

type Pet struct {
	ID   string
	Name string
}

// Definer en controller som holder en pointer til repositoryet vårt.
// Controlleren kan ha flere nyttige felt; logger,

type PetsService struct {
	db *repository.Database
}

func New(repo *repository.Database) (*PetsService, error) {
	s := PetsService{db: repo}
	return &s, nil
}

// implementer forretningslogikk som funksjoner

func (s *PetsService) Create(ctx context.Context, p Pet) (Pet, error) {
	data := repository.Pet{Name: p.Name}
	var err error
	data, err = s.db.Store(ctx, data)
	if err != nil {
		return Pet{}, err
	}

	log.Printf("pet stored %v", data)
	res := makePet(data)

	return res, nil
}

func (s *PetsService) GetAll(ctx context.Context) ([]Pet, error) {
	allPets, err := s.db.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// map repository sin datatype til servicens type
	res := make([]Pet, len(allPets))
	for i, p := range allPets {
		somePet := makePet(p)
		res[i] = somePet
	}

	return res, nil
}

func makePet(pet repository.Pet) Pet {
	ID := strconv.FormatInt(pet.ID, 10)
	res := Pet{Name: pet.Name, ID: ID}
	return res
}
