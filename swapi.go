package swapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const BaseURL = "http://swapi.co/api"

var ErrNotFound = errors.New("404: not found")

var client = &http.Client{}

// Get makes an API call to the path, it decodes the JSON response
// and stores it into out.
func Get(path string, out interface{}) error {
	url := path
	if path[:4] != "http" {
		url = BaseURL + path
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", "Go Swapi. github.com/leejarvis/swapi")

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 404 {
		return ErrNotFound
	}

	if err = json.NewDecoder(res.Body).Decode(out); err != nil {
		return err
	}

	return nil
}

// GetRoot returns a map of the API resource names and urls.
func GetRoot() (map[string]string, error) {
	var out map[string]string
	return out, Get("/", &out)
}

// getFilms fetches all films for each URL and returns a slice
// of Film types.
func getFilms(urls []string) (films []Film, err error) {
	for _, url := range urls {
		var f Film
		if err = Get(url, &f); err != nil {
			return
		}
		films = append(films, f)
	}

	return

}

// getPeople fetches all people for each URL and returns a slice
// of Person types
func getPeople(urls []string) (people []Person, err error) {
	for _, url := range urls {
		var p Person
		if err = Get(url, &p); err != nil {
			return
		}
		people = append(people, p)
	}

	return
}
