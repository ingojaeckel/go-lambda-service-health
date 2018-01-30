package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationNonNil(t *testing.T) {
	c, err := LoadConfiguration("../config.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestConfigurationTimeout(t *testing.T) {
	c, _ := LoadConfiguration("../config.yaml")
	assert.Equal(t, 1, c.Timeout)
}

func TestConfigurationServices(t *testing.T) {
	c, _ := LoadConfiguration("../config.yaml")
	assert.Equal(t, 2, len(c.Services))

	assert.True(t, len(c.Services[0].Name) > 0)
	assert.True(t, len(c.Services[0].URL) > 0)

	assert.True(t, len(c.Services[1].Name) > 0)
	assert.True(t, len(c.Services[1].URL) > 0)
}
