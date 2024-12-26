package models

import "time"

type User struct {
	Id        int
	MongoId   string
	Name      string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
