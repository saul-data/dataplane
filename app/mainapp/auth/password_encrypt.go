package auth

import (
	dpconfig "github.com/saul-data/dataplane/app/mainapp/config"

	"github.com/saul-data/dataplane/app/mainapp/logging"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(rawPassword string) (string, error) {
	password := []byte(rawPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		bcrypt.DefaultCost,
	)

	if err != nil {
		if dpconfig.Debug == "true" {
			logging.PrintSecretsRedact(err)
		}
		return "", err

	}

	return string(hashedPassword), nil
}
