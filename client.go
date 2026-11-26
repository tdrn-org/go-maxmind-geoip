//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package maxmindgeoip

import (
	"context"
	"net/http"
	"net/url"
)

// GeoIPEndpointURL is the Endpoint URL for the GeoIP web service.
var GeoIPEndpointURL *url.URL = func() *url.URL {
	url, err := url.Parse("https://geoip.maxmind.com/geoip/v2.1/")
	if err != nil {
		panic(err)
	}
	return url
}()

// GeoLiteEndpointURL is the Endpoint URL for the GeoLite web service.
var GeoLiteEndpointURL *url.URL = func() *url.URL {
	url, err := url.Parse("https://geolite.info/geoip/v2.1/")
	if err != nil {
		panic(err)
	}
	return url
}()

// NewClient creates a new Client instance using the given
// endpoint URL as well as Client options.
func NewClient(apiURL *url.URL, username, password string, options ...ClientOption) (*Client, error) {
	prepareReq := func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(username, password)
		req.Header.Set("Accept", "application/json")
		return nil
	}
	return newClient(apiURL, append(options, WithRequestEditors(prepareReq))...)
}
