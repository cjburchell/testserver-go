package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
}
