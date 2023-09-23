package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	// Hostname for base url of sprite/type image
	host string = "https://www.serebii.net"
)

// newPokemon creates a new Pokemon struct from a row of data.
// The row of data to use for populating the Pokemon.
// Return uhe newly created Pokemon struct
func newPokemon(row *goquery.Selection) Pokemon {
	return Pokemon{
		Row: row,
	}
}

// scraps sets the pokemon attribute
// return he pokemon to be
func (pokemon *Pokemon) scraps() *Pokemon {
	pokemon.setID()
	pokemon.setSprite()
	pokemon.setName()
	pokemon.setTypes()
	pokemon.setAbilities()
	pokemon.setBaseStats()

	return pokemon
}

// Set pokemon ID
func (pokemon *Pokemon) setID() {
	id := strings.TrimSpace(pokemon.Row.Find("td.fooinfo:nth-child(1)").Text())
	pokemon.ID = id
}

// Set pokemon Sprite
func (pokemon *Pokemon) setSprite() {
	sprite, _ := pokemon.Row.Find("td.fooinfo:nth-child(2) img").Attr("src")
	pokemon.Sprite = host + sprite
}

// Set pokemon Name
func (pokemon *Pokemon) setName() {
	name := pokemon.Row.Find("td.fooinfo:nth-child(3) a").Text()
	pokemon.Name = name
}

// Set pokemon Types
func (pokemon *Pokemon) setTypes() {
	types := []Type{}
	typeDOM := pokemon.Row.Find("td.fooinfo:nth-child(4) a img").Nodes
	for _, node := range typeDOM {
		typeNode := node.Attr
		typeUrl := typeNode[0].Val
		typeName, _ := SplitLastString(typeUrl)
		types = append(types, Type{
			Name: typeName,
			Gif:  host + typeUrl,
		})
	}

	pokemon.Types = types
}

// Set pokemon Abilities
func (pokemon *Pokemon) setAbilities() {
	abilities := []Ability{}
	abilityDOM := pokemon.Row.Find("td.fooinfo:nth-child(5) a").Nodes
	for _, node := range abilityDOM {
		abilityNode := node.Attr
		ability := abilityNode[0].Val
		abilityName, _ := SplitLastString(ability)
		abilities = append(abilities, Ability{
			Name: abilityName,
		})
	}

	pokemon.Abilities = abilities
}

// Set pokemon Base Stats
func (pokemon *Pokemon) setBaseStats() {
	hp, _ := strconv.Atoi(pokemon.Row.Find("td.fooinfo:nth-child(6)").Text())
	attack, _ := strconv.Atoi(pokemon.Row.Find("td.fooinfo:nth-child(7)").Text())
	defense, _ := strconv.Atoi(pokemon.Row.Find("td.fooinfo:nth-child(8)").Text())
	specialAttack, _ := strconv.Atoi(pokemon.Row.Find("td.fooinfo:nth-child(9)").Text())
	specialDefense, _ := strconv.Atoi(pokemon.Row.Find("td.fooinfo:nth-child(10)").Text())
	speed, _ := strconv.Atoi(pokemon.Row.Find("td.fooinfo:nth-child(11)").Text())
	baseStats := BaseStats{
		HP:             hp,
		Attack:         attack,
		Defense:        defense,
		SpecialAttack:  specialAttack,
		SpecialDefense: specialDefense,
		Speed:          speed,
	}

	pokemon.BaseStats = baseStats
}
