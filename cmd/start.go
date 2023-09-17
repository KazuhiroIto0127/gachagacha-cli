/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/KazuhiroIto0127/gachagacha-cli/pkg/randomizer"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Print random art.",
	Run: func(cmd *cobra.Command, args []string) {
		rarity := randomizer.GetRarity()
		rbb, err := os.ReadFile("resources/aa/rarities/" + rarity + ".txt")
		if err != nil {
      fmt.Println(err)
    }
		fmt.Print(string(rbb))

		item, err := randomizer.GetRandomItem(rarity)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(item.Name)

		aa, err := os.ReadFile("resources/aa/items/" + item.Aa)
		if err != nil {
      fmt.Println(err)
    }
		fmt.Print(string(aa))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
