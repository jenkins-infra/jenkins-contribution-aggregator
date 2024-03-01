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
	"testing"
)

func Test_isFileValid(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Happy case",
			args{"../test_data/not_a_csv.txt"},
			true,
		},
		{
			"File does not exist",
			args{"unexistantFile.txt"},
			false,
		},
		{
			"File is a directory in fact",
			args{"../test_data"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFileValid(tt.args.fileName); got != tt.want {
				t.Errorf("isFileValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateMonth(t *testing.T) {
	type args struct {
		month     string
		isVerbose bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"lowercase latest",
			args{"latest", true},
			true,
		},
		{
			"uppercase latest",
			args{"LATEST", true},
			true,
		},
		{
			"empty month",
			args{"", true},
			false,
		},
		{
			"happy case 1",
			args{"2023-08", true},
			true,
		},
		{
			"happy case 2",
			args{"2013-08", true},
			true,
		},
		{
			"happy case 3",
			args{"2023-12", true},
			true,
		},
		{
			"happy case 4",
			args{"2020-12", true},
			true,
		},
		{
			"invalid month 1",
			args{"2023-13", true},
			false,
		},
		{
			"invalid month 2",
			args{"2023-00", true},
			false,
		},
		{
			"invalid year (too old)",
			args{"2003-08", true},
			false,
		},
		{
			"plain junk 1",
			args{"2023", true},
			false,
		},
		{
			"plain junk 2",
			args{"blaah", true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidMonth(tt.args.month, tt.args.isVerbose); got != tt.want {
				t.Errorf("validateMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWithMDfileExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Markdown extension",
			args{filename: "myfile.md"},
			true,
		},
		{
			"Markdown extension (mixed case)",
			args{filename: "myfile.mD"},
			true,
		},
		{
			"CSV extension",
			args{filename: "myfile.csv"},
			false,
		},
		{
			"no extension",
			args{filename: "myfile"},
			false,
		},
		{
			"just the dot",
			args{filename: "myfile."},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWithMDfileExtension(tt.args.filename); got != tt.want {
				t.Errorf("isWithMDfileExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
