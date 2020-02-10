package transaction_demo

import "database/sql"

type ExampleModel struct {
	DB *sql.DB
}

func (m *ExampleModel) ExampleTransaction() error {
	// Calling the Begin() method on the connection pool creates a new sql.Tx
	// object, which represents the in-progress database transaction.
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	// Call Exec() on the transaction, passing in your statement and any
	// parameters. It's important to notice that tx.Exec() is called on the
	// transaction object just created, NOT the connection pool. Although we're
	// using tx.Exec() here you can also use tx.Query() and tx.QueryRow() in
	// exactly the same way.
	_, err = tx.Exec("INSERT INTO ...") //full statement goes here
	if err != nil {
		// If there is any error, we call the tx.Rollback() method on the
		// transaction. This will abort the transaction and no changes will be
		// made to the database.
		tx.Rollback()
		return err
	}
	// Carry out another transaction in exactly the same way.
	_, err = tx.Exec("UPDATE ...") //full statement goes here
	if err != nil {
		tx.Rollback()
		return err
	}
	// If there are no errors, the statements in the transaction can be committe
	// to the database with the tx.Commit() method. It's really important to ALW
	// call either Rollback() or Commit() before your function returns. If you
	// don't the connection will stay open and not be returned to the connectio
	// pool. This can lead to hitting your maximum connection limit/running out
	// resources.
	err = tx.Commit()
	return err
}
