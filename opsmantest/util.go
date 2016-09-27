package opsmantest

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"

	"io/ioutil"

	. "github.com/onsi/gomega"
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

func Md5FileSum(path string) (md5str string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	fmt.Println(len(b))
	md5str = Md5Bytes(b)
	return
}

func Md5Bytes(b []byte) (md5str string) {
	md5str = fmt.Sprintf("%x", md5.Sum(b))
	return
}

func VerifyUploadFile(path string, formValue string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		md5str, err := Md5FileSum(path)
		Ω(err).Should(BeNil())
		if err == nil {
			defer req.Body.Close()
			read_form , _ := req.MultipartReader()
			for {
				part, err_part := read_form.NextPart()
				if err_part == io.EOF {
					break
				}
				if part.FormName() == formValue {
					buf := new(bytes.Buffer)
					buf.ReadFrom(part)
					Ω(md5str).Should(Equal(Md5Bytes(buf.Bytes())), "Upload file md5sum mismatch")
					return
				}
			}
			Ω(true).Should(Equal(false), fmt.Sprintf("There is not form part: %s", formValue))
		}
	}
}

func CreateGarbageFile(content string) (path string, err error) {
	file, err := ioutil.TempFile("", "garbage")
	if err != nil {
		return
	}
	io.Copy(file, bytes.NewBufferString(content))
	path = file.Name()
	return
}
