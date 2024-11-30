package models

import "time"

type User struct {
	Id        int
	MongoId   string
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Task struct {
	Id          int
	Description string
	Attachments []string
	AssigneeId  int // USER ID
}
