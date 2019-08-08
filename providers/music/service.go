package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/unclepeddy/lynq/clients/seatgeek"
	"github.com/unclepeddy/lynq/clients/spotify"
)

type Service struct {
	Spotify  *spotify.Client
	SeatGeek *seatgeek.Client
}

func exampleMain() {
	userId := flag.String("user-id", "", "Spotify User ID")
	flag.Parse()

	if *userId == "" {
		log.Fatalf("user-id cannot be empty - please run auth module first to obtain your user ID")
	}

	service := NewService(*userId)
	concerts, _ := service.GetConcerts()
	log.Printf("%v", concerts)
}

// NewService returns a music service with all
// required API clients instantiated.
func NewService(spotifyUser string) *Service {
	sc := spotify.NewClient(spotifyUser)
	sgc := seatgeek.NewClient()
	return &Service{
		Spotify:  sc,
		SeatGeek: sgc}
}

// GetConcerts returns a list of the user's top K artists' upcoming concerts.
func (s *Service) GetConcerts() ([]seatgeek.Concert, error) {
	arts, err := s.Spotify.GetTopKArtists(5)
	if err != nil {
		log.Fatalf("Error(s) occured during retrieving events: %s", err)
	}

	c, err := s.SeatGeek.GetArtistsConcerts(arts)
	if err != nil {
		if len(c) == 0 {
			return nil, fmt.Errorf("Could not retrieve any events: %s", err)
		}
		log.Printf("Error(s) occured during retrieving events: %s", err)
	}

	return c, nil
}
