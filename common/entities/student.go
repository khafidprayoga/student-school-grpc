package entities

import "time"

type Student struct {
	Id        uint32        `gorm:"primaryKey;autoIncrement"`
	FullName  string        `gorm:"not null;size:256"`
	BirthDate time.Time     `gorm:"not null"`
	Class     string        `gorm:"not null;size:32"`
	Address   string        `gorm:"not null"`
	Phones    []PhoneNumber `foreignKey:"phone_number"`
}

func (s *Student) TableName() string {
	return "student"
}
