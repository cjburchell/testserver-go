package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTest(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
}
