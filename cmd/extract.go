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
	"sort"
	"strconv"

	"github.com/spf13/cobra"
)

// Variables set from the command line
var outputFileName string
var topSize int
var period int
var isVerboseExtract bool

type totalized_record struct {
	User string //Submitter name
	Pr   int    //Number of PRs
}

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract [input file]",
	Short: "Extracts the top submitters from the supplied pivot table",
	Long: `This command extract the top submitter for a given period (by default 12 months).
This interval is counted, by default, from the last month available in the pivot table.
The input file is first validated before being processed.
If not specified, the output file name is hardcoded to "top-submitters.csv". 

The "months" parameter is the number of months used to compute the top users, 
counting from backwards from the last month. If a 0 months is specified, all the 
available is counted.

The "topSize" parameter defines the number of users considered as top users.
If more submitters with the same amount of total PRs exist ("ex aequo"), they are included in 
the list (resulting in more thant the specified number of top users).  
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

		if !extractData(args[0], outputFileName, topSize, period, isVerboseExtract) {
			fmt.Print("Failed to extract data")
			os.Exit(1)
		}
	},
}

// Initialize the Cobra processor
func init() {
	rootCmd.AddCommand(extractCmd)

	// definition of flags and configuration settings.
	extractCmd.PersistentFlags().StringVarP(&outputFileName, "out", "o", "top-submitters.csv", "Output file name")
	extractCmd.PersistentFlags().IntVarP(&topSize, "topSize", "t", 35, "Number of top submitters to extract.")
	extractCmd.PersistentFlags().IntVarP(&period, "period", "p", 12, "Number of months to accumulate.")

	extractCmd.PersistentFlags().BoolVarP(&isVerboseExtract, "verbose", "v", false, "Displays useful info during the extraction")
}



// Extracts the top submitters for a given period and writes it to a file
func extractData(inputFilename string, outputFilename string, topSize int, period int, isVerboseExtract bool) bool {
	if isVerboseExtract {
		fmt.Printf("Extracting from \"%s\" the %d top submitters during the last %d months\n  and writing them to \"%s\"\n\n", inputFilename, topSize, period, outputFilename)
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

	firstDataColumn, lastDataColumn, oldestDate, mostRecentDate := getBoundaries(records, period)

	fmt.Printf("Accumulating data between %s and  %s (columns %d and %d)\n",
		oldestDate, mostRecentDate, firstDataColumn, lastDataColumn)

	//Slice that will contain all the totalized records	
	var new_output_slice []totalized_record

	for i, dataLine := range records {

		//Skip header line as it has already been checked
		if i == 0 {
			continue
		}

		recordTotal := 0
		for ii, column := range dataLine {
			if ii >= firstDataColumn && ii <= lastDataColumn {
				// fmt.Printf(", %s", column)

				// We don't treat conversion errors or negative values as the file has already been checked
				columnValue, _ := strconv.Atoi(column)
				recordTotal = recordTotal + columnValue
			}
		}

		//Add the total to the full list
		a_totalized_record := totalized_record{dataLine[0], recordTotal}
		new_output_slice = append(new_output_slice, a_totalized_record)
	}

	// Sort the slice, based on the number of PRs, in descending order
	sort.Slice(new_output_slice, func(i, j int) bool { return new_output_slice[i].Pr > new_output_slice[j].Pr })


	//Loop through list to find the top submitters (and ex-aequo) to load the final list
	current_total := 0
	isListComplete := false
	
	var csv_output_slice [][]string
	header_row := []string {"Submitter", "Total_PRs"}
	csv_output_slice = append(csv_output_slice, header_row)
	for i, total_record := range new_output_slice {
		if i < topSize {
			current_total = total_record.Pr

			var work_row []string
			work_row = append(work_row, total_record.User, strconv.Itoa(total_record.Pr))
			csv_output_slice = append(csv_output_slice, work_row)
		} else {
			if !isListComplete {
				if current_total == total_record.Pr {
					//This is an ex-aequo, so add it to the list
					var work_row []string
					work_row = append(work_row, total_record.User, strconv.Itoa(total_record.Pr))
					csv_output_slice = append(csv_output_slice, work_row)
				} else {
					// we have all we need
					isListComplete = true
				}
			}
		}
	}

	//Open output file
	out, err := os.Create(outputFilename)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

	//Write the collected data as a CSV file
	csv_out := csv.NewWriter(out)
	write_err := csv_out.WriteAll(csv_output_slice)
    if write_err != nil {
        log.Fatal(err)
    }
	csv_out.Flush()

	return true
}

// Based on the number of months requested, computes the start/end column and associated date for the given dataset
func getBoundaries(records [][]string, period int) (startColumn int, endColumn int, startMonth string, endMonth string) {
	nbrOfColumns := len(records[0])
	endColumn = nbrOfColumns - 1

	if period >= nbrOfColumns {
		period = 0
	}

	if period == 0 {
		startColumn = 1
	} else {
		startColumn = nbrOfColumns - (period)
	}
	startMonth = records[0][startColumn]
	endMonth = records[0][endColumn]

	return startColumn, endColumn, startMonth, endMonth
}
