// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// Defines values for ModelAuthType.
const (
	AuthTypeBasic  ModelAuthType = "basic"
	AuthTypeBearer ModelAuthType = "bearer"
	AuthTypeNone   ModelAuthType = "none"
)

// Defines values for ModelJobStatus.
const (
	JobStatusRunning ModelJobStatus = "RUNNING"
	JobStatusStopped ModelJobStatus = "STOPPED"
)

// Defines values for ModelJobType.
const (
	JobTypeAMQP ModelJobType = "AMQP"
	JobTypeHTTP ModelJobType = "HTTP"
)

// HandlersErrorResponse defines model for handlers.ErrorResponse.
type HandlersErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// ModelAMQPJob defines model for model.AMQPJob.
type ModelAMQPJob struct {
	// Body e.g., "Hello, world!"
	Body *string `json:"body,omitempty"`

	// Connection e.g., "amqp://guest:guest@localhost:5672/"
	Connection *string `json:"connection,omitempty"`

	// ContentType e.g., "text/plain"
	ContentType *string `json:"content_type,omitempty"`

	// Exchange e.g., "my_exchange"
	Exchange *string `json:"exchange,omitempty"`

	// Headers e.g., {"x-delay": 10000}
	Headers *map[string]map[string]interface{} `json:"headers,omitempty"`

	// RoutingKey e.g., "my_routing_key"
	RoutingKey *string `json:"routing_key,omitempty"`
}

// ModelAuth defines model for model.Auth.
type ModelAuth struct {
	// BearerToken for "bearer"
	BearerToken *string `json:"bearer_token,omitempty"`

	// Password for "basic"
	Password *string `json:"password,omitempty"`

	// Type e.g., "none", "basic", "bearer"
	Type *ModelAuthType `json:"type,omitempty"`

	// Username for "basic"
	Username *string `json:"username,omitempty"`
}

// ModelAuthType defines model for model.AuthType.
type ModelAuthType string

// ModelHTTPJob defines model for model.HTTPJob.
type ModelHTTPJob struct {
	// Auth e.g., {"type": "basic", "username": "foo", "password": "bar"}
	Auth *ModelAuth `json:"auth,omitempty"`

	// Body e.g., "{\"hello\": \"world\"}"
	Body *string `json:"body,omitempty"`

	// Headers e.g., {"Content-Type": "application/json"}
	Headers *map[string]string `json:"headers,omitempty"`

	// Method e.g., "GET", "POST", "PUT", "PATCH", "DELETE"
	Method *string `json:"method,omitempty"`

	// Url e.g., "https://example.com"
	Url *string `json:"url,omitempty"`

	// ValidResponseCodes e.g., [200, 201, 202]
	ValidResponseCodes *[]int `json:"valid_response_codes,omitempty"`
}

// ModelJob defines model for model.Job.
type ModelJob struct {
	AmqpJob   *ModelAMQPJob `json:"amqp_job,omitempty"`
	CreatedAt *string       `json:"created_at,omitempty"`

	// CronSchedule for recurring jobs
	CronSchedule *string `json:"cron_schedule,omitempty"`

	// ExecuteAt for one-off jobs
	ExecuteAt *string       `json:"execute_at,omitempty"`
	HttpJob   *ModelHTTPJob `json:"http_job,omitempty"`
	Id        *string       `json:"id,omitempty"`

	// NextRun when the job is scheduled to run next (can be null if the job is not scheduled to run again)
	NextRun   *string         `json:"next_run,omitempty"`
	Status    *ModelJobStatus `json:"status,omitempty"`
	Tags      *[]string       `json:"tags,omitempty"`
	Type      *ModelJobType   `json:"type,omitempty"`
	UpdatedAt *string         `json:"updated_at,omitempty"`
}

// ModelJobCreate defines model for model.JobCreate.
type ModelJobCreate struct {
	AmqpJob *ModelAMQPJob `json:"amqp_job,omitempty"`

	// CronSchedule for recurring jobs
	CronSchedule *string `json:"cron_schedule,omitempty"`

	// ExecuteAt ExecuteAt and CronSchedule are mutually exclusive.
	ExecuteAt *string `json:"execute_at,omitempty"`

	// HttpJob HTTPJob and AMQPJob are mutually exclusive.
	HttpJob *ModelHTTPJob `json:"http_job,omitempty"`
	Tags    *[]string     `json:"tags,omitempty"`

	// Type Job type
	Type *ModelJobType `json:"type,omitempty"`
}

// ModelJobExecution defines model for model.JobExecution.
type ModelJobExecution struct {
	EndTime      *string `json:"end_time,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	Id           *int    `json:"id,omitempty"`
	JobId        *string `json:"job_id,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	Success      *bool   `json:"success,omitempty"`
}

// ModelJobStatus defines model for model.JobStatus.
type ModelJobStatus string

// ModelJobType defines model for model.JobType.
type ModelJobType string

// ModelJobUpdate defines model for model.JobUpdate.
type ModelJobUpdate struct {
	Amqp         *ModelAMQPJob `json:"amqp,omitempty"`
	CronSchedule *string       `json:"cron_schedule,omitempty"`
	ExecuteAt    *string       `json:"execute_at,omitempty"`
	Http         *ModelHTTPJob `json:"http,omitempty"`
	Tags         *[]string     `json:"tags,omitempty"`
	Type         *ModelJobType `json:"type,omitempty"`
}

// GetJobsParams defines parameters for GetJobs.
type GetJobsParams struct {
	// Limit Limit
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Offset
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Tags Tags
	Tags *[]interface{} `form:"tags,omitempty" json:"tags,omitempty"`
}

// GetJobsIdExecutionsParams defines parameters for GetJobsIdExecutions.
type GetJobsIdExecutionsParams struct {
	// FailedOnly Failed Only
	FailedOnly *bool `form:"failedOnly,omitempty" json:"failedOnly,omitempty"`

	// Limit Limit
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Offset
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// PostJobsJSONRequestBody defines body for PostJobs for application/json ContentType.
type PostJobsJSONRequestBody = ModelJobCreate

// PutJobsIdJSONRequestBody defines body for PutJobsId for application/json ContentType.
type PutJobsIdJSONRequestBody = ModelJobUpdate

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
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
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
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

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetJobs request
	GetJobs(ctx context.Context, params *GetJobsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostJobsWithBody request with any body
	PostJobsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostJobs(ctx context.Context, body PostJobsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteJobsId request
	DeleteJobsId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetJobsId request
	GetJobsId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutJobsIdWithBody request with any body
	PutJobsIdWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutJobsId(ctx context.Context, id string, body PutJobsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetJobsIdExecutions request
	GetJobsIdExecutions(ctx context.Context, id string, params *GetJobsIdExecutionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetJobs(ctx context.Context, params *GetJobsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetJobsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostJobsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostJobsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostJobs(ctx context.Context, body PostJobsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostJobsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteJobsId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteJobsIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetJobsId(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetJobsIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutJobsIdWithBody(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutJobsIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutJobsId(ctx context.Context, id string, body PutJobsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutJobsIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetJobsIdExecutions(ctx context.Context, id string, params *GetJobsIdExecutionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetJobsIdExecutionsRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetJobsRequest generates requests for GetJobs
func NewGetJobsRequest(server string, params *GetJobsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Limit != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Offset != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "offset", runtime.ParamLocationQuery, *params.Offset); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Tags != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", false, "tags", runtime.ParamLocationQuery, *params.Tags); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostJobsRequest calls the generic PostJobs builder with application/json body
func NewPostJobsRequest(server string, body PostJobsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostJobsRequestWithBody(server, "application/json", bodyReader)
}

// NewPostJobsRequestWithBody generates requests for PostJobs with any type of body
func NewPostJobsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteJobsIdRequest generates requests for DeleteJobsId
func NewDeleteJobsIdRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetJobsIdRequest generates requests for GetJobsId
func NewGetJobsIdRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPutJobsIdRequest calls the generic PutJobsId builder with application/json body
func NewPutJobsIdRequest(server string, id string, body PutJobsIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutJobsIdRequestWithBody(server, id, "application/json", bodyReader)
}

// NewPutJobsIdRequestWithBody generates requests for PutJobsId with any type of body
func NewPutJobsIdRequestWithBody(server string, id string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetJobsIdExecutionsRequest generates requests for GetJobsIdExecutions
func NewGetJobsIdExecutionsRequest(server string, id string, params *GetJobsIdExecutionsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/jobs/%s/executions", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.FailedOnly != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "failedOnly", runtime.ParamLocationQuery, *params.FailedOnly); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Limit != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Offset != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "offset", runtime.ParamLocationQuery, *params.Offset); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetJobsWithResponse request
	GetJobsWithResponse(ctx context.Context, params *GetJobsParams, reqEditors ...RequestEditorFn) (*GetJobsResponse, error)

	// PostJobsWithBodyWithResponse request with any body
	PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostJobsResponse, error)

	PostJobsWithResponse(ctx context.Context, body PostJobsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostJobsResponse, error)

	// DeleteJobsIdWithResponse request
	DeleteJobsIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteJobsIdResponse, error)

	// GetJobsIdWithResponse request
	GetJobsIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetJobsIdResponse, error)

	// PutJobsIdWithBodyWithResponse request with any body
	PutJobsIdWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutJobsIdResponse, error)

	PutJobsIdWithResponse(ctx context.Context, id string, body PutJobsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutJobsIdResponse, error)

	// GetJobsIdExecutionsWithResponse request
	GetJobsIdExecutionsWithResponse(ctx context.Context, id string, params *GetJobsIdExecutionsParams, reqEditors ...RequestEditorFn) (*GetJobsIdExecutionsResponse, error)
}

type GetJobsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]ModelJob
	JSON400      *HandlersErrorResponse
	JSON500      *HandlersErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetJobsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetJobsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostJobsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ModelJob
	JSON400      *HandlersErrorResponse
	JSON500      *HandlersErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostJobsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostJobsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteJobsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *HandlersErrorResponse
	JSON500      *HandlersErrorResponse
}

// Status returns HTTPResponse.Status
func (r DeleteJobsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteJobsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetJobsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ModelJob
	JSON400      *HandlersErrorResponse
	JSON500      *HandlersErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetJobsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetJobsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutJobsIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ModelJob
	JSON400      *HandlersErrorResponse
	JSON500      *HandlersErrorResponse
}

// Status returns HTTPResponse.Status
func (r PutJobsIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutJobsIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetJobsIdExecutionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]ModelJobExecution
	JSON400      *HandlersErrorResponse
	JSON500      *HandlersErrorResponse
}

// Status returns HTTPResponse.Status
func (r GetJobsIdExecutionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetJobsIdExecutionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetJobsWithResponse request returning *GetJobsResponse
func (c *ClientWithResponses) GetJobsWithResponse(ctx context.Context, params *GetJobsParams, reqEditors ...RequestEditorFn) (*GetJobsResponse, error) {
	rsp, err := c.GetJobs(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetJobsResponse(rsp)
}

// PostJobsWithBodyWithResponse request with arbitrary body returning *PostJobsResponse
func (c *ClientWithResponses) PostJobsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostJobsResponse, error) {
	rsp, err := c.PostJobsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostJobsResponse(rsp)
}

func (c *ClientWithResponses) PostJobsWithResponse(ctx context.Context, body PostJobsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostJobsResponse, error) {
	rsp, err := c.PostJobs(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostJobsResponse(rsp)
}

// DeleteJobsIdWithResponse request returning *DeleteJobsIdResponse
func (c *ClientWithResponses) DeleteJobsIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*DeleteJobsIdResponse, error) {
	rsp, err := c.DeleteJobsId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteJobsIdResponse(rsp)
}

// GetJobsIdWithResponse request returning *GetJobsIdResponse
func (c *ClientWithResponses) GetJobsIdWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetJobsIdResponse, error) {
	rsp, err := c.GetJobsId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetJobsIdResponse(rsp)
}

// PutJobsIdWithBodyWithResponse request with arbitrary body returning *PutJobsIdResponse
func (c *ClientWithResponses) PutJobsIdWithBodyWithResponse(ctx context.Context, id string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutJobsIdResponse, error) {
	rsp, err := c.PutJobsIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutJobsIdResponse(rsp)
}

func (c *ClientWithResponses) PutJobsIdWithResponse(ctx context.Context, id string, body PutJobsIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutJobsIdResponse, error) {
	rsp, err := c.PutJobsId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutJobsIdResponse(rsp)
}

// GetJobsIdExecutionsWithResponse request returning *GetJobsIdExecutionsResponse
func (c *ClientWithResponses) GetJobsIdExecutionsWithResponse(ctx context.Context, id string, params *GetJobsIdExecutionsParams, reqEditors ...RequestEditorFn) (*GetJobsIdExecutionsResponse, error) {
	rsp, err := c.GetJobsIdExecutions(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetJobsIdExecutionsResponse(rsp)
}

// ParseGetJobsResponse parses an HTTP response from a GetJobsWithResponse call
func ParseGetJobsResponse(rsp *http.Response) (*GetJobsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetJobsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []ModelJob
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostJobsResponse parses an HTTP response from a PostJobsWithResponse call
func ParsePostJobsResponse(rsp *http.Response) (*PostJobsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostJobsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ModelJob
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteJobsIdResponse parses an HTTP response from a DeleteJobsIdWithResponse call
func ParseDeleteJobsIdResponse(rsp *http.Response) (*DeleteJobsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteJobsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetJobsIdResponse parses an HTTP response from a GetJobsIdWithResponse call
func ParseGetJobsIdResponse(rsp *http.Response) (*GetJobsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetJobsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ModelJob
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePutJobsIdResponse parses an HTTP response from a PutJobsIdWithResponse call
func ParsePutJobsIdResponse(rsp *http.Response) (*PutJobsIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutJobsIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ModelJob
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetJobsIdExecutionsResponse parses an HTTP response from a GetJobsIdExecutionsWithResponse call
func ParseGetJobsIdExecutionsResponse(rsp *http.Response) (*GetJobsIdExecutionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetJobsIdExecutionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []ModelJobExecution
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest HandlersErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
