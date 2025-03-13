// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DbLifeCycleManagementClient a client for DbLifeCycleManagement
type DbLifeCycleManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDbLifeCycleManagementClientWithConfigurationProvider Creates a new default DbLifeCycleManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDbLifeCycleManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DbLifeCycleManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("dblm"); !enabled {
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
	return newDbLifeCycleManagementClientFromBaseClient(baseClient, provider)
}

// NewDbLifeCycleManagementClientWithOboToken Creates a new default DbLifeCycleManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDbLifeCycleManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DbLifeCycleManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDbLifeCycleManagementClientFromBaseClient(baseClient, configProvider)
}

func newDbLifeCycleManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DbLifeCycleManagementClient, err error) {
	// DbLifeCycleManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DbLifeCycleManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DbLifeCycleManagementClient{BaseClient: baseClient}
	client.BasePath = "20240102"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DbLifeCycleManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dblm", "https://dbmgmt.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DbLifeCycleManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DbLifeCycleManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AbortPatchOperation Abort patch tasks for specified PatchOperationId or SubscribeDatabases.
// A default retry strategy applies to this operation AbortPatchOperation()
func (client DbLifeCycleManagementClient) AbortPatchOperation(ctx context.Context, request AbortPatchOperationRequest) (response AbortPatchOperationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.abortPatchOperation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AbortPatchOperationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AbortPatchOperationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AbortPatchOperationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AbortPatchOperationResponse")
	}
	return
}

// abortPatchOperation implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) abortPatchOperation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/patchOperations/{patchOperationId}/actions/abort", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AbortPatchOperationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "AbortPatchOperation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePatchOperationCompartment Moves a PatchOperation resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
// A default retry strategy applies to this operation ChangePatchOperationCompartment()
func (client DbLifeCycleManagementClient) ChangePatchOperationCompartment(ctx context.Context, request ChangePatchOperationCompartmentRequest) (response ChangePatchOperationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePatchOperationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePatchOperationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePatchOperationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePatchOperationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePatchOperationCompartmentResponse")
	}
	return
}

// changePatchOperationCompartment implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) changePatchOperationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/patchOperations/{patchOperationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePatchOperationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ChangePatchOperationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSoftwareImage Creates a new SoftwareImage.
// A default retry strategy applies to this operation CreateSoftwareImage()
func (client DbLifeCycleManagementClient) CreateSoftwareImage(ctx context.Context, request CreateSoftwareImageRequest) (response CreateSoftwareImageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSoftwareImageResponse")
	}
	return
}

// createSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) createSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareImages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "CreateSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSoftwareImageVersion Creates a SoftwareImageVersion for an existing SoftwareImage.
// A default retry strategy applies to this operation CreateSoftwareImageVersion()
func (client DbLifeCycleManagementClient) CreateSoftwareImageVersion(ctx context.Context, request CreateSoftwareImageVersionRequest) (response CreateSoftwareImageVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSoftwareImageVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSoftwareImageVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSoftwareImageVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSoftwareImageVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSoftwareImageVersionResponse")
	}
	return
}

// createSoftwareImageVersion implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) createSoftwareImageVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareImages/{softwareImageId}/softwareImageVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSoftwareImageVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "CreateSoftwareImageVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVulnerabilityScan Creates a VulnerabilityScan.
// A default retry strategy applies to this operation CreateVulnerabilityScan()
func (client DbLifeCycleManagementClient) CreateVulnerabilityScan(ctx context.Context, request CreateVulnerabilityScanRequest) (response CreateVulnerabilityScanResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVulnerabilityScan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVulnerabilityScanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVulnerabilityScanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVulnerabilityScanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVulnerabilityScanResponse")
	}
	return
}

// createVulnerabilityScan implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) createVulnerabilityScan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vulnerabilityScans", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVulnerabilityScanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "CreateVulnerabilityScan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePatchOperation Deletes a PatchOperation resource by identifier
// A default retry strategy applies to this operation DeletePatchOperation()
func (client DbLifeCycleManagementClient) DeletePatchOperation(ctx context.Context, request DeletePatchOperationRequest) (response DeletePatchOperationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePatchOperation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePatchOperationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePatchOperationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePatchOperationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePatchOperationResponse")
	}
	return
}

// deletePatchOperation implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) deletePatchOperation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/patchOperations/{patchOperationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePatchOperationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "DeletePatchOperation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSoftwareImage Deletes a SoftwareImage resource by identifier
// A default retry strategy applies to this operation DeleteSoftwareImage()
func (client DbLifeCycleManagementClient) DeleteSoftwareImage(ctx context.Context, request DeleteSoftwareImageRequest) (response DeleteSoftwareImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSoftwareImageResponse")
	}
	return
}

// deleteSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) deleteSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/softwareImages/{softwareImageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "DeleteSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSoftwareImageVersion Deletes a SoftwareImageVersion by softwareImageVersionKey
// A default retry strategy applies to this operation DeleteSoftwareImageVersion()
func (client DbLifeCycleManagementClient) DeleteSoftwareImageVersion(ctx context.Context, request DeleteSoftwareImageVersionRequest) (response DeleteSoftwareImageVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSoftwareImageVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSoftwareImageVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSoftwareImageVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSoftwareImageVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSoftwareImageVersionResponse")
	}
	return
}

// deleteSoftwareImageVersion implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) deleteSoftwareImageVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/softwareImages/{softwareImageId}/softwareImageVersions/{softwareImageVersionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSoftwareImageVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "DeleteSoftwareImageVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPatchManagement Overview of Patch Management.
// A default retry strategy applies to this operation GetPatchManagement()
func (client DbLifeCycleManagementClient) GetPatchManagement(ctx context.Context, request GetPatchManagementRequest) (response GetPatchManagementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPatchManagement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPatchManagementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPatchManagementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPatchManagementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPatchManagementResponse")
	}
	return
}

// getPatchManagement implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getPatchManagement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patchManagement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPatchManagementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetPatchManagement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPatchOperation Gets a PatchOperation by identifier
// A default retry strategy applies to this operation GetPatchOperation()
func (client DbLifeCycleManagementClient) GetPatchOperation(ctx context.Context, request GetPatchOperationRequest) (response GetPatchOperationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPatchOperation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPatchOperationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPatchOperationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPatchOperationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPatchOperationResponse")
	}
	return
}

// getPatchOperation implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getPatchOperation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patchOperations/{patchOperationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPatchOperationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetPatchOperation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPatchTask Gets a PatchTask by identifier
// A default retry strategy applies to this operation GetPatchTask()
func (client DbLifeCycleManagementClient) GetPatchTask(ctx context.Context, request GetPatchTaskRequest) (response GetPatchTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPatchTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPatchTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPatchTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPatchTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPatchTaskResponse")
	}
	return
}

// getPatchTask implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getPatchTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patchOperations/{patchOperationId}/patchTasks/{patchTaskKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPatchTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetPatchTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPatchTaskStep Gets a PatchTaskStep by identifier
// A default retry strategy applies to this operation GetPatchTaskStep()
func (client DbLifeCycleManagementClient) GetPatchTaskStep(ctx context.Context, request GetPatchTaskStepRequest) (response GetPatchTaskStepResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPatchTaskStep, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPatchTaskStepResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPatchTaskStepResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPatchTaskStepResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPatchTaskStepResponse")
	}
	return
}

// getPatchTaskStep implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getPatchTaskStep(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patchOperations/{patchOperationId}/patchTasks/{patchTaskKey}/patchTaskSteps/{patchTaskStepKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPatchTaskStepResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetPatchTaskStep", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSoftwareImage Gets a SoftwareImage by identifier
// A default retry strategy applies to this operation GetSoftwareImage()
func (client DbLifeCycleManagementClient) GetSoftwareImage(ctx context.Context, request GetSoftwareImageRequest) (response GetSoftwareImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSoftwareImageResponse")
	}
	return
}

// getSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareImages/{softwareImageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVulnerability Gets a Vulnerability
// A default retry strategy applies to this operation GetVulnerability()
func (client DbLifeCycleManagementClient) GetVulnerability(ctx context.Context, request GetVulnerabilityRequest) (response GetVulnerabilityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVulnerability, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVulnerabilityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVulnerabilityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVulnerabilityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVulnerabilityResponse")
	}
	return
}

// getVulnerability implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getVulnerability(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerability", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVulnerabilityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetVulnerability", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVulnerabilityScan Gets information about a VulnerabilityScan.
// A default retry strategy applies to this operation GetVulnerabilityScan()
func (client DbLifeCycleManagementClient) GetVulnerabilityScan(ctx context.Context, request GetVulnerabilityScanRequest) (response GetVulnerabilityScanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVulnerabilityScan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVulnerabilityScanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVulnerabilityScanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVulnerabilityScanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVulnerabilityScanResponse")
	}
	return
}

// getVulnerabilityScan implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getVulnerabilityScan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerabilityScans/{vulnerabilityScanId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVulnerabilityScanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetVulnerabilityScan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets details of the work request with the given ID.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DbLifeCycleManagementClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAggregatedVulnerabilityData Gets an AggregatedVulnerabilityData
// A default retry strategy applies to this operation ListAggregatedVulnerabilityData()
func (client DbLifeCycleManagementClient) ListAggregatedVulnerabilityData(ctx context.Context, request ListAggregatedVulnerabilityDataRequest) (response ListAggregatedVulnerabilityDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAggregatedVulnerabilityData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAggregatedVulnerabilityDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAggregatedVulnerabilityDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAggregatedVulnerabilityDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAggregatedVulnerabilityDataResponse")
	}
	return
}

// listAggregatedVulnerabilityData implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listAggregatedVulnerabilityData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerability/aggregatedVulnerabilityData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAggregatedVulnerabilityDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListAggregatedVulnerabilityData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabases Gets the list of databases
// A default retry strategy applies to this operation ListDatabases()
func (client DbLifeCycleManagementClient) ListDatabases(ctx context.Context, request ListDatabasesRequest) (response ListDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabasesResponse")
	}
	return
}

// listDatabases implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dblm/patchManagement/databases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNotifications List of notifications
// A default retry strategy applies to this operation ListNotifications()
func (client DbLifeCycleManagementClient) ListNotifications(ctx context.Context, request ListNotificationsRequest) (response ListNotificationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNotifications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNotificationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNotificationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNotificationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNotificationsResponse")
	}
	return
}

// listNotifications implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listNotifications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerability/notifications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNotificationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListNotifications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatchOperations Returns a list of PatchOperations.
// A default retry strategy applies to this operation ListPatchOperations()
func (client DbLifeCycleManagementClient) ListPatchOperations(ctx context.Context, request ListPatchOperationsRequest) (response ListPatchOperationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatchOperations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchOperationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchOperationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchOperationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchOperationsResponse")
	}
	return
}

// listPatchOperations implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listPatchOperations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patchOperations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchOperationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListPatchOperations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatchTaskSteps Returns a list of PatchTaskSteps.
// A default retry strategy applies to this operation ListPatchTaskSteps()
func (client DbLifeCycleManagementClient) ListPatchTaskSteps(ctx context.Context, request ListPatchTaskStepsRequest) (response ListPatchTaskStepsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatchTaskSteps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchTaskStepsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchTaskStepsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchTaskStepsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchTaskStepsResponse")
	}
	return
}

// listPatchTaskSteps implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listPatchTaskSteps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patchOperations/{patchOperationId}/patchTasks/{patchTaskKey}/patchTaskSteps", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchTaskStepsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListPatchTaskSteps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatchTasks Returns a list of PatchTasks.
// A default retry strategy applies to this operation ListPatchTasks()
func (client DbLifeCycleManagementClient) ListPatchTasks(ctx context.Context, request ListPatchTasksRequest) (response ListPatchTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatchTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatchTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatchTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatchTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatchTasksResponse")
	}
	return
}

// listPatchTasks implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listPatchTasks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/patchOperations/{patchOperationId}/patchTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatchTasksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListPatchTasks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSoftwareImages Returns a list of SoftwareImages.
// A default retry strategy applies to this operation ListSoftwareImages()
func (client DbLifeCycleManagementClient) ListSoftwareImages(ctx context.Context, request ListSoftwareImagesRequest) (response ListSoftwareImagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSoftwareImages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSoftwareImagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSoftwareImagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSoftwareImagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSoftwareImagesResponse")
	}
	return
}

// listSoftwareImages implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listSoftwareImages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/softwareImages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSoftwareImagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListSoftwareImages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVulnerabilities Gets the vulnerabilities summary list
// A default retry strategy applies to this operation ListVulnerabilities()
func (client DbLifeCycleManagementClient) ListVulnerabilities(ctx context.Context, request ListVulnerabilitiesRequest) (response ListVulnerabilitiesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listVulnerabilities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVulnerabilitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVulnerabilitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVulnerabilitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVulnerabilitiesResponse")
	}
	return
}

// listVulnerabilities implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listVulnerabilities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerability/vulnerabilities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVulnerabilitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListVulnerabilities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVulnerabilityResources Lists the summary of vulnerable and clean resourcees
// A default retry strategy applies to this operation ListVulnerabilityResources()
func (client DbLifeCycleManagementClient) ListVulnerabilityResources(ctx context.Context, request ListVulnerabilityResourcesRequest) (response ListVulnerabilityResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVulnerabilityResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVulnerabilityResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVulnerabilityResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVulnerabilityResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVulnerabilityResourcesResponse")
	}
	return
}

// listVulnerabilityResources implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listVulnerabilityResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerability/resources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVulnerabilityResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListVulnerabilityResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVulnerabilityScans Gets a list of VulnerabilityScans.
// A default retry strategy applies to this operation ListVulnerabilityScans()
func (client DbLifeCycleManagementClient) ListVulnerabilityScans(ctx context.Context, request ListVulnerabilityScansRequest) (response ListVulnerabilityScansResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVulnerabilityScans, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVulnerabilityScansResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVulnerabilityScansResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVulnerabilityScansResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVulnerabilityScansResponse")
	}
	return
}

// listVulnerabilityScans implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listVulnerabilityScans(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vulnerabilityScans", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVulnerabilityScansResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListVulnerabilityScans", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a (paginated) list of errors for the work request with the given ID.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DbLifeCycleManagementClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a (paginated) list of logs for the work request with the given ID.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DbLifeCycleManagementClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DbLifeCycleManagementClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchResources Creates a new PatchOperation to patch resources.
// A default retry strategy applies to this operation PatchResources()
func (client DbLifeCycleManagementClient) PatchResources(ctx context.Context, request PatchResourcesRequest) (response PatchResourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.patchResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchResourcesResponse")
	}
	return
}

// patchResources implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) patchResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/patchOperations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "PatchResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetryPatchOperation Retries existing PatchOperation to patch resources.
// A default retry strategy applies to this operation RetryPatchOperation()
func (client DbLifeCycleManagementClient) RetryPatchOperation(ctx context.Context, request RetryPatchOperationRequest) (response RetryPatchOperationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retryPatchOperation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetryPatchOperationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetryPatchOperationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetryPatchOperationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetryPatchOperationResponse")
	}
	return
}

// retryPatchOperation implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) retryPatchOperation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/patchOperations/{patchOperationId}/actions/retry", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetryPatchOperationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "RetryPatchOperation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Subscribe List of subscriber database objects are subscribed to a Software Image.
// A default retry strategy applies to this operation Subscribe()
func (client DbLifeCycleManagementClient) Subscribe(ctx context.Context, request SubscribeRequest) (response SubscribeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.subscribe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SubscribeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SubscribeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SubscribeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SubscribeResponse")
	}
	return
}

// subscribe implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) subscribe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareImages/{softwareImageId}/actions/subscribe", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SubscribeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "Subscribe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Unsubscribe Unsubscribe a list of subscriber database objects that are subscribed to a Software Image.
// A default retry strategy applies to this operation Unsubscribe()
func (client DbLifeCycleManagementClient) Unsubscribe(ctx context.Context, request UnsubscribeRequest) (response UnsubscribeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.unsubscribe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UnsubscribeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UnsubscribeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UnsubscribeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UnsubscribeResponse")
	}
	return
}

// unsubscribe implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) unsubscribe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/softwareImages/{softwareImageId}/actions/unsubscribe", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UnsubscribeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "Unsubscribe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePatchOperation Updates the PatchOperation
// A default retry strategy applies to this operation UpdatePatchOperation()
func (client DbLifeCycleManagementClient) UpdatePatchOperation(ctx context.Context, request UpdatePatchOperationRequest) (response UpdatePatchOperationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePatchOperation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePatchOperationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePatchOperationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePatchOperationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePatchOperationResponse")
	}
	return
}

// updatePatchOperation implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) updatePatchOperation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/patchOperations/{patchOperationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePatchOperationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "UpdatePatchOperation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSoftwareImage Updates the SoftwareImage
// A default retry strategy applies to this operation UpdateSoftwareImage()
func (client DbLifeCycleManagementClient) UpdateSoftwareImage(ctx context.Context, request UpdateSoftwareImageRequest) (response UpdateSoftwareImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSoftwareImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSoftwareImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSoftwareImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSoftwareImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSoftwareImageResponse")
	}
	return
}

// updateSoftwareImage implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) updateSoftwareImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/softwareImages/{softwareImageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSoftwareImageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "UpdateSoftwareImage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSoftwareImageVersion Updates the SoftwareImage
// A default retry strategy applies to this operation UpdateSoftwareImageVersion()
func (client DbLifeCycleManagementClient) UpdateSoftwareImageVersion(ctx context.Context, request UpdateSoftwareImageVersionRequest) (response UpdateSoftwareImageVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSoftwareImageVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSoftwareImageVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSoftwareImageVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSoftwareImageVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSoftwareImageVersionResponse")
	}
	return
}

// updateSoftwareImageVersion implements the OCIOperation interface (enables retrying operations)
func (client DbLifeCycleManagementClient) updateSoftwareImageVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/softwareImages/{softwareImageId}/softwareImageVersions/{softwareImageVersionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSoftwareImageVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DbLifeCycleManagement", "UpdateSoftwareImageVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
