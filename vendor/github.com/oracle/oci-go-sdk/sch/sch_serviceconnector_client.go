// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
//

package sch

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ServiceConnectorClient a client for ServiceConnector
type ServiceConnectorClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewServiceConnectorClientWithConfigurationProvider Creates a new default ServiceConnector client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewServiceConnectorClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ServiceConnectorClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newServiceConnectorClientFromBaseClient(baseClient, configProvider)
}

// NewServiceConnectorClientWithOboToken Creates a new default ServiceConnector client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewServiceConnectorClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ServiceConnectorClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newServiceConnectorClientFromBaseClient(baseClient, configProvider)
}

func newServiceConnectorClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ServiceConnectorClient, err error) {
	client = ServiceConnectorClient{BaseClient: baseClient}
	client.BasePath = "20200909"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ServiceConnectorClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("sch", "https://service-connector-hub.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ServiceConnectorClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ServiceConnectorClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateServiceConnector Activates the specified service connector.
// After you send your request, the service connector's state is temporarily
// UPDATING. When the state changes to ACTIVE, data begins transferring from the
// source service to the target service. For instructions on deactivating and
// activating service connectors, see
// To activate or deactivate a service connector (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
func (client ServiceConnectorClient) ActivateServiceConnector(ctx context.Context, request ActivateServiceConnectorRequest) (response ActivateServiceConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.activateServiceConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateServiceConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateServiceConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateServiceConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateServiceConnectorResponse")
	}
	return
}

// activateServiceConnector implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) activateServiceConnector(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/serviceConnectors/{serviceConnectorId}/actions/activate")
	if err != nil {
		return nil, err
	}

	var response ActivateServiceConnectorResponse
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

// ChangeServiceConnectorCompartment Moves a service connector into a different compartment within the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When provided, If-Match is checked against ETag values of the resource.
func (client ServiceConnectorClient) ChangeServiceConnectorCompartment(ctx context.Context, request ChangeServiceConnectorCompartmentRequest) (response ChangeServiceConnectorCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeServiceConnectorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeServiceConnectorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeServiceConnectorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeServiceConnectorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeServiceConnectorCompartmentResponse")
	}
	return
}

// changeServiceConnectorCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) changeServiceConnectorCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/serviceConnectors/{serviceConnectorId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeServiceConnectorCompartmentResponse
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

// CreateServiceConnector Creates a new service connector in the specified compartment.
// A service connector is a logically defined flow for moving data from
// a source service to a destination service in Oracle Cloud Infrastructure.
// For general information about service connectors, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
// For purposes of access control, you must provide the
// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where
// you want the service connector to reside. Notice that the service connector
// doesn't have to be in the same compartment as the source or target services.
// For information about access control and compartments, see
// Overview of the IAM Service (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
// After you send your request, the new service connector's state is temporarily
// CREATING. When the state changes to ACTIVE, data begins transferring from the
// source service to the target service. For instructions on deactivating and
// activating service connectors, see
// To activate or deactivate a service connector (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
func (client ServiceConnectorClient) CreateServiceConnector(ctx context.Context, request CreateServiceConnectorRequest) (response CreateServiceConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createServiceConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateServiceConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateServiceConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateServiceConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateServiceConnectorResponse")
	}
	return
}

// createServiceConnector implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) createServiceConnector(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/serviceConnectors")
	if err != nil {
		return nil, err
	}

	var response CreateServiceConnectorResponse
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

// DeactivateServiceConnector Deactivates the specified service connector.
// After you send your request, the service connector's state is temporarily
// UPDATING and any data transfer stops. The state then changes to INACTIVE.
// For instructions on deactivating and activating service connectors, see
// To activate or deactivate a service connector (https://docs.cloud.oracle.com/iaas/service-connector-hub/using/index.htm).
func (client ServiceConnectorClient) DeactivateServiceConnector(ctx context.Context, request DeactivateServiceConnectorRequest) (response DeactivateServiceConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deactivateServiceConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateServiceConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateServiceConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateServiceConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateServiceConnectorResponse")
	}
	return
}

// deactivateServiceConnector implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) deactivateServiceConnector(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/serviceConnectors/{serviceConnectorId}/actions/deactivate")
	if err != nil {
		return nil, err
	}

	var response DeactivateServiceConnectorResponse
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

// DeleteServiceConnector Deletes the specified service connector.
// After you send your request, the service connector's state is temporarily
// DELETING and any data transfer stops. The state then changes to DELETED.
func (client ServiceConnectorClient) DeleteServiceConnector(ctx context.Context, request DeleteServiceConnectorRequest) (response DeleteServiceConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteServiceConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteServiceConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteServiceConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteServiceConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteServiceConnectorResponse")
	}
	return
}

// deleteServiceConnector implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) deleteServiceConnector(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/serviceConnectors/{serviceConnectorId}")
	if err != nil {
		return nil, err
	}

	var response DeleteServiceConnectorResponse
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

// GetServiceConnector Gets the specified service connector's configuration information.
func (client ServiceConnectorClient) GetServiceConnector(ctx context.Context, request GetServiceConnectorRequest) (response GetServiceConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceConnectorResponse")
	}
	return
}

// getServiceConnector implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) getServiceConnector(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceConnectors/{serviceConnectorId}")
	if err != nil {
		return nil, err
	}

	var response GetServiceConnectorResponse
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

// GetWorkRequest Gets the details of the specified work request.
func (client ServiceConnectorClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client ServiceConnectorClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListServiceConnectors Lists service connectors in the specified compartment.
func (client ServiceConnectorClient) ListServiceConnectors(ctx context.Context, request ListServiceConnectorsRequest) (response ListServiceConnectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceConnectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceConnectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceConnectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceConnectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceConnectorsResponse")
	}
	return
}

// listServiceConnectors implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) listServiceConnectors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceConnectors")
	if err != nil {
		return nil, err
	}

	var response ListServiceConnectorsResponse
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

// ListWorkRequestErrors Lists work request errors for the specified work request. Results are paginated.
func (client ServiceConnectorClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client ServiceConnectorClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListWorkRequestLogs Lists logs for the specified work request. Results are paginated.
func (client ServiceConnectorClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client ServiceConnectorClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// ListWorkRequests Lists the work requests in the specified compartment.
func (client ServiceConnectorClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client ServiceConnectorClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
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

// UpdateServiceConnector Updates the configuration information for the specified service connector.
// After you send your request, the service connector's state is temporarily
// UPDATING and any data transfer pauses. The state then changes back to its
// original value: if ACTIVE, then data transfer resumes.
func (client ServiceConnectorClient) UpdateServiceConnector(ctx context.Context, request UpdateServiceConnectorRequest) (response UpdateServiceConnectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateServiceConnector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateServiceConnectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateServiceConnectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateServiceConnectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateServiceConnectorResponse")
	}
	return
}

// updateServiceConnector implements the OCIOperation interface (enables retrying operations)
func (client ServiceConnectorClient) updateServiceConnector(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/serviceConnectors/{serviceConnectorId}")
	if err != nil {
		return nil, err
	}

	var response UpdateServiceConnectorResponse
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
