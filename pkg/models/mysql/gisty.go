package mysql

import (
	"database/sql"

	"github.com/gisty/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type GistModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *GistModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *GistModel) Get(id int) (*models.Gist, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *GistModel) Latest() ([]*models.Gist, error) {
	return nil, nil
}
