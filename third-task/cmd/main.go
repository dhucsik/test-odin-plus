package main

import (
	"log"

	"third-task/csvwriter"
	"third-task/parser"
)

func main() {
	influencers, err := parser.GetInfluencers()
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(csvwriter.WriteInfluencers(influencers))
}
