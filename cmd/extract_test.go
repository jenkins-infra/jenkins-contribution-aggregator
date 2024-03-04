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
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var records_1 = [][]string{
	{"", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"},
	{"0x41head", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"},
	{"AScripnic", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"},
}

var records_2 = [][]string{
	{"", "2022-01", "2022-02"},
	{"0x41head", "1", "2"},
	{"AScripnic", "1", "2"},
}

var resultSlice_1 = [][]string{
	{"Submitter", "Total_PRs"},
	{"0x41head", "78"},
	{"AayushSaini101", "15"},
	{"ChadiEM", "8"},
	{"Artmorse", "6"},
	{"Abingcbc", "5"},
	{"CatherineKiiru", "4"},
	{"AndKiel", "4"},
	{"Absh-Day", "4"},
	{"BorisYaoA", "4"},
}

func Test_getBoundaries(t *testing.T) {

	type args struct {
		records     [][]string
		endMonthStr string
		months      int
		offset      int
	}
	tests := []struct {
		name            string
		args            args
		wantStartColumn int
		wantEndColumn   int
		wantStartMonth  string
		wantEndMonth    string
	}{
		{
			"Normal case",
			args{records: records_1, endMonthStr: "latest", months: 12, offset: 0},
			5, 16, "2022-05", "2023-04",
		},
		{
			"Get all available months",
			args{records: records_1, endMonthStr: "latest", months: 0, offset: 0},
			1, 16, "2022-01", "2023-04",
		},
		{
			"Get more months than available",
			args{records: records_1, endMonthStr: "latest", months: 20, offset: 0},
			1, 16, "2022-01", "2023-04",
		},
		{
			"Specify end month - normal case",
			args{records: records_1, endMonthStr: "2023-02", months: 6, offset: 0},
			9, 14, "2022-09", "2023-02",
		},
		{
			"Specify end month - get all available months",
			args{records: records_1, endMonthStr: "2023-02", months: 0, offset: 0},
			1, 14, "2022-01", "2023-02",
		},
		{
			"Specify end month - get more months than available",
			args{records: records_1, endMonthStr: "2023-02", months: 20, offset: 0},
			1, 14, "2022-01", "2023-02",
		},
		{
			"Specify end month - end month not found",
			args{records: records_1, endMonthStr: "2023-08", months: 12, offset: 0},
			5, 16, "2022-05", "2023-04",
		},
		{
			"short month set",
			args{records: records_2, endMonthStr: "latest", months: 12, offset: 0},
			1, 2, "2022-01", "2022-02",
		},
		{
			"Normal case with offset",
			args{records: records_1, endMonthStr: "latest", months: 12, offset: 1},
			4, 15, "2022-04", "2023-03",
		},
		{
			"Normal case with endMonth and offset",
			args{records: records_1, endMonthStr: "2023-02", months: 6, offset: 1},
			8, 13, "2022-08", "2023-01",
		},
		{
			"endMonth and offset out od bound",
			args{records: records_1, endMonthStr: "2023-02", months: 6, offset: 16},
			0, 0, "", "",
		},
		{
			"offset with latest out of dataset",
			args{records: records_1, endMonthStr: "latest", months: 12, offset: 16},
			0, 0, "", "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartColumn, gotEndColumn, gotStartMonth, gotEndMonth := getBoundaries(tt.args.records, tt.args.endMonthStr, tt.args.months, tt.args.offset)
			if gotStartColumn != tt.wantStartColumn {
				t.Errorf("getBoundaries() gotStartColumn = %v, want %v", gotStartColumn, tt.wantStartColumn)
			}
			if gotEndColumn != tt.wantEndColumn {
				t.Errorf("getBoundaries() gotEndColumn = %v, want %v", gotEndColumn, tt.wantEndColumn)
			}
			if gotStartMonth != tt.wantStartMonth {
				t.Errorf("getBoundaries() gotStartMonth = %v, want %v", gotStartMonth, tt.wantStartMonth)
			}
			if gotEndMonth != tt.wantEndMonth {
				t.Errorf("getBoundaries() gotEndMonth = %v, want %v", gotEndMonth, tt.wantEndMonth)
			}
		})
	}
}

func Test_extractData(t *testing.T) {
	type args struct {
		inputFilename    string
		topSize          int
		endMonth         string
		period           int
		offset           int
		isVerboseExtract bool
	}
	tests := []struct {
		name             string
		args             args
		wantResult       bool
		wantReal_endDate string
		wantOutputSlice  [][]string
	}{
		{
			"Happy case",
			args{
				inputFilename:    "../test_data/short_overview.csv",
				topSize:          7,
				endMonth:         "latest",
				period:           12,
				isVerboseExtract: false,
			},
			true, "2023-04", resultSlice_1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotReal_endDate, gotOutputSlice := extractData(tt.args.inputFilename, tt.args.topSize, tt.args.endMonth, tt.args.period, tt.args.offset, tt.args.isVerboseExtract)
			if gotResult != tt.wantResult {
				t.Errorf("extractData() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotReal_endDate != tt.wantReal_endDate {
				t.Errorf("extractData() gotReal_endDate = %v, want %v", gotReal_endDate, tt.wantReal_endDate)
			}
			if !reflect.DeepEqual(gotOutputSlice, tt.wantOutputSlice) {
				t.Errorf("extractData() gotOutputSlice = %v, want %v", gotOutputSlice, tt.wantOutputSlice)
			}
		})
	}
}

func Test_ExecuteExtractToMarkdown_integrationTest(t *testing.T) {
	// Setup test environment
	tempDir := t.TempDir()
	testOutputFilename := tempDir + "extract_markdown_output.md"
	goldenMarkdownFilename, err := duplicateFile("../test_data/extract_reference_output.md", tempDir)

	assert.NoError(t, err, "Unexpected Golden File duplication error")
	assert.NotEmpty(t, goldenMarkdownFilename, "Failure to duplicate Golden File")

	// setup the command line
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"extract", "../test_data/overview.csv", "--month=latest", "--period=12", "--topSize=35", "--out=" + testOutputFilename})

	// Execute the module under test
	error := rootCmd.Execute()

	// Check the results
	assert.NoError(t, error, "Unexpected failure")
	assert.True(t, isFileEquivalent(testOutputFilename, goldenMarkdownFilename))
}

//TODO: integration test for CSV output