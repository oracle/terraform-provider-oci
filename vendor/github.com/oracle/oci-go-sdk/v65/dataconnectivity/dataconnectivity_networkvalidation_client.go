// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//NetworkValidationClient a client for NetworkValidation
type NetworkValidationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewNetworkValidationClientWithConfigurationProvider Creates a new default NetworkValidation client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewNetworkValidationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client NetworkValidationClient, err error) {
	if enabled := common.CheckForEnabledServices("dataconnectivity"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newNetworkValidationClientFromBaseClient(baseClient, provider)
}

// NewNetworkValidationClientWithOboToken Creates a new default NetworkValidation client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewNetworkValidationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client NetworkValidationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newNetworkValidationClientFromBaseClient(baseClient, configProvider)
}

func newNetworkValidationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client NetworkValidationClient, err error) {
	// NetworkValidation service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("NetworkValidation"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = NetworkValidationClient{BaseClient: baseClient}
	client.BasePath = "20210217"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *NetworkValidationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dataconnectivity", "https://dataconnectivity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *NetworkValidationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *NetworkValidationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetNetworkConnectivityStatusCollection This API is used to get the network connectivity status fofor all the data assets attached to the provided private endpoint.
// A default retry strategy applies to this operation GetNetworkConnectivityStatusCollection()
func (client NetworkValidationClient) GetNetworkConnectivityStatusCollection(ctx context.Context, request GetNetworkConnectivityStatusCollectionRequest) (response GetNetworkConnectivityStatusCollectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.getNetworkConnectivityStatusCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNetworkConnectivityStatusCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNetworkConnectivityStatusCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNetworkConnectivityStatusCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNetworkConnectivityStatusCollectionResponse")
	}
	return
}

// getNetworkConnectivityStatusCollection implements the OCIOperation interface (enables retrying operations)
func (client NetworkValidationClient) getNetworkConnectivityStatusCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/registries/{registryId}/endpoints/{endpointKey}/networkConnectivityStatusCollection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNetworkConnectivityStatusCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "NetworkValidation", "GetNetworkConnectivityStatusCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
