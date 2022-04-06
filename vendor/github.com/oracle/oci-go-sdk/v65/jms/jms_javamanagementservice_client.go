// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//JavaManagementServiceClient a client for JavaManagementService
type JavaManagementServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewJavaManagementServiceClientWithConfigurationProvider Creates a new default JavaManagementService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewJavaManagementServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client JavaManagementServiceClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newJavaManagementServiceClientFromBaseClient(baseClient, provider)
}

// NewJavaManagementServiceClientWithOboToken Creates a new default JavaManagementService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewJavaManagementServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client JavaManagementServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newJavaManagementServiceClientFromBaseClient(baseClient, configProvider)
}

func newJavaManagementServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client JavaManagementServiceClient, err error) {
	// JavaManagementService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("JavaManagementService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = JavaManagementServiceClient{BaseClient: baseClient}
	client.BasePath = "20210610"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *JavaManagementServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("jms", "https://javamanagement.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *JavaManagementServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *JavaManagementServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Deletes the work request specified by an identifier.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client JavaManagementServiceClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
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
func (client JavaManagementServiceClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "JavaManagementService", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFleetCompartment Move a specified Fleet into the compartment identified in the POST form. When provided, If-Match is checked against ETag values of the resource.
// A default retry strategy applies to this operation ChangeFleetCompartment()
func (client JavaManagementServiceClient) ChangeFleetCompartment(ctx context.Context, request ChangeFleetCompartmentRequest) (response ChangeFleetCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFleetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFleetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFleetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFleetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFleetCompartmentResponse")
	}
	return
}

// changeFleetCompartment implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) changeFleetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFleetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/ChangeFleetCompartment"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ChangeFleetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBlocklist Add a new record to the fleet blocklist.
// A default retry strategy applies to this operation CreateBlocklist()
func (client JavaManagementServiceClient) CreateBlocklist(ctx context.Context, request CreateBlocklistRequest) (response CreateBlocklistResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBlocklist, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBlocklistResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBlocklistResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBlocklistResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBlocklistResponse")
	}
	return
}

// createBlocklist implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) createBlocklist(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/blocklists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBlocklistResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Blocklist/CreateBlocklist"
		err = common.PostProcessServiceError(err, "JavaManagementService", "CreateBlocklist", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFleet Create a new Fleet using the information provided.
// A default retry strategy applies to this operation CreateFleet()
func (client JavaManagementServiceClient) CreateFleet(ctx context.Context, request CreateFleetRequest) (response CreateFleetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFleetResponse")
	}
	return
}

// createFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) createFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/CreateFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "CreateFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBlocklist Deletes the blocklist record specified by an identifier.
// A default retry strategy applies to this operation DeleteBlocklist()
func (client JavaManagementServiceClient) DeleteBlocklist(ctx context.Context, request DeleteBlocklistRequest) (response DeleteBlocklistResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBlocklist, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBlocklistResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBlocklistResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBlocklistResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBlocklistResponse")
	}
	return
}

// deleteBlocklist implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deleteBlocklist(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}/blocklists/{blocklistKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBlocklistResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Blocklist/DeleteBlocklist"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeleteBlocklist", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFleet Deletes the Fleet specified by an identifier.
// A default retry strategy applies to this operation DeleteFleet()
func (client JavaManagementServiceClient) DeleteFleet(ctx context.Context, request DeleteFleetRequest) (response DeleteFleetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFleetResponse")
	}
	return
}

// deleteFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) deleteFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fleets/{fleetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/DeleteFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "DeleteFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFleet Retrieve a Fleet with the specified identifier.
// A default retry strategy applies to this operation GetFleet()
func (client JavaManagementServiceClient) GetFleet(ctx context.Context, request GetFleetRequest) (response GetFleetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFleetResponse")
	}
	return
}

// getFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/GetFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFleetAgentConfiguration Retrieve a Fleet Agent Configuration for the specified Fleet.
// A default retry strategy applies to this operation GetFleetAgentConfiguration()
func (client JavaManagementServiceClient) GetFleetAgentConfiguration(ctx context.Context, request GetFleetAgentConfigurationRequest) (response GetFleetAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFleetAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFleetAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFleetAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFleetAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFleetAgentConfigurationResponse")
	}
	return
}

// getFleetAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) getFleetAgentConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/agentConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFleetAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/FleetAgentConfiguration/GetFleetAgentConfiguration"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetFleetAgentConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Retrieve the details of a work request with the specified ID.
// A default retry strategy applies to this operation GetWorkRequest()
func (client JavaManagementServiceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client JavaManagementServiceClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "JavaManagementService", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBlocklists Returns a list of blocklist entities contained by a fleet.
// A default retry strategy applies to this operation ListBlocklists()
func (client JavaManagementServiceClient) ListBlocklists(ctx context.Context, request ListBlocklistsRequest) (response ListBlocklistsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlocklists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBlocklistsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBlocklistsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlocklistsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlocklistsResponse")
	}
	return
}

// listBlocklists implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listBlocklists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/blocklists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBlocklistsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Blocklist/ListBlocklists"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListBlocklists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFleets Returns a list of all the Fleets contained by a compartment. The query parameter `compartmentId`
// is required unless the query parameter `id` is specified.
// A default retry strategy applies to this operation ListFleets()
func (client JavaManagementServiceClient) ListFleets(ctx context.Context, request ListFleetsRequest) (response ListFleetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFleets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFleetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFleetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFleetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFleetsResponse")
	}
	return
}

// listFleets implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listFleets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFleetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/ListFleets"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListFleets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInstallationSites List Java installation sites in a Fleet filtered by query parameters.
// A default retry strategy applies to this operation ListInstallationSites()
func (client JavaManagementServiceClient) ListInstallationSites(ctx context.Context, request ListInstallationSitesRequest) (response ListInstallationSitesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstallationSites, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstallationSitesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstallationSitesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstallationSitesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstallationSitesResponse")
	}
	return
}

// listInstallationSites implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listInstallationSites(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/installationSites", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInstallationSitesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/InstallationSiteSummary/ListInstallationSites"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListInstallationSites", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJreUsage List Java Runtime usage in a specified host filtered by query parameters.
// A default retry strategy applies to this operation ListJreUsage()
func (client JavaManagementServiceClient) ListJreUsage(ctx context.Context, request ListJreUsageRequest) (response ListJreUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJreUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJreUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJreUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJreUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJreUsageResponse")
	}
	return
}

// listJreUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listJreUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listJreUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJreUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JreUsage/ListJreUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListJreUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkItems Retrieve a (paginated) list of work items for a specified work request.
// A default retry strategy applies to this operation ListWorkItems()
func (client JavaManagementServiceClient) ListWorkItems(ctx context.Context, request ListWorkItemsRequest) (response ListWorkItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkItemsResponse")
	}
	return
}

// listWorkItems implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) listWorkItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/workItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkItemSummary/ListWorkItems"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Retrieve a (paginated) list of errors for a specified work request.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client JavaManagementServiceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client JavaManagementServiceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Retrieve a (paginated) list of logs for a specified work request.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client JavaManagementServiceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client JavaManagementServiceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests List the work requests in a compartment. The query parameter `compartmentId` is required unless the query parameter `id` or `fleetId` is specified.
// A default retry strategy applies to this operation ListWorkRequests()
func (client JavaManagementServiceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client JavaManagementServiceClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "JavaManagementService", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveFleetInstallationSites Remove Java installation sites in a Fleet.
// A default retry strategy applies to this operation RemoveFleetInstallationSites()
func (client JavaManagementServiceClient) RemoveFleetInstallationSites(ctx context.Context, request RemoveFleetInstallationSitesRequest) (response RemoveFleetInstallationSitesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeFleetInstallationSites, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveFleetInstallationSitesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveFleetInstallationSitesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveFleetInstallationSitesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveFleetInstallationSitesResponse")
	}
	return
}

// removeFleetInstallationSites implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) removeFleetInstallationSites(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fleets/{fleetId}/actions/removeInstallationSites", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveFleetInstallationSitesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/InstallationSiteSummary/RemoveFleetInstallationSites"
		err = common.PostProcessServiceError(err, "JavaManagementService", "RemoveFleetInstallationSites", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeApplicationUsage List application usage in a Fleet filtered by query parameters.
// A default retry strategy applies to this operation SummarizeApplicationUsage()
func (client JavaManagementServiceClient) SummarizeApplicationUsage(ctx context.Context, request SummarizeApplicationUsageRequest) (response SummarizeApplicationUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeApplicationUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeApplicationUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeApplicationUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeApplicationUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeApplicationUsageResponse")
	}
	return
}

// summarizeApplicationUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeApplicationUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeApplicationUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeApplicationUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ApplicationUsage/SummarizeApplicationUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeApplicationUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeInstallationUsage List Java installation usage in a Fleet filtered by query parameters.
// A default retry strategy applies to this operation SummarizeInstallationUsage()
func (client JavaManagementServiceClient) SummarizeInstallationUsage(ctx context.Context, request SummarizeInstallationUsageRequest) (response SummarizeInstallationUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeInstallationUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeInstallationUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeInstallationUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeInstallationUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeInstallationUsageResponse")
	}
	return
}

// summarizeInstallationUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeInstallationUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeInstallationUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeInstallationUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/InstallationUsage/SummarizeInstallationUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeInstallationUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeJreUsage List Java Runtime usage in a specified Fleet filtered by query parameters.
// A default retry strategy applies to this operation SummarizeJreUsage()
func (client JavaManagementServiceClient) SummarizeJreUsage(ctx context.Context, request SummarizeJreUsageRequest) (response SummarizeJreUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeJreUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeJreUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeJreUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeJreUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeJreUsageResponse")
	}
	return
}

// summarizeJreUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeJreUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeJreUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeJreUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/JreUsage/SummarizeJreUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeJreUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeManagedInstanceUsage List managed instance usage in a Fleet filtered by query parameters.
// A default retry strategy applies to this operation SummarizeManagedInstanceUsage()
func (client JavaManagementServiceClient) SummarizeManagedInstanceUsage(ctx context.Context, request SummarizeManagedInstanceUsageRequest) (response SummarizeManagedInstanceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeManagedInstanceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeManagedInstanceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeManagedInstanceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeManagedInstanceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeManagedInstanceUsageResponse")
	}
	return
}

// summarizeManagedInstanceUsage implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeManagedInstanceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fleets/{fleetId}/actions/summarizeManagedInstanceUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeManagedInstanceUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/ManagedInstanceUsage/SummarizeManagedInstanceUsage"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeManagedInstanceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeResourceInventory Retrieve the inventory of JMS resources in the specified compartment: a list of the number of _active_ fleets, managed instances, Java Runtimes, Java installations, and applications.
// A default retry strategy applies to this operation SummarizeResourceInventory()
func (client JavaManagementServiceClient) SummarizeResourceInventory(ctx context.Context, request SummarizeResourceInventoryRequest) (response SummarizeResourceInventoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeResourceInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeResourceInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeResourceInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeResourceInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeResourceInventoryResponse")
	}
	return
}

// summarizeResourceInventory implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) summarizeResourceInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/summarizeResourceInventory", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeResourceInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/SummarizeResourceInventory"
		err = common.PostProcessServiceError(err, "JavaManagementService", "SummarizeResourceInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFleet Update the Fleet specified by an identifier.
// A default retry strategy applies to this operation UpdateFleet()
func (client JavaManagementServiceClient) UpdateFleet(ctx context.Context, request UpdateFleetRequest) (response UpdateFleetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFleet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFleetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFleetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFleetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFleetResponse")
	}
	return
}

// updateFleet implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) updateFleet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fleets/{fleetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFleetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/Fleet/UpdateFleet"
		err = common.PostProcessServiceError(err, "JavaManagementService", "UpdateFleet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFleetAgentConfiguration Update the Fleet Agent Configuration for the specified Fleet.
// A default retry strategy applies to this operation UpdateFleetAgentConfiguration()
func (client JavaManagementServiceClient) UpdateFleetAgentConfiguration(ctx context.Context, request UpdateFleetAgentConfigurationRequest) (response UpdateFleetAgentConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFleetAgentConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFleetAgentConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFleetAgentConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFleetAgentConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFleetAgentConfigurationResponse")
	}
	return
}

// updateFleetAgentConfiguration implements the OCIOperation interface (enables retrying operations)
func (client JavaManagementServiceClient) updateFleetAgentConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fleets/{fleetId}/agentConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFleetAgentConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms/20210610/FleetAgentConfiguration/UpdateFleetAgentConfiguration"
		err = common.PostProcessServiceError(err, "JavaManagementService", "UpdateFleetAgentConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
