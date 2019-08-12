package transport

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var mockContact1, mockContact2 *models.Contact

func init() {
	workTag, _ := models.GetTag("work")
	personalTag, _ := models.GetTag("personal")
	mockContact1 = &models.Contact{
		Id:     uuid.New(),
		Name:   "mockContact1",
		Emails: []*models.Email{&models.Email{Id: "testemail@abc.com", Tag: workTag}},
		PhNums: []*models.PhNum{&models.PhNum{Number: "9988776655", Tag: workTag}, &models.PhNum{Number: "5544332211", Tag: personalTag}},
	}
	mockContact2 = &models.Contact{
		Id:     uuid.New(),
		Name:   "mockContact2",
		Emails: []*models.Email{&models.Email{Id: "testcont2@abc.com", Tag: workTag}},
		PhNums: []*models.PhNum{&models.PhNum{Number: "6666666666", Tag: personalTag}, &models.PhNum{Number: "1199228833", Tag: personalTag}},
	}
}

type mockContactService struct {
}

func (mockSvc *mockContactService) SaveContact(contact *models.Contact) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (mockSvc *mockContactService) FindContacts(filterMap map[models.Filter]string) ([]models.Contact, error) {
	return []models.Contact{*mockContact1, *mockContact2}, nil

}

func (mockSvc *mockContactService) DeleteContact(contactId uuid.UUID) error {
	return nil
}

func (mockSvc *mockContactService) GetContact(contactId uuid.UUID) (*models.Contact, error) {
	return mockContact2, nil
}

func TestCreateContact(t *testing.T) {
	data, err := ioutil.ReadFile("./fixtures/create_contact_req.json")
	assert.NoError(t, err)
	var contact *CreateContactRequest
	err = json.Unmarshal(data, &contact)
	assert.NoError(t, err)
	svc := &mockContactService{}
	handler := &contactHandler{
		contactSvc: svc,
	}
	req, err := http.NewRequest("POST", "/contact", strings.NewReader(string(data)))
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	httphandler := http.HandlerFunc(handler.createContact)
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var respBody ContactId
	err = json.Unmarshal([]byte(rr.Body.String()), &respBody)
	assert.NoError(t, err)
	_, err = uuid.Parse(respBody.Id)
	assert.NoError(t, err)
}

func TestCreateContactErrors(t *testing.T) {
	data, err := ioutil.ReadFile("./fixtures/create_contact_req_errors.json")
	assert.NoError(t, err)
	var contacts []*CreateContactRequest
	err = json.Unmarshal(data, &contacts)
	assert.NoError(t, err)
	svc := &mockContactService{}
	handler := &contactHandler{
		contactSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.createContact)
	for _, contact := range contacts {
		contactByteArr, err := json.Marshal(contact)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/contact", strings.NewReader(string(contactByteArr)))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		httphandler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	}
}

func TestDeleteContact(t *testing.T) {
	contactId := uuid.New().String()
	svc := &mockContactService{}
	handler := &contactHandler{
		contactSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.deleteContact)
	req, err := http.NewRequest("DELETE", "/contact/"+contactId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": contactId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteContactError(t *testing.T) {
	contactId := "invalidId"
	svc := &mockContactService{}
	handler := &contactHandler{
		contactSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.deleteContact)
	req, err := http.NewRequest("DELETE", "/contact/"+contactId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": contactId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetContact(t *testing.T) {
	contactId := uuid.New().String()
	svc := &mockContactService{}
	handler := &contactHandler{
		contactSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.deleteContact)
	req, err := http.NewRequest("GET", "/contact/"+contactId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": contactId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetContactError(t *testing.T) {
	contactId := "invalidId"
	svc := &mockContactService{}
	handler := &contactHandler{
		contactSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.deleteContact)
	req, err := http.NewRequest("GET", "/contact/"+contactId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": contactId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
func TestSearchContact(t *testing.T) {
	svc := &mockContactService{}
	handler := &contactHandler{
		contactSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.searchContact)
	req, err := http.NewRequest("GET", "/contact/search?name=test&phnum=999&email=abc", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var respBody []*models.Contact
	err = json.Unmarshal([]byte(rr.Body.String()), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(respBody))
}
