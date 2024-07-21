package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type RoomAPI struct {
	apiClient *APIClient
}

// Issue is related to a room
type Room struct {
	RoomId int    `json:"roomId"`
	Number string `json:"number"`
}

// GetListOfRooms is a method that retrieves the list of rooms for a given theater.
func (r RoomAPI) GetListOfRooms(theaterId int) ([]Room, error) {
	var rooms []Room

	client := http.Client{}
	req, err := http.NewRequest("GET", r.apiClient.baseUrl+"/api/rooms", nil)
	if err != nil {
		log.Fatalf("Cannot create HTTP Client: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.apiClient.token)
	q := req.URL.Query()
	q.Add("theaterId", strconv.Itoa(theaterId))
	req.URL.RawQuery = q.Encode()

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

	err = json.Unmarshal(body, &rooms)
	if err != nil {
		log.Fatalf("Error while decoding JSON: %v", err)
	}

	return rooms, err
}

func NewRoomAPI(apiClient *APIClient) *RoomAPI {
	return &RoomAPI{
		apiClient: apiClient,
	}
}
