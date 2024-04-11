package jwt_service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type accessTokenCustomClaims struct {
	UserId uint64 `json:"user_id"`
	RoleId uint64 `json:"role_id"`
	Mail   string `json:"mail"`
	Type   string `json:"type"`
	jwt.StandardClaims
}

type generateAccessTokenIn struct {
	UserId uint64
	Mail   string
	RoleId uint64
}

func (j *JWT) generateAccessToken(params generateAccessTokenIn) (string, error) {
	claims := accessTokenCustomClaims{
		UserId: params.UserId,
		RoleId: params.RoleId,
		Mail:   params.Mail,
		Type:   accessTokenType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.iss,
			ExpiresAt: time.Now().Add(j.accessTokenExp).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signing)
}

func (j *JWT) parseAccessToken(token string) (*jwt.Token, error) {
	validatedToken, err := jwt.ParseWithClaims(token, &accessTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
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

func (j *JWT) prepareAccessToken(token *jwt.Token) (*accessTokenCustomClaims, error) {
	claims, ok := token.Claims.(*accessTokenCustomClaims)
	if !ok {
		return nil, errors.New("unable to get claims: [token.Claims.(accessTokenCustomClaims)]")
	}
	if claims.Type != accessTokenType {
		return nil, errors.New("invalid token type: [claims.Type != accessTokenType]")
	}

	iss := claims.Issuer
	if iss != j.iss {
		return nil, errors.New("invalid iss name: [iss != j.iss]")
	}

	return claims, nil
}
