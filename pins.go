package main

import (
	"strings"
)

const BaseUrl = "https://api.pinboard.in/v1"

var db = make(map[string]string)

type pinboardPin struct {
	href        string
	description string
	extended    string
	meta        string
	hash        string
	time        string
	shared      string
	toread      string
	tags        string
}

func getAuthAppendage(apiTok string) string {
	return "?auth_token=" + apiTok + "&format=json"
}

var params map[string]string

func getFullAppendage(params map[string]string, apiTok string) string {
	var appendageStringBuilder strings.Builder
	appendageStringBuilder.WriteByte('?')
	appendageStringBuilder.WriteString("auth_token=")
	appendageStringBuilder.WriteString(apiTok)

	if len(params) > 0 {
		appendageStringBuilder.WriteByte('&')
		for key, element := range params {
			if len(params) == 1 {
				appendageStringBuilder.WriteString(key)
				appendageStringBuilder.WriteByte('=')
				appendageStringBuilder.WriteString(element)
			}
		}
		fullAppendage := appendageStringBuilder.String()
		if len(fullAppendage) > 0 {
			fullAppendage = fullAppendage[:len(fullAppendage)-1]
		}
		return fullAppendage
	} else {
		return appendageStringBuilder.String()
	}

}

// func getAllPins() {
// 	response, err := http.Get()
// }
