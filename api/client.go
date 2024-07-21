package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// APIClient represents the configuration of the API
type APIClient struct {
	baseUrl    string
	hostName   string
	scheme     string
	token      string
	TheaterAPI *TheaterAPI
	RoomAPI    *RoomAPI
	IssueAPI   *IssueAPI
}

// TokenResponse contains the JWT token
type TokenResponse struct {
	Token string `json:"token"`
}

// User represents the user and its role
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

// login to the API
func (a *APIClient) Login(email string, password string) (int, error) {
	var tokenResponse TokenResponse

	loginBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    email,
		Password: password,
	}
	jsonBody, err := json.Marshal(loginBody)
	if err != nil {
		log.Fatalf("Cannot Marshal Body request: %v", err)
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", a.baseUrl+"/auth", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Impossible de créer le client HTTP: %v", err)
		return 0, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Do a POST request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erreur lors de la requête POST: %v", err)
		return 0, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
		return 0, err
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatalf("Error while decoding JSON: %v", err)
		return 0, err
	}
	a.token = tokenResponse.Token

	return response.StatusCode, err
}

// WhoAmI is a method that retrieves the information of the logged in user.
func (a *APIClient) WhoAmI() (User, error) {
	var user User

	client := http.Client{}
	req, err := http.NewRequest("GET", a.baseUrl+"/api/whoami", nil)
	if err != nil {
		log.Fatalf("Cannot create HTTP Client: %v", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.token)

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

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("Error while decoding JSON: %v", err)
	}

	return user, err
}

// Logout clears the JWT token. It disconnects the current user.
func (a *APIClient) Logout() {
	a.token = ""
}

// login to the API and intitialize the API client
func NewAPIClientInit(baseUrl string, email string, password string) (*APIClient, error) {
	var tokenResponse TokenResponse

	loginBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    email,
		Password: password,
	}
	jsonBody, err := json.Marshal(loginBody)

	fmt.Printf("Struct err : %v\n", err)
	fmt.Printf("Struct : %v\n", string(jsonBody[:]))

	client := http.Client{}
	req, err := http.NewRequest("POST", baseUrl+"/auth", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Cannot create HTTP Client: %v", err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Do a POST request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erreur lors de la requête POST: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
		return nil, err
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatalf("Error while decoding JSON: %v", err)
		return nil, err
	}

	url, _ := url.Parse(baseUrl)
	apiClient := &APIClient{
		baseUrl:  baseUrl,
		hostName: url.Hostname(),
		scheme:   url.Scheme,
		token:    tokenResponse.Token,
	}
	apiClient.TheaterAPI = NewTheaterAPI(apiClient)
	apiClient.RoomAPI = NewRoomAPI(apiClient)
	apiClient.IssueAPI = NewIssueAPI(apiClient)
	return apiClient, err
}

// login to the API and intitialize the API client
func NewAPIClient(baseUrl string) *APIClient {
	url, _ := url.Parse(baseUrl)
	apiClient := &APIClient{
		baseUrl:  baseUrl,
		hostName: url.Hostname(),
		scheme:   url.Scheme,
	}
	apiClient.TheaterAPI = NewTheaterAPI(apiClient)
	apiClient.RoomAPI = NewRoomAPI(apiClient)
	apiClient.IssueAPI = NewIssueAPI(apiClient)
	return apiClient
}
