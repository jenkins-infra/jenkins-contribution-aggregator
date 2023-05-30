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

import "testing"

var records_1 = [][]string{
	{"", "2022-01", "2022-02", "2022-03", "2022-04", "2022-05", "2022-06", "2022-07", "2022-08", "2022-09", "2022-10", "2022-11", "2022-12", "2023-01", "2023-02", "2023-03", "2023-04"},
	{"0x41head", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"},
	{"AScripnic", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"},
}

var records_2 = [][]string{
	{"", "2022-01"},
	{"0x41head", "1"},
	{"AScripnic", "1"},
}

func Test_getBoundaries(t *testing.T) {

	type args struct {
		records [][]string
		months  int
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
			args{records: records_1, months: 12},
			5, 16, "2022-05", "2023-04",
		},
		{
			"Get all available months",
			args{records: records_1, months: 0},
			1, 16, "2022-01", "2023-04",
		},
		{
			"Get more months than available",
			args{records: records_1, months: 20},
			1, 16, "2022-01", "2023-04",
		},		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartColumn, gotEndColumn, gotStartMonth, gotEndMonth := getBoundaries(tt.args.records, tt.args.months)
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
