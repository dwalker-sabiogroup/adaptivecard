package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestRefreshDefault(t *testing.T) {
	r := adaptivecard.NewRefresh()

	assert.Len(t, r.UserIDs, 0)
}

func TestRefreshAction(t *testing.T) {
	ae := adaptivecard.NewActionExecute()
	assert.Equal(t, *ae, adaptivecard.NewRefresh(
		adaptivecard.WithRefreshAction(*ae),
	).Action)
}

func TestRefreshExpires(t *testing.T) {
	assert.Equal(t, "2022-01-01T12:00:00Z", adaptivecard.NewRefresh(
		adaptivecard.WithRefreshExpires("2022-01-01T12:00:00Z"),
	).Expires)
}

func TestRefreshUserIDs(t *testing.T) {
	assert.Equal(t, "test-id", adaptivecard.NewRefresh(
		adaptivecard.WithRefreshUserIDs([]string{"test-id"}),
	).UserIDs[0])
}
