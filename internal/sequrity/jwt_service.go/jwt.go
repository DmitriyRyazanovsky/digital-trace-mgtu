package jwt_service

import (
	"mgtu/digital-trace/main-backend-service/internal/database"
	"mgtu/digital-trace/main-backend-service/internal/features"
	"mgtu/digital-trace/main-backend-service/internal/features/logging"
	"time"

	"github.com/pkg/errors"
)

const (
	accessTokenType  = "access"
	refreshTokenType = "refresh"
)

type JWT struct {
	db  *database.Database
	iss string // Издатель тот кто генерит jwt токены
	log logging.Logger

	sessionTokenLen   int
	sessionSigningLen int
	accessTokenExp    time.Duration
	refreshTokenExp   time.Duration

	signing []byte
}

type NewJWtServiceIn struct {
	Db  *database.Database
	Iss string // Издатель, тот кто генерит jwt токены
	Log logging.Logger

	SigningFilePath string

	SessionTokenLen   int
	SessionSigningLen int
	AccessTokenExp    time.Duration
	RefreshTokenExp   time.Duration
}

func NewJWtService(params NewJWtServiceIn) (*JWT, error) {
	exist, err := features.CreateFileIfNotExists(params.SigningFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "[features.CreateFileIfNotExists(params.SigningFilePath)]")
	}
	if !exist {
		content, err := features.GenerateRandomBytes(params.SessionSigningLen)
		if err != nil {
			return nil, errors.Wrap(err, "[features.GenerateRandomBytes(params.SessionSigningLen)]")
		}
		err = features.WriteContentInFile(params.SigningFilePath, string(content))
		if err != nil {
			return nil, errors.Wrap(err, "[features.WriteContentInFile(params.SigningFilePath, string(content))]")
		}
	}

	signing, err := features.ReadContentFromFile(params.SigningFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "[features.ReadContentFromFile(params.SigningFilePath)]")
	}

	return &JWT{
		db:                params.Db,
		iss:               params.Iss,
		log:               params.Log,
		sessionTokenLen:   params.SessionTokenLen,
		sessionSigningLen: params.SessionSigningLen,
		accessTokenExp:    params.AccessTokenExp,
		refreshTokenExp:   params.RefreshTokenExp,
		signing:           signing,
	}, nil
}

type CreateSessionIn struct {
	UserId uint64
}

type CreateSessionOut struct {
	AccessToken  string
	RefreshToken string
}

func (j *JWT) CreateSession(params CreateSessionIn) (*CreateSessionOut, error) {
	tx, err := j.db.OpenTransaction()
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.OpenTransaction()]")
	}

	user, err := j.findUserById(tx, params.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "[j.findUserById(tx, *session.UserId)]")
	}

	sessionToken, err := features.GenerateRandomBytes(j.sessionTokenLen)
	if err != nil {
		return nil, errors.Wrap(err, "[features.GenerateRandomBytes(j.sessionTokenLen)]")
	}

	findSessionOut, err := j.db.AddSession(tx, database.Session{
		UserId: &params.UserId,
		Token:  &sessionToken,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.AddSession()]")
	}

	out := &CreateSessionOut{}

	accessToken, err := j.generateAccessToken(generateAccessTokenIn{
		UserId: params.UserId,
		Mail:   *user.Email,
		RoleId: *user.RoleId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.generateAccessToken()]")
	}
	out.AccessToken = accessToken

	refreshToken, err := j.generateRefreshToken(generateRefreshTokenIn{
		UserId:       *user.Id,
		SessionId:    *findSessionOut.Session.Id,
		SessionToken: sessionToken,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.generateRefreshToken()]")
	}
	out.RefreshToken = refreshToken

	err = j.db.CommitTransaction(tx)
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.CommitTransaction(tx)]")
	}

	return out, err
}

func (j *JWT) ValidateAccessToken(tokenString string) (*accessTokenCustomClaims, error) {
	token, err := j.parseAccessToken(tokenString)
	if err != nil {
		return nil, errors.Wrap(err, "[j.parseAccessToken(tokenString)]")
	}

	accessTokenOut, err := j.prepareAccessToken(token)
	if err != nil {
		return nil, errors.Wrap(err, "[j.prepareAccessToken(token)]")
	}

	return accessTokenOut, nil
}

func (j *JWT) ValidateRefreshToken(tokenString string) (*refreshTokenCustomClaims, error) {
	token, err := j.parseRefreshToken(tokenString)
	if err != nil {
		return nil, errors.Wrap(err, "[j.parseRefreshToken(tokenString)]")
	}

	refreshTokenOut, err := j.prepareRefreshToken(token)
	if err != nil {
		return nil, errors.Wrap(err, "[j.parseRefreshToken(token)]")
	}

	return refreshTokenOut, nil
}

type UpdateSessionOut struct {
	AccessToken  string
	RefreshToken string
}

func (j *JWT) ReloadSession(tokenString string) (*UpdateSessionOut, error) {
	//* Создаём транзакцию
	tx, err := j.db.OpenTransaction()
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.OpenTransaction()]")
	}

	//* Проводим валидацию refresh токена для того
	payloads, err := j.ValidateRefreshToken(tokenString)
	if err != nil {
		return nil, errors.Wrap(err, "[j.ValidateRefreshToken(tokenString, j.signing)]")
	}

	//* Ищем сессию в БД, по id сессии и токену сессии указанных в refresh
	//* для получения id пользователя
	session, err := j.findSessionByIdAndToken(tx, payloads.SessionId, payloads.SessionToken)
	if err != nil {
		return nil, errors.Wrap(err, "[j.findSessionByIdAndToken(tx, payloads.SessionId, payloads.SessionToken)]")
	}

	//* Ищем сессию
	user, err := j.findUserById(tx, *session.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "[j.findUserById(tx, *session.UserId)]")
	}

	//* создаём объект на return
	out := &UpdateSessionOut{}

	//* генерируем новый access токен
	accessToken, err := j.generateAccessToken(generateAccessTokenIn{
		UserId: *user.Id,
		Mail:   *user.Email,
		RoleId: *user.RoleId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.generateAccessToken()]")
	}
	out.AccessToken = accessToken

	//* генерируем новый токен для сессии из БД
	sessionToken, err := features.GenerateRandomBytes(j.sessionTokenLen)
	if err != nil {
		return nil, errors.Wrap(err, "[features.GenerateRandomBytes(j.sessionTokenLen)]")
	}

	//* изменяем токен сессии в БД на новый токен
	_, err = j.db.ChangeSession(tx, database.Session{
		Token: &sessionToken,
	}, database.Session{
		Id: session.Id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.generateRefreshToken()]")
	}

	//* создаём refresh токен
	refreshToken, err := j.generateRefreshToken(generateRefreshTokenIn{
		UserId:       *user.Id,
		SessionId:    *session.Id,
		SessionToken: sessionToken,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.generateRefreshToken()]")
	}
	out.RefreshToken = refreshToken

	//* При успешной попытке завершаем транзакцию
	err = j.db.CommitTransaction(tx)
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.CommitTransaction(tx)]")
	}
	return out, nil
}
