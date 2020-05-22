package certificate

/*
	package for getting certifactes from apiClient
*/
import (
	api "../apiClient"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Certs struct {
	Key string
	Pfx string
}

func GetCerts(token string, symbol string, pin string) (Certs, error) { //returns the object
	errPrefix := "certificate.GetCerts error: "

	jsonData, err := api.TakeCerts(token, symbol, pin)
	if err != nil {
		return Certs{}, errors.New(errPrefix + err.Error())
	}

	fmt.Println("\n\n", jsonData)
	var certs Certs
	var result map[string]interface{}

	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		return Certs{}, errors.New(errPrefix + err.Error())
	}

	if result["IsError"].(bool) {
		return certs, errors.New(
			errPrefix + "Remote server returns [" + result["Message"].(string) + "]")
	}

	tokenCert := result["TokenCert"].(map[string]interface{})

	certs.Key = tokenCert["CertyfikatKlucz"].(string)
	certs.Pfx = tokenCert["CertyfikatPfx"].(string)
	return certs, nil
}

func GetRestAddress(token string) (string, error) {
	errPrefix := "certificate.GetRestAddress error: "

	data := api.TakeRoutingRules()
	table := strings.Split(data, "\n")
	chars := token[0:2]

	for _, row := range table {
		rowChars := row[0:2]
		if rowChars == chars {
			urls := strings.Split(row, ",")
			url := urls[1]
			return url, nil
		}
	}
	return "", errors.New(errPrefix + "No URL found")
}
