package helper

import "github.com/jmoiron/sqlx"

func CommitOrRollback(tx *sqlx.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
