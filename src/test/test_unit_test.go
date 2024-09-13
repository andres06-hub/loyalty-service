package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockTest__Unit(t *testing.T) {
	assr := assert.New(t)

	t.Log("mock test")

	assr.True(true)
}
