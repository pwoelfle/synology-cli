package client

var _ Request = new(request)

type request struct {
	cgiPath string
	apiName string
	version int
	method  string
	params  RequestParams
}

// NewRequest creates a new Request for .
func NewRequest(cgiPath, apiName string, version int, method string, params RequestParams) Request {
	return &request{
		cgiPath: cgiPath,
		apiName: apiName,
		version: version,
		method:  method,
		params:  params,
	}
}

func (r request) GetCGIPath() string {
	return r.cgiPath
}

func (r request) GetAPIName() string {
	return r.apiName
}

func (r request) GetVersion() int {
	return r.version
}

func (r request) GetMethod() string {
	return r.method
}

func (r request) GetParams() RequestParams {
	return r.params
}
