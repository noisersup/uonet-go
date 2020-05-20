package certificate

/*
	package for getting certifactes from apiClient
*/
import (
	api "../apiClient"
	u "../utils"
	"encoding/json"
	"strings"
)

type Certs struct {
	Key string
	Pfx string
}

func GetCerts(token string, symbol string, pin string) Certs { //returns the object
	jsonData := api.TakeCerts(token, symbol, pin)
	var certs Certs
	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		u.ErrLog(err)
	}
	tokenCert := result["TokenCert"].(map[string]interface{})

	certs.Key = tokenCert["CertyfikatKlucz"].(string)
	certs.Pfx = tokenCert["CertyfikatPfx"].(string)
	return certs
}

func SetRestAddress(token string) string {
	data := api.TakeRoutingRules()
	table := strings.Split(data, "\n")
	chars := token[0:2]

	for _, row := range table {
		rowChars := row[0:2]
		if rowChars == chars {
			urls := strings.Split(row, ",")
			url := urls[1]
			return url
		}
	}
	return ""
}
