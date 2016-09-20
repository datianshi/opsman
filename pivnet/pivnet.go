package pivnet

import "io"

type Pivnet struct {
	URL   string
	Token string
	Dest  io.Writer
}

func (p *Pivnet) Download() (err error) {
	//r := &rest.Rest{
	//	URL: p.URL,
	//}
	//resp, err := r.Build().WithHttpMethod(rest.POST).WithHttpHeader("Authorization", fmt.Sprintf("Token %s", p.Token)).Connect()
	//if err != nil {
	//	return
	//}
	//_, err = io.Copy(p.Dest, resp)
	return
}


