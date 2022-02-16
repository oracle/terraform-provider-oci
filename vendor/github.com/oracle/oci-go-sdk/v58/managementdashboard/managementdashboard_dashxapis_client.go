// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// API for the Management Dashboard micro-service. Use this API for dashboard and saved search metadata preservation and to perform  tasks such as creating a dashboard, creating a saved search, and obtaining a list of dashboards and saved searches in a compartment.
//
//

package managementdashboard

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//DashxApisClient a client for DashxApis
type DashxApisClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDashxApisClientWithConfigurationProvider Creates a new default DashxApis client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDashxApisClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DashxApisClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDashxApisClientFromBaseClient(baseClient, provider)
}

// NewDashxApisClientWithOboToken Creates a new default DashxApis client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewDashxApisClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DashxApisClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDashxApisClientFromBaseClient(baseClient, configProvider)
}

func newDashxApisClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DashxApisClient, err error) {
	// DashxApis service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DashxApisClient{BaseClient: baseClient}
	client.BasePath = "20200901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DashxApisClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("managementdashboard", "https://managementdashboard.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DashxApisClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DashxApisClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeManagementDashboardsCompartment Moves the dashboard from the existing compartment to a new compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ChangeManagementDashboardsCompartment.go.html to see an example of how to use ChangeManagementDashboardsCompartment API.
func (client DashxApisClient) ChangeManagementDashboardsCompartment(ctx context.Context, request ChangeManagementDashboardsCompartmentRequest) (response ChangeManagementDashboardsCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeManagementDashboardsCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagementDashboardsCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagementDashboardsCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagementDashboardsCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagementDashboardsCompartmentResponse")
	}
	return
}

// changeManagementDashboardsCompartment implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) changeManagementDashboardsCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementDashboards/{managementDashboardId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeManagementDashboardsCompartmentResponse
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

// ChangeManagementSavedSearchesCompartment Moves the saved search from the existing compartment to a new compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ChangeManagementSavedSearchesCompartment.go.html to see an example of how to use ChangeManagementSavedSearchesCompartment API.
func (client DashxApisClient) ChangeManagementSavedSearchesCompartment(ctx context.Context, request ChangeManagementSavedSearchesCompartmentRequest) (response ChangeManagementSavedSearchesCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeManagementSavedSearchesCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagementSavedSearchesCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagementSavedSearchesCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagementSavedSearchesCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagementSavedSearchesCompartmentResponse")
	}
	return
}

// changeManagementSavedSearchesCompartment implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) changeManagementSavedSearchesCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementSavedSearches/{managementSavedSearchId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeManagementSavedSearchesCompartmentResponse
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

// CreateManagementDashboard Creates a new dashboard.  Limit for number of saved searches in a dashboard is 20. Here's an example of how you can use CLI to create a dashboard. For information on the details that must be passed to CREATE, you can use the GET API to obtain the Create.json file:
// oci management-dashboard dashboard get --management-dashboard-id  "ocid1.managementdashboard.oc1..dashboardId1" --query data > Create.json.
// You can then modify the Create.json file by removing the"id" attribute and making other required changes, and use the oci management-dashboard dashboard create command.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/CreateManagementDashboard.go.html to see an example of how to use CreateManagementDashboard API.
func (client DashxApisClient) CreateManagementDashboard(ctx context.Context, request CreateManagementDashboardRequest) (response CreateManagementDashboardResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagementDashboardResponse")
	}
	return
}

// createManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) createManagementDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementDashboards", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateManagementDashboardResponse
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

// CreateManagementSavedSearch Creates a new saved search. Here's an example of how you can use CLI to create a saved search. For information on the details that must be passed to CREATE, you can use the GET API to obtain the Create.json file:
// oci management-dashboard saved-search get --management-saved-search-id ocid1.managementsavedsearch.oc1..savedsearchId1 --query data > Create.json.
// You can then modify the Create.json file by removing the "id" attribute and making other required changes, and use the oci management-dashboard saved-search create command.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/CreateManagementSavedSearch.go.html to see an example of how to use CreateManagementSavedSearch API.
func (client DashxApisClient) CreateManagementSavedSearch(ctx context.Context, request CreateManagementSavedSearchRequest) (response CreateManagementSavedSearchResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagementSavedSearchResponse")
	}
	return
}

// createManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) createManagementSavedSearch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementSavedSearches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateManagementSavedSearchResponse
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

// DeleteManagementDashboard Deletes a Dashboard by ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/DeleteManagementDashboard.go.html to see an example of how to use DeleteManagementDashboard API.
func (client DashxApisClient) DeleteManagementDashboard(ctx context.Context, request DeleteManagementDashboardRequest) (response DeleteManagementDashboardResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementDashboardResponse")
	}
	return
}

// deleteManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) deleteManagementDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementDashboards/{managementDashboardId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteManagementDashboardResponse
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

// DeleteManagementSavedSearch Deletes a saved search by ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/DeleteManagementSavedSearch.go.html to see an example of how to use DeleteManagementSavedSearch API.
func (client DashxApisClient) DeleteManagementSavedSearch(ctx context.Context, request DeleteManagementSavedSearchRequest) (response DeleteManagementSavedSearchResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementSavedSearchResponse")
	}
	return
}

// deleteManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) deleteManagementSavedSearch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementSavedSearches/{managementSavedSearchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteManagementSavedSearchResponse
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

// ExportDashboard Exports an array of dashboards and their saved searches. Export is designed to work with importDashboard. Here's an example of how you can use CLI to export a dashboard. $oci management-dashboard dashboard export --query data --export-dashboard-id "{\"dashboardIds\":[\"ocid1.managementdashboard.oc1..dashboardId1\"]}"  > dashboards.json
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ExportDashboard.go.html to see an example of how to use ExportDashboard API.
func (client DashxApisClient) ExportDashboard(ctx context.Context, request ExportDashboardRequest) (response ExportDashboardResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportDashboardResponse")
	}
	return
}

// exportDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) exportDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementDashboards/actions/exportDashboard/{exportDashboardId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportDashboardResponse
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

// GetManagementDashboard Gets a dashboard and its saved searches by ID.  Deleted or unauthorized saved searches are marked by tile's state property.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/GetManagementDashboard.go.html to see an example of how to use GetManagementDashboard API.
func (client DashxApisClient) GetManagementDashboard(ctx context.Context, request GetManagementDashboardRequest) (response GetManagementDashboardResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementDashboardResponse")
	}
	return
}

// getManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) getManagementDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementDashboards/{managementDashboardId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagementDashboardResponse
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

// GetManagementSavedSearch Gets a saved search by ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/GetManagementSavedSearch.go.html to see an example of how to use GetManagementSavedSearch API.
func (client DashxApisClient) GetManagementSavedSearch(ctx context.Context, request GetManagementSavedSearchRequest) (response GetManagementSavedSearchResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementSavedSearchResponse")
	}
	return
}

// getManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) getManagementSavedSearch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementSavedSearches/{managementSavedSearchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagementSavedSearchResponse
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

// ImportDashboard Imports an array of dashboards and their saved searches. Here's an example of how you can use CLI to import a dashboard. For information on the details that must be passed to IMPORT, you can use the EXPORT API to obtain the Import.json file:
// oci management-dashboard dashboard export --query data --export-dashboard-id "{\"dashboardIds\":[\"ocid1.managementdashboard.oc1..dashboardId1\"]}"  > Import.json.
// Note that import API updates the resource if it already exist, and creates a new resource if it does not exist. To import to a different compartment, edit and change the compartmentId to the desired compartment OCID.
// Here is an example of how you can use CLI to do import:
// oci management-dashboard dashboard import --from-json file://Import.json
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ImportDashboard.go.html to see an example of how to use ImportDashboard API.
func (client DashxApisClient) ImportDashboard(ctx context.Context, request ImportDashboardRequest) (response ImportDashboardResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportDashboardResponse")
	}
	return
}

// importDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) importDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementDashboards/actions/importDashboard", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportDashboardResponse
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

// ListManagementDashboards Gets the list of dashboards in a compartment with pagination.  Returned properties are the summary.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ListManagementDashboards.go.html to see an example of how to use ListManagementDashboards API.
func (client DashxApisClient) ListManagementDashboards(ctx context.Context, request ListManagementDashboardsRequest) (response ListManagementDashboardsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementDashboards, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementDashboardsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementDashboardsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementDashboardsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementDashboardsResponse")
	}
	return
}

// listManagementDashboards implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) listManagementDashboards(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementDashboards", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagementDashboardsResponse
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

// ListManagementSavedSearches Gets the list of saved searches in a compartment with pagination.  Returned properties are the summary.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ListManagementSavedSearches.go.html to see an example of how to use ListManagementSavedSearches API.
func (client DashxApisClient) ListManagementSavedSearches(ctx context.Context, request ListManagementSavedSearchesRequest) (response ListManagementSavedSearchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementSavedSearches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementSavedSearchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementSavedSearchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementSavedSearchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementSavedSearchesResponse")
	}
	return
}

// listManagementSavedSearches implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) listManagementSavedSearches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementSavedSearches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagementSavedSearchesResponse
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

// UpdateManagementDashboard Updates an existing dashboard identified by ID path parameter.  CompartmentId can be modified only by the changeCompartment API. Limit for number of saved searches in a dashboard is 20.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/UpdateManagementDashboard.go.html to see an example of how to use UpdateManagementDashboard API.
func (client DashxApisClient) UpdateManagementDashboard(ctx context.Context, request UpdateManagementDashboardRequest) (response UpdateManagementDashboardResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateManagementDashboard, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementDashboardResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementDashboardResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementDashboardResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementDashboardResponse")
	}
	return
}

// updateManagementDashboard implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) updateManagementDashboard(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementDashboards/{managementDashboardId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagementDashboardResponse
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

// UpdateManagementSavedSearch Updates an existing saved search identified by ID path parameter.  CompartmentId can be modified only by the changeCompartment API.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/UpdateManagementSavedSearch.go.html to see an example of how to use UpdateManagementSavedSearch API.
func (client DashxApisClient) UpdateManagementSavedSearch(ctx context.Context, request UpdateManagementSavedSearchRequest) (response UpdateManagementSavedSearchResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateManagementSavedSearch, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementSavedSearchResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementSavedSearchResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementSavedSearchResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementSavedSearchResponse")
	}
	return
}

// updateManagementSavedSearch implements the OCIOperation interface (enables retrying operations)
func (client DashxApisClient) updateManagementSavedSearch(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementSavedSearches/{managementSavedSearchId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagementSavedSearchResponse
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
