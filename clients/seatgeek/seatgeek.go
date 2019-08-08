package seatgeek

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gosimple/slug"
)

const (
	clientIdEnvVar = "SEATGEEK_CLIENT_ID"
)

type Client struct {
	http     *http.Client
	baseAddr string
	clientId string
}

// NewClient returns a client for SeatGeek API
func NewClient() *Client {
	clientId := os.Getenv(clientIdEnvVar)
	if clientId == "" {
		log.Fatalf("Please provide client id by setting %s", clientIdEnvVar)
	}
	return &Client{
		http:     &http.Client{},
		baseAddr: "https://api.seatGeek.com/2/",
		clientId: clientId}
}

type Concert struct {
	Id       string
	Title    string
	City     string
	Datetime string
}

// GetArtistsConcerts uses GetArtistConcerts to return a list of all provided
// artists' upcoming concerts.
func (c *Client) GetArtistsConcerts(names []string) ([]Concert, error) {
	concerts := make([]Concert, 0, len(names))
	// TODO: Make requests concurrently after validating seatgeek doesn't rate limit; dedup concerts
	for _, name := range names {
		es, err := c.GetArtistConcerts(name)
		if err != nil {
			return nil, fmt.Errorf("Could not retrieve concerts for artist `%s`: %s", name, err)
		}
		for _, e := range es {
			if (e != Concert{}) {
				concerts = append(concerts, e)
			}
		}
	}
	return concerts, nil
}

// GetArtistConcerts looks up artist by name and returns a list of
// their upcoming concerts.
func (c *Client) GetArtistConcerts(name string) ([]Concert, error) {
	name = slug.Make(name)
	path := fmt.Sprintf("events?performers.slug=%s", name)
	res, err := c.get(path)
	if err != nil {
		return nil, err
	}

	concertsMeta, ok := res["events"]
	if !ok {
		return nil, fmt.Errorf("Could not find 'events' array in response")
	}
	concerts := make([]Concert, len(concertsMeta.([]interface{})))
	for i, concertMeta := range concertsMeta.([]interface{}) {
		var ct Concert
		concert := concertMeta.(map[string]interface{})
		ct.Id = fmt.Sprintf("%.0f", concert["id"].(float64))
		ct.Title = concert["title"].(string)
		venue := concert["venue"].(map[string]interface{})
		ct.City = venue["city"].(string)
		ct.Datetime = concert["datetime_utc"].(string)
		concerts[i] = ct
	}
	return concerts, nil
}

// TODO: consider changing to the .NewRequest() and .Do model.
// get executes a GET call and returns a map of the response object
func (c *Client) get(path string) (map[string]interface{}, error) {
	url := c.baseAddr + path + "&client_id=" + c.clientId

	res, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return nil, fmt.Errorf("Received non-OK HTTP Status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}
	err = json.Unmarshal([]byte(body), &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
