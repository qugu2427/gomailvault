package mailvault

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (mv *MailVault) Login(username, password string) (passport *Passport, err error) {
	var id string
	var passwordSalt string
	var passwordHash string
	err = mv.db.QueryRow(queryStrLogin, username).Scan(&id, &passwordSalt, &passwordHash)
	if err == nil {
		hashErr := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordSalt+password))
		if hashErr == nil {
			passport = nil
		} else {
			passport = &Passport{id: id, username: username}
			_, err = mv.db.Exec(queryStrUpdateLastLogin, passport.id)
		}
	}
	return
}

func (mv *MailVault) CreateUser(username, domain, password string, deleteAfterDays int) error {
	passwordSalt := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(passwordSalt+password), 12)
	if err != nil {
		return err
	} else {
		passwordHash := string(passwordHashBytes)
		_, err = mv.db.Exec(
			queryStrCreateUser,
			username, domain, passwordSalt, passwordHash, deleteAfterDays)
		return err
	}
}

func (mv *MailVault) DeleteUser(passport *Passport) (err error) {
	_, err = mv.db.Exec(queryStrDelUser, passport.id)
	return
}

func (mv *MailVault) DeleteExpiredUsers() (err error) {
	_, err = mv.db.Exec(queryStrDelExpiredUsers)
	return
}
