// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//OperationsInsightsClient a client for OperationsInsights
type OperationsInsightsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOperationsInsightsClientWithConfigurationProvider Creates a new default OperationsInsights client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOperationsInsightsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OperationsInsightsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newOperationsInsightsClientFromBaseClient(baseClient, provider)
}

// NewOperationsInsightsClientWithOboToken Creates a new default OperationsInsights client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewOperationsInsightsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OperationsInsightsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOperationsInsightsClientFromBaseClient(baseClient, configProvider)
}

func newOperationsInsightsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OperationsInsightsClient, err error) {
	// OperationsInsights service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OperationsInsightsClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OperationsInsightsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("opsi", "https://operationsinsights.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OperationsInsightsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OperationsInsightsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddExadataInsightMembers Add new members (e.g. databases and hosts) to an Exadata system in Operations Insights. Exadata-related metric collection and analysis will be started.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/AddExadataInsightMembers.go.html to see an example of how to use AddExadataInsightMembers API.
func (client OperationsInsightsClient) AddExadataInsightMembers(ctx context.Context, request AddExadataInsightMembersRequest) (response AddExadataInsightMembersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addExadataInsightMembers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddExadataInsightMembersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddExadataInsightMembersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddExadataInsightMembersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddExadataInsightMembersResponse")
	}
	return
}

// addExadataInsightMembers implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) addExadataInsightMembers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInsights/{exadataInsightId}/actions/addMembers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddExadataInsightMembersResponse
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

// ChangeDatabaseInsightCompartment Moves a DatabaseInsight resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeDatabaseInsightCompartment.go.html to see an example of how to use ChangeDatabaseInsightCompartment API.
func (client OperationsInsightsClient) ChangeDatabaseInsightCompartment(ctx context.Context, request ChangeDatabaseInsightCompartmentRequest) (response ChangeDatabaseInsightCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseInsightCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseInsightCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseInsightCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseInsightCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseInsightCompartmentResponse")
	}
	return
}

// changeDatabaseInsightCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeDatabaseInsightCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseInsightCompartmentResponse
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

// ChangeEnterpriseManagerBridgeCompartment Moves a EnterpriseManagerBridge resource from one compartment to another. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeEnterpriseManagerBridgeCompartment.go.html to see an example of how to use ChangeEnterpriseManagerBridgeCompartment API.
func (client OperationsInsightsClient) ChangeEnterpriseManagerBridgeCompartment(ctx context.Context, request ChangeEnterpriseManagerBridgeCompartmentRequest) (response ChangeEnterpriseManagerBridgeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeEnterpriseManagerBridgeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEnterpriseManagerBridgeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEnterpriseManagerBridgeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEnterpriseManagerBridgeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEnterpriseManagerBridgeCompartmentResponse")
	}
	return
}

// changeEnterpriseManagerBridgeCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeEnterpriseManagerBridgeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/enterpriseManagerBridges/{enterpriseManagerBridgeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEnterpriseManagerBridgeCompartmentResponse
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

// ChangeExadataInsightCompartment Moves an Exadata insight resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeExadataInsightCompartment.go.html to see an example of how to use ChangeExadataInsightCompartment API.
func (client OperationsInsightsClient) ChangeExadataInsightCompartment(ctx context.Context, request ChangeExadataInsightCompartmentRequest) (response ChangeExadataInsightCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeExadataInsightCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeExadataInsightCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeExadataInsightCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeExadataInsightCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeExadataInsightCompartmentResponse")
	}
	return
}

// changeExadataInsightCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeExadataInsightCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInsights/{exadataInsightId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeExadataInsightCompartmentResponse
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

// ChangeHostInsightCompartment Moves a HostInsight resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeHostInsightCompartment.go.html to see an example of how to use ChangeHostInsightCompartment API.
func (client OperationsInsightsClient) ChangeHostInsightCompartment(ctx context.Context, request ChangeHostInsightCompartmentRequest) (response ChangeHostInsightCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeHostInsightCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeHostInsightCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeHostInsightCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeHostInsightCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeHostInsightCompartmentResponse")
	}
	return
}

// changeHostInsightCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeHostInsightCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/hostInsights/{hostInsightId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeHostInsightCompartmentResponse
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

// CreateAwrHub Create a AWR hub resource for the tenant in Operations Insights.
// This resource will be created in root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateAwrHub.go.html to see an example of how to use CreateAwrHub API.
func (client OperationsInsightsClient) CreateAwrHub(ctx context.Context, request CreateAwrHubRequest) (response CreateAwrHubResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAwrHub, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAwrHubResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAwrHubResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAwrHubResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAwrHubResponse")
	}
	return
}

// createAwrHub implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createAwrHub(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/awrHubs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAwrHubResponse
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

// CreateDatabaseInsight Create a Database Insight resource for a database in Operations Insights. The database will be enabled in Operations Insights. Database metric collection and analysis will be started.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateDatabaseInsight.go.html to see an example of how to use CreateDatabaseInsight API.
func (client OperationsInsightsClient) CreateDatabaseInsight(ctx context.Context, request CreateDatabaseInsightRequest) (response CreateDatabaseInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseInsightResponse")
	}
	return
}

// createDatabaseInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createDatabaseInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseInsightResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databaseinsight{})
	return response, err
}

// CreateEnterpriseManagerBridge Create a Enterprise Manager bridge in Operations Insights.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateEnterpriseManagerBridge.go.html to see an example of how to use CreateEnterpriseManagerBridge API.
func (client OperationsInsightsClient) CreateEnterpriseManagerBridge(ctx context.Context, request CreateEnterpriseManagerBridgeRequest) (response CreateEnterpriseManagerBridgeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEnterpriseManagerBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEnterpriseManagerBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEnterpriseManagerBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEnterpriseManagerBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEnterpriseManagerBridgeResponse")
	}
	return
}

// createEnterpriseManagerBridge implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createEnterpriseManagerBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/enterpriseManagerBridges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEnterpriseManagerBridgeResponse
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

// CreateExadataInsight Create an Exadata insight resource for an Exadata system in Operations Insights. The Exadata system will be enabled in Operations Insights. Exadata-related metric collection and analysis will be started.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateExadataInsight.go.html to see an example of how to use CreateExadataInsight API.
func (client OperationsInsightsClient) CreateExadataInsight(ctx context.Context, request CreateExadataInsightRequest) (response CreateExadataInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createExadataInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateExadataInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateExadataInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateExadataInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateExadataInsightResponse")
	}
	return
}

// createExadataInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createExadataInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateExadataInsightResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &exadatainsight{})
	return response, err
}

// CreateHostInsight Create a Host Insight resource for a host in Operations Insights. The host will be enabled in Operations Insights. Host metric collection and analysis will be started.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateHostInsight.go.html to see an example of how to use CreateHostInsight API.
func (client OperationsInsightsClient) CreateHostInsight(ctx context.Context, request CreateHostInsightRequest) (response CreateHostInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createHostInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateHostInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateHostInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateHostInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateHostInsightResponse")
	}
	return
}

// createHostInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createHostInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/hostInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateHostInsightResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &hostinsight{})
	return response, err
}

// CreateOperationsInsightsWarehouse Create a Operations Insights Warehouse resource for the tenant in Operations Insights. New ADW will be provisioned for this tenant.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateOperationsInsightsWarehouse.go.html to see an example of how to use CreateOperationsInsightsWarehouse API.
func (client OperationsInsightsClient) CreateOperationsInsightsWarehouse(ctx context.Context, request CreateOperationsInsightsWarehouseRequest) (response CreateOperationsInsightsWarehouseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOperationsInsightsWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOperationsInsightsWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOperationsInsightsWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOperationsInsightsWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOperationsInsightsWarehouseResponse")
	}
	return
}

// createOperationsInsightsWarehouse implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createOperationsInsightsWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operationsInsightsWarehouses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOperationsInsightsWarehouseResponse
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

// CreateOperationsInsightsWarehouseUser Create a Operations Insights Warehouse user resource for the tenant in Operations Insights.
// This resource will be created in root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateOperationsInsightsWarehouseUser.go.html to see an example of how to use CreateOperationsInsightsWarehouseUser API.
func (client OperationsInsightsClient) CreateOperationsInsightsWarehouseUser(ctx context.Context, request CreateOperationsInsightsWarehouseUserRequest) (response CreateOperationsInsightsWarehouseUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOperationsInsightsWarehouseUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOperationsInsightsWarehouseUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOperationsInsightsWarehouseUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOperationsInsightsWarehouseUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOperationsInsightsWarehouseUserResponse")
	}
	return
}

// createOperationsInsightsWarehouseUser implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createOperationsInsightsWarehouseUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operationsInsightsWarehouseUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOperationsInsightsWarehouseUserResponse
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

// DeleteAwrHub Deletes an AWR hub.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteAwrHub.go.html to see an example of how to use DeleteAwrHub API.
func (client OperationsInsightsClient) DeleteAwrHub(ctx context.Context, request DeleteAwrHubRequest) (response DeleteAwrHubResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAwrHub, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAwrHubResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAwrHubResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAwrHubResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAwrHubResponse")
	}
	return
}

// deleteAwrHub implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteAwrHub(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/awrHubs/{awrHubId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAwrHubResponse
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

// DeleteDatabaseInsight Deletes a database insight. The database insight will be deleted and cannot be enabled again.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteDatabaseInsight.go.html to see an example of how to use DeleteDatabaseInsight API.
func (client OperationsInsightsClient) DeleteDatabaseInsight(ctx context.Context, request DeleteDatabaseInsightRequest) (response DeleteDatabaseInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseInsightResponse")
	}
	return
}

// deleteDatabaseInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteDatabaseInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseInsights/{databaseInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseInsightResponse
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

// DeleteEnterpriseManagerBridge Deletes an Operations Insights Enterprise Manager bridge. If any database insight is still referencing this bridge, the operation will fail.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteEnterpriseManagerBridge.go.html to see an example of how to use DeleteEnterpriseManagerBridge API.
func (client OperationsInsightsClient) DeleteEnterpriseManagerBridge(ctx context.Context, request DeleteEnterpriseManagerBridgeRequest) (response DeleteEnterpriseManagerBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEnterpriseManagerBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEnterpriseManagerBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEnterpriseManagerBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEnterpriseManagerBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEnterpriseManagerBridgeResponse")
	}
	return
}

// deleteEnterpriseManagerBridge implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteEnterpriseManagerBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/enterpriseManagerBridges/{enterpriseManagerBridgeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEnterpriseManagerBridgeResponse
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

// DeleteExadataInsight Deletes an Exadata insight. The Exadata insight will be deleted and cannot be enabled again.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteExadataInsight.go.html to see an example of how to use DeleteExadataInsight API.
func (client OperationsInsightsClient) DeleteExadataInsight(ctx context.Context, request DeleteExadataInsightRequest) (response DeleteExadataInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteExadataInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteExadataInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteExadataInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteExadataInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteExadataInsightResponse")
	}
	return
}

// deleteExadataInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteExadataInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/exadataInsights/{exadataInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteExadataInsightResponse
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

// DeleteHostInsight Deletes a host insight. The host insight will be deleted and cannot be enabled again.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteHostInsight.go.html to see an example of how to use DeleteHostInsight API.
func (client OperationsInsightsClient) DeleteHostInsight(ctx context.Context, request DeleteHostInsightRequest) (response DeleteHostInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteHostInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteHostInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteHostInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteHostInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteHostInsightResponse")
	}
	return
}

// deleteHostInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteHostInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/hostInsights/{hostInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteHostInsightResponse
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

// DeleteOperationsInsightsWarehouse Deletes an Operations Insights Warehouse. There is only expected to be 1 warehouse per tenant.
// The warehouse is expected to be in the root compartment.
// User must delete AWR Hub resource for this warehouse before calling this operation.
// User must delete the warehouse users before calling this operation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteOperationsInsightsWarehouse.go.html to see an example of how to use DeleteOperationsInsightsWarehouse API.
func (client OperationsInsightsClient) DeleteOperationsInsightsWarehouse(ctx context.Context, request DeleteOperationsInsightsWarehouseRequest) (response DeleteOperationsInsightsWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOperationsInsightsWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOperationsInsightsWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOperationsInsightsWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOperationsInsightsWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOperationsInsightsWarehouseResponse")
	}
	return
}

// deleteOperationsInsightsWarehouse implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteOperationsInsightsWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/operationsInsightsWarehouses/{operationsInsightsWarehouseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOperationsInsightsWarehouseResponse
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

// DeleteOperationsInsightsWarehouseUser Deletes an Operations Insights Warehouse User.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteOperationsInsightsWarehouseUser.go.html to see an example of how to use DeleteOperationsInsightsWarehouseUser API.
func (client OperationsInsightsClient) DeleteOperationsInsightsWarehouseUser(ctx context.Context, request DeleteOperationsInsightsWarehouseUserRequest) (response DeleteOperationsInsightsWarehouseUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOperationsInsightsWarehouseUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOperationsInsightsWarehouseUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOperationsInsightsWarehouseUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOperationsInsightsWarehouseUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOperationsInsightsWarehouseUserResponse")
	}
	return
}

// deleteOperationsInsightsWarehouseUser implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteOperationsInsightsWarehouseUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/operationsInsightsWarehouseUsers/{operationsInsightsWarehouseUserId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOperationsInsightsWarehouseUserResponse
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

// DisableDatabaseInsight Disables a database in Operations Insights. Database metric collection and analysis will be stopped.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableDatabaseInsight.go.html to see an example of how to use DisableDatabaseInsight API.
func (client OperationsInsightsClient) DisableDatabaseInsight(ctx context.Context, request DisableDatabaseInsightRequest) (response DisableDatabaseInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableDatabaseInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableDatabaseInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableDatabaseInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableDatabaseInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableDatabaseInsightResponse")
	}
	return
}

// disableDatabaseInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) disableDatabaseInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/disable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableDatabaseInsightResponse
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

// DisableExadataInsight Disables an Exadata system in Operations Insights. Exadata-related metric collection and analysis will be stopped.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableExadataInsight.go.html to see an example of how to use DisableExadataInsight API.
func (client OperationsInsightsClient) DisableExadataInsight(ctx context.Context, request DisableExadataInsightRequest) (response DisableExadataInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExadataInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExadataInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExadataInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExadataInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExadataInsightResponse")
	}
	return
}

// disableExadataInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) disableExadataInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInsights/{exadataInsightId}/actions/disable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExadataInsightResponse
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

// DisableHostInsight Disables a host in Operations Insights. Host metric collection and analysis will be stopped.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableHostInsight.go.html to see an example of how to use DisableHostInsight API.
func (client OperationsInsightsClient) DisableHostInsight(ctx context.Context, request DisableHostInsightRequest) (response DisableHostInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableHostInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableHostInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableHostInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableHostInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableHostInsightResponse")
	}
	return
}

// disableHostInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) disableHostInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/hostInsights/{hostInsightId}/actions/disable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableHostInsightResponse
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

// DownloadOperationsInsightsWarehouseWallet Download the ADW wallet for Operations Insights Warehouse using which the Hub data is exposed.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DownloadOperationsInsightsWarehouseWallet.go.html to see an example of how to use DownloadOperationsInsightsWarehouseWallet API.
func (client OperationsInsightsClient) DownloadOperationsInsightsWarehouseWallet(ctx context.Context, request DownloadOperationsInsightsWarehouseWalletRequest) (response DownloadOperationsInsightsWarehouseWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadOperationsInsightsWarehouseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadOperationsInsightsWarehouseWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadOperationsInsightsWarehouseWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadOperationsInsightsWarehouseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadOperationsInsightsWarehouseWalletResponse")
	}
	return
}

// downloadOperationsInsightsWarehouseWallet implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) downloadOperationsInsightsWarehouseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operationsInsightsWarehouses/{operationsInsightsWarehouseId}/actions/downloadWarehouseWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadOperationsInsightsWarehouseWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableDatabaseInsight Enables a database in Operations Insights. Database metric collection and analysis will be started.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableDatabaseInsight.go.html to see an example of how to use EnableDatabaseInsight API.
func (client OperationsInsightsClient) EnableDatabaseInsight(ctx context.Context, request EnableDatabaseInsightRequest) (response EnableDatabaseInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableDatabaseInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableDatabaseInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableDatabaseInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableDatabaseInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableDatabaseInsightResponse")
	}
	return
}

// enableDatabaseInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) enableDatabaseInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/enable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableDatabaseInsightResponse
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

// EnableExadataInsight Enables an Exadata system in Operations Insights. Exadata-related metric collection and analysis will be started.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableExadataInsight.go.html to see an example of how to use EnableExadataInsight API.
func (client OperationsInsightsClient) EnableExadataInsight(ctx context.Context, request EnableExadataInsightRequest) (response EnableExadataInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableExadataInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableExadataInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableExadataInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableExadataInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableExadataInsightResponse")
	}
	return
}

// enableExadataInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) enableExadataInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/exadataInsights/{exadataInsightId}/actions/enable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableExadataInsightResponse
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

// EnableHostInsight Enables a host in Operations Insights. Host metric collection and analysis will be started.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableHostInsight.go.html to see an example of how to use EnableHostInsight API.
func (client OperationsInsightsClient) EnableHostInsight(ctx context.Context, request EnableHostInsightRequest) (response EnableHostInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableHostInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableHostInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableHostInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableHostInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableHostInsightResponse")
	}
	return
}

// enableHostInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) enableHostInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/hostInsights/{hostInsightId}/actions/enable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableHostInsightResponse
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

// GetAwrHub Gets details of an AWR hub.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrHub.go.html to see an example of how to use GetAwrHub API.
func (client OperationsInsightsClient) GetAwrHub(ctx context.Context, request GetAwrHubRequest) (response GetAwrHubResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAwrHub, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrHubResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrHubResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrHubResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrHubResponse")
	}
	return
}

// getAwrHub implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getAwrHub(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrHubResponse
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

// GetAwrReport Gets the AWR report for the specified source database in the AWR hub. The difference between the timeGreaterThanOrEqualTo and timeLessThanOrEqualTo should not be greater than 7 days.
// Either beginSnapshotIdentifierGreaterThanOrEqualTo & endSnapshotIdentifierLessThanOrEqualTo params Or timeGreaterThanOrEqualTo & timeLessThanOrEqualTo params are required.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrReport.go.html to see an example of how to use GetAwrReport API.
func (client OperationsInsightsClient) GetAwrReport(ctx context.Context, request GetAwrReportRequest) (response GetAwrReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAwrReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrReportResponse")
	}
	return
}

// getAwrReport implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getAwrReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrReportResponse
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

// GetDatabaseInsight Gets details of a database insight.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetDatabaseInsight.go.html to see an example of how to use GetDatabaseInsight API.
func (client OperationsInsightsClient) GetDatabaseInsight(ctx context.Context, request GetDatabaseInsightRequest) (response GetDatabaseInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseInsightResponse")
	}
	return
}

// getDatabaseInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getDatabaseInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/{databaseInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseInsightResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databaseinsight{})
	return response, err
}

// GetEnterpriseManagerBridge Gets details of an Operations Insights Enterprise Manager bridge.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetEnterpriseManagerBridge.go.html to see an example of how to use GetEnterpriseManagerBridge API.
func (client OperationsInsightsClient) GetEnterpriseManagerBridge(ctx context.Context, request GetEnterpriseManagerBridgeRequest) (response GetEnterpriseManagerBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEnterpriseManagerBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEnterpriseManagerBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEnterpriseManagerBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEnterpriseManagerBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEnterpriseManagerBridgeResponse")
	}
	return
}

// getEnterpriseManagerBridge implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getEnterpriseManagerBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enterpriseManagerBridges/{enterpriseManagerBridgeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEnterpriseManagerBridgeResponse
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

// GetExadataInsight Gets details of an Exadata insight.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetExadataInsight.go.html to see an example of how to use GetExadataInsight API.
func (client OperationsInsightsClient) GetExadataInsight(ctx context.Context, request GetExadataInsightRequest) (response GetExadataInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getExadataInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetExadataInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetExadataInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetExadataInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetExadataInsightResponse")
	}
	return
}

// getExadataInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getExadataInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/{exadataInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetExadataInsightResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &exadatainsight{})
	return response, err
}

// GetHostInsight Gets details of a host insight.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetHostInsight.go.html to see an example of how to use GetHostInsight API.
func (client OperationsInsightsClient) GetHostInsight(ctx context.Context, request GetHostInsightRequest) (response GetHostInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHostInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHostInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHostInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHostInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHostInsightResponse")
	}
	return
}

// getHostInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getHostInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/{hostInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetHostInsightResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &hostinsight{})
	return response, err
}

// GetOperationsInsightsWarehouse Gets details of an Operations Insights Warehouse.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOperationsInsightsWarehouse.go.html to see an example of how to use GetOperationsInsightsWarehouse API.
func (client OperationsInsightsClient) GetOperationsInsightsWarehouse(ctx context.Context, request GetOperationsInsightsWarehouseRequest) (response GetOperationsInsightsWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOperationsInsightsWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOperationsInsightsWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOperationsInsightsWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOperationsInsightsWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOperationsInsightsWarehouseResponse")
	}
	return
}

// getOperationsInsightsWarehouse implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getOperationsInsightsWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operationsInsightsWarehouses/{operationsInsightsWarehouseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOperationsInsightsWarehouseResponse
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

// GetOperationsInsightsWarehouseUser Gets details of an Operations Insights Warehouse User.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOperationsInsightsWarehouseUser.go.html to see an example of how to use GetOperationsInsightsWarehouseUser API.
func (client OperationsInsightsClient) GetOperationsInsightsWarehouseUser(ctx context.Context, request GetOperationsInsightsWarehouseUserRequest) (response GetOperationsInsightsWarehouseUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOperationsInsightsWarehouseUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOperationsInsightsWarehouseUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOperationsInsightsWarehouseUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOperationsInsightsWarehouseUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOperationsInsightsWarehouseUserResponse")
	}
	return
}

// getOperationsInsightsWarehouseUser implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getOperationsInsightsWarehouseUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operationsInsightsWarehouseUsers/{operationsInsightsWarehouseUserId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOperationsInsightsWarehouseUserResponse
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

// GetWorkRequest Gets the status of the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client OperationsInsightsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
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

// IngestDatabaseConfiguration This is a generic ingest endpoint for all database configuration metrics.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestDatabaseConfiguration.go.html to see an example of how to use IngestDatabaseConfiguration API.
func (client OperationsInsightsClient) IngestDatabaseConfiguration(ctx context.Context, request IngestDatabaseConfigurationRequest) (response IngestDatabaseConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestDatabaseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestDatabaseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestDatabaseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestDatabaseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestDatabaseConfigurationResponse")
	}
	return
}

// ingestDatabaseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestDatabaseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestDatabaseConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestDatabaseConfigurationResponse
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

// IngestHostConfiguration This is a generic ingest endpoint for all the host configuration metrics
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestHostConfiguration.go.html to see an example of how to use IngestHostConfiguration API.
func (client OperationsInsightsClient) IngestHostConfiguration(ctx context.Context, request IngestHostConfigurationRequest) (response IngestHostConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestHostConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestHostConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestHostConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestHostConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestHostConfigurationResponse")
	}
	return
}

// ingestHostConfiguration implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestHostConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/hostInsights/actions/ingestHostConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestHostConfigurationResponse
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

// IngestHostMetrics This is a generic ingest endpoint for all the host performance metrics
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestHostMetrics.go.html to see an example of how to use IngestHostMetrics API.
func (client OperationsInsightsClient) IngestHostMetrics(ctx context.Context, request IngestHostMetricsRequest) (response IngestHostMetricsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestHostMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestHostMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestHostMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestHostMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestHostMetricsResponse")
	}
	return
}

// ingestHostMetrics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestHostMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/hostInsights/actions/ingestHostMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestHostMetricsResponse
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

// IngestSqlBucket The sqlbucket endpoint takes in a JSON payload, persists it in Operations Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlBucket.go.html to see an example of how to use IngestSqlBucket API.
func (client OperationsInsightsClient) IngestSqlBucket(ctx context.Context, request IngestSqlBucketRequest) (response IngestSqlBucketResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestSqlBucket, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestSqlBucketResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestSqlBucketResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestSqlBucketResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestSqlBucketResponse")
	}
	return
}

// ingestSqlBucket implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestSqlBucket(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestSqlBucket", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestSqlBucketResponse
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

// IngestSqlPlanLines The SqlPlanLines endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlPlanLines.go.html to see an example of how to use IngestSqlPlanLines API.
func (client OperationsInsightsClient) IngestSqlPlanLines(ctx context.Context, request IngestSqlPlanLinesRequest) (response IngestSqlPlanLinesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestSqlPlanLines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestSqlPlanLinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestSqlPlanLinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestSqlPlanLinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestSqlPlanLinesResponse")
	}
	return
}

// ingestSqlPlanLines implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestSqlPlanLines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestSqlPlanLines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestSqlPlanLinesResponse
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

// IngestSqlStats The SQL Stats endpoint takes in a JSON payload, persists it in Operations Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlStats.go.html to see an example of how to use IngestSqlStats API.
func (client OperationsInsightsClient) IngestSqlStats(ctx context.Context, request IngestSqlStatsRequest) (response IngestSqlStatsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestSqlStats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestSqlStatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestSqlStatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestSqlStatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestSqlStatsResponse")
	}
	return
}

// ingestSqlStats implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestSqlStats(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestSqlStatsMetric", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestSqlStatsResponse
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

// IngestSqlText The SqlText endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
// Either databaseId or id must be specified.
// Disclaimer: SQL text being uploaded explicitly via APIs is not masked. Any sensitive literals contained in the sqlFullText column should be masked prior to ingestion.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlText.go.html to see an example of how to use IngestSqlText API.
func (client OperationsInsightsClient) IngestSqlText(ctx context.Context, request IngestSqlTextRequest) (response IngestSqlTextResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestSqlText, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestSqlTextResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestSqlTextResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestSqlTextResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestSqlTextResponse")
	}
	return
}

// ingestSqlText implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestSqlText(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestSqlText", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestSqlTextResponse
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

// ListAwrHubs Gets a list of AWR hubs. Either compartmentId or id must be specified. All these resources are expected to be in root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrHubs.go.html to see an example of how to use ListAwrHubs API.
func (client OperationsInsightsClient) ListAwrHubs(ctx context.Context, request ListAwrHubsRequest) (response ListAwrHubsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAwrHubs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrHubsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrHubsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrHubsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrHubsResponse")
	}
	return
}

// listAwrHubs implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAwrHubs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrHubsResponse
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

// ListAwrSnapshots Lists AWR snapshots for the specified source database in the AWR hub. The difference between the timeGreaterThanOrEqualTo and timeLessThanOrEqualTo should not exceed an elapsed range of 1 day.
// The timeGreaterThanOrEqualTo & timeLessThanOrEqualTo params are optional. If these params are not provided, by default last 1 day snapshots will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrSnapshots.go.html to see an example of how to use ListAwrSnapshots API.
func (client OperationsInsightsClient) ListAwrSnapshots(ctx context.Context, request ListAwrSnapshotsRequest) (response ListAwrSnapshotsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAwrSnapshots, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrSnapshotsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrSnapshotsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrSnapshotsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrSnapshotsResponse")
	}
	return
}

// listAwrSnapshots implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAwrSnapshots(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrSnapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrSnapshotsResponse
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

// ListDatabaseConfigurations Gets a list of database insight configurations based on the query parameters specified. Either compartmentId or databaseInsightId query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of database insight configurations in that compartment and in all sub-compartments will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseConfigurations.go.html to see an example of how to use ListDatabaseConfigurations API.
func (client OperationsInsightsClient) ListDatabaseConfigurations(ctx context.Context, request ListDatabaseConfigurationsRequest) (response ListDatabaseConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseConfigurationsResponse")
	}
	return
}

// listDatabaseConfigurations implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listDatabaseConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/databaseConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseConfigurationsResponse
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

// ListDatabaseInsights Gets a list of database insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of database insights in that compartment and in all sub-compartments will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseInsights.go.html to see an example of how to use ListDatabaseInsights API.
func (client OperationsInsightsClient) ListDatabaseInsights(ctx context.Context, request ListDatabaseInsightsRequest) (response ListDatabaseInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseInsightsResponse")
	}
	return
}

// listDatabaseInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listDatabaseInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseInsightsResponse
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

// ListEnterpriseManagerBridges Gets a list of Operations Insights Enterprise Manager bridges. Either compartmentId or id must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of bridges in that compartment and in all sub-compartments will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListEnterpriseManagerBridges.go.html to see an example of how to use ListEnterpriseManagerBridges API.
func (client OperationsInsightsClient) ListEnterpriseManagerBridges(ctx context.Context, request ListEnterpriseManagerBridgesRequest) (response ListEnterpriseManagerBridgesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEnterpriseManagerBridges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEnterpriseManagerBridgesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEnterpriseManagerBridgesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEnterpriseManagerBridgesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEnterpriseManagerBridgesResponse")
	}
	return
}

// listEnterpriseManagerBridges implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listEnterpriseManagerBridges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enterpriseManagerBridges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEnterpriseManagerBridgesResponse
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

// ListExadataConfigurations Gets a list of exadata insight configurations. Either compartmentId or exadataInsightsId query parameter must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListExadataConfigurations.go.html to see an example of how to use ListExadataConfigurations API.
func (client OperationsInsightsClient) ListExadataConfigurations(ctx context.Context, request ListExadataConfigurationsRequest) (response ListExadataConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExadataConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExadataConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExadataConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExadataConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExadataConfigurationsResponse")
	}
	return
}

// listExadataConfigurations implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listExadataConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/exadataConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExadataConfigurationsResponse
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

// ListExadataInsights Gets a list of Exadata insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of Exadata insights in that compartment and in all sub-compartments will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListExadataInsights.go.html to see an example of how to use ListExadataInsights API.
func (client OperationsInsightsClient) ListExadataInsights(ctx context.Context, request ListExadataInsightsRequest) (response ListExadataInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExadataInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExadataInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExadataInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExadataInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExadataInsightsResponse")
	}
	return
}

// listExadataInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listExadataInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExadataInsightsResponse
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

// ListHostConfigurations Gets a list of host insight configurations based on the query parameters specified. Either compartmentId or hostInsightId query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of host insight configurations in that compartment and in all sub-compartments will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostConfigurations.go.html to see an example of how to use ListHostConfigurations API.
func (client OperationsInsightsClient) ListHostConfigurations(ctx context.Context, request ListHostConfigurationsRequest) (response ListHostConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHostConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHostConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHostConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHostConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHostConfigurationsResponse")
	}
	return
}

// listHostConfigurations implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listHostConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/hostConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHostConfigurationsResponse
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

// ListHostInsights Gets a list of host insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of host insights in that compartment and in all sub-compartments will be returned.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostInsights.go.html to see an example of how to use ListHostInsights API.
func (client OperationsInsightsClient) ListHostInsights(ctx context.Context, request ListHostInsightsRequest) (response ListHostInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHostInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHostInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHostInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHostInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHostInsightsResponse")
	}
	return
}

// listHostInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listHostInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHostInsightsResponse
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

// ListHostedEntities Get a list of hosted entities details.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostedEntities.go.html to see an example of how to use ListHostedEntities API.
func (client OperationsInsightsClient) ListHostedEntities(ctx context.Context, request ListHostedEntitiesRequest) (response ListHostedEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHostedEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHostedEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHostedEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHostedEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHostedEntitiesResponse")
	}
	return
}

// listHostedEntities implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listHostedEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/hostedEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHostedEntitiesResponse
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

// ListImportableAgentEntities Gets a list of agent entities available to add a new hostInsight.  An agent entity is "available"
// and will be shown if all the following conditions are true:
//    1.  The agent OCID is not already being used for an existing hostInsight.
//    2.  The agent availabilityStatus = 'ACTIVE'
//    3.  The agent lifecycleState = 'ACTIVE'
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableAgentEntities.go.html to see an example of how to use ListImportableAgentEntities API.
func (client OperationsInsightsClient) ListImportableAgentEntities(ctx context.Context, request ListImportableAgentEntitiesRequest) (response ListImportableAgentEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listImportableAgentEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListImportableAgentEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListImportableAgentEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListImportableAgentEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListImportableAgentEntitiesResponse")
	}
	return
}

// listImportableAgentEntities implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listImportableAgentEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/importableAgentEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListImportableAgentEntitiesResponse
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

// ListImportableEnterpriseManagerEntities Gets a list of importable entities for an Operations Insights Enterprise Manager bridge that have not been imported before.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableEnterpriseManagerEntities.go.html to see an example of how to use ListImportableEnterpriseManagerEntities API.
func (client OperationsInsightsClient) ListImportableEnterpriseManagerEntities(ctx context.Context, request ListImportableEnterpriseManagerEntitiesRequest) (response ListImportableEnterpriseManagerEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listImportableEnterpriseManagerEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListImportableEnterpriseManagerEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListImportableEnterpriseManagerEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListImportableEnterpriseManagerEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListImportableEnterpriseManagerEntitiesResponse")
	}
	return
}

// listImportableEnterpriseManagerEntities implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listImportableEnterpriseManagerEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enterpriseManagerBridges/{enterpriseManagerBridgeId}/importableEnterpriseManagerEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListImportableEnterpriseManagerEntitiesResponse
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

// ListOperationsInsightsWarehouseUsers Gets a list of Operations Insights Warehouse users. Either compartmentId or id must be specified. All these resources are expected to be in root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsWarehouseUsers.go.html to see an example of how to use ListOperationsInsightsWarehouseUsers API.
func (client OperationsInsightsClient) ListOperationsInsightsWarehouseUsers(ctx context.Context, request ListOperationsInsightsWarehouseUsersRequest) (response ListOperationsInsightsWarehouseUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOperationsInsightsWarehouseUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOperationsInsightsWarehouseUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOperationsInsightsWarehouseUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOperationsInsightsWarehouseUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOperationsInsightsWarehouseUsersResponse")
	}
	return
}

// listOperationsInsightsWarehouseUsers implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listOperationsInsightsWarehouseUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operationsInsightsWarehouseUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOperationsInsightsWarehouseUsersResponse
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

// ListOperationsInsightsWarehouses Gets a list of Operations Insights warehouses. Either compartmentId or id must be specified.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsWarehouses.go.html to see an example of how to use ListOperationsInsightsWarehouses API.
func (client OperationsInsightsClient) ListOperationsInsightsWarehouses(ctx context.Context, request ListOperationsInsightsWarehousesRequest) (response ListOperationsInsightsWarehousesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOperationsInsightsWarehouses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOperationsInsightsWarehousesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOperationsInsightsWarehousesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOperationsInsightsWarehousesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOperationsInsightsWarehousesResponse")
	}
	return
}

// listOperationsInsightsWarehouses implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listOperationsInsightsWarehouses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operationsInsightsWarehouses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOperationsInsightsWarehousesResponse
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

// ListSqlPlans Query SQL Warehouse to list the plan xml for a given SQL execution plan. This returns a SqlPlanCollection object, but is currently limited to a single plan.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlPlans.go.html to see an example of how to use ListSqlPlans API.
func (client OperationsInsightsClient) ListSqlPlans(ctx context.Context, request ListSqlPlansRequest) (response ListSqlPlansResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlPlans, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlPlansResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlPlansResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlPlansResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlPlansResponse")
	}
	return
}

// listSqlPlans implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listSqlPlans(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlPlans", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlPlansResponse
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

// ListSqlSearches Search SQL by SQL Identifier across databases in a compartment and in all sub-compartments if specified.
// And get the SQL Text and the details of the databases executing the SQL for a given time period.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlSearches.go.html to see an example of how to use ListSqlSearches API.
func (client OperationsInsightsClient) ListSqlSearches(ctx context.Context, request ListSqlSearchesRequest) (response ListSqlSearchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlSearches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlSearchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlSearchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlSearchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlSearchesResponse")
	}
	return
}

// listSqlSearches implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listSqlSearches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlSearches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlSearchesResponse
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

// ListSqlTexts Query SQL Warehouse to get the full SQL Text for a SQL in a compartment and in all sub-compartments if specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlTexts.go.html to see an example of how to use ListSqlTexts API.
func (client OperationsInsightsClient) ListSqlTexts(ctx context.Context, request ListSqlTextsRequest) (response ListSqlTextsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSqlTexts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSqlTextsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSqlTextsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSqlTextsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSqlTextsResponse")
	}
	return
}

// listSqlTexts implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listSqlTexts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlTexts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSqlTextsResponse
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client OperationsInsightsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
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

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client OperationsInsightsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
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

// ListWorkRequests Lists the work requests in a compartment. Either compartmentId or id must be specified. Only one of id, resourceId or relatedResourceId can be specified optionally.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client OperationsInsightsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
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

// RotateOperationsInsightsWarehouseWallet Rotate the ADW wallet for Operations Insights Warehouse using which the Hub data is exposed.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/RotateOperationsInsightsWarehouseWallet.go.html to see an example of how to use RotateOperationsInsightsWarehouseWallet API.
func (client OperationsInsightsClient) RotateOperationsInsightsWarehouseWallet(ctx context.Context, request RotateOperationsInsightsWarehouseWalletRequest) (response RotateOperationsInsightsWarehouseWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.rotateOperationsInsightsWarehouseWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateOperationsInsightsWarehouseWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateOperationsInsightsWarehouseWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateOperationsInsightsWarehouseWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateOperationsInsightsWarehouseWalletResponse")
	}
	return
}

// rotateOperationsInsightsWarehouseWallet implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) rotateOperationsInsightsWarehouseWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operationsInsightsWarehouses/{operationsInsightsWarehouseId}/actions/rotateWarehouseWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateOperationsInsightsWarehouseWalletResponse
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

// SummarizeAwrSourcesSummaries Gets a list of summary of AWR Sources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrSourcesSummaries.go.html to see an example of how to use SummarizeAwrSourcesSummaries API.
func (client OperationsInsightsClient) SummarizeAwrSourcesSummaries(ctx context.Context, request SummarizeAwrSourcesSummariesRequest) (response SummarizeAwrSourcesSummariesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrSourcesSummaries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrSourcesSummariesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrSourcesSummariesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrSourcesSummariesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrSourcesSummariesResponse")
	}
	return
}

// summarizeAwrSourcesSummaries implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrSourcesSummaries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrSourcesSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrSourcesSummariesResponse
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

// SummarizeDatabaseInsightResourceCapacityTrend Returns response with time series data (endTimestamp, capacity, baseCapacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceCapacityTrend API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceCapacityTrend(ctx context.Context, request SummarizeDatabaseInsightResourceCapacityTrendRequest) (response SummarizeDatabaseInsightResourceCapacityTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceCapacityTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceCapacityTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceCapacityTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceCapacityTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceCapacityTrendResponse")
	}
	return
}

// summarizeDatabaseInsightResourceCapacityTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceCapacityTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceCapacityTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceCapacityTrendResponse
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

// SummarizeDatabaseInsightResourceForecastTrend Get Forecast predictions for CPU and Storage resources since a time in the past.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceForecastTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceForecastTrend API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceForecastTrend(ctx context.Context, request SummarizeDatabaseInsightResourceForecastTrendRequest) (response SummarizeDatabaseInsightResourceForecastTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceForecastTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceForecastTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceForecastTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceForecastTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceForecastTrendResponse")
	}
	return
}

// summarizeDatabaseInsightResourceForecastTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceForecastTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceForecastTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceForecastTrendResponse
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

// SummarizeDatabaseInsightResourceStatistics Lists the Resource statistics (usage,capacity, usage change percent, utilization percent, base capacity, isAutoScalingEnabled)
// for each database filtered by utilization level in a compartment and in all sub-compartments if specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceStatistics.go.html to see an example of how to use SummarizeDatabaseInsightResourceStatistics API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceStatistics(ctx context.Context, request SummarizeDatabaseInsightResourceStatisticsRequest) (response SummarizeDatabaseInsightResourceStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceStatisticsResponse")
	}
	return
}

// summarizeDatabaseInsightResourceStatistics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceStatistics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceStatistics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceStatisticsResponse
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

// SummarizeDatabaseInsightResourceUsage A cumulative distribution function is used to rank the usage data points per database within the specified time period.
// For each database, the minimum data point with a ranking > the percentile value is included in the summation.
// Linear regression functions are used to calculate the usage change percentage.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsage.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsage API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUsage(ctx context.Context, request SummarizeDatabaseInsightResourceUsageRequest) (response SummarizeDatabaseInsightResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceUsageResponse")
	}
	return
}

// summarizeDatabaseInsightResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceUsageSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceUsageResponse
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

// SummarizeDatabaseInsightResourceUsageTrend Returns response with time series data (endTimestamp, usage, capacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsageTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsageTrend API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUsageTrend(ctx context.Context, request SummarizeDatabaseInsightResourceUsageTrendRequest) (response SummarizeDatabaseInsightResourceUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceUsageTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceUsageTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceUsageTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceUsageTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceUsageTrendResponse")
	}
	return
}

// summarizeDatabaseInsightResourceUsageTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceUsageTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceUsageTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceUsageTrendResponse
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

// SummarizeDatabaseInsightResourceUtilizationInsight Gets resources with current utilization (high and low) and projected utilization (high and low) for a resource type over specified time period.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeDatabaseInsightResourceUtilizationInsight API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUtilizationInsight(ctx context.Context, request SummarizeDatabaseInsightResourceUtilizationInsightRequest) (response SummarizeDatabaseInsightResourceUtilizationInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightResourceUtilizationInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightResourceUtilizationInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightResourceUtilizationInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightResourceUtilizationInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightResourceUtilizationInsightResponse")
	}
	return
}

// summarizeDatabaseInsightResourceUtilizationInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightResourceUtilizationInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/resourceUtilizationInsight", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightResourceUtilizationInsightResponse
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

// SummarizeDatabaseInsightTablespaceUsageTrend Returns response with usage time series data (endTimestamp, usage, capacity) with breakdown by tablespaceName for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightTablespaceUsageTrend.go.html to see an example of how to use SummarizeDatabaseInsightTablespaceUsageTrend API.
func (client OperationsInsightsClient) SummarizeDatabaseInsightTablespaceUsageTrend(ctx context.Context, request SummarizeDatabaseInsightTablespaceUsageTrendRequest) (response SummarizeDatabaseInsightTablespaceUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeDatabaseInsightTablespaceUsageTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeDatabaseInsightTablespaceUsageTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeDatabaseInsightTablespaceUsageTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeDatabaseInsightTablespaceUsageTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeDatabaseInsightTablespaceUsageTrendResponse")
	}
	return
}

// summarizeDatabaseInsightTablespaceUsageTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeDatabaseInsightTablespaceUsageTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/tablespaceUsageTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeDatabaseInsightTablespaceUsageTrendResponse
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

// SummarizeExadataInsightResourceCapacityTrend Returns response with time series data (endTimestamp, capacity) for the time period specified for an exadata system for a resource metric.
// Additionally resources can be filtered using databaseInsightId, hostInsightId or storageServerName query parameters.
// Top five resources are returned if total exceeds the limit specified.
// Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE. Database name is returned in name field. DatabaseInsightId, cdbName and hostName query parameter applies to ResourceType DATABASE.
// Valid values for ResourceType HOST are CPU and MEMORY. HostName is returned in name field. HostInsightId and hostName query parameter applies to ResourceType HOST.
// Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT. Storage server name is returned in name field for resourceMetric IOPS and THROUGHPUT
// and asmName is returned in name field for resourceMetric STORAGE. StorageServerName query parameter applies to ResourceType STORAGE_SERVER.
// Valid values for ResourceType DISKGROUP is STORAGE. Comma delimited (asmName,diskgroupName) is returned in name field.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeExadataInsightResourceCapacityTrend API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceCapacityTrend(ctx context.Context, request SummarizeExadataInsightResourceCapacityTrendRequest) (response SummarizeExadataInsightResourceCapacityTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceCapacityTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceCapacityTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceCapacityTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceCapacityTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceCapacityTrendResponse")
	}
	return
}

// summarizeExadataInsightResourceCapacityTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceCapacityTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceCapacityTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceCapacityTrendResponse
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

// SummarizeExadataInsightResourceCapacityTrendAggregated Returns response with time series data (endTimestamp, capacity) for the time period specified for an exadata system or fleet aggregation for a resource metric.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE.
// Valid values for ResourceType HOST are CPU and MEMORY.
// Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceCapacityTrendAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceCapacityTrendAggregated API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceCapacityTrendAggregated(ctx context.Context, request SummarizeExadataInsightResourceCapacityTrendAggregatedRequest) (response SummarizeExadataInsightResourceCapacityTrendAggregatedResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceCapacityTrendAggregated, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceCapacityTrendAggregatedResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceCapacityTrendAggregatedResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceCapacityTrendAggregatedResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceCapacityTrendAggregatedResponse")
	}
	return
}

// summarizeExadataInsightResourceCapacityTrendAggregated implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceCapacityTrendAggregated(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceCapacityTrendAggregated", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceCapacityTrendAggregatedResponse
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

// SummarizeExadataInsightResourceForecastTrend Get historical usage and forecast predictions for an exadata system with breakdown by databases, hosts or storage servers.
// Additionally resources can be filtered using databaseInsightId, hostInsightId or storageServerName query parameters.
// Top five resources are returned if total exceeds the limit specified.
// Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE. Database name is returned in name field. DatabaseInsightId , cdbName and hostName query parameter applies to ResourceType DATABASE.
// Valid values for ResourceType HOST are CPU and MEMORY. HostName s returned in name field. HostInsightId and hostName query parameter applies to ResourceType HOST.
// Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT. Storage server name is returned in name field for resourceMetric IOPS and THROUGHPUT
// and asmName is returned in name field for resourceMetric STORAGE. StorageServerName query parameter applies to ResourceType STORAGE_SERVER.
// Valid value for ResourceType DISKGROUP is STORAGE. Comma delimited (asmName,diskgroupName) is returned in name field.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceForecastTrend.go.html to see an example of how to use SummarizeExadataInsightResourceForecastTrend API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceForecastTrend(ctx context.Context, request SummarizeExadataInsightResourceForecastTrendRequest) (response SummarizeExadataInsightResourceForecastTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceForecastTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceForecastTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceForecastTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceForecastTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceForecastTrendResponse")
	}
	return
}

// summarizeExadataInsightResourceForecastTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceForecastTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceForecastTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceForecastTrendResponse
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

// SummarizeExadataInsightResourceForecastTrendAggregated Get aggregated historical usage and forecast predictions for resources. Either compartmentId or exadataInsightsId query parameter must be specified.
// Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE.
// Valid values for ResourceType HOST are CPU and MEMORY.
// Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceForecastTrendAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceForecastTrendAggregated API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceForecastTrendAggregated(ctx context.Context, request SummarizeExadataInsightResourceForecastTrendAggregatedRequest) (response SummarizeExadataInsightResourceForecastTrendAggregatedResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceForecastTrendAggregated, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceForecastTrendAggregatedResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceForecastTrendAggregatedResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceForecastTrendAggregatedResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceForecastTrendAggregatedResponse")
	}
	return
}

// summarizeExadataInsightResourceForecastTrendAggregated implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceForecastTrendAggregated(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceForecastTrendAggregated", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceForecastTrendAggregatedResponse
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

// SummarizeExadataInsightResourceStatistics Lists the Resource statistics (usage, capacity, usage change percent, utilization percent) for each resource based on resourceMetric filtered by utilization level.
// Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE.
// Valid values for ResourceType HOST are CPU and MEMORY.
// Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS, THROUGHPUT.
// Valid value for ResourceType DISKGROUP is STORAGE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceStatistics.go.html to see an example of how to use SummarizeExadataInsightResourceStatistics API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceStatistics(ctx context.Context, request SummarizeExadataInsightResourceStatisticsRequest) (response SummarizeExadataInsightResourceStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceStatisticsResponse")
	}
	return
}

// summarizeExadataInsightResourceStatistics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceStatistics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceStatistics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceStatisticsResponse
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

// SummarizeExadataInsightResourceUsage A cumulative distribution function is used to rank the usage data points per resource within the specified time period.
// For each resource, the minimum data point with a ranking > the percentile value is included in the summation.
// Linear regression functions are used to calculate the usage change percentage.
// Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE.
// Valid values for ResourceType HOST are CPU and MEMORY.
// Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceUsage.go.html to see an example of how to use SummarizeExadataInsightResourceUsage API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceUsage(ctx context.Context, request SummarizeExadataInsightResourceUsageRequest) (response SummarizeExadataInsightResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceUsageResponse")
	}
	return
}

// summarizeExadataInsightResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceUsageSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceUsageResponse
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

// SummarizeExadataInsightResourceUsageAggregated A cumulative distribution function is used to rank the usage data points per database within the specified time period.
// For each database, the minimum data point with a ranking > the percentile value is included in the summation.
// Linear regression functions are used to calculate the usage change percentage.
// Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE.
// Valid values for ResourceType HOST are CPU and MEMORY.
// Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceUsageAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceUsageAggregated API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceUsageAggregated(ctx context.Context, request SummarizeExadataInsightResourceUsageAggregatedRequest) (response SummarizeExadataInsightResourceUsageAggregatedResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceUsageAggregated, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceUsageAggregatedResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceUsageAggregatedResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceUsageAggregatedResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceUsageAggregatedResponse")
	}
	return
}

// summarizeExadataInsightResourceUsageAggregated implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceUsageAggregated(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceUsageSummaryAggregated", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceUsageAggregatedResponse
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

// SummarizeExadataInsightResourceUtilizationInsight Gets current utilization, projected utilization and days to reach projectedUtilization for an exadata system over specified time period. Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE. Valid values for ResourceType HOST are CPU and MEMORY. Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeExadataInsightResourceUtilizationInsight API.
func (client OperationsInsightsClient) SummarizeExadataInsightResourceUtilizationInsight(ctx context.Context, request SummarizeExadataInsightResourceUtilizationInsightRequest) (response SummarizeExadataInsightResourceUtilizationInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataInsightResourceUtilizationInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataInsightResourceUtilizationInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataInsightResourceUtilizationInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataInsightResourceUtilizationInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataInsightResourceUtilizationInsightResponse")
	}
	return
}

// summarizeExadataInsightResourceUtilizationInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataInsightResourceUtilizationInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/resourceUtilizationInsight", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataInsightResourceUtilizationInsightResponse
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

// SummarizeExadataMembers Lists the software and hardware inventory of the Exadata System.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataMembers.go.html to see an example of how to use SummarizeExadataMembers API.
func (client OperationsInsightsClient) SummarizeExadataMembers(ctx context.Context, request SummarizeExadataMembersRequest) (response SummarizeExadataMembersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeExadataMembers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeExadataMembersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeExadataMembersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeExadataMembersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeExadataMembersResponse")
	}
	return
}

// summarizeExadataMembers implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeExadataMembers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/exadataInsights/exadataMembers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeExadataMembersResponse
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

// SummarizeHostInsightResourceCapacityTrend Returns response with time series data (endTimestamp, capacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeHostInsightResourceCapacityTrend API.
func (client OperationsInsightsClient) SummarizeHostInsightResourceCapacityTrend(ctx context.Context, request SummarizeHostInsightResourceCapacityTrendRequest) (response SummarizeHostInsightResourceCapacityTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightResourceCapacityTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightResourceCapacityTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightResourceCapacityTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightResourceCapacityTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightResourceCapacityTrendResponse")
	}
	return
}

// summarizeHostInsightResourceCapacityTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightResourceCapacityTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/resourceCapacityTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightResourceCapacityTrendResponse
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

// SummarizeHostInsightResourceForecastTrend Get Forecast predictions for CPU or memory resources since a time in the past.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceForecastTrend.go.html to see an example of how to use SummarizeHostInsightResourceForecastTrend API.
func (client OperationsInsightsClient) SummarizeHostInsightResourceForecastTrend(ctx context.Context, request SummarizeHostInsightResourceForecastTrendRequest) (response SummarizeHostInsightResourceForecastTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightResourceForecastTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightResourceForecastTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightResourceForecastTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightResourceForecastTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightResourceForecastTrendResponse")
	}
	return
}

// summarizeHostInsightResourceForecastTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightResourceForecastTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/resourceForecastTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightResourceForecastTrendResponse
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

// SummarizeHostInsightResourceStatistics Lists the resource statistics (usage, capacity, usage change percent, utilization percent, load) for each host filtered
// by utilization level in a compartment and in all sub-compartments if specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceStatistics.go.html to see an example of how to use SummarizeHostInsightResourceStatistics API.
func (client OperationsInsightsClient) SummarizeHostInsightResourceStatistics(ctx context.Context, request SummarizeHostInsightResourceStatisticsRequest) (response SummarizeHostInsightResourceStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightResourceStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightResourceStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightResourceStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightResourceStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightResourceStatisticsResponse")
	}
	return
}

// summarizeHostInsightResourceStatistics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightResourceStatistics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/resourceStatistics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightResourceStatisticsResponse
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

// SummarizeHostInsightResourceUsage A cumulative distribution function is used to rank the usage data points per host within the specified time period.
// For each host, the minimum data point with a ranking > the percentile value is included in the summation.
// Linear regression functions are used to calculate the usage change percentage.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUsage.go.html to see an example of how to use SummarizeHostInsightResourceUsage API.
func (client OperationsInsightsClient) SummarizeHostInsightResourceUsage(ctx context.Context, request SummarizeHostInsightResourceUsageRequest) (response SummarizeHostInsightResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightResourceUsageResponse")
	}
	return
}

// summarizeHostInsightResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/resourceUsageSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightResourceUsageResponse
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

// SummarizeHostInsightResourceUsageTrend Returns response with time series data (endTimestamp, usage, capacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUsageTrend.go.html to see an example of how to use SummarizeHostInsightResourceUsageTrend API.
func (client OperationsInsightsClient) SummarizeHostInsightResourceUsageTrend(ctx context.Context, request SummarizeHostInsightResourceUsageTrendRequest) (response SummarizeHostInsightResourceUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightResourceUsageTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightResourceUsageTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightResourceUsageTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightResourceUsageTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightResourceUsageTrendResponse")
	}
	return
}

// summarizeHostInsightResourceUsageTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightResourceUsageTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/resourceUsageTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightResourceUsageTrendResponse
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

// SummarizeHostInsightResourceUtilizationInsight Gets resources with current utilization (high and low) and projected utilization (high and low) for a resource type over specified time period.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeHostInsightResourceUtilizationInsight API.
func (client OperationsInsightsClient) SummarizeHostInsightResourceUtilizationInsight(ctx context.Context, request SummarizeHostInsightResourceUtilizationInsightRequest) (response SummarizeHostInsightResourceUtilizationInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightResourceUtilizationInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightResourceUtilizationInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightResourceUtilizationInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightResourceUtilizationInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightResourceUtilizationInsightResponse")
	}
	return
}

// summarizeHostInsightResourceUtilizationInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightResourceUtilizationInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/resourceUtilizationInsight", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightResourceUtilizationInsightResponse
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

// SummarizeOperationsInsightsWarehouseResourceUsage Gets the details of resources used by an Operations Insights Warehouse.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeOperationsInsightsWarehouseResourceUsage.go.html to see an example of how to use SummarizeOperationsInsightsWarehouseResourceUsage API.
func (client OperationsInsightsClient) SummarizeOperationsInsightsWarehouseResourceUsage(ctx context.Context, request SummarizeOperationsInsightsWarehouseResourceUsageRequest) (response SummarizeOperationsInsightsWarehouseResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeOperationsInsightsWarehouseResourceUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeOperationsInsightsWarehouseResourceUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeOperationsInsightsWarehouseResourceUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeOperationsInsightsWarehouseResourceUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeOperationsInsightsWarehouseResourceUsageResponse")
	}
	return
}

// summarizeOperationsInsightsWarehouseResourceUsage implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeOperationsInsightsWarehouseResourceUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operationsInsightsWarehouses/{operationsInsightsWarehouseId}/resourceUsageSummary", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeOperationsInsightsWarehouseResourceUsageResponse
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

// SummarizeSqlInsights Query SQL Warehouse to get the performance insights for SQLs taking greater than X% database time for a given
// time period across the given databases or database types in a compartment and in all sub-compartments if specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlInsights.go.html to see an example of how to use SummarizeSqlInsights API.
func (client OperationsInsightsClient) SummarizeSqlInsights(ctx context.Context, request SummarizeSqlInsightsRequest) (response SummarizeSqlInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlInsightsResponse")
	}
	return
}

// summarizeSqlInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlInsightsResponse
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

// SummarizeSqlPlanInsights Query SQL Warehouse to get the performance insights on the execution plans for a given SQL for a given time period.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlPlanInsights.go.html to see an example of how to use SummarizeSqlPlanInsights API.
func (client OperationsInsightsClient) SummarizeSqlPlanInsights(ctx context.Context, request SummarizeSqlPlanInsightsRequest) (response SummarizeSqlPlanInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlPlanInsights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlPlanInsightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlPlanInsightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlPlanInsightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlPlanInsightsResponse")
	}
	return
}

// summarizeSqlPlanInsights implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlPlanInsights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlPlanInsights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlPlanInsightsResponse
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

// SummarizeSqlResponseTimeDistributions Query SQL Warehouse to summarize the response time distribution of query executions for a given SQL for a given time period.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlResponseTimeDistributions.go.html to see an example of how to use SummarizeSqlResponseTimeDistributions API.
func (client OperationsInsightsClient) SummarizeSqlResponseTimeDistributions(ctx context.Context, request SummarizeSqlResponseTimeDistributionsRequest) (response SummarizeSqlResponseTimeDistributionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlResponseTimeDistributions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlResponseTimeDistributionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlResponseTimeDistributionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlResponseTimeDistributionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlResponseTimeDistributionsResponse")
	}
	return
}

// summarizeSqlResponseTimeDistributions implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlResponseTimeDistributions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlResponseTimeDistributions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlResponseTimeDistributionsResponse
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

// SummarizeSqlStatistics Query SQL Warehouse to get the performance statistics for SQLs taking greater than X% database time for a given
// time period across the given databases or database types in a compartment and in all sub-compartments if specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatistics.go.html to see an example of how to use SummarizeSqlStatistics API.
func (client OperationsInsightsClient) SummarizeSqlStatistics(ctx context.Context, request SummarizeSqlStatisticsRequest) (response SummarizeSqlStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlStatisticsResponse")
	}
	return
}

// summarizeSqlStatistics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlStatistics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlStatistics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlStatisticsResponse
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

// SummarizeSqlStatisticsTimeSeries Query SQL Warehouse to get the performance statistics time series for a given SQL across given databases for a
// given time period in a compartment and in all sub-compartments if specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatisticsTimeSeries.go.html to see an example of how to use SummarizeSqlStatisticsTimeSeries API.
func (client OperationsInsightsClient) SummarizeSqlStatisticsTimeSeries(ctx context.Context, request SummarizeSqlStatisticsTimeSeriesRequest) (response SummarizeSqlStatisticsTimeSeriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlStatisticsTimeSeries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlStatisticsTimeSeriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlStatisticsTimeSeriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlStatisticsTimeSeriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlStatisticsTimeSeriesResponse")
	}
	return
}

// summarizeSqlStatisticsTimeSeries implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlStatisticsTimeSeries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlStatisticsTimeSeries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlStatisticsTimeSeriesResponse
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

// SummarizeSqlStatisticsTimeSeriesByPlan Query SQL Warehouse to get the performance statistics time series for a given SQL by execution plans for a given time period.
// Either databaseId or id must be specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatisticsTimeSeriesByPlan.go.html to see an example of how to use SummarizeSqlStatisticsTimeSeriesByPlan API.
func (client OperationsInsightsClient) SummarizeSqlStatisticsTimeSeriesByPlan(ctx context.Context, request SummarizeSqlStatisticsTimeSeriesByPlanRequest) (response SummarizeSqlStatisticsTimeSeriesByPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeSqlStatisticsTimeSeriesByPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeSqlStatisticsTimeSeriesByPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeSqlStatisticsTimeSeriesByPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeSqlStatisticsTimeSeriesByPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeSqlStatisticsTimeSeriesByPlanResponse")
	}
	return
}

// summarizeSqlStatisticsTimeSeriesByPlan implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeSqlStatisticsTimeSeriesByPlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/sqlStatisticsTimeSeriesByPlan", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeSqlStatisticsTimeSeriesByPlanResponse
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

// UpdateAwrHub Updates the configuration of a hub .
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateAwrHub.go.html to see an example of how to use UpdateAwrHub API.
func (client OperationsInsightsClient) UpdateAwrHub(ctx context.Context, request UpdateAwrHubRequest) (response UpdateAwrHubResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAwrHub, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAwrHubResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAwrHubResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAwrHubResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAwrHubResponse")
	}
	return
}

// updateAwrHub implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateAwrHub(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/awrHubs/{awrHubId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAwrHubResponse
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

// UpdateDatabaseInsight Updates configuration of a database insight.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateDatabaseInsight.go.html to see an example of how to use UpdateDatabaseInsight API.
func (client OperationsInsightsClient) UpdateDatabaseInsight(ctx context.Context, request UpdateDatabaseInsightRequest) (response UpdateDatabaseInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseInsightResponse")
	}
	return
}

// updateDatabaseInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateDatabaseInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseInsights/{databaseInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseInsightResponse
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

// UpdateEnterpriseManagerBridge Updates configuration of an Operations Insights Enterprise Manager bridge.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateEnterpriseManagerBridge.go.html to see an example of how to use UpdateEnterpriseManagerBridge API.
func (client OperationsInsightsClient) UpdateEnterpriseManagerBridge(ctx context.Context, request UpdateEnterpriseManagerBridgeRequest) (response UpdateEnterpriseManagerBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEnterpriseManagerBridge, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEnterpriseManagerBridgeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEnterpriseManagerBridgeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEnterpriseManagerBridgeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEnterpriseManagerBridgeResponse")
	}
	return
}

// updateEnterpriseManagerBridge implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateEnterpriseManagerBridge(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/enterpriseManagerBridges/{enterpriseManagerBridgeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEnterpriseManagerBridgeResponse
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

// UpdateExadataInsight Updates configuration of an Exadata insight.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateExadataInsight.go.html to see an example of how to use UpdateExadataInsight API.
func (client OperationsInsightsClient) UpdateExadataInsight(ctx context.Context, request UpdateExadataInsightRequest) (response UpdateExadataInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateExadataInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateExadataInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateExadataInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateExadataInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateExadataInsightResponse")
	}
	return
}

// updateExadataInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateExadataInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/exadataInsights/{exadataInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateExadataInsightResponse
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

// UpdateHostInsight Updates configuration of a host insight.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateHostInsight.go.html to see an example of how to use UpdateHostInsight API.
func (client OperationsInsightsClient) UpdateHostInsight(ctx context.Context, request UpdateHostInsightRequest) (response UpdateHostInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateHostInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateHostInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateHostInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateHostInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateHostInsightResponse")
	}
	return
}

// updateHostInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateHostInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/hostInsights/{hostInsightId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateHostInsightResponse
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

// UpdateOperationsInsightsWarehouse Updates the configuration of an Operations Insights Warehouse.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateOperationsInsightsWarehouse.go.html to see an example of how to use UpdateOperationsInsightsWarehouse API.
func (client OperationsInsightsClient) UpdateOperationsInsightsWarehouse(ctx context.Context, request UpdateOperationsInsightsWarehouseRequest) (response UpdateOperationsInsightsWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOperationsInsightsWarehouse, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOperationsInsightsWarehouseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOperationsInsightsWarehouseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOperationsInsightsWarehouseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOperationsInsightsWarehouseResponse")
	}
	return
}

// updateOperationsInsightsWarehouse implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateOperationsInsightsWarehouse(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/operationsInsightsWarehouses/{operationsInsightsWarehouseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOperationsInsightsWarehouseResponse
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

// UpdateOperationsInsightsWarehouseUser Updates the configuration of an Operations Insights Warehouse User.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateOperationsInsightsWarehouseUser.go.html to see an example of how to use UpdateOperationsInsightsWarehouseUser API.
func (client OperationsInsightsClient) UpdateOperationsInsightsWarehouseUser(ctx context.Context, request UpdateOperationsInsightsWarehouseUserRequest) (response UpdateOperationsInsightsWarehouseUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOperationsInsightsWarehouseUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOperationsInsightsWarehouseUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOperationsInsightsWarehouseUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOperationsInsightsWarehouseUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOperationsInsightsWarehouseUserResponse")
	}
	return
}

// updateOperationsInsightsWarehouseUser implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateOperationsInsightsWarehouseUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/operationsInsightsWarehouseUsers/{operationsInsightsWarehouseUserId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOperationsInsightsWarehouseUserResponse
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
