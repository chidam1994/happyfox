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
	ContactId uuid.UUID `db:"contact_id" json:"-"`
	Id        string    `db:"email_id" json:"email_id"`
	Tag       Tag       `db:"tag" json:"tag"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

type PhNum struct {
	ContactId uuid.UUID `db:"contact_id" json:"-"`
	Number    string    `db:"phnum" json:"phnum"`
	Tag       Tag       `db:"tag" json:"tag"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

type Contact struct {
	Id        uuid.UUID `db:"id, primarykey" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
	Emails    []*Email  `db:"-" json:"emails,omitempty"`
	PhNums    []*PhNum  `db:"-" json:"phnums,omitempty"`
}

type Group struct {
	Id        uuid.UUID `db:"id, primarykey" json:"id"`
	Name      string    `db:"name" json:"name"`
	Members   []*Member `db:"-" json"members"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

type Member struct {
	GroupId   uuid.UUID `db:"group_id" json:"-"`
	MemberId  uuid.UUID `db:"member_id" json:"member_id"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}
