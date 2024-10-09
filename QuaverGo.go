package QuaverGo

import (
	"encoding/json"
	"github.com/AndrewwwGoodwin/QuaverGo/RateLimitManager"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient       *http.Client
	rateLimitManager *RateLimitManager.RateLimitManager
	baseURL          string
	//endpoint subdivision
	/*
		Clans        *Clans.Clans
		Leaderboards *Leaderboards.Leaderboards
		Multiplayer  *Multiplayer.Multiplayer
		Playlists    *Playlists.Playlists
		RankingQueue *RankingQueue.RankingQueue
	*/
	Download    *Download
	ServerStats *ServerStats
	Mapsets     *Mapsets
	Maps        *Maps
	Scores      *Scores
	Users       *Users
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
		initClient.Leaderboards = Leaderboards.Init(&initClient)
		initClient.Multiplayer = Multiplayer.Init(&initClient)
		initClient.Playlists = Playlists.Init(&initClient)
		initClient.RankingQueue = RankingQueue.Init(&initClient)
	*/
	initClient.Download = initDownload(&initClient)
	initClient.ServerStats = initServerStats(&initClient)
	initClient.Mapsets = initMapsets(&initClient)
	initClient.Maps = initMaps(&initClient)
	initClient.Scores = initScores(&initClient)
	initClient.Users = initUsers(&initClient)

	return &initClient
}

// fetchData takes in a target url, and the pointer to unMarshal into.
// only returns an error, as the data is unmarshalled into the pointer
func (c Client) fetchData(endpoint string, target interface{}) error {
	// see if the RateLimitManager lets us create a request
	err := c.rateLimitManager.Request()
	if err != nil {
		return err
	}

	// if we are allowed to make the request, send a GET to the provided endpoint
	response, err := http.Get(endpoint)

	if err != nil {
		return err
	}
	defer response.Body.Close()
	// read the response information
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// and finally unmarshal into the provided target interface
	return json.Unmarshal(data, target)
}
