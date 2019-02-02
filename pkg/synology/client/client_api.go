package client

type RequestParams map[string]string

type Request interface {
	GetCGIPath() string
	GetAPIName() string
	GetVersion() int
	GetMethod() string
	GetParams() RequestParams
}

type Object interface {
}

type Client interface {
	Call(Request, Object) error
}
