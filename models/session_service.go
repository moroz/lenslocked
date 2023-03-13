package models

import "database/sql"

type SessionService struct {
	DB        *sql.DB
	CookieKey string
}

func getUser(db *sql.DB, id int) *User {
	user := User{
		ID: id,
	}
	row := db.QueryRow(`select email from users where id = $1`, id)
	err := row.Scan(&user.Email)
	if err != nil {
		return nil
	}
	return &user
}

func AuthenticateUserByToken(db *sql.DB, token string) *User {
	id := DecodeSubjectFromAccessToken(token)
	if id == 0 {
		return nil
	}
	return getUser(db, id)
}
