package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	c := assert.New(t)
	c.Equal(1, 1, "hihi")
}
