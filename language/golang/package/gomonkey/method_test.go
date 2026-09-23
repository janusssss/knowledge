package gomonkey

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

type calculator struct {
}

func (c *calculator) AddMethod(a, b int) int {
	return a + b
}
func TestAddMethod(t *testing.T) {
	c := &calculator{}
	patch := gomonkey.ApplyMethod(c, "AddMethod",
		func(_ *calculator, a, b int) int {
			return 42
		})
	defer patch.Reset()
	result := c.AddMethod(1, 2)
	assert.Equal(t, 42, result, "Expected Add to return 42")
}
