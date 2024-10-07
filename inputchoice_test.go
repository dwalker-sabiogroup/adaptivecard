package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestInputChoiceDefault(t *testing.T) {
	ic := adaptivecard.NewInputChoice("title", "value")

	assert.Equal(t, "title", ic.Title)
	assert.Equal(t, "value", ic.Value)
}
