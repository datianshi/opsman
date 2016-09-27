package opsmantest

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"

	. "github.com/onsi/gomega"
	"io/ioutil"
)

type FakeTokenIssuer struct {
	Token        string
	ErrorControl error
}

func (i *FakeTokenIssuer) GetToken() (token string, err error) {
	token = i.Token
	err = i.ErrorControl
	return
}

func CompareMd5(buffer *bytes.Buffer, b_array *[]byte) bool {
	slice1 := md5.Sum(*b_array)
	slice2 := md5.Sum(buffer.Bytes())
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func Md5FileSum(file *os.File) (md5str string, err error) {
	b := &bytes.Buffer{}
	_, err = io.Copy(b, file)
	if err != nil {
		return
	}
	fmt.Println(len(b.Bytes()))
	md5str = Md5Bytes(b.Bytes())
	return
}

func Md5Bytes(b []byte) (md5str string) {
	md5str = fmt.Sprintf("%x", md5.Sum(b))
	return
}

func VerifyUploadFile(file *os.File) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		md5str, err := Md5FileSum(file)
		Ω(err).Should(BeNil())
		if err == nil {
			defer req.Body.Close()
			b := &bytes.Buffer{}
			fmt.Println(b.String())
			io.Copy(b, req.Body)
			fmt.Println(md5str)
			Ω(md5str).Should(Equal(Md5Bytes(b.Bytes())), "Upload file md5sum mismatch")
		}
	}
}

func CreateGarbageFile(content string) (file *os.File, err error) {
	file, err = ioutil.TempFile("", "garbage")
	if err != nil {
		return
	}
	io.Copy(file, bytes.NewBufferString(content))
	file.Close()
	fmt.Println(file.Name())
	return os.Open(file.Name())
}
