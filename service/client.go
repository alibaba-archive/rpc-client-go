package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
	"github.com/aliyun/rpc-client-go/utils"
)

// client is for calling oss api
type BaseClient struct {
	RegionId             string `json:"RegionId" xml:"RegionId"`
	Protocol             string `json:"Protocol" xml:"Protocol"`
	Endpoint             string `json:"Endpoint" xml:"Endpoint"`
	UserAgent            string `json:"UserAgent" xml:"UserAgent"`
	SecurityToken        string `json:"SecurityToken" xml:"SecurityToken"`
	ReadTimeout          int    `json:"ReadTimeout" xml:"ReadTimeout"`
	ConnectTimeout       int    `json:"ConnectTimeout" xml:"ConnectTimeout"`
	HttpProxy            string `json:"HttpProxy" xml:"HttpProxy"`
	HttpsProxy           string `json:"HttpsProxy" xml:"HttpsProxy"`
	NoProxy              string `json:"NoProxy" xml:"NoProxy"`
	LocalAddr            string `json:"LocalAddr" xml:"LocalAddr"`
	MaxIdleConns         int    `json:"MaxIdleConns" xml:"MaxIdleConns"`
	EndpointType         string `json:"EndpointType" xml:"EndpointType"`
	OpenPlatformEndpoint string `json:"OpenPlatformEndpoint" xml:"OpenPlatformEndpoint"`
	credential           credentials.Credential
}

var defaultUserAgent = fmt.Sprintf("AlibabaCloud (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), "0.01")

func (client *BaseClient) InitClient(config map[string]interface{}) (err error) {
	client.RegionId = getStringValue(config["regionId"])
	client.Protocol = getStringValue(config["protocol"])
	client.Endpoint = getStringValue(config["endpoint"])
	client.UserAgent = getStringValue(config["userAgent"])
	client.ReadTimeout = getIntValue(config["readTimeout"])
	client.ConnectTimeout = getIntValue(config["connectTimeout"])
	client.HttpProxy = getStringValue(config["httpProxy"])
	client.HttpsProxy = getStringValue(config["httpsProxy"])
	client.NoProxy = getStringValue(config["noProxy"])
	client.LocalAddr = getStringValue(config["localAddr"])
	client.MaxIdleConns = getIntValue(config["maxIdleConns"])
	client.EndpointType = getStringValue(config["endpointType"])
	client.OpenPlatformEndpoint = getStringValue(config["openPlatformEndpoint"])
	conf := &credentials.Configuration{
		AccessKeyID:     getStringValue(config["accessKeyId"]),
		AccessKeySecret: getStringValue(config["accessKeySecret"]),
		Type:            getStringValue(config["type"]),
	}
	if conf.Type == "" {
		conf.Type = "access_key"
	}
	client.credential, err = credentials.NewCredential(conf)
	if err != nil {
		return err
	}
	return nil
}

func (client *BaseClient) GetUserAgent(useragent string) string {
	if useragent == "" {
		return defaultUserAgent
	}
	return defaultUserAgent + " " + useragent
}

// Get Signature according to reqeust and bucketName
func (client *BaseClient) GetSignature(request *tea.Request, secret string) string {
	stringToSign := buildRpcStringToSign(request)
	signature := client.Sign(stringToSign, secret, "&")
	return signature
}

func (client *BaseClient) Sign(stringToSign, accessKeySecret, secretSuffix string) string {
	secret := accessKeySecret + secretSuffix
	signedBytes := shaHmac1(stringToSign, secret)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

// If num is not 0, return num, or return defaultNum
func (client *BaseClient) DefaultNumber(num int, defaultNum int) int {
	if num == 0 {
		return defaultNum
	}
	return num
}

// Parse filter to produce a map[string]string
func (client *BaseClient) Query(filter map[string]interface{}) map[string]string {
	tmp := make(map[string]interface{})
	byt, _ := json.Marshal(filter)
	_ = json.Unmarshal(byt, &tmp)

	result := make(map[string]string)
	for key, value := range tmp {
		filterValue := reflect.ValueOf(value)
		flatRepeatedList(filterValue, result, key)
	}

	return result
}

// Get Date in GMT
func (client *BaseClient) GetTimestamp() string {
	gmt := time.FixedZone("GMT", 0)
	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}

func (client *BaseClient) GetNonce() string {
	return utils.GetUUID()
}

func (client *BaseClient) Json(response *tea.Response) (result map[string]interface{}, err error) {
	body, err := response.ReadBody()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}

func (client *BaseClient) GetEndpoint(product string, regionid string) string {
	return client.Endpoint
}

func (client *BaseClient) GetAccessKeyId() string {
	if client.credential == nil {
		return ""
	}
	accesskey, err := client.credential.GetAccessKeyId()
	if err != nil {
		return ""
	}
	return accesskey
}

func (client *BaseClient) GetAccessKeySecret() string {
	if client.credential == nil {
		return ""
	}
	accesssecret, err := client.credential.GetAccessKeySecret()
	if err != nil {
		return ""
	}
	return accesssecret
}

// Determine whether the request failed
func (client *BaseClient) HasError(body map[string]interface{}) bool {
	if body == nil {
		return true
	}
	if obj := body["Code"]; obj != nil {
		if statusCode := body["Code"].(string); statusCode != "" {
			return true
		}
	}
	return false
}

// If realStr is not "", return realStr, or return defaultStr
func (client *BaseClient) Default(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}
	return str
}
