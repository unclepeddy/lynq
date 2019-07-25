package main

import (
	"encoding/json"

	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const (
	relTokenBaseDir = ".lynq/creds/"
)

func main() {
	userId := flag.String("user_id", "", "Spotify User ID")
	k := flag.Int("top_k", 5, "Number of top artists")
	flag.Parse()

	if *userId == "" {
		log.Fatalf("user_id cannot be empty - please run auth module first to obtain your user ID")
	}

	client := getClient(*userId)
	topArtists := getTopKArtists(client, *k)
	log.Print(topArtists)
}

// getClient tries to create a client with local credentials by calling loadToken()
func getClient(id string) *spotify.Client {
	tok, err := loadToken(id)
	if err != nil {
		log.Fatalf("Token not found locally for user `%v`: %v", id, err)
	}

	auth := spotify.NewAuthenticator("",
		spotify.ScopeUserReadPrivate, spotify.ScopeUserReadRecentlyPlayed)
	c := auth.NewClient(tok)
	return &c
}

// getTopKArtists returns the top K artist names found within the 20 most recently played tracks
func getTopKArtists(client *spotify.Client, k int) []string {
	items, err := client.PlayerRecentlyPlayed()
	if err != nil {
		log.Fatal("Could not get recently played tracks: ", err)
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

	topK := make([]string, 0, k)
	for _, ac := range aslice[:k] {
		topK = append(topK, ac.K)
	}

	return topK
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
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}
