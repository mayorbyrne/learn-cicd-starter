// let's write some tests for the auth package

package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "")

	auth, err := GetAPIKey(header)

	// we should have an ErrNoAuthHeaderIncluded error
	if auth != "" {
		t.Fatalf("expected: empty string, got: %s", auth)
	}
	// and the error should be ErrNoAuthHeaderIncluded
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	// check if the error is ErrNoAuthHeaderIncluded
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}

	// let's check if the error is equal to ErrNoAuthHeaderIncluded
	// we can use reflect.DeepEqual to compare the errors
	// this is a bit overkill, but it works
	// and we can use it to check if the error is nil
	if !reflect.DeepEqual(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}

	header = http.Header{}
	header.Set("Authorization", "oneWord")

	auth, err = GetAPIKey(header)

	// we should have an empty string and a malformed authorization header error
	if auth != "" {
		t.Fatalf("expected: empty string, got: %s", auth)
	}
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	// check if the error is a malformed authorization header error
	if err.Error() != "malformed authorization header" {
		t.Fatalf("expected: %s, got: %s", "malformed authorization header", err.Error())
	}
	// we can also check if the error is equal to the error we expect
	if !reflect.DeepEqual(err, errors.New("malformed authorization header")) {
		t.Fatalf("expected: %v, got: %v", errors.New("malformed authorization header"), err)
	}

	header = http.Header{}
	header.Set("Authorization", "ApiKey HelloWorld")

	auth, err = GetAPIKey(header)
	// we should have the API key and no error
	if auth != "HelloWorld" {
		t.Fatalf("expected: HelloWorld, got: %s", auth)
	}

	if err != nil {
		t.Fatalf("expected: nil, got: %v", err)
	}
}
