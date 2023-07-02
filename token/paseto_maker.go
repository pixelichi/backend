package token

import (
	"log"
	"net/http"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
	"shinypothos.com/api/common/server_error"
)

// PasetoMaker is a Paseto token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) Maker {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		log.Fatalf("symmetric key size: must be exactly %d characters long.", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker
}

func (maker *PasetoMaker) CreateToken(userID int64, duration time.Duration) (string, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) CreateTokenOrAbort(ctx *gin.Context, userID int64, duration time.Duration) string {
	token, err := maker.CreateToken(userID, duration)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError(err.Error()))
	}

	return token
}

// Check if the token is valid or not and if it is, it will return the payload in the token
func (maker *PasetoMaker) VerifyToken(token string) (Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return *payload, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return *payload, err
	}

	return *payload, nil
}
