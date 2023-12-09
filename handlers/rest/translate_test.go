package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"roblesdotdev/hello-api/handlers/rest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {
	// Arrange
	tt := []struct {
		Endpoint             string
		StatusCode           int
		ExpectedLanguage     string
		ExpectedTranslataion string
	}{
		{
			Endpoint:             "/hello",
			StatusCode:           http.StatusOK,
			ExpectedLanguage:     "english",
			ExpectedTranslataion: "hello",
		},
		{
			Endpoint:             "/hello?language=german",
			StatusCode:           http.StatusOK,
			ExpectedLanguage:     "german",
			ExpectedTranslataion: "hallo",
		},
		{
			Endpoint:             "/hello?language=dutch",
			StatusCode:           http.StatusNotFound,
			ExpectedLanguage:     "",
			ExpectedTranslataion: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	// Act
	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)

		// Assert
		if rr.Code != test.StatusCode {
			t.Errorf(`Expected status %d but received %d`, test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`Expected language %s but received "%s"`, test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslataion {
			t.Errorf(`Expected translation %s but received %s`, test.ExpectedTranslataion, resp.Translation)
		}
	}
}
