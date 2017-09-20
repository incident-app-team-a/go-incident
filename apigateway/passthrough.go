package apigateway

import (
	"github.com/epy0n0ff/go-incident/cloudfront"
	"github.com/epy0n0ff/go-incident/plivo"
)

type Request struct {
	BodyJson string `json:"body-json"`
	Params   Params `json:"params"`
}

type Params struct {
	Path        interface{} `json:"path"`
	QueryString interface{} `json:"querystring"`
	Header      Header      `json:"header"`
	Context     Context     `json:"context"`
}

type Header struct {
	Accept          string `json:"Accept"`
	AcceptEncoding  string `json:"Accept-Encoding"`
	ContentType     string `json:"Content-Type"`
	Host            string `json:"Host"`
	UserAgent       string `json:"User-Agent"`
	Via             string `json:"Via"`
	XForwardedFor   string `json:"X-Forwarded-For"`
	XForwardedPort  string `json:"X-Forwarded-Port"`
	XForwardedProto string `json:"X-Forwarded-Proto"`
	XAmzCfID        string `json:"X-Amz-Cf-Id"`
	XAmznTraceID    string `json:"X-Amzn-Trace-Id"`
	cloudfront.CloudFrontHeader
	plivo.PlivoHeader
}

type Context struct {
	AccountId                     string `json:"account-id"`
	ApiID                         string `json:"api-id"`
	ApiKey                        string `json:"api-key"`
	AuthorizerPrincipalID         string `json:"authorizer-principal-id"`
	Caller                        string `json:"caller"`
	CognitoAuthenticationProvider string `json:"cognito-authentication-provider"`
	CognitoAuthenticationType     string `json:"cognito-authentication-type"`
	CognitoIdentityID             string `json:"cognito-identity-id"`
	CognitoIdentityPoolID         string `json:"cognito-identity-pool-id"`
	HttpMethod                    string `json:"http-method"`
	Stage                         string `json:"stage"`
	SourceIP                      string `json:"source-ip"`
	User                          string `json:"user"`
	UserAgent                     string `json:"user-agent"`
	UserARN                       string `json:"user-arn"`
	RequestID                     string `json:"request-id"`
	ResourceID                    string `json:"resource-id"`
	ResourcePath                  string `json:"resource-path"`
}
