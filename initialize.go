package mailvault

func (mv *MailVault) initialize() (err error) {
	_, err = mv.db.Query(queryStrInit)
	return err
}
