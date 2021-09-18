package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getWinnerTotalGoals' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING competition
 *  2. INTEGER year
 */

const (
	COMP_URL  = "https://jsonmock.hackerrank.com/api/football_competitions?name=%s&year=%d"
	MATCH_URL = "https://jsonmock.hackerrank.com/api/football_matches?competition=%s&year=%d&team%d=%s&page=%d"
)

type CompetitionData struct {
	Name     string  `json:"name,omitempty"`
	Country  string  `json:"country,omitempty"`
	Year     float64 `json:"year,omitempty"`
	Winner   string  `json:"winner,omitempty"`
	Runnerup string  `json:"runnerup,omitempty"`
}

type Match struct {
	Competition string  `json:"competition"`
	Year        float64 `json:"year"`
	Round       string  `json:"round"`
	Team1       string  `json:"team1,omitempty"`
	Team2       string  `json:"team2,omitempty"`
	Team1Goals  string  `json:"team1goals"`
	Team2Goals  string  `json:"team2goals"`
}

type Matches []Match
type Competitions []CompetitionData

func getWinnerTotalGoals(competition string, year int32) int32 {
	r, err := http.Get(fmt.Sprintf(COMP_URL, url.QueryEscape(competition), year))
	if err != nil {
		log.Printf("Request %v", err)
		return 0
	}
	obj := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		log.Printf("Decoder %v", err)
		return 0
	}
	r.Body.Close()

	jsonComps, err := json.Marshal(obj["data"])
	if err != nil {
		return 0
	}
	comps := &Competitions{}
	err = json.Unmarshal(jsonComps, &comps)
	if err != nil {
		return 0
	}
	var comp CompetitionData

	for _, c := range *comps {
		if c.Name == competition && int32(c.Year) == year {
			comp = c
		}
	}

	goals := 0
	pages := 1
	for i := 1; i <= 2; i++ {
		for {
			r, err := http.Get(fmt.Sprintf(MATCH_URL, url.QueryEscape(competition), year, i, url.QueryEscape(comp.Winner), pages))
			if err != nil {
				log.Print(err)
				return 0
			}

			obj := make(map[string]interface{})

			if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
				log.Printf("Decoder %v", err)
				return 0
			}
			r.Body.Close()

			totalPages := int(obj["total_pages"].(float64))
			jsonMatches, err := json.Marshal(obj["data"])
			if err != nil {
				return 0
			}
			matches := &Matches{}
			err = json.Unmarshal(jsonMatches, &matches)

			for _, m := range *matches {
				var n int
				if i == 1 {
					n, _ = strconv.Atoi(m.Team1Goals)
				} else {
					n, _ = strconv.Atoi(m.Team2Goals)
				}
				goals += n
			}

			if totalPages > pages {
				pages++
			} else {
				break
			}
		}
	}

	return int32(goals)
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	competition := readLine(reader)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := getWinnerTotalGoals(competition, year)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
