package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permitions := jwt.MapClaims{}
	permitions["authorized"] = true
	permitions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permitions["usuarioID"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permitions)
	return token.SignedString([]byte("Secret"))
}
