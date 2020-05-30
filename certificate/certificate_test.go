package certificate

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetCerts(t *testing.T) {
	certs, err := GetCerts("", "", "")
	if err != nil {
		t.Errorf(err.Error())
	}
	if certs.Key == "" || certs.Pfx == "" {
		t.Errorf("Certs are empty")
	}
}
func TestGetRestAddress(t *testing.T) {
	testToken := "3S1Gds3"
	expectedU, err := url.Parse("https://lekcjaplus.vulcan.net.pl")
	if err != nil {
		t.Errorf(err.Error())
	}

	u, err := GetRestAddress(testToken)
	if err != nil {
		t.Errorf(err.Error())
	}

	if u == nil {
		t.Errorf("The url is empty")
	}

	if u.String() != expectedU.String() {
		fmt.Println(expectedU)                //idk why is this executed
		t.Errorf("The url is: " + u.String()) //TODO: Fix it
	}
}
