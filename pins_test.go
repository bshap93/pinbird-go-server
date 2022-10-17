package main

import (
	"regexp"
	"testing"
)

func TestGetAuthAppendage(t *testing.T) {
	token := "bshap93:1A2DC6C239D948A40F2C"
	authAppendage := getAuthAppendage(token)
	want := regexp.MustCompile(`\?auth_token=[a-z0-9]+:[A-Z0-9]+&format=json`)
	if !want.MatchString(authAppendage) {
		t.Fatalf(`getAuthAppendage(token) = %q, want match for %#q, nil`, authAppendage, want)
	}
}

func TestEmptyGetAuthAppendage(t *testing.T) {
	authAppendage := getAuthAppendage("")
	want := regexp.MustCompile(`\?auth_token=[a-z0-9]+:[A-Z0-9]+&format=json`)
	if want.MatchString(authAppendage) {
		t.Fatalf(`getAuthAppendage(token) = %q, want match for %#q, nil`, authAppendage, want)
	}
}

func TestGetFullAppendage(t *testing.T) {
	params := make(map[string]string)
	params["url"] = "https://pub.dev"
	params["tags"] = "flutter development"
	params["description"] = "A repository of Dart packages"
	params["extended"] = ""
	params["shared"] = "yes"
	params["toread"] = "no"
	token := "bshap93:1A2DC6C239D948A40F2C"
	fullAppendage := getFullAppendage(params, token)
	//TODO
	want := regexp.MustCompile(`regex`)
	if !want.MatchString(fullAppendage) {
		t.Fatalf(`getAuthAppendage(token) = %q, want match for %#q, nil`, fullAppendage, want)
	}

}

// // TestHelloEmpty calls greetings.Hello with an empty string,
// // checking for an error.
// func TestHelloEmpty(t *testing.T) {
// 	msg, err := Hello("")
// 	if msg != "" || err == nil {
// 		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
// 	}

// }
