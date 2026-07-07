// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// GuardedDataPipelineClient a client for GuardedDataPipeline
type GuardedDataPipelineClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGuardedDataPipelineClientWithConfigurationProvider Creates a new default GuardedDataPipeline client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGuardedDataPipelineClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GuardedDataPipelineClient, err error) {
	if enabled := common.CheckForEnabledServices("gdp"); !enabled {
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
	return newGuardedDataPipelineClientFromBaseClient(baseClient, provider)
}

// NewGuardedDataPipelineClientWithOboToken Creates a new default GuardedDataPipeline client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGuardedDataPipelineClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GuardedDataPipelineClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGuardedDataPipelineClientFromBaseClient(baseClient, configProvider)
}

func newGuardedDataPipelineClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GuardedDataPipelineClient, err error) {
	// GuardedDataPipeline service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GuardedDataPipeline"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GuardedDataPipelineClient{BaseClient: baseClient}
	client.BasePath = "20230301"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GuardedDataPipelineClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("gdp", "https://gdp.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GuardedDataPipelineClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GuardedDataPipelineClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeGdpMpPipelineCompartment Moves a pipeline resource from one compartment to another. When provided, if-match is checked against etag values of the resource.
// A default retry strategy applies to this operation ChangeGdpMpPipelineCompartment()
func (client GuardedDataPipelineClient) ChangeGdpMpPipelineCompartment(ctx context.Context, request ChangeGdpMpPipelineCompartmentRequest) (response ChangeGdpMpPipelineCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeGdpMpPipelineCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeGdpMpPipelineCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeGdpMpPipelineCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeGdpMpPipelineCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeGdpMpPipelineCompartmentResponse")
	}
	return
}

// changeGdpMpPipelineCompartment implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) changeGdpMpPipelineCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpMpPipelines/{gdpMpPipelineId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeGdpMpPipelineCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ChangeGdpMpPipelineCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ChangeGdpMpPipelineCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeGdpPipelineCompartment Moves a pipeline resource from one compartment to another. When provided, if-match is checked against etag values of the resource.
// A default retry strategy applies to this operation ChangeGdpPipelineCompartment()
func (client GuardedDataPipelineClient) ChangeGdpPipelineCompartment(ctx context.Context, request ChangeGdpPipelineCompartmentRequest) (response ChangeGdpPipelineCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeGdpPipelineCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeGdpPipelineCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeGdpPipelineCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeGdpPipelineCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeGdpPipelineCompartmentResponse")
	}
	return
}

// changeGdpPipelineCompartment implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) changeGdpPipelineCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpPipelines/{gdpPipelineId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeGdpPipelineCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ChangeGdpPipelineCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ChangeGdpPipelineCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGdpMpPipeline Creates a new pipeline.
// A default retry strategy applies to this operation CreateGdpMpPipeline()
func (client GuardedDataPipelineClient) CreateGdpMpPipeline(ctx context.Context, request CreateGdpMpPipelineRequest) (response CreateGdpMpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGdpMpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGdpMpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGdpMpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGdpMpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGdpMpPipelineResponse")
	}
	return
}

// createGdpMpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) createGdpMpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpMpPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGdpMpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "CreateGdpMpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "CreateGdpMpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGdpPipeline Creates a new pipeline.
// A default retry strategy applies to this operation CreateGdpPipeline()
func (client GuardedDataPipelineClient) CreateGdpPipeline(ctx context.Context, request CreateGdpPipelineRequest) (response CreateGdpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGdpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGdpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGdpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGdpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGdpPipelineResponse")
	}
	return
}

// createGdpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) createGdpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGdpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "CreateGdpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "CreateGdpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGdpMpPipeline Deletes a pipeline by identifier.
// A default retry strategy applies to this operation DeleteGdpMpPipeline()
func (client GuardedDataPipelineClient) DeleteGdpMpPipeline(ctx context.Context, request DeleteGdpMpPipelineRequest) (response DeleteGdpMpPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGdpMpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGdpMpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGdpMpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGdpMpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGdpMpPipelineResponse")
	}
	return
}

// deleteGdpMpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) deleteGdpMpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/gdpMpPipelines/{gdpMpPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGdpMpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "DeleteGdpMpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "DeleteGdpMpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGdpPipeline Deletes a pipeline by identifier.
// A default retry strategy applies to this operation DeleteGdpPipeline()
func (client GuardedDataPipelineClient) DeleteGdpPipeline(ctx context.Context, request DeleteGdpPipelineRequest) (response DeleteGdpPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGdpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGdpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGdpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGdpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGdpPipelineResponse")
	}
	return
}

// deleteGdpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) deleteGdpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/gdpPipelines/{gdpPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGdpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "DeleteGdpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "DeleteGdpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGdpMpPipeline Retrieves a pipeline by identifier.
// A default retry strategy applies to this operation GetGdpMpPipeline()
func (client GuardedDataPipelineClient) GetGdpMpPipeline(ctx context.Context, request GetGdpMpPipelineRequest) (response GetGdpMpPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGdpMpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGdpMpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGdpMpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGdpMpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGdpMpPipelineResponse")
	}
	return
}

// getGdpMpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) getGdpMpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpMpPipelines/{gdpMpPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGdpMpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "GetGdpMpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "GetGdpMpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGdpPipeline Retrieves a pipeline by identifier.
// A default retry strategy applies to this operation GetGdpPipeline()
func (client GuardedDataPipelineClient) GetGdpPipeline(ctx context.Context, request GetGdpPipelineRequest) (response GetGdpPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGdpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGdpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGdpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGdpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGdpPipelineResponse")
	}
	return
}

// getGdpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) getGdpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpPipelines/{gdpPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGdpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "GetGdpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "GetGdpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGdpWorkRequest Gets details of the work request with the given ID.
// A default retry strategy applies to this operation GetGdpWorkRequest()
func (client GuardedDataPipelineClient) GetGdpWorkRequest(ctx context.Context, request GetGdpWorkRequestRequest) (response GetGdpWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGdpWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGdpWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGdpWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGdpWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGdpWorkRequestResponse")
	}
	return
}

// getGdpWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) getGdpWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGdpWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "GetGdpWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "GetGdpWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTransferStatusPipelineTransferId Gets the transfer status for the provided pipeline and transfer identifiers.
// A default retry strategy applies to this operation GetTransferStatusPipelineTransferId()
func (client GuardedDataPipelineClient) GetTransferStatusPipelineTransferId(ctx context.Context, request GetTransferStatusPipelineTransferIdRequest) (response GetTransferStatusPipelineTransferIdResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTransferStatusPipelineTransferId, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTransferStatusPipelineTransferIdResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTransferStatusPipelineTransferIdResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTransferStatusPipelineTransferIdResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTransferStatusPipelineTransferIdResponse")
	}
	return
}

// getTransferStatusPipelineTransferId implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) getTransferStatusPipelineTransferId(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferStatus/pipeline/{pipelineId}/transferId/{transferStatusId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTransferStatusPipelineTransferIdResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "GetTransferStatusPipelineTransferId")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "GetTransferStatusPipelineTransferId", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGdpMpPipelines Returns a list of pipelines.
// A default retry strategy applies to this operation ListGdpMpPipelines()
func (client GuardedDataPipelineClient) ListGdpMpPipelines(ctx context.Context, request ListGdpMpPipelinesRequest) (response ListGdpMpPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGdpMpPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGdpMpPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGdpMpPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGdpMpPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGdpMpPipelinesResponse")
	}
	return
}

// listGdpMpPipelines implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) listGdpMpPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpMpPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGdpMpPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ListGdpMpPipelines")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ListGdpMpPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGdpPipelines Returns a list of pipelines.
// A default retry strategy applies to this operation ListGdpPipelines()
func (client GuardedDataPipelineClient) ListGdpPipelines(ctx context.Context, request ListGdpPipelinesRequest) (response ListGdpPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGdpPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGdpPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGdpPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGdpPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGdpPipelinesResponse")
	}
	return
}

// listGdpPipelines implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) listGdpPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGdpPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ListGdpPipelines")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ListGdpPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGdpWorkRequestErrors Returns a (paginated) list of errors for the work request with the given ID.
// A default retry strategy applies to this operation ListGdpWorkRequestErrors()
func (client GuardedDataPipelineClient) ListGdpWorkRequestErrors(ctx context.Context, request ListGdpWorkRequestErrorsRequest) (response ListGdpWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGdpWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGdpWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGdpWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGdpWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGdpWorkRequestErrorsResponse")
	}
	return
}

// listGdpWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) listGdpWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpWorkRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGdpWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ListGdpWorkRequestErrors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ListGdpWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGdpWorkRequestLogs Returns a (paginated) list of logs for the work request with the given ID.
// A default retry strategy applies to this operation ListGdpWorkRequestLogs()
func (client GuardedDataPipelineClient) ListGdpWorkRequestLogs(ctx context.Context, request ListGdpWorkRequestLogsRequest) (response ListGdpWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGdpWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGdpWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGdpWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGdpWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGdpWorkRequestLogsResponse")
	}
	return
}

// listGdpWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) listGdpWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpWorkRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGdpWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ListGdpWorkRequestLogs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ListGdpWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGdpWorkRequests Lists the work requests in a compartment.
// A default retry strategy applies to this operation ListGdpWorkRequests()
func (client GuardedDataPipelineClient) ListGdpWorkRequests(ctx context.Context, request ListGdpWorkRequestsRequest) (response ListGdpWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGdpWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGdpWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGdpWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGdpWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGdpWorkRequestsResponse")
	}
	return
}

// listGdpWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) listGdpWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/gdpWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGdpWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ListGdpWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ListGdpWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTransferStatusPipelines Lists transfer statuses for the provided pipeline, filename, and checksum.
// A default retry strategy applies to this operation ListTransferStatusPipelines()
func (client GuardedDataPipelineClient) ListTransferStatusPipelines(ctx context.Context, request ListTransferStatusPipelinesRequest) (response ListTransferStatusPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferStatusPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTransferStatusPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTransferStatusPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferStatusPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferStatusPipelinesResponse")
	}
	return
}

// listTransferStatusPipelines implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) listTransferStatusPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferStatus/pipeline/{pipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTransferStatusPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ListTransferStatusPipelines")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ListTransferStatusPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTransferStatuses Lists transfer statuses for a CDS pipeline, filename, and checksum combination.
// A default retry strategy applies to this operation ListTransferStatuses()
func (client GuardedDataPipelineClient) ListTransferStatuses(ctx context.Context, request ListTransferStatusesRequest) (response ListTransferStatusesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTransferStatuses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTransferStatusesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTransferStatusesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTransferStatusesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTransferStatusesResponse")
	}
	return
}

// listTransferStatuses implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) listTransferStatuses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/transferStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTransferStatusesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "ListTransferStatuses")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "ListTransferStatuses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PeerGdpMpPipeline Peers the pipeline.
// A default retry strategy applies to this operation PeerGdpMpPipeline()
func (client GuardedDataPipelineClient) PeerGdpMpPipeline(ctx context.Context, request PeerGdpMpPipelineRequest) (response PeerGdpMpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.peerGdpMpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PeerGdpMpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PeerGdpMpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PeerGdpMpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PeerGdpMpPipelineResponse")
	}
	return
}

// peerGdpMpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) peerGdpMpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpMpPipelines/{gdpMpPipelineId}/actions/peer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PeerGdpMpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "PeerGdpMpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "PeerGdpMpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PeerGdpPipeline Peers the pipeline.
// A default retry strategy applies to this operation PeerGdpPipeline()
func (client GuardedDataPipelineClient) PeerGdpPipeline(ctx context.Context, request PeerGdpPipelineRequest) (response PeerGdpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.peerGdpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PeerGdpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PeerGdpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PeerGdpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PeerGdpPipelineResponse")
	}
	return
}

// peerGdpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) peerGdpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpPipelines/{gdpPipelineId}/actions/peer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PeerGdpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "PeerGdpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "PeerGdpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateGdpMpPipelineKeys Rotates the pipeline keys by resending the control message with a fresh pair of keys.
// A default retry strategy applies to this operation RotateGdpMpPipelineKeys()
func (client GuardedDataPipelineClient) RotateGdpMpPipelineKeys(ctx context.Context, request RotateGdpMpPipelineKeysRequest) (response RotateGdpMpPipelineKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateGdpMpPipelineKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateGdpMpPipelineKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateGdpMpPipelineKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateGdpMpPipelineKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateGdpMpPipelineKeysResponse")
	}
	return
}

// rotateGdpMpPipelineKeys implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) rotateGdpMpPipelineKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpMpPipelines/{gdpMpPipelineId}/actions/rotateKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateGdpMpPipelineKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "RotateGdpMpPipelineKeys")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "RotateGdpMpPipelineKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateGdpPipelineKeys Rotates the pipeline keys by resending the control message with a fresh pair of keys.
// A default retry strategy applies to this operation RotateGdpPipelineKeys()
func (client GuardedDataPipelineClient) RotateGdpPipelineKeys(ctx context.Context, request RotateGdpPipelineKeysRequest) (response RotateGdpPipelineKeysResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rotateGdpPipelineKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateGdpPipelineKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateGdpPipelineKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateGdpPipelineKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateGdpPipelineKeysResponse")
	}
	return
}

// rotateGdpPipelineKeys implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) rotateGdpPipelineKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpPipelines/{gdpPipelineId}/actions/rotateKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateGdpPipelineKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "RotateGdpPipelineKeys")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "RotateGdpPipelineKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartGdpMpPipeline Starts the pipeline.
// A default retry strategy applies to this operation StartGdpMpPipeline()
func (client GuardedDataPipelineClient) StartGdpMpPipeline(ctx context.Context, request StartGdpMpPipelineRequest) (response StartGdpMpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startGdpMpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartGdpMpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartGdpMpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartGdpMpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartGdpMpPipelineResponse")
	}
	return
}

// startGdpMpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) startGdpMpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpMpPipelines/{gdpMpPipelineId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartGdpMpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "StartGdpMpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "StartGdpMpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartGdpPipeline Starts the pipeline.
// A default retry strategy applies to this operation StartGdpPipeline()
func (client GuardedDataPipelineClient) StartGdpPipeline(ctx context.Context, request StartGdpPipelineRequest) (response StartGdpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startGdpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartGdpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartGdpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartGdpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartGdpPipelineResponse")
	}
	return
}

// startGdpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) startGdpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpPipelines/{gdpPipelineId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartGdpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "StartGdpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "StartGdpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopGdpMpPipeline Stops the pipeline.
// A default retry strategy applies to this operation StopGdpMpPipeline()
func (client GuardedDataPipelineClient) StopGdpMpPipeline(ctx context.Context, request StopGdpMpPipelineRequest) (response StopGdpMpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopGdpMpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopGdpMpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopGdpMpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopGdpMpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopGdpMpPipelineResponse")
	}
	return
}

// stopGdpMpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) stopGdpMpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpMpPipelines/{gdpMpPipelineId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopGdpMpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "StopGdpMpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "StopGdpMpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopGdpPipeline Stops the pipeline.
// A default retry strategy applies to this operation StopGdpPipeline()
func (client GuardedDataPipelineClient) StopGdpPipeline(ctx context.Context, request StopGdpPipelineRequest) (response StopGdpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopGdpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopGdpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopGdpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopGdpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopGdpPipelineResponse")
	}
	return
}

// stopGdpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) stopGdpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/gdpPipelines/{gdpPipelineId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopGdpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "StopGdpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "StopGdpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGdpMpPipeline Updates the pipeline.
// A default retry strategy applies to this operation UpdateGdpMpPipeline()
func (client GuardedDataPipelineClient) UpdateGdpMpPipeline(ctx context.Context, request UpdateGdpMpPipelineRequest) (response UpdateGdpMpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateGdpMpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGdpMpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGdpMpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGdpMpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGdpMpPipelineResponse")
	}
	return
}

// updateGdpMpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) updateGdpMpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/gdpMpPipelines/{gdpMpPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGdpMpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "UpdateGdpMpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "UpdateGdpMpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGdpPipeline Updates the pipeline.
// A default retry strategy applies to this operation UpdateGdpPipeline()
func (client GuardedDataPipelineClient) UpdateGdpPipeline(ctx context.Context, request UpdateGdpPipelineRequest) (response UpdateGdpPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateGdpPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGdpPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGdpPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGdpPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGdpPipelineResponse")
	}
	return
}

// updateGdpPipeline implements the OCIOperation interface (enables retrying operations)
func (client GuardedDataPipelineClient) updateGdpPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/gdpPipelines/{gdpPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGdpPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "guardedDataPipeline", "UpdateGdpPipeline")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "GuardedDataPipeline", "UpdateGdpPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
