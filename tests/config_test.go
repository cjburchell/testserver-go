package tests

import (
	"testing"

	"github.com/cjburchell/docker-compose"
	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {

	redisFile := docker_compose.File{
		Version: "2",
		Services: map[string]docker_compose.Service{
			"redis": docker_compose.Service{
				Image: "redis:latest",
				Ports: []string{"6379:6379"},
			},
		},
	}

	assert.Equal(t, 123, 123, "they should be equal")
}
