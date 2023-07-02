package csvwriter

import (
	"encoding/csv"
	"os"
)

func WriteInfluencers(influencers [50][8]string) error {
	file, err := os.Create("influencers.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{"Rank", "Username", "Full name", "Category", "Followers", "Country", "Eng.(Auth.)", "Eng.(Avg.)"}

	writer.Write(headers)

	for _, influencer := range influencers {
		writer.Write(influencer[:])
	}

	defer writer.Flush()

	return nil
}
