package yodata_test

import (
	"github.com/juandunbar/yobirthday/yoconfig"
	"github.com/juandunbar/yobirthday/yodata"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

// This unit test is intended to test the functionality of the yodata package.
// It also unintentionally tests the yoconfig package

var dataService *yodata.DataService

// We will setup an im memory test database instead of mocking ( could mock using https://github.com/DATA-DOG/go-sqlmock )
func setupConfig() error {
	config, err := yoconfig.Load("../")
	if err != nil {
		return err
	}
	// Lets override some values to allow testing
	config.Set("database.dsn", "file::memory:?cache=shared")
	config.Set("sender_nickname", "test")
	// Set our config in the yodata package
	yodata.SetConfig(config)
	return nil
}
// Create our DataService object, it will use our set config values from above to allow testing
func setupService() error {
	service, err := yodata.NewService()
	if err != nil {
		return err
	}
	dataService = service
	return nil
}
// Create our test database and insert some test data
func createData() error {
	// Create our test birthday table
	createSql := `
		CREATE TABLE IF NOT EXISTS Birthdays (
		birthday_id INTEGER PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		nickname TEXT,
		email TEXT,
		phone TEXT,
		birthdate datetime,
		type TEXT);`
	_, err := dataService.Db.Exec(createSql)
	if err != nil {
		return err
	}
	// Get todays date, and not todays date in YYYY-MM-DD format
	today := time.Now().Format("2006-01-02")
	notToday := time.Now().AddDate(0, 0, 5).Format("2006-01-02")
	// Create our test birthday rows
	params := []interface{}{"john", "livingston", "jonny", "test@gmail.com", "555-867-5309", today, "default"}
	params = append(params, "sarah", "carter", "sarah", "test@yahoo.com", "555-444-7685", notToday, "friend")
	insertSql := `
		INSERT INTO Birthdays(first_name, last_name, nickname, email, phone, birthdate, type)
		VALUES(?, ?, ?, ?, ?, ?, ?),(?, ?, ?, ?, ?, ?, ?);`
	_, err = dataService.Db.Exec(insertSql, params...)
	if err != nil {
		return err
	}
	return nil
}

func TestDataService_GetEmails(t *testing.T) {
	// Setup config
	if err := setupConfig(); err != nil {
		t.Fatalf("failed to setup config with error: %v", err.Error())
	}
	// Setup our test db
	if err := setupService(); err != nil {
		t.Fatalf("failed to init test database with error: %v", err.Error())
	}
	// Insert test data into test db
	if err := createData(); err != nil {
		t.Fatalf("failed to create test data with error: %v", err.Error())
	}
	// Test our GetEmails code
	birthdays, err := dataService.GetEmails()
	if err != nil {
		t.Fatalf("failed to get emails from data with error: %v", err.Error())
	}
	// we expect 1 row returned
	got := len(birthdays)
	if got != 1 {
		t.Logf("%+v", birthdays)
		t.Fatalf("unexpected number of rows returned: %v, wanted: 1", got)
	}
}