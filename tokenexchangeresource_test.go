package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestTokenExchangeResourceDefault(t *testing.T) {
	ter := adaptivecard.NewTokenExchangeResource("id", "uri", "providerID")
	assert.Equal(t, "id", ter.ID)
	assert.Equal(t, "uri", ter.URI)
	assert.Equal(t, "providerID", ter.ProviderID)
}
