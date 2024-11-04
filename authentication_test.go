package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationDefault(t *testing.T) {
	a := adaptivecard.NewAuthentication()

	assert.Empty(t, a.Buttons)
}

func TestAuthenticationText(t *testing.T) {
	assert.Equal(t, "text", adaptivecard.NewAuthentication(
		adaptivecard.WithAuthenticationText("text"),
	).Text)
}

func TestAuthenticationConnectionName(t *testing.T) {
	assert.Equal(t, "connection_name", adaptivecard.NewAuthentication(
		adaptivecard.WithAuthenticationConnectionName("connection_name"),
	).ConnectionName)
}

func TestAuthenticationTokenExchnageResource(t *testing.T) {
	ter := adaptivecard.NewTokenExchangeResource("id", "uri", "provider_id")
	assert.Equal(t, *ter, adaptivecard.NewAuthentication(
		adaptivecard.WithAuthenticationTokenExchangeResource(*ter),
	).TokenExchangeResource)
}

func TestAuthenticationButtons(t *testing.T) {
	acb := adaptivecard.NewAuthCardButton("type", "value")
	assert.Equal(t, *acb, adaptivecard.NewAuthentication(
		adaptivecard.WithAuthenticationButtons([]adaptivecard.AuthCardButton{*acb}),
	).Buttons[0])
}
