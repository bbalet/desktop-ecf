package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type TheaterAPI struct {
	apiClient *APIClient
}

// Cinema represents the structure of a cinema.
type Theater struct {
	TheaterID int     `json:"theaterId"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// getListOfTheaters is a method that retrieves the list of theaters from the API.
func (t TheaterAPI) GetListOfTheaters() ([]Theater, error) {
	var theaters []Theater

	client := http.Client{}
	req, err := http.NewRequest("GET", t.apiClient.baseUrl+"/api/theaters", nil)
	if err != nil {
		log.Fatalf("Cannot create HTTP Client: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	// Do a GET request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erreur lors de la requête GET: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
	}

	err = json.Unmarshal(body, &theaters)
	if err != nil {
		log.Fatalf("Error while decoding JSON: %v", err)
	}

	return theaters, err
}

func NewTheaterAPI(apiClient *APIClient) *TheaterAPI {
	return &TheaterAPI{
		apiClient: apiClient,
	}
}
