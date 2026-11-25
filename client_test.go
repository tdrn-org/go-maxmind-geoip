//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package maxmindgeoip_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	maxmindgeoip "github.com/tdrn-org/go-maxmind-geoip"
	"github.com/tdrn-org/go-maxmind-geoip/api"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestGetCity(t *testing.T) {
	client := newClient(t)
	response, err := client.GetCity(t.Context(), api.IPAddress("8.8.8.8"), &api.GetCityParams{})
	require.NoError(t, err)
	_ = response
}

func TestGetCountry(t *testing.T) {
	client := newClient(t)
	response, err := client.GetCountry(t.Context(), api.IPAddress("8.8.8.8"), &api.GetCountryParams{})
	require.NoError(t, err)
	_ = response
}

func newClient(t *testing.T) *maxmindgeoip.Client {
	// Enable HTTP recording
	maskedAPIKey := "dGVzdDp0ZXN0"
	maskAuthorizationHeader := func(i *cassette.Interaction) error {
		i.Request.Headers["Authorization"] = []string{"Basic " + maskedAPIKey}
		return nil
	}
	vcr, err := recorder.New(filepath.Join("testdata", strings.ReplaceAll(t.Name(), "/", "_")), recorder.WithHook(maskAuthorizationHeader, recorder.AfterCaptureHook))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, vcr.Stop())
	})
	username := "test"
	password := "test"
	if vcr.IsRecording() {
		username = os.Getenv("ACCOUNT_ID")
		password = os.Getenv("LICENSE_KEY")
	}
	client, err := maxmindgeoip.NewClient(maxmindgeoip.GeoLiteEndpointURL, username, password, maxmindgeoip.WithHttpClient(vcr.GetDefaultClient()))
	require.NoError(t, err)
	return client
}
