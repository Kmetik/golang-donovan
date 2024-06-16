package ch4

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const api_key = "5bb0c2b8"

const api_url = "http://www.omdbapi.com/"

type Rating struct {
	Source string
	Value  string
}

type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Actors     string
	Plot       string
	Language   string
	Poster     string
	Ratings    []*Rating
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

var search = flag.String("s", " ", "movie to find")

func GetPoster() {
	flag.Parse()
	var movieResult Movie
	q := url.QueryEscape(*search)
	resp, err := http.Get(api_url + "?apikey=" + api_key + "&t=" + q)
	fmt.Println(resp.Request.URL)
	if err != nil {
		fmt.Printf("err %s, %s", *search, err.Error())
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		os.Exit(1)
	}

	if err := json.NewDecoder(resp.Body).Decode(&movieResult); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	resp2, err := http.Get(movieResult.Poster)
	if err != nil {
		fmt.Println(resp2.StatusCode)
		os.Exit(1)
	}
	os.Mkdir("posters", 0777)
	out, err := os.Create("posters/" + *search + ".jpg")
	if err != nil {
		fmt.Println("cannot write file", err)
		os.Exit(1)
	}
	defer out.Close()
	if written, err := io.Copy(out, resp2.Body); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(written)
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()
}
