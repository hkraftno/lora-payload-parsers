package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidHTTPMethod(t *testing.T) {
	r, _ := http.NewRequest("POST", "/03fd8e019c10001b", nil)
	w := httptest.NewRecorder()

	rootHandler(w, r)
	result := w.Result()
	if result.StatusCode == 200 {
		t.Errorf("Expected status code to be != 200 when not using GET")
	}
	body, err := ioutil.ReadAll(result.Body)
	if len(body) == 0 || err != nil {
		t.Errorf("Expected body to contain error was %+q %s", body, err)
	}
}

func TestInvalidHex(t *testing.T) {
	r, _ := http.NewRequest("GET", "/03fd8e019c10001b6", nil)
	w := httptest.NewRecorder()

	rootHandler(w, r)
	result := w.Result()
	if result.StatusCode == 200 {
		t.Errorf("Expected status code to be != 200 when not using valid HEX")
	}
	body, err := ioutil.ReadAll(result.Body)

	if len(body) == 0 || err != nil {
		t.Errorf("Expected body to contain error was %+q %s", body, err)
	}
}

func TestInvalidURL(t *testing.T) {
	r, _ := http.NewRequest("GET", "not-valid", nil)
	w := httptest.NewRecorder()

	rootHandler(w, r)
	result := w.Result()
	if result.StatusCode == 200 {
		t.Errorf("Expected status code to be != 200 when not using valid URL")
	}
	body, err := ioutil.ReadAll(result.Body)

	if len(body) == 0 || err != nil {
		t.Errorf("Expected body to contain error was %+q %s", body, err)
	}
}

func TestInvalidMessageFormat(t *testing.T) {
	r, _ := http.NewRequest("GET", "/09", nil)
	w := httptest.NewRecorder()

	rootHandler(w, r)
	result := w.Result()
	if result.StatusCode == 200 {
		t.Errorf("Expected status code to be != 200 when not using valid URL")
	}
	body, err := ioutil.ReadAll(result.Body)

	if len(body) == 0 || err != nil {
		t.Errorf("Expected body to contain error was %+q %s", body, err)
	}
}

func TestParserExampleHex1(t *testing.T) {
	r, _ := http.NewRequest("GET", "/pir_lab_xxns/02fc8e019c10001b1b00", nil)
	w := httptest.NewRecorder()

	rootHandler(w, r)
	result := w.Result()
	body, _ := ioutil.ReadAll(result.Body)
	var expected = `{"id":2,"battery_level":99,"internal_data":"8e019c10","counter":1776384}`
	if expected != string(body) {
		t.Errorf(
			"Expected JSON to be\n%s\nbut was\n%s",
			expected,
			body,
		)
	}
}
