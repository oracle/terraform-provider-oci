// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	client.BasePath = "20230831"
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

// CreateRunbook Creates a new Runbook.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateRunbook.go.html to see an example of how to use CreateRunbook API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Runbook/CreateRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "CreateRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTaskRecord Creates a new Task.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateTaskRecord.go.html to see an example of how to use CreateTaskRecord API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/TaskRecord/CreateTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "CreateTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRunbook Deletes a Runbook resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteRunbook.go.html to see an example of how to use DeleteRunbook API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Runbook/DeleteRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "DeleteRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTaskRecord Deletes a Task Record resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteTaskRecord.go.html to see an example of how to use DeleteTaskRecord API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/TaskRecord/DeleteTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "DeleteTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRunbook Get the details of a runbook in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetRunbook.go.html to see an example of how to use GetRunbook API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Runbook/GetRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTaskRecord Gets a Task by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetTaskRecord.go.html to see an example of how to use GetTaskRecord API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/TaskRecord/GetTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRunbooks List runbooks in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbooks.go.html to see an example of how to use ListRunbooks API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/RunbookCollection/ListRunbooks"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListRunbooks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaskRecords Returns a list of TaskRecords.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListTaskRecords.go.html to see an example of how to use ListTaskRecords API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/TaskRecordCollection/ListTaskRecords"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListTaskRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishRunbook Publish a Runbook.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/PublishRunbook.go.html to see an example of how to use PublishRunbook API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Runbook/PublishRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "PublishRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetDefaultRunbook Publish a Runbook.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/SetDefaultRunbook.go.html to see an example of how to use SetDefaultRunbook API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Runbook/SetDefaultRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "SetDefaultRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRunbook Updates the Ronbook
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateRunbook.go.html to see an example of how to use UpdateRunbook API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Runbook/UpdateRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "UpdateRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTaskRecord Updates the Task
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateTaskRecord.go.html to see an example of how to use UpdateTaskRecord API.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/TaskRecord/UpdateTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "UpdateTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
