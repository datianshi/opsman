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
	SkipSsl  bool
	tokenIssuer uaa.TokenIssuer
}

func CreateOpsman(url string, skipssl bool, tokenIssuer uaa.TokenIssuer) *OpsMan{
	return &OpsMan{
		OpsManUrl: url,
		SkipSsl: skipssl,
		tokenIssuer: tokenIssuer,
	}
}

const UPLOAD_PRODUCT_PATH string = "/api/v0/available_products"

func (p *OpsMan) Upload(file *os.File) error{
	token, err := p.tokenIssuer.GetToken()
	if(err!=nil){
		return err
	}
	r := &rest.Rest{
		URL: fmt.Sprintf("%s%s", p.OpsManUrl, UPLOAD_PRODUCT_PATH),
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