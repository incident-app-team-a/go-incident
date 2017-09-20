package apigateway

import (
	"encoding/json"
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {
	rawJson := `
{
    "body-json": "Direction=.....",
    "params": {
        "path": {},
        "querystring": {},
        "header": {
            "Accept": "*/*",
            "Accept-Encoding": "gzip, deflate",
            "CloudFront-Forwarded-Proto": "https",
            "CloudFront-Is-Desktop-Viewer": "true",
            "CloudFront-Is-Mobile-Viewer": "false",
            "CloudFront-Is-SmartTV-Viewer": "false",
            "CloudFront-Is-Tablet-Viewer": "false",
            "CloudFront-Viewer-Country": "US",
            "Content-Type": "application/x-www-form-urlencoded",
            "Host": "xxxxxx.execute-api.ap-northeast-1.amazonaws.com",
            "User-Agent": "Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.35 Safari/535.1",
            "Via": "1.1 ABC.cloudfront.net (CloudFront)",
            "X-Amz-Cf-Id": "ABC",
            "X-Amzn-Trace-Id": "Root=1-0123456798",
            "X-Forwarded-For": "127.0.0.1",
            "X-Forwarded-Port": "443",
            "X-Forwarded-Proto": "https",
            "X-Plivo-Cloud": "v1",
            "X-Plivo-Signature": "ABCD=="
        }
    },
    "context": {
        "account-id": "",
        "api-id": "abcdefg",
        "api-key": "",
        "authorizer-principal-id": "",
        "caller": "",
        "cognito-authentication-provider": "",
        "cognito-authentication-type": "",
        "cognito-identity-id": "",
        "cognito-identity-pool-id": "",
        "http-method": "POST",
        "stage": "stg",
        "source-ip": "127.0.0.1",
        "user": "",
        "user-agent": "Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.35 Safari/535.1",
        "user-arn": "",
        "request-id": "abcdefg-013546",
        "resource-id": "abcdefg",
        "resource-path": "/plivo/ring"
    }
}`

	var req Request
	if err := json.Unmarshal([]byte(rawJson), &req); err != nil {
		t.Fatalf("unexpected exception: %v", err)
	}
}
