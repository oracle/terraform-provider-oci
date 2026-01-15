// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ScheduledQueryClient a client for ScheduledQuery
type ScheduledQueryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewScheduledQueryClientWithConfigurationProvider Creates a new default ScheduledQuery client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewScheduledQueryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ScheduledQueryClient, err error) {
	if enabled := common.CheckForEnabledServices("apmtraces"); !enabled {
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
	return newScheduledQueryClientFromBaseClient(baseClient, provider)
}

// NewScheduledQueryClientWithOboToken Creates a new default ScheduledQuery client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewScheduledQueryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ScheduledQueryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newScheduledQueryClientFromBaseClient(baseClient, configProvider)
}

func newScheduledQueryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ScheduledQueryClient, err error) {
	// ScheduledQuery service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ScheduledQuery"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ScheduledQueryClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ScheduledQueryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmtraces", "https://apm-trace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ScheduledQueryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ScheduledQueryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateScheduledQuery Create a scheduled query in the APM Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/CreateScheduledQuery.go.html to see an example of how to use CreateScheduledQuery API.
func (client ScheduledQueryClient) CreateScheduledQuery(ctx context.Context, request CreateScheduledQueryRequest) (response CreateScheduledQueryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createScheduledQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScheduledQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScheduledQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScheduledQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScheduledQueryResponse")
	}
	return
}

// createScheduledQuery implements the OCIOperation interface (enables retrying operations)
func (client ScheduledQueryClient) createScheduledQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scheduledQueries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScheduledQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/ScheduledQuery/CreateScheduledQuery"
		err = common.PostProcessServiceError(err, "ScheduledQuery", "CreateScheduledQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteScheduledQuery Delete a scheduled query in the APM Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/DeleteScheduledQuery.go.html to see an example of how to use DeleteScheduledQuery API.
func (client ScheduledQueryClient) DeleteScheduledQuery(ctx context.Context, request DeleteScheduledQueryRequest) (response DeleteScheduledQueryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteScheduledQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScheduledQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScheduledQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScheduledQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScheduledQueryResponse")
	}
	return
}

// deleteScheduledQuery implements the OCIOperation interface (enables retrying operations)
func (client ScheduledQueryClient) deleteScheduledQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/scheduledQueries/{scheduledQueryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScheduledQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/ScheduledQuery/DeleteScheduledQuery"
		err = common.PostProcessServiceError(err, "ScheduledQuery", "DeleteScheduledQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScheduledQuery Retrieve a scheduled query in the APM Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetScheduledQuery.go.html to see an example of how to use GetScheduledQuery API.
func (client ScheduledQueryClient) GetScheduledQuery(ctx context.Context, request GetScheduledQueryRequest) (response GetScheduledQueryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getScheduledQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledQueryResponse")
	}
	return
}

// getScheduledQuery implements the OCIOperation interface (enables retrying operations)
func (client ScheduledQueryClient) getScheduledQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledQueries/{scheduledQueryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduledQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/ScheduledQuery/GetScheduledQuery"
		err = common.PostProcessServiceError(err, "ScheduledQuery", "GetScheduledQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledQueries Returns a list of all scheduled queries in the APM Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/ListScheduledQueries.go.html to see an example of how to use ListScheduledQueries API.
func (client ScheduledQueryClient) ListScheduledQueries(ctx context.Context, request ListScheduledQueriesRequest) (response ListScheduledQueriesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listScheduledQueries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledQueriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledQueriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledQueriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledQueriesResponse")
	}
	return
}

// listScheduledQueries implements the OCIOperation interface (enables retrying operations)
func (client ScheduledQueryClient) listScheduledQueries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scheduledQueries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledQueriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/ScheduledQueryCollection/ListScheduledQueries"
		err = common.PostProcessServiceError(err, "ScheduledQuery", "ListScheduledQueries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateScheduledQuery Update a scheduled query in the APM Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/UpdateScheduledQuery.go.html to see an example of how to use UpdateScheduledQuery API.
func (client ScheduledQueryClient) UpdateScheduledQuery(ctx context.Context, request UpdateScheduledQueryRequest) (response UpdateScheduledQueryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateScheduledQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScheduledQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScheduledQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScheduledQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScheduledQueryResponse")
	}
	return
}

// updateScheduledQuery implements the OCIOperation interface (enables retrying operations)
func (client ScheduledQueryClient) updateScheduledQuery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/scheduledQueries/{scheduledQueryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScheduledQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/ScheduledQuery/UpdateScheduledQuery"
		err = common.PostProcessServiceError(err, "ScheduledQuery", "UpdateScheduledQuery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
