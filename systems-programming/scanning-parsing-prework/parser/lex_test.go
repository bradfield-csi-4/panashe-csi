package parser

import (
	"fmt"
	"testing"
)

func testExpectEqual(t *testing.T, expected, got interface{}) {
	if expected != got {
		t.Errorf("expected: %#v, got %#v", expected, got)
	}
}

// ["shumac", or{}, or{}, and{}, "sovereignness"]
// ["subtransparent", or{}, "trame", "uninflicted", "unclassifiable"]
// [and{}, not{}, "vulcanizable", and{}, "pillary"]
// [or{}, or{}, or{}, not{}, or{}]
// [and{}, and{}, and{}, and{}, "term"]

func TestProducesTokens(t *testing.T) {
	cases := []struct {
		src      string
		expected []Token
	}{
		{src: "hello world OR alice -bob", expected: []Token{
			ident{name: "hello"}, ident{name: "world"}, or{}, ident{name: "alice"}, not{}, ident{name: "bob"},
		},
		},
	}
	for i := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			// testExpectEqual(t, cases[
		})
	}
	// tokens := Tokenize(src)
	// for i := range tokens {
	// 	testExpectEqual(t, expectedTokens[i], tokens[i])
	// }
}
