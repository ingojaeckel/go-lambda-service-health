package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {
	c, err := LoadConfiguration("../config.yaml")
	assert.Nil(t, err)
	assert.Equal(t, 5, c.Timeout)
}
