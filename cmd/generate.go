package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gaeulbyul/RedEyesFilterGenerator/bloomfilter"
	"github.com/muesli/coral"
)

var generateCmd = &coral.Command{
	Use:   "generate -i [input file] -o [filter file]",
	Short: "Generate bloomfilter file.",
	Long:  `Generate bloomfilter file from the identifiers.`,
	Run:   bloomFilterGenerate,
	// Args: func(cmd *cobra.Command, args []string) error {},
}

var inputfile string
var outputfile string

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&inputfile, "input", "i", "", "A path to text file that contains identifiers.")
	generateCmd.Flags().StringVarP(&outputfile, "output", "o", "", "A path to generated bloom-filter file.")
	generateCmd.MarkFlagRequired("input")
	generateCmd.MarkFlagRequired("output")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func bloomFilterGenerate(cmd *coral.Command, args []string) {
	fmt.Printf("Reading from '%s'...\n", inputfile)
	inputTextFile, err := ioutil.ReadFile(inputfile)
	if err != nil {
		panic(err)
	}
	inputText := string(inputTextFile)
	splitted := strings.Split(inputText, "\n")
	lenOfInput := len(splitted)
	m := bloomfilter.EstimateParameters(lenOfInput, 1.0e-6)
	fmt.Printf("Generating... (len = %d, m = %d)\n", lenOfInput, m)
	bf := bloomfilter.New(m, 20)
	before := time.Now().UnixMilli()
	for i, line := range splitted {
		identifier := strings.ToLower(strings.TrimSpace(line))
		if identifier == "" {
			continue
		}
		if strings.HasPrefix(identifier, "#") || strings.HasPrefix(identifier, " ") || strings.HasPrefix(identifier, "\t") {
			continue
		}
		count := i + 1
		if count%5000 == 0 {
			fmt.Printf("\r[%d / %d]", count, lenOfInput)
		}
		bf.Add([]byte(identifier))
	}
	fmt.Printf("\r[%d / %d]", lenOfInput, lenOfInput)
	after := time.Now().UnixMilli()
	fmt.Printf("\nWriting to '%s'...\n", outputfile)
	if err := ioutil.WriteFile(outputfile, bf.ToBytes(), 0644); err != nil {
		panic(err)
	}
	difftime := after - before
	fmt.Printf("DONE (%d ms)", difftime)
}
