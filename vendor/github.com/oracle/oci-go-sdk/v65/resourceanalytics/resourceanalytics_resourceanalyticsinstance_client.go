// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ResourceAnalyticsInstanceClient a client for ResourceAnalyticsInstance
type ResourceAnalyticsInstanceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewResourceAnalyticsInstanceClientWithConfigurationProvider Creates a new default ResourceAnalyticsInstance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewResourceAnalyticsInstanceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ResourceAnalyticsInstanceClient, err error) {
	if enabled := common.CheckForEnabledServices("resourceanalytics"); !enabled {
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
	return newResourceAnalyticsInstanceClientFromBaseClient(baseClient, provider)
}

// NewResourceAnalyticsInstanceClientWithOboToken Creates a new default ResourceAnalyticsInstance client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewResourceAnalyticsInstanceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ResourceAnalyticsInstanceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newResourceAnalyticsInstanceClientFromBaseClient(baseClient, configProvider)
}

func newResourceAnalyticsInstanceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ResourceAnalyticsInstanceClient, err error) {
	// ResourceAnalyticsInstance service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ResourceAnalyticsInstance"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ResourceAnalyticsInstanceClient{BaseClient: baseClient}
	client.BasePath = "20241031"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ResourceAnalyticsInstanceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("resourceanalytics", "https://resource-analytics.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ResourceAnalyticsInstanceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ResourceAnalyticsInstanceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Cancels a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client ResourceAnalyticsInstanceClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeResourceAnalyticsInstanceCompartment Moves a ResourceAnalyticsInstance into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ChangeResourceAnalyticsInstanceCompartment.go.html to see an example of how to use ChangeResourceAnalyticsInstanceCompartment API.
// A default retry strategy applies to this operation ChangeResourceAnalyticsInstanceCompartment()
func (client ResourceAnalyticsInstanceClient) ChangeResourceAnalyticsInstanceCompartment(ctx context.Context, request ChangeResourceAnalyticsInstanceCompartmentRequest) (response ChangeResourceAnalyticsInstanceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeResourceAnalyticsInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeResourceAnalyticsInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeResourceAnalyticsInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeResourceAnalyticsInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeResourceAnalyticsInstanceCompartmentResponse")
	}
	return
}

// changeResourceAnalyticsInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) changeResourceAnalyticsInstanceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resourceAnalyticsInstances/{resourceAnalyticsInstanceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeResourceAnalyticsInstanceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstance/ChangeResourceAnalyticsInstanceCompartment"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "ChangeResourceAnalyticsInstanceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateResourceAnalyticsInstance Creates a ResourceAnalyticsInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/CreateResourceAnalyticsInstance.go.html to see an example of how to use CreateResourceAnalyticsInstance API.
// A default retry strategy applies to this operation CreateResourceAnalyticsInstance()
func (client ResourceAnalyticsInstanceClient) CreateResourceAnalyticsInstance(ctx context.Context, request CreateResourceAnalyticsInstanceRequest) (response CreateResourceAnalyticsInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createResourceAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateResourceAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateResourceAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateResourceAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateResourceAnalyticsInstanceResponse")
	}
	return
}

// createResourceAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) createResourceAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resourceAnalyticsInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateResourceAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstance/CreateResourceAnalyticsInstance"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "CreateResourceAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteResourceAnalyticsInstance Deletes a ResourceAnalyticsInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/DeleteResourceAnalyticsInstance.go.html to see an example of how to use DeleteResourceAnalyticsInstance API.
// A default retry strategy applies to this operation DeleteResourceAnalyticsInstance()
func (client ResourceAnalyticsInstanceClient) DeleteResourceAnalyticsInstance(ctx context.Context, request DeleteResourceAnalyticsInstanceRequest) (response DeleteResourceAnalyticsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteResourceAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteResourceAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteResourceAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteResourceAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteResourceAnalyticsInstanceResponse")
	}
	return
}

// deleteResourceAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) deleteResourceAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/resourceAnalyticsInstances/{resourceAnalyticsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteResourceAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstance/DeleteResourceAnalyticsInstance"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "DeleteResourceAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResourceAnalyticsInstance Gets information about a ResourceAnalyticsInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/GetResourceAnalyticsInstance.go.html to see an example of how to use GetResourceAnalyticsInstance API.
// A default retry strategy applies to this operation GetResourceAnalyticsInstance()
func (client ResourceAnalyticsInstanceClient) GetResourceAnalyticsInstance(ctx context.Context, request GetResourceAnalyticsInstanceRequest) (response GetResourceAnalyticsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResourceAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResourceAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResourceAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResourceAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResourceAnalyticsInstanceResponse")
	}
	return
}

// getResourceAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) getResourceAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceAnalyticsInstances/{resourceAnalyticsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResourceAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstance/GetResourceAnalyticsInstance"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "GetResourceAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ResourceAnalyticsInstanceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client ResourceAnalyticsInstanceClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceAnalyticsInstances Gets a list of ResourceAnalyticsInstances.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListResourceAnalyticsInstances.go.html to see an example of how to use ListResourceAnalyticsInstances API.
// A default retry strategy applies to this operation ListResourceAnalyticsInstances()
func (client ResourceAnalyticsInstanceClient) ListResourceAnalyticsInstances(ctx context.Context, request ListResourceAnalyticsInstancesRequest) (response ListResourceAnalyticsInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceAnalyticsInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceAnalyticsInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceAnalyticsInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceAnalyticsInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceAnalyticsInstancesResponse")
	}
	return
}

// listResourceAnalyticsInstances implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) listResourceAnalyticsInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceAnalyticsInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceAnalyticsInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstanceCollection/ListResourceAnalyticsInstances"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "ListResourceAnalyticsInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists the errors for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ResourceAnalyticsInstanceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client ResourceAnalyticsInstanceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ResourceAnalyticsInstanceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client ResourceAnalyticsInstanceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ResourceAnalyticsInstanceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client ResourceAnalyticsInstanceClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResourceAnalyticsInstanceDisableOac Disable an OAC instance for a ResourceAnalyticsInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ResourceAnalyticsInstanceDisableOac.go.html to see an example of how to use ResourceAnalyticsInstanceDisableOac API.
func (client ResourceAnalyticsInstanceClient) ResourceAnalyticsInstanceDisableOac(ctx context.Context, request ResourceAnalyticsInstanceDisableOacRequest) (response ResourceAnalyticsInstanceDisableOacResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.resourceAnalyticsInstanceDisableOac, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResourceAnalyticsInstanceDisableOacResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResourceAnalyticsInstanceDisableOacResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResourceAnalyticsInstanceDisableOacResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResourceAnalyticsInstanceDisableOacResponse")
	}
	return
}

// resourceAnalyticsInstanceDisableOac implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) resourceAnalyticsInstanceDisableOac(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resourceAnalyticsInstances/{resourceAnalyticsInstanceId}/actions/disableOac", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResourceAnalyticsInstanceDisableOacResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstance/ResourceAnalyticsInstanceDisableOac"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "ResourceAnalyticsInstanceDisableOac", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResourceAnalyticsInstanceEnableOac Attaches an OAC instance to a ResourceAnalyticsInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ResourceAnalyticsInstanceEnableOac.go.html to see an example of how to use ResourceAnalyticsInstanceEnableOac API.
func (client ResourceAnalyticsInstanceClient) ResourceAnalyticsInstanceEnableOac(ctx context.Context, request ResourceAnalyticsInstanceEnableOacRequest) (response ResourceAnalyticsInstanceEnableOacResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.resourceAnalyticsInstanceEnableOac, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResourceAnalyticsInstanceEnableOacResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResourceAnalyticsInstanceEnableOacResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResourceAnalyticsInstanceEnableOacResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResourceAnalyticsInstanceEnableOacResponse")
	}
	return
}

// resourceAnalyticsInstanceEnableOac implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) resourceAnalyticsInstanceEnableOac(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resourceAnalyticsInstances/{resourceAnalyticsInstanceId}/actions/enableOac", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResourceAnalyticsInstanceEnableOacResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstance/ResourceAnalyticsInstanceEnableOac"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "ResourceAnalyticsInstanceEnableOac", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResourceAnalyticsInstance Updates a ResourceAnalyticsInstance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/UpdateResourceAnalyticsInstance.go.html to see an example of how to use UpdateResourceAnalyticsInstance API.
// A default retry strategy applies to this operation UpdateResourceAnalyticsInstance()
func (client ResourceAnalyticsInstanceClient) UpdateResourceAnalyticsInstance(ctx context.Context, request UpdateResourceAnalyticsInstanceRequest) (response UpdateResourceAnalyticsInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResourceAnalyticsInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResourceAnalyticsInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResourceAnalyticsInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResourceAnalyticsInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResourceAnalyticsInstanceResponse")
	}
	return
}

// updateResourceAnalyticsInstance implements the OCIOperation interface (enables retrying operations)
func (client ResourceAnalyticsInstanceClient) updateResourceAnalyticsInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/resourceAnalyticsInstances/{resourceAnalyticsInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateResourceAnalyticsInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/ResourceAnalyticsInstance/UpdateResourceAnalyticsInstance"
		err = common.PostProcessServiceError(err, "ResourceAnalyticsInstance", "UpdateResourceAnalyticsInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
