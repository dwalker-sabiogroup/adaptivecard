package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestFactDefault(t *testing.T) {
	f := adaptivecard.NewFact("title", "value")

	assert.Equal(t, "title", f.Title)
	assert.Equal(t, "value", f.Value)
}
