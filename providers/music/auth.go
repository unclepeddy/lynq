package main

import (
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const (
	relTokenBaseDir = ".lynq/creds/"
	state           = "lynq"
	baseAddr        = "http://localhost"
	netAddr         = ":8080"
	redirEndpoint   = "/lynq_callback"
)

func main() {
	ch := make(chan *spotify.Client)
	auth := spotify.NewAuthenticator(baseAddr+netAddr+redirEndpoint,
		spotify.ScopeUserReadPrivate, spotify.ScopeUserReadRecentlyPlayed)
	startAuthWebhook(ch, auth)
	<-ch
}

// saveToken saves given token as a json object in a file named the given id
// followed by '-token.json' suffix in tokenBaseDir relative to user's HOME.
func saveToken(id string, token *oauth2.Token) {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to save oauth token: %v", err)
	}
	dir := filepath.Join(usr.HomeDir, relTokenBaseDir)
	os.MkdirAll(dir, 0700)

	file := fmt.Sprintf("%s-token.json", id)
	file = filepath.Join(dir, file)

	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to save oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// startAuthWebhook sets up auth webhook and prints OAuth service endpoint url
// Note: if this is ever called in a long-running process, we should explicitly do shutdown
func startAuthWebhook(ch chan *spotify.Client, auth spotify.Authenticator) {
	log.Print("Starting webserver at address ", netAddr)

	http.HandleFunc(redirEndpoint, completeAuth(ch, auth))
	go http.ListenAndServe(netAddr, nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
}

// completeAuth uses response from Spotify's auth service to create client and caches the token
func completeAuth(ch chan *spotify.Client, auth spotify.Authenticator) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tok, err := auth.Token(state, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			log.Fatal(err)
		}
		if st := r.FormValue("state"); st != state {
			http.NotFound(w, r)
			log.Fatalf("State mismatch: %s != %s\n", st, state)
		}
		client := auth.NewClient(tok)
		user, err := client.CurrentUser()
		if err != nil {
			log.Fatal(err)
		}
		saveToken(user.ID, tok)
		fmt.Fprintf(w, "Welcome %s - You are logged in as user `%s`",
			user.DisplayName, user.ID)
		ch <- &client
	}
}
