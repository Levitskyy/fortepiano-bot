package model

import "time"

type User struct {
	Id        int64
	Name      string
	Email     string
	CreatedAt time.Time
}

type Group struct {
	Id         int64
	Name       string
	InviteLink string
}

type Subscription struct {
	Id      int
	UserId  int64
	GroupId int64
	EndDate time.Time
}
