// file: main_test.go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetBatting(t *testing.T) {
	req, err := http.NewRequest("GET", "/batting/", nil)
	if err != nil {
		t.Fatal(err)
	}

	MockServer := &StatServer{
		store: NewMockStatStore(),
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(MockServer.GetBattingHandler)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"Player":"Player1","AVG":".200"},{"Player":"Player2", "AVG":".250"},{"Player":"Player3","AVG":".300"}]`
	if !IsJsonEqual(response.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", response.Body.String(), expected)
	}
}

func TestGetPitching(t *testing.T) {
	req, err := http.NewRequest("GET", "/pitching/", nil)
	if err != nil {
		t.Fatal(err)
	}

	MockServer := &StatServer{
		store: NewMockStatStore(),
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(MockServer.GetPitchingHandler)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Player":"Player1","ERA":"3.00"},{"Player":"Player2", "ERA":"2.50"},{"Player":"Player3","ERA":"2.00"}`
	if !IsJsonEqual(response.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", response.Body.String(), expected)
	}
}

func TestGetFielding(t *testing.T) {
	req, err := http.NewRequest("GET", "/fielding/", nil)
	if err != nil {
		t.Fatal(err)
	}

	MockServer := &StatServer{
		store: NewMockStatStore(),
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(MockServer.GetFieldingHandler)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Player":"Player1","FLDP":".950"},{"Player":"Player2","FLDP":".960"},{"Player":"Player3","FLDP":".970"}`
	if !IsJsonEqual(response.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", response.Body.String(), expected)
	}
}

func IsJsonEqual(obj1, obj2 string) bool {
	var o1, o2 map[string]any
	json.Unmarshal([]byte(obj1), &o1)
	json.Unmarshal([]byte(obj2), &o2)
	return reflect.DeepEqual(o1, o2)
}
