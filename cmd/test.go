package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/gaeulbyul/RedEyesFilterGenerator/bloomfilter"
	"github.com/muesli/coral"
)

var testCmd = &coral.Command{
	Use:   "test [filter file] [...identifiers]",
	Short: "Test given identifier are included in bloomFilter data.",
	Long:  `Test given identifier are included in bloomFilter data.`,
	Run:   bloomFilterTestIdentifier,
	Args: func(cmd *coral.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("please specity one or more identifiers.\n")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func bloomFilterTestIdentifier(cmd *coral.Command, args []string) {
	bloomFilterFile, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	bf := bloomfilter.NewFromBytes(bloomFilterFile, 20)
	for _, argItem := range args[1:] {
		argItemAsBytes := []byte(argItem)
		testResult := bf.Test(argItemAsBytes)
		fmt.Printf("%s = %t\n", argItem, testResult)
	}
}
