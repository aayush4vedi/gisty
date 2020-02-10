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
	stmt := `SELECT id, title, content, created, expires FROM gisty 
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// Never forget
	defer rows.Close()

	gists := []*models.Gist{}

	for rows.Next() {
		s := &models.Gist{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		gists = append(gists, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gists, nil
}
