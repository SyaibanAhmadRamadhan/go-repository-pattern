package helpers

import "database/sql"

func Defer(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfErr(errRollback)
		panic(err)
	} else {
		err := tx.Commit()
		PanicIfErr(err)
	}
}
