package models

type User struct {
	Id       int64  `json:"id" bson:"id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
