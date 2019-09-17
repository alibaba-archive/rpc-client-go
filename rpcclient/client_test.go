package rpcclient

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
)

func Test_InitWithConfig(t *testing.T) {
	config := map[string]string{
		"accessKeyId": "accessKeyId",
	}
	client := new(Client)
	originCredential := credential
	err := client.InitWithConfig(config)
	defer func() {
		credential = originCredential
	}()
	assert.Equal(t, "AccessKeySecret cannot be empty", err.Error())

	config["accessKeySecret"] = "accessKeySecret"
	err = client.InitWithConfig(config)
	assert.Nil(t, err)

	accessKeyId := client.GetAccessKeyId()
	assert.Equal(t, "accessKeyId", accessKeyId)

	accessKeySecret := client.GetAccessKeySecret()
	assert.Equal(t, "accessKeySecret", accessKeySecret)
}

func Test_GetUserAgent(t *testing.T) {
	client := new(Client)
	useragent := client.GetUserAgent("")
	assert.Equal(t, "AlibabaCloud (darwin; amd64) Golang/1.11.4 Core/0.01", useragent)

	useragent = client.GetUserAgent("test")
	assert.Equal(t, "AlibabaCloud (darwin; amd64) Golang/1.11.4 Core/0.01 test", useragent)
}

func Test_GetSignature(t *testing.T) {
	req := tea.NewRequest()
	req.Query["test"] = "ok"

	client := new(Client)
	sign := client.GetSignature(req, "accessKeySecret")
	assert.Equal(t, "jHx/oHoHNrbVfhncHEvPdHXZwHU=", sign)
}

func Test_Validator(t *testing.T) {
	client := new(Client)
	err := client.Validator(nil)
	assert.Nil(t, err)
}

func Test_DefaultNumber(t *testing.T) {
	client := new(Client)
	num := client.DefaultNumber(nil, 1)
	assert.Equal(t, 1, num)

	num = client.DefaultNumber(2, 1)
	assert.Equal(t, 2, num)
}

func Test_Query(t *testing.T) {
	client := new(Client)
	filter := map[string]interface{}{
		"client": "test",
	}

	result := client.Query(filter)
	assert.Equal(t, "test", result["client"])
}

func Test_GetTimestamp(t *testing.T) {
	client := new(Client)

	stamp := client.GetTimestamp()
	assert.NotNil(t, stamp)
}

func Test_GetNonce(t *testing.T) {
	client := new(Client)

	nonce := client.GetNonce()
	assert.Equal(t, 32, len(nonce))
}

func Test_Json(t *testing.T) {
	client := new(Client)
	httpresp := &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"cleint":"test"}`)),
	}
	resp := tea.NewResponse(httpresp)
	result, err := client.Json(resp)
	assert.Nil(t, err)
	assert.Equal(t, "test", result["cleint"])
}

func Test_GetEndpoint(t *testing.T) {
	client := new(Client)
	client.Endpoint = "client.aliyuncs.com"

	endpoint := client.GetEndpoint("", "")
	assert.Equal(t, "client.aliyuncs.com", endpoint)
}

func Test_HasError(t *testing.T) {
	client := new(Client)
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
	client := new(Client)
	str := client.Default(nil, "client")
	assert.Equal(t, "client", str)

	str = client.Default("", "client")
	assert.Equal(t, "", str)
}
