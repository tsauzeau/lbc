package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsauzeau/lbc/cmd/lbc/config"
)

// Example test to show usage of `make test`
func TestDummy(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestLoadConfig(t *testing.T) {
	if err := config.LoadConfig("/config"); err != nil {
		assert.Error(t, err)
	}
}
