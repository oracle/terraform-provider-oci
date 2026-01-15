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

// InternalDemandSignalClient a client for InternalDemandSignal
type InternalDemandSignalClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewInternalDemandSignalClientWithConfigurationProvider Creates a new default InternalDemandSignal client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewInternalDemandSignalClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client InternalDemandSignalClient, err error) {
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
	return newInternalDemandSignalClientFromBaseClient(baseClient, provider)
}

// NewInternalDemandSignalClientWithOboToken Creates a new default InternalDemandSignal client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewInternalDemandSignalClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client InternalDemandSignalClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newInternalDemandSignalClientFromBaseClient(baseClient, configProvider)
}

func newInternalDemandSignalClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client InternalDemandSignalClient, err error) {
	// InternalDemandSignal service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("InternalDemandSignal"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = InternalDemandSignalClient{BaseClient: baseClient}
	client.BasePath = "20231107"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *InternalDemandSignalClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("capacitymanagement", "https://control-center-cp.{region}.oci.{secondLevelDomain}", "control-center-cp")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *InternalDemandSignalClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *InternalDemandSignalClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateInternalOccmDemandSignalDelivery This is a post API which is used to create a demand signal delivery resource.
// operationId: CreateInternalOccmDemandSignalDelivery
// summary: A post call to create a demand signal delivery.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/CreateInternalOccmDemandSignalDelivery.go.html to see an example of how to use CreateInternalOccmDemandSignalDelivery API.
// A default retry strategy applies to this operation CreateInternalOccmDemandSignalDelivery()
func (client InternalDemandSignalClient) CreateInternalOccmDemandSignalDelivery(ctx context.Context, request CreateInternalOccmDemandSignalDeliveryRequest) (response CreateInternalOccmDemandSignalDeliveryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createInternalOccmDemandSignalDelivery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInternalOccmDemandSignalDeliveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInternalOccmDemandSignalDeliveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInternalOccmDemandSignalDeliveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInternalOccmDemandSignalDeliveryResponse")
	}
	return
}

// createInternalOccmDemandSignalDelivery implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) createInternalOccmDemandSignalDelivery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/internal/occmDemandSignalDeliveries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateInternalOccmDemandSignalDeliveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignalDelivery/CreateInternalOccmDemandSignalDelivery"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "CreateInternalOccmDemandSignalDelivery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteInternalOccmDemandSignalDelivery This is an internal DELETE API which is used to delete a demand signal delivery resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/DeleteInternalOccmDemandSignalDelivery.go.html to see an example of how to use DeleteInternalOccmDemandSignalDelivery API.
// A default retry strategy applies to this operation DeleteInternalOccmDemandSignalDelivery()
func (client InternalDemandSignalClient) DeleteInternalOccmDemandSignalDelivery(ctx context.Context, request DeleteInternalOccmDemandSignalDeliveryRequest) (response DeleteInternalOccmDemandSignalDeliveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInternalOccmDemandSignalDelivery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteInternalOccmDemandSignalDeliveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteInternalOccmDemandSignalDeliveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInternalOccmDemandSignalDeliveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInternalOccmDemandSignalDeliveryResponse")
	}
	return
}

// deleteInternalOccmDemandSignalDelivery implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) deleteInternalOccmDemandSignalDelivery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/internal/occmDemandSignalDeliveries/{occmDemandSignalDeliveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteInternalOccmDemandSignalDeliveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignalDelivery/DeleteInternalOccmDemandSignalDelivery"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "DeleteInternalOccmDemandSignalDelivery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInternalOccmDemandSignal This is an internal GET API which gets the detailed information about a specific demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetInternalOccmDemandSignal.go.html to see an example of how to use GetInternalOccmDemandSignal API.
// A default retry strategy applies to this operation GetInternalOccmDemandSignal()
func (client InternalDemandSignalClient) GetInternalOccmDemandSignal(ctx context.Context, request GetInternalOccmDemandSignalRequest) (response GetInternalOccmDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInternalOccmDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInternalOccmDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInternalOccmDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInternalOccmDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInternalOccmDemandSignalResponse")
	}
	return
}

// getInternalOccmDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) getInternalOccmDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignals/{occmDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInternalOccmDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignal/GetInternalOccmDemandSignal"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "GetInternalOccmDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInternalOccmDemandSignalCatalog This API helps in getting the details about a specific occm demand signal catalog.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetInternalOccmDemandSignalCatalog.go.html to see an example of how to use GetInternalOccmDemandSignalCatalog API.
// A default retry strategy applies to this operation GetInternalOccmDemandSignalCatalog()
func (client InternalDemandSignalClient) GetInternalOccmDemandSignalCatalog(ctx context.Context, request GetInternalOccmDemandSignalCatalogRequest) (response GetInternalOccmDemandSignalCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInternalOccmDemandSignalCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInternalOccmDemandSignalCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInternalOccmDemandSignalCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInternalOccmDemandSignalCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInternalOccmDemandSignalCatalogResponse")
	}
	return
}

// getInternalOccmDemandSignalCatalog implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) getInternalOccmDemandSignalCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignalCatalog/{occmDemandSignalCatalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInternalOccmDemandSignalCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalCatalogResource/GetInternalOccmDemandSignalCatalog"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "GetInternalOccmDemandSignalCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInternalOccmDemandSignalDelivery This is an internal GET API to get the details of a demand signal delivery resource corresponding to a demand signal item.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetInternalOccmDemandSignalDelivery.go.html to see an example of how to use GetInternalOccmDemandSignalDelivery API.
// A default retry strategy applies to this operation GetInternalOccmDemandSignalDelivery()
func (client InternalDemandSignalClient) GetInternalOccmDemandSignalDelivery(ctx context.Context, request GetInternalOccmDemandSignalDeliveryRequest) (response GetInternalOccmDemandSignalDeliveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInternalOccmDemandSignalDelivery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInternalOccmDemandSignalDeliveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInternalOccmDemandSignalDeliveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInternalOccmDemandSignalDeliveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInternalOccmDemandSignalDeliveryResponse")
	}
	return
}

// getInternalOccmDemandSignalDelivery implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) getInternalOccmDemandSignalDelivery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignalDeliveries/{occmDemandSignalDeliveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInternalOccmDemandSignalDeliveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignalDelivery/GetInternalOccmDemandSignalDelivery"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "GetInternalOccmDemandSignalDelivery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalOccmDemandSignalCatalogResources This API will list all the  resources across all demand signal catalogs for a given namespace and customer group.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalCatalogResources.go.html to see an example of how to use ListInternalOccmDemandSignalCatalogResources API.
// A default retry strategy applies to this operation ListInternalOccmDemandSignalCatalogResources()
func (client InternalDemandSignalClient) ListInternalOccmDemandSignalCatalogResources(ctx context.Context, request ListInternalOccmDemandSignalCatalogResourcesRequest) (response ListInternalOccmDemandSignalCatalogResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalOccmDemandSignalCatalogResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalOccmDemandSignalCatalogResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalOccmDemandSignalCatalogResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalOccmDemandSignalCatalogResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalOccmDemandSignalCatalogResourcesResponse")
	}
	return
}

// listInternalOccmDemandSignalCatalogResources implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) listInternalOccmDemandSignalCatalogResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignalCatalogResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalOccmDemandSignalCatalogResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignalCatalogResource/ListInternalOccmDemandSignalCatalogResources"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "ListInternalOccmDemandSignalCatalogResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalOccmDemandSignalCatalogs This API will list demand signal catalogs for a given customer group.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalCatalogs.go.html to see an example of how to use ListInternalOccmDemandSignalCatalogs API.
// A default retry strategy applies to this operation ListInternalOccmDemandSignalCatalogs()
func (client InternalDemandSignalClient) ListInternalOccmDemandSignalCatalogs(ctx context.Context, request ListInternalOccmDemandSignalCatalogsRequest) (response ListInternalOccmDemandSignalCatalogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalOccmDemandSignalCatalogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalOccmDemandSignalCatalogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalOccmDemandSignalCatalogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalOccmDemandSignalCatalogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalOccmDemandSignalCatalogsResponse")
	}
	return
}

// listInternalOccmDemandSignalCatalogs implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) listInternalOccmDemandSignalCatalogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignalCatalog", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalOccmDemandSignalCatalogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccmDemandSignalCatalog/ListInternalOccmDemandSignalCatalogs"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "ListInternalOccmDemandSignalCatalogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalOccmDemandSignalDeliveries This GET call is used to list all demand signal delivery resources within the customer group passed as a query parameter.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalDeliveries.go.html to see an example of how to use ListInternalOccmDemandSignalDeliveries API.
// A default retry strategy applies to this operation ListInternalOccmDemandSignalDeliveries()
func (client InternalDemandSignalClient) ListInternalOccmDemandSignalDeliveries(ctx context.Context, request ListInternalOccmDemandSignalDeliveriesRequest) (response ListInternalOccmDemandSignalDeliveriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalOccmDemandSignalDeliveries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalOccmDemandSignalDeliveriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalOccmDemandSignalDeliveriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalOccmDemandSignalDeliveriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalOccmDemandSignalDeliveriesResponse")
	}
	return
}

// listInternalOccmDemandSignalDeliveries implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) listInternalOccmDemandSignalDeliveries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignalDeliveries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalOccmDemandSignalDeliveriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignalDelivery/ListInternalOccmDemandSignalDeliveries"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "ListInternalOccmDemandSignalDeliveries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalOccmDemandSignalItems This internal API will list the detailed information about the resources demanded as part of the demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalItems.go.html to see an example of how to use ListInternalOccmDemandSignalItems API.
// A default retry strategy applies to this operation ListInternalOccmDemandSignalItems()
func (client InternalDemandSignalClient) ListInternalOccmDemandSignalItems(ctx context.Context, request ListInternalOccmDemandSignalItemsRequest) (response ListInternalOccmDemandSignalItemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalOccmDemandSignalItems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalOccmDemandSignalItemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalOccmDemandSignalItemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalOccmDemandSignalItemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalOccmDemandSignalItemsResponse")
	}
	return
}

// listInternalOccmDemandSignalItems implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) listInternalOccmDemandSignalItems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignalItems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalOccmDemandSignalItemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignalItemCollection/ListInternalOccmDemandSignalItems"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "ListInternalOccmDemandSignalItems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalOccmDemandSignals This is an internal GET call is used to list all demand signals within the compartment passed as a query parameter.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignals.go.html to see an example of how to use ListInternalOccmDemandSignals API.
// A default retry strategy applies to this operation ListInternalOccmDemandSignals()
func (client InternalDemandSignalClient) ListInternalOccmDemandSignals(ctx context.Context, request ListInternalOccmDemandSignalsRequest) (response ListInternalOccmDemandSignalsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalOccmDemandSignals, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalOccmDemandSignalsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalOccmDemandSignalsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalOccmDemandSignalsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalOccmDemandSignalsResponse")
	}
	return
}

// listInternalOccmDemandSignals implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) listInternalOccmDemandSignals(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occmDemandSignals", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalOccmDemandSignalsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignal/ListInternalOccmDemandSignals"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "ListInternalOccmDemandSignals", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateInternalOccmDemandSignal This is a internal PUT API which shall be used to update the metadata of the demand signal.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateInternalOccmDemandSignal.go.html to see an example of how to use UpdateInternalOccmDemandSignal API.
// A default retry strategy applies to this operation UpdateInternalOccmDemandSignal()
func (client InternalDemandSignalClient) UpdateInternalOccmDemandSignal(ctx context.Context, request UpdateInternalOccmDemandSignalRequest) (response UpdateInternalOccmDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateInternalOccmDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInternalOccmDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInternalOccmDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInternalOccmDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInternalOccmDemandSignalResponse")
	}
	return
}

// updateInternalOccmDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) updateInternalOccmDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/internal/occmDemandSignals/{occmDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInternalOccmDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignal/UpdateInternalOccmDemandSignal"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "UpdateInternalOccmDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateInternalOccmDemandSignalDelivery This is an internal PUT API which is used to update the demand signal delivery resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateInternalOccmDemandSignalDelivery.go.html to see an example of how to use UpdateInternalOccmDemandSignalDelivery API.
// A default retry strategy applies to this operation UpdateInternalOccmDemandSignalDelivery()
func (client InternalDemandSignalClient) UpdateInternalOccmDemandSignalDelivery(ctx context.Context, request UpdateInternalOccmDemandSignalDeliveryRequest) (response UpdateInternalOccmDemandSignalDeliveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateInternalOccmDemandSignalDelivery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInternalOccmDemandSignalDeliveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInternalOccmDemandSignalDeliveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInternalOccmDemandSignalDeliveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInternalOccmDemandSignalDeliveryResponse")
	}
	return
}

// updateInternalOccmDemandSignalDelivery implements the OCIOperation interface (enables retrying operations)
func (client InternalDemandSignalClient) updateInternalOccmDemandSignalDelivery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/internal/occmDemandSignalDeliveries/{occmDemandSignalDeliveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInternalOccmDemandSignalDeliveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/InternalOccmDemandSignalDelivery/UpdateInternalOccmDemandSignalDelivery"
		err = common.PostProcessServiceError(err, "InternalDemandSignal", "UpdateInternalOccmDemandSignalDelivery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
