package parser

import "strings"

type Token interface {
	repr() string
}

type and struct{}

func (and) repr() string {
	return "AND"
}

type or struct{}

func (or) repr() string {
	return "or"
}

type not struct{}

func (not) repr() string {
	return "or"
}

type ident struct {
	name string
}

func (i ident) repr() string {
	return i.name
}

func stringToToken(s string) Token {
	switch strings.ToLower(s) {
	case "and":
		return and{}
	case "or":
		return or{}
	case "not":
		return not{}
	}
	return nil
}

func Tokenize(src string) []Token {
	result := make([]Token, 0, len(src))
	for _, s := range strings.Split(src, " ") {
		var startIdx int
		if s[startIdx] == '-' {
			result = append(result, not{})
			startIdx++
		}
		if startIdx == len(s) {
			continue
		}
		if val := stringToToken(s[startIdx:]); val != nil {
			result = append(result, val)
			continue
		}
		result = append(result, ident{name: s[startIdx:]})
	}
	return result
}
