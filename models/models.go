package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Tag string

const (
	Work     Tag = "work"
	Personal Tag = "personal"
)

func (tag Tag) String() string {
	tags := map[Tag]string{
		Work:     "work",
		Personal: "personal",
	}
	return tags[tag]
}

func GetTag(tagstr string) (Tag, error) {
	tags := map[string]Tag{
		"work":     Work,
		"personal": Personal,
	}
	tag, ok := tags[tagstr]
	if !ok {
		return Tag(""), fmt.Errorf("error converting string: %s to Tag", tagstr)

	}
	return tag, nil
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
