package helpers

type HttpHeaders struct {
	Key   string
	Value string
}

type HttpReq struct {
	Method  string
	Headers []HttpHeaders
	Url     string
}
