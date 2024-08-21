// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// AnalyticsClient a client for Analytics
type AnalyticsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAnalyticsClientWithConfigurationProvider Creates a new default Analytics client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAnalyticsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AnalyticsClient, err error) {
	if enabled := common.CheckForEnabledServices("analytics"); !enabled {
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
	return newAnalyticsClientFromBaseClient(baseClient, provider)
}

// NewAnalyticsClientWithOboToken Creates a new default Analytics client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAnalyticsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AnalyticsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAnalyticsClientFromBaseClient(baseClient, configProvider)
}

func newAnalyticsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AnalyticsClient, err error) {
	// Analytics service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Analytics"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AnalyticsClient{BaseClient: baseClient}
	client.BasePath = "20190331"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AnalyticsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("analytics", "https://analytics.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AnalyticsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AnalyticsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAnalyticsInstanceCompartment Change the compartment of an Analytics instance. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ChangeAnalyticsInstanceCompartment.go.html to see an example of how to use ChangeAnalyticsInstanceCompartment API.
// A default retry strategy applies to this operation ChangeAnalyticsInstanceCompartment()
func (client AnalyticsClient) ChangeAnalyticsInstanceCompartment(ctx context.Context, request ChangeAnalyticsInstanceCompartmentRequest) (response ChangeAnalyticsInstanceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAnalyticsInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAnalyticsInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAnalyticsInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAnalyticsInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAnalyticsInstanceCompartmentResponse")
	}
	return
}

// changeAnalyticsInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) changeAnalyticsInstanceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAnalyticsInstanceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/ChangeAnalyticsInstanceCompartment"
		err = common.PostProcessServiceError(err, "Analytics", "ChangeAnalyticsInstanceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAnalyticsInstanceNetworkEndpoint Change an Analytics instance network endpoint. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ChangeAnalyticsInstanceNetworkEndpoint.go.html to see an example of how to use ChangeAnalyticsInstanceNetworkEndpoint API.
// A default retry strategy applies to this operation ChangeAnalyticsInstanceNetworkEndpoint()
func (client AnalyticsClient) ChangeAnalyticsInstanceNetworkEndpoint(ctx context.Context, request ChangeAnalyticsInstanceNetworkEndpointRequest) (response ChangeAnalyticsInstanceNetworkEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAnalyticsInstanceNetworkEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAnalyticsInstanceNetworkEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAnalyticsInstanceNetworkEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAnalyticsInstanceNetworkEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAnalyticsInstanceNetworkEndpointResponse")
	}
	return
}

// changeAnalyticsInstanceNetworkEndpoint implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) changeAnalyticsInstanceNetworkEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/actions/changeNetworkEndpoint", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAnalyticsInstanceNetworkEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/ChangeAnalyticsInstanceNetworkEndpoint"
		err = common.PostProcessServiceError(err, "Analytics", "ChangeAnalyticsInstanceNetworkEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAnalyticsInstance Create a new AnalyticsInstance in the specified compartment. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/CreateAnalyticsInstance.go.html to see an example of how to use CreateAnalyticsInstance API.
// A default retry strategy applies to this operation CreateAnalyticsInstance()
func (client AnalyticsClient) CreateAnalyticsInstance(ctx context.Context, request CreateAnalyticsInstanceRequest) (response CreateAnalyticsInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAnalyticsInstanceResponse")
	}
	return
}

// createAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) createAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Analytics", "CreateAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePrivateAccessChannel Create an Private access Channel for the Analytics instance. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/CreatePrivateAccessChannel.go.html to see an example of how to use CreatePrivateAccessChannel API.
// A default retry strategy applies to this operation CreatePrivateAccessChannel()
func (client AnalyticsClient) CreatePrivateAccessChannel(ctx context.Context, request CreatePrivateAccessChannelRequest) (response CreatePrivateAccessChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPrivateAccessChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePrivateAccessChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePrivateAccessChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePrivateAccessChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePrivateAccessChannelResponse")
	}
	return
}

// createPrivateAccessChannel implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) createPrivateAccessChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/privateAccessChannels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePrivateAccessChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/CreatePrivateAccessChannel"
		err = common.PostProcessServiceError(err, "Analytics", "CreatePrivateAccessChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVanityUrl Allows specifying a custom host name to be used to access the analytics instance.  This requires prior setup of DNS entry and certificate
// for this host.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/CreateVanityUrl.go.html to see an example of how to use CreateVanityUrl API.
// A default retry strategy applies to this operation CreateVanityUrl()
func (client AnalyticsClient) CreateVanityUrl(ctx context.Context, request CreateVanityUrlRequest) (response CreateVanityUrlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVanityUrl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVanityUrlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVanityUrlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVanityUrlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVanityUrlResponse")
	}
	return
}

// createVanityUrl implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) createVanityUrl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/vanityUrls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVanityUrlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/CreateVanityUrl"
		err = common.PostProcessServiceError(err, "Analytics", "CreateVanityUrl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAnalyticsInstance Terminates the specified Analytics instance. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/DeleteAnalyticsInstance.go.html to see an example of how to use DeleteAnalyticsInstance API.
// A default retry strategy applies to this operation DeleteAnalyticsInstance()
func (client AnalyticsClient) DeleteAnalyticsInstance(ctx context.Context, request DeleteAnalyticsInstanceRequest) (response DeleteAnalyticsInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAnalyticsInstanceResponse")
	}
	return
}

// deleteAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) deleteAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/analyticsInstances/{analyticsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/DeleteAnalyticsInstance"
		err = common.PostProcessServiceError(err, "Analytics", "DeleteAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePrivateAccessChannel Delete an Analytics instance's Private access channel with the given unique identifier key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/DeletePrivateAccessChannel.go.html to see an example of how to use DeletePrivateAccessChannel API.
// A default retry strategy applies to this operation DeletePrivateAccessChannel()
func (client AnalyticsClient) DeletePrivateAccessChannel(ctx context.Context, request DeletePrivateAccessChannelRequest) (response DeletePrivateAccessChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deletePrivateAccessChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePrivateAccessChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePrivateAccessChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePrivateAccessChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePrivateAccessChannelResponse")
	}
	return
}

// deletePrivateAccessChannel implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) deletePrivateAccessChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/analyticsInstances/{analyticsInstanceId}/privateAccessChannels/{privateAccessChannelKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePrivateAccessChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/DeletePrivateAccessChannel"
		err = common.PostProcessServiceError(err, "Analytics", "DeletePrivateAccessChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVanityUrl Allows deleting a previously created vanity url.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/DeleteVanityUrl.go.html to see an example of how to use DeleteVanityUrl API.
// A default retry strategy applies to this operation DeleteVanityUrl()
func (client AnalyticsClient) DeleteVanityUrl(ctx context.Context, request DeleteVanityUrlRequest) (response DeleteVanityUrlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteVanityUrl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVanityUrlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVanityUrlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVanityUrlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVanityUrlResponse")
	}
	return
}

// deleteVanityUrl implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) deleteVanityUrl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/analyticsInstances/{analyticsInstanceId}/vanityUrls/{vanityUrlKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVanityUrlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/DeleteVanityUrl"
		err = common.PostProcessServiceError(err, "Analytics", "DeleteVanityUrl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWorkRequest Cancel a work request that has not started yet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/DeleteWorkRequest.go.html to see an example of how to use DeleteWorkRequest API.
func (client AnalyticsClient) DeleteWorkRequest(ctx context.Context, request DeleteWorkRequestRequest) (response DeleteWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWorkRequestResponse")
	}
	return
}

// deleteWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) deleteWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/WorkRequest/DeleteWorkRequest"
		err = common.PostProcessServiceError(err, "Analytics", "DeleteWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAnalyticsInstance Info for a specific Analytics instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/GetAnalyticsInstance.go.html to see an example of how to use GetAnalyticsInstance API.
// A default retry strategy applies to this operation GetAnalyticsInstance()
func (client AnalyticsClient) GetAnalyticsInstance(ctx context.Context, request GetAnalyticsInstanceRequest) (response GetAnalyticsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAnalyticsInstanceResponse")
	}
	return
}

// getAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) getAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/analyticsInstances/{analyticsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/GetAnalyticsInstance"
		err = common.PostProcessServiceError(err, "Analytics", "GetAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivateAccessChannel Retrieve private access channel in the specified Analytics Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/GetPrivateAccessChannel.go.html to see an example of how to use GetPrivateAccessChannel API.
// A default retry strategy applies to this operation GetPrivateAccessChannel()
func (client AnalyticsClient) GetPrivateAccessChannel(ctx context.Context, request GetPrivateAccessChannelRequest) (response GetPrivateAccessChannelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateAccessChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateAccessChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateAccessChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateAccessChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateAccessChannelResponse")
	}
	return
}

// getPrivateAccessChannel implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) getPrivateAccessChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/analyticsInstances/{analyticsInstanceId}/privateAccessChannels/{privateAccessChannelKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivateAccessChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/GetPrivateAccessChannel"
		err = common.PostProcessServiceError(err, "Analytics", "GetPrivateAccessChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Get the details of a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client AnalyticsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "Analytics", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAnalyticsInstances List Analytics instances.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ListAnalyticsInstances.go.html to see an example of how to use ListAnalyticsInstances API.
// A default retry strategy applies to this operation ListAnalyticsInstances()
func (client AnalyticsClient) ListAnalyticsInstances(ctx context.Context, request ListAnalyticsInstancesRequest) (response ListAnalyticsInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAnalyticsInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAnalyticsInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAnalyticsInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAnalyticsInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAnalyticsInstancesResponse")
	}
	return
}

// listAnalyticsInstances implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) listAnalyticsInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/analyticsInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAnalyticsInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstanceSummary/ListAnalyticsInstances"
		err = common.PostProcessServiceError(err, "Analytics", "ListAnalyticsInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Get the errors of a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client AnalyticsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Analytics", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Get the logs of a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client AnalyticsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/WorkRequestLog/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Analytics", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests List all work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client AnalyticsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/WorkRequestSummary/ListWorkRequests"
		err = common.PostProcessServiceError(err, "Analytics", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScaleAnalyticsInstance Scale an Analytics instance up or down. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/ScaleAnalyticsInstance.go.html to see an example of how to use ScaleAnalyticsInstance API.
// A default retry strategy applies to this operation ScaleAnalyticsInstance()
func (client AnalyticsClient) ScaleAnalyticsInstance(ctx context.Context, request ScaleAnalyticsInstanceRequest) (response ScaleAnalyticsInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scaleAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScaleAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScaleAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScaleAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScaleAnalyticsInstanceResponse")
	}
	return
}

// scaleAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) scaleAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/actions/scale", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScaleAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/ScaleAnalyticsInstance"
		err = common.PostProcessServiceError(err, "Analytics", "ScaleAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetFeatureBundle Set the external service properties of an Analytics instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/SetFeatureBundle.go.html to see an example of how to use SetFeatureBundle API.
// A default retry strategy applies to this operation SetFeatureBundle()
func (client AnalyticsClient) SetFeatureBundle(ctx context.Context, request SetFeatureBundleRequest) (response SetFeatureBundleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.setFeatureBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetFeatureBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetFeatureBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetFeatureBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetFeatureBundleResponse")
	}
	return
}

// setFeatureBundle implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) setFeatureBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/actions/setFeatureBundle", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetFeatureBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/SetFeatureBundle"
		err = common.PostProcessServiceError(err, "Analytics", "SetFeatureBundle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetKmsKey Encrypts the customer data of this Analytics instance using either a customer OCI Vault Key or Oracle managed default key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/SetKmsKey.go.html to see an example of how to use SetKmsKey API.
// A default retry strategy applies to this operation SetKmsKey()
func (client AnalyticsClient) SetKmsKey(ctx context.Context, request SetKmsKeyRequest) (response SetKmsKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.setKmsKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetKmsKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetKmsKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetKmsKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetKmsKeyResponse")
	}
	return
}

// setKmsKey implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) setKmsKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/actions/setKmsKey", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetKmsKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/SetKmsKey"
		err = common.PostProcessServiceError(err, "Analytics", "SetKmsKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartAnalyticsInstance Starts the specified Analytics instance. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/StartAnalyticsInstance.go.html to see an example of how to use StartAnalyticsInstance API.
// A default retry strategy applies to this operation StartAnalyticsInstance()
func (client AnalyticsClient) StartAnalyticsInstance(ctx context.Context, request StartAnalyticsInstanceRequest) (response StartAnalyticsInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartAnalyticsInstanceResponse")
	}
	return
}

// startAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) startAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/StartAnalyticsInstance"
		err = common.PostProcessServiceError(err, "Analytics", "StartAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopAnalyticsInstance Stop the specified Analytics instance. The operation is long-running
// and creates a new WorkRequest.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/StopAnalyticsInstance.go.html to see an example of how to use StopAnalyticsInstance API.
// A default retry strategy applies to this operation StopAnalyticsInstance()
func (client AnalyticsClient) StopAnalyticsInstance(ctx context.Context, request StopAnalyticsInstanceRequest) (response StopAnalyticsInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopAnalyticsInstanceResponse")
	}
	return
}

// stopAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) stopAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/analyticsInstances/{analyticsInstanceId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/StopAnalyticsInstance"
		err = common.PostProcessServiceError(err, "Analytics", "StopAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAnalyticsInstance Updates certain fields of an Analytics instance. Fields that are not provided in the
// request will not be updated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/UpdateAnalyticsInstance.go.html to see an example of how to use UpdateAnalyticsInstance API.
func (client AnalyticsClient) UpdateAnalyticsInstance(ctx context.Context, request UpdateAnalyticsInstanceRequest) (response UpdateAnalyticsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAnalyticsInstanceResponse")
	}
	return
}

// updateAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) updateAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/analyticsInstances/{analyticsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/UpdateAnalyticsInstance"
		err = common.PostProcessServiceError(err, "Analytics", "UpdateAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePrivateAccessChannel Update the Private Access Channel with the given unique identifier key in the specified Analytics Instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/UpdatePrivateAccessChannel.go.html to see an example of how to use UpdatePrivateAccessChannel API.
// A default retry strategy applies to this operation UpdatePrivateAccessChannel()
func (client AnalyticsClient) UpdatePrivateAccessChannel(ctx context.Context, request UpdatePrivateAccessChannelRequest) (response UpdatePrivateAccessChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updatePrivateAccessChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePrivateAccessChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePrivateAccessChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePrivateAccessChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePrivateAccessChannelResponse")
	}
	return
}

// updatePrivateAccessChannel implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) updatePrivateAccessChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/analyticsInstances/{analyticsInstanceId}/privateAccessChannels/{privateAccessChannelKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePrivateAccessChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/UpdatePrivateAccessChannel"
		err = common.PostProcessServiceError(err, "Analytics", "UpdatePrivateAccessChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVanityUrl Allows uploading a new certificate for a vanity url, which will have to be done when the current certificate is expiring.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/analytics/UpdateVanityUrl.go.html to see an example of how to use UpdateVanityUrl API.
// A default retry strategy applies to this operation UpdateVanityUrl()
func (client AnalyticsClient) UpdateVanityUrl(ctx context.Context, request UpdateVanityUrlRequest) (response UpdateVanityUrlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateVanityUrl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVanityUrlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVanityUrlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVanityUrlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVanityUrlResponse")
	}
	return
}

// updateVanityUrl implements the OCIOperation interface (enables retrying operations)
func (client AnalyticsClient) updateVanityUrl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/analyticsInstances/{analyticsInstanceId}/vanityUrls/{vanityUrlKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVanityUrlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/analytics/20190331/AnalyticsInstance/UpdateVanityUrl"
		err = common.PostProcessServiceError(err, "Analytics", "UpdateVanityUrl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
