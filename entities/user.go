package entities

import (
	"time"
)

type UserID string

func NewUser(id UserID, name string, age int, createdAt time.Time) User {
	return User{
		id:        id,
		name:      name,
		age:       age,
		createdAt: createdAt,
	}
}

func (a UserID) String() string {
	return string(a)
}

type (
	User struct {
		id        UserID
		name      string
		age       int
		createdAt time.Time
	}
)

func (a User) ID() UserID {
	return a.id
}

func (a User) Name() string {
	return a.name
}

func (a User) Age() int {
	return a.age
}

func (a User) CreatedAt() time.Time {
	return a.createdAt
}
