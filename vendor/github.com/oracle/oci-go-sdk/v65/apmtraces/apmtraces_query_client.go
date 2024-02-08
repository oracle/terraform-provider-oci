// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// QueryClient a client for Query
type QueryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewQueryClientWithConfigurationProvider Creates a new default Query client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewQueryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client QueryClient, err error) {
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
	return newQueryClientFromBaseClient(baseClient, provider)
}

// NewQueryClientWithOboToken Creates a new default Query client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewQueryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client QueryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newQueryClientFromBaseClient(baseClient, configProvider)
}

func newQueryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client QueryClient, err error) {
	// Query service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Query"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = QueryClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *QueryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmtraces", "https://apm-trace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *QueryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *QueryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListQuickPicks Returns a list of predefined Quick Pick queries intended to assist the user
// to choose a query to run.  There is no sorting applied on the results.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/ListQuickPicks.go.html to see an example of how to use ListQuickPicks API.
func (client QueryClient) ListQuickPicks(ctx context.Context, request ListQuickPicksRequest) (response ListQuickPicksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQuickPicks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListQuickPicksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListQuickPicksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQuickPicksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQuickPicksResponse")
	}
	return
}

// listQuickPicks implements the OCIOperation interface (enables retrying operations)
func (client QueryClient) listQuickPicks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/queries/quickPicks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListQuickPicksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/QuickPickSummary/ListQuickPicks"
		err = common.PostProcessServiceError(err, "Query", "ListQuickPicks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Query Retrieves the results (selected attributes and aggregations) of a query constructed according to the Application Performance Monitoring Defined Query Syntax.
// Query results are filtered by the filter criteria specified in the where clause.
// Further query results are grouped by the attributes specified in the group by clause.  Finally,
// ordering (asc/desc) is done by the specified attributes in the order by clause.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/Query.go.html to see an example of how to use Query API.
func (client QueryClient) Query(ctx context.Context, request QueryRequest) (response QueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.query, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryResponse")
	}
	return
}

// query implements the OCIOperation interface (enables retrying operations)
func (client QueryClient) query(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/queries/actions/runQuery", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/QueryResultResponse/Query"
		err = common.PostProcessServiceError(err, "Query", "Query", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
