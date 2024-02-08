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

// DiagnosabilityClient a client for Diagnosability
type DiagnosabilityClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDiagnosabilityClientWithConfigurationProvider Creates a new default Diagnosability client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDiagnosabilityClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DiagnosabilityClient, err error) {
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
	return newDiagnosabilityClientFromBaseClient(baseClient, provider)
}

// NewDiagnosabilityClientWithOboToken Creates a new default Diagnosability client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDiagnosabilityClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DiagnosabilityClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDiagnosabilityClientFromBaseClient(baseClient, configProvider)
}

func newDiagnosabilityClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DiagnosabilityClient, err error) {
	// Diagnosability service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Diagnosability"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DiagnosabilityClient{BaseClient: baseClient}
	client.BasePath = "20201101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DiagnosabilityClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemanagement", "https://dbmgmt.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DiagnosabilityClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DiagnosabilityClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListAlertLogs Lists the alert logs for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAlertLogs.go.html to see an example of how to use ListAlertLogs API.
func (client DiagnosabilityClient) ListAlertLogs(ctx context.Context, request ListAlertLogsRequest) (response ListAlertLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlertLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlertLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlertLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlertLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlertLogsResponse")
	}
	return
}

// listAlertLogs implements the OCIOperation interface (enables retrying operations)
func (client DiagnosabilityClient) listAlertLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/alertLogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlertLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListAlertLogs"
		err = common.PostProcessServiceError(err, "Diagnosability", "ListAlertLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAttentionLogs Lists the attention logs for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAttentionLogs.go.html to see an example of how to use ListAttentionLogs API.
func (client DiagnosabilityClient) ListAttentionLogs(ctx context.Context, request ListAttentionLogsRequest) (response ListAttentionLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttentionLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAttentionLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAttentionLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttentionLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttentionLogsResponse")
	}
	return
}

// listAttentionLogs implements the OCIOperation interface (enables retrying operations)
func (client DiagnosabilityClient) listAttentionLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/attentionLogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAttentionLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ListAttentionLogs"
		err = common.PostProcessServiceError(err, "Diagnosability", "ListAttentionLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAlertLogCounts Get the counts of alert logs for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAlertLogCounts.go.html to see an example of how to use SummarizeAlertLogCounts API.
func (client DiagnosabilityClient) SummarizeAlertLogCounts(ctx context.Context, request SummarizeAlertLogCountsRequest) (response SummarizeAlertLogCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAlertLogCounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAlertLogCountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAlertLogCountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAlertLogCountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAlertLogCountsResponse")
	}
	return
}

// summarizeAlertLogCounts implements the OCIOperation interface (enables retrying operations)
func (client DiagnosabilityClient) summarizeAlertLogCounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/alertLogCounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAlertLogCountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAlertLogCounts"
		err = common.PostProcessServiceError(err, "Diagnosability", "SummarizeAlertLogCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAttentionLogCounts Get the counts of attention logs for the specified Managed Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAttentionLogCounts.go.html to see an example of how to use SummarizeAttentionLogCounts API.
func (client DiagnosabilityClient) SummarizeAttentionLogCounts(ctx context.Context, request SummarizeAttentionLogCountsRequest) (response SummarizeAttentionLogCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAttentionLogCounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAttentionLogCountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAttentionLogCountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAttentionLogCountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAttentionLogCountsResponse")
	}
	return
}

// summarizeAttentionLogCounts implements the OCIOperation interface (enables retrying operations)
func (client DiagnosabilityClient) summarizeAttentionLogCounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedDatabases/{managedDatabaseId}/attentionLogCounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAttentionLogCountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/SummarizeAttentionLogCounts"
		err = common.PostProcessServiceError(err, "Diagnosability", "SummarizeAttentionLogCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
