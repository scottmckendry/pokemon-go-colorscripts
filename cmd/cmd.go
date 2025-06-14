package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

const version = "0.2.0" // x-release-please-version

var cmd = &cobra.Command{
	Use: "pokemon-go-colorscripts",

	Run: func(cmd *cobra.Command, args []string) {
		flags := flags{
			list:    mustGetBool(cmd, "list"),
			name:    mustGetString(cmd, "name"),
			form:    mustGetString(cmd, "form"),
			noTitle: mustGetBool(cmd, "no-title"),
			shiny:   mustGetBool(cmd, "shiny"),
			big:     mustGetBool(cmd, "big"),
			random:  mustGetBool(cmd, "random"),
			version: mustGetBool(cmd, "version"),
		}

		if flags.version {
			fmt.Printf("pokemon-go-colorscripts version %s\n", version)
		}

		loadPokemon()

		if flags.list {
			listPokemon()
		}

		if flags.name != "" {
			showPokemonByName(flags)
		}

		if flags.random {
			showRandomPokemon(flags)
		}

		cmd.Help()
		os.Exit(1)
	},
}

func Execute() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cmd.Flags().BoolP("list", "l", false, "list all available pokemon")
	cmd.Flags().StringP("form", "f", "", "Show an alternate form of a pokemon")
	cmd.Flags().Bool("no-title", false, "Do not display pokemon name")
	cmd.Flags().BoolP("shiny", "s", false, "Show the shiny version of a pokemon instead")
	cmd.Flags().BoolP("big", "b", false, "Show a larger version of the sprite")
	cmd.Flags().BoolP("random", "r", false, "display a random pokemon")
	cmd.Flags().BoolP("version", "v", false, "display version information")
	cmd.Flags().
		StringP("name", "n", "", "Select a pokemon by name. Generally spelled like in the games.a few exceptions are nidoran-f, nidoran-m, mr-mime, farfetchd, flabebe type-null etc. Perhaps grep the output of --list if in doubt.")
}

// Unmarshall the pokemon.json file into a slice of Pokemon structs and a map of Pokemon structs for easy lookup
func loadPokemon() {
	json.Unmarshal(pokemon_json, &pokemon_list)
	pokemon_map = make(map[string]pokemon)
	for _, pokemon := range pokemon_list {
		pokemon_map[pokemon.Name] = pokemon
	}
}

// Print a list of all available pokemon
func listPokemon() {
	for _, pokemon := range pokemon_list {
		fmt.Println(pokemon.Name)
	}
	os.Exit(0)
}

// Show a pokemon by name
func showPokemonByName(flags flags) {
	sizeSubdir := "small"
	if flags.big {
		sizeSubdir = "large"
	}

	pokemon, ok := pokemon_map[flags.name]
	if !ok {
		fmt.Printf("Invald pokemon %s\n", flags.name)
		os.Exit(1)
	}

	shinySubdir := "regular"
	if flags.shiny {
		shinySubdir = "shiny"
	}

	if flags.form != "" {
		if slices.Contains(pokemon.Forms, flags.form) {
			flags.name += "-" + flags.form
		} else {
			if len(pokemon.Forms) == 1 {
				fmt.Printf("No alternate forms available for %s\n", flags.name)
			} else {
				fmt.Printf("Avaliable forms for %s are:\n", flags.name)
				for _, form := range pokemon.Forms {
					fmt.Println(form)
				}
			}
			os.Exit(1)
		}
	}

	if !flags.noTitle {
		if flags.shiny {
			fmt.Println(flags.name + " (shiny)")
		} else {
			fmt.Println(flags.name)
		}
	}

	path := fmt.Sprintf(
		"pokemon-colorscripts/colorscripts/%s/%s/%s",
		sizeSubdir,
		shinySubdir,
		flags.name,
	)

	pokemon_file, _ := colorscripts.ReadFile(path)
	fmt.Print(string(pokemon_file))
	os.Exit(0)
}

// Show a random pokemon
func showRandomPokemon(flags flags) {
	flags.name = pokemon_list[rand.Intn(len(pokemon_list))].Name
	flags.form = pokemon_map[flags.name].Forms[rand.Intn(len(pokemon_map[flags.name].Forms))]
	if flags.form == "regular" {
		flags.form = ""
	}

	// if shiny flag is not passed, set a small random chance for the pokemon to be shiny
	if !flags.shiny && rand.Float64() < shinyrate {
		flags.shiny = true
	}
	showPokemonByName(flags)
}

// Get a boolean flag from a command, and exit if it fails
func mustGetBool(cmd *cobra.Command, name string) bool {
	val, err := cmd.Flags().GetBool(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return val
}

// Get a string flag from a command, and exit if it fails
func mustGetString(cmd *cobra.Command, name string) string {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return val
}
