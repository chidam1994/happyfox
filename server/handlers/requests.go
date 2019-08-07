package handlers

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/chidam1994/happyfox/models"
)

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var rxPhNum = regexp.MustCompile("^(\\d{10})$")

type CreateContactRequest struct {
	Name   string   `json:"name"`
	Emails []*Email `json:"emails"`
	PhNums []*PhNum `json:"phnums"`
}

type ContactId struct {
	Id string `json:"id"`
}

type Email struct {
	Id  string `json:"email_id"`
	Tag string `json:"tag"`
}

type PhNum struct {
	Number string `json:"number"`
	Tag    string `json:"tag"`
}

func (request *CreateContactRequest) Validate() (*models.Contact, error) {
	result := &models.Contact{}
	if request.Name == "" {
		return nil, errors.New("name cannot be blank")
	}
	result.Name = request.Name
	for i := range request.Emails {
		if request.Emails[i].Id == "" {
			return nil, errors.New("email_id cant be empty")
		}
		if !rxEmail.Match([]byte(request.Emails[i].Id)) {
			return nil, fmt.Errorf("the email_id %s is not valid", request.Emails[i].Id)
		}
		tag, err := models.GetTag(request.Emails[i].Tag)
		if err != nil {
			return nil, errors.New("email tag can only either of \"work\" or \"personal\" values")
		}
		email := &models.Email{
			Id:  request.Emails[i].Id,
			Tag: tag,
		}
		result.Emails = append(result.Emails, email)
	}
	for i := range request.PhNums {
		if request.PhNums[i].Number == "" {
			return nil, errors.New("phone number cant be empty")
		}
		if !rxPhNum.Match([]byte(request.PhNums[i].Number)) {
			return nil, fmt.Errorf("the number %s is not valid", request.PhNums[i].Number)
		}
		tag, err := models.GetTag(request.Emails[i].Tag)
		if err != nil {
			return nil, errors.New("email tag can only either of \"work\" or \"personal\" values")
		}
		phNum := &models.PhNum{
			Number: request.Emails[i].Id,
			Tag:    tag,
		}
		result.PhNums = append(result.PhNums, phNum)
	}
	return result, nil
}
