package rocket

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var apiKey = os.Getenv("API_KEY")
var client = NewRocketClient(apiKey)

func TestRocketClient(t *testing.T) {
	client := NewRocketClient("123456")
	assert.Equal(t, "123456", client.ApiKey)
}

func TestPrepareRequest(t *testing.T) {
	req := client.PrepareRequest("/test/")
	assert.Equal(t, client.ApiKey, req.Header.Get("Authorization"))
	assert.Equal(t, "/test/", req.URL.Path)
}

func TestRocketGet(t *testing.T) {
	time.Sleep(1000)
	response, _ := client.Get("data/platforms")
	assert.Equal(t, response.Ok, true)
	assert.Equal(t, response.StatusCode, 200)
}

func TestPlatform(t *testing.T) {
	time.Sleep(1000)
	response := client.Platforms()
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, `[{"id":1,"name":"Steam"},{"id":2,"name":"Ps4"},{"id":3,"name":"XboxOne"}]`, response.Content)
}

func TestSeasons(t *testing.T) {
	time.Sleep(1000)
	response := client.Seasons()
	assert.Equal(t, 200, response.StatusCode)
}

func TestPlaylists(t *testing.T) {
	time.Sleep(1000)
	response := client.Playlists()
	assert.Equal(t, 200, response.StatusCode)
}

func TestTiers(t *testing.T) {
	time.Sleep(1000)
	response := client.Tiers()
	assert.Equal(t, 200, response.StatusCode)
}

func TestPlayer(t *testing.T) {
	time.Sleep(1000)
	response := client.Player("76561198079681869", 1)
	assert.Equal(t, 200, response.StatusCode)
}

func TestSearchPlayers(t *testing.T) {
	time.Sleep(1000)
	response := client.SearchPlayers("Girbons")
	assert.Equal(t, 200, response.StatusCode)
}

func TestRankedLeaderboard(t *testing.T) {
	time.Sleep(1000)
	response := client.RankedLeaderboard(10)
	assert.Equal(t, 200, response.StatusCode)
}

func TestStatLeaderboard(t *testing.T) {
	time.Sleep(1000)
	response := client.StatLeaderboard("wins")
	assert.Equal(t, 200, response.StatusCode)
}
