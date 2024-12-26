package models

type Task struct {
	Id          int
	Description string
	Attachments []string
	AssigneeId  int // USER ID
}
