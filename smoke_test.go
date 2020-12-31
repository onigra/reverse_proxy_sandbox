package reverse_proxy_sandbox_test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestWeb1(t *testing.T) {
	// Given
	url := "http://web1.nginx.internal"

	// And
	expected := "web1"

	// When
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	actual, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}

	// Then
	if resp.StatusCode != 200 {
		t.Error("a response code is not 200")
	}

	if string(actual) != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}

func TestWeb2(t *testing.T) {
	// Given
	url := "http://web2.nginx.internal"

	// And
	expected := "web2"

	// When
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	actual, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}

	// Then
	if resp.StatusCode != 200 {
		t.Error("a response code is not 200")
	}

	if string(actual) != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}

func TestWeb3(t *testing.T) {
	// Given
	url := "http://web3.nginx.internal"

	// And
	expected := "web3"

	// When
	resp, err := http.Get(url)
	if err != nil {
		t.Error(err)
	}

	actual, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}

	// Then
	if resp.StatusCode != 200 {
		t.Error("a response code is not 200")
	}

	if string(actual) != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}
