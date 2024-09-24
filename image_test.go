package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestImageHeight(t *testing.T) {
	assert.Equal(t, adaptivecard.BlockElementHeightStretch, adaptivecard.NewImage("test",
		adaptivecard.WithHeight(adaptivecard.BlockElementHeightStretch),
	).Height)

	assert.Equal(t, "50px", adaptivecard.NewImage("test",
		adaptivecard.WithHeight("50px"),
	).Height)
}
