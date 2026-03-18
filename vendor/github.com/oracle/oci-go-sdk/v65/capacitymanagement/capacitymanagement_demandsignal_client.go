// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DemandSignalClient a client for DemandSignal
type DemandSignalClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDemandSignalClientWithConfigurationProvider Creates a new default DemandSignal client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDemandSignalClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DemandSignalClient, err error) {
	if enabled := common.CheckForEnabledServices("capacitymanagement"); !enabled {
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
	return newDemandSignalClientFromBaseClient(baseClient, provider)
}

// NewDemandSignalClientWithOboToken Creates a new default DemandSignal client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDemandSignalClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DemandSignalClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDemandSignalClientFromBaseClient(baseClient, configProvider)
}

func newDemandSignalClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DemandSignalClient, err error) {
	// DemandSignal service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DemandSignal"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DemandSignalClient{BaseClient: baseClient}
	client.BasePath = "20231107"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DemandSignalClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("capacitymanagement", "https://control-center-cp.{region}.oci.{secondLevelDomain}", "control-center-cp")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DemandSignalClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DemandSignalClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkCreateOccmDemandSignalItem This API will help in bulk creation of demand signal items. This API is atomic i.e either all the demand signal item resources will be created or none will be created.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/BulkCreateOccmDemandSignalItem.go.html to see an example of how to use BulkCreateOccmDemandSignalItem API.
// A default retry strategy applies to this operation BulkCreateOccmDemandSignalItem()
func (client DemandSignalClient) BulkCreateOccmDemandSignalItem(ctx context.Context, request BulkCreateOccmDemandSignalItemRequest) (response BulkCreateOccmDemandSignalItemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkCreateOccmDemandSignalItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkCreateOccmDemandSignalItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkCreateOccmDemandSignalItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkCreateOccmDemandSignalItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkCreateOccmDemandSignalItemResponse")
	}
	return
}

// bulkCreateOccmDemandSignalItem implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) bulkCreateOccmDemandSignalItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occmDemandSignalItems/actions/bulkCreateDemandSignalItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkCreateOccmDemandSignalItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalItem/BulkCreateOccmDemandSignalItem"
		err = common.PostProcessServiceError(err, "DemandSignal", "BulkCreateOccmDemandSignalItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOccmDemandSignal This is a post API to create occm demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/CreateOccmDemandSignal.go.html to see an example of how to use CreateOccmDemandSignal API.
// A default retry strategy applies to this operation CreateOccmDemandSignal()
func (client DemandSignalClient) CreateOccmDemandSignal(ctx context.Context, request CreateOccmDemandSignalRequest) (response CreateOccmDemandSignalResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccmDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccmDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccmDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccmDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccmDemandSignalResponse")
	}
	return
}

// createOccmDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) createOccmDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occmDemandSignals", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccmDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignal/CreateOccmDemandSignal"
		err = common.PostProcessServiceError(err, "DemandSignal", "CreateOccmDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOccmDemandSignalItem This API will create a demand signal item representing a resource request. This needs to be grouped under a demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/CreateOccmDemandSignalItem.go.html to see an example of how to use CreateOccmDemandSignalItem API.
// A default retry strategy applies to this operation CreateOccmDemandSignalItem()
func (client DemandSignalClient) CreateOccmDemandSignalItem(ctx context.Context, request CreateOccmDemandSignalItemRequest) (response CreateOccmDemandSignalItemResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccmDemandSignalItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccmDemandSignalItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccmDemandSignalItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccmDemandSignalItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccmDemandSignalItemResponse")
	}
	return
}

// createOccmDemandSignalItem implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) createOccmDemandSignalItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occmDemandSignalItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccmDemandSignalItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalItem/CreateOccmDemandSignalItem"
		err = common.PostProcessServiceError(err, "DemandSignal", "CreateOccmDemandSignalItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccmDemandSignal This is a DELETE API which deletes a demand signal with the provided demand signal ocid.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/DeleteOccmDemandSignal.go.html to see an example of how to use DeleteOccmDemandSignal API.
// A default retry strategy applies to this operation DeleteOccmDemandSignal()
func (client DemandSignalClient) DeleteOccmDemandSignal(ctx context.Context, request DeleteOccmDemandSignalRequest) (response DeleteOccmDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccmDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccmDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccmDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccmDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccmDemandSignalResponse")
	}
	return
}

// deleteOccmDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) deleteOccmDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occmDemandSignals/{occmDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccmDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignal/DeleteOccmDemandSignal"
		err = common.PostProcessServiceError(err, "DemandSignal", "DeleteOccmDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccmDemandSignalItem This is a DELETE API which deletes a demand signal item with the provided demand signal item ocid.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/DeleteOccmDemandSignalItem.go.html to see an example of how to use DeleteOccmDemandSignalItem API.
// A default retry strategy applies to this operation DeleteOccmDemandSignalItem()
func (client DemandSignalClient) DeleteOccmDemandSignalItem(ctx context.Context, request DeleteOccmDemandSignalItemRequest) (response DeleteOccmDemandSignalItemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccmDemandSignalItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccmDemandSignalItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccmDemandSignalItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccmDemandSignalItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccmDemandSignalItemResponse")
	}
	return
}

// deleteOccmDemandSignalItem implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) deleteOccmDemandSignalItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occmDemandSignalItems/{occmDemandSignalItemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccmDemandSignalItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalItem/DeleteOccmDemandSignalItem"
		err = common.PostProcessServiceError(err, "DemandSignal", "DeleteOccmDemandSignalItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccmDemandSignal This is a GET API which gets the detailed information about a specific demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetOccmDemandSignal.go.html to see an example of how to use GetOccmDemandSignal API.
// A default retry strategy applies to this operation GetOccmDemandSignal()
func (client DemandSignalClient) GetOccmDemandSignal(ctx context.Context, request GetOccmDemandSignalRequest) (response GetOccmDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOccmDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccmDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccmDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccmDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccmDemandSignalResponse")
	}
	return
}

// getOccmDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) getOccmDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occmDemandSignals/{occmDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccmDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignal/GetOccmDemandSignal"
		err = common.PostProcessServiceError(err, "DemandSignal", "GetOccmDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccmDemandSignalItem This is a GET API to get the details of a demand signal item resource representing the details of the resource demanded by you.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetOccmDemandSignalItem.go.html to see an example of how to use GetOccmDemandSignalItem API.
// A default retry strategy applies to this operation GetOccmDemandSignalItem()
func (client DemandSignalClient) GetOccmDemandSignalItem(ctx context.Context, request GetOccmDemandSignalItemRequest) (response GetOccmDemandSignalItemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOccmDemandSignalItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccmDemandSignalItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccmDemandSignalItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccmDemandSignalItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccmDemandSignalItemResponse")
	}
	return
}

// getOccmDemandSignalItem implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) getOccmDemandSignalItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occmDemandSignalItems/{occmDemandSignalItemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccmDemandSignalItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalItem/GetOccmDemandSignalItem"
		err = common.PostProcessServiceError(err, "DemandSignal", "GetOccmDemandSignalItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccmDemandSignalCatalogResources This API will list all the  resources across all demand signal catalogs for a given namespace and customer group containing the caller compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignalCatalogResources.go.html to see an example of how to use ListOccmDemandSignalCatalogResources API.
// A default retry strategy applies to this operation ListOccmDemandSignalCatalogResources()
func (client DemandSignalClient) ListOccmDemandSignalCatalogResources(ctx context.Context, request ListOccmDemandSignalCatalogResourcesRequest) (response ListOccmDemandSignalCatalogResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccmDemandSignalCatalogResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccmDemandSignalCatalogResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccmDemandSignalCatalogResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccmDemandSignalCatalogResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccmDemandSignalCatalogResourcesResponse")
	}
	return
}

// listOccmDemandSignalCatalogResources implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) listOccmDemandSignalCatalogResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occmDemandSignalCatalogResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccmDemandSignalCatalogResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalCatalogResource/ListOccmDemandSignalCatalogResources"
		err = common.PostProcessServiceError(err, "DemandSignal", "ListOccmDemandSignalCatalogResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccmDemandSignalDeliveries This GET call is used to list all demand signals delivery resources within the compartment passed as a query param.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignalDeliveries.go.html to see an example of how to use ListOccmDemandSignalDeliveries API.
// A default retry strategy applies to this operation ListOccmDemandSignalDeliveries()
func (client DemandSignalClient) ListOccmDemandSignalDeliveries(ctx context.Context, request ListOccmDemandSignalDeliveriesRequest) (response ListOccmDemandSignalDeliveriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccmDemandSignalDeliveries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccmDemandSignalDeliveriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccmDemandSignalDeliveriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccmDemandSignalDeliveriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccmDemandSignalDeliveriesResponse")
	}
	return
}

// listOccmDemandSignalDeliveries implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) listOccmDemandSignalDeliveries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occmDemandSignalDeliveries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccmDemandSignalDeliveriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalDeliveryCollection/ListOccmDemandSignalDeliveries"
		err = common.PostProcessServiceError(err, "DemandSignal", "ListOccmDemandSignalDeliveries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccmDemandSignalItems This API will list the detailed information about the resources demanded as part of the demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignalItems.go.html to see an example of how to use ListOccmDemandSignalItems API.
// A default retry strategy applies to this operation ListOccmDemandSignalItems()
func (client DemandSignalClient) ListOccmDemandSignalItems(ctx context.Context, request ListOccmDemandSignalItemsRequest) (response ListOccmDemandSignalItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccmDemandSignalItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccmDemandSignalItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccmDemandSignalItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccmDemandSignalItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccmDemandSignalItemsResponse")
	}
	return
}

// listOccmDemandSignalItems implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) listOccmDemandSignalItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occmDemandSignalItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccmDemandSignalItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalItem/ListOccmDemandSignalItems"
		err = common.PostProcessServiceError(err, "DemandSignal", "ListOccmDemandSignalItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccmDemandSignals This GET call is used to list all demand signals within the compartment passed as a query parameter.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignals.go.html to see an example of how to use ListOccmDemandSignals API.
// A default retry strategy applies to this operation ListOccmDemandSignals()
func (client DemandSignalClient) ListOccmDemandSignals(ctx context.Context, request ListOccmDemandSignalsRequest) (response ListOccmDemandSignalsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccmDemandSignals, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccmDemandSignalsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccmDemandSignalsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccmDemandSignalsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccmDemandSignalsResponse")
	}
	return
}

// listOccmDemandSignals implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) listOccmDemandSignals(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occmDemandSignals", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccmDemandSignalsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignal/ListOccmDemandSignals"
		err = common.PostProcessServiceError(err, "DemandSignal", "ListOccmDemandSignals", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccmDemandSignal This is a PUT API which shall be used to update the metadata of the demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateOccmDemandSignal.go.html to see an example of how to use UpdateOccmDemandSignal API.
// A default retry strategy applies to this operation UpdateOccmDemandSignal()
func (client DemandSignalClient) UpdateOccmDemandSignal(ctx context.Context, request UpdateOccmDemandSignalRequest) (response UpdateOccmDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccmDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccmDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccmDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccmDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccmDemandSignalResponse")
	}
	return
}

// updateOccmDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) updateOccmDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occmDemandSignals/{occmDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccmDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignal/UpdateOccmDemandSignal"
		err = common.PostProcessServiceError(err, "DemandSignal", "UpdateOccmDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccmDemandSignalItem This is a PUT API which can be used to update the demand signal item resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateOccmDemandSignalItem.go.html to see an example of how to use UpdateOccmDemandSignalItem API.
// A default retry strategy applies to this operation UpdateOccmDemandSignalItem()
func (client DemandSignalClient) UpdateOccmDemandSignalItem(ctx context.Context, request UpdateOccmDemandSignalItemRequest) (response UpdateOccmDemandSignalItemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccmDemandSignalItem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccmDemandSignalItemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccmDemandSignalItemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccmDemandSignalItemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccmDemandSignalItemResponse")
	}
	return
}

// updateOccmDemandSignalItem implements the OCIOperation interface (enables retrying operations)
func (client DemandSignalClient) updateOccmDemandSignalItem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occmDemandSignalItems/{occmDemandSignalItemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccmDemandSignalItemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalItem/UpdateOccmDemandSignalItem"
		err = common.PostProcessServiceError(err, "DemandSignal", "UpdateOccmDemandSignalItem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
