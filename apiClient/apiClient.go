package apiClient
/*
	Package responsible for sending requests.
	It's the lowest layer of this app
*/
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	u "../utils"
	"strconv"
	"time"
)
var client = http.Client{}

func TakeRoutingRules() string {
	resp, err := http.Get("http://komponenty.vulcan.net.pl/UonetPlusMobile/RoutingRules.txt")
	if err != nil {u.ErrLog(err)}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {u.ErrLog(err)}

	return string(body)
}

func TakeCerts(token string, symbol string, pin string) string{
	//address := cert.GetRestAddress(token) + "/" + symbol + "/mobile-api/Uczen.v3.UczenStart/Certyfikat" //Doesn't work
	address := "" //Hardcoded for now, TODO: repair GetRestAddress()
	fmt.Println(address)

	timeNow := time.Now().Unix()
	time1 := strconv.Itoa(int(timeNow))
	time2 := strconv.Itoa(int(timeNow) - 1)
	//TODO: make this code cleaner, add random uuid, change device name

	jsonstr := []byte(`{
		"PIN": "` + pin + `",
		"TokenKey": "` + token + `",
		"AppVersion": "18.4.1.388",
		"DeviceId": "a4f98332-6a5d-4a53-bd40-a6dd559a9fae",
		"DeviceName": "Galaxy#S7",
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

	req, err := http.NewRequest("POST", address, bytes.NewBuffer(jsonstr))
	req.Header.Set("RequestMobileType", "RegisterDevice")
	req.Header.Set("User-Agent", "MobileUserAgent")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {u.ErrLog(err)}

	resp, err := client.Do(req)
	if err != nil {u.ErrLog(err)}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}