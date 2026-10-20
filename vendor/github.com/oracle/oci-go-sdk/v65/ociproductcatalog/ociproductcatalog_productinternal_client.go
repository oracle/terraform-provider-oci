// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Product Catalog API
//
// Apis to manage the products used by the OCI service teams
//

package ociproductcatalog

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ProductInternalClient a client for ProductInternal
type ProductInternalClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewProductInternalClientWithConfigurationProvider Creates a new default ProductInternal client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewProductInternalClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ProductInternalClient, err error) {
	if enabled := common.CheckForEnabledServices("ociproductcatalog"); !enabled {
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
	return newProductInternalClientFromBaseClient(baseClient, provider)
}

// NewProductInternalClientWithOboToken Creates a new default ProductInternal client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewProductInternalClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ProductInternalClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newProductInternalClientFromBaseClient(baseClient, configProvider)
}

func newProductInternalClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ProductInternalClient, err error) {
	// ProductInternal service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ProductInternal"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ProductInternalClient{BaseClient: baseClient}
	client.BasePath = "20250610"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ProductInternalClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ociproductcatalog", "https://cp.product-catalog.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ProductInternalClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ProductInternalClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeInternalProductCompartment Change product compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/ChangeInternalProductCompartment.go.html to see an example of how to use ChangeInternalProductCompartment API.
// A default retry strategy applies to this operation ChangeInternalProductCompartment()
func (client ProductInternalClient) ChangeInternalProductCompartment(ctx context.Context, request ChangeInternalProductCompartmentRequest) (response ChangeInternalProductCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeInternalProductCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeInternalProductCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeInternalProductCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeInternalProductCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeInternalProductCompartmentResponse")
	}
	return
}

// changeInternalProductCompartment implements the OCIOperation interface (enables retrying operations)
func (client ProductInternalClient) changeInternalProductCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/internal/product/{productId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeInternalProductCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productInternal", "ChangeInternalProductCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductInternal", "ChangeInternalProductCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateInternalProduct Used by the OCI service teams to create a product to control the availability of the relative resources
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/CreateInternalProduct.go.html to see an example of how to use CreateInternalProduct API.
// A default retry strategy applies to this operation CreateInternalProduct()
func (client ProductInternalClient) CreateInternalProduct(ctx context.Context, request CreateInternalProductRequest) (response CreateInternalProductResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createInternalProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInternalProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInternalProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInternalProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInternalProductResponse")
	}
	return
}

// createInternalProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductInternalClient) createInternalProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/internal/product", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateInternalProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productInternal", "CreateInternalProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductInternal", "CreateInternalProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteInternalProduct Delete the product
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/DeleteInternalProduct.go.html to see an example of how to use DeleteInternalProduct API.
// A default retry strategy applies to this operation DeleteInternalProduct()
func (client ProductInternalClient) DeleteInternalProduct(ctx context.Context, request DeleteInternalProductRequest) (response DeleteInternalProductResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInternalProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteInternalProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteInternalProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInternalProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInternalProductResponse")
	}
	return
}

// deleteInternalProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductInternalClient) deleteInternalProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/internal/product/{productId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteInternalProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productInternal", "DeleteInternalProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductInternal", "DeleteInternalProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInternalProduct Get the product by ID
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/GetInternalProduct.go.html to see an example of how to use GetInternalProduct API.
// A default retry strategy applies to this operation GetInternalProduct()
func (client ProductInternalClient) GetInternalProduct(ctx context.Context, request GetInternalProductRequest) (response GetInternalProductResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInternalProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInternalProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInternalProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInternalProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInternalProductResponse")
	}
	return
}

// getInternalProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductInternalClient) getInternalProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/product/{productId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInternalProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productInternal", "GetInternalProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductInternal", "GetInternalProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalProducts List the products with the filters
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/ListInternalProducts.go.html to see an example of how to use ListInternalProducts API.
// A default retry strategy applies to this operation ListInternalProducts()
func (client ProductInternalClient) ListInternalProducts(ctx context.Context, request ListInternalProductsRequest) (response ListInternalProductsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalProducts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalProductsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalProductsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalProductsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalProductsResponse")
	}
	return
}

// listInternalProducts implements the OCIOperation interface (enables retrying operations)
func (client ProductInternalClient) listInternalProducts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/products", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalProductsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productInternal", "ListInternalProducts")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductInternal", "ListInternalProducts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateInternalProduct Update the product
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/UpdateInternalProduct.go.html to see an example of how to use UpdateInternalProduct API.
// A default retry strategy applies to this operation UpdateInternalProduct()
func (client ProductInternalClient) UpdateInternalProduct(ctx context.Context, request UpdateInternalProductRequest) (response UpdateInternalProductResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateInternalProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInternalProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInternalProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInternalProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInternalProductResponse")
	}
	return
}

// updateInternalProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductInternalClient) updateInternalProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/internal/product/{productId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInternalProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productInternal", "UpdateInternalProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductInternal", "UpdateInternalProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
