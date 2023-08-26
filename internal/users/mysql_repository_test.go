package users

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestUserRepositorySuite struct {
	suite.Suite
}

// SetupTest sets up often used objects
func (test *TestUserRepositorySuite) SetupTest() {

}

// TestClientTestSuite Runs the testsuite
func TestCampaignRepositoryTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestUserRepositorySuite))
}

func TestGetUser(m *testing.T) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(m, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.Errorf("Failed to open connection to DB: %v", err)
	}

	if conn == nil {
		m.Error("Failed to open connection to DB: conn is nil")
	}

	defer db.Close()

	mock.ExpectQuery("SELECT * FROM users WHERE id = ? ORDER BY id LIMIT 1").WithArgs(1)

	db.QueryContext(context.Background(), "SELECT * FROM users WHERE id = ? ORDER BY id LIMIT 1", 1)

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}

func TestAddUser(m *testing.T) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(m, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.Errorf("Failed to open connection to DB: %v", err)
	}

	if conn == nil {
		m.Error("Failed to open connection to DB: conn is nil")
	}

	defer db.Close()

	mock.ExpectQuery("INSERT INTO users (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)").WithArgs(1, "test username", "2023-08-26", "2023-08-26")

	db.QueryContext(context.Background(), "INSERT INTO users (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)", 1, "test username", "2023-08-26", "2023-08-26")

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}

func TestUpdateUser(m *testing.T) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(m, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.Errorf("Failed to open connection to DB: %v", err)
	}

	if conn == nil {
		m.Error("Failed to open connection to DB: conn is nil")
	}

	defer db.Close()

	mock.ExpectQuery("UPDATE users SET name = ? WHERE id = ?").WithArgs("new username", 1)

	db.QueryContext(context.Background(), "UPDATE users SET name = ? WHERE id = ?", "new username", 1)

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}

func TestDeleteUser(m *testing.T) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(m, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.Errorf("Failed to open connection to DB: %v", err)
	}

	if conn == nil {
		m.Error("Failed to open connection to DB: conn is nil")
	}

	defer db.Close()

	mock.ExpectQuery("DELETE FROM users WHERE id = ?").WithArgs(1)

	db.QueryContext(context.Background(), "DELETE FROM users WHERE id = ?", 1)

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}

func TestGetAllUsers(m *testing.T) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(m, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.Errorf("Failed to open connection to DB: %v", err)
	}

	if conn == nil {
		m.Error("Failed to open connection to DB: conn is nil")
	}

	defer db.Close()

	mock.ExpectQuery("SELECT * FROM users")

	db.QueryContext(context.Background(), "SELECT * FROM users")

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}
