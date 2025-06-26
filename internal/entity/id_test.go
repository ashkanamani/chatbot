package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDTypeValue(t *testing.T) {
	assert.Equal(t, "type", NewID("type", "val").Type())
	assert.Equal(t, "val", NewID("type", "val").ID())
}
