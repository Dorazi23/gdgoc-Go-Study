package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const APIKEY = "193ef3a"

type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
}

func main() {
	body, _ := SearchById("tt3896198")
	fmt.Println(body.Title)

	body, _ = SearchByName("Game of")
	fmt.Println(body.Title)
}

func SearchByName(name string) (*MovieInfo, error) {
	parms := url.Values{}
	parms.Set("apikey", APIKEY)
	parms.Set("t", name)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func SearchById(id string) (*MovieInfo, error) {
	parms := url.Values{}
	parms.Set("apikey", APIKEY)
	parms.Set("i", id)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody:" + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func sendGetRequest(url string) (string, error) {
	// 섹션 1
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// 섹션 2
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 섹션 3
	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}

	// 섹션 4
	return string(body), nil
}
