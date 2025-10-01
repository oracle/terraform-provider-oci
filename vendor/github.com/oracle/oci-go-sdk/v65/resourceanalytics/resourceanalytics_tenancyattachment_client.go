// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// TenancyAttachmentClient a client for TenancyAttachment
type TenancyAttachmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewTenancyAttachmentClientWithConfigurationProvider Creates a new default TenancyAttachment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewTenancyAttachmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client TenancyAttachmentClient, err error) {
	if enabled := common.CheckForEnabledServices("resourceanalytics"); !enabled {
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
	return newTenancyAttachmentClientFromBaseClient(baseClient, provider)
}

// NewTenancyAttachmentClientWithOboToken Creates a new default TenancyAttachment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewTenancyAttachmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client TenancyAttachmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newTenancyAttachmentClientFromBaseClient(baseClient, configProvider)
}

func newTenancyAttachmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client TenancyAttachmentClient, err error) {
	// TenancyAttachment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("TenancyAttachment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = TenancyAttachmentClient{BaseClient: baseClient}
	client.BasePath = "20241031"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *TenancyAttachmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("resourceanalytics", "https://resource-analytics.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *TenancyAttachmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *TenancyAttachmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateTenancyAttachment Creates a TenancyAttachment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/CreateTenancyAttachment.go.html to see an example of how to use CreateTenancyAttachment API.
// A default retry strategy applies to this operation CreateTenancyAttachment()
func (client TenancyAttachmentClient) CreateTenancyAttachment(ctx context.Context, request CreateTenancyAttachmentRequest) (response CreateTenancyAttachmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTenancyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTenancyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTenancyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTenancyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTenancyAttachmentResponse")
	}
	return
}

// createTenancyAttachment implements the OCIOperation interface (enables retrying operations)
func (client TenancyAttachmentClient) createTenancyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tenancyAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTenancyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/TenancyAttachment/CreateTenancyAttachment"
		err = common.PostProcessServiceError(err, "TenancyAttachment", "CreateTenancyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTenancyAttachment Deletes a TenancyAttachment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/DeleteTenancyAttachment.go.html to see an example of how to use DeleteTenancyAttachment API.
// A default retry strategy applies to this operation DeleteTenancyAttachment()
func (client TenancyAttachmentClient) DeleteTenancyAttachment(ctx context.Context, request DeleteTenancyAttachmentRequest) (response DeleteTenancyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTenancyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTenancyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTenancyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTenancyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTenancyAttachmentResponse")
	}
	return
}

// deleteTenancyAttachment implements the OCIOperation interface (enables retrying operations)
func (client TenancyAttachmentClient) deleteTenancyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/tenancyAttachments/{tenancyAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTenancyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/TenancyAttachment/DeleteTenancyAttachment"
		err = common.PostProcessServiceError(err, "TenancyAttachment", "DeleteTenancyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTenancyAttachment Gets information about a TenancyAttachment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/GetTenancyAttachment.go.html to see an example of how to use GetTenancyAttachment API.
// A default retry strategy applies to this operation GetTenancyAttachment()
func (client TenancyAttachmentClient) GetTenancyAttachment(ctx context.Context, request GetTenancyAttachmentRequest) (response GetTenancyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTenancyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTenancyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTenancyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTenancyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTenancyAttachmentResponse")
	}
	return
}

// getTenancyAttachment implements the OCIOperation interface (enables retrying operations)
func (client TenancyAttachmentClient) getTenancyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancyAttachments/{tenancyAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTenancyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/TenancyAttachment/GetTenancyAttachment"
		err = common.PostProcessServiceError(err, "TenancyAttachment", "GetTenancyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTenancyAttachments Gets a list of TenancyAttachments.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListTenancyAttachments.go.html to see an example of how to use ListTenancyAttachments API.
// A default retry strategy applies to this operation ListTenancyAttachments()
func (client TenancyAttachmentClient) ListTenancyAttachments(ctx context.Context, request ListTenancyAttachmentsRequest) (response ListTenancyAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTenancyAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTenancyAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTenancyAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTenancyAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTenancyAttachmentsResponse")
	}
	return
}

// listTenancyAttachments implements the OCIOperation interface (enables retrying operations)
func (client TenancyAttachmentClient) listTenancyAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancyAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTenancyAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/TenancyAttachmentCollection/ListTenancyAttachments"
		err = common.PostProcessServiceError(err, "TenancyAttachment", "ListTenancyAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTenancyAttachment Updates a TenancyAttachment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/UpdateTenancyAttachment.go.html to see an example of how to use UpdateTenancyAttachment API.
// A default retry strategy applies to this operation UpdateTenancyAttachment()
func (client TenancyAttachmentClient) UpdateTenancyAttachment(ctx context.Context, request UpdateTenancyAttachmentRequest) (response UpdateTenancyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTenancyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTenancyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTenancyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTenancyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTenancyAttachmentResponse")
	}
	return
}

// updateTenancyAttachment implements the OCIOperation interface (enables retrying operations)
func (client TenancyAttachmentClient) updateTenancyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/tenancyAttachments/{tenancyAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTenancyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/resource-analytics/20241031/TenancyAttachment/UpdateTenancyAttachment"
		err = common.PostProcessServiceError(err, "TenancyAttachment", "UpdateTenancyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
