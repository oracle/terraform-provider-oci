// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//AccountClient a client for Account
type AccountClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAccountClientWithConfigurationProvider Creates a new default Account client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAccountClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AccountClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newAccountClientFromBaseClient(baseClient, provider)
}

// NewAccountClientWithOboToken Creates a new default Account client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewAccountClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AccountClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAccountClientFromBaseClient(baseClient, configProvider)
}

func newAccountClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AccountClient, err error) {
	// Account service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AccountClient{BaseClient: baseClient}
	client.BasePath = "20181001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AccountClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplace", "https://marketplace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AccountClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AccountClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetLaunchEligibility Returns Tenant eligibility and other information for launching a PIC image
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetLaunchEligibility.go.html to see an example of how to use GetLaunchEligibility API.
func (client AccountClient) GetLaunchEligibility(ctx context.Context, request GetLaunchEligibilityRequest) (response GetLaunchEligibilityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLaunchEligibility, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLaunchEligibilityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLaunchEligibilityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLaunchEligibilityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLaunchEligibilityResponse")
	}
	return
}

// getLaunchEligibility implements the OCIOperation interface (enables retrying operations)
func (client AccountClient) getLaunchEligibility(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/launchEligibility", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLaunchEligibilityResponse
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

// GetThirdPartyPaidListingEligibility Returns eligibility details of the tenancy to see and launch third party paid listings
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetThirdPartyPaidListingEligibility.go.html to see an example of how to use GetThirdPartyPaidListingEligibility API.
func (client AccountClient) GetThirdPartyPaidListingEligibility(ctx context.Context, request GetThirdPartyPaidListingEligibilityRequest) (response GetThirdPartyPaidListingEligibilityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getThirdPartyPaidListingEligibility, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetThirdPartyPaidListingEligibilityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetThirdPartyPaidListingEligibilityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetThirdPartyPaidListingEligibilityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetThirdPartyPaidListingEligibilityResponse")
	}
	return
}

// getThirdPartyPaidListingEligibility implements the OCIOperation interface (enables retrying operations)
func (client AccountClient) getThirdPartyPaidListingEligibility(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/thirdPartyPaidListingEligibility", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetThirdPartyPaidListingEligibilityResponse
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
