package uaa

import (
	"github.com/datianshi/rest-func/rest"
	"fmt"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

const client_id string = "opsman"

type UAA struct {
	URL      string
	Username string
	Password string
	SkipSsl  bool
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (u *UAA) GetToken() (token string, err error) {
	r := &rest.Rest{
		URL: fmt.Sprintf("%s/oauth/token", u.URL),
	}
	formValues := url.Values{"grant_type": {"password"}, "username": {u.Username}, "password": {u.Password}}
	response, err := r.Build().
		WithBasicAuth(client_id,"").
		WithFormValue(formValues).
		SkipSslVerify(u.SkipSsl).
		WithHttpMethod(rest.POST).
		Connect()
	if (err != nil) {
		return
	}
	body := response.Body
	defer body.Close()
	tokenResp := &TokenResponse{}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &tokenResp)
	if err != nil {
		return
	}
	token = tokenResp.AccessToken
	return

}