package mailvault

import (
	"database/sql"
	"fmt"
)

type MailVaultConfig struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DbName string
}

func CreateMailVault(cfg MailVaultConfig) (mailVault *MailVault, err error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}
	mailVault.db = db
	err = mailVault.initialize()
	return
}

type MailVault struct {
	db *sql.DB
}

type Passport struct {
	id       string
	username string
}
