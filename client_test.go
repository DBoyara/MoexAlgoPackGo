package moexalgopackgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlgoClient(t *testing.T) {
	login := "testuser"
	password := "testpassword"

	client := NewAlgoClient(login, password)

	assert.NotNil(t, client, "NewAlgoClient returned nil")

}