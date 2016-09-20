package opsman

import (

	"fmt"
	"github.com/datianshi/opsman/uaa"
	"github.com/datianshi/rest-func/rest"
	"os"
	"errors"
)

type OpsMan struct{
	OpsManUrl string
	Username string
	Password string
	SkipSsl  bool
}

func (p *OpsMan) Upload(file *os.File) error{
	uaa:= &uaa.UAA{
		URL: fmt.Sprintf("%s/uaa", p.OpsManUrl),
		Username: p.Username,
		Password: p.Password,
		SkipSsl: p.SkipSsl,
	}
	token, err :=uaa.GetToken()
	if(err!=nil){
		return err
	}
	r := &rest.Rest{
		URL: fmt.Sprintf("%s/api/v0/available_products", p.OpsManUrl),
	}
	resp, err:=r.Build().
		WithHttpMethod(rest.POST).
		SkipSslVerify(p.SkipSsl).
		WithHttpHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		WithMultipartForm("product[file]", file).
		Connect()
	if(err != nil){
		return err
	}
	if (resp.Status != "200 OK"){
		return errors.New(fmt.Sprintf("No 200 status code received. Response code is: %s", resp.Status))
	}
	return nil
}