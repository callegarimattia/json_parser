package parser_test

import (
	"testing"

	"github.com/callegarimattia/json_parser/internal/parser"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	json          string
	expectedCode  int
	expectedError error
}

func TestParser_Parse(t *testing.T) {
	p := parser.NewParser()

	testcases := []testcase{
		{``, 0, nil},
	}

	for _, tc := range testcases {
		code, err := p.Parse([]byte(tc.json))
		assert.Equal(t, tc.expectedCode, code)
		assert.Equal(t, tc.expectedError, err)
	}
}
