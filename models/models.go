package models

import (
	"fmt"
	"strings"
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

func GetTag(tagstr string) (Tag, error) {
	tags := [...]string{"Work", "Personal"}
	for i := range tags {
		if strings.ToLower(tags[i]) == strings.ToLower(tagstr) {
			return Tag(i), nil
		}
	}
	return Tag(0), fmt.Errorf("error converting string: %s to Tag", tagstr)
}

type Email struct {
	ContactId uuid.UUID `db:"contact_id"`
	Id        string    `db:"email_id"`
	Tag       Tag       `db:"tag"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PhNum struct {
	ContactId uuid.UUID `db:"contact_id"`
	Number    string    `db:"phnum"`
	Tag       Tag       `db:"tag"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Contact struct {
	Id        uuid.UUID `db:"id, primarykey"`
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
