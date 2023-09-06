package aplus

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// HttpRequestDoer Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// NewClient Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

type ClientInterface interface {
	SearchContentDocuments(ctx context.Context, params *SearchContentDocumentsParams) (*http.Response, error)
	SearchContentPublishRecords(ctx context.Context, asin string, params *SearchContentPublishRecordsParams) (*http.Response, error)
}

// NewSearchContentDocumentsRequest generates requests for SearchContentDocuments
func NewSearchContentDocumentsRequest(endpoint string, params *SearchContentDocumentsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := "/aplus/2020-11-01/contentDocuments"
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()
	queryValues.Add("marketplaceId", params.MarketplaceId)

	if params.PageToken != nil {
		queryValues.Add("pageToken", *params.PageToken)
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) SearchContentDocuments(ctx context.Context, params *SearchContentDocumentsParams) (*http.Response, error) {
	req, err := NewSearchContentDocumentsRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewSearchContentPublishRecordsRequest generates requests for SearchContentPublishRecords
func NewSearchContentPublishRecordsRequest(endpoint string, asin string, params *SearchContentPublishRecordsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := "/aplus/2020-11-01/contentPublishRecords"
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()
	queryValues.Add("marketplaceId", params.MarketplaceId)
	queryValues.Add("asin", asin)

	if params.PageToken != nil {
		queryValues.Add("pageToken", *params.PageToken)
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) SearchContentPublishRecords(ctx context.Context, asin string, params *SearchContentPublishRecordsParams) (*http.Response, error) {
	req, err := NewSearchContentPublishRecordsRequest(c.Endpoint, asin, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

type ClientWithResponseInterface interface {
	SearchContentDocumentsWithResponse(ctx context.Context, params *SearchContentDocumentsParams) (*http.Response, error)
	SearchContentPublishRecordsWithResponse(ctx context.Context, asin string, params *SearchContentPublishRecordsParams) (*http.Response, error)
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

func (c *ClientWithResponses) SearchContentDocumentsWithResponse(ctx context.Context, params *SearchContentDocumentsParams) (*SearchContentDocumentsResp, error) {
	rsp, err := c.SearchContentDocuments(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseSearchContentDocumentsResp(rsp)
}

func ParseSearchContentDocumentsResp(rsp *http.Response) (*SearchContentDocumentsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &SearchContentDocumentsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest SearchContentDocumentsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

func (c *ClientWithResponses) SearchContentPublishRecordsWithResponse(ctx context.Context, asin string, params *SearchContentPublishRecordsParams) (*SearchContentPublishRecordsResp, error) {
	rsp, err := c.SearchContentPublishRecords(ctx, asin, params)
	if err != nil {
		return nil, err
	}
	return ParseSearchContentPublishRecordsResp(rsp)
}

func ParseSearchContentPublishRecordsResp(rsp *http.Response) (*SearchContentPublishRecordsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &SearchContentPublishRecordsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest SearchContentPublishRecordsResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// StatusCode returns HTTPResponse.StatusCode
func (r SearchContentPublishRecordsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
