package utils_jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	sut "utils/jwt"
)

func TestUtilsJwt_GenerateJWT(t *testing.T) {
	var input = "unitestusername"
	
	calculated, err := sut.GenerateJWT(input)
	assert.Nil(t, err)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyIjoidW5pdGVzdHVzZXJuYW1lIn0.0eS90T6SDLTFZ3ubS25dnfqyF-Ff4k-uSHLI_Sdrq-s", calculated)
}

func TestUtilsJwt_GenerateJWT_empty_username(t *testing.T) {
	var input = ""
	
	calculated, err := sut.GenerateJWT(input)
	assert.NotNil(t, err)
	assert.Equal(t, "", calculated)
}

func TestUtilsJwt_GetUsernameFromToken(t *testing.T) {
	var input = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyIjoidW5pdGVzdHVzZXJuYW1lIn0.0eS90T6SDLTFZ3ubS25dnfqyF-Ff4k-uSHLI_Sdrq-s"

	calculated, err := sut.GetUsernameFromToken(input)
	assert.Nil(t, err)
	assert.Equal(t, "unitestusername", calculated)
}

func TestUtilsJwt_GetUsernameFromToken_empty_token(t *testing.T) {
	var input = ""

	calculated, err := sut.GetUsernameFromToken(input)
	assert.NotNil(t, err)
	assert.Equal(t, "", calculated)
}