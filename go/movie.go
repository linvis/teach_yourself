package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const omdbURL = "http://www.omdbapi.com"
const apiKey = "&apikey=d1e4eadc"

type OmdbMovie struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Runtime  string
	Poster   string
}

func OmdbRequest(name string) (*OmdbMovie, error) {
	movieURL := omdbURL + "?t=" + url.QueryEscape(name) + apiKey
	fmt.Println(movieURL)

	resp, err := http.Get(movieURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query movie error")
	}

	var res OmdbMovie
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("json decode error")
	}

	return &res, nil
}

func main() {
	var name string
	for _, arg := range os.Args[1:] {
		name = name + arg + " "
	}
	name = name[0 : len(name)-1]

	info, err := OmdbRequest(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", *info)
}
