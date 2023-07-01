package config_test

import (
	"testing"
	"todo-go/config"

	"github.com/magiconair/properties/assert"
)

func TestConfig(t *testing.T) {
	got, err := config.New()
	assert.Equal(t, &config.Config{}, got)
	assert.Equal(t, nil, err)
}