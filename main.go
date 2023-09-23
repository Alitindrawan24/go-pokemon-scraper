package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

const (
	// Target url fpr scrapper
	targetUrl string = "https://www.serebii.net/pokemon/nationalpokedex.shtml"

	// Output filename
	output string = "pokemon.json"
)

func main() {

	c := colly.NewCollector()
	var pokemons []Pokemon

	// Create json file
	defer func() {
		file, _ := json.MarshalIndent(pokemons, "", "\t")
		_ = os.WriteFile(output, file, 0644)
	}()

	c.OnHTML("table.dextable > tbody > tr", func(e *colly.HTMLElement) {
		row := e.DOM
		pokemon := newPokemon(row)
		pokemon.scraps()

		if pokemon.ID != "" {
			pokemons = append(pokemons, pokemon)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.Visit(targetUrl)
}
