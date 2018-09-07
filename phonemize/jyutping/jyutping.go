/*
Package jyutping implements the hangulize.Phonemizer interface for Cantonese
Hanzu. It's a variant of the package pinyin.
*/
package jyutping

import (
	"bytes"
	"strings"

	goJyutping "github.com/sublee/go-jyutping"
)

// P is the Jyutping phonemizer.
var P jyutpingPhonemizer

// ----------------------------------------------------------------------------

type jyutpingPhonemizer struct{}

func (jyutpingPhonemizer) ID() string {
	return "jyutping"
}

func (p *jyutpingPhonemizer) Phonemize(word string) string {
	var chunks []string
	var buf bytes.Buffer

	a := goJyutping.NewArgs()

	for _, ch := range word {
		pyn := goJyutping.SinglePinyin(ch, a)

		if len(pyn) == 0 {
			buf.WriteRune(ch)
		} else {
			if buf.Len() != 0 {
				chunks = append(chunks, buf.String())
				buf.Reset()
			}
			chunks = append(chunks, pyn[0])
		}
	}
	if buf.Len() != 0 {
		chunks = append(chunks, buf.String())
	}

	return strings.Join(chunks, "\u200b")
}
