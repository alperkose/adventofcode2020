package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMaskedInput(t *testing.T) {
	testCases := []struct {
		desc        string
		input       string
		mask        uint64
		maskedValue uint64
	}{
		{
			desc:        "Simple mask of 1",
			input:       "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1",
			mask:        1,
			maskedValue: 1,
		}, {
			desc:        "Simple mask of 0",
			input:       "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX0",
			mask:        1,
			maskedValue: 0,
		}, {
			desc:        "A mask from excercise",
			input:       "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			mask:        66,
			maskedValue: 64,
		}, {
			desc:        "A mask with 1 as most significant bit",
			input:       "1XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			mask:        1 << 35,
			maskedValue: 1 << 35,
		}, {
			desc:        "A mask with 0 as most significant bit",
			input:       "0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			mask:        1 << 35,
			maskedValue: 0,
		}, {
			desc:        "A mask of all 0",
			input:       "000000000000000000000000000000000000",
			mask:        0xfffffffff,
			maskedValue: 0,
		}, {
			desc:        "A mask of all 1",
			input:       "111111111111111111111111111111111111",
			mask:        0xfffffffff,
			maskedValue: 0xfffffffff,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m := FromMaskString(tC.input)
			assert.Equal(t, tC.mask, m.Mask)
			assert.Equal(t, tC.maskedValue, m.Value)
		})
	}
}

func TestMappingMaskToInputs(t *testing.T) {
	testCases := []struct {
		desc   string
		mask   Mask
		input  uint64
		output uint64
	}{
		{
			desc:   "All forced 1",
			mask:   FromMaskString("111111111111111111111111111111111111"),
			input:  42,
			output: 0xfffffffff,
		}, {
			desc:   "All forced 0",
			mask:   FromMaskString("000000000000000000000000000000000000"),
			input:  42,
			output: 0,
		}, {
			desc:   "None all ignored",
			mask:   FromMaskString("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			input:  42,
			output: 42,
		}, {
			desc:   "Partially forced 1 all allowed",
			mask:   FromMaskString("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XX11XX"),
			input:  42,
			output: 110,
		}, {
			desc:   "Partially forced 0 all allowed",
			mask:   FromMaskString("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX00XX00"),
			input:  42,
			output: 8,
		}, {
			desc:   "Partially forced 0 and 1 all allowed",
			mask:   FromMaskString("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"),
			input:  11,
			output: 73,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.mask.MapTo(tC.input)
			assert.Equal(t, tC.output, actual)
		})
	}
}

func TestMappingMaskToAlternatives(t *testing.T) {
	testCases := []struct {
		desc   string
		mask   Mask
		input  uint64
		output []uint64
	}{
		{
			desc:   "All forced 1",
			mask:   FromMaskString("111111111111111111111111111111111111"),
			input:  42,
			output: []uint64{0xfffffffff},
		}, {
			desc:   "All forced 0",
			mask:   FromMaskString("000000000000000000000000000000000000"),
			input:  42,
			output: []uint64{42},
		}, {
			desc:   "Two floating bit",
			mask:   FromMaskString("000000000000000000000000000000X1001X"),
			input:  42,
			output: []uint64{26, 27, 58, 59},
		}, {
			desc:   "Three floating bit",
			mask:   FromMaskString("00000000000000000000000000000000X0XX"),
			input:  26,
			output: []uint64{16, 17, 18, 19, 24, 25, 26, 27},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.mask.MapToAlternatives(tC.input)
			assert.ElementsMatch(t, tC.output, actual)
		})
	}
}
