package service

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
)

func Test_InitWithConfig(t *testing.T) {
	config := map[string]interface{}{
		"accessKeyId": "accessKeyId",
	}
	client := new(BaseClient)
	err := client.InitClient(config)

	assert.Equal(t, "AccessKeySecret cannot be empty", err.Error())

	config["accessKeySecret"] = "accessKeySecret"
	err = client.InitClient(config)
	assert.Nil(t, err)

	accessKeyId := client.GetAccessKeyId()
	assert.Equal(t, "accessKeyId", accessKeyId)

	accessKeySecret := client.GetAccessKeySecret()
	assert.Equal(t, "accessKeySecret", accessKeySecret)
}

func Test_UserAgent(t *testing.T) {
	client := new(BaseClient)
	assert.Contains(t, client.GetUserAgent(""), "AlibabaCloud")
}

func Test_GetSignature(t *testing.T) {
	req := tea.NewRequest()
	req.Query["test"] = "ok"

	client := new(BaseClient)
	sign := client.GetSignature(req, "accessKeySecret")
	assert.Equal(t, "jHx/oHoHNrbVfhncHEvPdHXZwHU=", sign)
}

func Test_DefaultNumber(t *testing.T) {
	client := new(BaseClient)
	num := client.DefaultNumber(nil, 1)
	assert.Equal(t, 1, num)

	num = client.DefaultNumber(2, 1)
	assert.Equal(t, 2, num)
}

func Test_Query(t *testing.T) {
	client := new(BaseClient)
	filter := map[string]interface{}{
		"client": "test",
	}

	result := client.Query(filter)
	assert.Equal(t, "test", result["client"])
}

func Test_GetTimestamp(t *testing.T) {
	client := new(BaseClient)

	stamp := client.GetTimestamp()
	assert.NotNil(t, stamp)
}

func Test_GetNonce(t *testing.T) {
	client := new(BaseClient)

	nonce := client.GetNonce()
	assert.Equal(t, 32, len(nonce))
}

func Test_Json(t *testing.T) {
	client := new(BaseClient)
	httpresp := &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"cleint":"test"}`)),
	}
	resp := tea.NewResponse(httpresp)
	result, err := client.Json(resp)
	assert.Nil(t, err)
	assert.Equal(t, "test", result["cleint"])
}

func Test_GetEndpoint(t *testing.T) {
	client := new(BaseClient)
	client.Endpoint = "client.aliyuncs.com"

	endpoint := client.GetEndpoint("", "")
	assert.Equal(t, "client.aliyuncs.com", endpoint)
}

func Test_HasError(t *testing.T) {
	client := new(BaseClient)
	iserror := client.HasError(nil)
	assert.True(t, iserror)

	body := map[string]interface{}{
		"Code": "200",
	}
	iserror = client.HasError(body)
	assert.True(t, iserror)

	body = make(map[string]interface{})
	iserror = client.HasError(body)
	assert.False(t, iserror)
}

func Test_Default(t *testing.T) {
	client := new(BaseClient)
	str := client.Default(nil, "client")
	assert.Equal(t, "client", str)

	str = client.Default("", "client")
	assert.Equal(t, "", str)
}
