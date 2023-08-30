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
