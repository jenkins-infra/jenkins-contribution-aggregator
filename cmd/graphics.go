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

// TODO: add type for legends
// TODO: how is the data passed so that it can be formatted
// TODO: add parameter to limit the size of the data displayed
// Plots the passed data in a png file named after the user in the specified directory
func plot_bargraph(plotDirectory string, name string, xLabels []string, values []string) error {
	plotFileName := path.Join(plotDirectory, name+".png")

	//FIXME: type of graph (PR or Comments)

	p := plot.New()

	// x_legend := "2020-01,2020-02,2020-03,2020-04,2020-05,2020-06,2020-07,2020-08,2020-09,2020-10,2020-11,2020-12,2021-01,2021-02,2021-03,2021-04,2021-05,2021-06,2021-07,2021-08,2021-09,2021-10,2021-11,2021-12,2022-01,2022-02,2022-03,2022-04,2022-05,2022-06,2022-07,2022-08,2022-09,2022-10,2022-11,2022-12,2023-01,2023-02,2023-03,2023-04,2023-05,2023-06,2023-07,2023-08,2023-09,2023-10,2023-11,2023-12,2024-01,2024-02,2024-03,2024-04"

	p.Title.Text = "Submissions by " + name
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

	return p.Save(5*vg.Inch, 3*vg.Inch, plotFileName)
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
