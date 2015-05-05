package abbr

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var Keywords = map[string]map[string]bool{
	"go": {
		"break":       true,
		"case":        true,
		"chan":        true,
		"const":       true,
		"continue":    true,
		"default":     true,
		"defer":       true,
		"else":        true,
		"error":       true,
		"fallthrough": true,
		"for":         true,
		"func":        true,
		"go":          true,
		"goto":        true,
		"if":          true,
		"import":      true,
		"interface":   true,
		"map":         true,
		"package":     true,
		"range":       true,
		"return":      true,
		"select":      true,
		"string":      true,
		"struct":      true,
		"switch":      true,
		"type":        true,
		"var":         true,
	},
	"php": {
		"__halt_compiler": true,
		"abstract":        true,
		"and":             true,
		"array":           true,
		"as":              true,
		"break":           true,
		"callable":        true,
		"case":            true,
		"catch":           true,
		"class":           true,
		"clone":           true,
		"const":           true,
		"continue":        true,
		"declare":         true,
		"default":         true,
		"die":             true,
		"do":              true,
		"echo":            true,
		"else":            true,
		"elseif":          true,
		"empty":           true,
		"enddeclare":      true,
		"endfor":          true,
		"endforeach":      true,
		"endif":           true,
		"endswitch":       true,
		"endwhile":        true,
		"eval":            true,
		"exit":            true,
		"extends":         true,
		"final":           true,
		"finally":         true,
		"for":             true,
		"foreach":         true,
		"function":        true,
		"global":          true,
		"goto":            true,
		"if":              true,
		"implements":      true,
		"include":         true,
		"include_once":    true,
		"instanceof":      true,
		"insteadof":       true,
		"interface":       true,
		"isset":           true,
		"list":            true,
		"namespace":       true,
		"new":             true,
		"or":              true,
		"print":           true,
		"private":         true,
		"protected":       true,
		"public":          true,
		"require":         true,
		"require_once":    true,
		"return":          true,
		"static":          true,
		"switch":          true,
		"throw":           true,
		"trait":           true,
		"try":             true,
		"unset":           true,
		"use":             true,
		"var":             true,
		"while":           true,
	},
}

// IsKeyword checks if a word is a keyword of the given language.
func IsKeyword(word string, lang string) bool {
	lowWord := strings.ToLower(word)
	lowLang := strings.ToLower(lang)
	if keywords[lowLang][lowWord] {
		return true
	}
	return false
}
