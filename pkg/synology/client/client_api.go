package client

type RequestParams map[string]string

type ErrorCodeMapper func(errCode ErrorCode) string

type Request struct {
	CGIPath         string
	APIName         string
	Version         int
	Method          string
	Params          RequestParams
	ErrorCodeMapper ErrorCodeMapper
}

type Object interface {
}

type Client interface {
	Call(Request, Object) error
}

func NewRequest(cgiPath, apiName string, version int, method string, params RequestParams, errorCodeMapper ErrorCodeMapper) Request {
	return Request{
		CGIPath:         cgiPath,
		APIName:         apiName,
		Version:         version,
		Method:          method,
		Params:          params,
		ErrorCodeMapper: errorCodeMapper,
	}
}
