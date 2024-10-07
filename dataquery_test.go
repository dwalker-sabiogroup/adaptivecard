package adaptivecard_test

import (
	"testing"

	"github.com/dwalker-sabiogroup/adaptivecard"
	"github.com/stretchr/testify/assert"
)

func TestDataQueryDefault(t *testing.T) {
	dq := adaptivecard.NewDataQuery("dataset")

	assert.Equal(t, "dataset", dq.Dataset)
	assert.Empty(t, dq.Count)
	assert.Empty(t, dq.Skip)
}

func TestDataQueryCount(t *testing.T) {
	assert.Equal(t, 2, adaptivecard.NewDataQuery("dataset",
		adaptivecard.WithDataQueryCount(2),
	).Count)
}

func TestDataQuerySkip(t *testing.T) {
	assert.Equal(t, 2, adaptivecard.NewDataQuery("dataset",
		adaptivecard.WithDataQuerySkip(2),
	).Skip)
}
