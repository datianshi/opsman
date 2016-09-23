package pivnet

import (
	"io"
	"github.com/datianshi/rest-func/rest"
	"fmt"
	"strconv"
	"time"
	"gopkg.in/cheggaaa/pb.v1"
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
	sizeStr := resp.Header.Get("Content-Length")
	var size int64
	if size, err = strconv.ParseInt(sizeStr, 10, 64); err != nil {
		return
	}
	bar := pb.New64(size).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.Start()
	reader:= bar.NewProxyReader(resp.Body)
	_, err = io.Copy(dest, reader)
	return
}


