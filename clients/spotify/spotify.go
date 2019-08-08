package spotify

import (
	"encoding/json"

	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"

	upstream "github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type Client upstream.Client

// NewClient tries to create a client with local credentials by calling loadToken()
func NewClient(id string) *Client {
	tok, err := loadToken(id)
	if err != nil {
		log.Fatalf("Token not found locally for user `%v`: %v", id, err)
	}

	_, idOk := os.LookupEnv("SPOTIFY_ID")
	_, secOk := os.LookupEnv("SPOTIFY_SECRET")
	if !idOk || !secOk {
		log.Fatalf("Please provide SPOTIFY_ID and SPOTIFY_SECRET")
	}

	auth := upstream.NewAuthenticator("",
		upstream.ScopeUserReadPrivate, upstream.ScopeUserReadRecentlyPlayed)
	c := Client(auth.NewClient(tok))
	return &c
}

// GetTopKArtists returns the top K artist names found within the 20 most recently played tracks
func (c *Client) GetTopKArtists(k int) ([]string, error) {
	client := (*upstream.Client)(c)
	items, err := client.PlayerRecentlyPlayed()
	if err != nil {
		return nil, fmt.Errorf("Could not get recently played tracks: %s", err)
	}
	// Build map of artist counts
	// TODO(unclepeddy): change index to spotify.ID and defer the name lookup to after aggregation
	arts := make(map[string]int)
	for _, item := range items {
		for _, art := range item.Track.Artists {
			if _, ok := arts[art.Name]; !ok {
				arts[art.Name] = 0
			}
			arts[art.Name] += 1
		}
	}

	// Extract top k artists
	type KV struct {
		K string
		V int
	}
	aslice := make([]KV, 0, len(arts))
	for k := range arts {
		aslice = append(aslice, KV{k, arts[k]})
	}
	sort.Slice(aslice, func(i, j int) bool { return aslice[i].V > aslice[j].V })

	if len(aslice) < k {
		k = len(aslice)
	}
	topK := make([]string, k)
	for i, ac := range aslice[:k] {
		topK[i] = ac.K
	}

	return topK, nil
}

// loadToken loads auth token from a file named by the given id
// followed by '-token.json' suffix in tokenBaseDir relative to user's HOME.
func loadToken(id string) (*oauth2.Token, error) {
	file := fmt.Sprintf("%s-token.json", id)
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	file = filepath.Join(usr.HomeDir, relTokenBaseDir, file)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	return t, err
}
