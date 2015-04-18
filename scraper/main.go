package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type chant struct {
	Title    string `json:"title"`
	Lyric    string `json:"lyric"`
	AudioURL string `json:"audioURL"`
}

const baseURL = "http://fanchants.co.uk/football-team/"

func SanitiseLyric(l string) string {
	l = strings.TrimSpace(l)
	return l
}

func ScrapeChant(url string) *chant {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var c chant
	c.Title = doc.Find("div#heading-txt > h1 > span").Last().Text()
	c.Lyric = SanitiseLyric(doc.Find("div.lyrics > p").Text())
	c.AudioURL, _ = doc.Find("div#player-audio audio source[type='audio/mp3']").Attr("src")

	return &c
}

func ScrapeChants(team string) {
	doc, err := goquery.NewDocument(baseURL + team)
	if err != nil {
		log.Fatal(err)
	}

	var nPages int
	nPages, err = strconv.Atoi(doc.Find("span.paginator-centre a").Last().Text())
	if err != nil {
		log.Fatal(err)
	}

	if nPages > 3 {
		nPages = 3
	}

	log.Printf("%d number of pages", nPages)
	totalChants := 0
	chants := make([]*chant, 21*nPages, 30*nPages)

	for i := 2; i < nPages; i++ {
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("a.audio").Each(func(i int, s *goquery.Selection) {
			if url, exist := s.Attr("href"); exist {
				chants[totalChants] = ScrapeChant(url)
				totalChants++
			}
		})

		doc, err = goquery.NewDocument(baseURL + team + "/" + strconv.Itoa(i))
	}

	chants = chants[0:totalChants]

	var chantsData []byte
	if chantsData, err = json.Marshal(chants); err != nil {
		log.Fatal(err)
	}

	var f *os.File
	if f, err = os.Create(team + ".json"); err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	f.Write(chantsData)
}

func main() {
	ScrapeChants("liverpool")
}
