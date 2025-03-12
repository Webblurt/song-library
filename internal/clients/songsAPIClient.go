package clients

import (
	"encoding/json"
	"net/http"
	"net/url"
	models "song-library/internal/models"
)

func (c *ExternalAPIClient) FetchSongInfo(group, song string) (models.SongDetail, error) {
	apiURL, err := url.Parse(c.URL + "/info")
	if err != nil {
		return models.SongDetail{}, err
	}

	query := apiURL.Query()
	query.Set("group", group)
	query.Set("song", song)
	apiURL.RawQuery = query.Encode()
	c.log.Debug("Fetching song info from:", apiURL.String())

	resp, err := http.Get(apiURL.String())
	if err != nil {
		return models.SongDetail{}, err
	}
	defer resp.Body.Close()

	c.log.Debug("Status code: ", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return models.SongDetail{}, err
	}

	var songDetail models.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return models.SongDetail{}, err
	}

	c.log.Debug("Fetched song info successfully:", songDetail)
	return songDetail, nil
}
