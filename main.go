package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apparatno/sample-webservice/pets"
	"github.com/apparatno/sample-webservice/repository"
)

func main() {
	// ikke bruk default mux men lag din egen
	mux := http.NewServeMux()

	// opprett instanser av avhengigheter

	data := make(map[int64]repository.Pet)
	repo, err := repository.New(data)
	if err != nil {
		log.Fatal(err)
	}

	service, err := pets.New(repo)
	if err != nil {
		log.Fatal(err)
	}

	handlePetsFunc := makeHandlePetsFunc(service)

	// opprett handlers for endepunkt
	mux.HandleFunc("/pets", handlePetsFunc)

	// Start server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

type Pet struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// denne funksjonen implementerer selve håndteringen av http requesten.
// Den har til ansvar å mappe innkommende data til det formatet som servicen
// skjønner (typisk deserialisere JSON til lokale typer), luke ut feil,
// og mappe feil fra servicen til HTTP status koder.
// Hvordan man velger å representere data i dette laget er litt opp til en selv.
// I dette eksempelet har jeg valgt å lage en struct som definerer formatet,
// men det kunne også være en map for å gjøre det litt enklere.
// På denne måten kan man også lage middleware som for eksempel autentisering.
func makeHandlePetsFunc(s *pets.PetsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Dette kunne simplifiseres ved å enten implementere hver enkelt http method
		// i sin egen funksjon eller bruke en annen mux som for eksempel Gorilla Mux.
		switch r.Method {
		case http.MethodGet:
			// GET /pets henter alle pets fra servicen, mapper og returnerer JSON
			all, err := s.GetAll(ctx)
			if err != nil {
				status := errToCode(err)
				w.WriteHeader(status)
				w.Write([]byte(err.Error()))
				return
			}

			res := make([]Pet, len(all))
			for i, p := range all {
				res[i] = Pet{ID: p.ID, Name: p.Name}
			}

			encoder := json.NewEncoder(w)
			if err := encoder.Encode(res); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
		case http.MethodPost:
			// POST /pets forventer en Pet JSON og sender den til servicen.
			var data Pet
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("unable to decode input data"))
				return
			}

			p := pets.Pet{Name: data.Name}

			created, err := s.Create(ctx, p)
			if err != nil {
				status := errToCode(err)
				w.WriteHeader(status)
				w.Write([]byte(err.Error()))
				return
			}

			pet := Pet{
				ID:   created.ID,
				Name: created.Name,
			}

			encoder := json.NewEncoder(w)
			if err = encoder.Encode(&pet); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("unable to write response"))
				return
			}

			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func errToCode(err error) int {
	switch err {
	case repository.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
