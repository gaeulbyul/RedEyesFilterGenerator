package cmd

import (
	"github.com/muesli/coral"
)

var rootCmd = &coral.Command{
	Use:   "FilterGenerator",
	Short: "Generate a BloomFilter file for the RedEyes or Shinigami-Eyes",
	Long:  `Generate a BloomFilter file for the RedEyes or Shinigami-Eyes`,
	Version: "v0.0.0.1",
}

func Execute() {
	coral.CheckErr(rootCmd.Execute())
}

func init() {}
