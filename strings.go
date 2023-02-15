package gutils

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

var (
	reNonASCII     = regexp.MustCompile("[[:^ascii:]]")
	reNonPrintable = regexp.MustCompile("[[:^print:]]")
)

func ImplodeLines(lines []string) string {
	return strings.Join(lines, "\n")
}

func ExplodeLines(lines string) []string {
	return strings.Split(lines, "\n")
}

func ReplaceNonASCII(str, repl string) string {
	return reNonASCII.ReplaceAllLiteralString(str, repl)
}

func ReplaceNonPrintable(str, repl string) string {
	return reNonPrintable.ReplaceAllLiteralString(str, repl)
}

func RemoveNonPrintable(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
}

func ChunkString(s string, sep string, chunkSize int) [][]string {
	res := [][]string{}
	pieces := strings.Split(s, sep)
	var chunk []string
	for len(pieces) > 0 {
		if len(pieces) > chunkSize {
			chunk = pieces[0:chunkSize]
			pieces = pieces[chunkSize:]
		} else {
			chunk = pieces[0:]
			pieces = []string{}
		}
		res = append(res, chunk)
	}
	return res
}

func RemoveCommandFlags(parts []string) []string {
	res := []string{}
	for _, p := range parts {
		if !strings.HasPrefix(p, "-") {
			res = append(res, p)
		}
	}
	return res
}

// GeneratePseudoEmptyString creates a string that looks empty when printed in a terminal.
//
// The string consists of `n` times a space followed by a backspace.
//
// If `n` is zero, the function will use a random value between 1 and 1000 (inclusive).
func GeneratePseudoEmptyString(n int) string {
	if n == 0 {
		n = GetRandomInt(1, 1000)
	}
	return strings.Repeat(" \u0008", n) // space follow by backspace
}

// GenerateGarbageString produces a string (length is randomly chosen between 1 and `n`)
// consisting of random (non)-printable characters.
func GenerateGarbageString(n int) string {
	garbage := make([]byte, GetRandomInt(1, n))
	rand.Read(garbage)
	return string(garbage)
}
