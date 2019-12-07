package wows

import (
	"encoding/json"
	"github.com/imroc/req"
	"log"
)

type Realm string

const (
	ASIA Realm = "https://api.worldofwarships.asia/wows/"
	NA   Realm = "https://api.worldofwarships.com/wows/"
	EU   Realm = "https://api.worldofwarships.eu/wows/"
	RU   Realm = "https://api.worldofwarships.ru/wows/"
)

type Client struct {
	applicationId string
	apiBase       string
}

func NewClient(applicationId string, realm Realm) *Client {
	return &Client{
		apiBase:       string(realm),
		applicationId: applicationId,
	}
}

func (c *Client) httpGet(api string, params req.Param) *req.Resp {
	params["application_id"] = c.applicationId
	resp, err := req.Get(c.apiBase+api, params)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func (c *Client) mustUnmarshal(data []byte, target interface{}) {
	err := json.Unmarshal(data, target)
	if err != nil {
		log.Fatal(err)
	}
}
