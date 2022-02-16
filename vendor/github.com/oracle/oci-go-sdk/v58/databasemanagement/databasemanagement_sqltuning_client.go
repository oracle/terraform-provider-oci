// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//SqlTuningClient a client for SqlTuning
type SqlTuningClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSqlTuningClientWithConfigurationProvider Creates a new default SqlTuning client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSqlTuningClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SqlTuningClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newSqlTuningClientFromBaseClient(baseClient, provider)
}

// NewSqlTuningClientWithOboToken Creates a new default SqlTuning client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewSqlTuningClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SqlTuningClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSqlTuningClientFromBaseClient(baseClient, configProvider)
}

func newSqlTuningClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SqlTuningClient, err error) {
	// SqlTuning service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SqlTuningClient{BaseClient: baseClient}
	client.BasePath = "20201101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SqlTuningClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemanagement", "https://dbmgmt.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SqlTuningClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SqlTuningClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CloneSqlTuningTask Clones and runs a SQL tuning task in the database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CloneSqlTuningTask.go.html to see an example of how to use CloneSqlTuningTask API.
func (client SqlTuningClient) CloneSqlTuningTask(ctx context.Context, request CloneSqlTuningTaskRequest) (response CloneSqlTuningTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.cloneSqlTuningTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CloneSqlTuningTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CloneSqlTuningTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CloneSqlTuningTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CloneSqlTuningTaskResponse")
	}
	return
}

// cloneSqlTuningTask implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) cloneSqlTuningTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/cloneSqlTuningTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CloneSqlTuningTaskResponse
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

// DropSqlTuningTask Drops a SQL tuning task and its related results from the database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DropSqlTuningTask.go.html to see an example of how to use DropSqlTuningTask API.
func (client SqlTuningClient) DropSqlTuningTask(ctx context.Context, request DropSqlTuningTaskRequest) (response DropSqlTuningTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.dropSqlTuningTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DropSqlTuningTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DropSqlTuningTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DropSqlTuningTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DropSqlTuningTaskResponse")
	}
	return
}

// dropSqlTuningTask implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) dropSqlTuningTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/dropSqlTuningTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DropSqlTuningTaskResponse
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

// GetExecutionPlanStatsComparision Retrieves a comparison of the existing SQL execution plan and a new plan.
// A SQL tuning task may suggest a new execution plan for a SQL,
// and this API retrieves the comparison report of the statistics of the two plans.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetExecutionPlanStatsComparision.go.html to see an example of how to use GetExecutionPlanStatsComparision API.
func (client SqlTuningClient) GetExecutionPlanStatsComparision(ctx context.Context, request GetExecutionPlanStatsComparisionRequest) (response GetExecutionPlanStatsComparisionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExecutionPlanStatsComparision, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExecutionPlanStatsComparisionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExecutionPlanStatsComparisionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExecutionPlanStatsComparisionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExecutionPlanStatsComparisionResponse")
	}
	return
}

// getExecutionPlanStatsComparision implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) getExecutionPlanStatsComparision(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlTuningAdvisorTasks/{sqlTuningAdvisorTaskId}/executionPlanStatsComparision", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExecutionPlanStatsComparisionResponse
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

// GetSqlExecutionPlan Retrieves a SQL execution plan for the SQL being tuned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetSqlExecutionPlan.go.html to see an example of how to use GetSqlExecutionPlan API.
func (client SqlTuningClient) GetSqlExecutionPlan(ctx context.Context, request GetSqlExecutionPlanRequest) (response GetSqlExecutionPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSqlExecutionPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSqlExecutionPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSqlExecutionPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSqlExecutionPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSqlExecutionPlanResponse")
	}
	return
}

// getSqlExecutionPlan implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) getSqlExecutionPlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlTuningAdvisorTasks/{sqlTuningAdvisorTaskId}/sqlExecutionPlan", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSqlExecutionPlanResponse
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

// GetSqlTuningAdvisorTaskSummaryReport Gets the summary report for the specified SQL Tuning Advisor task.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetSqlTuningAdvisorTaskSummaryReport.go.html to see an example of how to use GetSqlTuningAdvisorTaskSummaryReport API.
func (client SqlTuningClient) GetSqlTuningAdvisorTaskSummaryReport(ctx context.Context, request GetSqlTuningAdvisorTaskSummaryReportRequest) (response GetSqlTuningAdvisorTaskSummaryReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSqlTuningAdvisorTaskSummaryReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSqlTuningAdvisorTaskSummaryReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSqlTuningAdvisorTaskSummaryReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSqlTuningAdvisorTaskSummaryReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSqlTuningAdvisorTaskSummaryReportResponse")
	}
	return
}

// getSqlTuningAdvisorTaskSummaryReport implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) getSqlTuningAdvisorTaskSummaryReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlTuningAdvisorTasks/{sqlTuningAdvisorTaskId}/summaryReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSqlTuningAdvisorTaskSummaryReportResponse
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

// ListSqlTuningAdvisorTaskFindings Gets an array of the details of the findings that match specific filters.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningAdvisorTaskFindings.go.html to see an example of how to use ListSqlTuningAdvisorTaskFindings API.
func (client SqlTuningClient) ListSqlTuningAdvisorTaskFindings(ctx context.Context, request ListSqlTuningAdvisorTaskFindingsRequest) (response ListSqlTuningAdvisorTaskFindingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlTuningAdvisorTaskFindings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlTuningAdvisorTaskFindingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlTuningAdvisorTaskFindingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlTuningAdvisorTaskFindingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlTuningAdvisorTaskFindingsResponse")
	}
	return
}

// listSqlTuningAdvisorTaskFindings implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) listSqlTuningAdvisorTaskFindings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlTuningAdvisorTasks/{sqlTuningAdvisorTaskId}/findings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlTuningAdvisorTaskFindingsResponse
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

// ListSqlTuningAdvisorTaskRecommendations Gets the findings and possible actions for a given object in a SQL tuning task.
// The task ID and object ID are used to retrieve the findings and recommendations.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningAdvisorTaskRecommendations.go.html to see an example of how to use ListSqlTuningAdvisorTaskRecommendations API.
func (client SqlTuningClient) ListSqlTuningAdvisorTaskRecommendations(ctx context.Context, request ListSqlTuningAdvisorTaskRecommendationsRequest) (response ListSqlTuningAdvisorTaskRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlTuningAdvisorTaskRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlTuningAdvisorTaskRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlTuningAdvisorTaskRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlTuningAdvisorTaskRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlTuningAdvisorTaskRecommendationsResponse")
	}
	return
}

// listSqlTuningAdvisorTaskRecommendations implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) listSqlTuningAdvisorTaskRecommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlTuningAdvisorTasks/{sqlTuningAdvisorTaskId}/recommendations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlTuningAdvisorTaskRecommendationsResponse
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

// ListSqlTuningAdvisorTasks Lists the SQL Tuning Advisor tasks for the specified Managed Database.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningAdvisorTasks.go.html to see an example of how to use ListSqlTuningAdvisorTasks API.
func (client SqlTuningClient) ListSqlTuningAdvisorTasks(ctx context.Context, request ListSqlTuningAdvisorTasksRequest) (response ListSqlTuningAdvisorTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlTuningAdvisorTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlTuningAdvisorTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlTuningAdvisorTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlTuningAdvisorTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlTuningAdvisorTasksResponse")
	}
	return
}

// listSqlTuningAdvisorTasks implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) listSqlTuningAdvisorTasks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlTuningAdvisorTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlTuningAdvisorTasksResponse
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

// StartSqlTuningTask Starts a SQL tuning task for a given set of SQL statements from the active session history top SQL statements.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/StartSqlTuningTask.go.html to see an example of how to use StartSqlTuningTask API.
func (client SqlTuningClient) StartSqlTuningTask(ctx context.Context, request StartSqlTuningTaskRequest) (response StartSqlTuningTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startSqlTuningTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartSqlTuningTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartSqlTuningTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartSqlTuningTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartSqlTuningTaskResponse")
	}
	return
}

// startSqlTuningTask implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) startSqlTuningTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/startSqlTuningTask", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartSqlTuningTaskResponse
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
