package cmd

import "embed"

type pokemon struct {
	Name  string   `json:"name"`
	Forms []string `json:"forms"`
}

type flags struct {
	list          bool
	name          string
	form          string
	noTitle       bool
	shiny         bool
	big           bool
	random        bool
	randomByNames string
}

var generations = map[string][2]int{
	"1": {1, 151},
	"2": {152, 251},
	"3": {252, 386},
	"4": {387, 493},
	"5": {494, 649},
	"6": {650, 721},
	"7": {722, 809},
	"8": {810, 898},
}

var pokemon_map map[string]pokemon
var pokemon_list []pokemon
var shinyrate = 1 / 128

//go:embed pokemon-colorscripts/colorscripts/*
var colorscripts embed.FS

//go:embed pokemon-colorscripts/pokemon.json
var pokemon_json []byte
