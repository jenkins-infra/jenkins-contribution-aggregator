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
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plot_bargraph(t *testing.T) {
	//setup environment
	tempDir := t.TempDir()

	labels := []string{"2020-01", "2020-02", "2020-03", "2020-04", "2020-05", "2020-06", "2020-07", "2020-08", "2020-09", "2020-10", "2020-11", "2020-12",
		"2021-01", "2021-02", "2021-03", "2021-04", "2021-05", "2021-06", "2021-07", "2021-08", "2021-09", "2021-10", "2021-11", "2021-12",
		"2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12",
		"2023-01", "2023-02", "2023-03", "2023-04", "2023-05", "2023-06", "2023-07", "2023-08", "2023-09", "2023-10", "2023-11", "2023-12",
		"2024-01", "2024-02", "2024-03", "2024-04"}

	values := []string{"43", "39", "42", "48", "31", "36", "31", "32", "38", "53", "51", "35",
		"43", "64", "58", "45", "28", "17", "37", "38", "54", "76", "89", "43",
		"52", "48", "104", "99", "43", "19", "34", "76", "61", "39", "172", "91",
		"147", "54", "43", "65", "59", "126", "136", "171", "85", "113", "81", "143",
		"76", "22", "44", "31"}

	err := plot_bargraph(tempDir, "test", labels, values)
	assert.NoError(t, err, "Function should not have failed")
	outputPngFileName := filepath.Join(tempDir, "test.png")
	assert.FileExists(t,outputPngFileName,"No graphic file generated")
	fmt.Printf("To view the file generated: open %s\n",outputPngFileName)
}

func Test_generateAxisLabels(t *testing.T) {
	type args struct {
		inputLabels []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Simple case",
			args{inputLabels: []string{"2020-01", "2020-02", "2020-03", "2020-04", "2020-05", "2020-06", "2020-07", "2020-08", "2020-09", "2020-10", "2020-11", "2020-12", "2021-01", "2021-02"}},
			[]string{"2020", "", "", "", "", "", "", "", "", "", "", "", "2021", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyAxisLabels(tt.args.inputLabels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateAxisLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertValuesToFloats(t *testing.T) {
	type args struct {
		stringValues []string
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			"Happy case",
			args{stringValues: []string{"1", "2", "3"}},
			[]float64{1, 2, 3},
			false,
		},
		//TODO: empty input
		{
			"empty element",
			args{stringValues: []string{"1", "2", ""}},
			nil,
			true,
		},
		{
			"space element",
			args{stringValues: []string{"1", "2", " "}},
			nil,
			true,
		},
		{
			"invalid value",
			args{stringValues: []string{"1", "2", "junk"}},
			nil,
			true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertValuesToInts(tt.args.stringValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertValuesToInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertValuesToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
