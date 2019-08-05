package models

import (
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
	Id     uuid.UUID
	Name   string
	Emails []Email
	PhNums []PhNum
}

type Group struct {
	Id      uuid.UUID
	Name    string
	Members []uuid.UUID
}
