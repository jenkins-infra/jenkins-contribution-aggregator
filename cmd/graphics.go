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
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// TODO: add type for legends
// TODO: the name of the contributor is the file name. Just specify the output directory
// TODO: how is the data passed so that it can be formatted
// TODO: add parameter to limit the size of the data displayed 
func plot_bargraph(plotFileName string) error {
	p := plot.New()

	// x_legend := "2020-01,2020-02,2020-03,2020-04,2020-05,2020-06,2020-07,2020-08,2020-09,2020-10,2020-11,2020-12,2021-01,2021-02,2021-03,2021-04,2021-05,2021-06,2021-07,2021-08,2021-09,2021-10,2021-11,2021-12,2022-01,2022-02,2022-03,2022-04,2022-05,2022-06,2022-07,2022-08,2022-09,2022-10,2022-11,2022-12,2023-01,2023-02,2023-03,2023-04,2023-05,2023-06,2023-07,2023-08,2023-09,2023-10,2023-11,2023-12,2024-01,2024-02,2024-03,2024-04"

	p.Title.Text = "Submissions by XYZ"
	p.Y.Label.Text = "Count"

	w := vg.Points(20)

	// groupA := plotter.Values{20, 35, 30, 35, 27}
	groupA := plotter.Values{43, 39, 42, 48, 31, 36, 31, 32, 38, 53, 51, 35, 43, 64, 58, 45, 28, 17, 37, 38, 54, 76, 89, 43, 52, 48, 104, 99, 43, 19, 34, 76, 61, 39, 172, 91, 147, 54, 43, 65, 59, 126, 136, 171, 85, 113, 81, 143, 76, 22, 44, 31}

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		return err
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	// barsA.Offset = -w

	p.Add(barsA)
	// p.Legend.Add("Pull Requests", barsA)

	p.NominalX("2020", "", "", "", "", "", "", "", "", "", "", "", "2021", "", "", "", "", "", "", "", "", "", "", "", "2022", "", "", "", "", "", "", "", "", "", "", "", "2023", "", "", "", "", "", "", "", "", "", "", "", "2024", "", "", "")

	return p.Save(5*vg.Inch, 3*vg.Inch, plotFileName)
}
