// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// InvoiceServiceClient a client for InvoiceService
type InvoiceServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewInvoiceServiceClientWithConfigurationProvider Creates a new default InvoiceService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewInvoiceServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client InvoiceServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("ospgateway"); !enabled {
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
	return newInvoiceServiceClientFromBaseClient(baseClient, provider)
}

// NewInvoiceServiceClientWithOboToken Creates a new default InvoiceService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewInvoiceServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client InvoiceServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newInvoiceServiceClientFromBaseClient(baseClient, configProvider)
}

func newInvoiceServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client InvoiceServiceClient, err error) {
	// InvoiceService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("InvoiceService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = InvoiceServiceClient{BaseClient: baseClient}
	client.BasePath = "20191001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *InvoiceServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ospgateway", "https://osp-oci-integ.osp.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *InvoiceServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *InvoiceServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DownloadPdfContent Returns an invoice in pdf format
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/DownloadPdfContent.go.html to see an example of how to use DownloadPdfContent API.
// A default retry strategy applies to this operation DownloadPdfContent()
func (client InvoiceServiceClient) DownloadPdfContent(ctx context.Context, request DownloadPdfContentRequest) (response DownloadPdfContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadPdfContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadPdfContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadPdfContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadPdfContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadPdfContentResponse")
	}
	return
}

// downloadPdfContent implements the OCIOperation interface (enables retrying operations)
func (client InvoiceServiceClient) downloadPdfContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/invoices/{internalInvoiceId}/actions/downloadPdfContent", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadPdfContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "InvoiceService", "DownloadPdfContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInvoice Returns an invoice by invoice id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/GetInvoice.go.html to see an example of how to use GetInvoice API.
// A default retry strategy applies to this operation GetInvoice()
func (client InvoiceServiceClient) GetInvoice(ctx context.Context, request GetInvoiceRequest) (response GetInvoiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInvoice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInvoiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInvoiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInvoiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInvoiceResponse")
	}
	return
}

// getInvoice implements the OCIOperation interface (enables retrying operations)
func (client InvoiceServiceClient) getInvoice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/invoices/{internalInvoiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInvoiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "InvoiceService", "GetInvoice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInvoiceLines Returns the invoice product list by invoice id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/ListInvoiceLines.go.html to see an example of how to use ListInvoiceLines API.
// A default retry strategy applies to this operation ListInvoiceLines()
func (client InvoiceServiceClient) ListInvoiceLines(ctx context.Context, request ListInvoiceLinesRequest) (response ListInvoiceLinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInvoiceLines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInvoiceLinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInvoiceLinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInvoiceLinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInvoiceLinesResponse")
	}
	return
}

// listInvoiceLines implements the OCIOperation interface (enables retrying operations)
func (client InvoiceServiceClient) listInvoiceLines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/invoices/{internalInvoiceId}/invoiceLines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInvoiceLinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "InvoiceService", "ListInvoiceLines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInvoices Returns a list of invoices
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/ListInvoices.go.html to see an example of how to use ListInvoices API.
// A default retry strategy applies to this operation ListInvoices()
func (client InvoiceServiceClient) ListInvoices(ctx context.Context, request ListInvoicesRequest) (response ListInvoicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInvoices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInvoicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInvoicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInvoicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInvoicesResponse")
	}
	return
}

// listInvoices implements the OCIOperation interface (enables retrying operations)
func (client InvoiceServiceClient) listInvoices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/invoices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInvoicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "InvoiceService", "ListInvoices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PayInvoice Pay an invoice
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/PayInvoice.go.html to see an example of how to use PayInvoice API.
// A default retry strategy applies to this operation PayInvoice()
func (client InvoiceServiceClient) PayInvoice(ctx context.Context, request PayInvoiceRequest) (response PayInvoiceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.payInvoice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PayInvoiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PayInvoiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PayInvoiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PayInvoiceResponse")
	}
	return
}

// payInvoice implements the OCIOperation interface (enables retrying operations)
func (client InvoiceServiceClient) payInvoice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/invoices/{internalInvoiceId}/actions/pay", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PayInvoiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "InvoiceService", "PayInvoice", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
