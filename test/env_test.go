package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ylanzinhoy/tools_scraping_py/internal/controller"
	"github.com/ylanzinhoy/tools_scraping_py/internal/services"
)

func EnvTest(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(nil, controller.CreateEnv)
}


