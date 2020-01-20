package spotify

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
	relTokenBaseDir      = ".lynq/creds/"
	state                = "lynq"
	baseAddr             = "http://localhost"
	netAddr              = ":8080"
	spotifyRedirEndpoint = "/spotify_callback"
)

func exampleAuth() {
	auth := spotify.NewAuthenticator(baseAddr+netAddr+spotifyRedirEndpoint,
		spotify.ScopeUserReadPrivate, spotify.ScopeUserReadRecentlyPlayed)
	client := authWebhook(auth)
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Obtained credentials for user `%s`\n", user.ID)
}

// saveToken saves given token as a json object in a file named the given id
// followed by '-token.json' suffix in tokenBaseDir relative to user's HOME.
func saveToken(id string, token *oauth2.Token) {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to save oauth token: %v", err)
	}
	dir := filepath.Join(usr.HomeDir, relTokenBaseDir)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		log.Fatalf("Unable to save oauth token: %v", err)
	}

	file := fmt.Sprintf("%s-token.json", id)
	file = filepath.Join(dir, file)

	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to save oauth token: %v", err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatalf("Unable to save oauth token: %v", err)
	}
}

// authWebhook sets up auth webhook, prints OAuth service endpoint url and blocks
// until user has authenticated with Spotify, at which point returns spotify client.
// Note: if this is ever called in a long-running process, we should explicitly do shutdown
func authWebhook(auth spotify.Authenticator) *spotify.Client {
	ch := make(chan *spotify.Client)
	log.Print("Starting webserver at address ", netAddr)

	http.HandleFunc(spotifyRedirEndpoint, completeAuth(ch, auth))
	go http.ListenAndServe(netAddr, nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	return <-ch
}

// completeAuth uses response from Spotify's auth service to create client and caches the token
func completeAuth(ch chan *spotify.Client, auth spotify.Authenticator) http.HandlerFunc {
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
