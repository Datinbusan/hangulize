package hangulize

import (
	"testing"

	"github.com/hangulize/hre"
	"github.com/stretchr/testify/assert"
)

func TestRuleString(t *testing.T) {
	p, _ := hre.NewPattern("foo", nil, nil)
	rp := hre.NewRPattern("bar", nil, nil)
	r := Rule{p, rp}
	assert.Equal(t, `"foo" -> "bar"`, r.String())
}

func TestRuleReplacements(t *testing.T) {
	p, _ := hre.NewPattern("foo", nil, nil)
	rp := hre.NewRPattern("bar", nil, nil)
	r := Rule{p, rp}

	repls := r.replacements("abcfoodef")

	assert.Len(t, repls, 1)
	assert.Equal(t, 3, repls[0].start)
	assert.Equal(t, 6, repls[0].stop)
	assert.Equal(t, "bar", repls[0].word)
}

func TestRuleReplace(t *testing.T) {
	p, _ := hre.NewPattern("foo", nil, nil)
	rp := hre.NewRPattern("bar", nil, nil)
	r := Rule{p, rp}
	assert.Equal(t, "abcbardef", r.Replace("abcfoodef"))
}
