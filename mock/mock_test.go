//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package mock_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	maxmindgeoip "github.com/tdrn-org/go-maxmind-geoip"
	"github.com/tdrn-org/go-maxmind-geoip/api"
	"github.com/tdrn-org/go-maxmind-geoip/mock"
)

var testGeonameId int = 123456
var testCityNames api.Names = map[string]string{"en": "a city"}
var testCity *api.City = &api.City{
	GeonameId: &testGeonameId,
	Names:     &testCityNames,
}

func TestGetCity(t *testing.T) {
	mock := mock.StartMockServer(&mockHandler{})
	defer mock.Stop(t.Context())
	client, err := maxmindgeoip.NewClient(mock.APIURL(), "", "")
	require.NoError(t, err)

	response, err := client.GetCity(t.Context(), api.IPAddress("8.8.8.8"), &api.GetCityParams{})
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, http.StatusOK, response.StatusCode())
	require.NotNil(t, response.ApplicationvndMaxmindComCityJSON200)
	require.NotNil(t, response.ApplicationvndMaxmindComCityJSON200.City)
	require.Equal(t, testCity, response.ApplicationvndMaxmindComCityJSON200.City)
}

type mockHandler struct{}

func (h *mockHandler) GetCity(ctx context.Context, request api.GetCityRequestObject) (api.GetCityResponseObject, error) {
	return &api.GetCity200ApplicationVndMaxmindComCityPlusJSONResponse{City: testCity}, nil
}
