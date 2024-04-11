package jwt_service

import (
	"database/sql"
	"mgtu/digital-trace/main-backend-service/internal/database"

	"github.com/pkg/errors"
)

func (j *JWT) findUserById(tx *sql.Tx, userId uint64) (*database.User, error) {
	findUserOut, err := j.db.FindUser(tx, database.User{
		Id: &userId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.FindUser()]")
	}
	if !findUserOut.IsFound {
		return nil, errors.New("unable to find user by id: [!findUserOut.IsFound]")
	}
	if len(findUserOut.User) != 1 {
		return nil, errors.New("[len(findUserOut.User) != 1]]")
	}

	return &findUserOut.User[0], nil
}

func (j *JWT) findSessionByIdAndToken(tx *sql.Tx, sessionId uint64, sessionToken []byte) (*database.Session, error) {
	findSessionOut, err := j.db.FindSession(tx, database.Session{
		Id:    &sessionId,
		Token: &sessionToken,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[j.db.FindSession]")
	}
	if !findSessionOut.IsFound {
		return nil, errors.New("сессия не найдена: [!findSessionOut.IsFound]")
	}
	if len(findSessionOut.Session) != 1 {
		return nil, errors.New("[len(findSessionOut.Session) != 1]")
	}

	return &findSessionOut.Session[0], nil
}
