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
	"reflect"
	"testing"
)

var searchSubmitter_dataset = [][]string{
	{"submitter", "PR"},
	{"alpha", "0"},
	{"bravo", "0"},
	{"charly", "0"},
	{"delta", "0"},
}

func Test_isSubmitterFound(t *testing.T) {
	type args struct {
		dataset   [][]string
		submitter string
	}
	tests := []struct {
		name      string
		args      args
		wantFound bool
	}{
		{
			"happy case",
			args{dataset: searchSubmitter_dataset, submitter: "bravo"},
			true,
		},
		{
			"Not found",
			args{dataset: searchSubmitter_dataset, submitter: "foxtrot"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFound := isSubmitterFound(tt.args.dataset, tt.args.submitter); gotFound != tt.wantFound {
				t.Errorf("isSubmitterFound() = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

var dataset_1 = [][]string{
	{"submitter", "PR"},
	{"alpha", "1"},
	{"bravo", "2"},
	{"charly", "3"},
	{"delta", "4"},
}
var dataset_2 = [][]string{
	{"submitter", "PR"},
	{"alpha", "1"},
	{"charly", "2"},
	{"delta", "3"},
	{"zebra", "4"},
}

var dataset_result = [][]string{
	{"Submitter", "Total_PRs", "Status"},
	{"alpha", "1", ""},
	{"bravo", "2", "new"},
	{"charly", "3", ""},
	{"delta", "4", ""},
	{"zebra", "", "churned"},
}

func Test_compareExtractedData(t *testing.T) {
	type args struct {
		recentData [][]string
		oldData    [][]string
	}
	tests := []struct {
		name                      string
		args                      args
		wantEnrichedExtractedData [][]string
	}{
		{
			"case 1",
			args{recentData: dataset_1, oldData: dataset_2},
			dataset_result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEnrichedExtractedData := compareExtractedData(tt.args.recentData, tt.args.oldData); !reflect.DeepEqual(gotEnrichedExtractedData, tt.wantEnrichedExtractedData) {
				t.Errorf("compareExtractedData() = %v, want %v", gotEnrichedExtractedData, tt.wantEnrichedExtractedData)
			}
		})
	}
}
