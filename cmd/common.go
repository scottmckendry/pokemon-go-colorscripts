package cmd

import "embed"

type pokemon struct {
	Name  string   `json:"name"`
	Forms []string `json:"forms"`
}

type flags struct {
	list    bool
	name    string
	form    string
	noTitle bool
	shiny   bool
	big     bool
	random  bool
	version bool
}

var pokemon_map map[string]pokemon
var pokemon_list []pokemon
var shinyrate float64 = 0.01

//go:embed pokemon-colorscripts/colorscripts/*
var colorscripts embed.FS

//go:embed pokemon-colorscripts/pokemon.json
var pokemon_json []byte
