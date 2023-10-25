package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {

	type TestCase struct {
		name string
		in   string
		out  string
	}

	testCases := []TestCase{
		{name: "1", in: "a4bc2d5e", out: "aaaabccddddde"},
		{name: "2", in: "abcd", out: "abcd"},
		{name: "3", in: "45", out: ""},
		{name: "4", in: "", out: ""},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, Unpacking(v.in), v.out)
		})
	}
}

func TestEscape(t *testing.T) {

	type TestCase struct {
		name string
		in   string
		out  string
	}

	testCases := []TestCase{
		{name: "1", in: `qwe\4\5`, out: `qwe45`},
		{name: "2", in: `qwe\45`, out: `qwe44444`},
		{name: "3", in: `qwe\\5`, out: `qwe\\\\\`},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, v.out, Unpacking(v.in))
		})
	}
}
