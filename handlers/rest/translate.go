package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"roblesdotdev/hello-api/translation"
	"strings"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	fmt.Println("Language: ", language)
	fmt.Println("Word: ", word)
	translation := translation.Translate(word, language)
	if translation == "" {
		language = ""
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := Resp{
		Language:    language,
		Translation: translation,
	}
	if err := enc.Encode(resp); err != nil {
		panic("unable to code response")
	}
}
