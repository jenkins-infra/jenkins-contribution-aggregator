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

func Test_checkFile(t *testing.T) {
	type args struct {
		fileName string
		isSilent bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Bad first column",
			args{
				fileName: "../test_data/bad_first_column.csv",
				isSilent: false,
			},
			false,
		},
		{
			"Bad date column",
			args{
				fileName: "../test_data/bad_date_column.csv",
				isSilent: false,
			},
			false,
		},
		{
			"Bad submitter name",
			args{
				fileName: "../test_data/bad_submitter_name.csv",
				isSilent: false,
			},
			false,
		},
		{
			"deleted user case",
			args{
				fileName: "../test_data/deleted_user_case.csv",
				isSilent: false,
			},
			true,
		},
		{
			"non integer data value",
			args{
				fileName: "../test_data/bad_data_value.csv",
				isSilent: false,
			},
			false,
		},
		{
			"negative data value",
			args{
				fileName: "../test_data/bad_data_negative_value.csv",
				isSilent: false,
			},
			false,
		},
		{
			"file not found",
			args{
				fileName: "../test_data/blaah.csv",
				isSilent: false,
			},
			false,
		},
		{
			"Happy case",
			args{
				fileName: "../test_data/overview.csv",
				isSilent: false,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkFile(tt.args.fileName, tt.args.isSilent); got != tt.want {
				t.Errorf("checkFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
