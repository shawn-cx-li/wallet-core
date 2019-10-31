package eth

import (
	"testing"
)

type checksumCase struct {
	input  string
	output string
}

var tests = []checksumCase{
	{
		"0xfb6916095ca1df60bb79ce92ce3ea74c37c5d359",
		"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
	},
	{
		"abc",
		"abc",
	},
	{
		"1MUz4VMYui5qY1mxUiG8BQ1Luv6tqkvaiL",
		"1MUz4VMYui5qY1mxUiG8BQ1Luv6tqkvaiL",
	},
}

func TestToChecksum(t *testing.T) {
	for i := range tests {
		c := tests[i]
		o := ToChecksum(c.input)
		if o != c.output {
			t.Errorf("failed to convert checksum %s. expect %s, got %s", c.input, c.output, o)
		}
	}
}
