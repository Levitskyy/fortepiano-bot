package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Email     sql.NullString
	CreatedAt time.Time `db:"created_at"`
}

type Group struct {
	Id         int64
	Name       string
	InviteLink string `db:"invite_link"`
}

type Subscription struct {
	Id      int
	UserId  int64     `db:"user_id"`
	GroupId int64     `db:"group_id"`
	EndDate time.Time `db:"end_date"`
}
