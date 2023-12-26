package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub      string
	EventID  string `json:"event_id"`
	TokenUse string `json:"token_use"`
	Scope    string
	AuthTime int `json:"auth_time"`
	Iss      string
	Exp      int
	Iat      int
	ClientID string `json:"client_id"`
	Username string
}

func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("El token no es valido")
		return false, nil, "El token no es valido"
	}

	userInfo, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token: " + err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON

	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura JSPON " + err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Fecha de eexpiracion de token = " + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expirado !!"
	}
	return true, nil, string(tkj.Username)
}
