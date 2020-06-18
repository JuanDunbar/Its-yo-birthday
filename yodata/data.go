package yodata

import (
	"database/sql"
	"fmt"
	"github.com/juandunbar/yobirthday/yoconfig"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var config *yoconfig.Config

type DataService struct {
	Db *sql.DB
}

type Email struct {
	FirstName string
	LastName  string
	Nickname  string
	Email     string
	Phone     string
	Birthday  string
	Type 	  string
	SenderNickname string
}

func NewService() (*DataService, error) {
	// Open a database connection using our config dsn
	dsn := config.GetString("database.dsn")
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	// Test our connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// Create the DataService object and return
	ds := &DataService{Db: db}
	return ds, nil
}

// This function will hit our database and return any birthdays that match today's date
func (ds *DataService) GetEmails() ([]Email, error) {
	// Get any matching birthdays for our current date
	query, sqlParam := ds.GetEmailQuery()
	rows, err := ds.Db.Query(query, sqlParam)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	senderNickname := config.GetString("sender_nickname")
	emails := make([]Email, 0)
	for rows.Next() {
		var email Email
		err = rows.Scan(&email.FirstName, &email.LastName, &email.Nickname,
			&email.Email, &email.Phone, &email.Birthday, &email.Type)
		if err != nil {
			return nil, err
		}
		// Set our sender nickname
		email.SenderNickname = senderNickname
		emails = append(emails, email)
	}

	return emails, nil
}

// This function generates a query to get any birthdays for today's date
// sqlParam needs to be in format "DD:MM"
func (ds *DataService) GetEmailQuery() (query string, sqlParam string) {
	_, month, day := time.Now().Date()
	strMonth, strDay := fixMonthDay(month, day)
	sqlParam = fmt.Sprintf("%v:%v", strMonth, strDay)
	query = `SELECT first_name, last_name, nickname, 
				email, phone, birthdate, type 
			FROM Birthdays
			WHERE strftime('%m:%d', birthdate) = ?`

	return
}

// This function is to get our month and day values into the same format as our sql query
// example: "01" vs "1"
func fixMonthDay(month time.Month, day int) (string, string) {
	strMonth := ""
	strDay := ""
	// Lets get our month into "MM" format
	intMonth := int(month)
	if intMonth < 10 {
		strMonth = fmt.Sprintf("0%v", intMonth)
	} else {
		strMonth = fmt.Sprintf("%v", intMonth)
	}
	// Lets get our day into "DD" format
	if day < 10 {
		strDay = fmt.Sprintf("0%v", day)
	} else {
		strDay = fmt.Sprintf("%v", day)
	}
	return strMonth, strDay
}

func SetConfig(c *yoconfig.Config) {
	config = c
}