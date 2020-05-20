package uonetApi

/*
	The "endpoint" of this api
*/
import (
	u "../utils"
	"encoding/base64"
	"fmt"
	"software.sslmate.com/src/go-pkcs12"
)

func Sign(content string, pfx string) {
	cert, err := base64.StdEncoding.DecodeString(pfx)
	if err != nil {
		u.ErrLog(err)
	}

	_, certificate, err := pkcs12.Decode(cert, "CE75EA598C7743AD9B0B7328DED85B06")
	fmt.Println(certificate, err)
	//TODO: Finish him!
}
func GetApi() {

}
