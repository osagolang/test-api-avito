package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {

	token, err := GenerateToken(777)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

}

func TestParseToken_valid(t *testing.T) {

	userID := 777
	token, _ := GenerateToken(userID)

	parsedToken, err := ParseToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, parsedToken)

}

func TestParseToken_invalid(t *testing.T) {

	token := "invalid.token"

	parsedToken, err := ParseToken(token)

	assert.Error(t, err)
	assert.Nil(t, parsedToken)

}
