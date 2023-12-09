// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// UsageapiClient a client for Usageapi
type UsageapiClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewUsageapiClientWithConfigurationProvider Creates a new default Usageapi client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewUsageapiClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client UsageapiClient, err error) {
	if enabled := common.CheckForEnabledServices("usageapi"); !enabled {
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
	return newUsageapiClientFromBaseClient(baseClient, provider)
}

// NewUsageapiClientWithOboToken Creates a new default Usageapi client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewUsageapiClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client UsageapiClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newUsageapiClientFromBaseClient(baseClient, configProvider)
}

func newUsageapiClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client UsageapiClient, err error) {
	// Usageapi service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Usageapi"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = UsageapiClient{BaseClient: baseClient}
	client.BasePath = "20200107"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *UsageapiClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("usageapi", "https://usageapi.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *UsageapiClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *UsageapiClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateCustomTable Returns the created custom table.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/CreateCustomTable.go.html to see an example of how to use CreateCustomTable API.
func (client UsageapiClient) CreateCustomTable(ctx context.Context, request CreateCustomTableRequest) (response CreateCustomTableResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCustomTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCustomTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCustomTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCustomTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCustomTableResponse")
	}
	return
}

// createCustomTable implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) createCustomTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/customTables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCustomTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/CustomTable/CreateCustomTable"
		err = common.PostProcessServiceError(err, "Usageapi", "CreateCustomTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateQuery Returns the created query.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/CreateQuery.go.html to see an example of how to use CreateQuery API.
func (client UsageapiClient) CreateQuery(ctx context.Context, request CreateQueryRequest) (response CreateQueryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateQueryResponse")
	}
	return
}

// createQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) createQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Query/CreateQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "CreateQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSchedule Returns the created schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/CreateSchedule.go.html to see an example of how to use CreateSchedule API.
func (client UsageapiClient) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (response CreateScheduleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScheduleResponse")
	}
	return
}

// createSchedule implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) createSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/schedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Schedule/CreateSchedule"
		err = common.PostProcessServiceError(err, "Usageapi", "CreateSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUsageCarbonEmissionsQuery Returns the created usage carbon emissions query.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/CreateUsageCarbonEmissionsQuery.go.html to see an example of how to use CreateUsageCarbonEmissionsQuery API.
func (client UsageapiClient) CreateUsageCarbonEmissionsQuery(ctx context.Context, request CreateUsageCarbonEmissionsQueryRequest) (response CreateUsageCarbonEmissionsQueryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createUsageCarbonEmissionsQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUsageCarbonEmissionsQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUsageCarbonEmissionsQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUsageCarbonEmissionsQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUsageCarbonEmissionsQueryResponse")
	}
	return
}

// createUsageCarbonEmissionsQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) createUsageCarbonEmissionsQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/usageCarbonEmissionsQueries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUsageCarbonEmissionsQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/UsageCarbonEmissionsQuery/CreateUsageCarbonEmissionsQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "CreateUsageCarbonEmissionsQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCustomTable Delete a saved custom table by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/DeleteCustomTable.go.html to see an example of how to use DeleteCustomTable API.
func (client UsageapiClient) DeleteCustomTable(ctx context.Context, request DeleteCustomTableRequest) (response DeleteCustomTableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCustomTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCustomTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCustomTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCustomTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCustomTableResponse")
	}
	return
}

// deleteCustomTable implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) deleteCustomTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/customTables/{customTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCustomTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/CustomTable/DeleteCustomTable"
		err = common.PostProcessServiceError(err, "Usageapi", "DeleteCustomTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteQuery Delete a saved query by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/DeleteQuery.go.html to see an example of how to use DeleteQuery API.
func (client UsageapiClient) DeleteQuery(ctx context.Context, request DeleteQueryRequest) (response DeleteQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteQueryResponse")
	}
	return
}

// deleteQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) deleteQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/queries/{queryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Query/DeleteQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "DeleteQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSchedule Delete a saved scheduled report by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/DeleteSchedule.go.html to see an example of how to use DeleteSchedule API.
func (client UsageapiClient) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) (response DeleteScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScheduleResponse")
	}
	return
}

// deleteSchedule implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) deleteSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/schedules/{scheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Schedule/DeleteSchedule"
		err = common.PostProcessServiceError(err, "Usageapi", "DeleteSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUsageCarbonEmissionsQuery Delete a usage carbon emissions saved query by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/DeleteUsageCarbonEmissionsQuery.go.html to see an example of how to use DeleteUsageCarbonEmissionsQuery API.
func (client UsageapiClient) DeleteUsageCarbonEmissionsQuery(ctx context.Context, request DeleteUsageCarbonEmissionsQueryRequest) (response DeleteUsageCarbonEmissionsQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUsageCarbonEmissionsQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUsageCarbonEmissionsQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUsageCarbonEmissionsQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUsageCarbonEmissionsQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUsageCarbonEmissionsQueryResponse")
	}
	return
}

// deleteUsageCarbonEmissionsQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) deleteUsageCarbonEmissionsQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/usageCarbonEmissionsQueries/{usageCarbonEmissionsQueryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUsageCarbonEmissionsQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/UsageCarbonEmissionsQuery/DeleteUsageCarbonEmissionsQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "DeleteUsageCarbonEmissionsQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCustomTable Returns the saved custom table.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/GetCustomTable.go.html to see an example of how to use GetCustomTable API.
func (client UsageapiClient) GetCustomTable(ctx context.Context, request GetCustomTableRequest) (response GetCustomTableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCustomTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCustomTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCustomTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCustomTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCustomTableResponse")
	}
	return
}

// getCustomTable implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) getCustomTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/customTables/{customTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCustomTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/CustomTable/GetCustomTable"
		err = common.PostProcessServiceError(err, "Usageapi", "GetCustomTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetQuery Returns the saved query.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/GetQuery.go.html to see an example of how to use GetQuery API.
func (client UsageapiClient) GetQuery(ctx context.Context, request GetQueryRequest) (response GetQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQueryResponse")
	}
	return
}

// getQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) getQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queries/{queryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Query/GetQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "GetQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSchedule Returns the saved schedule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/GetSchedule.go.html to see an example of how to use GetSchedule API.
func (client UsageapiClient) GetSchedule(ctx context.Context, request GetScheduleRequest) (response GetScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduleResponse")
	}
	return
}

// getSchedule implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) getSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedules/{scheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Schedule/GetSchedule"
		err = common.PostProcessServiceError(err, "Usageapi", "GetSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScheduledRun Returns the saved schedule run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/GetScheduledRun.go.html to see an example of how to use GetScheduledRun API.
func (client UsageapiClient) GetScheduledRun(ctx context.Context, request GetScheduledRunRequest) (response GetScheduledRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScheduledRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledRunResponse")
	}
	return
}

// getScheduledRun implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) getScheduledRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledRuns/{scheduledRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduledRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/ScheduledRun/GetScheduledRun"
		err = common.PostProcessServiceError(err, "Usageapi", "GetScheduledRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUsageCarbonEmissionsQuery Returns the usage carbon emissions saved query.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/GetUsageCarbonEmissionsQuery.go.html to see an example of how to use GetUsageCarbonEmissionsQuery API.
func (client UsageapiClient) GetUsageCarbonEmissionsQuery(ctx context.Context, request GetUsageCarbonEmissionsQueryRequest) (response GetUsageCarbonEmissionsQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUsageCarbonEmissionsQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUsageCarbonEmissionsQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUsageCarbonEmissionsQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUsageCarbonEmissionsQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUsageCarbonEmissionsQueryResponse")
	}
	return
}

// getUsageCarbonEmissionsQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) getUsageCarbonEmissionsQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usageCarbonEmissionsQueries/{usageCarbonEmissionsQueryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUsageCarbonEmissionsQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/UsageCarbonEmissionsQuery/GetUsageCarbonEmissionsQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "GetUsageCarbonEmissionsQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCustomTables Returns the saved custom table list.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListCustomTables.go.html to see an example of how to use ListCustomTables API.
func (client UsageapiClient) ListCustomTables(ctx context.Context, request ListCustomTablesRequest) (response ListCustomTablesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCustomTables, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCustomTablesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCustomTablesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCustomTablesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCustomTablesResponse")
	}
	return
}

// listCustomTables implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) listCustomTables(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/customTables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCustomTablesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/CustomTable/ListCustomTables"
		err = common.PostProcessServiceError(err, "Usageapi", "ListCustomTables", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListQueries Returns the saved query list.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListQueries.go.html to see an example of how to use ListQueries API.
func (client UsageapiClient) ListQueries(ctx context.Context, request ListQueriesRequest) (response ListQueriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQueries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListQueriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListQueriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQueriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQueriesResponse")
	}
	return
}

// listQueries implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) listQueries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListQueriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Query/ListQueries"
		err = common.PostProcessServiceError(err, "Usageapi", "ListQueries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledRuns Returns schedule history list.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListScheduledRuns.go.html to see an example of how to use ListScheduledRuns API.
func (client UsageapiClient) ListScheduledRuns(ctx context.Context, request ListScheduledRunsRequest) (response ListScheduledRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledRunsResponse")
	}
	return
}

// listScheduledRuns implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) listScheduledRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/ScheduledRun/ListScheduledRuns"
		err = common.PostProcessServiceError(err, "Usageapi", "ListScheduledRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSchedules Returns the saved schedule list.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListSchedules.go.html to see an example of how to use ListSchedules API.
func (client UsageapiClient) ListSchedules(ctx context.Context, request ListSchedulesRequest) (response ListSchedulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSchedules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSchedulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSchedulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSchedulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSchedulesResponse")
	}
	return
}

// listSchedules implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) listSchedules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/schedules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSchedulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Schedule/ListSchedules"
		err = common.PostProcessServiceError(err, "Usageapi", "ListSchedules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUsageCarbonEmissionsQueries Returns the usage carbon emissions saved query list.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListUsageCarbonEmissionsQueries.go.html to see an example of how to use ListUsageCarbonEmissionsQueries API.
func (client UsageapiClient) ListUsageCarbonEmissionsQueries(ctx context.Context, request ListUsageCarbonEmissionsQueriesRequest) (response ListUsageCarbonEmissionsQueriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUsageCarbonEmissionsQueries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsageCarbonEmissionsQueriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsageCarbonEmissionsQueriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsageCarbonEmissionsQueriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsageCarbonEmissionsQueriesResponse")
	}
	return
}

// listUsageCarbonEmissionsQueries implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) listUsageCarbonEmissionsQueries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usageCarbonEmissionsQueries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUsageCarbonEmissionsQueriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/UsageCarbonEmissionsQuery/ListUsageCarbonEmissionsQueries"
		err = common.PostProcessServiceError(err, "Usageapi", "ListUsageCarbonEmissionsQueries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestAverageCarbonEmission Returns the average carbon emissions summary by SKU.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/RequestAverageCarbonEmission.go.html to see an example of how to use RequestAverageCarbonEmission API.
func (client UsageapiClient) RequestAverageCarbonEmission(ctx context.Context, request RequestAverageCarbonEmissionRequest) (response RequestAverageCarbonEmissionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestAverageCarbonEmission, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestAverageCarbonEmissionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestAverageCarbonEmissionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestAverageCarbonEmissionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestAverageCarbonEmissionResponse")
	}
	return
}

// requestAverageCarbonEmission implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestAverageCarbonEmission(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/averageCarbonEmissions/{skuPartNumber}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestAverageCarbonEmissionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/AverageCarbonEmission/RequestAverageCarbonEmission"
		err = common.PostProcessServiceError(err, "Usageapi", "RequestAverageCarbonEmission", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestCleanEnergyUsage Returns the clean energy usage summary by region.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/RequestCleanEnergyUsage.go.html to see an example of how to use RequestCleanEnergyUsage API.
func (client UsageapiClient) RequestCleanEnergyUsage(ctx context.Context, request RequestCleanEnergyUsageRequest) (response RequestCleanEnergyUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestCleanEnergyUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestCleanEnergyUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestCleanEnergyUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestCleanEnergyUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestCleanEnergyUsageResponse")
	}
	return
}

// requestCleanEnergyUsage implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestCleanEnergyUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/cleanEnergyUsages/{region}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestCleanEnergyUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/CleanEnergyUsage/RequestCleanEnergyUsage"
		err = common.PostProcessServiceError(err, "Usageapi", "RequestCleanEnergyUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedConfigurations Returns the configurations list for the UI drop-down list.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/RequestSummarizedConfigurations.go.html to see an example of how to use RequestSummarizedConfigurations API.
func (client UsageapiClient) RequestSummarizedConfigurations(ctx context.Context, request RequestSummarizedConfigurationsRequest) (response RequestSummarizedConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedConfigurationsResponse")
	}
	return
}

// requestSummarizedConfigurations implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestSummarizedConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Configuration/RequestSummarizedConfigurations"
		err = common.PostProcessServiceError(err, "Usageapi", "RequestSummarizedConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedUsages Returns usage for the given account.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/RequestSummarizedUsages.go.html to see an example of how to use RequestSummarizedUsages API.
func (client UsageapiClient) RequestSummarizedUsages(ctx context.Context, request RequestSummarizedUsagesRequest) (response RequestSummarizedUsagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedUsages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedUsagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedUsagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedUsagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedUsagesResponse")
	}
	return
}

// requestSummarizedUsages implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestSummarizedUsages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/usage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedUsagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/UsageSummary/RequestSummarizedUsages"
		err = common.PostProcessServiceError(err, "Usageapi", "RequestSummarizedUsages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestUsageCarbonEmissionConfig Returns the configuration list for the UI drop-down list of carbon emission console.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/RequestUsageCarbonEmissionConfig.go.html to see an example of how to use RequestUsageCarbonEmissionConfig API.
func (client UsageapiClient) RequestUsageCarbonEmissionConfig(ctx context.Context, request RequestUsageCarbonEmissionConfigRequest) (response RequestUsageCarbonEmissionConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestUsageCarbonEmissionConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestUsageCarbonEmissionConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestUsageCarbonEmissionConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestUsageCarbonEmissionConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestUsageCarbonEmissionConfigResponse")
	}
	return
}

// requestUsageCarbonEmissionConfig implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestUsageCarbonEmissionConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usageCarbonEmissionsConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestUsageCarbonEmissionConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Configuration/RequestUsageCarbonEmissionConfig"
		err = common.PostProcessServiceError(err, "Usageapi", "RequestUsageCarbonEmissionConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestUsageCarbonEmissions Returns usage carbon emission for the given account.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/RequestUsageCarbonEmissions.go.html to see an example of how to use RequestUsageCarbonEmissions API.
func (client UsageapiClient) RequestUsageCarbonEmissions(ctx context.Context, request RequestUsageCarbonEmissionsRequest) (response RequestUsageCarbonEmissionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestUsageCarbonEmissions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestUsageCarbonEmissionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestUsageCarbonEmissionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestUsageCarbonEmissionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestUsageCarbonEmissionsResponse")
	}
	return
}

// requestUsageCarbonEmissions implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) requestUsageCarbonEmissions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/usageCarbonEmissions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestUsageCarbonEmissionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/UsageCarbonEmissionSummary/RequestUsageCarbonEmissions"
		err = common.PostProcessServiceError(err, "Usageapi", "RequestUsageCarbonEmissions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCustomTable Update a saved custom table by table id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/UpdateCustomTable.go.html to see an example of how to use UpdateCustomTable API.
func (client UsageapiClient) UpdateCustomTable(ctx context.Context, request UpdateCustomTableRequest) (response UpdateCustomTableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCustomTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCustomTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCustomTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCustomTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCustomTableResponse")
	}
	return
}

// updateCustomTable implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) updateCustomTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/customTables/{customTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCustomTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/CustomTable/UpdateCustomTable"
		err = common.PostProcessServiceError(err, "Usageapi", "UpdateCustomTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateQuery Update a saved query by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/UpdateQuery.go.html to see an example of how to use UpdateQuery API.
func (client UsageapiClient) UpdateQuery(ctx context.Context, request UpdateQueryRequest) (response UpdateQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateQueryResponse")
	}
	return
}

// updateQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) updateQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/queries/{queryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Query/UpdateQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "UpdateQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSchedule Update a saved schedule
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/UpdateSchedule.go.html to see an example of how to use UpdateSchedule API.
func (client UsageapiClient) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (response UpdateScheduleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSchedule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScheduleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScheduleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScheduleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScheduleResponse")
	}
	return
}

// updateSchedule implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) updateSchedule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/schedules/{scheduleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScheduleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/Schedule/UpdateSchedule"
		err = common.PostProcessServiceError(err, "Usageapi", "UpdateSchedule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateUsageCarbonEmissionsQuery Update a usage carbon emissions saved query by the OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/UpdateUsageCarbonEmissionsQuery.go.html to see an example of how to use UpdateUsageCarbonEmissionsQuery API.
func (client UsageapiClient) UpdateUsageCarbonEmissionsQuery(ctx context.Context, request UpdateUsageCarbonEmissionsQueryRequest) (response UpdateUsageCarbonEmissionsQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUsageCarbonEmissionsQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUsageCarbonEmissionsQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUsageCarbonEmissionsQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUsageCarbonEmissionsQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUsageCarbonEmissionsQueryResponse")
	}
	return
}

// updateUsageCarbonEmissionsQuery implements the OCIOperation interface (enables retrying operations)
func (client UsageapiClient) updateUsageCarbonEmissionsQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/usageCarbonEmissionsQueries/{usageCarbonEmissionsQueryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUsageCarbonEmissionsQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage/20200107/UsageCarbonEmissionsQuery/UpdateUsageCarbonEmissionsQuery"
		err = common.PostProcessServiceError(err, "Usageapi", "UpdateUsageCarbonEmissionsQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
