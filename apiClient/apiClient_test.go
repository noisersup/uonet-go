package apiClient

import (
	"encoding/json"
	"testing"
)

func TestTakeRoutingRules(t *testing.T) {
	var rules string = TakeRoutingRules()
	if rules == "" {
		t.Errorf("RoutingRules file is empty")
	}
}
func TestTakeCerts(t *testing.T) {
	certs, err := TakeCerts("", "", "")
	if err != nil {
		t.Errorf(err.Error())
	}

	if certs == "" {
		t.Errorf("Body of Certs response is empty")
	} else {
		var result map[string]interface{}

		err := json.Unmarshal([]byte(certs), &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		if result["IsError"].(bool) {
			t.Errorf("Server error message: " + result["Message"].(string))
		}
	}
}
