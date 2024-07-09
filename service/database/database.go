/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"wasaphoto/service/api/models"
	
)

// AppDatabase is the high-level interface for the DB
type AppDatabase interface {
	AddPhoto(ownerID, address string, date time.Time, description string) (int64, error)
	DeletePhoto(userID, photoID string) (string, error)

	AddComment(userID, pictureID, content string) error
	DeleteComment(userID, pictureID, commentID string) error
	GetComments(pictureID string) ([]models.Comment, error)

	AddLike(userID, pictureID string) error
	DeleteLike(userID, pictureID string) error
	GetLikes(pictureID string) ([]string, error)

	AddBan(userID, bannedID string) error
	DeleteBan(userID, bannedID string) error
	GetBan(userID string) ([]string, error)
	GetIfBan(userID, bannerID string) (bool, error)

	AddFollower(followerID, followingID string) error
	DeleteFollower(followerID, followingID string) error
	GetFollowers(userID string) ([]models.User, error)
	GetFollowing(userID string) ([]models.User, error)

	AddUser(username string) (int64, error)
	Login(username string) (int64, error)
	SetUsername(userID, username string) error
	GetUsername(userID string) (string, error)
	GetFeed(userID string) ([]models.Picture, error)
	GetUsers(usernamePrefix string) ([]models.User, error)

	Ping() error
	GetDB() *sql.DB
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building an AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createDatabase(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}


func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

func (db *appdbimpl) GetDB() *sql.DB {
	return db.c
}

func createDatabase(db *sql.DB) error {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS ban (
			userID INTEGER NOT NULL,
			banID INTEGER NOT NULL,
			PRIMARY KEY(userID, banID)
		);`,

		`CREATE TABLE IF NOT EXISTS comments (
			commentID INTEGER UNIQUE,
			pictureID INTEGER NOT NULL,
			text TEXT NOT NULL,
			ownerID INTEGER NOT NULL,
			PRIMARY KEY(commentID AUTOINCREMENT)
		);`,

		`CREATE TABLE IF NOT EXISTS followers (
			FollowerID INTEGER NOT NULL,
			FollowingID INTEGER NOT NULL,
			PRIMARY KEY(FollowerID, FollowingID)
		);`,

		`CREATE TABLE IF NOT EXISTS likes (
			userID INTEGER NOT NULL,
			pictureID INTEGER NOT NULL,
			PRIMARY KEY(pictureID, userID)
		);`,

		`CREATE TABLE IF NOT EXISTS pictures (
			ownerID INTEGER NOT NULL,
			pictureID INTEGER NOT NULL UNIQUE,
			address TEXT,
			date TEXT,
			description TEXT,
			PRIMARY KEY(pictureID AUTOINCREMENT)
		);`,

		`CREATE TABLE IF NOT EXISTS users (
			userID INTEGER NOT NULL UNIQUE,
			username TEXT NOT NULL UNIQUE,
			PRIMARY KEY(userID)
		);`,
	}

	for _, sqlStmt := range tables {
		_, err := db.Exec(sqlStmt)
		if err != nil {
			return fmt.Errorf("error executing statement: %v\n%v", err, sqlStmt)
		}
	}

	return nil
}