package mailvault

const (
	queryStrInit string = `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS user_account (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		username TEXT UNIQUE NOT NULL,
		domain TEXT NOT NULL,
		password_salt VARCHAR(6) NOT NULL,
		password_hash TEXT NOT NULL,
		delete_after_days INTEGER DEFAULT 365 NOT NULL,
		last_login TIMESTAMP DEFAULT NOW() NOT NULL,
		created TIMESTAMP DEFAULT NOW() NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS mail (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		mail_from TEXT NOT NULL,
		rcpt_to TEXT NOT NULL,
		body TEXT NOT NULL,
		recieved TIMESTAMP DEFAULT NOW() NOT NULL,
		rcpt_username TEXT NOT NULL REFERENCES user_account(username) ON DELETE CASCADE
	);
	`
	queryStrLogin            string = "SELECT id, password_salt, password_hash FROM user_account WHERE username=$1"
	queryStrDelUser          string = "DELETE FROM user_account WHERE id=$1"
	queryStrUpdateLastLogin  string = "UPDATE user_account SET last_login=NOW() WHERE id=$1"
	queryStrDelExpiredUsers  string = "DELETE FROM user_account WHERE EXTRACT(DAY FROM age(last_login, NOW())) > delete_after_days"
	queryStrCreateUser       string = "INSERT INTO user_account (username, domain, password_salt, password_hash, delete_after_days) VALUES ($1, $2, $3, $4, $5)"
	queryStrInsertMail       string = "INSERT INTO mail (mail_from, rcpt_to, body, key_encrypted, rcpt_username) VALUES ($1, $2, $3, $4, $5)"
	queryStrDelMail          string = "DELETE FROM mail WHERE id=$1"
	queryStrGetUserMailCount string = "SELECT COUNT(*) FROM mail WHERE rcpt_username=$1"
	queryStrGetMailCount     string = "SELECT COUNT(*) FROM mail"
)
