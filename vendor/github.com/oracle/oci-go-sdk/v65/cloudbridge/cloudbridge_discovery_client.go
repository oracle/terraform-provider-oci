// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DiscoveryClient a client for Discovery
type DiscoveryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDiscoveryClientWithConfigurationProvider Creates a new default Discovery client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDiscoveryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DiscoveryClient, err error) {
	if enabled := common.CheckForEnabledServices("cloudbridge"); !enabled {
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
	return newDiscoveryClientFromBaseClient(baseClient, provider)
}

// NewDiscoveryClientWithOboToken Creates a new default Discovery client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDiscoveryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DiscoveryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDiscoveryClientFromBaseClient(baseClient, configProvider)
}

func newDiscoveryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DiscoveryClient, err error) {
	// Discovery service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Discovery"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DiscoveryClient{BaseClient: baseClient}
	client.BasePath = "20220509"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DiscoveryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cloudbridge", "https://cloudbridge.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DiscoveryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DiscoveryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAssetSourceCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ChangeAssetSourceCompartment.go.html to see an example of how to use ChangeAssetSourceCompartment API.
// A default retry strategy applies to this operation ChangeAssetSourceCompartment()
func (client DiscoveryClient) ChangeAssetSourceCompartment(ctx context.Context, request ChangeAssetSourceCompartmentRequest) (response ChangeAssetSourceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAssetSourceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAssetSourceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAssetSourceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAssetSourceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAssetSourceCompartmentResponse")
	}
	return
}

// changeAssetSourceCompartment implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) changeAssetSourceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assetSources/{assetSourceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAssetSourceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/ChangeAssetSourceCompartment"
		err = common.PostProcessServiceError(err, "Discovery", "ChangeAssetSourceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDiscoveryScheduleCompartment Moves the specified discovery schedule into a different compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ChangeDiscoveryScheduleCompartment.go.html to see an example of how to use ChangeDiscoveryScheduleCompartment API.
// A default retry strategy applies to this operation ChangeDiscoveryScheduleCompartment()
func (client DiscoveryClient) ChangeDiscoveryScheduleCompartment(ctx context.Context, request ChangeDiscoveryScheduleCompartmentRequest) (response ChangeDiscoveryScheduleCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDiscoveryScheduleCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDiscoveryScheduleCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDiscoveryScheduleCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDiscoveryScheduleCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDiscoveryScheduleCompartmentResponse")
	}
	return
}

// changeDiscoveryScheduleCompartment implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) changeDiscoveryScheduleCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/discoverySchedules/{discoveryScheduleId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDiscoveryScheduleCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/DiscoverySchedule/ChangeDiscoveryScheduleCompartment"
		err = common.PostProcessServiceError(err, "Discovery", "ChangeDiscoveryScheduleCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAssetSource Creates an asset source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/CreateAssetSource.go.html to see an example of how to use CreateAssetSource API.
// A default retry strategy applies to this operation CreateAssetSource()
func (client DiscoveryClient) CreateAssetSource(ctx context.Context, request CreateAssetSourceRequest) (response CreateAssetSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAssetSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAssetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAssetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAssetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAssetSourceResponse")
	}
	return
}

// createAssetSource implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) createAssetSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assetSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAssetSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/CreateAssetSource"
		err = common.PostProcessServiceError(err, "Discovery", "CreateAssetSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assetsource{})
	return response, err
}

// CreateDiscoverySchedule Creates the discovery schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/CreateDiscoverySchedule.go.html to see an example of how to use CreateDiscoverySchedule API.
// A default retry strategy applies to this operation CreateDiscoverySchedule()
func (client DiscoveryClient) CreateDiscoverySchedule(ctx context.Context, request CreateDiscoveryScheduleRequest) (response CreateDiscoveryScheduleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDiscoverySchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDiscoveryScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDiscoveryScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDiscoveryScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDiscoveryScheduleResponse")
	}
	return
}

// createDiscoverySchedule implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) createDiscoverySchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/discoverySchedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDiscoveryScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/DiscoverySchedule/CreateDiscoverySchedule"
		err = common.PostProcessServiceError(err, "Discovery", "CreateDiscoverySchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAssetSource Deletes the asset source by ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/DeleteAssetSource.go.html to see an example of how to use DeleteAssetSource API.
// A default retry strategy applies to this operation DeleteAssetSource()
func (client DiscoveryClient) DeleteAssetSource(ctx context.Context, request DeleteAssetSourceRequest) (response DeleteAssetSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAssetSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAssetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAssetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAssetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAssetSourceResponse")
	}
	return
}

// deleteAssetSource implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) deleteAssetSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/assetSources/{assetSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAssetSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/DeleteAssetSource"
		err = common.PostProcessServiceError(err, "Discovery", "DeleteAssetSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDiscoverySchedule Deletes the specified discovery schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/DeleteDiscoverySchedule.go.html to see an example of how to use DeleteDiscoverySchedule API.
// A default retry strategy applies to this operation DeleteDiscoverySchedule()
func (client DiscoveryClient) DeleteDiscoverySchedule(ctx context.Context, request DeleteDiscoveryScheduleRequest) (response DeleteDiscoveryScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDiscoverySchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDiscoveryScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDiscoveryScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDiscoveryScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDiscoveryScheduleResponse")
	}
	return
}

// deleteDiscoverySchedule implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) deleteDiscoverySchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/discoverySchedules/{discoveryScheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDiscoveryScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/DiscoverySchedule/DeleteDiscoverySchedule"
		err = common.PostProcessServiceError(err, "Discovery", "DeleteDiscoverySchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssetSource Gets the asset source by ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetAssetSource.go.html to see an example of how to use GetAssetSource API.
// A default retry strategy applies to this operation GetAssetSource()
func (client DiscoveryClient) GetAssetSource(ctx context.Context, request GetAssetSourceRequest) (response GetAssetSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssetSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssetSourceResponse")
	}
	return
}

// getAssetSource implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) getAssetSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assetSources/{assetSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssetSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/GetAssetSource"
		err = common.PostProcessServiceError(err, "Discovery", "GetAssetSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &assetsource{})
	return response, err
}

// GetDiscoverySchedule Reads information about the specified discovery schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetDiscoverySchedule.go.html to see an example of how to use GetDiscoverySchedule API.
// A default retry strategy applies to this operation GetDiscoverySchedule()
func (client DiscoveryClient) GetDiscoverySchedule(ctx context.Context, request GetDiscoveryScheduleRequest) (response GetDiscoveryScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDiscoverySchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDiscoveryScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDiscoveryScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDiscoveryScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDiscoveryScheduleResponse")
	}
	return
}

// getDiscoverySchedule implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) getDiscoverySchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoverySchedules/{discoveryScheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDiscoveryScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/DiscoverySchedule/GetDiscoverySchedule"
		err = common.PostProcessServiceError(err, "Discovery", "GetDiscoverySchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssetSourceConnections Gets known connections to the asset source by the asset source ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAssetSourceConnections.go.html to see an example of how to use ListAssetSourceConnections API.
// A default retry strategy applies to this operation ListAssetSourceConnections()
func (client DiscoveryClient) ListAssetSourceConnections(ctx context.Context, request ListAssetSourceConnectionsRequest) (response ListAssetSourceConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssetSourceConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssetSourceConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssetSourceConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssetSourceConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssetSourceConnectionsResponse")
	}
	return
}

// listAssetSourceConnections implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) listAssetSourceConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assetSources/{assetSourceId}/actions/listConnections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssetSourceConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/ListAssetSourceConnections"
		err = common.PostProcessServiceError(err, "Discovery", "ListAssetSourceConnections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssetSources Returns a list of asset sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAssetSources.go.html to see an example of how to use ListAssetSources API.
// A default retry strategy applies to this operation ListAssetSources()
func (client DiscoveryClient) ListAssetSources(ctx context.Context, request ListAssetSourcesRequest) (response ListAssetSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssetSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssetSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssetSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssetSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssetSourcesResponse")
	}
	return
}

// listAssetSources implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) listAssetSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assetSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssetSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/ListAssetSources"
		err = common.PostProcessServiceError(err, "Discovery", "ListAssetSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDiscoverySchedules Lists discovery schedules.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListDiscoverySchedules.go.html to see an example of how to use ListDiscoverySchedules API.
// A default retry strategy applies to this operation ListDiscoverySchedules()
func (client DiscoveryClient) ListDiscoverySchedules(ctx context.Context, request ListDiscoverySchedulesRequest) (response ListDiscoverySchedulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDiscoverySchedules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDiscoverySchedulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDiscoverySchedulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDiscoverySchedulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDiscoverySchedulesResponse")
	}
	return
}

// listDiscoverySchedules implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) listDiscoverySchedules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoverySchedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDiscoverySchedulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/DiscoverySchedule/ListDiscoverySchedules"
		err = common.PostProcessServiceError(err, "Discovery", "ListDiscoverySchedules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshAssetSource Initiates the process of asset metadata synchronization with the related asset source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/RefreshAssetSource.go.html to see an example of how to use RefreshAssetSource API.
// A default retry strategy applies to this operation RefreshAssetSource()
func (client DiscoveryClient) RefreshAssetSource(ctx context.Context, request RefreshAssetSourceRequest) (response RefreshAssetSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshAssetSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshAssetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshAssetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshAssetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshAssetSourceResponse")
	}
	return
}

// refreshAssetSource implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) refreshAssetSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assetSources/{assetSourceId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshAssetSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/RefreshAssetSource"
		err = common.PostProcessServiceError(err, "Discovery", "RefreshAssetSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAssetSource Updates the asset source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdateAssetSource.go.html to see an example of how to use UpdateAssetSource API.
// A default retry strategy applies to this operation UpdateAssetSource()
func (client DiscoveryClient) UpdateAssetSource(ctx context.Context, request UpdateAssetSourceRequest) (response UpdateAssetSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAssetSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAssetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAssetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAssetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAssetSourceResponse")
	}
	return
}

// updateAssetSource implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) updateAssetSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/assetSources/{assetSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAssetSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetSource/UpdateAssetSource"
		err = common.PostProcessServiceError(err, "Discovery", "UpdateAssetSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDiscoverySchedule Updates the specified discovery schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdateDiscoverySchedule.go.html to see an example of how to use UpdateDiscoverySchedule API.
// A default retry strategy applies to this operation UpdateDiscoverySchedule()
func (client DiscoveryClient) UpdateDiscoverySchedule(ctx context.Context, request UpdateDiscoveryScheduleRequest) (response UpdateDiscoveryScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDiscoverySchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDiscoveryScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDiscoveryScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDiscoveryScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDiscoveryScheduleResponse")
	}
	return
}

// updateDiscoverySchedule implements the OCIOperation interface (enables retrying operations)
func (client DiscoveryClient) updateDiscoverySchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/discoverySchedules/{discoveryScheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDiscoveryScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/DiscoverySchedule/UpdateDiscoverySchedule"
		err = common.PostProcessServiceError(err, "Discovery", "UpdateDiscoverySchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
