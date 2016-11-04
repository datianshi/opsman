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
const UPLOAD_PRODUCT_FORM_PARAM string = "product[file]"

const UPLOAD_STEMCELL_PATH string = "/api/v0/stemcells"
const UPLOAD_STEMCELL_FORM_PARAM string = "stemcell[file]"

//type Upload func(file *os.File) *rest.ConnectParams

type UploadMethod func(string, *os.File) *rest.ConnectParams

func (p *OpsMan) upload(file *os.File, method UploadMethod) error{
	token, err := p.tokenIssuer.GetToken()
	if(err!=nil){
		return err
	}
	restCon:= method(token, file)
	resp,err:= restCon.Connect()
	if(err != nil){
		return err
	}
	if (resp.Status != "200 OK"){
		return errors.New(fmt.Sprintf("No 200 status code received. Response code is: %s", resp.Status))
	}
	return nil
}

func (p *OpsMan) UploadProduct(file *os.File) error{
	var product UploadMethod = func(token string, file *os.File) *rest.ConnectParams{
		return p.uploadBuilder(UPLOAD_PRODUCT_PATH, UPLOAD_PRODUCT_FORM_PARAM, token, file)
	}
	return p.upload(file, product)
}

func (p *OpsMan) UploadStemcell(file *os.File) error{
	var stemcell UploadMethod = func(token string, file *os.File) *rest.ConnectParams{
		return p.uploadBuilder(UPLOAD_STEMCELL_PATH, UPLOAD_STEMCELL_FORM_PARAM, token, file)
	}
	return p.upload(file, stemcell)
}

func (p *OpsMan) uploadBuilder(uploadUrl, uploadForm, token string, file *os.File) *rest.ConnectParams{
	r := &rest.Rest{
		URL: fmt.Sprintf("%s%s", p.OpsManUrl, uploadUrl),
	}
	return r.Build().
		WithHttpMethod(rest.POST).
		SkipSslVerify(p.SkipSsl).
		WithHttpHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		WithMultipartForm(uploadForm, file)
}