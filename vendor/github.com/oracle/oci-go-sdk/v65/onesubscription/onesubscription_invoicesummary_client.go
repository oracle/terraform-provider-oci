// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// InvoiceSummaryClient a client for InvoiceSummary
type InvoiceSummaryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewInvoiceSummaryClientWithConfigurationProvider Creates a new default InvoiceSummary client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewInvoiceSummaryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client InvoiceSummaryClient, err error) {
	if enabled := common.CheckForEnabledServices("onesubscription"); !enabled {
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
	return newInvoiceSummaryClientFromBaseClient(baseClient, provider)
}

// NewInvoiceSummaryClientWithOboToken Creates a new default InvoiceSummary client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewInvoiceSummaryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client InvoiceSummaryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newInvoiceSummaryClientFromBaseClient(baseClient, configProvider)
}

func newInvoiceSummaryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client InvoiceSummaryClient, err error) {
	// InvoiceSummary service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("InvoiceSummary"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = InvoiceSummaryClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *InvoiceSummaryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("onesubscription", "https://identity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *InvoiceSummaryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *InvoiceSummaryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListInvoicelineComputedUsages This is a collection API which returns a list of Invoiced Computed Usages for given Invoiceline id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListInvoicelineComputedUsages.go.html to see an example of how to use ListInvoicelineComputedUsages API.
func (client InvoiceSummaryClient) ListInvoicelineComputedUsages(ctx context.Context, request ListInvoicelineComputedUsagesRequest) (response ListInvoicelineComputedUsagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInvoicelineComputedUsages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInvoicelineComputedUsagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInvoicelineComputedUsagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInvoicelineComputedUsagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInvoicelineComputedUsagesResponse")
	}
	return
}

// listInvoicelineComputedUsages implements the OCIOperation interface (enables retrying operations)
func (client InvoiceSummaryClient) listInvoicelineComputedUsages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/invoiceLineComputedUsages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInvoicelineComputedUsagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "InvoiceSummary", "ListInvoicelineComputedUsages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInvoices This is a collection API which returns a list of Invoices for given filters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/onesubscription/ListInvoices.go.html to see an example of how to use ListInvoices API.
func (client InvoiceSummaryClient) ListInvoices(ctx context.Context, request ListInvoicesRequest) (response ListInvoicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client InvoiceSummaryClient) listInvoices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/invoice", binaryReqBody, extraHeaders)
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
		err = common.PostProcessServiceError(err, "InvoiceSummary", "ListInvoices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
