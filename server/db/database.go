package db

import (
	"database/sql"
	"errors"

	"github.com/NickBrisebois/InteractivePersistentStartpage/server/config"

	// go-sqlite3 is required so that we can open a sqlite3 database.
	_ "github.com/mattn/go-sqlite3"
)

// Schema is the database schema for the startpage API
type LinksSchema struct {
	ID   string
	Name string
	URL  string
}

// Database holds the database itself
type Database struct {
	DB *sql.DB
}

// NewDatabase creates a new sqlite3 database handling object
func NewDatabase(config *config.Config) (*Database, error) {
	db, err := sql.Open("sqlite3", config.DBPath)

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("db was nil")
	}

	// Create the tables if they don't already exist
	createLinksTable(db)

	return &Database{
		DB: db,
	}, nil
}

// createLinksTable creates the sqlite3 holding link data
func createLinksTable(db *sql.DB) error {
	linksTable := `
	CREATE TABLE IF NOT EXISTS links (
		ID INTEGER NOT NULL PRIMARY KEY,
		Name TEXT,
		URL TEXT
	);
	`

	if _, err := db.Exec(linksTable); err != nil {
		return err
	}

	return nil
}

// DBAddLink handles adding new links to the SQlite3 db
func (db *Database) DBAddLink(name string, url string) error {
	addStmt, err := db.DB.Prepare("INSERT INTO links (Name, URL) VALUES (?, ?)")

	if err != nil {
		return err
	}

	addStmt.Exec(name, url)
	return nil
}

// DBDelLink handles deleting links from the SQLite3 db
func (db *Database) DBDelLink(id int) error {
	delStmt, err := db.DB.Prepare("DELETE FROM links WHERE id=?")

	if err != nil {
		return err
	}

	delStmt.Exec(id)
	return nil
}
