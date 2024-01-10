// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// SqlTuningClient a client for SqlTuning
type SqlTuningClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSqlTuningClientWithConfigurationProvider Creates a new default SqlTuning client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSqlTuningClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SqlTuningClient, err error) {
	if enabled := common.CheckForEnabledServices("databasemanagement"); !enabled {
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
	return newSqlTuningClientFromBaseClient(baseClient, provider)
}

// NewSqlTuningClientWithOboToken Creates a new default SqlTuning client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewSqlTuningClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SqlTuningClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSqlTuningClientFromBaseClient(baseClient, configProvider)
}

func newSqlTuningClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SqlTuningClient, err error) {
	// SqlTuning service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("SqlTuning"))
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
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *SqlTuningClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CloneSqlTuningTask Clones and runs a SQL tuning task in the database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/CloneSqlTuningTask"
		err = common.PostProcessServiceError(err, "SqlTuning", "CloneSqlTuningTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSqlTuningSet Creates an empty Sql tuning set within the Managed Database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/CreateSqlTuningSet.go.html to see an example of how to use CreateSqlTuningSet API.
// A default retry strategy applies to this operation CreateSqlTuningSet()
func (client SqlTuningClient) CreateSqlTuningSet(ctx context.Context, request CreateSqlTuningSetRequest) (response CreateSqlTuningSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSqlTuningSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSqlTuningSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSqlTuningSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSqlTuningSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSqlTuningSetResponse")
	}
	return
}

// createSqlTuningSet implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) createSqlTuningSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlTuningSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSqlTuningSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SqlTuningSet/CreateSqlTuningSet"
		err = common.PostProcessServiceError(err, "SqlTuning", "CreateSqlTuningSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DropSqlTuningSet Drops the Sql tuning set specified by sqlTuningSet within the Managed Database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DropSqlTuningSet.go.html to see an example of how to use DropSqlTuningSet API.
// A default retry strategy applies to this operation DropSqlTuningSet()
func (client SqlTuningClient) DropSqlTuningSet(ctx context.Context, request DropSqlTuningSetRequest) (response DropSqlTuningSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.dropSqlTuningSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DropSqlTuningSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DropSqlTuningSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DropSqlTuningSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DropSqlTuningSetResponse")
	}
	return
}

// dropSqlTuningSet implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) dropSqlTuningSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlTuningSets/{sqlTuningSetId}/actions/dropSqlTuningSet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DropSqlTuningSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SqlTuningSet/DropSqlTuningSet"
		err = common.PostProcessServiceError(err, "SqlTuning", "DropSqlTuningSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DropSqlTuningTask Drops a SQL tuning task and its related results from the database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/DropSqlTuningTask"
		err = common.PostProcessServiceError(err, "SqlTuning", "DropSqlTuningTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DropSqlsInSqlTuningSet Deletes the Sqls in the specified Sql tuning set that matches the filter criteria provided in the basicFilter.
// If basicFilter criteria is not provided, then entire Sqls in the Sql tuning set is deleted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/DropSqlsInSqlTuningSet.go.html to see an example of how to use DropSqlsInSqlTuningSet API.
// A default retry strategy applies to this operation DropSqlsInSqlTuningSet()
func (client SqlTuningClient) DropSqlsInSqlTuningSet(ctx context.Context, request DropSqlsInSqlTuningSetRequest) (response DropSqlsInSqlTuningSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.dropSqlsInSqlTuningSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DropSqlsInSqlTuningSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DropSqlsInSqlTuningSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DropSqlsInSqlTuningSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DropSqlsInSqlTuningSetResponse")
	}
	return
}

// dropSqlsInSqlTuningSet implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) dropSqlsInSqlTuningSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlTuningSets/{sqlTuningSetId}/actions/dropSqlsInSqlTuningSet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DropSqlsInSqlTuningSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SqlTuningSet/DropSqlsInSqlTuningSet"
		err = common.PostProcessServiceError(err, "SqlTuning", "DropSqlsInSqlTuningSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FetchSqlTuningSet Fetch the details of Sql statements in the Sql tuning set specified by name, owner and optional filter parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/FetchSqlTuningSet.go.html to see an example of how to use FetchSqlTuningSet API.
// A default retry strategy applies to this operation FetchSqlTuningSet()
func (client SqlTuningClient) FetchSqlTuningSet(ctx context.Context, request FetchSqlTuningSetRequest) (response FetchSqlTuningSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.fetchSqlTuningSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FetchSqlTuningSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FetchSqlTuningSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FetchSqlTuningSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FetchSqlTuningSetResponse")
	}
	return
}

// fetchSqlTuningSet implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) fetchSqlTuningSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlTuningSets/{sqlTuningSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FetchSqlTuningSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SqlTuningSet/FetchSqlTuningSet"
		err = common.PostProcessServiceError(err, "SqlTuning", "FetchSqlTuningSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExecutionPlanStatsComparision Retrieves a comparison of the existing SQL execution plan and a new plan.
// A SQL tuning task may suggest a new execution plan for a SQL,
// and this API retrieves the comparison report of the statistics of the two plans.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetExecutionPlanStatsComparision"
		err = common.PostProcessServiceError(err, "SqlTuning", "GetExecutionPlanStatsComparision", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSqlExecutionPlan Retrieves a SQL execution plan for the SQL being tuned.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetSqlExecutionPlan"
		err = common.PostProcessServiceError(err, "SqlTuning", "GetSqlExecutionPlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSqlTuningAdvisorTaskSummaryReport Gets the summary report for the specified SQL Tuning Advisor task.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/GetSqlTuningAdvisorTaskSummaryReport"
		err = common.PostProcessServiceError(err, "SqlTuning", "GetSqlTuningAdvisorTaskSummaryReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlTuningAdvisorTaskFindings Gets an array of the details of the findings that match specific filters.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListSqlTuningAdvisorTaskFindings"
		err = common.PostProcessServiceError(err, "SqlTuning", "ListSqlTuningAdvisorTaskFindings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlTuningAdvisorTaskRecommendations Gets the findings and possible actions for a given object in a SQL tuning task.
// The task ID and object ID are used to retrieve the findings and recommendations.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListSqlTuningAdvisorTaskRecommendations"
		err = common.PostProcessServiceError(err, "SqlTuning", "ListSqlTuningAdvisorTaskRecommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlTuningAdvisorTasks Lists the SQL Tuning Advisor tasks for the specified Managed Database.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListSqlTuningAdvisorTasks"
		err = common.PostProcessServiceError(err, "SqlTuning", "ListSqlTuningAdvisorTasks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlTuningSets Lists the SQL tuning sets for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningSets.go.html to see an example of how to use ListSqlTuningSets API.
func (client SqlTuningClient) ListSqlTuningSets(ctx context.Context, request ListSqlTuningSetsRequest) (response ListSqlTuningSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlTuningSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlTuningSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlTuningSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlTuningSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlTuningSetsResponse")
	}
	return
}

// listSqlTuningSets implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) listSqlTuningSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/sqlTuningSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlTuningSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListSqlTuningSets"
		err = common.PostProcessServiceError(err, "SqlTuning", "ListSqlTuningSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LoadSqlTuningSet Load Sql statements into the Sql tuning set specified by name and optional filter parameters within the Managed Database specified by managedDatabaseId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/LoadSqlTuningSet.go.html to see an example of how to use LoadSqlTuningSet API.
// A default retry strategy applies to this operation LoadSqlTuningSet()
func (client SqlTuningClient) LoadSqlTuningSet(ctx context.Context, request LoadSqlTuningSetRequest) (response LoadSqlTuningSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.loadSqlTuningSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LoadSqlTuningSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LoadSqlTuningSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LoadSqlTuningSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LoadSqlTuningSetResponse")
	}
	return
}

// loadSqlTuningSet implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) loadSqlTuningSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlTuningSets/{sqlTuningSetId}/actions/loadSqlTuningSet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LoadSqlTuningSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SqlTuningSet/LoadSqlTuningSet"
		err = common.PostProcessServiceError(err, "SqlTuning", "LoadSqlTuningSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SaveSqlTuningSetAs Saves the specified list of Sqls statements into another new Sql tuning set or loads into an existing Sql tuning set'.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SaveSqlTuningSetAs.go.html to see an example of how to use SaveSqlTuningSetAs API.
// A default retry strategy applies to this operation SaveSqlTuningSetAs()
func (client SqlTuningClient) SaveSqlTuningSetAs(ctx context.Context, request SaveSqlTuningSetAsRequest) (response SaveSqlTuningSetAsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.saveSqlTuningSetAs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SaveSqlTuningSetAsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SaveSqlTuningSetAsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SaveSqlTuningSetAsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SaveSqlTuningSetAsResponse")
	}
	return
}

// saveSqlTuningSetAs implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) saveSqlTuningSetAs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlTuningSets/{sqlTuningSetId}/actions/saveAs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SaveSqlTuningSetAsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SqlTuningSet/SaveSqlTuningSetAs"
		err = common.PostProcessServiceError(err, "SqlTuning", "SaveSqlTuningSetAs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartSqlTuningTask Starts a SQL tuning task for a given set of SQL statements from the active session history top SQL statements.
//
// # See also
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/StartSqlTuningTask"
		err = common.PostProcessServiceError(err, "SqlTuning", "StartSqlTuningTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateBasicFilter Executes a SQL query to check whether user entered basic filter criteria is valid or not.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ValidateBasicFilter.go.html to see an example of how to use ValidateBasicFilter API.
// A default retry strategy applies to this operation ValidateBasicFilter()
func (client SqlTuningClient) ValidateBasicFilter(ctx context.Context, request ValidateBasicFilterRequest) (response ValidateBasicFilterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateBasicFilter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateBasicFilterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateBasicFilterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateBasicFilterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateBasicFilterResponse")
	}
	return
}

// validateBasicFilter implements the OCIOperation interface (enables retrying operations)
func (client SqlTuningClient) validateBasicFilter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/sqlTuningSets/{sqlTuningSetId}/actions/validateBasicFilter", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateBasicFilterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/SqlTuningSet/ValidateBasicFilter"
		err = common.PostProcessServiceError(err, "SqlTuning", "ValidateBasicFilter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
