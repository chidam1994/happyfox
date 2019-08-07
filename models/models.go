package models

import (
	"time"

	"github.com/google/uuid"
)

type Tag int

const (
	Work Tag = iota
	Personal
)

func (tag Tag) String() string {
	tags := [...]string{"Work", "Personal"}
	return tags[tag]
}

type Email struct {
	Id  string
	Tag Tag
}

type PhNum struct {
	Number string
	Tag    Tag
}

type Contact struct {
	Id        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Emails    []*Email  `db:"-"`
	PhNums    []*PhNum  `db:"-"`
}

type Group struct {
	Id      uuid.UUID
	Name    string
	Members []uuid.UUID
}
