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
func (rocket *RocketClient) Get(endpoint string) (string, error) {
	request := rocket.PrepareRequest(fmt.Sprintf("%s%s", API_ENDPOINT, endpoint))
	response, err := rocket.Client.Do(request)
	return handleResponse(response, err)
}

// Platforms endpoint.
func (rocket *RocketClient) Platforms() string {
	response, _ := rocket.Get("data/platformsssss")
	return response
}

// Seasons endpoint.
func (rocket *RocketClient) Seasons() string {
	response, _ := rocket.Get("data/seasons")
	return response
}

// Playlists endpoint.
func (rocket *RocketClient) Playlists() string {
	response, _ := rocket.Get("data/playlists")
	return response
}

// Tiers endpoint.
func (rocket *RocketClient) Tiers() string {
	response, _ := rocket.Get("data/tiers")
	return response
}

// Player info endpoint.
func (rocket *RocketClient) Player(id, platform int) string {
	playerEndpoint := fmt.Sprintf("player?unique_id=%d&platform_id=%d", id, platform)
	response, _ := rocket.Get(playerEndpoint)
	return response
}

// SearchPlayers endpoint.
func (rocket *RocketClient) SearchPlayers(name string) string {
	searchPlayersEndpoint := fmt.Sprintf("search/players?display_name=%s", name)
	response, _ := rocket.Get(searchPlayersEndpoint)
	return response
}

// RankedLeaderboard endpoint.
func (rocket *RocketClient) RankedLeaderboard(id int) string {
	rankedLeaderboardEndpoint := fmt.Sprintf("leaderboard/ranked?playlist_id=%d", id)
	response, _ := rocket.Get(rankedLeaderboardEndpoint)
	return response
}

//StatLeaderboard endpoint.
func (rocket *RocketClient) StatLeaderboard(statType string) string {
	statLeaderboardEndpoint := fmt.Sprintf("leaderboard/stat?type=%s", statType)
	response, _ := rocket.Get(statLeaderboardEndpoint)
	return response
}
