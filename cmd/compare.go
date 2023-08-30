/*
Copyright © 2023 Jean-Marc Meessen jean-marc@meessen-web.org

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var compareWith int

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compares two top Submitters extractions to show \"churned\" or \"new\" submitters.",
	Long: `The COMPARE command will will extract a the Top Submitters as with the EXTRACT command and than
compare it with an extraction with the same settings but with an X amount of months before`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if !isFileValid(args[0]) {
			return fmt.Errorf("Invalid file\n")
		}
		if !isValidMonth(endMonth, isVerboseExtract) {
			return fmt.Errorf("Invalid month\n")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// When called standalone, we want to give the minimal information
		isSilent := true

		if !checkFile(args[0], isSilent) {
			fmt.Print("Invalid input file.")
			os.Exit(1)
		}

		if outputFileName == "top-submitters_YYYY-MM.csv" {
			outputFileName = "top-submitters_" + strings.ToUpper(endMonth) + ".csv"
		}

		// if !extractData(args[0], outputFileName, topSize, endMonth, period, isVerboseExtract) {
		// 	fmt.Print("Failed to extract data")
		// 	os.Exit(1)
		// }
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)

	// Here you will define your flags and configuration settings.
	compareCmd.PersistentFlags().StringVarP(&outputFileName, "out", "o", "top-submitters_YYYY-MM.csv", "Output file name.")
	compareCmd.PersistentFlags().IntVarP(&topSize, "topSize", "t", 35, "Number of top submitters to extract.")
	compareCmd.PersistentFlags().IntVarP(&period, "period", "p", 12, "Number of months to accumulate.")
	compareCmd.PersistentFlags().IntVarP(&compareWith, "compare", "c", 3, "Number of months back to compare with.")
	compareCmd.PersistentFlags().StringVarP(&endMonth, "month", "m", "latest", "Month to extract top submitters.")

	compareCmd.PersistentFlags().BoolVarP(&isVerboseExtract, "verbose", "v", false, "Displays useful info during the extraction")
}
