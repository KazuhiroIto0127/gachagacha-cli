/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/KazuhiroIto0127/gachagacha-cli/pkg/randomizer"
	"github.com/KazuhiroIto0127/gachagacha-cli/pkg/utils"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This command allows you to pull the gacha. The rarity levels are SSR, SR, and R.",
	Run: func(cmd *cobra.Command, args []string) {
		rarity := randomizer.GetRarity()
		fmt.Println(utils.AaFromFile("rarities/" + rarity + ".txt"))

		item, err := randomizer.GetRandomItem(rarity)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(item.Name)

		fmt.Println(utils.AaFromFile("items/" + strings.ToLower(rarity) + "/" + item.Aa))
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
