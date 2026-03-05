// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OciCacheEngineOptionsClient a client for OciCacheEngineOptions
type OciCacheEngineOptionsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOciCacheEngineOptionsClientWithConfigurationProvider Creates a new default OciCacheEngineOptions client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOciCacheEngineOptionsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OciCacheEngineOptionsClient, err error) {
	if enabled := common.CheckForEnabledServices("redis"); !enabled {
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
	return newOciCacheEngineOptionsClientFromBaseClient(baseClient, provider)
}

// NewOciCacheEngineOptionsClientWithOboToken Creates a new default OciCacheEngineOptions client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOciCacheEngineOptionsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OciCacheEngineOptionsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOciCacheEngineOptionsClientFromBaseClient(baseClient, configProvider)
}

func newOciCacheEngineOptionsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OciCacheEngineOptionsClient, err error) {
	// OciCacheEngineOptions service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OciCacheEngineOptions"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OciCacheEngineOptionsClient{BaseClient: baseClient}
	client.BasePath = "20220315"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OciCacheEngineOptionsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("redis", "https://redis.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OciCacheEngineOptionsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OciCacheEngineOptionsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListOciCacheEngineOptions Lists OCI Cache Engine options
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheEngineOptions.go.html to see an example of how to use ListOciCacheEngineOptions API.
// A default retry strategy applies to this operation ListOciCacheEngineOptions()
func (client OciCacheEngineOptionsClient) ListOciCacheEngineOptions(ctx context.Context, request ListOciCacheEngineOptionsRequest) (response ListOciCacheEngineOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOciCacheEngineOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOciCacheEngineOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOciCacheEngineOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOciCacheEngineOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOciCacheEngineOptionsResponse")
	}
	return
}

// listOciCacheEngineOptions implements the OCIOperation interface (enables retrying operations)
func (client OciCacheEngineOptionsClient) listOciCacheEngineOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheEngineOptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOciCacheEngineOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheEngineOptionSummary/ListOciCacheEngineOptions"
		err = common.PostProcessServiceError(err, "OciCacheEngineOptions", "ListOciCacheEngineOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
