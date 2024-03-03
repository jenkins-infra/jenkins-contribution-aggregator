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
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Validates that the input file is a real file (and not a directory)
func isFileValid(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// validates whether  the month parameter has the correct format ("YYYY-MM" or "latest")
func isValidMonth(month string, isVerbose bool) bool {
	if month == "" {
		if isVerbose {
			fmt.Print("Empty month\n")
		}
		return false
	}
	if strings.ToUpper(month) == "LATEST" {
		return true
	}

	regexpMonth := regexp.MustCompile(`20[12][0-9]-(0[1-9]|1[0-2])`)
	if !regexpMonth.MatchString(month) {
		if isVerbose {
			fmt.Printf("Supplied data (%s) is not in a valid month format. Should be \"YYYY-MM\" and later than 2010\n", month)
		}
		return false
	}

	return true
}

// Write the string slice to a file formatted as a CSV
func writeCSVtoFile(outputFileName string, csv_output_slice [][]string) {
	//Open output file
	out, err := os.Create(outputFileName)
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
}

// returns true if the file extention is .md.
// It returns false in other cases, thus assuming a CSV output
func isWithMDfileExtension(filename string) bool {
	extension := filepath.Ext(filename)
	if strings.ToLower(extension) == ".md" {
		return true
	} else {
		return false
	}
}

// Writes the data as Markdown
func writeDataAsMarkdown(outputFileName string, output_data_slice [][]string) {
	//Open output file
	f, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	out := bufio.NewWriter(f)

	fmt.Fprintf(out, "%s\n", "# Extract")
	fmt.Fprintf(out, " \n")
	fmt.Fprintf(out, " \n")
	out.Flush()
}
