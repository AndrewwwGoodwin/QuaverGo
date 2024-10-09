package QuaverGo

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Mapsets struct {
	APIClient         *Client
	EndpointExtension string
}

func initMapsets(apiClient *Client) *Mapsets {
	return &Mapsets{APIClient: apiClient, EndpointExtension: "/mapset/"}
}

func (m Mapsets) ByID(id int) (Mapset, error) {
	var returnedMapset MapsetJson
	err := m.APIClient.fetchData(fmt.Sprintf("%s%s%d", m.APIClient.baseURL, m.EndpointExtension, id), &returnedMapset)
	if err != nil {
		return Mapset{}, err
	}
	return returnedMapset.Mapset, nil
}

type Mapset struct {
	Id              int       `json:"id"`
	PackageMd5      string    `json:"package_md5"`
	CreatorId       int       `json:"creator_id"`
	CreatorUsername string    `json:"creator_username"`
	Artist          string    `json:"artist"`
	Title           string    `json:"title"`
	Source          string    `json:"source"`
	Tags            string    `json:"tags"`
	Description     string    `json:"description"`
	DateSubmitted   time.Time `json:"date_submitted"`
	DateLastUpdated time.Time `json:"date_last_updated"`
	IsVisible       bool      `json:"is_visible"`
	IsExplicit      bool      `json:"is_explicit"`
	Maps            []Map     `json:"maps"`
}

type MapsetJson struct {
	Mapset Mapset `json:"mapset"`
}

type MapsetsJson struct {
	Mapsets []Mapset `json:"mapsets"`
}

// RankedList returns an []int of all the ranked maps in quaver
func (m Mapsets) RankedList() ([]int, error) {
	var returnedMapset RankedList
	err := m.APIClient.fetchData(fmt.Sprintf("%s%sranked", m.APIClient.baseURL, m.EndpointExtension), &returnedMapset)
	if err != nil {
		return nil, err
	}
	return returnedMapset.RankedMapsets, nil
}

type RankedList struct {
	RankedMapsets []int `json:"ranked_mapsets"`
}

// OffsetList returns a list of all mapsets that have online offsets.
func (m Mapsets) OffsetList() ([]OffsetStruct, error) {
	var returnedOffsets OffsetList
	err := m.APIClient.fetchData(fmt.Sprintf("%s%soffsets", m.APIClient.baseURL, m.EndpointExtension), &returnedOffsets)
	if err != nil {
		return nil, err
	}
	return returnedOffsets.OnlineOffsets, nil
}

type OffsetList struct {
	OnlineOffsets []OffsetStruct `json:"online_offsets"`
}

type OffsetStruct struct {
	ID     int `json:"id"`
	Offset int `json:"offset"`
}

// Search todo mapset search
// Search uses url parameters, and I haven't built the fetchData function to handle that.
func (m Mapsets) Search(options *MapsetSearchOptions) ([]Mapset, error) {
	//yikes, here we go I guess.
	// we're gonna actually use a URL builder here.
	// next we parse our url into the builder
	urlObj, err := url.Parse(fmt.Sprintf("%s%ssearch", m.APIClient.baseURL, m.EndpointExtension))
	if err != nil {
		return nil, err
	}
	q := urlObj.Query()
	//Generate the search query parameters based on what gets passed in the SearchOptions variable
	params := map[string]string{
		"search":                options.Query,
		"ranked_status":         options.RankedStatus,
		"mode":                  options.Mode,
		"page":                  options.Page,
		"min_difficulty_rating": options.MinDifficultyRating,
		"max_difficulty_rating": options.MaxDifficultyRating,
		"min_bpm":               options.MinBPM,
		"max_bpm":               options.MaxBPM,
		"min_length":            options.MinLength,
		"max_length":            options.MaxLength,
		"min_long_note_percent": options.MinLongNotePercent,
		"max_long_note_percent": options.MaxLongNotePercent,
		"min_play_count":        options.MinPlayCount,
		"max_play_count":        options.MaxPlayCount,
		"min_combo":             options.MinCombo,
		"max_combo":             options.MaxCombo,
		"min_date_submitted":    options.MinDateSubmitted,
		"max_date_submitted":    options.MaxDateSubmitted, // this might not be a thing. please test
		"min_last_updated":      options.MinLastUpdated,
		"max_last_updated":      options.MaxLastUpdated,
		"show_explicit":         options.ShowExplicit,
	}
	for key, value := range params {
		if value == "" {
			continue
		}
		q.Set(key, value)
	}
	urlObj.RawQuery = q.Encode()
	//fmt.Println(urlObj.String())
	returnMapsetJson := MapsetsJson{}
	err = m.APIClient.fetchData(urlObj.String(), &returnMapsetJson)
	if err != nil {
		return nil, err
	}
	return returnMapsetJson.Mapsets, nil
}

// https://api.quavergame.com/v2/mapset/search?search=swan&ranked_status=1&mode=1&limit=1&min_difficulty_rating=20&max_difficulty_rating=30
type MapsetSearchOptions struct {
	Query               string
	RankedStatus        string
	Mode                string
	Page                string
	MinDifficultyRating string
	MaxDifficultyRating string
	MinBPM              string
	MaxBPM              string
	MinLength           string
	MaxLength           string
	MinLongNotePercent  string
	MaxLongNotePercent  string
	MinPlayCount        string
	MaxPlayCount        string
	MinCombo            string
	MaxCombo            string
	MinDateSubmitted    string //unix timestamp
	MaxDateSubmitted    string
	MinLastUpdated      string //unix timestamp
	MaxLastUpdated      string
	ShowExplicit        string
}

func MapsetSearchOptionsBuilder(searchQuery string) *MapsetSearchOptions {
	return &MapsetSearchOptions{Query: searchQuery}
}

func (options *MapsetSearchOptions) SetRankedStatus(ranked bool) *MapsetSearchOptions {
	if ranked {
		//ranked
		options.RankedStatus = "1"
	} else {
		//unranked
		options.RankedStatus = "2"
	}
	return options
}

// SetMode 1 is 4k, 2 is 7k. will add a builder thing later that handles this more clearly
func (options *MapsetSearchOptions) SetMode(mode int) *MapsetSearchOptions {
	options.Mode = strconv.Itoa(mode)
	return options
}
func (options *MapsetSearchOptions) SetPage(page int) *MapsetSearchOptions {
	options.Page = strconv.Itoa(page)
	return options
}
func (options *MapsetSearchOptions) SetMinDifficultyRating(minDiff float64) *MapsetSearchOptions {
	options.MinDifficultyRating = strconv.FormatFloat(minDiff, 'f', -1, 64)
	return options
}
func (options *MapsetSearchOptions) SetMaxDifficultyRating(maxDiff float64) *MapsetSearchOptions {
	options.MinDifficultyRating = strconv.FormatFloat(maxDiff, 'f', -1, 64)
	return options
}
func (options *MapsetSearchOptions) SetMinBPM(minBPM float64) *MapsetSearchOptions {
	options.MinBPM = strconv.FormatFloat(minBPM, 'f', -1, 64)
	return options
}
func (options *MapsetSearchOptions) SetMaxBPM(maxBPM float64) *MapsetSearchOptions {
	options.MaxBPM = strconv.FormatFloat(maxBPM, 'f', -1, 64)
	return options
}
func (options *MapsetSearchOptions) SetMinLength(minLength int) *MapsetSearchOptions {
	options.MinLength = strconv.Itoa(minLength)
	return options
}
func (options *MapsetSearchOptions) SetMaxLength(maxLength int) *MapsetSearchOptions {
	options.MaxLength = strconv.Itoa(maxLength)
	return options
}
func (options *MapsetSearchOptions) SetMinLongNotePercent(minLongNotePercent int) *MapsetSearchOptions {
	options.MinLongNotePercent = strconv.Itoa(minLongNotePercent)
	return options
}
func (options *MapsetSearchOptions) SetMaxLongNotePercent(maxLongNotePercent int) *MapsetSearchOptions {
	options.MaxLongNotePercent = strconv.Itoa(maxLongNotePercent)
	return options
}
func (options *MapsetSearchOptions) SetMinPlayCount(minPlayCount int) *MapsetSearchOptions {
	options.MinPlayCount = strconv.Itoa(minPlayCount)
	return options
}
func (options *MapsetSearchOptions) SetMaxPlayCount(maxPlayCount int) *MapsetSearchOptions {
	options.MaxPlayCount = strconv.Itoa(maxPlayCount)
	return options
}
func (options *MapsetSearchOptions) SetMinCombo(minCombo int) *MapsetSearchOptions {
	options.MinCombo = strconv.Itoa(minCombo)
	return options
}
func (options *MapsetSearchOptions) SetMaxCombo(maxCombo int) *MapsetSearchOptions {
	options.MaxCombo = strconv.Itoa(maxCombo)
	return options
}

// SetShowExplicit these values are just kinda a guess since I cant find a map marked as explicit lol
func (options *MapsetSearchOptions) SetShowExplicit(showExplicit bool) *MapsetSearchOptions {
	if showExplicit {
		//explicit
		options.ShowExplicit = "1"
	} else {
		//not explicit
		options.ShowExplicit = "0"
	}
	return options
}
func (options *MapsetSearchOptions) SetMinDateSubmitted(date time.Time) *MapsetSearchOptions {
	options.MinDateSubmitted = strconv.FormatInt(date.Unix(), 10)
	return options
}
func (options *MapsetSearchOptions) SetMaxDateSubmitted(date time.Time) *MapsetSearchOptions {
	options.MaxDateSubmitted = strconv.FormatInt(date.Unix(), 10)
	return options
}
func (options *MapsetSearchOptions) SetMinLastUpdated(date time.Time) *MapsetSearchOptions {
	options.MinLastUpdated = strconv.FormatInt(date.Unix(), 10)
	return options
}
func (options *MapsetSearchOptions) SetMaxLastUpdated(date time.Time) *MapsetSearchOptions {
	options.MaxLastUpdated = strconv.FormatInt(date.Unix(), 10)
	return options
}
