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

// FleetAppsManagementOperationsClient a client for FleetAppsManagementOperations
type FleetAppsManagementOperationsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFleetAppsManagementOperationsClientWithConfigurationProvider Creates a new default FleetAppsManagementOperations client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFleetAppsManagementOperationsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FleetAppsManagementOperationsClient, err error) {
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
	return newFleetAppsManagementOperationsClientFromBaseClient(baseClient, provider)
}

// NewFleetAppsManagementOperationsClientWithOboToken Creates a new default FleetAppsManagementOperations client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFleetAppsManagementOperationsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FleetAppsManagementOperationsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFleetAppsManagementOperationsClientFromBaseClient(baseClient, configProvider)
}

func newFleetAppsManagementOperationsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FleetAppsManagementOperationsClient, err error) {
	// FleetAppsManagementOperations service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FleetAppsManagementOperations"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FleetAppsManagementOperationsClient{BaseClient: baseClient}
	client.BasePath = "20230831"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FleetAppsManagementOperationsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fleetappsmanagement", "https://fams.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FleetAppsManagementOperationsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FleetAppsManagementOperationsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreatePatch Creates a new Patch.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreatePatch.go.html to see an example of how to use CreatePatch API.
// A default retry strategy applies to this operation CreatePatch()
func (client FleetAppsManagementOperationsClient) CreatePatch(ctx context.Context, request CreatePatchRequest) (response CreatePatchResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePatchResponse")
	}
	return
}

// createPatch implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) createPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Patch/CreatePatch"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "CreatePatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSchedulerDefinition Create a SchedulerDefinition to perform lifecycle operations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateSchedulerDefinition.go.html to see an example of how to use CreateSchedulerDefinition API.
// A default retry strategy applies to this operation CreateSchedulerDefinition()
func (client FleetAppsManagementOperationsClient) CreateSchedulerDefinition(ctx context.Context, request CreateSchedulerDefinitionRequest) (response CreateSchedulerDefinitionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSchedulerDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSchedulerDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSchedulerDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSchedulerDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSchedulerDefinitionResponse")
	}
	return
}

// createSchedulerDefinition implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) createSchedulerDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/schedulerDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSchedulerDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerDefinition/CreateSchedulerDefinition"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "CreateSchedulerDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePatch Deletes a Patch resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeletePatch.go.html to see an example of how to use DeletePatch API.
// A default retry strategy applies to this operation DeletePatch()
func (client FleetAppsManagementOperationsClient) DeletePatch(ctx context.Context, request DeletePatchRequest) (response DeletePatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePatchResponse")
	}
	return
}

// deletePatch implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) deletePatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/patches/{patchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Patch/DeletePatch"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "DeletePatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSchedulerDefinition Deletes a SchedulerDefinition resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteSchedulerDefinition.go.html to see an example of how to use DeleteSchedulerDefinition API.
// A default retry strategy applies to this operation DeleteSchedulerDefinition()
func (client FleetAppsManagementOperationsClient) DeleteSchedulerDefinition(ctx context.Context, request DeleteSchedulerDefinitionRequest) (response DeleteSchedulerDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSchedulerDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSchedulerDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSchedulerDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSchedulerDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSchedulerDefinitionResponse")
	}
	return
}

// deleteSchedulerDefinition implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) deleteSchedulerDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/schedulerDefinitions/{schedulerDefinitionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSchedulerDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerDefinition/DeleteSchedulerDefinition"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "DeleteSchedulerDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSchedulerJob Delete a lifecycle operation schedule in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteSchedulerJob.go.html to see an example of how to use DeleteSchedulerJob API.
// A default retry strategy applies to this operation DeleteSchedulerJob()
func (client FleetAppsManagementOperationsClient) DeleteSchedulerJob(ctx context.Context, request DeleteSchedulerJobRequest) (response DeleteSchedulerJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSchedulerJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSchedulerJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSchedulerJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSchedulerJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSchedulerJobResponse")
	}
	return
}

// deleteSchedulerJob implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) deleteSchedulerJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/schedulerJobs/{schedulerJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSchedulerJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerJob/DeleteSchedulerJob"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "DeleteSchedulerJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportComplianceReport Generate Compliance Report
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ExportComplianceReport.go.html to see an example of how to use ExportComplianceReport API.
// A default retry strategy applies to this operation ExportComplianceReport()
func (client FleetAppsManagementOperationsClient) ExportComplianceReport(ctx context.Context, request ExportComplianceReportRequest) (response ExportComplianceReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportComplianceReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportComplianceReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportComplianceReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportComplianceReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportComplianceReportResponse")
	}
	return
}

// exportComplianceReport implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) exportComplianceReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/complianceRecords/actions/exportComplianceReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportComplianceReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/ComplianceRecord/ExportComplianceReport"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ExportComplianceReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExecution Get Task Execution by Identifier for a Resource within an action group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetExecution.go.html to see an example of how to use GetExecution API.
// A default retry strategy applies to this operation GetExecution()
func (client FleetAppsManagementOperationsClient) GetExecution(ctx context.Context, request GetExecutionRequest) (response GetExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExecutionResponse")
	}
	return
}

// getExecution implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) getExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobs/{schedulerJobId}/jobActivities/{jobActivityId}/resources/{resourceId}/executions/{executionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Execution/GetExecution"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "GetExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobActivity Gets activity details by identifier for a job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetJobActivity.go.html to see an example of how to use GetJobActivity API.
// A default retry strategy applies to this operation GetJobActivity()
func (client FleetAppsManagementOperationsClient) GetJobActivity(ctx context.Context, request GetJobActivityRequest) (response GetJobActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobActivityResponse")
	}
	return
}

// getJobActivity implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) getJobActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobs/{schedulerJobId}/jobActivities/{jobActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/JobActivity/GetJobActivity"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "GetJobActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPatch Gets a Patch by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetPatch.go.html to see an example of how to use GetPatch API.
// A default retry strategy applies to this operation GetPatch()
func (client FleetAppsManagementOperationsClient) GetPatch(ctx context.Context, request GetPatchRequest) (response GetPatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPatchResponse")
	}
	return
}

// getPatch implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) getPatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patches/{patchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Patch/GetPatch"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "GetPatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSchedulerDefinition Get the details of a SchedulerDefinition that performs lifecycle management operations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetSchedulerDefinition.go.html to see an example of how to use GetSchedulerDefinition API.
// A default retry strategy applies to this operation GetSchedulerDefinition()
func (client FleetAppsManagementOperationsClient) GetSchedulerDefinition(ctx context.Context, request GetSchedulerDefinitionRequest) (response GetSchedulerDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSchedulerDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSchedulerDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSchedulerDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSchedulerDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSchedulerDefinitionResponse")
	}
	return
}

// getSchedulerDefinition implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) getSchedulerDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerDefinitions/{schedulerDefinitionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSchedulerDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerDefinition/GetSchedulerDefinition"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "GetSchedulerDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSchedulerJob Get the details of a lifecycle management operations job in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetSchedulerJob.go.html to see an example of how to use GetSchedulerJob API.
// A default retry strategy applies to this operation GetSchedulerJob()
func (client FleetAppsManagementOperationsClient) GetSchedulerJob(ctx context.Context, request GetSchedulerJobRequest) (response GetSchedulerJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSchedulerJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSchedulerJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSchedulerJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSchedulerJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSchedulerJobResponse")
	}
	return
}

// getSchedulerJob implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) getSchedulerJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobs/{schedulerJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSchedulerJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerJob/GetSchedulerJob"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "GetSchedulerJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListComplianceRecords Gets a list of complianceDetails.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListComplianceRecords.go.html to see an example of how to use ListComplianceRecords API.
// A default retry strategy applies to this operation ListComplianceRecords()
func (client FleetAppsManagementOperationsClient) ListComplianceRecords(ctx context.Context, request ListComplianceRecordsRequest) (response ListComplianceRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listComplianceRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListComplianceRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListComplianceRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListComplianceRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListComplianceRecordsResponse")
	}
	return
}

// listComplianceRecords implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listComplianceRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/complianceRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListComplianceRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/ComplianceRecordCollection/ListComplianceRecords"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListComplianceRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExecutions Returns a list of Task Executions for a Resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListExecutions.go.html to see an example of how to use ListExecutions API.
// A default retry strategy applies to this operation ListExecutions()
func (client FleetAppsManagementOperationsClient) ListExecutions(ctx context.Context, request ListExecutionsRequest) (response ListExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExecutionsResponse")
	}
	return
}

// listExecutions implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobs/{schedulerJobId}/jobActivities/{jobActivityId}/resources/{resourceId}/executions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/ExecutionCollection/ListExecutions"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatches Returns a list of Patches.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListPatches.go.html to see an example of how to use ListPatches API.
// A default retry strategy applies to this operation ListPatches()
func (client FleetAppsManagementOperationsClient) ListPatches(ctx context.Context, request ListPatchesRequest) (response ListPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchesResponse")
	}
	return
}

// listPatches implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/PatchCollection/ListPatches"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResources Returns a list of resources for an Activity Execution.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListResources.go.html to see an example of how to use ListResources API.
// A default retry strategy applies to this operation ListResources()
func (client FleetAppsManagementOperationsClient) ListResources(ctx context.Context, request ListResourcesRequest) (response ListResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourcesResponse")
	}
	return
}

// listResources implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobs/{schedulerJobId}/jobActivities/{jobActivityId}/resources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/ResourceCollection/ListResources"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledFleets Returns a list of ScheduledFleets.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListScheduledFleets.go.html to see an example of how to use ListScheduledFleets API.
// A default retry strategy applies to this operation ListScheduledFleets()
func (client FleetAppsManagementOperationsClient) ListScheduledFleets(ctx context.Context, request ListScheduledFleetsRequest) (response ListScheduledFleetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledFleets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledFleetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledFleetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledFleetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledFleetsResponse")
	}
	return
}

// listScheduledFleets implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listScheduledFleets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerDefinitions/{schedulerDefinitionId}/scheduledFleets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledFleetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/ScheduledFleetCollection/ListScheduledFleets"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListScheduledFleets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSchedulerDefinitions List all lifecycle management schedules in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListSchedulerDefinitions.go.html to see an example of how to use ListSchedulerDefinitions API.
// A default retry strategy applies to this operation ListSchedulerDefinitions()
func (client FleetAppsManagementOperationsClient) ListSchedulerDefinitions(ctx context.Context, request ListSchedulerDefinitionsRequest) (response ListSchedulerDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSchedulerDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSchedulerDefinitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSchedulerDefinitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSchedulerDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSchedulerDefinitionsResponse")
	}
	return
}

// listSchedulerDefinitions implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listSchedulerDefinitions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSchedulerDefinitionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerDefinitionCollection/ListSchedulerDefinitions"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListSchedulerDefinitions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSchedulerJobs List scheduled lifecycle operation jobs in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListSchedulerJobs.go.html to see an example of how to use ListSchedulerJobs API.
// A default retry strategy applies to this operation ListSchedulerJobs()
func (client FleetAppsManagementOperationsClient) ListSchedulerJobs(ctx context.Context, request ListSchedulerJobsRequest) (response ListSchedulerJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSchedulerJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSchedulerJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSchedulerJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSchedulerJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSchedulerJobsResponse")
	}
	return
}

// listSchedulerJobs implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listSchedulerJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSchedulerJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerJobCollection/ListSchedulerJobs"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListSchedulerJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSteps Returns a list of Steps for an Activity Execution.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListSteps.go.html to see an example of how to use ListSteps API.
// A default retry strategy applies to this operation ListSteps()
func (client FleetAppsManagementOperationsClient) ListSteps(ctx context.Context, request ListStepsRequest) (response ListStepsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSteps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStepsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStepsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStepsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStepsResponse")
	}
	return
}

// listSteps implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) listSteps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobs/{schedulerJobId}/jobActivities/{jobActivityId}/steps", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStepsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/StepCollection/ListSteps"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ListSteps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ManageJobExecution Manage execution actions for a Job like retrying or pausing a task.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ManageJobExecution.go.html to see an example of how to use ManageJobExecution API.
// A default retry strategy applies to this operation ManageJobExecution()
func (client FleetAppsManagementOperationsClient) ManageJobExecution(ctx context.Context, request ManageJobExecutionRequest) (response ManageJobExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.manageJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ManageJobExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ManageJobExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ManageJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ManageJobExecutionResponse")
	}
	return
}

// manageJobExecution implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) manageJobExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/schedulerJobs/{schedulerJobId}/actions/manageJobExecution", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ManageJobExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerJob/ManageJobExecution"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "ManageJobExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeComplianceRecordCounts Retrieve  aggregated summary information of ComplianceRecords within a Tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/SummarizeComplianceRecordCounts.go.html to see an example of how to use SummarizeComplianceRecordCounts API.
// A default retry strategy applies to this operation SummarizeComplianceRecordCounts()
func (client FleetAppsManagementOperationsClient) SummarizeComplianceRecordCounts(ctx context.Context, request SummarizeComplianceRecordCountsRequest) (response SummarizeComplianceRecordCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeComplianceRecordCounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeComplianceRecordCountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeComplianceRecordCountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeComplianceRecordCountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeComplianceRecordCountsResponse")
	}
	return
}

// summarizeComplianceRecordCounts implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) summarizeComplianceRecordCounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/complianceRecordCounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeComplianceRecordCountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/ComplianceRecordAggregationCollection/SummarizeComplianceRecordCounts"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "SummarizeComplianceRecordCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeManagedEntityCounts Retrieve  aggregated summary information of Managed Entities within a Tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/SummarizeManagedEntityCounts.go.html to see an example of how to use SummarizeManagedEntityCounts API.
// A default retry strategy applies to this operation SummarizeManagedEntityCounts()
func (client FleetAppsManagementOperationsClient) SummarizeManagedEntityCounts(ctx context.Context, request SummarizeManagedEntityCountsRequest) (response SummarizeManagedEntityCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeManagedEntityCounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeManagedEntityCountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeManagedEntityCountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeManagedEntityCountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeManagedEntityCountsResponse")
	}
	return
}

// summarizeManagedEntityCounts implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) summarizeManagedEntityCounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedEntityCounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeManagedEntityCountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/ManagedEntityAggregationCollection/SummarizeManagedEntityCounts"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "SummarizeManagedEntityCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSchedulerJobCounts Retrieve aggregated summary information of Scheduler Jobs within a Tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/SummarizeSchedulerJobCounts.go.html to see an example of how to use SummarizeSchedulerJobCounts API.
// A default retry strategy applies to this operation SummarizeSchedulerJobCounts()
func (client FleetAppsManagementOperationsClient) SummarizeSchedulerJobCounts(ctx context.Context, request SummarizeSchedulerJobCountsRequest) (response SummarizeSchedulerJobCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSchedulerJobCounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSchedulerJobCountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSchedulerJobCountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSchedulerJobCountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSchedulerJobCountsResponse")
	}
	return
}

// summarizeSchedulerJobCounts implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) summarizeSchedulerJobCounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedulerJobCounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSchedulerJobCountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerJobAggregationCollection/SummarizeSchedulerJobCounts"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "SummarizeSchedulerJobCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePatch Updates the Patch
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdatePatch.go.html to see an example of how to use UpdatePatch API.
// A default retry strategy applies to this operation UpdatePatch()
func (client FleetAppsManagementOperationsClient) UpdatePatch(ctx context.Context, request UpdatePatchRequest) (response UpdatePatchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePatch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePatchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePatchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePatchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePatchResponse")
	}
	return
}

// updatePatch implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) updatePatch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/patches/{patchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePatchResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Patch/UpdatePatch"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "UpdatePatch", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSchedulerDefinition Update the details of a SchedulerDefinition that performs lifecycle management operations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateSchedulerDefinition.go.html to see an example of how to use UpdateSchedulerDefinition API.
// A default retry strategy applies to this operation UpdateSchedulerDefinition()
func (client FleetAppsManagementOperationsClient) UpdateSchedulerDefinition(ctx context.Context, request UpdateSchedulerDefinitionRequest) (response UpdateSchedulerDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSchedulerDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSchedulerDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSchedulerDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSchedulerDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSchedulerDefinitionResponse")
	}
	return
}

// updateSchedulerDefinition implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) updateSchedulerDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/schedulerDefinitions/{schedulerDefinitionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSchedulerDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerDefinition/UpdateSchedulerDefinition"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "UpdateSchedulerDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSchedulerJob Update a lifecycle operation job schedule in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateSchedulerJob.go.html to see an example of how to use UpdateSchedulerJob API.
// A default retry strategy applies to this operation UpdateSchedulerJob()
func (client FleetAppsManagementOperationsClient) UpdateSchedulerJob(ctx context.Context, request UpdateSchedulerJobRequest) (response UpdateSchedulerJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSchedulerJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSchedulerJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSchedulerJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSchedulerJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSchedulerJobResponse")
	}
	return
}

// updateSchedulerJob implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementOperationsClient) updateSchedulerJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/schedulerJobs/{schedulerJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSchedulerJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/SchedulerJob/UpdateSchedulerJob"
		err = common.PostProcessServiceError(err, "FleetAppsManagementOperations", "UpdateSchedulerJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
