// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//DataplaneClient a client for Dataplane
type DataplaneClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataplaneClientWithConfigurationProvider Creates a new default Dataplane client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataplaneClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataplaneClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDataplaneClientFromBaseClient(baseClient, provider)
}

// NewDataplaneClientWithOboToken Creates a new default Dataplane client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDataplaneClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataplaneClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataplaneClientFromBaseClient(baseClient, configProvider)
}

func newDataplaneClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataplaneClient, err error) {
	// Dataplane service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataplaneClient{BaseClient: baseClient}
	client.BasePath = "v1"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataplaneClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("identitydataplane", "https://auth.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataplaneClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DataplaneClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GenerateScopedAccessToken Based on the calling principal and the input payload, derive the claims and create a security token.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydataplane/GenerateScopedAccessToken.go.html to see an example of how to use GenerateScopedAccessToken API.
func (client DataplaneClient) GenerateScopedAccessToken(ctx context.Context, request GenerateScopedAccessTokenRequest) (response GenerateScopedAccessTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateScopedAccessToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateScopedAccessTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateScopedAccessTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateScopedAccessTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateScopedAccessTokenResponse")
	}
	return
}

// generateScopedAccessToken implements the OCIOperation interface (enables retrying operations)
func (client DataplaneClient) generateScopedAccessToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/generateScopedAccessToken", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateScopedAccessTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
