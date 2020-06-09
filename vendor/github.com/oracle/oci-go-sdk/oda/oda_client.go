// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//OdaClient a client for Oda
type OdaClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOdaClientWithConfigurationProvider Creates a new default Oda client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOdaClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OdaClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newOdaClientFromBaseClient(baseClient, configProvider)
}

// NewOdaClientWithOboToken Creates a new default Oda client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewOdaClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OdaClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newOdaClientFromBaseClient(baseClient, configProvider)
}

func newOdaClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OdaClient, err error) {
	client = OdaClient{BaseClient: baseClient}
	client.BasePath = "20190506"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OdaClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("oda", "https://digitalassistant-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OdaClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OdaClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOdaInstanceCompartment Moves an Digital Assistant instance into a different compartment. When provided, If-Match is checked against
// ETag values of the resource.
func (client OdaClient) ChangeOdaInstanceCompartment(ctx context.Context, request ChangeOdaInstanceCompartmentRequest) (response ChangeOdaInstanceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeOdaInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOdaInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOdaInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOdaInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOdaInstanceCompartmentResponse")
	}
	return
}

// changeOdaInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) changeOdaInstanceCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeOdaInstanceCompartmentResponse
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

// CreateOdaInstance Starts an asynchronous job to create a Digital Assistant instance.
// To monitor the status of the job, take the `opc-work-request-id` response
// header value and use it to call `GET /workRequests/{workRequestID}`.
func (client OdaClient) CreateOdaInstance(ctx context.Context, request CreateOdaInstanceRequest) (response CreateOdaInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createOdaInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOdaInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOdaInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOdaInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOdaInstanceResponse")
	}
	return
}

// createOdaInstance implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) createOdaInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances")
	if err != nil {
		return nil, err
	}

	var response CreateOdaInstanceResponse
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

// DeleteOdaInstance Starts an asynchronous job to delete the specified Digital Assistant instance.
// To monitor the status of the job, take the `opc-work-request-id` response header value and use it to call `GET /workRequests/{workRequestID}`.
func (client OdaClient) DeleteOdaInstance(ctx context.Context, request DeleteOdaInstanceRequest) (response DeleteOdaInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOdaInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOdaInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOdaInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOdaInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOdaInstanceResponse")
	}
	return
}

// deleteOdaInstance implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) deleteOdaInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}")
	if err != nil {
		return nil, err
	}

	var response DeleteOdaInstanceResponse
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

// GetOdaInstance Gets the specified Digital Assistant instance.
func (client OdaClient) GetOdaInstance(ctx context.Context, request GetOdaInstanceRequest) (response GetOdaInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOdaInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOdaInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOdaInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOdaInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOdaInstanceResponse")
	}
	return
}

// getOdaInstance implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) getOdaInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}")
	if err != nil {
		return nil, err
	}

	var response GetOdaInstanceResponse
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

// GetWorkRequest Gets information about the work request with the specified ID, including its status.
// You can use this operation to monitor the status of jobs that you
// requested to create, delete, and update instances.
func (client OdaClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OdaClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
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

// ListOdaInstances Returns a page of Digital Assistant instances that belong to the specified
// compartment.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client OdaClient) ListOdaInstances(ctx context.Context, request ListOdaInstancesRequest) (response ListOdaInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOdaInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOdaInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOdaInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOdaInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOdaInstancesResponse")
	}
	return
}

// listOdaInstances implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) listOdaInstances(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances")
	if err != nil {
		return nil, err
	}

	var response ListOdaInstancesResponse
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

// ListWorkRequestErrors Returns a page of errors for the specified work request.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client OdaClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OdaClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
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

// ListWorkRequestLogs Returns a page of of log messages for a given work request.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client OdaClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OdaClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
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

// ListWorkRequests Returns a page of work requests for the specified compartment.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client OdaClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client OdaClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
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

// StartOdaInstance Starts an inactive Digital Assistant instance. Once active, the instance will be accessible and metering
// of requests will be started again.
func (client OdaClient) StartOdaInstance(ctx context.Context, request StartOdaInstanceRequest) (response StartOdaInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startOdaInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartOdaInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartOdaInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartOdaInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartOdaInstanceResponse")
	}
	return
}

// startOdaInstance implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) startOdaInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/actions/start")
	if err != nil {
		return nil, err
	}

	var response StartOdaInstanceResponse
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

// StopOdaInstance Stops an active Digital Assistant instance. Once inactive, the instance will not be accessible and metering
// of requests will be stopped until the instance is started again. Data associated with the instance
// is not affected.
func (client OdaClient) StopOdaInstance(ctx context.Context, request StopOdaInstanceRequest) (response StopOdaInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.stopOdaInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopOdaInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopOdaInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopOdaInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopOdaInstanceResponse")
	}
	return
}

// stopOdaInstance implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) stopOdaInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/actions/stop")
	if err != nil {
		return nil, err
	}

	var response StopOdaInstanceResponse
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

// UpdateOdaInstance Updates the specified Digital Assistant instance with the information in the request body.
func (client OdaClient) UpdateOdaInstance(ctx context.Context, request UpdateOdaInstanceRequest) (response UpdateOdaInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOdaInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOdaInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOdaInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOdaInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOdaInstanceResponse")
	}
	return
}

// updateOdaInstance implements the OCIOperation interface (enables retrying operations)
func (client OdaClient) updateOdaInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}")
	if err != nil {
		return nil, err
	}

	var response UpdateOdaInstanceResponse
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
