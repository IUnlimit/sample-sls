package es

import (
	"fmt"
	global "github.com/IUnlimit/sample-sls/internal"
	"github.com/IUnlimit/sample-sls/internal/model"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
)

func Invoke() {
	config := global.Config.ES

	client := resty.New()
	info := &model.ESInfo{}
	client.R().SetResult(info).Get(config.Url)
	fmt.Println(spew.Sdump(info))
}
