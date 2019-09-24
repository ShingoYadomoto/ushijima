package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoiSlice(t *testing.T) {
	type testcase struct {
		got   []string
		exp   []int
		isErr bool
	}

	tcl := []testcase{
		{got: []string{"12", "34", "6598", "300"}, exp: []int{12, 34, 6598, 300}, isErr: false},
		{got: []string{"00", "000", "0000"}, exp: []int{0, 0, 0}, isErr: false},
		{got: []string{"20"}, exp: []int{20}, isErr: false},

		{got: []string{"wwww"}, exp: nil, isErr: true},
		{got: []string{"12", "34", "fail", "300"}, exp: nil, isErr: true},
	}
	for _, tc := range tcl {
		got, err := AtoiSlice(tc.got)

		if tc.isErr {
			assert.Error(t, err)
		} else {
			assert.Equal(t, tc.exp, got)
		}
	}

}
