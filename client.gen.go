// Generated from api.ClientWithResponsesInterface
package maxmindgeoip

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/tdrn-org/go-maxmind-geoip/api"
)

// ErrClientFailure indicates a system error while invoking the REST API.
var ErrClientFailure = errors.New("client call failure")

// ErrAPIFailure indicates an API error while invoking the REST API.
var ErrAPIFailure = errors.New("API failure")

// Client instances are used to access the REST API.
type Client struct {
	apiURL         *url.URL
	apiClient      api.ClientWithResponsesInterface
	httpClient     *http.Client
	logger         *slog.Logger
	requestEditors []api.RequestEditorFn
}

// ClientOption interface is used to set Client options during Client creation.
type ClientOption interface {
	// Apply sets one or more options in the given Client instance.
	Apply(client *Client)
}

// ClientOptionFunc type is used to wrap functions into a ClientOption instance.
type ClientOptionFunc func(client *Client)

// Apply sets one or more options in the given Client instance.
func (f ClientOptionFunc) Apply(client *Client) {
	f(client)
}

// WithHttpClient sets a Client instance's [http.Client] which is used
// to access the REST API.
func WithHttpClient(httpClient *http.Client) ClientOptionFunc {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

// WithLogger sets a Client instance's [slog.Logger] which is used
// for logging.
func WithLogger(logger *slog.Logger) ClientOptionFunc {
	return func(client *Client) {
		client.logger = logger
	}
}

// WithRequestEditors sets the [api.RequestEditorFn]s to applied to to each
// request issued via this Client instance.
func WithRequestEditors(requestEditors ...api.RequestEditorFn) ClientOptionFunc {
	return func(client *Client) {
		client.requestEditors = append(client.requestEditors, requestEditors...)
	}
}

func newClient(apiURL *url.URL, options ...ClientOption) (*Client, error) {
	client := &Client{
		apiURL: apiURL,
	}
	for _, option := range options {
		option.Apply(client)
	}
	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}
	if client.logger == nil {
		client.logger = slog.With(slog.String("client", apiURL.String()))
	}
	apiClient, err := api.NewClientWithResponses(apiURL.String(), api.WithHTTPClient(client.httpClient), api.WithRequestEditorFn(client.prepareRequest()))
	if err != nil {
		return nil, fmt.Errorf("failed to create client (cause: %w)", err)
	}
	client.apiClient = apiClient
	return client, nil
}

// See [api.GetCityWithResponse]
func (client *Client) GetCity(ctx context.Context, ipAddress api.IPAddress, params *api.GetCityParams, reqEditors ...api.RequestEditorFn) (*api.GetCityResponse, error) {
	response, err := client.apiClient.GetCityWithResponse(ctx, ipAddress, params, reqEditors...)
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

// See [api.GetCountryWithResponse]
func (client *Client) GetCountry(ctx context.Context, ipAddress api.IPAddress, params *api.GetCountryParams, reqEditors ...api.RequestEditorFn) (*api.GetCountryResponse, error) {
	response, err := client.apiClient.GetCountryWithResponse(ctx, ipAddress, params, reqEditors...)
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

// See [api.GetInsightsWithResponse]
func (client *Client) GetInsights(ctx context.Context, ipAddress api.IPAddress, params *api.GetInsightsParams, reqEditors ...api.RequestEditorFn) (*api.GetInsightsResponse, error) {
	response, err := client.apiClient.GetInsightsWithResponse(ctx, ipAddress, params, reqEditors...)
	if err != nil {
		return nil, client.wrapSystemError(err)
	}
	err = client.checkAPIResponse(response.HTTPResponse)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (client *Client) prepareRequest() api.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		for _, requestEditor := range client.requestEditors {
			err := requestEditor(ctx, req)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (client *Client) operationID() string {
	pc, _, _, _ := runtime.Caller(3)
	caller := runtime.FuncForPC(pc)
	callerName := caller.Name()
	return callerName[strings.LastIndex(callerName, ".")+1:]
}

func (client *Client) wrapSystemError(err error) error {
	return fmt.Errorf("%w %s (cause: %w)", ErrClientFailure, client.operationID(), err)
}

func (client *Client) checkAPIResponse(httpResponse *http.Response) error {
	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("%w %s (status: %s)", ErrAPIFailure, client.operationID(), httpResponse.Status)
	}
	return nil
}
