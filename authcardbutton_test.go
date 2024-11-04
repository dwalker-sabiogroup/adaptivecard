package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestAuthCardButtonDefault(t *testing.T) {
	acb := adaptivecard.NewAuthCardButton("type", "value")

	assert.Equal(t, "type", acb.Type)
	assert.Equal(t, "value", acb.Value)
	assert.Empty(t, acb.Title)
	assert.Empty(t, acb.Image)
}

func TestAuthCardButtonTitle(t *testing.T) {
	assert.Equal(t, "title", adaptivecard.NewAuthCardButton("type", "value",
		adaptivecard.WithAuthCardButtonTitle("title"),
	).Title)
}

func TestAuthCardButtonImage(t *testing.T) {
	assert.Equal(t, "image", adaptivecard.NewAuthCardButton("type", "value",
		adaptivecard.WithAuthCardButtonImage("image"),
	).Image)
}
