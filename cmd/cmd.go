package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "pokemon-go-colorscripts",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WIP")
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
	cmd.Flags().String("no-title", "", "Do not display pokemon name")
	cmd.Flags().StringP("shiny", "s", "", "Show the shiny version of a pokemon instead")
	cmd.Flags().BoolP("big", "b", false, "Show a larger version of the sprite")
	cmd.Flags().BoolP("random", "r", false, "display a random pokemon")
	cmd.Flags().
		StringP("random-by-names", "R", "", "Show a random pokemon chosen in the provided list of names. This list is in form (poke_1,poke_2,...,poke_n) only separated by commas WITHOUT whitespace (e.g. charmander,bulbasaur,squirtle)")
}
