// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// OciCacheDefaultConfigSetClient a client for OciCacheDefaultConfigSet
type OciCacheDefaultConfigSetClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOciCacheDefaultConfigSetClientWithConfigurationProvider Creates a new default OciCacheDefaultConfigSet client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOciCacheDefaultConfigSetClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OciCacheDefaultConfigSetClient, err error) {
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
	return newOciCacheDefaultConfigSetClientFromBaseClient(baseClient, provider)
}

// NewOciCacheDefaultConfigSetClientWithOboToken Creates a new default OciCacheDefaultConfigSet client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOciCacheDefaultConfigSetClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OciCacheDefaultConfigSetClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOciCacheDefaultConfigSetClientFromBaseClient(baseClient, configProvider)
}

func newOciCacheDefaultConfigSetClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OciCacheDefaultConfigSetClient, err error) {
	// OciCacheDefaultConfigSet service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OciCacheDefaultConfigSet"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OciCacheDefaultConfigSetClient{BaseClient: baseClient}
	client.BasePath = "20220315"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OciCacheDefaultConfigSetClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("redis", "https://redis.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OciCacheDefaultConfigSetClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OciCacheDefaultConfigSetClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetOciCacheDefaultConfigSet Retrieves the specified OCI Cache Default Config Set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/GetOciCacheDefaultConfigSet.go.html to see an example of how to use GetOciCacheDefaultConfigSet API.
// A default retry strategy applies to this operation GetOciCacheDefaultConfigSet()
func (client OciCacheDefaultConfigSetClient) GetOciCacheDefaultConfigSet(ctx context.Context, request GetOciCacheDefaultConfigSetRequest) (response GetOciCacheDefaultConfigSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOciCacheDefaultConfigSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOciCacheDefaultConfigSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOciCacheDefaultConfigSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOciCacheDefaultConfigSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOciCacheDefaultConfigSetResponse")
	}
	return
}

// getOciCacheDefaultConfigSet implements the OCIOperation interface (enables retrying operations)
func (client OciCacheDefaultConfigSetClient) getOciCacheDefaultConfigSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheDefaultConfigSets/{ociCacheDefaultConfigSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOciCacheDefaultConfigSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheDefaultConfigSet/GetOciCacheDefaultConfigSet"
		err = common.PostProcessServiceError(err, "OciCacheDefaultConfigSet", "GetOciCacheDefaultConfigSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOciCacheDefaultConfigSets Lists the OCI Cache Default Config Sets in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheDefaultConfigSets.go.html to see an example of how to use ListOciCacheDefaultConfigSets API.
// A default retry strategy applies to this operation ListOciCacheDefaultConfigSets()
func (client OciCacheDefaultConfigSetClient) ListOciCacheDefaultConfigSets(ctx context.Context, request ListOciCacheDefaultConfigSetsRequest) (response ListOciCacheDefaultConfigSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOciCacheDefaultConfigSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOciCacheDefaultConfigSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOciCacheDefaultConfigSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOciCacheDefaultConfigSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOciCacheDefaultConfigSetsResponse")
	}
	return
}

// listOciCacheDefaultConfigSets implements the OCIOperation interface (enables retrying operations)
func (client OciCacheDefaultConfigSetClient) listOciCacheDefaultConfigSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheDefaultConfigSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOciCacheDefaultConfigSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheDefaultConfigSetSummary/ListOciCacheDefaultConfigSets"
		err = common.PostProcessServiceError(err, "OciCacheDefaultConfigSet", "ListOciCacheDefaultConfigSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
