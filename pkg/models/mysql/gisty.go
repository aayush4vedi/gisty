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
	// Write the SQL statement we want to execute
	stmt := `SELECT id, title, content, created, expires FROM gisty 
	WHERE expires > UTC_TIMESTAMP() AND id = ?` 
	
	// Use the QueryRow() method on the connection pool to execute the SQL statement.
	// This returns a pointer to a sql.Row object which holds the result from the database. 
	row := m.DB.QueryRow(stmt, id) 
	
	// Initialize a pointer to a new zeroed Snippet struct. 
	s := &models.Gist{} 
	
	// Use row.Scan() to copy the values from each field in sql.Row to the 
	// corresponding field in the Snippet struct. Notice that the arguments 
	// to row.Scan are *pointers* 
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires) 
	if err == sql.ErrNoRows { 
		return nil, models.ErrNoRecord 
	} else if err != nil { 
		return nil, err 
	} 
	
	// If everything went OK then return the Snippet object. 
	return s, nil 
}

func (m *GistModel) Latest() ([]*models.Gist, error) {
	return nil, nil
}
