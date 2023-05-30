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
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Variables set from the command line
var outputFileName string
var topSize int
var months int
var isVerboseExtract bool

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract [input file]",
	Short: "Extracts the top submitters from the supplied pivot table",
	Long: `This command extract the top submitter for a given period (by default 12 months).
The input file is first validated before being processed.
If not specified, the output file name is hardcoded to "top-submitters.csv". 
The "months" parameter is the number of months used to compute the top users, 
counting from backwards from the last day.
The "topSize" parameter defines the number of users considered as top users. 
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if !isFileValid(args[0]) {
			return fmt.Errorf("Invalid file")
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

		if !extractData(args[0], outputFileName, topSize, months, isVerbose) {
			fmt.Print("Failed to extract data")
			os.Exit(1)
		}
	},
}

// Initialize the Cobra processor
func init() {
	rootCmd.AddCommand(extractCmd)

	// Here you will define your flags and configuration settings.
	extractCmd.PersistentFlags().StringVar(&outputFileName, "out", "top-submitters.csv", "Output file name")
	extractCmd.PersistentFlags().IntVar(&topSize, "topSize", 35, "Number of top submitters to extract.")
	extractCmd.PersistentFlags().IntVar(&months, "months", 12, "Accumulated number of months.")

	// checkCmd.PersistentFlags().BoolVar(&isVerboseExtract, "verboseExtract", false, "Displays useful info during the extraction")
}

// Extracts the top submitters for a given period and writes it to a file
func extractData(inputFilename string, outputFilename string, topSize int, months int, isVerboseExtract bool) bool {
	if isVerboseExtract {
		fmt.Printf("Extracting from \"%s\" the %d top submitters during the last %d months\n  and writing them to \"%s\"\n\n", inputFilename, topSize, months, outputFilename)
	}

	//At this stage of the processing, we assume that the input file is correctly formatted
	f, err := os.Open(inputFilename)
	if err != nil {
		log.Printf("Unable to read input file "+inputFilename+"\n", err)
		return false
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Printf("Unexpected error loading"+inputFilename+"\n", err)
		return false
	}

	firstDataColumn, lastDataColumn, oldestDate, mostRecentDate := getBoundaries(records, months)

	fmt.Printf("Accumulating data between %s and  %s (between columns %d and %d)\n", oldestDate, mostRecentDate, firstDataColumn, lastDataColumn)

	return true
}

func getBoundaries(records [][]string, months int) (startColumn int, endColumn int, startMonth string, endMonth string) {
	nbrOfColumns := len(records[0])
	endColumn = nbrOfColumns - 1

	if (months >= nbrOfColumns) {
		months = 0
	}

	if months == 0 {
		startColumn = 1
	} else {
		startColumn = nbrOfColumns - (months)
	}
	startMonth = records[0][startColumn]
	endMonth = records[0][endColumn]

	return startColumn, endColumn, startMonth, endMonth
}
