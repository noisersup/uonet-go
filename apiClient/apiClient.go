package apiClient

/*
	Package responsible for sending requests.
	It's the lowest layer of this app
*/
import (
	cert "../certificate"
	u "../utils"
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var client = http.Client{}

func TakeRoutingRules() string {
	resp, err := http.Get("http://komponenty.vulcan.net.pl/UonetPlusMobile/RoutingRules.txt")
	if err != nil {
		u.ErrLog(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		u.ErrLog(err)
	}
	return string(body)
}

func TakeCerts(token string, symbol string, pin string) (string, error) {
	errPrefix := "apiClient.TakeCerts error: "

	routingAddress, err := cert.GetRestAddress(token)
	if err != nil {
		u.ErrLog(err)
	}

	address, err := url.Parse(routingAddress.String() + "/" + symbol + "/mobile-api/Uczen.v3.UczenStart/Certyfikat") //Doesn't work
	if err != nil {
		u.ErrLog(err)
	}

	timeNow := time.Now().Unix()
	time1 := strconv.Itoa(int(timeNow))
	time2 := strconv.Itoa(int(timeNow) - 1)
	//TODO: make this code cleaner, add random uuid, change device name

	jsonstr := []byte(`{
		"PIN": "` + pin + `",
		"TokenKey": "` + token + `",
		"AppVersion": "18.4.1.388",
		"DeviceId": "a4f98332-6a5d-4a53-bd40-a6dd559a9fae",
		"DeviceName": "Galaxy#S8",
		"DeviceNameUser": "",
			"DeviceDescription": "",
		"DeviceSystemType": "Android",
		"DeviceSystemVersion": "6.0.1",
		"RemoteMobileTimeKey": ` + time1 + `,
		"TimeKey": ` + time2 + `,
		"RequestId": "a4f98332-6a5d-4a53-bd40-a6dd559a9fae",
		"RemoteMobileAppVersion": "18.4.1.388",
		"RemoteMobileAppName": "VULCAN-Android-ModulUcznia"
	}`)

	req, err := http.NewRequest("POST", address.String(), bytes.NewBuffer(jsonstr))
	if err != nil {
		return "", errors.New(errPrefix + err.Error())
	}

	req.Header.Set("RequestMobileType", "RegisterDevice")
	req.Header.Set("User-Agent", "MobileUserAgent")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(errPrefix + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(errPrefix + err.Error())
	}

	return string(body), nil
}
