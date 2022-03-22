// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"github.com/oracle/oci-go-sdk/v63/common/auth"
	"net/http"
)

//ApiGatewayClient a client for ApiGateway
type ApiGatewayClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApiGatewayClientWithConfigurationProvider Creates a new default ApiGateway client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApiGatewayClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApiGatewayClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newApiGatewayClientFromBaseClient(baseClient, provider)
}

// NewApiGatewayClientWithOboToken Creates a new default ApiGateway client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewApiGatewayClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApiGatewayClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newApiGatewayClientFromBaseClient(baseClient, configProvider)
}

func newApiGatewayClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApiGatewayClient, err error) {
	// ApiGateway service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ApiGatewayClient{BaseClient: baseClient}
	client.BasePath = "20190501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApiGatewayClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apigateway", "https://apigateway.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApiGatewayClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApiGatewayClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeApiCompartment Changes the API compartment.
func (client ApiGatewayClient) ChangeApiCompartment(ctx context.Context, request ChangeApiCompartmentRequest) (response ChangeApiCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeApiCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeApiCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeApiCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeApiCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeApiCompartmentResponse")
	}
	return
}

// changeApiCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) changeApiCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/apis/{apiId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeApiCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Api/ChangeApiCompartment"
		err = common.PostProcessServiceError(err, "ApiGateway", "ChangeApiCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCatalogCompartment Changes the API Catalog compartment.
func (client ApiGatewayClient) ChangeCatalogCompartment(ctx context.Context, request ChangeCatalogCompartmentRequest) (response ChangeCatalogCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCatalogCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCatalogCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCatalogCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCatalogCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCatalogCompartmentResponse")
	}
	return
}

// changeCatalogCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) changeCatalogCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCatalogCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Catalog/ChangeCatalogCompartment"
		err = common.PostProcessServiceError(err, "ApiGateway", "ChangeCatalogCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCertificateCompartment Changes the certificate compartment.
func (client ApiGatewayClient) ChangeCertificateCompartment(ctx context.Context, request ChangeCertificateCompartmentRequest) (response ChangeCertificateCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCertificateCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCertificateCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCertificateCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCertificateCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCertificateCompartmentResponse")
	}
	return
}

// changeCertificateCompartment implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) changeCertificateCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates/{certificateId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCertificateCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Certificate/ChangeCertificateCompartment"
		err = common.PostProcessServiceError(err, "ApiGateway", "ChangeCertificateCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateApi Creates a new API.
func (client ApiGatewayClient) CreateApi(ctx context.Context, request CreateApiRequest) (response CreateApiResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createApi, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateApiResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateApiResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApiResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApiResponse")
	}
	return
}

// createApi implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) createApi(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/apis", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateApiResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ApiGateway", "CreateApi", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCatalog Creates a new catalog.
func (client ApiGatewayClient) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (response CreateCatalogResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCatalogResponse")
	}
	return
}

// createCatalog implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) createCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Catalog/CreateCatalog"
		err = common.PostProcessServiceError(err, "ApiGateway", "CreateCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCertificate Creates a new Certificate.
func (client ApiGatewayClient) CreateCertificate(ctx context.Context, request CreateCertificateRequest) (response CreateCertificateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCertificateResponse")
	}
	return
}

// createCertificate implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) createCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ApiGateway", "CreateCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSdk Creates a new SDK.
func (client ApiGatewayClient) CreateSdk(ctx context.Context, request CreateSdkRequest) (response CreateSdkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSdk, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSdkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSdkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSdkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSdkResponse")
	}
	return
}

// createSdk implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) createSdk(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sdks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSdkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Sdk/CreateSdk"
		err = common.PostProcessServiceError(err, "ApiGateway", "CreateSdk", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApi Deletes the API with the given identifier.
func (client ApiGatewayClient) DeleteApi(ctx context.Context, request DeleteApiRequest) (response DeleteApiResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApi, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApiResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApiResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApiResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApiResponse")
	}
	return
}

// deleteApi implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) deleteApi(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/apis/{apiId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApiResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Api/DeleteApi"
		err = common.PostProcessServiceError(err, "ApiGateway", "DeleteApi", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCatalog Deletes the API Catalog with the given identifier.
func (client ApiGatewayClient) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (response DeleteCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCatalogResponse")
	}
	return
}

// deleteCatalog implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) deleteCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Catalog/DeleteCatalog"
		err = common.PostProcessServiceError(err, "ApiGateway", "DeleteCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCertificate Deletes the certificate with the given identifier.
func (client ApiGatewayClient) DeleteCertificate(ctx context.Context, request DeleteCertificateRequest) (response DeleteCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCertificateResponse")
	}
	return
}

// deleteCertificate implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) deleteCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/certificates/{certificateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Certificate/DeleteCertificate"
		err = common.PostProcessServiceError(err, "ApiGateway", "DeleteCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSdk Deletes provided SDK.
func (client ApiGatewayClient) DeleteSdk(ctx context.Context, request DeleteSdkRequest) (response DeleteSdkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSdk, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSdkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSdkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSdkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSdkResponse")
	}
	return
}

// deleteSdk implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) deleteSdk(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sdks/{sdkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSdkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Sdk/DeleteSdk"
		err = common.PostProcessServiceError(err, "ApiGateway", "DeleteSdk", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApi Gets an API by identifier.
func (client ApiGatewayClient) GetApi(ctx context.Context, request GetApiRequest) (response GetApiResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApi, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApiResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApiResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApiResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApiResponse")
	}
	return
}

// getApi implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) getApi(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apis/{apiId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApiResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Api/GetApi"
		err = common.PostProcessServiceError(err, "ApiGateway", "GetApi", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApiContent Get the raw API content.
func (client ApiGatewayClient) GetApiContent(ctx context.Context, request GetApiContentRequest) (response GetApiContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApiContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApiContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApiContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApiContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApiContentResponse")
	}
	return
}

// getApiContent implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) getApiContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apis/{apiId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApiContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Api/GetApiContent"
		err = common.PostProcessServiceError(err, "ApiGateway", "GetApiContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApiDeploymentSpecification Gets an API Deployment specification by identifier.
func (client ApiGatewayClient) GetApiDeploymentSpecification(ctx context.Context, request GetApiDeploymentSpecificationRequest) (response GetApiDeploymentSpecificationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApiDeploymentSpecification, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApiDeploymentSpecificationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApiDeploymentSpecificationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApiDeploymentSpecificationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApiDeploymentSpecificationResponse")
	}
	return
}

// getApiDeploymentSpecification implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) getApiDeploymentSpecification(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apis/{apiId}/deploymentSpecification", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApiDeploymentSpecificationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/ApiSpecification/GetApiDeploymentSpecification"
		err = common.PostProcessServiceError(err, "ApiGateway", "GetApiDeploymentSpecification", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApiValidations Gets the API validation results.
func (client ApiGatewayClient) GetApiValidations(ctx context.Context, request GetApiValidationsRequest) (response GetApiValidationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApiValidations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApiValidationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApiValidationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApiValidationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApiValidationsResponse")
	}
	return
}

// getApiValidations implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) getApiValidations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apis/{apiId}/validations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApiValidationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/ApiValidations/GetApiValidations"
		err = common.PostProcessServiceError(err, "ApiGateway", "GetApiValidations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCatalog Gets an API Catalog by identifier.
func (client ApiGatewayClient) GetCatalog(ctx context.Context, request GetCatalogRequest) (response GetCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogResponse")
	}
	return
}

// getCatalog implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) getCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Catalog/GetCatalog"
		err = common.PostProcessServiceError(err, "ApiGateway", "GetCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificate Gets a certificate by identifier.
func (client ApiGatewayClient) GetCertificate(ctx context.Context, request GetCertificateRequest) (response GetCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCertificateResponse")
	}
	return
}

// getCertificate implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) getCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificates/{certificateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Certificate/GetCertificate"
		err = common.PostProcessServiceError(err, "ApiGateway", "GetCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSdk Return object store downloadable URL and metadata.
func (client ApiGatewayClient) GetSdk(ctx context.Context, request GetSdkRequest) (response GetSdkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSdk, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSdkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSdkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSdkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSdkResponse")
	}
	return
}

// getSdk implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) getSdk(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sdks/{sdkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSdkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Sdk/GetSdk"
		err = common.PostProcessServiceError(err, "ApiGateway", "GetSdk", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApis Returns a list of APIs.
func (client ApiGatewayClient) ListApis(ctx context.Context, request ListApisRequest) (response ListApisResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApis, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApisResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApisResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApisResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApisResponse")
	}
	return
}

// listApis implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) listApis(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/apis", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApisResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Api/ListApis"
		err = common.PostProcessServiceError(err, "ApiGateway", "ListApis", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCatalogs Returns a list of catalogs.
func (client ApiGatewayClient) ListCatalogs(ctx context.Context, request ListCatalogsRequest) (response ListCatalogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCatalogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCatalogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogsResponse")
	}
	return
}

// listCatalogs implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) listCatalogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCatalogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Catalog/ListCatalogs"
		err = common.PostProcessServiceError(err, "ApiGateway", "ListCatalogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCertificates Returns a list of certificates.
func (client ApiGatewayClient) ListCertificates(ctx context.Context, request ListCertificatesRequest) (response ListCertificatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCertificates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCertificatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCertificatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCertificatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCertificatesResponse")
	}
	return
}

// listCertificates implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) listCertificates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Certificate/ListCertificates"
		err = common.PostProcessServiceError(err, "ApiGateway", "ListCertificates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSdkLanguageTypes Lists programming languages in which SDK can be generated.
func (client ApiGatewayClient) ListSdkLanguageTypes(ctx context.Context, request ListSdkLanguageTypesRequest) (response ListSdkLanguageTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSdkLanguageTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSdkLanguageTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSdkLanguageTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSdkLanguageTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSdkLanguageTypesResponse")
	}
	return
}

// listSdkLanguageTypes implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) listSdkLanguageTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sdkLanguageTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSdkLanguageTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/SdkLanguageTypeSummary/ListSdkLanguageTypes"
		err = common.PostProcessServiceError(err, "ApiGateway", "ListSdkLanguageTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSdks Returns list of generated SDKs.
func (client ApiGatewayClient) ListSdks(ctx context.Context, request ListSdksRequest) (response ListSdksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSdks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSdksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSdksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSdksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSdksResponse")
	}
	return
}

// listSdks implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) listSdks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sdks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSdksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Sdk/ListSdks"
		err = common.PostProcessServiceError(err, "ApiGateway", "ListSdks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateApi Updates the API with the given identifier.
func (client ApiGatewayClient) UpdateApi(ctx context.Context, request UpdateApiRequest) (response UpdateApiResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateApi, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateApiResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateApiResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateApiResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateApiResponse")
	}
	return
}

// updateApi implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) updateApi(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/apis/{apiId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateApiResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Api/UpdateApi"
		err = common.PostProcessServiceError(err, "ApiGateway", "UpdateApi", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCatalog Updates the API Catalog with the given identifier.
func (client ApiGatewayClient) UpdateCatalog(ctx context.Context, request UpdateCatalogRequest) (response UpdateCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCatalogResponse")
	}
	return
}

// updateCatalog implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) updateCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Catalog/UpdateCatalog"
		err = common.PostProcessServiceError(err, "ApiGateway", "UpdateCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCertificate Updates a certificate with the given identifier
func (client ApiGatewayClient) UpdateCertificate(ctx context.Context, request UpdateCertificateRequest) (response UpdateCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCertificateResponse")
	}
	return
}

// updateCertificate implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) updateCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/certificates/{certificateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Certificate/UpdateCertificate"
		err = common.PostProcessServiceError(err, "ApiGateway", "UpdateCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSdk Updates the SDK with the given identifier.
func (client ApiGatewayClient) UpdateSdk(ctx context.Context, request UpdateSdkRequest) (response UpdateSdkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSdk, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSdkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSdkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSdkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSdkResponse")
	}
	return
}

// updateSdk implements the OCIOperation interface (enables retrying operations)
func (client ApiGatewayClient) updateSdk(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sdks/{sdkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSdkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/api-gateway/20190501/Sdk/UpdateSdk"
		err = common.PostProcessServiceError(err, "ApiGateway", "UpdateSdk", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
