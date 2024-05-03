package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	assert.Equal(t, "世界でいちばんおひめさま", Hello())
}
