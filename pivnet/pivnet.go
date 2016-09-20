package pivnet

import (
	"io"
	"github.com/datianshi/rest-func/rest"
	"fmt"
)

type Pivnet struct {
	URL   string
	Token string
}

func (p *Pivnet) Download(dest io.Writer) (err error) {
	r := &rest.Rest{
		URL: p.URL,
	}
	resp, err := r.Build().WithHttpMethod(rest.POST).WithHttpHeader("Authorization", fmt.Sprintf("Token %s", p.Token)).Connect()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(dest, resp.Body)
	return
}


