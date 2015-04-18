package apihack

import (
	"fmt"
	"log"
	"net/http"
)

func Matches(league string) (*http.Response, error) {
	url := fmt.Sprintf("%s/leagues/%s/fixtures", baseURI, league)
	log.Printf("Requesting API Hack: %s", url)

	return http.Get(url)
}
