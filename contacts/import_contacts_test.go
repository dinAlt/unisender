package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestImportContactsRequest_FieldNames(t *testing.T) {
	expectedFieldNames := test.RandomStringSlice(12, 36)
	var givenFieldNames []string

	expectedResponse := randomImportContactsResponse()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenFieldNames = strings.Split(req.FormValue("field_names"), ",")

		result := api.Response{
			Result: expectedResponse,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.ImportContacts(req, randomCollection()).
		FieldNames(expectedFieldNames...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedFieldNames, givenFieldNames) {
		t.Fatal("Expected and given field names should be equal")
	}
}

func TestImportContactsRequest_OverwriteTags(t *testing.T) {
	expectedOverwriteTags := 1
	var givenOverwriteTags int

	expectedResponse := randomImportContactsResponse()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOverwriteTags, _ = strconv.Atoi(req.FormValue("overwrite_tags"))

		result := api.Response{
			Result: expectedResponse,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.ImportContacts(req, randomCollection()).
		OverwriteTags().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOverwriteTags != givenOverwriteTags {
		t.Fatalf(`Param "overwrite_tags" should be %d, %d given`, expectedOverwriteTags, givenOverwriteTags)
	}
}

func TestImportContactsRequest_OverwriteLists(t *testing.T) {
	expectedOverwriteLists := 1
	var givenOverwriteLists int

	expectedResponse := randomImportContactsResponse()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOverwriteLists, _ = strconv.Atoi(req.FormValue("overwrite_lists"))

		result := api.Response{
			Result: expectedResponse,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.ImportContacts(req, randomCollection()).
		OverwriteLists().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOverwriteLists != givenOverwriteLists {
		t.Fatalf(`Param "overwrite_lists" should be %d, %d given`, expectedOverwriteLists, givenOverwriteLists)
	}
}

func TestImportContactsRequest_Execute(t *testing.T) {
	expectedResponse := randomImportContactsResponse()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedResponse,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResponse, err := contacts.ImportContacts(req, randomCollection()).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResponse, givenResponse) {
		t.Fatal("Expected and given responses should be equal")
	}
}

func randomCollection() (collection *contacts.Collection) {
	collection = contacts.NewCollection()

	n := test.RandomInt(12, 36)
	for i := 0; i < n; i++ {
		contact := collection.Email(test.RandomString(12, 36))

		listIDs := test.RandomInt64Slice(12, 32)
		for _, listID := range listIDs {
			contact.AddListID(listID, test.RandomTime(12, 365))
		}
	}

	return
}

func randomImportContactsResponse() *contacts.ImportContactsResponse {
	inserted := test.RandomInt(99, 999)
	updated := test.RandomInt(99, 999)
	deleted := test.RandomInt(99, 999)

	var logs []contacts.ImportContactsResponseLogMessage
	for i := 0; i < test.RandomInt(9, 99); i++ {
		logs = append(logs, contacts.ImportContactsResponseLogMessage{
			Index:   i,
			Code:    test.RandomString(12, 36),
			Message: test.RandomString(12, 36),
		})
	}

	return &contacts.ImportContactsResponse{
		Total:     inserted + updated + deleted,
		Inserted:  inserted,
		Updated:   updated,
		Deleted:   deleted,
		NewEmails: test.RandomInt(99, 999),
		Invalid:   test.RandomInt(99, 999),
		Log:       logs,
	}
}
