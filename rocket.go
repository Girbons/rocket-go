package rocket

import (
	"fmt"
	"net/http"
)

const API_ENDPOINT = "https://api.rocketleaguestats.com/v1/"

// RocketClient.
type RocketClient struct {
	ApiKey string
	Client *http.Client
}

// NewRocketCliet.
func NewRocketClient(apiKey string) *RocketClient {
	return &RocketClient{
		ApiKey: apiKey,
		Client: &http.Client{},
	}
}

// PrepareRequest with Authorization Header.
func (rocket *RocketClient) PrepareRequest(endpoint string) *http.Request {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Authorization", rocket.ApiKey)
	return req
}

// Get make the request.
func (rocket *RocketClient) Get(endpoint string) (*Response, error) {
	request := rocket.PrepareRequest(fmt.Sprintf("%s%s", API_ENDPOINT, endpoint))
	response, err := rocket.Client.Do(request)
	return handleResponse(response, err)
}

// Platforms endpoint.
func (rocket *RocketClient) Platforms() *Response {
	response, _ := rocket.Get("data/platforms")
	return response
}

// Seasons endpoint.
func (rocket *RocketClient) Seasons() *Response {
	response, _ := rocket.Get("data/seasons")
	return response
}

// Playlists endpoint.
func (rocket *RocketClient) Playlists() *Response {
	response, _ := rocket.Get("data/playlists")
	return response
}

// Tiers endpoint.
func (rocket *RocketClient) Tiers() *Response {
	response, _ := rocket.Get("data/tiers")
	return response
}

// Player info endpoint.
func (rocket *RocketClient) Player(id string, platform int) *Response {
	playerEndpoint := fmt.Sprintf("player?unique_id=%s&platform_id=%d", id, platform)
	response, _ := rocket.Get(playerEndpoint)
	return response
}

// SearchPlayers endpoint.
func (rocket *RocketClient) SearchPlayers(name string) *Response {
	searchPlayersEndpoint := fmt.Sprintf("search/players?display_name=%s", name)
	response, _ := rocket.Get(searchPlayersEndpoint)
	return response
}

// RankedLeaderboard endpoint.
func (rocket *RocketClient) RankedLeaderboard(id int) *Response {
	rankedLeaderboardEndpoint := fmt.Sprintf("leaderboard/ranked?playlist_id=%d", id)
	response, _ := rocket.Get(rankedLeaderboardEndpoint)
	return response
}

//StatLeaderboard endpoint.
func (rocket *RocketClient) StatLeaderboard(statType string) *Response {
	statLeaderboardEndpoint := fmt.Sprintf("leaderboard/stat?type=%s", statType)
	response, _ := rocket.Get(statLeaderboardEndpoint)
	return response
}
