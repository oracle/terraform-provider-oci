// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OperationsInsightsClient a client for OperationsInsights
type OperationsInsightsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOperationsInsightsClientWithConfigurationProvider Creates a new default OperationsInsights client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOperationsInsightsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OperationsInsightsClient, err error) {
	if enabled := common.CheckForEnabledServices("opsi"); !enabled {
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
	return newOperationsInsightsClientFromBaseClient(baseClient, provider)
}

// NewOperationsInsightsClientWithOboToken Creates a new default OperationsInsights client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOperationsInsightsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OperationsInsightsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOperationsInsightsClientFromBaseClient(baseClient, configProvider)
}

func newOperationsInsightsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OperationsInsightsClient, err error) {
	// OperationsInsights service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OperationsInsights"))
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
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *OperationsInsightsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddExadataInsightMembers Add new members (e.g. databases and hosts) to an Exadata system in Operations Insights. Exadata-related metric collection and analysis will be started.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/AddExadataInsightMembers.go.html to see an example of how to use AddExadataInsightMembers API.
// A default retry strategy applies to this operation AddExadataInsightMembers()
func (client OperationsInsightsClient) AddExadataInsightMembers(ctx context.Context, request AddExadataInsightMembersRequest) (response AddExadataInsightMembersResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/AddExadataInsightMembers"
		err = common.PostProcessServiceError(err, "OperationsInsights", "AddExadataInsightMembers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAutonomousDatabaseInsightAdvancedFeatures Update connection detail for advanced features of Autonomous Database in Operations Insights.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeAutonomousDatabaseInsightAdvancedFeatures.go.html to see an example of how to use ChangeAutonomousDatabaseInsightAdvancedFeatures API.
// A default retry strategy applies to this operation ChangeAutonomousDatabaseInsightAdvancedFeatures()
func (client OperationsInsightsClient) ChangeAutonomousDatabaseInsightAdvancedFeatures(ctx context.Context, request ChangeAutonomousDatabaseInsightAdvancedFeaturesRequest) (response ChangeAutonomousDatabaseInsightAdvancedFeaturesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAutonomousDatabaseInsightAdvancedFeatures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAutonomousDatabaseInsightAdvancedFeaturesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAutonomousDatabaseInsightAdvancedFeaturesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAutonomousDatabaseInsightAdvancedFeaturesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAutonomousDatabaseInsightAdvancedFeaturesResponse")
	}
	return
}

// changeAutonomousDatabaseInsightAdvancedFeatures implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeAutonomousDatabaseInsightAdvancedFeatures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/changeAutonomousDatabaseInsightAdvancedFeatures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAutonomousDatabaseInsightAdvancedFeaturesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ChangeAutonomousDatabaseInsightAdvancedFeatures"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeAutonomousDatabaseInsightAdvancedFeatures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAwrHubSourceCompartment Moves an AwrHubSource resource from one compartment to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeAwrHubSourceCompartment.go.html to see an example of how to use ChangeAwrHubSourceCompartment API.
// A default retry strategy applies to this operation ChangeAwrHubSourceCompartment()
func (client OperationsInsightsClient) ChangeAwrHubSourceCompartment(ctx context.Context, request ChangeAwrHubSourceCompartmentRequest) (response ChangeAwrHubSourceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAwrHubSourceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAwrHubSourceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAwrHubSourceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAwrHubSourceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAwrHubSourceCompartmentResponse")
	}
	return
}

// changeAwrHubSourceCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeAwrHubSourceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/awrHubSources/{awrHubSourceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAwrHubSourceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/ChangeAwrHubSourceCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeAwrHubSourceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseInsightCompartment Moves a DatabaseInsight resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeDatabaseInsightCompartment.go.html to see an example of how to use ChangeDatabaseInsightCompartment API.
// A default retry strategy applies to this operation ChangeDatabaseInsightCompartment()
func (client OperationsInsightsClient) ChangeDatabaseInsightCompartment(ctx context.Context, request ChangeDatabaseInsightCompartmentRequest) (response ChangeDatabaseInsightCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ChangeDatabaseInsightCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeDatabaseInsightCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEnterpriseManagerBridgeCompartment Moves a EnterpriseManagerBridge resource from one compartment to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeEnterpriseManagerBridgeCompartment.go.html to see an example of how to use ChangeEnterpriseManagerBridgeCompartment API.
// A default retry strategy applies to this operation ChangeEnterpriseManagerBridgeCompartment()
func (client OperationsInsightsClient) ChangeEnterpriseManagerBridgeCompartment(ctx context.Context, request ChangeEnterpriseManagerBridgeCompartmentRequest) (response ChangeEnterpriseManagerBridgeCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/EnterpriseManagerBridges/ChangeEnterpriseManagerBridgeCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeEnterpriseManagerBridgeCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeExadataInsightCompartment Moves an Exadata insight resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeExadataInsightCompartment.go.html to see an example of how to use ChangeExadataInsightCompartment API.
// A default retry strategy applies to this operation ChangeExadataInsightCompartment()
func (client OperationsInsightsClient) ChangeExadataInsightCompartment(ctx context.Context, request ChangeExadataInsightCompartmentRequest) (response ChangeExadataInsightCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/ChangeExadataInsightCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeExadataInsightCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeHostInsightCompartment Moves a HostInsight resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeHostInsightCompartment.go.html to see an example of how to use ChangeHostInsightCompartment API.
// A default retry strategy applies to this operation ChangeHostInsightCompartment()
func (client OperationsInsightsClient) ChangeHostInsightCompartment(ctx context.Context, request ChangeHostInsightCompartmentRequest) (response ChangeHostInsightCompartmentResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/ChangeHostInsightCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeHostInsightCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMacsManagedCloudDatabaseInsightConnection Change the connection details of a Cloud MACS-managed database insight. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeMacsManagedCloudDatabaseInsightConnection.go.html to see an example of how to use ChangeMacsManagedCloudDatabaseInsightConnection API.
// A default retry strategy applies to this operation ChangeMacsManagedCloudDatabaseInsightConnection()
func (client OperationsInsightsClient) ChangeMacsManagedCloudDatabaseInsightConnection(ctx context.Context, request ChangeMacsManagedCloudDatabaseInsightConnectionRequest) (response ChangeMacsManagedCloudDatabaseInsightConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMacsManagedCloudDatabaseInsightConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMacsManagedCloudDatabaseInsightConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMacsManagedCloudDatabaseInsightConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMacsManagedCloudDatabaseInsightConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMacsManagedCloudDatabaseInsightConnectionResponse")
	}
	return
}

// changeMacsManagedCloudDatabaseInsightConnection implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeMacsManagedCloudDatabaseInsightConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/changeMacsManagedCloudDatabaseInsightConnectionDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMacsManagedCloudDatabaseInsightConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ChangeMacsManagedCloudDatabaseInsightConnection"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeMacsManagedCloudDatabaseInsightConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeNewsReportCompartment Moves a news report resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeNewsReportCompartment.go.html to see an example of how to use ChangeNewsReportCompartment API.
// A default retry strategy applies to this operation ChangeNewsReportCompartment()
func (client OperationsInsightsClient) ChangeNewsReportCompartment(ctx context.Context, request ChangeNewsReportCompartmentRequest) (response ChangeNewsReportCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeNewsReportCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeNewsReportCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeNewsReportCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeNewsReportCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeNewsReportCompartmentResponse")
	}
	return
}

// changeNewsReportCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeNewsReportCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/newsReports/{newsReportId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeNewsReportCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/NewsReports/ChangeNewsReportCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeNewsReportCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOperationsInsightsPrivateEndpointCompartment Moves a private endpoint from one compartment to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeOperationsInsightsPrivateEndpointCompartment.go.html to see an example of how to use ChangeOperationsInsightsPrivateEndpointCompartment API.
// A default retry strategy applies to this operation ChangeOperationsInsightsPrivateEndpointCompartment()
func (client OperationsInsightsClient) ChangeOperationsInsightsPrivateEndpointCompartment(ctx context.Context, request ChangeOperationsInsightsPrivateEndpointCompartmentRequest) (response ChangeOperationsInsightsPrivateEndpointCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOperationsInsightsPrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOperationsInsightsPrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOperationsInsightsPrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOperationsInsightsPrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOperationsInsightsPrivateEndpointCompartmentResponse")
	}
	return
}

// changeOperationsInsightsPrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeOperationsInsightsPrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operationsInsightsPrivateEndpoints/{operationsInsightsPrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOperationsInsightsPrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsPrivateEndpoint/ChangeOperationsInsightsPrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeOperationsInsightsPrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOperationsInsightsWarehouseCompartment Moves a Operations Insights Warehouse resource from one compartment to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeOperationsInsightsWarehouseCompartment.go.html to see an example of how to use ChangeOperationsInsightsWarehouseCompartment API.
// A default retry strategy applies to this operation ChangeOperationsInsightsWarehouseCompartment()
func (client OperationsInsightsClient) ChangeOperationsInsightsWarehouseCompartment(ctx context.Context, request ChangeOperationsInsightsWarehouseCompartmentRequest) (response ChangeOperationsInsightsWarehouseCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOperationsInsightsWarehouseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOperationsInsightsWarehouseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOperationsInsightsWarehouseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOperationsInsightsWarehouseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOperationsInsightsWarehouseCompartmentResponse")
	}
	return
}

// changeOperationsInsightsWarehouseCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeOperationsInsightsWarehouseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operationsInsightsWarehouses/{operationsInsightsWarehouseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOperationsInsightsWarehouseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/ChangeOperationsInsightsWarehouseCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeOperationsInsightsWarehouseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeOpsiConfigurationCompartment Moves an OpsiConfiguration resource from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangeOpsiConfigurationCompartment.go.html to see an example of how to use ChangeOpsiConfigurationCompartment API.
// A default retry strategy applies to this operation ChangeOpsiConfigurationCompartment()
func (client OperationsInsightsClient) ChangeOpsiConfigurationCompartment(ctx context.Context, request ChangeOpsiConfigurationCompartmentRequest) (response ChangeOpsiConfigurationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOpsiConfigurationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOpsiConfigurationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOpsiConfigurationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOpsiConfigurationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOpsiConfigurationCompartmentResponse")
	}
	return
}

// changeOpsiConfigurationCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changeOpsiConfigurationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/opsiConfigurations/{opsiConfigurationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOpsiConfigurationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiConfigurations/ChangeOpsiConfigurationCompartment"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangeOpsiConfigurationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePeComanagedDatabaseInsight Change the connection details of a co-managed  database insight. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ChangePeComanagedDatabaseInsight.go.html to see an example of how to use ChangePeComanagedDatabaseInsight API.
// A default retry strategy applies to this operation ChangePeComanagedDatabaseInsight()
func (client OperationsInsightsClient) ChangePeComanagedDatabaseInsight(ctx context.Context, request ChangePeComanagedDatabaseInsightRequest) (response ChangePeComanagedDatabaseInsightResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePeComanagedDatabaseInsight, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePeComanagedDatabaseInsightResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePeComanagedDatabaseInsightResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePeComanagedDatabaseInsightResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePeComanagedDatabaseInsightResponse")
	}
	return
}

// changePeComanagedDatabaseInsight implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) changePeComanagedDatabaseInsight(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/changePeComanagedDatabaseInsightDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePeComanagedDatabaseInsightResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ChangePeComanagedDatabaseInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ChangePeComanagedDatabaseInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAwrHub Create a AWR hub resource for the tenant in Operations Insights.
// This resource will be created in root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateAwrHub.go.html to see an example of how to use CreateAwrHub API.
// A default retry strategy applies to this operation CreateAwrHub()
func (client OperationsInsightsClient) CreateAwrHub(ctx context.Context, request CreateAwrHubRequest) (response CreateAwrHubResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/CreateAwrHub"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateAwrHub", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAwrHubSource Register Awr Hub source
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateAwrHubSource.go.html to see an example of how to use CreateAwrHubSource API.
// A default retry strategy applies to this operation CreateAwrHubSource()
func (client OperationsInsightsClient) CreateAwrHubSource(ctx context.Context, request CreateAwrHubSourceRequest) (response CreateAwrHubSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAwrHubSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAwrHubSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAwrHubSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAwrHubSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAwrHubSourceResponse")
	}
	return
}

// createAwrHubSource implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createAwrHubSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/awrHubSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAwrHubSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/CreateAwrHubSource"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateAwrHubSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatabaseInsight Create a Database Insight resource for a database in Operations Insights. The database will be enabled in Operations Insights. Database metric collection and analysis will be started.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateDatabaseInsight.go.html to see an example of how to use CreateDatabaseInsight API.
// A default retry strategy applies to this operation CreateDatabaseInsight()
func (client OperationsInsightsClient) CreateDatabaseInsight(ctx context.Context, request CreateDatabaseInsightRequest) (response CreateDatabaseInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/CreateDatabaseInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateDatabaseInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databaseinsight{})
	return response, err
}

// CreateEnterpriseManagerBridge Create a Enterprise Manager bridge in Operations Insights.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateEnterpriseManagerBridge.go.html to see an example of how to use CreateEnterpriseManagerBridge API.
// A default retry strategy applies to this operation CreateEnterpriseManagerBridge()
func (client OperationsInsightsClient) CreateEnterpriseManagerBridge(ctx context.Context, request CreateEnterpriseManagerBridgeRequest) (response CreateEnterpriseManagerBridgeResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/EnterpriseManagerBridges/CreateEnterpriseManagerBridge"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateEnterpriseManagerBridge", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateExadataInsight Create an Exadata insight resource for an Exadata system in Operations Insights. The Exadata system will be enabled in Operations Insights. Exadata-related metric collection and analysis will be started.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateExadataInsight.go.html to see an example of how to use CreateExadataInsight API.
// A default retry strategy applies to this operation CreateExadataInsight()
func (client OperationsInsightsClient) CreateExadataInsight(ctx context.Context, request CreateExadataInsightRequest) (response CreateExadataInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/CreateExadataInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateExadataInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &exadatainsight{})
	return response, err
}

// CreateHostInsight Create a Host Insight resource for a host in Ops Insights. The host will be enabled in Ops Insights. Host metric collection and analysis will be started.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateHostInsight.go.html to see an example of how to use CreateHostInsight API.
// A default retry strategy applies to this operation CreateHostInsight()
func (client OperationsInsightsClient) CreateHostInsight(ctx context.Context, request CreateHostInsightRequest) (response CreateHostInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/CreateHostInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateHostInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &hostinsight{})
	return response, err
}

// CreateNewsReport Create a news report in Ops Insights. The report will be enabled in Ops Insights. Insights will be emailed as per selected frequency.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateNewsReport.go.html to see an example of how to use CreateNewsReport API.
// A default retry strategy applies to this operation CreateNewsReport()
func (client OperationsInsightsClient) CreateNewsReport(ctx context.Context, request CreateNewsReportRequest) (response CreateNewsReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createNewsReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNewsReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNewsReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNewsReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNewsReportResponse")
	}
	return
}

// createNewsReport implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createNewsReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/newsReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNewsReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/NewsReports/CreateNewsReport"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateNewsReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOperationsInsightsPrivateEndpoint Create a private endpoint resource for the tenant in Ops Insights.
// This resource will be created in customer compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateOperationsInsightsPrivateEndpoint.go.html to see an example of how to use CreateOperationsInsightsPrivateEndpoint API.
// A default retry strategy applies to this operation CreateOperationsInsightsPrivateEndpoint()
func (client OperationsInsightsClient) CreateOperationsInsightsPrivateEndpoint(ctx context.Context, request CreateOperationsInsightsPrivateEndpointRequest) (response CreateOperationsInsightsPrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOperationsInsightsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOperationsInsightsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOperationsInsightsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOperationsInsightsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOperationsInsightsPrivateEndpointResponse")
	}
	return
}

// createOperationsInsightsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createOperationsInsightsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operationsInsightsPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOperationsInsightsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsPrivateEndpoint/CreateOperationsInsightsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateOperationsInsightsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOperationsInsightsWarehouse Create a Ops Insights Warehouse resource for the tenant in Ops Insights. New ADW will be provisioned for this tenant.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment. If the 'opsi-warehouse-type'
// header is passed to the API, a warehouse resource without ADW or Schema provisioning is created.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateOperationsInsightsWarehouse.go.html to see an example of how to use CreateOperationsInsightsWarehouse API.
// A default retry strategy applies to this operation CreateOperationsInsightsWarehouse()
func (client OperationsInsightsClient) CreateOperationsInsightsWarehouse(ctx context.Context, request CreateOperationsInsightsWarehouseRequest) (response CreateOperationsInsightsWarehouseResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/CreateOperationsInsightsWarehouse"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateOperationsInsightsWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOperationsInsightsWarehouseUser Create a Operations Insights Warehouse user resource for the tenant in Operations Insights.
// This resource will be created in root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateOperationsInsightsWarehouseUser.go.html to see an example of how to use CreateOperationsInsightsWarehouseUser API.
// A default retry strategy applies to this operation CreateOperationsInsightsWarehouseUser()
func (client OperationsInsightsClient) CreateOperationsInsightsWarehouseUser(ctx context.Context, request CreateOperationsInsightsWarehouseUserRequest) (response CreateOperationsInsightsWarehouseUserResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouseUsers/CreateOperationsInsightsWarehouseUser"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateOperationsInsightsWarehouseUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOpsiConfiguration Create an OPSI configuration resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateOpsiConfiguration.go.html to see an example of how to use CreateOpsiConfiguration API.
// A default retry strategy applies to this operation CreateOpsiConfiguration()
func (client OperationsInsightsClient) CreateOpsiConfiguration(ctx context.Context, request CreateOpsiConfigurationRequest) (response CreateOpsiConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOpsiConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOpsiConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOpsiConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOpsiConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOpsiConfigurationResponse")
	}
	return
}

// createOpsiConfiguration implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) createOpsiConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/opsiConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOpsiConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiConfigurations/CreateOpsiConfiguration"
		err = common.PostProcessServiceError(err, "OperationsInsights", "CreateOpsiConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &opsiconfiguration{})
	return response, err
}

// DeleteAwrHub Deletes an AWR hub.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteAwrHub.go.html to see an example of how to use DeleteAwrHub API.
// A default retry strategy applies to this operation DeleteAwrHub()
func (client OperationsInsightsClient) DeleteAwrHub(ctx context.Context, request DeleteAwrHubRequest) (response DeleteAwrHubResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/DeleteAwrHub"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteAwrHub", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAwrHubObject Deletes an Awr Hub object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteAwrHubObject.go.html to see an example of how to use DeleteAwrHubObject API.
// A default retry strategy applies to this operation DeleteAwrHubObject()
func (client OperationsInsightsClient) DeleteAwrHubObject(ctx context.Context, request DeleteAwrHubObjectRequest) (response DeleteAwrHubObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAwrHubObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAwrHubObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAwrHubObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAwrHubObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAwrHubObjectResponse")
	}
	return
}

// deleteAwrHubObject implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteAwrHubObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/awrHubObjects/awrHubSources/{awrHubSourceId}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAwrHubObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubObjects/DeleteAwrHubObject"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteAwrHubObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAwrHubSource Deletes an Awr Hub source object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteAwrHubSource.go.html to see an example of how to use DeleteAwrHubSource API.
// A default retry strategy applies to this operation DeleteAwrHubSource()
func (client OperationsInsightsClient) DeleteAwrHubSource(ctx context.Context, request DeleteAwrHubSourceRequest) (response DeleteAwrHubSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAwrHubSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAwrHubSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAwrHubSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAwrHubSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAwrHubSourceResponse")
	}
	return
}

// deleteAwrHubSource implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteAwrHubSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/awrHubSources/{awrHubSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAwrHubSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/DeleteAwrHubSource"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteAwrHubSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseInsight Deletes a database insight. The database insight will be deleted and cannot be enabled again.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteDatabaseInsight.go.html to see an example of how to use DeleteDatabaseInsight API.
// A default retry strategy applies to this operation DeleteDatabaseInsight()
func (client OperationsInsightsClient) DeleteDatabaseInsight(ctx context.Context, request DeleteDatabaseInsightRequest) (response DeleteDatabaseInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/DeleteDatabaseInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteDatabaseInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEnterpriseManagerBridge Deletes an Operations Insights Enterprise Manager bridge. If any database insight is still referencing this bridge, the operation will fail.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteEnterpriseManagerBridge.go.html to see an example of how to use DeleteEnterpriseManagerBridge API.
// A default retry strategy applies to this operation DeleteEnterpriseManagerBridge()
func (client OperationsInsightsClient) DeleteEnterpriseManagerBridge(ctx context.Context, request DeleteEnterpriseManagerBridgeRequest) (response DeleteEnterpriseManagerBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/EnterpriseManagerBridges/DeleteEnterpriseManagerBridge"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteEnterpriseManagerBridge", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteExadataInsight Deletes an Exadata insight. The Exadata insight will be deleted and cannot be enabled again.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteExadataInsight.go.html to see an example of how to use DeleteExadataInsight API.
// A default retry strategy applies to this operation DeleteExadataInsight()
func (client OperationsInsightsClient) DeleteExadataInsight(ctx context.Context, request DeleteExadataInsightRequest) (response DeleteExadataInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/DeleteExadataInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteExadataInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteHostInsight Deletes a host insight. The host insight will be deleted and cannot be enabled again.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteHostInsight.go.html to see an example of how to use DeleteHostInsight API.
// A default retry strategy applies to this operation DeleteHostInsight()
func (client OperationsInsightsClient) DeleteHostInsight(ctx context.Context, request DeleteHostInsightRequest) (response DeleteHostInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/DeleteHostInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteHostInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNewsReport Deletes a news report. The news report will be deleted and cannot be enabled again.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteNewsReport.go.html to see an example of how to use DeleteNewsReport API.
// A default retry strategy applies to this operation DeleteNewsReport()
func (client OperationsInsightsClient) DeleteNewsReport(ctx context.Context, request DeleteNewsReportRequest) (response DeleteNewsReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNewsReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNewsReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNewsReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNewsReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNewsReportResponse")
	}
	return
}

// deleteNewsReport implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteNewsReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/newsReports/{newsReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNewsReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/NewsReports/DeleteNewsReport"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteNewsReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOperationsInsightsPrivateEndpoint Deletes a private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteOperationsInsightsPrivateEndpoint.go.html to see an example of how to use DeleteOperationsInsightsPrivateEndpoint API.
// A default retry strategy applies to this operation DeleteOperationsInsightsPrivateEndpoint()
func (client OperationsInsightsClient) DeleteOperationsInsightsPrivateEndpoint(ctx context.Context, request DeleteOperationsInsightsPrivateEndpointRequest) (response DeleteOperationsInsightsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOperationsInsightsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOperationsInsightsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOperationsInsightsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOperationsInsightsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOperationsInsightsPrivateEndpointResponse")
	}
	return
}

// deleteOperationsInsightsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteOperationsInsightsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/operationsInsightsPrivateEndpoints/{operationsInsightsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOperationsInsightsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsPrivateEndpoint/DeleteOperationsInsightsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteOperationsInsightsPrivateEndpoint", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteOperationsInsightsWarehouse.go.html to see an example of how to use DeleteOperationsInsightsWarehouse API.
// A default retry strategy applies to this operation DeleteOperationsInsightsWarehouse()
func (client OperationsInsightsClient) DeleteOperationsInsightsWarehouse(ctx context.Context, request DeleteOperationsInsightsWarehouseRequest) (response DeleteOperationsInsightsWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/DeleteOperationsInsightsWarehouse"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteOperationsInsightsWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOperationsInsightsWarehouseUser Deletes an Operations Insights Warehouse User.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteOperationsInsightsWarehouseUser.go.html to see an example of how to use DeleteOperationsInsightsWarehouseUser API.
// A default retry strategy applies to this operation DeleteOperationsInsightsWarehouseUser()
func (client OperationsInsightsClient) DeleteOperationsInsightsWarehouseUser(ctx context.Context, request DeleteOperationsInsightsWarehouseUserRequest) (response DeleteOperationsInsightsWarehouseUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouseUsers/DeleteOperationsInsightsWarehouseUser"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteOperationsInsightsWarehouseUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOpsiConfiguration Deletes an OPSI configuration resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DeleteOpsiConfiguration.go.html to see an example of how to use DeleteOpsiConfiguration API.
// A default retry strategy applies to this operation DeleteOpsiConfiguration()
func (client OperationsInsightsClient) DeleteOpsiConfiguration(ctx context.Context, request DeleteOpsiConfigurationRequest) (response DeleteOpsiConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOpsiConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOpsiConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOpsiConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOpsiConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOpsiConfigurationResponse")
	}
	return
}

// deleteOpsiConfiguration implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) deleteOpsiConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/opsiConfigurations/{opsiConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOpsiConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiConfigurations/DeleteOpsiConfiguration"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DeleteOpsiConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableAutonomousDatabaseInsightAdvancedFeatures Disable advanced features for an Autonomous Database in Operations Insights. The connection detail and advanced features will be removed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableAutonomousDatabaseInsightAdvancedFeatures.go.html to see an example of how to use DisableAutonomousDatabaseInsightAdvancedFeatures API.
// A default retry strategy applies to this operation DisableAutonomousDatabaseInsightAdvancedFeatures()
func (client OperationsInsightsClient) DisableAutonomousDatabaseInsightAdvancedFeatures(ctx context.Context, request DisableAutonomousDatabaseInsightAdvancedFeaturesRequest) (response DisableAutonomousDatabaseInsightAdvancedFeaturesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableAutonomousDatabaseInsightAdvancedFeatures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAutonomousDatabaseInsightAdvancedFeaturesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAutonomousDatabaseInsightAdvancedFeaturesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAutonomousDatabaseInsightAdvancedFeaturesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAutonomousDatabaseInsightAdvancedFeaturesResponse")
	}
	return
}

// disableAutonomousDatabaseInsightAdvancedFeatures implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) disableAutonomousDatabaseInsightAdvancedFeatures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/disableAutonomousDatabaseInsightAdvancedFeatures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAutonomousDatabaseInsightAdvancedFeaturesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/DisableAutonomousDatabaseInsightAdvancedFeatures"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DisableAutonomousDatabaseInsightAdvancedFeatures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableAwrHubSource Disables a Awr Hub source database in Operations Insights. This will stop the Awr data flow for the given Awr Hub source.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableAwrHubSource.go.html to see an example of how to use DisableAwrHubSource API.
// A default retry strategy applies to this operation DisableAwrHubSource()
func (client OperationsInsightsClient) DisableAwrHubSource(ctx context.Context, request DisableAwrHubSourceRequest) (response DisableAwrHubSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableAwrHubSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAwrHubSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAwrHubSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAwrHubSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAwrHubSourceResponse")
	}
	return
}

// disableAwrHubSource implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) disableAwrHubSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/awrHubSources/{awrHubSourceId}/actions/disable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAwrHubSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/DisableAwrHubSource"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DisableAwrHubSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableDatabaseInsight Disables a database in Operations Insights. Database metric collection and analysis will be stopped.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableDatabaseInsight.go.html to see an example of how to use DisableDatabaseInsight API.
// A default retry strategy applies to this operation DisableDatabaseInsight()
func (client OperationsInsightsClient) DisableDatabaseInsight(ctx context.Context, request DisableDatabaseInsightRequest) (response DisableDatabaseInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/DisableDatabaseInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DisableDatabaseInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExadataInsight Disables an Exadata system in Operations Insights. Exadata-related metric collection and analysis will be stopped.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableExadataInsight.go.html to see an example of how to use DisableExadataInsight API.
// A default retry strategy applies to this operation DisableExadataInsight()
func (client OperationsInsightsClient) DisableExadataInsight(ctx context.Context, request DisableExadataInsightRequest) (response DisableExadataInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/DisableExadataInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DisableExadataInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableHostInsight Disables a host in Ops Insights. Host metric collection and analysis will be stopped.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DisableHostInsight.go.html to see an example of how to use DisableHostInsight API.
// A default retry strategy applies to this operation DisableHostInsight()
func (client OperationsInsightsClient) DisableHostInsight(ctx context.Context, request DisableHostInsightRequest) (response DisableHostInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/DisableHostInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DisableHostInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadOperationsInsightsWarehouseWallet Download the ADW wallet for Operations Insights Warehouse using which the Hub data is exposed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/DownloadOperationsInsightsWarehouseWallet.go.html to see an example of how to use DownloadOperationsInsightsWarehouseWallet API.
// A default retry strategy applies to this operation DownloadOperationsInsightsWarehouseWallet()
func (client OperationsInsightsClient) DownloadOperationsInsightsWarehouseWallet(ctx context.Context, request DownloadOperationsInsightsWarehouseWalletRequest) (response DownloadOperationsInsightsWarehouseWalletResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/DownloadOperationsInsightsWarehouseWallet"
		err = common.PostProcessServiceError(err, "OperationsInsights", "DownloadOperationsInsightsWarehouseWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableAutonomousDatabaseInsightAdvancedFeatures Enables advanced features for an Autonomous Database in Operations Insights. A direct connection will be available for further collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableAutonomousDatabaseInsightAdvancedFeatures.go.html to see an example of how to use EnableAutonomousDatabaseInsightAdvancedFeatures API.
// A default retry strategy applies to this operation EnableAutonomousDatabaseInsightAdvancedFeatures()
func (client OperationsInsightsClient) EnableAutonomousDatabaseInsightAdvancedFeatures(ctx context.Context, request EnableAutonomousDatabaseInsightAdvancedFeaturesRequest) (response EnableAutonomousDatabaseInsightAdvancedFeaturesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableAutonomousDatabaseInsightAdvancedFeatures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableAutonomousDatabaseInsightAdvancedFeaturesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableAutonomousDatabaseInsightAdvancedFeaturesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableAutonomousDatabaseInsightAdvancedFeaturesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableAutonomousDatabaseInsightAdvancedFeaturesResponse")
	}
	return
}

// enableAutonomousDatabaseInsightAdvancedFeatures implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) enableAutonomousDatabaseInsightAdvancedFeatures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/{databaseInsightId}/actions/enableAutonomousDatabaseInsightAdvancedFeatures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableAutonomousDatabaseInsightAdvancedFeaturesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/EnableAutonomousDatabaseInsightAdvancedFeatures"
		err = common.PostProcessServiceError(err, "OperationsInsights", "EnableAutonomousDatabaseInsightAdvancedFeatures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableAwrHubSource Enables a Awr Hub source database in Operations Insights. This will resume the Awr data flow for the given Awr Hub source if it was stopped earlier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableAwrHubSource.go.html to see an example of how to use EnableAwrHubSource API.
// A default retry strategy applies to this operation EnableAwrHubSource()
func (client OperationsInsightsClient) EnableAwrHubSource(ctx context.Context, request EnableAwrHubSourceRequest) (response EnableAwrHubSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableAwrHubSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableAwrHubSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableAwrHubSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableAwrHubSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableAwrHubSourceResponse")
	}
	return
}

// enableAwrHubSource implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) enableAwrHubSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/awrHubSources/{awrHubSourceId}/actions/enable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableAwrHubSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/EnableAwrHubSource"
		err = common.PostProcessServiceError(err, "OperationsInsights", "EnableAwrHubSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableDatabaseInsight Enables a database in Operations Insights. Database metric collection and analysis will be started.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableDatabaseInsight.go.html to see an example of how to use EnableDatabaseInsight API.
// A default retry strategy applies to this operation EnableDatabaseInsight()
func (client OperationsInsightsClient) EnableDatabaseInsight(ctx context.Context, request EnableDatabaseInsightRequest) (response EnableDatabaseInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/EnableDatabaseInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "EnableDatabaseInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableExadataInsight Enables an Exadata system in Operations Insights. Exadata-related metric collection and analysis will be started.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableExadataInsight.go.html to see an example of how to use EnableExadataInsight API.
// A default retry strategy applies to this operation EnableExadataInsight()
func (client OperationsInsightsClient) EnableExadataInsight(ctx context.Context, request EnableExadataInsightRequest) (response EnableExadataInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/EnableExadataInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "EnableExadataInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableHostInsight Enables a host in Ops Insights. Host metric collection and analysis will be started.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/EnableHostInsight.go.html to see an example of how to use EnableHostInsight API.
// A default retry strategy applies to this operation EnableHostInsight()
func (client OperationsInsightsClient) EnableHostInsight(ctx context.Context, request EnableHostInsightRequest) (response EnableHostInsightResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/EnableHostInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "EnableHostInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrDatabaseReport Gets the AWR report for the specified database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrDatabaseReport.go.html to see an example of how to use GetAwrDatabaseReport API.
// A default retry strategy applies to this operation GetAwrDatabaseReport()
func (client OperationsInsightsClient) GetAwrDatabaseReport(ctx context.Context, request GetAwrDatabaseReportRequest) (response GetAwrDatabaseReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAwrDatabaseReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrDatabaseReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrDatabaseReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrDatabaseReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrDatabaseReportResponse")
	}
	return
}

// getAwrDatabaseReport implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getAwrDatabaseReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrDatabaseReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/GetAwrDatabaseReport"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetAwrDatabaseReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrDatabaseSqlReport Gets the SQL health check report for one SQL of the specified database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrDatabaseSqlReport.go.html to see an example of how to use GetAwrDatabaseSqlReport API.
// A default retry strategy applies to this operation GetAwrDatabaseSqlReport()
func (client OperationsInsightsClient) GetAwrDatabaseSqlReport(ctx context.Context, request GetAwrDatabaseSqlReportRequest) (response GetAwrDatabaseSqlReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAwrDatabaseSqlReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrDatabaseSqlReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrDatabaseSqlReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrDatabaseSqlReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrDatabaseSqlReportResponse")
	}
	return
}

// getAwrDatabaseSqlReport implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getAwrDatabaseSqlReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseSqlReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrDatabaseSqlReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/GetAwrDatabaseSqlReport"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetAwrDatabaseSqlReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrHub Gets details of an AWR hub.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrHub.go.html to see an example of how to use GetAwrHub API.
// A default retry strategy applies to this operation GetAwrHub()
func (client OperationsInsightsClient) GetAwrHub(ctx context.Context, request GetAwrHubRequest) (response GetAwrHubResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/GetAwrHub"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetAwrHub", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrHubObject Gets the Awr Hub object metadata and body.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrHubObject.go.html to see an example of how to use GetAwrHubObject API.
// A default retry strategy applies to this operation GetAwrHubObject()
func (client OperationsInsightsClient) GetAwrHubObject(ctx context.Context, request GetAwrHubObjectRequest) (response GetAwrHubObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAwrHubObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrHubObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrHubObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrHubObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrHubObjectResponse")
	}
	return
}

// getAwrHubObject implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getAwrHubObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubObjects/awrHubSources/{awrHubSourceId}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrHubObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubObjects/GetAwrHubObject"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetAwrHubObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrHubSource Gets the Awr Hub source object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrHubSource.go.html to see an example of how to use GetAwrHubSource API.
// A default retry strategy applies to this operation GetAwrHubSource()
func (client OperationsInsightsClient) GetAwrHubSource(ctx context.Context, request GetAwrHubSourceRequest) (response GetAwrHubSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAwrHubSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAwrHubSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAwrHubSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAwrHubSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAwrHubSourceResponse")
	}
	return
}

// getAwrHubSource implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getAwrHubSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubSources/{awrHubSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAwrHubSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/GetAwrHubSource"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetAwrHubSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAwrReport Gets the AWR report for the specified source database in the AWR hub. The difference between the timeGreaterThanOrEqualTo and timeLessThanOrEqualTo should not be greater than 7 days.
// Either beginSnapshotIdentifierGreaterThanOrEqualTo & endSnapshotIdentifierLessThanOrEqualTo params Or timeGreaterThanOrEqualTo & timeLessThanOrEqualTo params are required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetAwrReport.go.html to see an example of how to use GetAwrReport API.
// A default retry strategy applies to this operation GetAwrReport()
func (client OperationsInsightsClient) GetAwrReport(ctx context.Context, request GetAwrReportRequest) (response GetAwrReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/GetAwrReport"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetAwrReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseInsight Gets details of a database insight.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetDatabaseInsight.go.html to see an example of how to use GetDatabaseInsight API.
// A default retry strategy applies to this operation GetDatabaseInsight()
func (client OperationsInsightsClient) GetDatabaseInsight(ctx context.Context, request GetDatabaseInsightRequest) (response GetDatabaseInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/GetDatabaseInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetDatabaseInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &databaseinsight{})
	return response, err
}

// GetEnterpriseManagerBridge Gets details of an Operations Insights Enterprise Manager bridge.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetEnterpriseManagerBridge.go.html to see an example of how to use GetEnterpriseManagerBridge API.
// A default retry strategy applies to this operation GetEnterpriseManagerBridge()
func (client OperationsInsightsClient) GetEnterpriseManagerBridge(ctx context.Context, request GetEnterpriseManagerBridgeRequest) (response GetEnterpriseManagerBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/EnterpriseManagerBridges/GetEnterpriseManagerBridge"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetEnterpriseManagerBridge", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetExadataInsight Gets details of an Exadata insight.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetExadataInsight.go.html to see an example of how to use GetExadataInsight API.
// A default retry strategy applies to this operation GetExadataInsight()
func (client OperationsInsightsClient) GetExadataInsight(ctx context.Context, request GetExadataInsightRequest) (response GetExadataInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/GetExadataInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetExadataInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &exadatainsight{})
	return response, err
}

// GetHostInsight Gets details of a host insight.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetHostInsight.go.html to see an example of how to use GetHostInsight API.
// A default retry strategy applies to this operation GetHostInsight()
func (client OperationsInsightsClient) GetHostInsight(ctx context.Context, request GetHostInsightRequest) (response GetHostInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/GetHostInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetHostInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &hostinsight{})
	return response, err
}

// GetNewsReport Gets details of a news report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetNewsReport.go.html to see an example of how to use GetNewsReport API.
// A default retry strategy applies to this operation GetNewsReport()
func (client OperationsInsightsClient) GetNewsReport(ctx context.Context, request GetNewsReportRequest) (response GetNewsReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNewsReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNewsReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNewsReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNewsReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNewsReportResponse")
	}
	return
}

// getNewsReport implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getNewsReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/newsReports/{newsReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNewsReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/NewsReports/GetNewsReport"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetNewsReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOperationsInsightsPrivateEndpoint Gets the details of the specified private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOperationsInsightsPrivateEndpoint.go.html to see an example of how to use GetOperationsInsightsPrivateEndpoint API.
// A default retry strategy applies to this operation GetOperationsInsightsPrivateEndpoint()
func (client OperationsInsightsClient) GetOperationsInsightsPrivateEndpoint(ctx context.Context, request GetOperationsInsightsPrivateEndpointRequest) (response GetOperationsInsightsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOperationsInsightsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOperationsInsightsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOperationsInsightsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOperationsInsightsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOperationsInsightsPrivateEndpointResponse")
	}
	return
}

// getOperationsInsightsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getOperationsInsightsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operationsInsightsPrivateEndpoints/{operationsInsightsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOperationsInsightsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsPrivateEndpoint/GetOperationsInsightsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetOperationsInsightsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOperationsInsightsWarehouse Gets details of an Ops Insights Warehouse.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOperationsInsightsWarehouse.go.html to see an example of how to use GetOperationsInsightsWarehouse API.
// A default retry strategy applies to this operation GetOperationsInsightsWarehouse()
func (client OperationsInsightsClient) GetOperationsInsightsWarehouse(ctx context.Context, request GetOperationsInsightsWarehouseRequest) (response GetOperationsInsightsWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/GetOperationsInsightsWarehouse"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetOperationsInsightsWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOperationsInsightsWarehouseUser Gets details of an Operations Insights Warehouse User.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOperationsInsightsWarehouseUser.go.html to see an example of how to use GetOperationsInsightsWarehouseUser API.
// A default retry strategy applies to this operation GetOperationsInsightsWarehouseUser()
func (client OperationsInsightsClient) GetOperationsInsightsWarehouseUser(ctx context.Context, request GetOperationsInsightsWarehouseUserRequest) (response GetOperationsInsightsWarehouseUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouseUsers/GetOperationsInsightsWarehouseUser"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetOperationsInsightsWarehouseUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOpsiConfiguration Gets details of an OPSI configuration resource.
// Values specified in configItemField and configItemCustomStatus query params will be considered, only if configItems field is requested as part of opsiConfigField query param.
// Values specified in configItemCustomStatus will determine whether only customized configuration items or only non-customized configuration items or both have to be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOpsiConfiguration.go.html to see an example of how to use GetOpsiConfiguration API.
// A default retry strategy applies to this operation GetOpsiConfiguration()
func (client OperationsInsightsClient) GetOpsiConfiguration(ctx context.Context, request GetOpsiConfigurationRequest) (response GetOpsiConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOpsiConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOpsiConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOpsiConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOpsiConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOpsiConfigurationResponse")
	}
	return
}

// getOpsiConfiguration implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getOpsiConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opsiConfigurations/{opsiConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOpsiConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiConfigurations/GetOpsiConfiguration"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetOpsiConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &opsiconfiguration{})
	return response, err
}

// GetOpsiDataObject Gets details of an OPSI data object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOpsiDataObject.go.html to see an example of how to use GetOpsiDataObject API.
// A default retry strategy applies to this operation GetOpsiDataObject()
func (client OperationsInsightsClient) GetOpsiDataObject(ctx context.Context, request GetOpsiDataObjectRequest) (response GetOpsiDataObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOpsiDataObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOpsiDataObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOpsiDataObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOpsiDataObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOpsiDataObjectResponse")
	}
	return
}

// getOpsiDataObject implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) getOpsiDataObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opsiDataObjects/{opsiDataObjectIdentifier}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOpsiDataObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiDataObjects/GetOpsiDataObject"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetOpsiDataObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &opsidataobject{})
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client OperationsInsightsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/WorkRequests/GetWorkRequest"
		err = common.PostProcessServiceError(err, "OperationsInsights", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// HeadAwrHubObject Gets the Awr Hub object's user-defined metadata and entity tag (ETag).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/HeadAwrHubObject.go.html to see an example of how to use HeadAwrHubObject API.
// A default retry strategy applies to this operation HeadAwrHubObject()
func (client OperationsInsightsClient) HeadAwrHubObject(ctx context.Context, request HeadAwrHubObjectRequest) (response HeadAwrHubObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.headAwrHubObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = HeadAwrHubObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = HeadAwrHubObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(HeadAwrHubObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into HeadAwrHubObjectResponse")
	}
	return
}

// headAwrHubObject implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) headAwrHubObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodHead, "/awrHubObjects/awrHubSources/{awrHubSourceId}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response HeadAwrHubObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubObjects/HeadAwrHubObject"
		err = common.PostProcessServiceError(err, "OperationsInsights", "HeadAwrHubObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestAddmReports This endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestAddmReports.go.html to see an example of how to use IngestAddmReports API.
// A default retry strategy applies to this operation IngestAddmReports()
func (client OperationsInsightsClient) IngestAddmReports(ctx context.Context, request IngestAddmReportsRequest) (response IngestAddmReportsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestAddmReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestAddmReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestAddmReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestAddmReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestAddmReportsResponse")
	}
	return
}

// ingestAddmReports implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestAddmReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestAddmReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestAddmReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestAddmReports"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestAddmReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestDatabaseConfiguration This is a generic ingest endpoint for all database configuration metrics.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestDatabaseConfiguration.go.html to see an example of how to use IngestDatabaseConfiguration API.
// A default retry strategy applies to this operation IngestDatabaseConfiguration()
func (client OperationsInsightsClient) IngestDatabaseConfiguration(ctx context.Context, request IngestDatabaseConfigurationRequest) (response IngestDatabaseConfigurationResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestDatabaseConfiguration"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestDatabaseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestHostConfiguration This is a generic ingest endpoint for all the host configuration metrics
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestHostConfiguration.go.html to see an example of how to use IngestHostConfiguration API.
// A default retry strategy applies to this operation IngestHostConfiguration()
func (client OperationsInsightsClient) IngestHostConfiguration(ctx context.Context, request IngestHostConfigurationRequest) (response IngestHostConfigurationResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/IngestHostConfiguration"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestHostConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestHostMetrics This is a generic ingest endpoint for all the host performance metrics
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestHostMetrics.go.html to see an example of how to use IngestHostMetrics API.
// A default retry strategy applies to this operation IngestHostMetrics()
func (client OperationsInsightsClient) IngestHostMetrics(ctx context.Context, request IngestHostMetricsRequest) (response IngestHostMetricsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/IngestHostMetrics"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestHostMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestMySqlSqlStats The MySql SQL Stats endpoint takes in a JSON payload, persists it in Ops Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestMySqlSqlStats.go.html to see an example of how to use IngestMySqlSqlStats API.
// A default retry strategy applies to this operation IngestMySqlSqlStats()
func (client OperationsInsightsClient) IngestMySqlSqlStats(ctx context.Context, request IngestMySqlSqlStatsRequest) (response IngestMySqlSqlStatsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestMySqlSqlStats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestMySqlSqlStatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestMySqlSqlStatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestMySqlSqlStatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestMySqlSqlStatsResponse")
	}
	return
}

// ingestMySqlSqlStats implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestMySqlSqlStats(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestMySqlSqlStatsMetric", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestMySqlSqlStatsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestMySqlSqlStats"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestMySqlSqlStats", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestMySqlSqlText The SqlText endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
// Either databaseId or id must be specified.
// Disclaimer: SQL text being uploaded explicitly via APIs is already masked. All sensitive literals contained in the sqlFullText column are masked prior to ingestion.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestMySqlSqlText.go.html to see an example of how to use IngestMySqlSqlText API.
// A default retry strategy applies to this operation IngestMySqlSqlText()
func (client OperationsInsightsClient) IngestMySqlSqlText(ctx context.Context, request IngestMySqlSqlTextRequest) (response IngestMySqlSqlTextResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestMySqlSqlText, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestMySqlSqlTextResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestMySqlSqlTextResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestMySqlSqlTextResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestMySqlSqlTextResponse")
	}
	return
}

// ingestMySqlSqlText implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) ingestMySqlSqlText(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/ingestMySqlSqlText", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestMySqlSqlTextResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestMySqlSqlText"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestMySqlSqlText", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestSqlBucket The sqlbucket endpoint takes in a JSON payload, persists it in Ops Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlBucket.go.html to see an example of how to use IngestSqlBucket API.
// A default retry strategy applies to this operation IngestSqlBucket()
func (client OperationsInsightsClient) IngestSqlBucket(ctx context.Context, request IngestSqlBucketRequest) (response IngestSqlBucketResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestSqlBucket"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestSqlBucket", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestSqlPlanLines The SqlPlanLines endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlPlanLines.go.html to see an example of how to use IngestSqlPlanLines API.
// A default retry strategy applies to this operation IngestSqlPlanLines()
func (client OperationsInsightsClient) IngestSqlPlanLines(ctx context.Context, request IngestSqlPlanLinesRequest) (response IngestSqlPlanLinesResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestSqlPlanLines"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestSqlPlanLines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestSqlStats The SQL Stats endpoint takes in a JSON payload, persists it in Ops Insights ingest pipeline.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlStats.go.html to see an example of how to use IngestSqlStats API.
// A default retry strategy applies to this operation IngestSqlStats()
func (client OperationsInsightsClient) IngestSqlStats(ctx context.Context, request IngestSqlStatsRequest) (response IngestSqlStatsResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestSqlStats"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestSqlStats", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IngestSqlText The SqlText endpoint takes in a JSON payload, persists it in Operation Insights ingest pipeline.
// Either databaseId or id must be specified.
// Disclaimer: SQL text being uploaded explicitly via APIs is not masked. Any sensitive literals contained in the sqlFullText column should be masked prior to ingestion.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/IngestSqlText.go.html to see an example of how to use IngestSqlText API.
// A default retry strategy applies to this operation IngestSqlText()
func (client OperationsInsightsClient) IngestSqlText(ctx context.Context, request IngestSqlTextRequest) (response IngestSqlTextResponse, err error) {
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/IngestSqlText"
		err = common.PostProcessServiceError(err, "OperationsInsights", "IngestSqlText", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddmDbFindingCategories Gets list of ADDM finding categories.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAddmDbFindingCategories.go.html to see an example of how to use ListAddmDbFindingCategories API.
// A default retry strategy applies to this operation ListAddmDbFindingCategories()
func (client OperationsInsightsClient) ListAddmDbFindingCategories(ctx context.Context, request ListAddmDbFindingCategoriesRequest) (response ListAddmDbFindingCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddmDbFindingCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddmDbFindingCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddmDbFindingCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddmDbFindingCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddmDbFindingCategoriesResponse")
	}
	return
}

// listAddmDbFindingCategories implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAddmDbFindingCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbFindingCategories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddmDbFindingCategoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListAddmDbFindingCategories"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAddmDbFindingCategories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddmDbFindingsTimeSeries Get the ADDM findings time series for the specified databases for a given time period.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAddmDbFindingsTimeSeries.go.html to see an example of how to use ListAddmDbFindingsTimeSeries API.
// A default retry strategy applies to this operation ListAddmDbFindingsTimeSeries()
func (client OperationsInsightsClient) ListAddmDbFindingsTimeSeries(ctx context.Context, request ListAddmDbFindingsTimeSeriesRequest) (response ListAddmDbFindingsTimeSeriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddmDbFindingsTimeSeries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddmDbFindingsTimeSeriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddmDbFindingsTimeSeriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddmDbFindingsTimeSeriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddmDbFindingsTimeSeriesResponse")
	}
	return
}

// listAddmDbFindingsTimeSeries implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAddmDbFindingsTimeSeries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbFindingsTimeSeries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddmDbFindingsTimeSeriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListAddmDbFindingsTimeSeries"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAddmDbFindingsTimeSeries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddmDbParameterCategories Gets list of ADDM database parameter categories for the specified databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAddmDbParameterCategories.go.html to see an example of how to use ListAddmDbParameterCategories API.
// A default retry strategy applies to this operation ListAddmDbParameterCategories()
func (client OperationsInsightsClient) ListAddmDbParameterCategories(ctx context.Context, request ListAddmDbParameterCategoriesRequest) (response ListAddmDbParameterCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddmDbParameterCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddmDbParameterCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddmDbParameterCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddmDbParameterCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddmDbParameterCategoriesResponse")
	}
	return
}

// listAddmDbParameterCategories implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAddmDbParameterCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbParameterCategories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddmDbParameterCategoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListAddmDbParameterCategories"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAddmDbParameterCategories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddmDbRecommendationCategories Gets list of ADDM recommendation categories for the specified databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAddmDbRecommendationCategories.go.html to see an example of how to use ListAddmDbRecommendationCategories API.
// A default retry strategy applies to this operation ListAddmDbRecommendationCategories()
func (client OperationsInsightsClient) ListAddmDbRecommendationCategories(ctx context.Context, request ListAddmDbRecommendationCategoriesRequest) (response ListAddmDbRecommendationCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddmDbRecommendationCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddmDbRecommendationCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddmDbRecommendationCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddmDbRecommendationCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddmDbRecommendationCategoriesResponse")
	}
	return
}

// listAddmDbRecommendationCategories implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAddmDbRecommendationCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbRecommendationCategories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddmDbRecommendationCategoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListAddmDbRecommendationCategories"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAddmDbRecommendationCategories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddmDbRecommendationsTimeSeries Gets time series data for ADDM recommendations for the specified databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAddmDbRecommendationsTimeSeries.go.html to see an example of how to use ListAddmDbRecommendationsTimeSeries API.
// A default retry strategy applies to this operation ListAddmDbRecommendationsTimeSeries()
func (client OperationsInsightsClient) ListAddmDbRecommendationsTimeSeries(ctx context.Context, request ListAddmDbRecommendationsTimeSeriesRequest) (response ListAddmDbRecommendationsTimeSeriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddmDbRecommendationsTimeSeries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddmDbRecommendationsTimeSeriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddmDbRecommendationsTimeSeriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddmDbRecommendationsTimeSeriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddmDbRecommendationsTimeSeriesResponse")
	}
	return
}

// listAddmDbRecommendationsTimeSeries implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAddmDbRecommendationsTimeSeries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbRecommendationsTimeSeries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddmDbRecommendationsTimeSeriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListAddmDbRecommendationsTimeSeries"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAddmDbRecommendationsTimeSeries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddmDbs Gets a list of ADDM database information
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAddmDbs.go.html to see an example of how to use ListAddmDbs API.
// A default retry strategy applies to this operation ListAddmDbs()
func (client OperationsInsightsClient) ListAddmDbs(ctx context.Context, request ListAddmDbsRequest) (response ListAddmDbsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddmDbs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddmDbsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddmDbsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddmDbsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddmDbsResponse")
	}
	return
}

// listAddmDbs implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAddmDbs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddmDbsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListAddmDbs"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAddmDbs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrDatabaseSnapshots Lists AWR snapshots for the specified database in the AWR.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrDatabaseSnapshots.go.html to see an example of how to use ListAwrDatabaseSnapshots API.
// A default retry strategy applies to this operation ListAwrDatabaseSnapshots()
func (client OperationsInsightsClient) ListAwrDatabaseSnapshots(ctx context.Context, request ListAwrDatabaseSnapshotsRequest) (response ListAwrDatabaseSnapshotsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAwrDatabaseSnapshots, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrDatabaseSnapshotsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrDatabaseSnapshotsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrDatabaseSnapshotsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrDatabaseSnapshotsResponse")
	}
	return
}

// listAwrDatabaseSnapshots implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAwrDatabaseSnapshots(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseSnapshots", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrDatabaseSnapshotsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/ListAwrDatabaseSnapshots"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAwrDatabaseSnapshots", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrDatabases Gets the list of databases and their snapshot summary details available in the AWRHub.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrDatabases.go.html to see an example of how to use ListAwrDatabases API.
// A default retry strategy applies to this operation ListAwrDatabases()
func (client OperationsInsightsClient) ListAwrDatabases(ctx context.Context, request ListAwrDatabasesRequest) (response ListAwrDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAwrDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrDatabasesResponse")
	}
	return
}

// listAwrDatabases implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAwrDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/ListAwrDatabases"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAwrDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrHubObjects Gets a list of Awr Hub objects. Awr Hub id needs to specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrHubObjects.go.html to see an example of how to use ListAwrHubObjects API.
// A default retry strategy applies to this operation ListAwrHubObjects()
func (client OperationsInsightsClient) ListAwrHubObjects(ctx context.Context, request ListAwrHubObjectsRequest) (response ListAwrHubObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAwrHubObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrHubObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrHubObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrHubObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrHubObjectsResponse")
	}
	return
}

// listAwrHubObjects implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAwrHubObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubObjects/awrHubSources/{awrHubSourceId}/o", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrHubObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubObjects/ListAwrHubObjects"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAwrHubObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrHubSources Gets a list of Awr Hub source objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrHubSources.go.html to see an example of how to use ListAwrHubSources API.
// A default retry strategy applies to this operation ListAwrHubSources()
func (client OperationsInsightsClient) ListAwrHubSources(ctx context.Context, request ListAwrHubSourcesRequest) (response ListAwrHubSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAwrHubSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAwrHubSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAwrHubSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAwrHubSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAwrHubSourcesResponse")
	}
	return
}

// listAwrHubSources implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listAwrHubSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAwrHubSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/ListAwrHubSources"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAwrHubSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrHubs Gets a list of AWR hubs. Either compartmentId or id must be specified. All these resources are expected to be in root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrHubs.go.html to see an example of how to use ListAwrHubs API.
// A default retry strategy applies to this operation ListAwrHubs()
func (client OperationsInsightsClient) ListAwrHubs(ctx context.Context, request ListAwrHubsRequest) (response ListAwrHubsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/ListAwrHubs"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAwrHubs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAwrSnapshots Lists AWR snapshots for the specified source database in the AWR hub. The difference between the timeGreaterThanOrEqualTo and timeLessThanOrEqualTo should not exceed an elapsed range of 1 day.
// The timeGreaterThanOrEqualTo & timeLessThanOrEqualTo params are optional. If these params are not provided, by default last 1 day snapshots will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrSnapshots.go.html to see an example of how to use ListAwrSnapshots API.
// A default retry strategy applies to this operation ListAwrSnapshots()
func (client OperationsInsightsClient) ListAwrSnapshots(ctx context.Context, request ListAwrSnapshotsRequest) (response ListAwrSnapshotsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/ListAwrSnapshots"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListAwrSnapshots", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseConfigurations Gets a list of database insight configurations based on the query parameters specified. Either compartmentId or databaseInsightId query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of database insight configurations in that compartment and in all sub-compartments will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseConfigurations.go.html to see an example of how to use ListDatabaseConfigurations API.
// A default retry strategy applies to this operation ListDatabaseConfigurations()
func (client OperationsInsightsClient) ListDatabaseConfigurations(ctx context.Context, request ListDatabaseConfigurationsRequest) (response ListDatabaseConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListDatabaseConfigurations"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListDatabaseConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseInsights Gets a list of database insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of database insights in that compartment and in all sub-compartments will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListDatabaseInsights.go.html to see an example of how to use ListDatabaseInsights API.
// A default retry strategy applies to this operation ListDatabaseInsights()
func (client OperationsInsightsClient) ListDatabaseInsights(ctx context.Context, request ListDatabaseInsightsRequest) (response ListDatabaseInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListDatabaseInsights"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListDatabaseInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEnterpriseManagerBridges Gets a list of Ops Insights Enterprise Manager bridges. Either compartmentId or id must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of bridges in that compartment and in all sub-compartments will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListEnterpriseManagerBridges.go.html to see an example of how to use ListEnterpriseManagerBridges API.
// A default retry strategy applies to this operation ListEnterpriseManagerBridges()
func (client OperationsInsightsClient) ListEnterpriseManagerBridges(ctx context.Context, request ListEnterpriseManagerBridgesRequest) (response ListEnterpriseManagerBridgesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/EnterpriseManagerBridges/ListEnterpriseManagerBridges"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListEnterpriseManagerBridges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExadataConfigurations Gets a list of exadata insight configurations. Either compartmentId or exadataInsightsId query parameter must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListExadataConfigurations.go.html to see an example of how to use ListExadataConfigurations API.
// A default retry strategy applies to this operation ListExadataConfigurations()
func (client OperationsInsightsClient) ListExadataConfigurations(ctx context.Context, request ListExadataConfigurationsRequest) (response ListExadataConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/ListExadataConfigurations"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListExadataConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExadataInsights Gets a list of Exadata insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of Exadata insights in that compartment and in all sub-compartments will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListExadataInsights.go.html to see an example of how to use ListExadataInsights API.
// A default retry strategy applies to this operation ListExadataInsights()
func (client OperationsInsightsClient) ListExadataInsights(ctx context.Context, request ListExadataInsightsRequest) (response ListExadataInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/ListExadataInsights"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListExadataInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHostConfigurations Gets a list of host insight configurations based on the query parameters specified. Either compartmentId or hostInsightId query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of host insight configurations in that compartment and in all sub-compartments will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostConfigurations.go.html to see an example of how to use ListHostConfigurations API.
// A default retry strategy applies to this operation ListHostConfigurations()
func (client OperationsInsightsClient) ListHostConfigurations(ctx context.Context, request ListHostConfigurationsRequest) (response ListHostConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/ListHostConfigurations"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListHostConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHostInsights Gets a list of host insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
// When both compartmentId and compartmentIdInSubtree are specified, a list of host insights in that compartment and in all sub-compartments will be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostInsights.go.html to see an example of how to use ListHostInsights API.
// A default retry strategy applies to this operation ListHostInsights()
func (client OperationsInsightsClient) ListHostInsights(ctx context.Context, request ListHostInsightsRequest) (response ListHostInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/ListHostInsights"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListHostInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHostedEntities Get a list of hosted entities details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListHostedEntities.go.html to see an example of how to use ListHostedEntities API.
// A default retry strategy applies to this operation ListHostedEntities()
func (client OperationsInsightsClient) ListHostedEntities(ctx context.Context, request ListHostedEntitiesRequest) (response ListHostedEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/ListHostedEntities"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListHostedEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListImportableAgentEntities Gets a list of agent entities available to add a new hostInsight.  An agent entity is "available"
// and will be shown if all the following conditions are true:
//  1. The agent OCID is not already being used for an existing hostInsight.
//  2. The agent availabilityStatus = 'ACTIVE'
//  3. The agent lifecycleState = 'ACTIVE'
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableAgentEntities.go.html to see an example of how to use ListImportableAgentEntities API.
// A default retry strategy applies to this operation ListImportableAgentEntities()
func (client OperationsInsightsClient) ListImportableAgentEntities(ctx context.Context, request ListImportableAgentEntitiesRequest) (response ListImportableAgentEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/ListImportableAgentEntities"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListImportableAgentEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListImportableComputeEntities Gets a list of available compute intances running cloud agent to add a new hostInsight.  An Compute entity is "available"
// and will be shown if all the following conditions are true:
//  1. Compute is running OCA
//  2. OCI Management Agent is not enabled or If OCI Management Agent is enabled
//     2.1 The agent OCID is not already being used for an existing hostInsight.
//     2.2 The agent availabilityStatus = 'ACTIVE'
//     2.3 The agent lifecycleState = 'ACTIVE'
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableComputeEntities.go.html to see an example of how to use ListImportableComputeEntities API.
// A default retry strategy applies to this operation ListImportableComputeEntities()
func (client OperationsInsightsClient) ListImportableComputeEntities(ctx context.Context, request ListImportableComputeEntitiesRequest) (response ListImportableComputeEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listImportableComputeEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListImportableComputeEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListImportableComputeEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListImportableComputeEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListImportableComputeEntitiesResponse")
	}
	return
}

// listImportableComputeEntities implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listImportableComputeEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/importableComputeEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListImportableComputeEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/ListImportableComputeEntities"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListImportableComputeEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListImportableEnterpriseManagerEntities Gets a list of importable entities for an Operations Insights Enterprise Manager bridge that have not been imported before.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListImportableEnterpriseManagerEntities.go.html to see an example of how to use ListImportableEnterpriseManagerEntities API.
// A default retry strategy applies to this operation ListImportableEnterpriseManagerEntities()
func (client OperationsInsightsClient) ListImportableEnterpriseManagerEntities(ctx context.Context, request ListImportableEnterpriseManagerEntitiesRequest) (response ListImportableEnterpriseManagerEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/EnterpriseManagerBridges/ListImportableEnterpriseManagerEntities"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListImportableEnterpriseManagerEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNewsReports Gets a list of news reports based on the query parameters specified. Either compartmentId or id query parameter must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListNewsReports.go.html to see an example of how to use ListNewsReports API.
// A default retry strategy applies to this operation ListNewsReports()
func (client OperationsInsightsClient) ListNewsReports(ctx context.Context, request ListNewsReportsRequest) (response ListNewsReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNewsReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNewsReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNewsReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNewsReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNewsReportsResponse")
	}
	return
}

// listNewsReports implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listNewsReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/newsReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNewsReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/NewsReport/ListNewsReports"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListNewsReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOperationsInsightsPrivateEndpoints Gets a list of Operation Insights private endpoints.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsPrivateEndpoints.go.html to see an example of how to use ListOperationsInsightsPrivateEndpoints API.
// A default retry strategy applies to this operation ListOperationsInsightsPrivateEndpoints()
func (client OperationsInsightsClient) ListOperationsInsightsPrivateEndpoints(ctx context.Context, request ListOperationsInsightsPrivateEndpointsRequest) (response ListOperationsInsightsPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOperationsInsightsPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOperationsInsightsPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOperationsInsightsPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOperationsInsightsPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOperationsInsightsPrivateEndpointsResponse")
	}
	return
}

// listOperationsInsightsPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listOperationsInsightsPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operationsInsightsPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOperationsInsightsPrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsPrivateEndpoint/ListOperationsInsightsPrivateEndpoints"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListOperationsInsightsPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOperationsInsightsWarehouseUsers Gets a list of Operations Insights Warehouse users. Either compartmentId or id must be specified. All these resources are expected to be in root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsWarehouseUsers.go.html to see an example of how to use ListOperationsInsightsWarehouseUsers API.
// A default retry strategy applies to this operation ListOperationsInsightsWarehouseUsers()
func (client OperationsInsightsClient) ListOperationsInsightsWarehouseUsers(ctx context.Context, request ListOperationsInsightsWarehouseUsersRequest) (response ListOperationsInsightsWarehouseUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouseUsers/ListOperationsInsightsWarehouseUsers"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListOperationsInsightsWarehouseUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOperationsInsightsWarehouses Gets a list of Ops Insights warehouses. Either compartmentId or id must be specified.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsWarehouses.go.html to see an example of how to use ListOperationsInsightsWarehouses API.
// A default retry strategy applies to this operation ListOperationsInsightsWarehouses()
func (client OperationsInsightsClient) ListOperationsInsightsWarehouses(ctx context.Context, request ListOperationsInsightsWarehousesRequest) (response ListOperationsInsightsWarehousesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/ListOperationsInsightsWarehouses"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListOperationsInsightsWarehouses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOpsiConfigurations Gets a list of OPSI configuration resources based on the query parameters specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOpsiConfigurations.go.html to see an example of how to use ListOpsiConfigurations API.
// A default retry strategy applies to this operation ListOpsiConfigurations()
func (client OperationsInsightsClient) ListOpsiConfigurations(ctx context.Context, request ListOpsiConfigurationsRequest) (response ListOpsiConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOpsiConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOpsiConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOpsiConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOpsiConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOpsiConfigurationsResponse")
	}
	return
}

// listOpsiConfigurations implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listOpsiConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opsiConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOpsiConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiConfigurations/ListOpsiConfigurations"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListOpsiConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOpsiDataObjects Gets a list of OPSI data objects based on the query parameters specified. CompartmentId id query parameter must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOpsiDataObjects.go.html to see an example of how to use ListOpsiDataObjects API.
// A default retry strategy applies to this operation ListOpsiDataObjects()
func (client OperationsInsightsClient) ListOpsiDataObjects(ctx context.Context, request ListOpsiDataObjectsRequest) (response ListOpsiDataObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOpsiDataObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOpsiDataObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOpsiDataObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOpsiDataObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOpsiDataObjectsResponse")
	}
	return
}

// listOpsiDataObjects implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listOpsiDataObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opsiDataObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOpsiDataObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiDataObjects/ListOpsiDataObjects"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListOpsiDataObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlPlans Query SQL Warehouse to list the plan xml for a given SQL execution plan. This returns a SqlPlanCollection object, but is currently limited to a single plan.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlPlans.go.html to see an example of how to use ListSqlPlans API.
// A default retry strategy applies to this operation ListSqlPlans()
func (client OperationsInsightsClient) ListSqlPlans(ctx context.Context, request ListSqlPlansRequest) (response ListSqlPlansResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListSqlPlans"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListSqlPlans", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlSearches Search SQL by SQL Identifier across databases in a compartment and in all sub-compartments if specified.
// And get the SQL Text and the details of the databases executing the SQL for a given time period.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlSearches.go.html to see an example of how to use ListSqlSearches API.
// A default retry strategy applies to this operation ListSqlSearches()
func (client OperationsInsightsClient) ListSqlSearches(ctx context.Context, request ListSqlSearchesRequest) (response ListSqlSearchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListSqlSearches"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListSqlSearches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSqlTexts Query SQL Warehouse to get the full SQL Text for a SQL in a compartment and in all sub-compartments if specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListSqlTexts.go.html to see an example of how to use ListSqlTexts API.
// A default retry strategy applies to this operation ListSqlTexts()
func (client OperationsInsightsClient) ListSqlTexts(ctx context.Context, request ListSqlTextsRequest) (response ListSqlTextsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/ListSqlTexts"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListSqlTexts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWarehouseDataObjects Gets a list of Warehouse data objects (e.g: views, tables), based on the query parameters specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWarehouseDataObjects.go.html to see an example of how to use ListWarehouseDataObjects API.
// A default retry strategy applies to this operation ListWarehouseDataObjects()
func (client OperationsInsightsClient) ListWarehouseDataObjects(ctx context.Context, request ListWarehouseDataObjectsRequest) (response ListWarehouseDataObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWarehouseDataObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWarehouseDataObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWarehouseDataObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWarehouseDataObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWarehouseDataObjectsResponse")
	}
	return
}

// listWarehouseDataObjects implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) listWarehouseDataObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/{warehouseType}/{warehouseId}/dataObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWarehouseDataObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiWarehouseDataObjects/ListWarehouseDataObjects"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListWarehouseDataObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client OperationsInsightsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/WorkRequests/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client OperationsInsightsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/WorkRequests/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment. Either compartmentId or id must be specified. Only one of id, resourceId or relatedResourceId can be specified optionally.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client OperationsInsightsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/WorkRequests/ListWorkRequests"
		err = common.PostProcessServiceError(err, "OperationsInsights", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutAwrHubObject Creates a new object or overwrites an existing object with the same name to the Awr Hub.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/PutAwrHubObject.go.html to see an example of how to use PutAwrHubObject API.
// A default retry strategy applies to this operation PutAwrHubObject()
func (client OperationsInsightsClient) PutAwrHubObject(ctx context.Context, request PutAwrHubObjectRequest) (response PutAwrHubObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putAwrHubObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutAwrHubObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutAwrHubObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutAwrHubObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutAwrHubObjectResponse")
	}
	return
}

// putAwrHubObject implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) putAwrHubObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/awrHubObjects/awrHubSources/{awrHubSourceId}/o/{objectName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutAwrHubObjectResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubObjects/PutAwrHubObject"
		err = common.PostProcessServiceError(err, "OperationsInsights", "PutAwrHubObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// QueryOpsiDataObjectData Queries an OPSI data object with the inputs provided and sends the result set back. Either analysisTimeInterval
// or timeIntervalStart and timeIntervalEnd parameters need to be passed as well.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/QueryOpsiDataObjectData.go.html to see an example of how to use QueryOpsiDataObjectData API.
// A default retry strategy applies to this operation QueryOpsiDataObjectData()
func (client OperationsInsightsClient) QueryOpsiDataObjectData(ctx context.Context, request QueryOpsiDataObjectDataRequest) (response QueryOpsiDataObjectDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.queryOpsiDataObjectData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryOpsiDataObjectDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryOpsiDataObjectDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryOpsiDataObjectDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryOpsiDataObjectDataResponse")
	}
	return
}

// queryOpsiDataObjectData implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) queryOpsiDataObjectData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/opsiDataObjects/actions/queryData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response QueryOpsiDataObjectDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiDataObjects/QueryOpsiDataObjectData"
		err = common.PostProcessServiceError(err, "OperationsInsights", "QueryOpsiDataObjectData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &querydataobjectresultsetrowscollection{})
	return response, err
}

// QueryWarehouseDataObjectData Queries Warehouse data objects (e.g: views, tables) with the inputs provided and sends the result set back.
// Any data to which an OperationsInsightsWarehouseUser with a permission to the corresponding Warehouse can be queried.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/QueryWarehouseDataObjectData.go.html to see an example of how to use QueryWarehouseDataObjectData API.
// A default retry strategy applies to this operation QueryWarehouseDataObjectData()
func (client OperationsInsightsClient) QueryWarehouseDataObjectData(ctx context.Context, request QueryWarehouseDataObjectDataRequest) (response QueryWarehouseDataObjectDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.queryWarehouseDataObjectData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryWarehouseDataObjectDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryWarehouseDataObjectDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryWarehouseDataObjectDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryWarehouseDataObjectDataResponse")
	}
	return
}

// queryWarehouseDataObjectData implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) queryWarehouseDataObjectData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/{warehouseType}/{warehouseId}/actions/queryData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response QueryWarehouseDataObjectDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiWarehouseDataObjects/QueryWarehouseDataObjectData"
		err = common.PostProcessServiceError(err, "OperationsInsights", "QueryWarehouseDataObjectData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &querydataobjectresultsetrowscollection{})
	return response, err
}

// RotateOperationsInsightsWarehouseWallet Rotate the ADW wallet for Operations Insights Warehouse using which the Hub data is exposed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/RotateOperationsInsightsWarehouseWallet.go.html to see an example of how to use RotateOperationsInsightsWarehouseWallet API.
// A default retry strategy applies to this operation RotateOperationsInsightsWarehouseWallet()
func (client OperationsInsightsClient) RotateOperationsInsightsWarehouseWallet(ctx context.Context, request RotateOperationsInsightsWarehouseWalletRequest) (response RotateOperationsInsightsWarehouseWalletResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/RotateOperationsInsightsWarehouseWallet"
		err = common.PostProcessServiceError(err, "OperationsInsights", "RotateOperationsInsightsWarehouseWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAddmDbFindings Summarizes ADDM findings for the specified databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbFindings.go.html to see an example of how to use SummarizeAddmDbFindings API.
// A default retry strategy applies to this operation SummarizeAddmDbFindings()
func (client OperationsInsightsClient) SummarizeAddmDbFindings(ctx context.Context, request SummarizeAddmDbFindingsRequest) (response SummarizeAddmDbFindingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAddmDbFindings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAddmDbFindingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAddmDbFindingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAddmDbFindingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAddmDbFindingsResponse")
	}
	return
}

// summarizeAddmDbFindings implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAddmDbFindings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbFindings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAddmDbFindingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeAddmDbFindings"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAddmDbFindings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAddmDbParameterChanges Summarizes the AWR database parameter change history for the specified parameter. There will
// be one element for each time that parameter changed during the specified time period.
// This API is limited to only one parameter per request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbParameterChanges.go.html to see an example of how to use SummarizeAddmDbParameterChanges API.
// A default retry strategy applies to this operation SummarizeAddmDbParameterChanges()
func (client OperationsInsightsClient) SummarizeAddmDbParameterChanges(ctx context.Context, request SummarizeAddmDbParameterChangesRequest) (response SummarizeAddmDbParameterChangesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAddmDbParameterChanges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAddmDbParameterChangesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAddmDbParameterChangesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAddmDbParameterChangesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAddmDbParameterChangesResponse")
	}
	return
}

// summarizeAddmDbParameterChanges implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAddmDbParameterChanges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbParameterChanges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAddmDbParameterChangesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeAddmDbParameterChanges"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAddmDbParameterChanges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAddmDbParameters Summarizes database parameter history information for the specified databases. Return a list of parameters
// with information on whether the parameter values were changed or not within the specified
// time period. The response does not include the individual parameter changes within the time
// period.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbParameters.go.html to see an example of how to use SummarizeAddmDbParameters API.
// A default retry strategy applies to this operation SummarizeAddmDbParameters()
func (client OperationsInsightsClient) SummarizeAddmDbParameters(ctx context.Context, request SummarizeAddmDbParametersRequest) (response SummarizeAddmDbParametersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAddmDbParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAddmDbParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAddmDbParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAddmDbParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAddmDbParametersResponse")
	}
	return
}

// summarizeAddmDbParameters implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAddmDbParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbParameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAddmDbParametersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeAddmDbParameters"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAddmDbParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAddmDbRecommendations Summarizes ADDM recommendations for the specified databases.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbRecommendations.go.html to see an example of how to use SummarizeAddmDbRecommendations API.
// A default retry strategy applies to this operation SummarizeAddmDbRecommendations()
func (client OperationsInsightsClient) SummarizeAddmDbRecommendations(ctx context.Context, request SummarizeAddmDbRecommendationsRequest) (response SummarizeAddmDbRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAddmDbRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAddmDbRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAddmDbRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAddmDbRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAddmDbRecommendationsResponse")
	}
	return
}

// summarizeAddmDbRecommendations implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAddmDbRecommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbRecommendations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAddmDbRecommendationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeAddmDbRecommendations"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAddmDbRecommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAddmDbSchemaObjects Summarizes Schema objects for the specified databases for the specified objectIdentifiers
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbSchemaObjects.go.html to see an example of how to use SummarizeAddmDbSchemaObjects API.
// A default retry strategy applies to this operation SummarizeAddmDbSchemaObjects()
func (client OperationsInsightsClient) SummarizeAddmDbSchemaObjects(ctx context.Context, request SummarizeAddmDbSchemaObjectsRequest) (response SummarizeAddmDbSchemaObjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAddmDbSchemaObjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAddmDbSchemaObjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAddmDbSchemaObjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAddmDbSchemaObjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAddmDbSchemaObjectsResponse")
	}
	return
}

// summarizeAddmDbSchemaObjects implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAddmDbSchemaObjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbSchemaObjects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAddmDbSchemaObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeAddmDbSchemaObjects"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAddmDbSchemaObjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAddmDbSqlStatements Summarizes SQL Statements for the specified databases for the specified sqlIdentifiers
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbSqlStatements.go.html to see an example of how to use SummarizeAddmDbSqlStatements API.
// A default retry strategy applies to this operation SummarizeAddmDbSqlStatements()
func (client OperationsInsightsClient) SummarizeAddmDbSqlStatements(ctx context.Context, request SummarizeAddmDbSqlStatementsRequest) (response SummarizeAddmDbSqlStatementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAddmDbSqlStatements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAddmDbSqlStatementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAddmDbSqlStatementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAddmDbSqlStatementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAddmDbSqlStatementsResponse")
	}
	return
}

// summarizeAddmDbSqlStatements implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAddmDbSqlStatements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseInsights/addmDbSqlStatements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAddmDbSqlStatementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeAddmDbSqlStatements"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAddmDbSqlStatements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseCpuUsages Summarizes the AWR CPU resource limits and metrics for the specified database in AWR.
// Based on the time range provided as part of query param, the metrics points will be returned in the response as below.
// - if time range is <=7 days then the metrics points will be for every MINUTES
// - if time range is <=2 hours then the metrics points will be for every 10 SECONDS
// - if time range is >7 days then the metrics points will be for every HOUR.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseCpuUsages.go.html to see an example of how to use SummarizeAwrDatabaseCpuUsages API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseCpuUsages()
func (client OperationsInsightsClient) SummarizeAwrDatabaseCpuUsages(ctx context.Context, request SummarizeAwrDatabaseCpuUsagesRequest) (response SummarizeAwrDatabaseCpuUsagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseCpuUsages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseCpuUsagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseCpuUsagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseCpuUsagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseCpuUsagesResponse")
	}
	return
}

// summarizeAwrDatabaseCpuUsages implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseCpuUsages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseCpuUsages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseCpuUsagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseCpuUsages"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseCpuUsages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseMetrics Summarizes the metric samples for the specified database in the AWR. The metric samples are summarized based on the Time dimension for each metric.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseMetrics.go.html to see an example of how to use SummarizeAwrDatabaseMetrics API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseMetrics()
func (client OperationsInsightsClient) SummarizeAwrDatabaseMetrics(ctx context.Context, request SummarizeAwrDatabaseMetricsRequest) (response SummarizeAwrDatabaseMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseMetricsResponse")
	}
	return
}

// summarizeAwrDatabaseMetrics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseMetrics"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseParameterChanges Summarizes the database parameter change history for one database parameter of the specified database in AWR. One change history record contains
// the previous value, the changed value, and the corresponding time range. If the database parameter value was changed multiple times within the time range, then multiple change history records are created for the same parameter.
// Note that this API only returns information on change history details for one database parameter.
// To get a list of all the database parameters whose values were changed during a specified time range, use the following API endpoint:
// /awrHubs/{awrHubId}/awrDbParameters?awrSourceDatabaseIdentifier={awrSourceDbId}
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseParameterChanges.go.html to see an example of how to use SummarizeAwrDatabaseParameterChanges API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseParameterChanges()
func (client OperationsInsightsClient) SummarizeAwrDatabaseParameterChanges(ctx context.Context, request SummarizeAwrDatabaseParameterChangesRequest) (response SummarizeAwrDatabaseParameterChangesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseParameterChanges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseParameterChangesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseParameterChangesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseParameterChangesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseParameterChangesResponse")
	}
	return
}

// summarizeAwrDatabaseParameterChanges implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseParameterChanges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseParameterChanges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseParameterChangesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseParameterChanges"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseParameterChanges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseParameters Summarizes the database parameter history for the specified database in AWR. This includes the list of database
// parameters, with information on whether the parameter values were modified within the query time range. Note that
// each database parameter is only listed once. Depending on the optional query parameters, the returned summary gets all the database parameters, which include:
// Queryparam (valueChanged ="Y") - Each parameter whose value was changed during the time range, "isChanged : true" in response for the DB params.
// Queryparam (valueChanged ="N") - Each parameter whose value was unchanged during the time range, "isChanged : false" in response for the DB params.
// Queryparam (valueChanged ="Y"  and valueModified = "SYSTEM_MOD") - Each parameter whose value was changed at the system level during the time range, "isChanged : true" & "valueModified : SYSTEM_MOD" in response for the DB params.
// Queryparam (valueChanged ="N" and  valueDefault = "FALSE") - Each parameter whose value was unchanged during the time range, however, the value is not the default value, "isChanged : true" & "isDefault : false" in response for the DB params.
// Note that this API does not return information on the number of times each database parameter has been changed within the time range. To get the database parameter value change history for a specific parameter, use the following API endpoint:
// /awrHubs/{awrHubId}/awrDbParameterChanges?awrSourceDatabaseIdentifier={awrSourceDbId}
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseParameters.go.html to see an example of how to use SummarizeAwrDatabaseParameters API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseParameters()
func (client OperationsInsightsClient) SummarizeAwrDatabaseParameters(ctx context.Context, request SummarizeAwrDatabaseParametersRequest) (response SummarizeAwrDatabaseParametersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseParametersResponse")
	}
	return
}

// summarizeAwrDatabaseParameters implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseParameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseParametersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseParameters"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseSnapshotRanges Summarizes the AWR snapshot ranges that contain continuous snapshots, for the specified AWRHub.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseSnapshotRanges.go.html to see an example of how to use SummarizeAwrDatabaseSnapshotRanges API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseSnapshotRanges()
func (client OperationsInsightsClient) SummarizeAwrDatabaseSnapshotRanges(ctx context.Context, request SummarizeAwrDatabaseSnapshotRangesRequest) (response SummarizeAwrDatabaseSnapshotRangesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseSnapshotRanges, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseSnapshotRangesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseSnapshotRangesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseSnapshotRangesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseSnapshotRangesResponse")
	}
	return
}

// summarizeAwrDatabaseSnapshotRanges implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseSnapshotRanges(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseSnapshotRanges", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseSnapshotRangesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseSnapshotRanges"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseSnapshotRanges", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseSysstats Summarizes the AWR SYSSTAT sample data for the specified database in AWR. The statistical data is summarized based on the Time dimension for each statistic.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseSysstats.go.html to see an example of how to use SummarizeAwrDatabaseSysstats API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseSysstats()
func (client OperationsInsightsClient) SummarizeAwrDatabaseSysstats(ctx context.Context, request SummarizeAwrDatabaseSysstatsRequest) (response SummarizeAwrDatabaseSysstatsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseSysstats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseSysstatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseSysstatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseSysstatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseSysstatsResponse")
	}
	return
}

// summarizeAwrDatabaseSysstats implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseSysstats(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseSysstats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseSysstatsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseSysstats"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseSysstats", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseTopWaitEvents Summarizes the AWR top wait events.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseTopWaitEvents.go.html to see an example of how to use SummarizeAwrDatabaseTopWaitEvents API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseTopWaitEvents()
func (client OperationsInsightsClient) SummarizeAwrDatabaseTopWaitEvents(ctx context.Context, request SummarizeAwrDatabaseTopWaitEventsRequest) (response SummarizeAwrDatabaseTopWaitEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseTopWaitEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseTopWaitEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseTopWaitEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseTopWaitEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseTopWaitEventsResponse")
	}
	return
}

// summarizeAwrDatabaseTopWaitEvents implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseTopWaitEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseTopWaitEvents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseTopWaitEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseTopWaitEvents"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseTopWaitEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseWaitEventBuckets Summarizes AWR wait event data into value buckets and frequency, for the specified database in the AWR.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseWaitEventBuckets.go.html to see an example of how to use SummarizeAwrDatabaseWaitEventBuckets API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseWaitEventBuckets()
func (client OperationsInsightsClient) SummarizeAwrDatabaseWaitEventBuckets(ctx context.Context, request SummarizeAwrDatabaseWaitEventBucketsRequest) (response SummarizeAwrDatabaseWaitEventBucketsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseWaitEventBuckets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseWaitEventBucketsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseWaitEventBucketsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseWaitEventBucketsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseWaitEventBucketsResponse")
	}
	return
}

// summarizeAwrDatabaseWaitEventBuckets implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseWaitEventBuckets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseWaitEventBuckets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseWaitEventBucketsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseWaitEventBuckets"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseWaitEventBuckets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrDatabaseWaitEvents Summarizes the AWR wait event sample data for the specified database in the AWR. The event data is summarized based on the Time dimension for each event.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrDatabaseWaitEvents.go.html to see an example of how to use SummarizeAwrDatabaseWaitEvents API.
// A default retry strategy applies to this operation SummarizeAwrDatabaseWaitEvents()
func (client OperationsInsightsClient) SummarizeAwrDatabaseWaitEvents(ctx context.Context, request SummarizeAwrDatabaseWaitEventsRequest) (response SummarizeAwrDatabaseWaitEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAwrDatabaseWaitEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAwrDatabaseWaitEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAwrDatabaseWaitEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAwrDatabaseWaitEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAwrDatabaseWaitEventsResponse")
	}
	return
}

// summarizeAwrDatabaseWaitEvents implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeAwrDatabaseWaitEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/awrHubs/{awrHubId}/awrDatabaseWaitEvents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAwrDatabaseWaitEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrDatabaseWaitEvents"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrDatabaseWaitEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAwrSourcesSummaries Gets a list of summary of AWR Sources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAwrSourcesSummaries.go.html to see an example of how to use SummarizeAwrSourcesSummaries API.
// A default retry strategy applies to this operation SummarizeAwrSourcesSummaries()
func (client OperationsInsightsClient) SummarizeAwrSourcesSummaries(ctx context.Context, request SummarizeAwrSourcesSummariesRequest) (response SummarizeAwrSourcesSummariesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/SummarizeAwrSourcesSummaries"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeAwrSourcesSummaries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeConfigurationItems Gets the applicable configuration items based on the query parameters specified. Configuration items for an opsiConfigType with respect to a compartmentId can be fetched.
// Values specified in configItemField param will determine what fields for each configuration items have to be returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeConfigurationItems.go.html to see an example of how to use SummarizeConfigurationItems API.
// A default retry strategy applies to this operation SummarizeConfigurationItems()
func (client OperationsInsightsClient) SummarizeConfigurationItems(ctx context.Context, request SummarizeConfigurationItemsRequest) (response SummarizeConfigurationItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeConfigurationItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeConfigurationItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeConfigurationItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeConfigurationItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeConfigurationItemsResponse")
	}
	return
}

// summarizeConfigurationItems implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeConfigurationItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opsiConfigurations/configurationItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeConfigurationItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiConfigurations/SummarizeConfigurationItems"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeConfigurationItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &configurationitemscollection{})
	return response, err
}

// SummarizeDatabaseInsightResourceCapacityTrend Returns response with time series data (endTimestamp, capacity, baseCapacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceCapacityTrend API.
// A default retry strategy applies to this operation SummarizeDatabaseInsightResourceCapacityTrend()
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceCapacityTrend(ctx context.Context, request SummarizeDatabaseInsightResourceCapacityTrendRequest) (response SummarizeDatabaseInsightResourceCapacityTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeDatabaseInsightResourceCapacityTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeDatabaseInsightResourceCapacityTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeDatabaseInsightResourceForecastTrend Get Forecast predictions for CPU and Storage resources since a time in the past.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceForecastTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceForecastTrend API.
// A default retry strategy applies to this operation SummarizeDatabaseInsightResourceForecastTrend()
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceForecastTrend(ctx context.Context, request SummarizeDatabaseInsightResourceForecastTrendRequest) (response SummarizeDatabaseInsightResourceForecastTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeDatabaseInsightResourceForecastTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeDatabaseInsightResourceForecastTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeDatabaseInsightResourceStatistics Lists the Resource statistics (usage,capacity, usage change percent, utilization percent, base capacity, isAutoScalingEnabled)
// for each database filtered by utilization level in a compartment and in all sub-compartments if specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceStatistics.go.html to see an example of how to use SummarizeDatabaseInsightResourceStatistics API.
// A default retry strategy applies to this operation SummarizeDatabaseInsightResourceStatistics()
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceStatistics(ctx context.Context, request SummarizeDatabaseInsightResourceStatisticsRequest) (response SummarizeDatabaseInsightResourceStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeDatabaseInsightResourceStatistics"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeDatabaseInsightResourceStatistics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsage.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsage API.
// A default retry strategy applies to this operation SummarizeDatabaseInsightResourceUsage()
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUsage(ctx context.Context, request SummarizeDatabaseInsightResourceUsageRequest) (response SummarizeDatabaseInsightResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeDatabaseInsightResourceUsage"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeDatabaseInsightResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeDatabaseInsightResourceUsageTrend Returns response with time series data (endTimestamp, usage, capacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUsageTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceUsageTrend API.
// A default retry strategy applies to this operation SummarizeDatabaseInsightResourceUsageTrend()
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUsageTrend(ctx context.Context, request SummarizeDatabaseInsightResourceUsageTrendRequest) (response SummarizeDatabaseInsightResourceUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeDatabaseInsightResourceUsageTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeDatabaseInsightResourceUsageTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeDatabaseInsightResourceUtilizationInsight Gets resources with current utilization (high and low) and projected utilization (high and low) for a resource type over specified time period.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeDatabaseInsightResourceUtilizationInsight API.
// A default retry strategy applies to this operation SummarizeDatabaseInsightResourceUtilizationInsight()
func (client OperationsInsightsClient) SummarizeDatabaseInsightResourceUtilizationInsight(ctx context.Context, request SummarizeDatabaseInsightResourceUtilizationInsightRequest) (response SummarizeDatabaseInsightResourceUtilizationInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeDatabaseInsightResourceUtilizationInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeDatabaseInsightResourceUtilizationInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeDatabaseInsightTablespaceUsageTrend Returns response with usage time series data (endTimestamp, usage, capacity) with breakdown by tablespaceName for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightTablespaceUsageTrend.go.html to see an example of how to use SummarizeDatabaseInsightTablespaceUsageTrend API.
// A default retry strategy applies to this operation SummarizeDatabaseInsightTablespaceUsageTrend()
func (client OperationsInsightsClient) SummarizeDatabaseInsightTablespaceUsageTrend(ctx context.Context, request SummarizeDatabaseInsightTablespaceUsageTrendRequest) (response SummarizeDatabaseInsightTablespaceUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeDatabaseInsightTablespaceUsageTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeDatabaseInsightTablespaceUsageTrend", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeExadataInsightResourceCapacityTrend API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceCapacityTrend()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceCapacityTrend(ctx context.Context, request SummarizeExadataInsightResourceCapacityTrendRequest) (response SummarizeExadataInsightResourceCapacityTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceCapacityTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceCapacityTrend", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceCapacityTrendAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceCapacityTrendAggregated API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceCapacityTrendAggregated()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceCapacityTrendAggregated(ctx context.Context, request SummarizeExadataInsightResourceCapacityTrendAggregatedRequest) (response SummarizeExadataInsightResourceCapacityTrendAggregatedResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceCapacityTrendAggregated"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceCapacityTrendAggregated", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceForecastTrend.go.html to see an example of how to use SummarizeExadataInsightResourceForecastTrend API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceForecastTrend()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceForecastTrend(ctx context.Context, request SummarizeExadataInsightResourceForecastTrendRequest) (response SummarizeExadataInsightResourceForecastTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceForecastTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceForecastTrend", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceForecastTrendAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceForecastTrendAggregated API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceForecastTrendAggregated()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceForecastTrendAggregated(ctx context.Context, request SummarizeExadataInsightResourceForecastTrendAggregatedRequest) (response SummarizeExadataInsightResourceForecastTrendAggregatedResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceForecastTrendAggregated"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceForecastTrendAggregated", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceStatistics.go.html to see an example of how to use SummarizeExadataInsightResourceStatistics API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceStatistics()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceStatistics(ctx context.Context, request SummarizeExadataInsightResourceStatisticsRequest) (response SummarizeExadataInsightResourceStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceStatistics"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceStatistics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceUsage.go.html to see an example of how to use SummarizeExadataInsightResourceUsage API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceUsage()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceUsage(ctx context.Context, request SummarizeExadataInsightResourceUsageRequest) (response SummarizeExadataInsightResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceUsage"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceUsage", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceUsageAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceUsageAggregated API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceUsageAggregated()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceUsageAggregated(ctx context.Context, request SummarizeExadataInsightResourceUsageAggregatedRequest) (response SummarizeExadataInsightResourceUsageAggregatedResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceUsageAggregated"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceUsageAggregated", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeExadataInsightResourceUtilizationInsight Gets current utilization, projected utilization and days to reach projectedUtilization for an exadata system over specified time period. Valid values for ResourceType DATABASE are CPU,MEMORY,IO and STORAGE. Valid values for ResourceType HOST are CPU and MEMORY. Valid values for ResourceType STORAGE_SERVER are STORAGE, IOPS and THROUGHPUT.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeExadataInsightResourceUtilizationInsight API.
// A default retry strategy applies to this operation SummarizeExadataInsightResourceUtilizationInsight()
func (client OperationsInsightsClient) SummarizeExadataInsightResourceUtilizationInsight(ctx context.Context, request SummarizeExadataInsightResourceUtilizationInsightRequest) (response SummarizeExadataInsightResourceUtilizationInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataInsightResourceUtilizationInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataInsightResourceUtilizationInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeExadataMembers Lists the software and hardware inventory of the Exadata System.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataMembers.go.html to see an example of how to use SummarizeExadataMembers API.
// A default retry strategy applies to this operation SummarizeExadataMembers()
func (client OperationsInsightsClient) SummarizeExadataMembers(ctx context.Context, request SummarizeExadataMembersRequest) (response SummarizeExadataMembersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/SummarizeExadataMembers"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeExadataMembers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightDiskStatistics Returns response with disk(s) statistics for a host.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightDiskStatistics.go.html to see an example of how to use SummarizeHostInsightDiskStatistics API.
// A default retry strategy applies to this operation SummarizeHostInsightDiskStatistics()
func (client OperationsInsightsClient) SummarizeHostInsightDiskStatistics(ctx context.Context, request SummarizeHostInsightDiskStatisticsRequest) (response SummarizeHostInsightDiskStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightDiskStatistics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightDiskStatisticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightDiskStatisticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightDiskStatisticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightDiskStatisticsResponse")
	}
	return
}

// summarizeHostInsightDiskStatistics implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightDiskStatistics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/diskStatistics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightDiskStatisticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightDiskStatistics"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightDiskStatistics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightHostRecommendation Returns response with some recommendations if apply for a host.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightHostRecommendation.go.html to see an example of how to use SummarizeHostInsightHostRecommendation API.
// A default retry strategy applies to this operation SummarizeHostInsightHostRecommendation()
func (client OperationsInsightsClient) SummarizeHostInsightHostRecommendation(ctx context.Context, request SummarizeHostInsightHostRecommendationRequest) (response SummarizeHostInsightHostRecommendationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightHostRecommendation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightHostRecommendationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightHostRecommendationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightHostRecommendationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightHostRecommendationResponse")
	}
	return
}

// summarizeHostInsightHostRecommendation implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightHostRecommendation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/hostRecommendation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightHostRecommendationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightHostRecommendation"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightHostRecommendation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightNetworkUsageTrend Returns response with usage time series data with breakdown by network interface for the time period specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightNetworkUsageTrend.go.html to see an example of how to use SummarizeHostInsightNetworkUsageTrend API.
// A default retry strategy applies to this operation SummarizeHostInsightNetworkUsageTrend()
func (client OperationsInsightsClient) SummarizeHostInsightNetworkUsageTrend(ctx context.Context, request SummarizeHostInsightNetworkUsageTrendRequest) (response SummarizeHostInsightNetworkUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightNetworkUsageTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightNetworkUsageTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightNetworkUsageTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightNetworkUsageTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightNetworkUsageTrendResponse")
	}
	return
}

// summarizeHostInsightNetworkUsageTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightNetworkUsageTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/networkUsageTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightNetworkUsageTrendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightNetworkUsageTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightNetworkUsageTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightResourceCapacityTrend Returns response with time series data (endTimestamp, capacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeHostInsightResourceCapacityTrend API.
// A default retry strategy applies to this operation SummarizeHostInsightResourceCapacityTrend()
func (client OperationsInsightsClient) SummarizeHostInsightResourceCapacityTrend(ctx context.Context, request SummarizeHostInsightResourceCapacityTrendRequest) (response SummarizeHostInsightResourceCapacityTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightResourceCapacityTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightResourceCapacityTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightResourceForecastTrend Get Forecast predictions for CPU or memory resources since a time in the past.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceForecastTrend.go.html to see an example of how to use SummarizeHostInsightResourceForecastTrend API.
// A default retry strategy applies to this operation SummarizeHostInsightResourceForecastTrend()
func (client OperationsInsightsClient) SummarizeHostInsightResourceForecastTrend(ctx context.Context, request SummarizeHostInsightResourceForecastTrendRequest) (response SummarizeHostInsightResourceForecastTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightResourceForecastTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightResourceForecastTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightResourceStatistics Lists the resource statistics (usage, capacity, usage change percent, utilization percent, load) for each host filtered
// by utilization level in a compartment and in all sub-compartments if specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceStatistics.go.html to see an example of how to use SummarizeHostInsightResourceStatistics API.
// A default retry strategy applies to this operation SummarizeHostInsightResourceStatistics()
func (client OperationsInsightsClient) SummarizeHostInsightResourceStatistics(ctx context.Context, request SummarizeHostInsightResourceStatisticsRequest) (response SummarizeHostInsightResourceStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightResourceStatistics"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightResourceStatistics", apiReferenceLink)
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
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUsage.go.html to see an example of how to use SummarizeHostInsightResourceUsage API.
// A default retry strategy applies to this operation SummarizeHostInsightResourceUsage()
func (client OperationsInsightsClient) SummarizeHostInsightResourceUsage(ctx context.Context, request SummarizeHostInsightResourceUsageRequest) (response SummarizeHostInsightResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightResourceUsage"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightResourceUsageTrend Returns response with time series data (endTimestamp, usage, capacity) for the time period specified.
// The maximum time range for analysis is 2 years, hence this is intentionally not paginated.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUsageTrend.go.html to see an example of how to use SummarizeHostInsightResourceUsageTrend API.
// A default retry strategy applies to this operation SummarizeHostInsightResourceUsageTrend()
func (client OperationsInsightsClient) SummarizeHostInsightResourceUsageTrend(ctx context.Context, request SummarizeHostInsightResourceUsageTrendRequest) (response SummarizeHostInsightResourceUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightResourceUsageTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightResourceUsageTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightResourceUtilizationInsight Gets resources with current utilization (high and low) and projected utilization (high and low) for a resource type over specified time period.
// If compartmentIdInSubtree is specified, aggregates resources in a compartment and in all sub-compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeHostInsightResourceUtilizationInsight API.
// A default retry strategy applies to this operation SummarizeHostInsightResourceUtilizationInsight()
func (client OperationsInsightsClient) SummarizeHostInsightResourceUtilizationInsight(ctx context.Context, request SummarizeHostInsightResourceUtilizationInsightRequest) (response SummarizeHostInsightResourceUtilizationInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightResourceUtilizationInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightResourceUtilizationInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightStorageUsageTrend Returns response with usage time series data with breakdown by filesystem for the time period specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightStorageUsageTrend.go.html to see an example of how to use SummarizeHostInsightStorageUsageTrend API.
// A default retry strategy applies to this operation SummarizeHostInsightStorageUsageTrend()
func (client OperationsInsightsClient) SummarizeHostInsightStorageUsageTrend(ctx context.Context, request SummarizeHostInsightStorageUsageTrendRequest) (response SummarizeHostInsightStorageUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightStorageUsageTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightStorageUsageTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightStorageUsageTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightStorageUsageTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightStorageUsageTrendResponse")
	}
	return
}

// summarizeHostInsightStorageUsageTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightStorageUsageTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/storageUsageTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightStorageUsageTrendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightStorageUsageTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightStorageUsageTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightTopProcessesUsage Returns response with aggregated data (timestamp, usageData) for top processes on a specific date.
// Data is aggregated for the time specified and processes are sorted descendent by the process metric specified (CPU, MEMORY, VIRTUAL_MEMORY).
// hostInsightId, processMetric must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightTopProcessesUsage.go.html to see an example of how to use SummarizeHostInsightTopProcessesUsage API.
// A default retry strategy applies to this operation SummarizeHostInsightTopProcessesUsage()
func (client OperationsInsightsClient) SummarizeHostInsightTopProcessesUsage(ctx context.Context, request SummarizeHostInsightTopProcessesUsageRequest) (response SummarizeHostInsightTopProcessesUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightTopProcessesUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightTopProcessesUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightTopProcessesUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightTopProcessesUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightTopProcessesUsageResponse")
	}
	return
}

// summarizeHostInsightTopProcessesUsage implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightTopProcessesUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/topProcessesUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightTopProcessesUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightTopProcessesUsage"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightTopProcessesUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeHostInsightTopProcessesUsageTrend Returns response with aggregated time series data (timeIntervalstart, timeIntervalEnd, commandArgs, usageData) for top processes.
// Data is aggregated for the time period specified and proceses are sorted descendent by the proces metric specified (CPU, MEMORY, VIRTUAL_MEMORY).
// HostInsight Id and Process metric must be specified
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightTopProcessesUsageTrend.go.html to see an example of how to use SummarizeHostInsightTopProcessesUsageTrend API.
// A default retry strategy applies to this operation SummarizeHostInsightTopProcessesUsageTrend()
func (client OperationsInsightsClient) SummarizeHostInsightTopProcessesUsageTrend(ctx context.Context, request SummarizeHostInsightTopProcessesUsageTrendRequest) (response SummarizeHostInsightTopProcessesUsageTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeHostInsightTopProcessesUsageTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeHostInsightTopProcessesUsageTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeHostInsightTopProcessesUsageTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeHostInsightTopProcessesUsageTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeHostInsightTopProcessesUsageTrendResponse")
	}
	return
}

// summarizeHostInsightTopProcessesUsageTrend implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) summarizeHostInsightTopProcessesUsageTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/hostInsights/topProcessesUsageTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeHostInsightTopProcessesUsageTrendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/SummarizeHostInsightTopProcessesUsageTrend"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeHostInsightTopProcessesUsageTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeOperationsInsightsWarehouseResourceUsage Gets the details of resources used by an Operations Insights Warehouse.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeOperationsInsightsWarehouseResourceUsage.go.html to see an example of how to use SummarizeOperationsInsightsWarehouseResourceUsage API.
// A default retry strategy applies to this operation SummarizeOperationsInsightsWarehouseResourceUsage()
func (client OperationsInsightsClient) SummarizeOperationsInsightsWarehouseResourceUsage(ctx context.Context, request SummarizeOperationsInsightsWarehouseResourceUsageRequest) (response SummarizeOperationsInsightsWarehouseResourceUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/SummarizeOperationsInsightsWarehouseResourceUsage"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeOperationsInsightsWarehouseResourceUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlInsights Query SQL Warehouse to get the performance insights for SQLs taking greater than X% database time for a given
// time period across the given databases or database types in a compartment and in all sub-compartments if specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlInsights.go.html to see an example of how to use SummarizeSqlInsights API.
// A default retry strategy applies to this operation SummarizeSqlInsights()
func (client OperationsInsightsClient) SummarizeSqlInsights(ctx context.Context, request SummarizeSqlInsightsRequest) (response SummarizeSqlInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeSqlInsights"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeSqlInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlPlanInsights Query SQL Warehouse to get the performance insights on the execution plans for a given SQL for a given time period.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlPlanInsights.go.html to see an example of how to use SummarizeSqlPlanInsights API.
// A default retry strategy applies to this operation SummarizeSqlPlanInsights()
func (client OperationsInsightsClient) SummarizeSqlPlanInsights(ctx context.Context, request SummarizeSqlPlanInsightsRequest) (response SummarizeSqlPlanInsightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeSqlPlanInsights"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeSqlPlanInsights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlResponseTimeDistributions Query SQL Warehouse to summarize the response time distribution of query executions for a given SQL for a given time period.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlResponseTimeDistributions.go.html to see an example of how to use SummarizeSqlResponseTimeDistributions API.
// A default retry strategy applies to this operation SummarizeSqlResponseTimeDistributions()
func (client OperationsInsightsClient) SummarizeSqlResponseTimeDistributions(ctx context.Context, request SummarizeSqlResponseTimeDistributionsRequest) (response SummarizeSqlResponseTimeDistributionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeSqlResponseTimeDistributions"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeSqlResponseTimeDistributions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlStatistics Query SQL Warehouse to get the performance statistics for SQLs taking greater than X% database time for a given
// time period across the given databases or database types in a compartment and in all sub-compartments if specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatistics.go.html to see an example of how to use SummarizeSqlStatistics API.
// A default retry strategy applies to this operation SummarizeSqlStatistics()
func (client OperationsInsightsClient) SummarizeSqlStatistics(ctx context.Context, request SummarizeSqlStatisticsRequest) (response SummarizeSqlStatisticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeSqlStatistics"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeSqlStatistics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlStatisticsTimeSeries Query SQL Warehouse to get the performance statistics time series for a given SQL across given databases for a
// given time period in a compartment and in all sub-compartments if specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatisticsTimeSeries.go.html to see an example of how to use SummarizeSqlStatisticsTimeSeries API.
// A default retry strategy applies to this operation SummarizeSqlStatisticsTimeSeries()
func (client OperationsInsightsClient) SummarizeSqlStatisticsTimeSeries(ctx context.Context, request SummarizeSqlStatisticsTimeSeriesRequest) (response SummarizeSqlStatisticsTimeSeriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeSqlStatisticsTimeSeries"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeSqlStatisticsTimeSeries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeSqlStatisticsTimeSeriesByPlan Query SQL Warehouse to get the performance statistics time series for a given SQL by execution plans for a given time period.
// Either databaseId or id must be specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeSqlStatisticsTimeSeriesByPlan.go.html to see an example of how to use SummarizeSqlStatisticsTimeSeriesByPlan API.
// A default retry strategy applies to this operation SummarizeSqlStatisticsTimeSeriesByPlan()
func (client OperationsInsightsClient) SummarizeSqlStatisticsTimeSeriesByPlan(ctx context.Context, request SummarizeSqlStatisticsTimeSeriesByPlanRequest) (response SummarizeSqlStatisticsTimeSeriesByPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/SummarizeSqlStatisticsTimeSeriesByPlan"
		err = common.PostProcessServiceError(err, "OperationsInsights", "SummarizeSqlStatisticsTimeSeriesByPlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestMacsManagedCloudDatabaseInsightConnection Test the connection details of a Cloud MACS-managed database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/TestMacsManagedCloudDatabaseInsightConnection.go.html to see an example of how to use TestMacsManagedCloudDatabaseInsightConnection API.
// A default retry strategy applies to this operation TestMacsManagedCloudDatabaseInsightConnection()
func (client OperationsInsightsClient) TestMacsManagedCloudDatabaseInsightConnection(ctx context.Context, request TestMacsManagedCloudDatabaseInsightConnectionRequest) (response TestMacsManagedCloudDatabaseInsightConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.testMacsManagedCloudDatabaseInsightConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestMacsManagedCloudDatabaseInsightConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestMacsManagedCloudDatabaseInsightConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestMacsManagedCloudDatabaseInsightConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestMacsManagedCloudDatabaseInsightConnectionResponse")
	}
	return
}

// testMacsManagedCloudDatabaseInsightConnection implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) testMacsManagedCloudDatabaseInsightConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseInsights/actions/testMacsManagedCloudDatabaseInsightConnectionDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestMacsManagedCloudDatabaseInsightConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/TestMacsManagedCloudDatabaseInsightConnection"
		err = common.PostProcessServiceError(err, "OperationsInsights", "TestMacsManagedCloudDatabaseInsightConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAwrHub Updates the configuration of a hub .
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateAwrHub.go.html to see an example of how to use UpdateAwrHub API.
// A default retry strategy applies to this operation UpdateAwrHub()
func (client OperationsInsightsClient) UpdateAwrHub(ctx context.Context, request UpdateAwrHubRequest) (response UpdateAwrHubResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubs/UpdateAwrHub"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateAwrHub", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAwrHubSource Update Awr Hub Source object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateAwrHubSource.go.html to see an example of how to use UpdateAwrHubSource API.
// A default retry strategy applies to this operation UpdateAwrHubSource()
func (client OperationsInsightsClient) UpdateAwrHubSource(ctx context.Context, request UpdateAwrHubSourceRequest) (response UpdateAwrHubSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAwrHubSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAwrHubSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAwrHubSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAwrHubSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAwrHubSourceResponse")
	}
	return
}

// updateAwrHubSource implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateAwrHubSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/awrHubSources/{awrHubSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAwrHubSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/AwrHubSources/UpdateAwrHubSource"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateAwrHubSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseInsight Updates configuration of a database insight.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateDatabaseInsight.go.html to see an example of how to use UpdateDatabaseInsight API.
// A default retry strategy applies to this operation UpdateDatabaseInsight()
func (client OperationsInsightsClient) UpdateDatabaseInsight(ctx context.Context, request UpdateDatabaseInsightRequest) (response UpdateDatabaseInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/DatabaseInsights/UpdateDatabaseInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateDatabaseInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEnterpriseManagerBridge Updates configuration of an Operations Insights Enterprise Manager bridge.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateEnterpriseManagerBridge.go.html to see an example of how to use UpdateEnterpriseManagerBridge API.
// A default retry strategy applies to this operation UpdateEnterpriseManagerBridge()
func (client OperationsInsightsClient) UpdateEnterpriseManagerBridge(ctx context.Context, request UpdateEnterpriseManagerBridgeRequest) (response UpdateEnterpriseManagerBridgeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/EnterpriseManagerBridges/UpdateEnterpriseManagerBridge"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateEnterpriseManagerBridge", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateExadataInsight Updates configuration of an Exadata insight.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateExadataInsight.go.html to see an example of how to use UpdateExadataInsight API.
// A default retry strategy applies to this operation UpdateExadataInsight()
func (client OperationsInsightsClient) UpdateExadataInsight(ctx context.Context, request UpdateExadataInsightRequest) (response UpdateExadataInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/ExadataInsights/UpdateExadataInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateExadataInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateHostInsight Updates configuration of a host insight.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateHostInsight.go.html to see an example of how to use UpdateHostInsight API.
// A default retry strategy applies to this operation UpdateHostInsight()
func (client OperationsInsightsClient) UpdateHostInsight(ctx context.Context, request UpdateHostInsightRequest) (response UpdateHostInsightResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/HostInsights/UpdateHostInsight"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateHostInsight", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNewsReport Updates the  configuration of a news report.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateNewsReport.go.html to see an example of how to use UpdateNewsReport API.
// A default retry strategy applies to this operation UpdateNewsReport()
func (client OperationsInsightsClient) UpdateNewsReport(ctx context.Context, request UpdateNewsReportRequest) (response UpdateNewsReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNewsReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNewsReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNewsReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNewsReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNewsReportResponse")
	}
	return
}

// updateNewsReport implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateNewsReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/newsReports/{newsReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNewsReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/NewsReports/UpdateNewsReport"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateNewsReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOperationsInsightsPrivateEndpoint Updates one or more attributes of the specified private endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateOperationsInsightsPrivateEndpoint.go.html to see an example of how to use UpdateOperationsInsightsPrivateEndpoint API.
// A default retry strategy applies to this operation UpdateOperationsInsightsPrivateEndpoint()
func (client OperationsInsightsClient) UpdateOperationsInsightsPrivateEndpoint(ctx context.Context, request UpdateOperationsInsightsPrivateEndpointRequest) (response UpdateOperationsInsightsPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOperationsInsightsPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOperationsInsightsPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOperationsInsightsPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOperationsInsightsPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOperationsInsightsPrivateEndpointResponse")
	}
	return
}

// updateOperationsInsightsPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateOperationsInsightsPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/operationsInsightsPrivateEndpoints/{operationsInsightsPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOperationsInsightsPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsPrivateEndpoint/UpdateOperationsInsightsPrivateEndpoint"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateOperationsInsightsPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOperationsInsightsWarehouse Updates the configuration of an Ops Insights Warehouse.
// There is only expected to be 1 warehouse per tenant. The warehouse is expected to be in the root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateOperationsInsightsWarehouse.go.html to see an example of how to use UpdateOperationsInsightsWarehouse API.
// A default retry strategy applies to this operation UpdateOperationsInsightsWarehouse()
func (client OperationsInsightsClient) UpdateOperationsInsightsWarehouse(ctx context.Context, request UpdateOperationsInsightsWarehouseRequest) (response UpdateOperationsInsightsWarehouseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouses/UpdateOperationsInsightsWarehouse"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateOperationsInsightsWarehouse", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOperationsInsightsWarehouseUser Updates the configuration of an Operations Insights Warehouse User.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateOperationsInsightsWarehouseUser.go.html to see an example of how to use UpdateOperationsInsightsWarehouseUser API.
// A default retry strategy applies to this operation UpdateOperationsInsightsWarehouseUser()
func (client OperationsInsightsClient) UpdateOperationsInsightsWarehouseUser(ctx context.Context, request UpdateOperationsInsightsWarehouseUserRequest) (response UpdateOperationsInsightsWarehouseUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OperationsInsightsWarehouseUsers/UpdateOperationsInsightsWarehouseUser"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateOperationsInsightsWarehouseUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOpsiConfiguration Updates an OPSI configuration resource with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/UpdateOpsiConfiguration.go.html to see an example of how to use UpdateOpsiConfiguration API.
// A default retry strategy applies to this operation UpdateOpsiConfiguration()
func (client OperationsInsightsClient) UpdateOpsiConfiguration(ctx context.Context, request UpdateOpsiConfigurationRequest) (response UpdateOpsiConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOpsiConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOpsiConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOpsiConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOpsiConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOpsiConfigurationResponse")
	}
	return
}

// updateOpsiConfiguration implements the OCIOperation interface (enables retrying operations)
func (client OperationsInsightsClient) updateOpsiConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/opsiConfigurations/{opsiConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOpsiConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/operations-insights/20200630/OpsiConfigurations/UpdateOpsiConfiguration"
		err = common.PostProcessServiceError(err, "OperationsInsights", "UpdateOpsiConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
