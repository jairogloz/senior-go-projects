package server_test

import (
	"errors"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/mocks"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/server"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_SaveEmployee_ErrorSaving(t *testing.T) {
	reqBody := strings.NewReader(`{"name":"John Doe"}`)
	req, err := http.NewRequest("POST", "/employee", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := mocks.NewMockEmployeeStorage(ctrl)

	s := server.NewServer(storageMock)

	// create a new recorder to capture the response
	rr := httptest.NewRecorder()

	storageMock.EXPECT().SaveEmployee(gomock.Any()).Return(errors.New("error saving employee"))

	// call the router's ServeHTTP method with the request and recorder
	s.Router.ServeHTTP(rr, req)

	// check that the response status code is 500 OK
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}
