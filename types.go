package main

import "github.com/PuerkitoBio/goquery"

type Pokemon struct {
	Row       *goquery.Selection `json:"-"`
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Sprite    string             `json:"sprite"`
	Types     []Type             `json:"types"`
	Abilities []Ability          `json:"abilities"`
	BaseStats BaseStats          `json:"base_stats"`
}

type Type struct {
	Name string `json:"name"`
	Gif  string `json:"gif"`
}

type Ability struct {
	Name string `json:"name"`
}

type BaseStats struct {
	HP             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"special_attack"`
	SpecialDefense int `json:"special_defense"`
	Speed          int `json:"speed"`
}
