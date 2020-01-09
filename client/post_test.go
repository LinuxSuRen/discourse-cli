package client_test

import (
	"github.com/linuxsuren/discourse-cli/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTopic(t *testing.T)  {
	dis := &client.DiscouseClient{
		RootURL: "https://community.jenkins-zh.cn/",
	}

	topic := client.Topic{}
	err := dis.CreatePost(topic)
	assert.Nil(t, err)
}
