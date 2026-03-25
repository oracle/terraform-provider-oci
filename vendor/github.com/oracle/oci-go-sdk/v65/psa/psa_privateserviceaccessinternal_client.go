// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage Private Service Access (PSA) endpoints. PSA endpoints are used to create private access between resources in a VCN or on-premises and services in Oracle services network. For important details about how PSA endpoints work, see Access to Oracle Services: Private Service Access Endpoints (https://docs.oracle.com/iaas/Content/Network/Concepts/private-service-access.htm).
//

package psa

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// PrivateServiceAccessInternalClient a client for PrivateServiceAccessInternal
type PrivateServiceAccessInternalClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPrivateServiceAccessInternalClientWithConfigurationProvider Creates a new default PrivateServiceAccessInternal client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPrivateServiceAccessInternalClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PrivateServiceAccessInternalClient, err error) {
	if enabled := common.CheckForEnabledServices("psa"); !enabled {
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
	return newPrivateServiceAccessInternalClientFromBaseClient(baseClient, provider)
}

// NewPrivateServiceAccessInternalClientWithOboToken Creates a new default PrivateServiceAccessInternal client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewPrivateServiceAccessInternalClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PrivateServiceAccessInternalClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPrivateServiceAccessInternalClientFromBaseClient(baseClient, configProvider)
}

func newPrivateServiceAccessInternalClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PrivateServiceAccessInternalClient, err error) {
	// PrivateServiceAccessInternal service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("PrivateServiceAccessInternal"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PrivateServiceAccessInternalClient{BaseClient: baseClient}
	client.BasePath = "20240301"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PrivateServiceAccessInternalClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("psa", "https://psasvc.{region}.oci.{secondLevelDomain}", "psasvc")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PrivateServiceAccessInternalClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PrivateServiceAccessInternalClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangePsaEndpointServiceCompartment Moves a Psa Endpoint Service into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation ChangePsaEndpointServiceCompartment()
func (client PrivateServiceAccessInternalClient) ChangePsaEndpointServiceCompartment(ctx context.Context, request ChangePsaEndpointServiceCompartmentRequest) (response ChangePsaEndpointServiceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePsaEndpointServiceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePsaEndpointServiceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePsaEndpointServiceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePsaEndpointServiceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePsaEndpointServiceCompartmentResponse")
	}
	return
}

// changePsaEndpointServiceCompartment implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessInternalClient) changePsaEndpointServiceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/psaEndpointServices/{psaEndpointServiceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePsaEndpointServiceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "privateServiceAccessInternal", "ChangePsaEndpointServiceCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/psasvc/20240301/PsaEndpointService/ChangePsaEndpointServiceCompartment"
		err = common.PostProcessServiceError(err, "PrivateServiceAccessInternal", "ChangePsaEndpointServiceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePsaEndpointService Creates a psa endpoint service resource in the compartment for a particular service.
// A default retry strategy applies to this operation CreatePsaEndpointService()
func (client PrivateServiceAccessInternalClient) CreatePsaEndpointService(ctx context.Context, request CreatePsaEndpointServiceRequest) (response CreatePsaEndpointServiceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPsaEndpointService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePsaEndpointServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePsaEndpointServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePsaEndpointServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePsaEndpointServiceResponse")
	}
	return
}

// createPsaEndpointService implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessInternalClient) createPsaEndpointService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/psaEndpointServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePsaEndpointServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "privateServiceAccessInternal", "CreatePsaEndpointService")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/psasvc/20240301/PsaEndpointService/CreatePsaEndpointService"
		err = common.PostProcessServiceError(err, "PrivateServiceAccessInternal", "CreatePsaEndpointService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePsaEndpointService Deletes a PsaEndpointService.
// A default retry strategy applies to this operation DeletePsaEndpointService()
func (client PrivateServiceAccessInternalClient) DeletePsaEndpointService(ctx context.Context, request DeletePsaEndpointServiceRequest) (response DeletePsaEndpointServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePsaEndpointService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePsaEndpointServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePsaEndpointServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePsaEndpointServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePsaEndpointServiceResponse")
	}
	return
}

// deletePsaEndpointService implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessInternalClient) deletePsaEndpointService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/psaEndpointServices/{psaEndpointServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePsaEndpointServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "privateServiceAccessInternal", "DeletePsaEndpointService")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/psasvc/20240301/PsaEndpointService/DeletePsaEndpointService"
		err = common.PostProcessServiceError(err, "PrivateServiceAccessInternal", "DeletePsaEndpointService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPsaEndpointService Gets information about a PsaEndPointService Resource.
// A default retry strategy applies to this operation GetPsaEndpointService()
func (client PrivateServiceAccessInternalClient) GetPsaEndpointService(ctx context.Context, request GetPsaEndpointServiceRequest) (response GetPsaEndpointServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPsaEndpointService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPsaEndpointServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPsaEndpointServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPsaEndpointServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPsaEndpointServiceResponse")
	}
	return
}

// getPsaEndpointService implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessInternalClient) getPsaEndpointService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaEndpointServices/{psaEndpointServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPsaEndpointServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "privateServiceAccessInternal", "GetPsaEndpointService")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/psasvc/20240301/PsaEndpointService/GetPsaEndpointService"
		err = common.PostProcessServiceError(err, "PrivateServiceAccessInternal", "GetPsaEndpointService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPsaEndpointServices List the psa endpoint services.
// A default retry strategy applies to this operation ListPsaEndpointServices()
func (client PrivateServiceAccessInternalClient) ListPsaEndpointServices(ctx context.Context, request ListPsaEndpointServicesRequest) (response ListPsaEndpointServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPsaEndpointServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPsaEndpointServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPsaEndpointServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPsaEndpointServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPsaEndpointServicesResponse")
	}
	return
}

// listPsaEndpointServices implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessInternalClient) listPsaEndpointServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaEndpointServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPsaEndpointServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "privateServiceAccessInternal", "ListPsaEndpointServices")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/psasvc/20240301/PsaEndpointServiceCollection/ListPsaEndpointServices"
		err = common.PostProcessServiceError(err, "PrivateServiceAccessInternal", "ListPsaEndpointServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPsaInternalData List the PSA internal only data for a VCN, sorted by service id.
// A default retry strategy applies to this operation ListPsaInternalData()
func (client PrivateServiceAccessInternalClient) ListPsaInternalData(ctx context.Context, request ListPsaInternalDataRequest) (response ListPsaInternalDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPsaInternalData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPsaInternalDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPsaInternalDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPsaInternalDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPsaInternalDataResponse")
	}
	return
}

// listPsaInternalData implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessInternalClient) listPsaInternalData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaInternalData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPsaInternalDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "privateServiceAccessInternal", "ListPsaInternalData")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/psasvc/20240301/PsaInternalDataCollection/ListPsaInternalData"
		err = common.PostProcessServiceError(err, "PrivateServiceAccessInternal", "ListPsaInternalData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePsaEndpointService Updates a PsaEndpointService.
// A default retry strategy applies to this operation UpdatePsaEndpointService()
func (client PrivateServiceAccessInternalClient) UpdatePsaEndpointService(ctx context.Context, request UpdatePsaEndpointServiceRequest) (response UpdatePsaEndpointServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePsaEndpointService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePsaEndpointServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePsaEndpointServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePsaEndpointServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePsaEndpointServiceResponse")
	}
	return
}

// updatePsaEndpointService implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessInternalClient) updatePsaEndpointService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/psaEndpointServices/{psaEndpointServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePsaEndpointServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "privateServiceAccessInternal", "UpdatePsaEndpointService")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/psasvc/20240301/PsaEndpointService/UpdatePsaEndpointService"
		err = common.PostProcessServiceError(err, "PrivateServiceAccessInternal", "UpdatePsaEndpointService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
