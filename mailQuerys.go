package mailvault

func (mv *MailVault) InsertMail(mailFrom, rcptTo, body, rcptUsername string) (err error) {
	_, err = mv.db.Exec(
		queryStrInsertMail,
		mailFrom,
		rcptTo,
		body,
		rcptUsername)
	return err
}

func (mv *MailVault) DeleteMail(ids []string) (err error) {
	for _, id := range ids {
		_, err = mv.db.Exec(
			queryStrDelMail,
			id,
		)
		if err != nil {
			return err
		}
	}
	return
}

func (mv *MailVault) CountUserMail(password *Passport) (count int, err error) {
	err = mv.db.QueryRow(queryStrGetUserMailCount, password.username).Scan(&count)
	return
}

func (mv *MailVault) CountMail() (count int, err error) {
	err = mv.db.QueryRow(queryStrGetMailCount).Scan(&count)
	return
}
