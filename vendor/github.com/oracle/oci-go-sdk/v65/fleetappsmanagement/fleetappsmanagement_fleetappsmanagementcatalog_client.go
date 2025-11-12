// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FleetAppsManagementCatalogClient a client for FleetAppsManagementCatalog
type FleetAppsManagementCatalogClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFleetAppsManagementCatalogClientWithConfigurationProvider Creates a new default FleetAppsManagementCatalog client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFleetAppsManagementCatalogClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FleetAppsManagementCatalogClient, err error) {
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
	return newFleetAppsManagementCatalogClientFromBaseClient(baseClient, provider)
}

// NewFleetAppsManagementCatalogClientWithOboToken Creates a new default FleetAppsManagementCatalog client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFleetAppsManagementCatalogClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FleetAppsManagementCatalogClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFleetAppsManagementCatalogClientFromBaseClient(baseClient, configProvider)
}

func newFleetAppsManagementCatalogClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FleetAppsManagementCatalogClient, err error) {
	// FleetAppsManagementCatalog service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FleetAppsManagementCatalog"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FleetAppsManagementCatalogClient{BaseClient: baseClient}
	client.BasePath = "20250228"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FleetAppsManagementCatalogClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fleetappsmanagement", "https://fams.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FleetAppsManagementCatalogClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FleetAppsManagementCatalogClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeCatalogItemCompartment Moves a CatalogItem into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ChangeCatalogItemCompartment.go.html to see an example of how to use ChangeCatalogItemCompartment API.
// A default retry strategy applies to this operation ChangeCatalogItemCompartment()
func (client FleetAppsManagementCatalogClient) ChangeCatalogItemCompartment(ctx context.Context, request ChangeCatalogItemCompartmentRequest) (response ChangeCatalogItemCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCatalogItemCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCatalogItemCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCatalogItemCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCatalogItemCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCatalogItemCompartmentResponse")
	}
	return
}

// changeCatalogItemCompartment implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) changeCatalogItemCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogItems/{catalogItemId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCatalogItemCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItem/ChangeCatalogItemCompartment"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "ChangeCatalogItemCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CloneCatalogItem Clones a CatalogItem into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CloneCatalogItem.go.html to see an example of how to use CloneCatalogItem API.
// A default retry strategy applies to this operation CloneCatalogItem()
func (client FleetAppsManagementCatalogClient) CloneCatalogItem(ctx context.Context, request CloneCatalogItemRequest) (response CloneCatalogItemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cloneCatalogItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CloneCatalogItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CloneCatalogItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CloneCatalogItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CloneCatalogItemResponse")
	}
	return
}

// cloneCatalogItem implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) cloneCatalogItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogItems/{catalogItemId}/actions/cloneCatalogItem", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CloneCatalogItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItem/CloneCatalogItem"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "CloneCatalogItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureCatalogItem Configures a CatalogItem. Creating new Catalog Item.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ConfigureCatalogItem.go.html to see an example of how to use ConfigureCatalogItem API.
// A default retry strategy applies to this operation ConfigureCatalogItem()
func (client FleetAppsManagementCatalogClient) ConfigureCatalogItem(ctx context.Context, request ConfigureCatalogItemRequest) (response ConfigureCatalogItemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureCatalogItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureCatalogItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureCatalogItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureCatalogItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureCatalogItemResponse")
	}
	return
}

// configureCatalogItem implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) configureCatalogItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogItems/{catalogItemId}/actions/configure", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureCatalogItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItem/ConfigureCatalogItem"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "ConfigureCatalogItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCatalogItem Creates a CatalogItem.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateCatalogItem.go.html to see an example of how to use CreateCatalogItem API.
// A default retry strategy applies to this operation CreateCatalogItem()
func (client FleetAppsManagementCatalogClient) CreateCatalogItem(ctx context.Context, request CreateCatalogItemRequest) (response CreateCatalogItemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCatalogItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCatalogItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCatalogItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCatalogItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCatalogItemResponse")
	}
	return
}

// createCatalogItem implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) createCatalogItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCatalogItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItem/CreateCatalogItem"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "CreateCatalogItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCatalogItem Deletes a CatalogItem.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteCatalogItem.go.html to see an example of how to use DeleteCatalogItem API.
// A default retry strategy applies to this operation DeleteCatalogItem()
func (client FleetAppsManagementCatalogClient) DeleteCatalogItem(ctx context.Context, request DeleteCatalogItemRequest) (response DeleteCatalogItemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCatalogItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCatalogItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCatalogItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCatalogItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCatalogItemResponse")
	}
	return
}

// deleteCatalogItem implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) deleteCatalogItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogItems/{catalogItemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCatalogItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItem/DeleteCatalogItem"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "DeleteCatalogItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCatalogItem Gets information about a CatalogItem.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetCatalogItem.go.html to see an example of how to use GetCatalogItem API.
// A default retry strategy applies to this operation GetCatalogItem()
func (client FleetAppsManagementCatalogClient) GetCatalogItem(ctx context.Context, request GetCatalogItemRequest) (response GetCatalogItemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalogItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCatalogItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCatalogItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogItemResponse")
	}
	return
}

// getCatalogItem implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) getCatalogItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogItems/{catalogItemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCatalogItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItem/GetCatalogItem"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "GetCatalogItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCatalogItemVariablesDefinition Gets information about a CatalogItem Variables.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetCatalogItemVariablesDefinition.go.html to see an example of how to use GetCatalogItemVariablesDefinition API.
// A default retry strategy applies to this operation GetCatalogItemVariablesDefinition()
func (client FleetAppsManagementCatalogClient) GetCatalogItemVariablesDefinition(ctx context.Context, request GetCatalogItemVariablesDefinitionRequest) (response GetCatalogItemVariablesDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalogItemVariablesDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCatalogItemVariablesDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCatalogItemVariablesDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogItemVariablesDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogItemVariablesDefinitionResponse")
	}
	return
}

// getCatalogItemVariablesDefinition implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) getCatalogItemVariablesDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogItems/{catalogItemId}/variablesDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCatalogItemVariablesDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItemVariablesDefinition/GetCatalogItemVariablesDefinition"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "GetCatalogItemVariablesDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCatalogItems Gets a list of Catalog Items in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListCatalogItems.go.html to see an example of how to use ListCatalogItems API.
// A default retry strategy applies to this operation ListCatalogItems()
func (client FleetAppsManagementCatalogClient) ListCatalogItems(ctx context.Context, request ListCatalogItemsRequest) (response ListCatalogItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCatalogItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCatalogItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogItemsResponse")
	}
	return
}

// listCatalogItems implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) listCatalogItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCatalogItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItemCollection/ListCatalogItems"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "ListCatalogItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCatalogItem Updates a CatalogItem.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateCatalogItem.go.html to see an example of how to use UpdateCatalogItem API.
// A default retry strategy applies to this operation UpdateCatalogItem()
func (client FleetAppsManagementCatalogClient) UpdateCatalogItem(ctx context.Context, request UpdateCatalogItemRequest) (response UpdateCatalogItemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCatalogItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCatalogItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCatalogItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCatalogItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCatalogItemResponse")
	}
	return
}

// updateCatalogItem implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementCatalogClient) updateCatalogItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogItems/{catalogItemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCatalogItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20250228/CatalogItem/UpdateCatalogItem"
		err = common.PostProcessServiceError(err, "FleetAppsManagementCatalog", "UpdateCatalogItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
