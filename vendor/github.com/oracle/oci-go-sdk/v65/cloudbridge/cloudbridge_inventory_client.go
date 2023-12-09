// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// InventoryClient a client for Inventory
type InventoryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewInventoryClientWithConfigurationProvider Creates a new default Inventory client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewInventoryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client InventoryClient, err error) {
	if enabled := common.CheckForEnabledServices("cloudbridge"); !enabled {
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
	return newInventoryClientFromBaseClient(baseClient, provider)
}

// NewInventoryClientWithOboToken Creates a new default Inventory client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewInventoryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client InventoryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newInventoryClientFromBaseClient(baseClient, configProvider)
}

func newInventoryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client InventoryClient, err error) {
	// Inventory service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Inventory"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = InventoryClient{BaseClient: baseClient}
	client.BasePath = "20220509"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *InventoryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cloudbridge", "https://cloudbridge.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *InventoryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *InventoryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AnalyzeAssets Returns an aggregation of assets. Aggregation groups are sorted by groupBy property.
// Default sort order is ascending, but can be overridden by the sortOrder parameter.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/AnalyzeAssets.go.html to see an example of how to use AnalyzeAssets API.
// A default retry strategy applies to this operation AnalyzeAssets()
func (client InventoryClient) AnalyzeAssets(ctx context.Context, request AnalyzeAssetsRequest) (response AnalyzeAssetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.analyzeAssets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AnalyzeAssetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AnalyzeAssetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AnalyzeAssetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AnalyzeAssetsResponse")
	}
	return
}

// analyzeAssets implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) analyzeAssets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assetAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AnalyzeAssetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetAggregation/AnalyzeAssets"
		err = common.PostProcessServiceError(err, "Inventory", "AnalyzeAssets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAssetCompartment Moves an asset resource from one compartment to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ChangeAssetCompartment.go.html to see an example of how to use ChangeAssetCompartment API.
// A default retry strategy applies to this operation ChangeAssetCompartment()
func (client InventoryClient) ChangeAssetCompartment(ctx context.Context, request ChangeAssetCompartmentRequest) (response ChangeAssetCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAssetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAssetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAssetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAssetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAssetCompartmentResponse")
	}
	return
}

// changeAssetCompartment implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) changeAssetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assets/{assetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAssetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Asset/ChangeAssetCompartment"
		err = common.PostProcessServiceError(err, "Inventory", "ChangeAssetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAssetTags Change an asset's tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ChangeAssetTags.go.html to see an example of how to use ChangeAssetTags API.
// A default retry strategy applies to this operation ChangeAssetTags()
func (client InventoryClient) ChangeAssetTags(ctx context.Context, request ChangeAssetTagsRequest) (response ChangeAssetTagsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAssetTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAssetTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAssetTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAssetTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAssetTagsResponse")
	}
	return
}

// changeAssetTags implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) changeAssetTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assets/{assetId}/actions/changeTags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAssetTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Asset/ChangeAssetTags"
		err = common.PostProcessServiceError(err, "Inventory", "ChangeAssetTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &asset{})
	return response, err
}

// CreateAsset Creates an asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/CreateAsset.go.html to see an example of how to use CreateAsset API.
// A default retry strategy applies to this operation CreateAsset()
func (client InventoryClient) CreateAsset(ctx context.Context, request CreateAssetRequest) (response CreateAssetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAssetResponse")
	}
	return
}

// createAsset implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) createAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Asset/CreateAsset"
		err = common.PostProcessServiceError(err, "Inventory", "CreateAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &asset{})
	return response, err
}

// CreateInventory Creates an inventory.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/CreateInventory.go.html to see an example of how to use CreateInventory API.
// A default retry strategy applies to this operation CreateInventory()
func (client InventoryClient) CreateInventory(ctx context.Context, request CreateInventoryRequest) (response CreateInventoryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInventoryResponse")
	}
	return
}

// createInventory implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) createInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/inventories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Inventory/CreateInventory"
		err = common.PostProcessServiceError(err, "Inventory", "CreateInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAsset Deletes an asset resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/DeleteAsset.go.html to see an example of how to use DeleteAsset API.
// A default retry strategy applies to this operation DeleteAsset()
func (client InventoryClient) DeleteAsset(ctx context.Context, request DeleteAssetRequest) (response DeleteAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAssetResponse")
	}
	return
}

// deleteAsset implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) deleteAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/assets/{assetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Asset/DeleteAsset"
		err = common.PostProcessServiceError(err, "Inventory", "DeleteAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteInventory Deletes an inventory resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/DeleteInventory.go.html to see an example of how to use DeleteInventory API.
// A default retry strategy applies to this operation DeleteInventory()
func (client InventoryClient) DeleteInventory(ctx context.Context, request DeleteInventoryRequest) (response DeleteInventoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInventoryResponse")
	}
	return
}

// deleteInventory implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) deleteInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/inventories/{inventoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Inventory/DeleteInventory"
		err = common.PostProcessServiceError(err, "Inventory", "DeleteInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAsset Gets an asset by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetAsset.go.html to see an example of how to use GetAsset API.
// A default retry strategy applies to this operation GetAsset()
func (client InventoryClient) GetAsset(ctx context.Context, request GetAssetRequest) (response GetAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssetResponse")
	}
	return
}

// getAsset implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) getAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assets/{assetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Asset/GetAsset"
		err = common.PostProcessServiceError(err, "Inventory", "GetAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &asset{})
	return response, err
}

// GetInventory Gets an inventory by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetInventory.go.html to see an example of how to use GetInventory API.
// A default retry strategy applies to this operation GetInventory()
func (client InventoryClient) GetInventory(ctx context.Context, request GetInventoryRequest) (response GetInventoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInventoryResponse")
	}
	return
}

// getInventory implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) getInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/inventories/{inventoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Inventory/GetInventory"
		err = common.PostProcessServiceError(err, "Inventory", "GetInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportInventory Import resources in inventory.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ImportInventory.go.html to see an example of how to use ImportInventory API.
// A default retry strategy applies to this operation ImportInventory()
func (client InventoryClient) ImportInventory(ctx context.Context, request ImportInventoryRequest) (response ImportInventoryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportInventoryResponse")
	}
	return
}

// importInventory implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) importInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/inventories/{inventoryId}/actions/import", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Inventory/ImportInventory"
		err = common.PostProcessServiceError(err, "Inventory", "ImportInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssets Returns a list of assets.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAssets.go.html to see an example of how to use ListAssets API.
// A default retry strategy applies to this operation ListAssets()
func (client InventoryClient) ListAssets(ctx context.Context, request ListAssetsRequest) (response ListAssetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssetsResponse")
	}
	return
}

// listAssets implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) listAssets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AssetCollection/ListAssets"
		err = common.PostProcessServiceError(err, "Inventory", "ListAssets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHistoricalMetrics List asset historical metrics.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListHistoricalMetrics.go.html to see an example of how to use ListHistoricalMetrics API.
// A default retry strategy applies to this operation ListHistoricalMetrics()
func (client InventoryClient) ListHistoricalMetrics(ctx context.Context, request ListHistoricalMetricsRequest) (response ListHistoricalMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHistoricalMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHistoricalMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHistoricalMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHistoricalMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHistoricalMetricsResponse")
	}
	return
}

// listHistoricalMetrics implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) listHistoricalMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/assets/{assetId}/historicalMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHistoricalMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/HistoricalMetric/ListHistoricalMetrics"
		err = common.PostProcessServiceError(err, "Inventory", "ListHistoricalMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInventories Returns a list of inventories.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListInventories.go.html to see an example of how to use ListInventories API.
// A default retry strategy applies to this operation ListInventories()
func (client InventoryClient) ListInventories(ctx context.Context, request ListInventoriesRequest) (response ListInventoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInventories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInventoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInventoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInventoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInventoriesResponse")
	}
	return
}

// listInventories implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) listInventories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/inventories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInventoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Inventory/ListInventories"
		err = common.PostProcessServiceError(err, "Inventory", "ListInventories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SubmitHistoricalMetrics Creates or updates all metrics related to the asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/SubmitHistoricalMetrics.go.html to see an example of how to use SubmitHistoricalMetrics API.
// A default retry strategy applies to this operation SubmitHistoricalMetrics()
func (client InventoryClient) SubmitHistoricalMetrics(ctx context.Context, request SubmitHistoricalMetricsRequest) (response SubmitHistoricalMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.submitHistoricalMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SubmitHistoricalMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SubmitHistoricalMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SubmitHistoricalMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SubmitHistoricalMetricsResponse")
	}
	return
}

// submitHistoricalMetrics implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) submitHistoricalMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/assets/{assetId}/actions/submitHistoricalMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SubmitHistoricalMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/HistoricalMetric/SubmitHistoricalMetrics"
		err = common.PostProcessServiceError(err, "Inventory", "SubmitHistoricalMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAsset Updates the asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdateAsset.go.html to see an example of how to use UpdateAsset API.
// A default retry strategy applies to this operation UpdateAsset()
func (client InventoryClient) UpdateAsset(ctx context.Context, request UpdateAssetRequest) (response UpdateAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAssetResponse")
	}
	return
}

// updateAsset implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) updateAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/assets/{assetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Asset/UpdateAsset"
		err = common.PostProcessServiceError(err, "Inventory", "UpdateAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &asset{})
	return response, err
}

// UpdateInventory Updates an inventory.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdateInventory.go.html to see an example of how to use UpdateInventory API.
// A default retry strategy applies to this operation UpdateInventory()
func (client InventoryClient) UpdateInventory(ctx context.Context, request UpdateInventoryRequest) (response UpdateInventoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInventoryResponse")
	}
	return
}

// updateInventory implements the OCIOperation interface (enables retrying operations)
func (client InventoryClient) updateInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/inventories/{inventoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Inventory/UpdateInventory"
		err = common.PostProcessServiceError(err, "Inventory", "UpdateInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
