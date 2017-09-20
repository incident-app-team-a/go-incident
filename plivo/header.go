package plivo

type PlivoHeader struct {
	XPlivoCloud     string `json:"X-Plivo-Cloud"`
	XPlivoSignature string `json:"X-Plivo-Signature"`
}
