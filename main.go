package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/basbeu/JudobaseToWiki/internal"
	"github.com/basbeu/JudobaseToWiki/judobase"
)

func main() {
	input := flag.String("input", "./input.json", "input json file from judobase.ijf.org")
	output := flag.String("output", "./output.txt", "output text file containing the result in a wikipedia format")

	flag.Parse()

	jsonFile, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	var judobaseComp judobase.Competition
	err = json.Unmarshal(byteValue, &judobaseComp)
	if err != nil {
		log.Fatal(err)
		return
	}
	competition := internal.NewCompetition(judobaseComp)
	formatter := internal.CompetitionFrenchWikiFormatter{}
	out := formatter.Format(competition)

	os.WriteFile(*output, []byte(out), 0666)
}
