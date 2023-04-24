package helper

import (
	"crypto"
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"process-loan/lib/constant"
	"process-loan/lib/log"
	"reflect"
	"runtime"
	"time"
)

func NilInterface(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}

func ConvertStringToDate(s string, layoutISO string) (time.Time, error) {
	//layoutISO := "2006-01-02T15:04:05"
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		fmt.Println("Error Convert Date => ", err)
		return t, err
	}
	return t, nil
}

func ValidateSignatureRsa(publicKey *rsa.PublicKey, data, signature string) error {
	PrintHeader()

	sDec, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		log.LogData("", "validate", "base64.StdEncoding.DecodeStrin", constant.LEVEL_LOG_ERROR, err)
		return err
	}

	h := sha256.New()
	_, err = h.Write([]byte(data))
	if err != nil {
		log.LogData("", "validate", "base64.StdEncoding.DecodeStrin", constant.LEVEL_LOG_ERROR, err)
		return err
	}
	d := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, d, sDec)
	if err != nil {
		log.LogData("", "validate", "base64.StdEncoding.DecodeStrin", constant.LEVEL_LOG_ERROR, err)
		return err
	}

	return nil
}

func StringToSignHMAC512(method, url, timeStamp string, minifyReq []byte) string {
	PrintHeader()
	h := sha256.New()
	h.Write(minifyReq)

	fmt.Println("Data Minify : ", string(minifyReq))

	data := method + "&" + url + "&" + string(minifyReq) + "&" + timeStamp
	return data
}

func ValidateSignatureHMAC512(key string, data []byte, signature string) bool {
	PrintHeader()

	verify := CreateSignatureHMAC(key, data)
	fmt.Println("Signature : ", signature)
	fmt.Println("Verify : ", verify)
	if signature == verify {
		return true
	}
	return false
	// bypass validation signature
	// return true
}

func CreateSignatureHMAC(sKey string, data []byte) string {
	PrintHeader()

	h := hmac.New(sha512.New, []byte(sKey))

	h.Write(data)
	verify := hex.EncodeToString(h.Sum(nil))
	// verify := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return verify
}

func PrintHeader() {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Printf("<======> %s <======>", runtime.FuncForPC(pc).Name())
	fmt.Println()
}
