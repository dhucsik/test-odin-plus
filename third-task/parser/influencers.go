package parser

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://hypeauditor.com/top-instagram-all-russia/"

func GetInfluencers() ([50][8]string, error) {
	var influencers [50][8]string

	client := GetHttpClient()

	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return influencers, err
	}

	res, err := client.Do(request)
	if err != nil {
		return influencers, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return influencers, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return influencers, err
	}

	doc.Find("div.ranking-card").Find("div.row").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i > 50 {
			return false
		}

		if i == 0 {
			return true
		}

		influencers[i-1][0] = s.Find("div.rank").Find("span").First().Text()
		influencers[i-1][1] = s.Find("div.contributor__name-content").Text()
		influencers[i-1][2] = s.Find("div.contributor__title").Text()
		influencers[i-1][3] = s.Find("div.category").Find("div.ellipsis").First().Text()
		influencers[i-1][4] = s.Find("div.subscribers").Text()
		influencers[i-1][5] = s.Find("div.audience").Text()
		influencers[i-1][6] = s.Find("div.authentic").Text()
		influencers[i-1][7] = s.Find("div.engagement").Text()

		return true
	})

	return influencers, nil
}
