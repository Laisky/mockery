package cmd

import (
	"testing"

	"github.com/Laisky/testify/assert"
)

func TestNewShowConfigCmd(t *testing.T) {
	cmd := NewShowConfigCmd()
	assert.Equal(t, "showconfig", cmd.Name())
}
