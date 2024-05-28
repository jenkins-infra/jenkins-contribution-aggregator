/*
Copyright © 2024 Jean-Marc Meessen jean-marc@meessen-web.org

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
	"path"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func plotAllHistoryFiles(plotDirectory string, historicDataSlice [][]string) error {

	var header []string

	for row, historyRow := range historicDataSlice {
		if row == 0 {
			header = historyRow[1:]
		} else {
			err := plot_bargraph(plotDirectory, historyRow[0], header, historyRow[1:])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// TODO: add type for legends
// TODO: how is the data passed so that it can be formatted
// TODO: add parameter to limit the size of the data displayed
// Plots the passed data in a png file named after the user in the specified directory
func plot_bargraph(plotDirectory string, name string, xLabels []string, values []string) error {


	//FIXME: type of graph (PR or Comments)

	p := plot.New()

	//In case of a compare, the name is appended with "new" or "churned". So we need to clean it up
	name_element := strings.Split(name, " ")
	cleanedName := name_element[0]

	plotFileName := path.Join(plotDirectory, cleanedName+".png")

	p.Title.Text = "Submissions by " + cleanedName
	p.Y.Label.Text = "Count"

	w := vg.Points(20)

	floatValues, err := convertValuesToInts(values)
	if err != nil {
		return err
	}
	groupA := plotter.Values(floatValues)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		return err
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)

	p.Add(barsA)
	simplifiedLabels := simplifyAxisLabels(xLabels)

	p.NominalX(simplifiedLabels...)

	return p.Save(10*vg.Inch, 6*vg.Inch, plotFileName)
}

// take the list of months and transforms this to a lighter list that can be displayed on the graph
func simplifyAxisLabels(inputLabels []string) []string {
	var outputLabels []string
	currentYear := ""

	for _, oldLabel := range inputLabels {
		//Labels come in the form YYYY-MM
		splittedLabel := strings.Split(oldLabel, "-")
		labelsYear := splittedLabel[0]
		if labelsYear != currentYear {
			outputLabels = append(outputLabels, labelsYear)
			currentYear = labelsYear
		} else {
			outputLabels = append(outputLabels, "")
		}
	}

	return outputLabels
}

// Converts a slice of numerical strings into a slice of floats.
func convertValuesToInts(stringValues []string) ([]float64, error) {
	var floatValues []float64

	for _, stringValue := range stringValues {
		value, err := strconv.ParseFloat(strings.TrimSpace(stringValue), 64)
		if err != nil {
			return nil, fmt.Errorf("Unexpected error converting data to int (%v)", err)
		} else {
			floatValues = append(floatValues, value)
		}
	}
	return floatValues, nil
}
