// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// RatecardClient a client for Ratecard
type RatecardClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRatecardClientWithConfigurationProvider Creates a new default Ratecard client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRatecardClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RatecardClient, err error) {
	if enabled := common.CheckForEnabledServices("onesubscription"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newRatecardClientFromBaseClient(baseClient, provider)
}

// NewRatecardClientWithOboToken Creates a new default Ratecard client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewRatecardClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RatecardClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRatecardClientFromBaseClient(baseClient, configProvider)
}

func newRatecardClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RatecardClient, err error) {
	// Ratecard service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Ratecard"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RatecardClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RatecardClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("onesubscription", "https://identity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RatecardClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *RatecardClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListRateCards List API that returns all ratecards for given Subscription Id and Account ID (if provided) and
// for a particular date range
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListRateCards.go.html to see an example of how to use ListRateCards API.
func (client RatecardClient) ListRateCards(ctx context.Context, request ListRateCardsRequest) (response ListRateCardsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRateCards, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRateCardsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRateCardsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRateCardsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRateCardsResponse")
	}
	return
}

// listRateCards implements the OCIOperation interface (enables retrying operations)
func (client RatecardClient) listRateCards(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ratecards", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRateCardsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Ratecard", "ListRateCards", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
