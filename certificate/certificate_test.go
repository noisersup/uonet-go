package certificate

import (
	"fmt"
	"testing"
)

func TestGetCerts(t *testing.T) {
	certs, err := GetCerts("3S16HNB", "warszawawola", "429895")
	if err != nil {
		t.Errorf(err.Error())
	}
	if certs.Key == "" || certs.Pfx == "" {
		t.Errorf("Certs are empty")
	}
}
func TestGetRestAddress(t *testing.T) {
	testToken := "3S1Gds3"
	expectedUrl := "https://lekcjaplus.vulcan.net.pl"

	url, err := GetRestAddress(testToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if url == "" {
		t.Errorf("The url is empty")
	}

	if url != expectedUrl {
		fmt.Println(expectedUrl)       //idk why is this executed
		t.Errorf("The url is: " + url) //TODO: Fix it
	}
}
