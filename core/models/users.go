package models

type User struct {
	ID            string `gorm:"size:512"`
	FirstName     string `gorm:"size:512"`
	LastName      string `gorm:"size:512"`
	Age           int
	RecordingDate int64
}
