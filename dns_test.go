package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	hw := "hello world"
	assert.Equal(t, "hello world", hw)
}
