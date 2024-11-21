package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "pokemon-go-colorscripts",

	Run: func(cmd *cobra.Command, args []string) {
		flags := flags{
			list:          mustGetBool(cmd, "list"),
			name:          mustGetString(cmd, "name"),
			form:          mustGetString(cmd, "form"),
			noTitle:       mustGetBool(cmd, "no-title"),
			shiny:         mustGetBool(cmd, "shiny"),
			big:           mustGetBool(cmd, "big"),
			random:        mustGetBool(cmd, "random"),
			randomByNames: mustGetString(cmd, "random-by-names"),
		}

		loadPokemon()

		if flags.list {
			listPokemon()
		}

		size := "small"
		if flags.big {
			size = "large"
		}

		parent := "regular"
		if flags.shiny {
			parent = "shiny"
		}

		file_content, err := colorscripts.ReadFile(
			"pokemon-colorscripts/colorscripts/" + size + "/" + parent + "/" + flags.name,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if !flags.noTitle {
			fmt.Println(flags.name)
		}
		fmt.Print(string(file_content))
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
	cmd.Flags().
		StringP("name", "n", "", "Select a pokemon by name. Generally spelled like in the games.a few exceptions are nidoran-f, nidoran-m, mr-mime, farfetchd, flabebe type-null etc. Perhaps grep the output of --list if in doubt.")
	cmd.Flags().StringP("form", "f", "", "Show an alternate form of a pokemon")
	cmd.Flags().Bool("no-title", false, "Do not display pokemon name")
	cmd.Flags().BoolP("shiny", "s", false, "Show the shiny version of a pokemon instead")
	cmd.Flags().BoolP("big", "b", false, "Show a larger version of the sprite")
	cmd.Flags().BoolP("random", "r", false, "display a random pokemon")
	cmd.Flags().
		StringP("random-by-names", "R", "", "Show a random pokemon chosen in the provided list of names. This list is in form (poke_1,poke_2,...,poke_n) only separated by commas WITHOUT whitespace (e.g. charmander,bulbasaur,squirtle)")
}

// Print a list of all available pokemon
func listPokemon() {
	for _, pokemon := range pokemon_list {
		fmt.Println(pokemon.Name)
	}
	os.Exit(0)
}

// Unmarshall the pokemon.json file into a slice of Pokemon structs and a map of Pokemon structs for easy lookup
func loadPokemon() {
	json.Unmarshal(pokemon_json, &pokemon_list)
	pokemon_map = make(map[string]pokemon)
	for _, pokemon := range pokemon_list {
		pokemon_map[pokemon.Name] = pokemon
	}
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
