package impl

import (
	"fmt"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/restclientlib"
	"log"
)

type IDGenClientImpl struct {
	Host string
	Port uint
}

func (client *IDGenClientImpl) HostPort() string {
	return fmt.Sprintf("https://%s:%d", client.Host, client.Port)
}

func (client *IDGenClientImpl) NextID(tableName string) (*idgenmodel.IDGenResp,
	errorlib.AppError) {
	restClient := restclientlib.NewRestClient()
	hostPort := client.HostPort()
	url := hostPort + nextIDUrl
	log.Println("url =", url)
	req := idgenmodel.IDGenReq{
		TableName: &tableName,
	}
	var resp idgenmodel.IDGenResp
	appErr := restClient.Post(url, req, &resp)
	return &resp, appErr
}
