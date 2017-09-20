package cloudfront

type CloudFrontHeader struct {
	CloudFrontForwardedProto  string `json:"CloudFront-Forwarded-Proto"`
	CloudFrontIsDesktopViewer string `json:"CloudFront-Is-Desktop-Viewer"`
	CloudFrontIsMobileViewer  string `json:"CloudFront-Is-Mobile-Viewer"`
	CloudFrontIsSmartTVViewer string `json:"CloudFront-Is-SmartTV-Viewer"`
	CloudFrontIsTabletViewer  string `json:"CloudFront-Is-Tablet-Viewer"`
	CloudFrontViewerCountry   string `json:"CloudFront-Viewer-Country"`
}
