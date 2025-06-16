// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ApiMetadataClient a client for ApiMetadata
type ApiMetadataClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApiMetadataClientWithConfigurationProvider Creates a new default ApiMetadata client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApiMetadataClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApiMetadataClient, err error) {
	if enabled := common.CheckForEnabledServices("apiaccesscontrol"); !enabled {
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
	return newApiMetadataClientFromBaseClient(baseClient, provider)
}

// NewApiMetadataClientWithOboToken Creates a new default ApiMetadata client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewApiMetadataClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApiMetadataClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newApiMetadataClientFromBaseClient(baseClient, configProvider)
}

func newApiMetadataClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApiMetadataClient, err error) {
	// ApiMetadata service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ApiMetadata"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ApiMetadataClient{BaseClient: baseClient}
	client.BasePath = "20241130"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApiMetadataClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apiaccesscontrol", "https://pactl.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApiMetadataClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApiMetadataClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetApiMetadata Gets information about a ApiMetadata.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/GetApiMetadata.go.html to see an example of how to use GetApiMetadata API.
// A default retry strategy applies to this operation GetApiMetadata()
func (client ApiMetadataClient) GetApiMetadata(ctx context.Context, request GetApiMetadataRequest) (response GetApiMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApiMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApiMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApiMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApiMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApiMetadataResponse")
	}
	return
}

// getApiMetadata implements the OCIOperation interface (enables retrying operations)
func (client ApiMetadataClient) getApiMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apiMetadatas/{apiMetadataId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApiMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ApiMetadata", "GetApiMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApiMetadata Gets a list of ApiMetadata.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ListApiMetadata.go.html to see an example of how to use ListApiMetadata API.
// A default retry strategy applies to this operation ListApiMetadata()
func (client ApiMetadataClient) ListApiMetadata(ctx context.Context, request ListApiMetadataRequest) (response ListApiMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApiMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApiMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApiMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApiMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApiMetadataResponse")
	}
	return
}

// listApiMetadata implements the OCIOperation interface (enables retrying operations)
func (client ApiMetadataClient) listApiMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apiMetadatas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApiMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ApiMetadata", "ListApiMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApiMetadataByEntityTypes Gets a list of ApiMetadata Grouped By Entity Types.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ListApiMetadataByEntityTypes.go.html to see an example of how to use ListApiMetadataByEntityTypes API.
// A default retry strategy applies to this operation ListApiMetadataByEntityTypes()
func (client ApiMetadataClient) ListApiMetadataByEntityTypes(ctx context.Context, request ListApiMetadataByEntityTypesRequest) (response ListApiMetadataByEntityTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApiMetadataByEntityTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApiMetadataByEntityTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApiMetadataByEntityTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApiMetadataByEntityTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApiMetadataByEntityTypesResponse")
	}
	return
}

// listApiMetadataByEntityTypes implements the OCIOperation interface (enables retrying operations)
func (client ApiMetadataClient) listApiMetadataByEntityTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apiMetadatas/byEntityType", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApiMetadataByEntityTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ApiMetadata", "ListApiMetadataByEntityTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
