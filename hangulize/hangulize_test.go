package hangulize

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func hangulize(spec *Spec, word string) string {
	h := NewHangulizer(spec)
	return h.Hangulize(word)
}

// TestLang generates subtests for bundled lang specs.
func TestLang(t *testing.T) {
	for _, lang := range ListLangs() {
		t.Run(lang, func(t *testing.T) {
			spec, ok := LoadSpec(lang)

			assert.Truef(t, ok, `failed to load "%s" spec`, lang)

			h := NewHangulizer(spec)

			for _, testCase := range spec.Test {
				loanword := testCase.Left()
				expected := testCase.Right()[0]

				ch := make(chan Trace, 1000)
				got := h.HangulizeTrace(loanword, ch)

				if got == expected {
					continue
				}

				// Trace result to understand the failure reason.
				f := bytes.NewBufferString("")
				hr := strings.Repeat("-", 30)

				// Render failure message.
				fmt.Fprintln(f, hr)
				fmt.Fprintf(f, `lang: "%s"`, lang)
				fmt.Fprintln(f)
				fmt.Fprintf(f, `word: %#v`, loanword)
				fmt.Fprintln(f)
				fmt.Fprintln(f, hr)
				for e := range ch {
					fmt.Fprintln(f, e.String())
				}
				fmt.Fprintln(f, hr)

				assert.Equal(t, expected, got, f.String())
			}
		})
	}
}

// TestSlash tests "/" in input word.  The original Hangulize removes the "/"
// so the result was "글로르이아" instead of "글로르/이아".
func TestSlash(t *testing.T) {
	assert.Equal(t, "글로르/이아", Hangulize("ita", "glor/ia"))
}

func TestHyphen(t *testing.T) {
	spec := parseSpec(`
	config:
		markers = "-"

	transcribe:
		"e" -> "ㅔ"
		"x" -> "-ㄱㅅ"
	`)
	assert.Equal(t, "엑스", hangulize(spec, "ex"))
}
