package main

import (
	"fmt"

	"github.com/fioncat/tergame-2048/game"
	"github.com/spf13/cobra"
)

var cfg game.Config

func init() {
	cmd.PersistentFlags().IntVarP(&cfg.Size, "size", "s", 4, "area size.")
	cmd.PersistentFlags().IntSliceVarP(&cfg.Options, "options", "o", []int{2, 4}, "range of generated numbers.")
	cmd.PersistentFlags().IntVarP(&cfg.Grow, "grow", "g", 1, "how many numbers are generated each step.")
}

var cmd = &cobra.Command{
	Use: "ter2048",

	Short: "2048 game in terminal.",
	Long:  "2048 game in terminal.",

	Run: func(cmd *cobra.Command, args []string) {
		if cfg.Size < 4 {
			cfg.Size = 4
		}
		if len(cfg.Options) == 0 {
			cfg.Options = []int{2, 4}
		}
		if cfg.Grow <= 0 {
			cfg.Grow = 1
		}
		game.Start(&cfg)
	},
}

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("exec failed: %v\n", err)
	}
}
