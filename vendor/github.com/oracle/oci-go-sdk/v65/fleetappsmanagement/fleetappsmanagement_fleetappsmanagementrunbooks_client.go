// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FleetAppsManagementRunbooksClient a client for FleetAppsManagementRunbooks
type FleetAppsManagementRunbooksClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFleetAppsManagementRunbooksClientWithConfigurationProvider Creates a new default FleetAppsManagementRunbooks client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFleetAppsManagementRunbooksClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FleetAppsManagementRunbooksClient, err error) {
	if enabled := common.CheckForEnabledServices("fleetappsmanagement"); !enabled {
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
	return newFleetAppsManagementRunbooksClientFromBaseClient(baseClient, provider)
}

// NewFleetAppsManagementRunbooksClientWithOboToken Creates a new default FleetAppsManagementRunbooks client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFleetAppsManagementRunbooksClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FleetAppsManagementRunbooksClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFleetAppsManagementRunbooksClientFromBaseClient(baseClient, configProvider)
}

func newFleetAppsManagementRunbooksClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FleetAppsManagementRunbooksClient, err error) {
	// FleetAppsManagementRunbooks service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FleetAppsManagementRunbooks"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FleetAppsManagementRunbooksClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FleetAppsManagementRunbooksClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fleetappsmanagement", "https://fams.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FleetAppsManagementRunbooksClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FleetAppsManagementRunbooksClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeRunbookCompartment Moves a Runbook into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ChangeRunbookCompartment.go.html to see an example of how to use ChangeRunbookCompartment API.
// A default retry strategy applies to this operation ChangeRunbookCompartment()
func (client FleetAppsManagementRunbooksClient) ChangeRunbookCompartment(ctx context.Context, request ChangeRunbookCompartmentRequest) (response ChangeRunbookCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRunbookCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRunbookCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRunbookCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRunbookCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRunbookCompartmentResponse")
	}
	return
}

// changeRunbookCompartment implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) changeRunbookCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/{runbookId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRunbookCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/ChangeRunbookCompartment"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ChangeRunbookCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeTaskRecordCompartment Moves a task record into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ChangeTaskRecordCompartment.go.html to see an example of how to use ChangeTaskRecordCompartment API.
// A default retry strategy applies to this operation ChangeTaskRecordCompartment()
func (client FleetAppsManagementRunbooksClient) ChangeTaskRecordCompartment(ctx context.Context, request ChangeTaskRecordCompartmentRequest) (response ChangeTaskRecordCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTaskRecordCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTaskRecordCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTaskRecordCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTaskRecordCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTaskRecordCompartmentResponse")
	}
	return
}

// changeTaskRecordCompartment implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) changeTaskRecordCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/taskRecords/{taskRecordId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTaskRecordCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/TaskRecord/ChangeTaskRecordCompartment"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ChangeTaskRecordCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRunbook Creates a runbook.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateRunbook.go.html to see an example of how to use CreateRunbook API.
// A default retry strategy applies to this operation CreateRunbook()
func (client FleetAppsManagementRunbooksClient) CreateRunbook(ctx context.Context, request CreateRunbookRequest) (response CreateRunbookResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRunbookResponse")
	}
	return
}

// createRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) createRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/CreateRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "CreateRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRunbookVersion Add RunbookVersion in Fleet Application Management.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateRunbookVersion.go.html to see an example of how to use CreateRunbookVersion API.
// A default retry strategy applies to this operation CreateRunbookVersion()
func (client FleetAppsManagementRunbooksClient) CreateRunbookVersion(ctx context.Context, request CreateRunbookVersionRequest) (response CreateRunbookVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRunbookVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRunbookVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRunbookVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRunbookVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRunbookVersionResponse")
	}
	return
}

// createRunbookVersion implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) createRunbookVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbookVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRunbookVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookVersion/CreateRunbookVersion"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "CreateRunbookVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTaskRecord Creates a new task record.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateTaskRecord.go.html to see an example of how to use CreateTaskRecord API.
// A default retry strategy applies to this operation CreateTaskRecord()
func (client FleetAppsManagementRunbooksClient) CreateTaskRecord(ctx context.Context, request CreateTaskRecordRequest) (response CreateTaskRecordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTaskRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTaskRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTaskRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTaskRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTaskRecordResponse")
	}
	return
}

// createTaskRecord implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) createTaskRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/taskRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTaskRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/TaskRecord/CreateTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "CreateTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRunbook Deletes a runbook specified by the identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteRunbook.go.html to see an example of how to use DeleteRunbook API.
// A default retry strategy applies to this operation DeleteRunbook()
func (client FleetAppsManagementRunbooksClient) DeleteRunbook(ctx context.Context, request DeleteRunbookRequest) (response DeleteRunbookResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRunbookResponse")
	}
	return
}

// deleteRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) deleteRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/runbooks/{runbookId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/DeleteRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "DeleteRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRunbookVersion Removes a Runbook Version from the runbook in Fleet Application Management.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteRunbookVersion.go.html to see an example of how to use DeleteRunbookVersion API.
// A default retry strategy applies to this operation DeleteRunbookVersion()
func (client FleetAppsManagementRunbooksClient) DeleteRunbookVersion(ctx context.Context, request DeleteRunbookVersionRequest) (response DeleteRunbookVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRunbookVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRunbookVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRunbookVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRunbookVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRunbookVersionResponse")
	}
	return
}

// deleteRunbookVersion implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) deleteRunbookVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/runbookVersions/{runbookVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRunbookVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookVersion/DeleteRunbookVersion"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "DeleteRunbookVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTaskRecord Deletes the task record specified by an identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteTaskRecord.go.html to see an example of how to use DeleteTaskRecord API.
// A default retry strategy applies to this operation DeleteTaskRecord()
func (client FleetAppsManagementRunbooksClient) DeleteTaskRecord(ctx context.Context, request DeleteTaskRecordRequest) (response DeleteTaskRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTaskRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTaskRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTaskRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTaskRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTaskRecordResponse")
	}
	return
}

// deleteTaskRecord implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) deleteTaskRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/taskRecords/{taskRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTaskRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/TaskRecord/DeleteTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "DeleteTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportRunbook Export the specified version of the runbook.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ExportRunbook.go.html to see an example of how to use ExportRunbook API.
// A default retry strategy applies to this operation ExportRunbook()
func (client FleetAppsManagementRunbooksClient) ExportRunbook(ctx context.Context, request ExportRunbookRequest) (response ExportRunbookResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportRunbookResponse")
	}
	return
}

// exportRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) exportRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/{runbookId}/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/ExportRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ExportRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportRunbookVersion Export the specified version of the runbook.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ExportRunbookVersion.go.html to see an example of how to use ExportRunbookVersion API.
// A default retry strategy applies to this operation ExportRunbookVersion()
func (client FleetAppsManagementRunbooksClient) ExportRunbookVersion(ctx context.Context, request ExportRunbookVersionRequest) (response ExportRunbookVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportRunbookVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportRunbookVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportRunbookVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportRunbookVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportRunbookVersionResponse")
	}
	return
}

// exportRunbookVersion implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) exportRunbookVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbookVersions/{runbookVersionId}/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportRunbookVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/ExportRunbookVersion"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ExportRunbookVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FindRunbookExportDependency Find runbook export Dependencies
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/FindRunbookExportDependency.go.html to see an example of how to use FindRunbookExportDependency API.
// A default retry strategy applies to this operation FindRunbookExportDependency()
func (client FleetAppsManagementRunbooksClient) FindRunbookExportDependency(ctx context.Context, request FindRunbookExportDependencyRequest) (response FindRunbookExportDependencyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.findRunbookExportDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FindRunbookExportDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FindRunbookExportDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FindRunbookExportDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FindRunbookExportDependencyResponse")
	}
	return
}

// findRunbookExportDependency implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) findRunbookExportDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/actions/findExportDependencies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FindRunbookExportDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookExportDependencyCollection/FindRunbookExportDependency"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "FindRunbookExportDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FindRunbookImportDependency Find runbook import Dependencies
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/FindRunbookImportDependency.go.html to see an example of how to use FindRunbookImportDependency API.
// A default retry strategy applies to this operation FindRunbookImportDependency()
func (client FleetAppsManagementRunbooksClient) FindRunbookImportDependency(ctx context.Context, request FindRunbookImportDependencyRequest) (response FindRunbookImportDependencyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.findRunbookImportDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FindRunbookImportDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FindRunbookImportDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FindRunbookImportDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FindRunbookImportDependencyResponse")
	}
	return
}

// findRunbookImportDependency implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) findRunbookImportDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/actions/findImportDependencies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FindRunbookImportDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookImportDependencyCollection/FindRunbookImportDependency"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "FindRunbookImportDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRunbook Get the details of a runbook in Fleet Application Management.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetRunbook.go.html to see an example of how to use GetRunbook API.
// A default retry strategy applies to this operation GetRunbook()
func (client FleetAppsManagementRunbooksClient) GetRunbook(ctx context.Context, request GetRunbookRequest) (response GetRunbookResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRunbookResponse")
	}
	return
}

// getRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) getRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks/{runbookId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/GetRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRunbookExport Get the runbook export status for provided runbook and exportId.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetRunbookExport.go.html to see an example of how to use GetRunbookExport API.
// A default retry strategy applies to this operation GetRunbookExport()
func (client FleetAppsManagementRunbooksClient) GetRunbookExport(ctx context.Context, request GetRunbookExportRequest) (response GetRunbookExportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRunbookExport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRunbookExportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRunbookExportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRunbookExportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRunbookExportResponse")
	}
	return
}

// getRunbookExport implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) getRunbookExport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks/{runbookId}/exports/{exportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRunbookExportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookExport/GetRunbookExport"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetRunbookExport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRunbookImport Get the runbook import status for provided runbook and importId.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetRunbookImport.go.html to see an example of how to use GetRunbookImport API.
// A default retry strategy applies to this operation GetRunbookImport()
func (client FleetAppsManagementRunbooksClient) GetRunbookImport(ctx context.Context, request GetRunbookImportRequest) (response GetRunbookImportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRunbookImport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRunbookImportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRunbookImportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRunbookImportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRunbookImportResponse")
	}
	return
}

// getRunbookImport implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) getRunbookImport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks/{runbookId}/imports/{importId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRunbookImportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookImport/GetRunbookImport"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetRunbookImport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRunbookVersion Gets a Runbook Version by identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetRunbookVersion.go.html to see an example of how to use GetRunbookVersion API.
// A default retry strategy applies to this operation GetRunbookVersion()
func (client FleetAppsManagementRunbooksClient) GetRunbookVersion(ctx context.Context, request GetRunbookVersionRequest) (response GetRunbookVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRunbookVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRunbookVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRunbookVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRunbookVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRunbookVersionResponse")
	}
	return
}

// getRunbookVersion implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) getRunbookVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbookVersions/{runbookVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRunbookVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookVersion/GetRunbookVersion"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetRunbookVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTaskRecord Gets information for the specified task record.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetTaskRecord.go.html to see an example of how to use GetTaskRecord API.
// A default retry strategy applies to this operation GetTaskRecord()
func (client FleetAppsManagementRunbooksClient) GetTaskRecord(ctx context.Context, request GetTaskRecordRequest) (response GetTaskRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTaskRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTaskRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTaskRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTaskRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTaskRecordResponse")
	}
	return
}

// getTaskRecord implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) getTaskRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taskRecords/{taskRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTaskRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/TaskRecord/GetTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportRunbook Import the specified version of the runbook.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ImportRunbook.go.html to see an example of how to use ImportRunbook API.
// A default retry strategy applies to this operation ImportRunbook()
func (client FleetAppsManagementRunbooksClient) ImportRunbook(ctx context.Context, request ImportRunbookRequest) (response ImportRunbookResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportRunbookResponse")
	}
	return
}

// importRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) importRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/actions/import", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/ImportRunbookDetails/ImportRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ImportRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportRunbookPrecheck Precheck for import runbook.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ImportRunbookPrecheck.go.html to see an example of how to use ImportRunbookPrecheck API.
// A default retry strategy applies to this operation ImportRunbookPrecheck()
func (client FleetAppsManagementRunbooksClient) ImportRunbookPrecheck(ctx context.Context, request ImportRunbookPrecheckRequest) (response ImportRunbookPrecheckResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importRunbookPrecheck, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportRunbookPrecheckResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportRunbookPrecheckResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportRunbookPrecheckResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportRunbookPrecheckResponse")
	}
	return
}

// importRunbookPrecheck implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) importRunbookPrecheck(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/actions/importPrecheck", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportRunbookPrecheckResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/ImportRunbookPrecheckDetails/ImportRunbookPrecheck"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ImportRunbookPrecheck", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportRunbookVersion Export the specified version of the runbook.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ImportRunbookVersion.go.html to see an example of how to use ImportRunbookVersion API.
// A default retry strategy applies to this operation ImportRunbookVersion()
func (client FleetAppsManagementRunbooksClient) ImportRunbookVersion(ctx context.Context, request ImportRunbookVersionRequest) (response ImportRunbookVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importRunbookVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportRunbookVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportRunbookVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportRunbookVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportRunbookVersionResponse")
	}
	return
}

// importRunbookVersion implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) importRunbookVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbookVersions/actions/import", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportRunbookVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/ImportRunbookVersionDetails/ImportRunbookVersion"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ImportRunbookVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRunbookExportStatuses Returns a list of all the Runbook export status in the specified compartment.
// The query parameter `compartmentId` is required.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbookExportStatuses.go.html to see an example of how to use ListRunbookExportStatuses API.
// A default retry strategy applies to this operation ListRunbookExportStatuses()
func (client FleetAppsManagementRunbooksClient) ListRunbookExportStatuses(ctx context.Context, request ListRunbookExportStatusesRequest) (response ListRunbookExportStatusesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRunbookExportStatuses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRunbookExportStatusesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRunbookExportStatusesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRunbookExportStatusesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRunbookExportStatusesResponse")
	}
	return
}

// listRunbookExportStatuses implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) listRunbookExportStatuses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks/exportStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRunbookExportStatusesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookExportStatusCollection/ListRunbookExportStatuses"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListRunbookExportStatuses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRunbookImportStatuses Returns a list of all the Runbook import status in the specified compartment.
// The query parameter `compartmentId` is required.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbookImportStatuses.go.html to see an example of how to use ListRunbookImportStatuses API.
// A default retry strategy applies to this operation ListRunbookImportStatuses()
func (client FleetAppsManagementRunbooksClient) ListRunbookImportStatuses(ctx context.Context, request ListRunbookImportStatusesRequest) (response ListRunbookImportStatusesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRunbookImportStatuses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRunbookImportStatusesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRunbookImportStatusesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRunbookImportStatusesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRunbookImportStatusesResponse")
	}
	return
}

// listRunbookImportStatuses implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) listRunbookImportStatuses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks/importStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRunbookImportStatusesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookImportStatusCollection/ListRunbookImportStatuses"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListRunbookImportStatuses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRunbookVersions List versions for a runbook in Fleet Application Management.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbookVersions.go.html to see an example of how to use ListRunbookVersions API.
// A default retry strategy applies to this operation ListRunbookVersions()
func (client FleetAppsManagementRunbooksClient) ListRunbookVersions(ctx context.Context, request ListRunbookVersionsRequest) (response ListRunbookVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRunbookVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRunbookVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRunbookVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRunbookVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRunbookVersionsResponse")
	}
	return
}

// listRunbookVersions implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) listRunbookVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbookVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRunbookVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookVersionCollection/ListRunbookVersions"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListRunbookVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRunbooks Returns a list of all the runbooks in the specified compartment.
// The query parameter `compartmentId` is required unless the query parameter `id` is specified.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbooks.go.html to see an example of how to use ListRunbooks API.
// A default retry strategy applies to this operation ListRunbooks()
func (client FleetAppsManagementRunbooksClient) ListRunbooks(ctx context.Context, request ListRunbooksRequest) (response ListRunbooksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRunbooks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRunbooksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRunbooksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRunbooksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRunbooksResponse")
	}
	return
}

// listRunbooks implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) listRunbooks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRunbooksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookCollection/ListRunbooks"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListRunbooks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaskRecords Returns a list of all the task records in the specified compartment.
// The query parameter `compartmentId` is required unless the query parameter `id` is specified.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListTaskRecords.go.html to see an example of how to use ListTaskRecords API.
// A default retry strategy applies to this operation ListTaskRecords()
func (client FleetAppsManagementRunbooksClient) ListTaskRecords(ctx context.Context, request ListTaskRecordsRequest) (response ListTaskRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaskRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaskRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaskRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaskRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaskRecordsResponse")
	}
	return
}

// listTaskRecords implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) listTaskRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taskRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaskRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/TaskRecordCollection/ListTaskRecords"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListTaskRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishRunbook Publish the specified version of the runbook.
// The specified version of the runbook becomes acitve when it is published.Only active versions of runbook can be used in execution.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/PublishRunbook.go.html to see an example of how to use PublishRunbook API.
// A default retry strategy applies to this operation PublishRunbook()
func (client FleetAppsManagementRunbooksClient) PublishRunbook(ctx context.Context, request PublishRunbookRequest) (response PublishRunbookResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.publishRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishRunbookResponse")
	}
	return
}

// publishRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) publishRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/actions/publishRunbook", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/PublishRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "PublishRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetDefaultRunbook Set a runbook as default.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/SetDefaultRunbook.go.html to see an example of how to use SetDefaultRunbook API.
// A default retry strategy applies to this operation SetDefaultRunbook()
func (client FleetAppsManagementRunbooksClient) SetDefaultRunbook(ctx context.Context, request SetDefaultRunbookRequest) (response SetDefaultRunbookResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.setDefaultRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetDefaultRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetDefaultRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetDefaultRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetDefaultRunbookResponse")
	}
	return
}

// setDefaultRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) setDefaultRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/runbooks/actions/setDefaultRunbook", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetDefaultRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/SetDefaultRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "SetDefaultRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRunbook Updates the runbook specified by the identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateRunbook.go.html to see an example of how to use UpdateRunbook API.
// A default retry strategy applies to this operation UpdateRunbook()
func (client FleetAppsManagementRunbooksClient) UpdateRunbook(ctx context.Context, request UpdateRunbookRequest) (response UpdateRunbookResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRunbookResponse")
	}
	return
}

// updateRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) updateRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/runbooks/{runbookId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/Runbook/UpdateRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "UpdateRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRunbookVersion Updates the RunbookVersion.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateRunbookVersion.go.html to see an example of how to use UpdateRunbookVersion API.
// A default retry strategy applies to this operation UpdateRunbookVersion()
func (client FleetAppsManagementRunbooksClient) UpdateRunbookVersion(ctx context.Context, request UpdateRunbookVersionRequest) (response UpdateRunbookVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRunbookVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRunbookVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRunbookVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRunbookVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRunbookVersionResponse")
	}
	return
}

// updateRunbookVersion implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) updateRunbookVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/runbookVersions/{runbookVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRunbookVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/RunbookVersion/UpdateRunbookVersion"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "UpdateRunbookVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTaskRecord Updates certain attributes for the specified task record.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateTaskRecord.go.html to see an example of how to use UpdateTaskRecord API.
// A default retry strategy applies to this operation UpdateTaskRecord()
func (client FleetAppsManagementRunbooksClient) UpdateTaskRecord(ctx context.Context, request UpdateTaskRecordRequest) (response UpdateTaskRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTaskRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTaskRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTaskRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTaskRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTaskRecordResponse")
	}
	return
}

// updateTaskRecord implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) updateTaskRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/taskRecords/{taskRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTaskRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/TaskRecord/UpdateTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "UpdateTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
