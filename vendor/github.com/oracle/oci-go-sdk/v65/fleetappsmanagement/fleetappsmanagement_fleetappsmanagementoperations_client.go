// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
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

// CreateSchedulerDefinition Creates a new SchedulerDefinition.
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

// DeleteSchedulerJob Deletes a SchedulerJob resource by identifier
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

// GetExecution Gets a JobActivity by identifier
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

// GetJobActivity Gets a JobActivity by identifier
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

// GetSchedulerDefinition Gets a SchedulerDefinition by identifier
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

// GetSchedulerJob Gets a SchedulerJob by identifier
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

// ListScheduledFleets Returns a list of ScheduledFleets
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

// ListSchedulerDefinitions Returns a list of SchedulerDefinitions.
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

// ListSchedulerJobs Returns a list of SchedulerJobs.
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

// UpdateSchedulerDefinition Updates the SchedulerDefinition
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

// UpdateSchedulerJob Updates the SchedulerJob
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
