package gomonkey

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func AddFunction(a, b int) int {
	return a + b
}

func TestAddFunction(t *testing.T) {
	patch := gomonkey.ApplyFunc(AddFunction, func(a, b int) int {
		return 42
	})
	defer patch.Reset()
	result := AddFunction(1, 2)
	assert.Equal(t, 42, result, "Expected Add to return 42")
}
