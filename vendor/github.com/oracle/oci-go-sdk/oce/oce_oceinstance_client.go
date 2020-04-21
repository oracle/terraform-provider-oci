// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//OceInstanceClient a client for OceInstance
type OceInstanceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOceInstanceClientWithConfigurationProvider Creates a new default OceInstance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOceInstanceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OceInstanceClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newOceInstanceClientFromBaseClient(baseClient, configProvider)
}

// NewOceInstanceClientWithOboToken Creates a new default OceInstance client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewOceInstanceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OceInstanceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newOceInstanceClientFromBaseClient(baseClient, configProvider)
}

func newOceInstanceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OceInstanceClient, err error) {
	client = OceInstanceClient{BaseClient: baseClient}
	client.BasePath = "20190912"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OceInstanceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("oce", "https://cp.oce.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OceInstanceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OceInstanceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOceInstanceCompartment Moves a OceInstance into a different compartment
func (client OceInstanceClient) ChangeOceInstanceCompartment(ctx context.Context, request ChangeOceInstanceCompartmentRequest) (response ChangeOceInstanceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeOceInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOceInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOceInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOceInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOceInstanceCompartmentResponse")
	}
	return
}

// changeOceInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client OceInstanceClient) changeOceInstanceCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oceInstances/{oceInstanceId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeOceInstanceCompartmentResponse
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

// CreateOceInstance Creates a new OceInstance.
func (client OceInstanceClient) CreateOceInstance(ctx context.Context, request CreateOceInstanceRequest) (response CreateOceInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createOceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOceInstanceResponse")
	}
	return
}

// createOceInstance implements the OCIOperation interface (enables retrying operations)
func (client OceInstanceClient) createOceInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/oceInstances")
	if err != nil {
		return nil, err
	}

	var response CreateOceInstanceResponse
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

// DeleteOceInstance Deletes a OceInstance resource by identifier
func (client OceInstanceClient) DeleteOceInstance(ctx context.Context, request DeleteOceInstanceRequest) (response DeleteOceInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOceInstanceResponse")
	}
	return
}

// deleteOceInstance implements the OCIOperation interface (enables retrying operations)
func (client OceInstanceClient) deleteOceInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/oceInstances/{oceInstanceId}")
	if err != nil {
		return nil, err
	}

	var response DeleteOceInstanceResponse
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

// GetOceInstance Gets a OceInstance by identifier
func (client OceInstanceClient) GetOceInstance(ctx context.Context, request GetOceInstanceRequest) (response GetOceInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOceInstanceResponse")
	}
	return
}

// getOceInstance implements the OCIOperation interface (enables retrying operations)
func (client OceInstanceClient) getOceInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oceInstances/{oceInstanceId}")
	if err != nil {
		return nil, err
	}

	var response GetOceInstanceResponse
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

// GetWorkRequest Gets the status of the work request with the given ID.
func (client OceInstanceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client OceInstanceClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListOceInstances Returns a list of OceInstances.
func (client OceInstanceClient) ListOceInstances(ctx context.Context, request ListOceInstancesRequest) (response ListOceInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOceInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOceInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOceInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOceInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOceInstancesResponse")
	}
	return
}

// listOceInstances implements the OCIOperation interface (enables retrying operations)
func (client OceInstanceClient) listOceInstances(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/oceInstances")
	if err != nil {
		return nil, err
	}

	var response ListOceInstancesResponse
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
func (client OceInstanceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client OceInstanceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
func (client OceInstanceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client OceInstanceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListWorkRequests Lists the work requests in a compartment.
func (client OceInstanceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client OceInstanceClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// UpdateOceInstance Updates the OceInstance
func (client OceInstanceClient) UpdateOceInstance(ctx context.Context, request UpdateOceInstanceRequest) (response UpdateOceInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOceInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOceInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOceInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOceInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOceInstanceResponse")
	}
	return
}

// updateOceInstance implements the OCIOperation interface (enables retrying operations)
func (client OceInstanceClient) updateOceInstance(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/oceInstances/{oceInstanceId}")
	if err != nil {
		return nil, err
	}

	var response UpdateOceInstanceResponse
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
