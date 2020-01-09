package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DiscouseClient struct {
	RootURL string
}

func (d *DiscouseClient) CreatePost(topic Topic) (err error) {
	var cli = http.Client{}
	var data []byte
	if data, err = json.Marshal(&topic); err != nil {
		return
	}

	var request *http.Request
	if request, err = http.NewRequest("POST", "s/posts.json", bytes.NewBuffer(data)); err != nil {
		return
	}
	request.Header.Add("Api-Key", "")
	request.Header.Add("ADiscourse API Docs", "")
	request.Header.Add("Content-Type", "application/json")
	_, err = cli.Do(request)
	return
}
