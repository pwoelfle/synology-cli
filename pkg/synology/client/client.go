package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

var _ Client = new(client)

type client struct {
	baseURL    url.URL
	httpClient *http.Client
}

// NewClient creates a new Client object with the given parameters.
func NewClient(host string) (Client, error) {
	if len(host) == 0 {
		return nil, fmt.Errorf("host must be defined")
	}

	baseURL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := client{
		baseURL: *baseURL,
		httpClient: &http.Client{
			Jar: cookieJar,
		},
	}

	return client, nil
}

func (c client) requestGet(path string, params map[string]string, object Object) error {
	requestURL := c.baseURL

	if len(requestURL.Path) > 0 {
		requestURL.Path = requestURL.Path + "/" + path
	} else {
		requestURL.Path = path
	}

	requestParams := requestURL.Query()
	for k, v := range params {
		requestParams.Add(k, v)
	}

	requestURL.RawQuery = requestParams.Encode()

	response, err := c.httpClient.Get(requestURL.String())
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseObject map[string]*json.RawMessage
	err = json.Unmarshal(body, &responseObject)
	if err != nil {
		return err
	}

	var success bool
	err = json.Unmarshal(*responseObject["success"], &success)
	if err != nil {
		return err
	}

	if !success {
		var Error Error
		if err = json.Unmarshal(*responseObject["error"], &Error); err != nil {
			return err
		}
		return Error
	}

	if object != nil {
		if err = json.Unmarshal(*responseObject["data"], object); err != nil {
			return err
		}
	}

	return nil
}

func (c client) Call(req Request, obj Object) error {
	path := fmt.Sprintf("/webapi/%s", req.GetCGIPath())
	params := map[string]string{
		"api":     req.GetAPIName(),
		"version": strconv.Itoa(req.GetVersion()),
		"method":  req.GetMethod(),
	}

	for k, v := range req.GetParams() {
		params[k] = v
	}
	return c.requestGet(path, params, obj)
}
