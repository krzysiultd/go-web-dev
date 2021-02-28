package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://movie-database-imdb-alternative.p.rapidapi.com/?s=Avengers%20Endgame&page=1&r=json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "undefined")
	req.Header.Add("x-rapidapi-host", "movie-database-imdb-alternative.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	// need to register credit card to use free plan (up to 1000 requests / day)
}
