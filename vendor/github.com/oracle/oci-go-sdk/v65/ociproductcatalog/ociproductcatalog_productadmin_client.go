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

// ProductAdminClient a client for ProductAdmin
type ProductAdminClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewProductAdminClientWithConfigurationProvider Creates a new default ProductAdmin client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewProductAdminClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ProductAdminClient, err error) {
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
	return newProductAdminClientFromBaseClient(baseClient, provider)
}

// NewProductAdminClientWithOboToken Creates a new default ProductAdmin client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewProductAdminClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ProductAdminClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newProductAdminClientFromBaseClient(baseClient, configProvider)
}

func newProductAdminClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ProductAdminClient, err error) {
	// ProductAdmin service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ProductAdmin"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ProductAdminClient{BaseClient: baseClient}
	client.BasePath = "20250610"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ProductAdminClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ociproductcatalog", "https://cp.product-catalog.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ProductAdminClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ProductAdminClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAdminProductCompartment Change product compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/ChangeAdminProductCompartment.go.html to see an example of how to use ChangeAdminProductCompartment API.
// A default retry strategy applies to this operation ChangeAdminProductCompartment()
func (client ProductAdminClient) ChangeAdminProductCompartment(ctx context.Context, request ChangeAdminProductCompartmentRequest) (response ChangeAdminProductCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAdminProductCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAdminProductCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAdminProductCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAdminProductCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAdminProductCompartmentResponse")
	}
	return
}

// changeAdminProductCompartment implements the OCIOperation interface (enables retrying operations)
func (client ProductAdminClient) changeAdminProductCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/internal/admin/product/{productId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAdminProductCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productAdmin", "ChangeAdminProductCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductAdmin", "ChangeAdminProductCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAdminProduct Used by the OCI service teams to create a product to control the availability of the relative resources
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/CreateAdminProduct.go.html to see an example of how to use CreateAdminProduct API.
// A default retry strategy applies to this operation CreateAdminProduct()
func (client ProductAdminClient) CreateAdminProduct(ctx context.Context, request CreateAdminProductRequest) (response CreateAdminProductResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAdminProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAdminProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAdminProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAdminProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAdminProductResponse")
	}
	return
}

// createAdminProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductAdminClient) createAdminProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/internal/admin/product", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAdminProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productAdmin", "CreateAdminProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductAdmin", "CreateAdminProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAdminProduct Delete the product
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/DeleteAdminProduct.go.html to see an example of how to use DeleteAdminProduct API.
// A default retry strategy applies to this operation DeleteAdminProduct()
func (client ProductAdminClient) DeleteAdminProduct(ctx context.Context, request DeleteAdminProductRequest) (response DeleteAdminProductResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAdminProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAdminProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAdminProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAdminProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAdminProductResponse")
	}
	return
}

// deleteAdminProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductAdminClient) deleteAdminProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/internal/admin/product/{productId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAdminProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productAdmin", "DeleteAdminProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductAdmin", "DeleteAdminProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAdminProduct Get the product by ID
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/GetAdminProduct.go.html to see an example of how to use GetAdminProduct API.
// A default retry strategy applies to this operation GetAdminProduct()
func (client ProductAdminClient) GetAdminProduct(ctx context.Context, request GetAdminProductRequest) (response GetAdminProductResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAdminProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAdminProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAdminProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAdminProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAdminProductResponse")
	}
	return
}

// getAdminProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductAdminClient) getAdminProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/admin/product/{productId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAdminProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productAdmin", "GetAdminProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductAdmin", "GetAdminProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminProducts List the products with the filters
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/ListAdminProducts.go.html to see an example of how to use ListAdminProducts API.
// A default retry strategy applies to this operation ListAdminProducts()
func (client ProductAdminClient) ListAdminProducts(ctx context.Context, request ListAdminProductsRequest) (response ListAdminProductsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminProducts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminProductsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminProductsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminProductsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminProductsResponse")
	}
	return
}

// listAdminProducts implements the OCIOperation interface (enables retrying operations)
func (client ProductAdminClient) listAdminProducts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/admin/products", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminProductsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productAdmin", "ListAdminProducts")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductAdmin", "ListAdminProducts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAdminProduct Update the product
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ociproductcatalog/UpdateAdminProduct.go.html to see an example of how to use UpdateAdminProduct API.
// A default retry strategy applies to this operation UpdateAdminProduct()
func (client ProductAdminClient) UpdateAdminProduct(ctx context.Context, request UpdateAdminProductRequest) (response UpdateAdminProductResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAdminProduct, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAdminProductResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAdminProductResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAdminProductResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAdminProductResponse")
	}
	return
}

// updateAdminProduct implements the OCIOperation interface (enables retrying operations)
func (client ProductAdminClient) updateAdminProduct(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/internal/admin/product/{productId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAdminProductResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "productAdmin", "UpdateAdminProduct")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ProductAdmin", "UpdateAdminProduct", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
