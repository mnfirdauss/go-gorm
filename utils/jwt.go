package utils

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["name"] = name
	claims["exp"] = json.Number(strconv.FormatInt(time.Now().Add(time.Minute*15).Unix(), 10))
	claims["aud"] = "Audience" // OPTIONAL AUDIENCE

	JWT_SECRET := `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIGaozMA951amsyyAjz/C3FUhdspS1Kqi3s5EdbJeop0boAoGCCqGSM49
AwEHoUQDQgAEPvB35tXsy4P4ZKpH3jAGGWA4ZVOnQsiLPBrWfjk76UXnrXqZO5LW
EHK9AyZbafH3s+QwFG5zIrv8gf6Fx5qItw==
-----END EC PRIVATE KEY-----`
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = "KEY-IDENTIFIER"

	block, _ := pem.Decode([]byte(JWT_SECRET))

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key ", err.Error())
		return "", err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println("[CreateToken] err:", err)
		return "", err
	}

	log.Println("[CreateToken] token:", tokenString)

	return tokenString, nil
}
