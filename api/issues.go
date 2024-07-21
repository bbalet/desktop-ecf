package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

type IssueAPI struct {
	apiClient *APIClient
}

// Issue is related to a room
type Issue struct {
	IssueID     int    `json:"issueId"`
	RoomId      int    `json:"roomId"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

// GetListOfIssues is a method that retrieves the list of issues for a given room
func (i IssueAPI) GetListOfIssues(roomId int) ([]Issue, error) {
	var issues []Issue

	client := http.Client{}
	req, err := http.NewRequest("GET", i.apiClient.baseUrl+"/api/issues", nil)
	if err != nil {
		log.Fatalf("Impossible de créer le client HTTP: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+i.apiClient.token)
	q := req.URL.Query()
	q.Add("roomId", strconv.Itoa(roomId))
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

	err = json.Unmarshal(body, &issues)
	if err != nil {
		log.Fatalf("Error while decoding JSON: %v", err)
	}

	return issues, err
}

// CreateNewIssue is a method that create a new issue
func (i IssueAPI) CreateNewIssue(issue Issue) (int, error) {

	jsonBody, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("Impossible de construire le corps de la requête: %v", err)
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", i.apiClient.baseUrl+"/api/issues", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Impossible de créer le client HTTP: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+i.apiClient.token)

	bytes, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("Out CreateNewIssue : %v\n", string(bytes[:]))

	// Do a GET request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erreur lors de la requête POST: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	/*body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
	}*/

	return response.StatusCode, err
}

// UpdateIssue is a method that patch an issue
func (i IssueAPI) UpdateIssue(issue Issue) (int, error) {

	jsonBody, err := json.Marshal(issue)
	if err != nil {
		log.Fatalf("Impossible de construire le corps de la requête: %v", err)
	}
	client := http.Client{}
	req, err := http.NewRequest("PATCH", i.apiClient.baseUrl+"/api/issues/"+strconv.Itoa(issue.IssueID), bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Impossible de créer le client HTTP: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+i.apiClient.token)
	q := req.URL.Query()
	q.Add("issueId", strconv.Itoa(issue.IssueID))
	req.URL.RawQuery = q.Encode()

	bytes, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("Out CreateNewIssue : %v\n", string(bytes[:]))

	// Do a GET request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erreur lors de la requête PATCH: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
	}

	fmt.Printf("Out CreateNewIssue : %v\n", string(body[:]))

	return response.StatusCode, err
}

func NewIssueAPI(apiClient *APIClient) *IssueAPI {
	return &IssueAPI{
		apiClient: apiClient,
	}
}
