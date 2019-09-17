package rpcclient

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
	"github.com/aliyun/rpc-client/utils"
)

// client is for calling oss api
type Client struct {
	RegionId       string `json:"RegionId" xml:"RegionId"`
	Protocol       string `json:"Protocol" xml:"Protocol"`
	Endpoint       string `json:"Endpoint" xml:"Endpoint"`
	UserAgent      string `json:"UserAgent" xml:"UserAgent"`
	SecurityToken  string `json:"SecurityToken" xml:"SecurityToken"`
	ReadTimeout    int    `json:"ReadTimeout" xml:"ReadTimeout"`
	ConnectTimeout int    `json:"ConnectTimeout" xml:"ConnectTimeout"`
	HttpProxy      string `json:"HttpProxy" xml:"HttpProxy"`
	HttpsProxy     string `json:"HttpsProxy" xml:"HttpsProxy"`
	NoProxy        string `json:"NoProxy" xml:"NoProxy"`
	LocalAddr      string `json:"LocalAddr" xml:"LocalAddr"`
	MaxIdleConns   int    `json:"MaxIdleConns" xml:"MaxIdleConns"`
}

var defaultUserAgent = fmt.Sprintf("AlibabaCloud (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), "0.01")
var credential credentials.Credential

func (client *Client) InitWithConfig(config map[string]string) (err error) {
	client.RegionId = config["regionId"]
	client.Protocol = config["protocol"]
	client.Endpoint = config["endpoint"]
	conf := &credentials.Configuration{
		AccessKeyID:     config["accessKeyId"],
		AccessKeySecret: config["accessKeySecret"],
		Type:            config["type"],
	}
	if conf.Type == "" {
		conf.Type = "access_key"
	}
	credential, err = credentials.NewCredential(conf)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetUserAgent(useragent string) string {
	if useragent == "" {
		return defaultUserAgent
	}
	return defaultUserAgent + " " + useragent
}

// Get Signature according to reqeust and bucketName
func (client *Client) GetSignature(request *tea.Request, secret string) string {
	stringToSign := buildRpcStringToSign(request)
	signature := client.Sign(stringToSign, secret, "&")
	return signature
}

func (client *Client) Sign(stringToSign, accessKeySecret, secretSuffix string) string {
	secret := accessKeySecret + secretSuffix
	signedBytes := shaHmac1(stringToSign, secret)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

func shaHmac1(source, secret string) []byte {
	key := []byte(secret)
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(source))
	return hmac.Sum(nil)
}

func buildRpcStringToSign(request *tea.Request) (stringToSign string) {
	signParams := make(map[string]string)
	for key, value := range request.Query {
		signParams[key] = value
	}

	stringToSign = utils.GetUrlFormedMap(signParams)
	stringToSign = strings.Replace(stringToSign, "+", "%20", -1)
	stringToSign = strings.Replace(stringToSign, "*", "%2A", -1)
	stringToSign = strings.Replace(stringToSign, "%7E", "~", -1)
	stringToSign = url.QueryEscape(stringToSign)
	stringToSign = request.Method + "&%2F&" + stringToSign
	return
}

// Verify whether the parameters meet the requirements
func (client *Client) Validator(params interface{}) error {
	return nil
}

// If num is not 0, return num, or return defaultNum
func (client *Client) DefaultNumber(num interface{}, defaultNum int) int {
	if num == nil {
		return defaultNum
	}
	realnum := num.(int)
	return realnum
}

// Parse filter to produce a map[string]string
func (client *Client) Query(filter map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for key, value := range filter {
		filterValue := reflect.ValueOf(value)
		err := flatRepeatedList(filterValue, result, key)
		if err != nil {
			return nil
		}
	}

	return result
}

// Get Date in GMT
func (client *Client) GetTimestamp() string {
	gmt := time.FixedZone("GMT", 0)
	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}

func (client *Client) GetNonce() string {
	return utils.GetUUID()
}

func (client *Client) Json(response *tea.Response) (result map[string]interface{}, err error) {
	body, err := response.ReadBody()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}

func (client *Client) GetEndpoint(product string, regionid string) string {
	return client.Endpoint
}

func (client *Client) GetAccessKeyId() string {
	accesskey, err := credential.GetAccessKeyID()
	if err != nil {
		return ""
	}
	return accesskey
}

func (client *Client) GetAccessKeySecret() string {
	accesssecret, err := credential.GetAccessSecret()
	if err != nil {
		return ""
	}
	return accesssecret
}

// Determine whether the request failed
func (client *Client) HasError(body map[string]interface{}) bool {
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
func (client *Client) Default(str interface{}, defaultStr string) string {
	if str == nil {
		return defaultStr
	}
	realstr := str.(string)
	return realstr
}
