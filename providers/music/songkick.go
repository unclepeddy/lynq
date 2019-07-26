## Feel free to ignore this file 
package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
)

const songkickBaseAddr = "https://api.songkick.com/api/3.0"

type Client struct {
	http   *http.Client
	apiKey string
}

// NewSongkickClient returns a client for Songkick API
func NewSongkickClient(apiKey string) Client {
	return Client{http: &http.Client{}, apiKey: apiKey}
}

// GetArtistEvents looks up artist by name and returns a list of
// their upcoming events.
func (c *Client) GetArtistEvents(name string) (string, error) {
	id, err := c.getArtistId(name)
	if err != nil {
		return "", err
	}
	// TODO: call getArtistEvents(id)
	return id, nil
}

// getArtistEvents returns a list of upcoming events for artist with given id
func (c *Client) getArtistEvents(id string) (string, error) {
	url := fmt.Sprintf("%s/artists/%s/calendar.json?apikey=%s", songkickBaseAddr, id, c.apiKey)
	var res struct {
		Artist Artist `json:"artist"`
	}
	err := c.get(url, &res)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not retrieve artist id for `%s`: %s", name, err))
	}
	return res.Artist.ID, nil
}

type Artist struct {
	ID          string `json:"id"`
	URI         string `json:"uri"`
	DisplayName string `json:"displayName"`
	OnTourUntil string `json:"onTourUntil"`
}

// getArtistId retrieves Songkick artist id for a given name
func (c *Client) getArtistId(name string) (string, error) {
	url := fmt.Sprintf("%s/search/artists.json?apikey=%s&query=%s", songkickBaseAddr, c.apiKey, name)
	var res struct {
		Artist Artist `json:"artist"`
	}
	err := get(url, &res)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not retrieve artist id for `%s`: %s", name, err))
	}
	return res.Artist.ID, nil
}

// get executes a GET call and decodes the response using provided interface
func (c *Client) get(url string, result interface{}) error {
	var r struct {
		Results result `json:"results"`
	}

	// TODO: validate you're initing this correctly
	var p struct {
		ResultsPage Results `json:"resultsPage"`
	}

	res, err := c.http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if err := json.NewDecoder(res).Decode(p); err != nil {
		return "", errors.New(fmt.Sprintf("Could not decode response from `%s` with type %s: %s", url, result, err))
	}
	return nil
}
