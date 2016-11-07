package pivnet

import (
	"fmt"
	"github.com/datianshi/rest-func/rest"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"strconv"
	"time"
	"encoding/json"
	"bytes"
)

type Pivnet struct {
	Token string
	PivURL string
}

type Product struct{
	Id int64
	Verstion string
	AcceptUrl string
	Files []ProductFile
}

type ProductFile struct{
	Name string
	DownloadUrl string
}

type LatestResponse struct{
	Id int64
	Version string
	Eula Eula
	Product_files []ProductFileResponse
	Links Links `json:"_links"`
}

type ProductFileResponse struct{
	Id int64
	Name string
	Links Links `json:"_links"`
}

type Eula struct{
	Links Links `json:"_links"`
}

type Links struct{
	Self URL
	Download URL
	EulaURL URL `json:"eula_acceptance"`
}

type URL struct{
	Href string
}

func (p *Pivnet) Download(dest io.Writer, url string) (err error) {
	r := &rest.Rest{
		URL: url,
	}
	resp, err := r.Build().WithHttpMethod(rest.POST).WithHttpHeader("Authorization", fmt.Sprintf("Token %s", p.Token)).Connect()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	sizeStr := resp.Header.Get("Content-Length")
	var size int64
	if size, err = strconv.ParseInt(sizeStr, 10, 64); err != nil {
		return
	}
	bar := pb.New64(size).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.Start()
	reader := bar.NewProxyReader(resp.Body)
	_, err = io.Copy(dest, reader)
	bar.Finish()
	return
}

func (p *Pivnet) LatestProduct(productName string) (product Product, err error) {
	r1 := &rest.Rest{
		URL: fmt.Sprintf("%s/api/v2/products/%s/releases/latest", p.PivURL, productName),
	}

	resp, err := r1.Build().WithHttpHeader("Authorization", fmt.Sprintf("Token %s", p.Token)).Connect()
	if err!=nil{
		return
	}
	defer resp.Body.Close()
	var productResponse LatestResponse
	var b bytes.Buffer
	io.Copy(&b, resp.Body)
	err = json.Unmarshal(b.Bytes(), &productResponse)
	if (err!=nil){
		return
	}
	files := make([]ProductFile, 0)
	for _, r := range productResponse.Product_files{
		file:= ProductFile{
			Name: r.Name,
			DownloadUrl: r.Links.Download.Href,
		}
		files = append(files, file)
	}
	product = Product{
		Id: productResponse.Id,
		AcceptUrl: productResponse.Links.EulaURL.Href,
		Files: files,
	}
	return
}
