package mysql

import (
	"database/sql"

	"github.com/gisty/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type GistModel struct {
	DB *sql.DB
}

func (m *GistModel) Insert(title, content, expires string) (int, error) {
	// Write the SQL statement we want to execute.
	stmt := `INSERT INTO gisty (title, content, created, expires)
	VALUES (?, ?, date(now()), adddate(now(), ?))`

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. This method returns a sql.Result object, which contains some basic
	// information about what happened when the statement was execute
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result object to get the ID of our
	// newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type before returning
	return int(id), nil
}

func (m *GistModel) Get(id int) (*models.Gist, error) {
	return nil, nil
}

func (m *GistModel) Latest() ([]*models.Gist, error) {
	return nil, nil
}
