package jwt_service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type refreshTokenCustomClaims struct {
	UserId       uint64 `json:"user_id"`
	SessionId    uint64 `json:"session_id"`
	SessionToken []byte `json:"session_token"`
	Type         string `json:"type"`
	jwt.StandardClaims
}

type generateRefreshTokenIn struct {
	UserId       uint64
	SessionId    uint64
	SessionToken []byte
}

func (j *JWT) generateRefreshToken(params generateRefreshTokenIn) (string, error) {
	claims := refreshTokenCustomClaims{
		UserId:       params.UserId,
		SessionId:    params.SessionId,
		SessionToken: params.SessionToken,
		Type:         refreshTokenType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.iss,
			ExpiresAt: time.Now().Add(j.refreshTokenExp).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signing)
}

func (j *JWT) parseRefreshToken(token string) (*jwt.Token, error) {
	validatedToken, err := jwt.ParseWithClaims(token, &refreshTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signing, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "[jwt.Parse()]")
	}

	if !validatedToken.Valid {
		return nil, errors.New("[!validatedToken.Valid]")
	}

	return validatedToken, nil
}

func (j *JWT) prepareRefreshToken(token *jwt.Token) (*refreshTokenCustomClaims, error) {
	claims, ok := token.Claims.(*refreshTokenCustomClaims)
	if !ok {
		return nil, errors.New("unable to get claims: [token.Claims.(refreshTokenCustomClaims)]")
	}
	if claims.Type != refreshTokenType {
		return nil, errors.New("invalid token type: [claims.Type != refreshTokenType]")
	}

	iss := claims.Issuer
	if iss != j.iss {
		return nil, errors.New("invalid iss name: [iss != j.iss]")
	}

	return claims, nil
}
