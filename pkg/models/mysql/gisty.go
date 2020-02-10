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
	stmt := `INSERT INTO gisty (title, content, created, expires)
	VALUES (?, ?, date(now()), adddate(now(), ?))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *GistModel) Get(id int) (*models.Gist, error) {
	s := &models.Gist{}
	err := m.DB.QueryRow(`SELECT id, title, content, created, expires 
	FROM gisty WHERE expires > UTC_TIMESTAMP() AND id = ?`, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *GistModel) Latest() ([]*models.Gist, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT id, title, content, created, expires FROM gisty 
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	// Use the Query() method to execute stmt
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// Never forget
	defer rows.Close()

	gists := []*models.Gist{}

	// Use rows.Next to iterate through the rows in the resultset
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		s := &models.Gist{}
		// Use rows.Scan() to copy the values from each field in the row to the
		// new Snippet object that we created. Again, the arguments to row.Scan
		// must be pointers to the place you want to copy the data into, and the
		// number of arguments must be exactly the same as the number of
		// columns returned by your statement.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		gists = append(gists, s)
	}
	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume that a successful iteration was completed
	// over the whole resultset.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// If everything went OK then return the Snippets slice.
	return gists, nil
}
