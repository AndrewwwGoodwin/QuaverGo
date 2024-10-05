package QuaverGo

import (
	"QuaverGo/RateLimitManager"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	httpClient       *http.Client
	rateLimitManager *RateLimitManager.RateLimitManager
	baseURL          string
	//endpoint subdivision
	/*
		Clans        *Clans.Clans
		Download     *Downloads.Downloads
		Leaderboards *Leaderboards.Leaderboards
		Maps         *Maps.Maps
		Mapsets      *Mapsets.Mapsets
		Multiplayer  *Multiplayer.Multiplayer
		Playlists    *Playlists.Playlists
		RankingQueue *RankingQueue.RankingQueue
		ServerStats  *ServerStats.ServerStats
		Scores       *Scores.Scores
	*/
	Users *Users
}

func Init() *Client {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	rateLimiter := RateLimitManager.Init(100)
	initClient := Client{httpClient: httpClient, rateLimitManager: rateLimiter}
	initClient.baseURL = "https://api.quavergame.com/v2"
	/*
		initClient.Clans = Clans.Init(&initClient)
		initClient.Download = Downloads.Init(&initClient)
		initClient.Leaderboards = Leaderboards.Init(&initClient)
		initClient.Maps = Maps.Init(&initClient)
		initClient.Mapsets = Mapsets.Init(&initClient)
		initClient.Multiplayer = Multiplayer.Init(&initClient)
		initClient.Playlists = Playlists.Init(&initClient)
		initClient.RankingQueue = RankingQueue.Init(&initClient)
		initClient.ServerStats = ServerStats.Init(&initClient)
		initClient.Scores = Scores.Init(&initClient)
	*/
	initClient.Users = initUsers(&initClient)

	return &initClient
}

// AttemptRequest takes in a formatted string of an api endpoint, performs a GET request to the specified URL, and returns the response or error
func (c Client) AttemptRequest(URL string) (*http.Response, error) {
	// tell the rate limit handler that we want to make a new request
	err := c.rateLimitManager.Request()
	if err != nil {
		// if rate limiter says no, we've exceeded our limit and the request fails
		return nil, err
	}
	// otherwise, proceed and attempt to create the request
	response, err := c.httpClient.Get(URL)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("received bad status code " + strconv.Itoa(response.StatusCode))
	}

	return response, nil
}
