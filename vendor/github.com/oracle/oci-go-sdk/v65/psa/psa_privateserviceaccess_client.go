// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage privateServiceAccess.
//

package psa

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// PrivateServiceAccessClient a client for PrivateServiceAccess
type PrivateServiceAccessClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPrivateServiceAccessClientWithConfigurationProvider Creates a new default PrivateServiceAccess client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPrivateServiceAccessClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PrivateServiceAccessClient, err error) {
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
	return newPrivateServiceAccessClientFromBaseClient(baseClient, provider)
}

// NewPrivateServiceAccessClientWithOboToken Creates a new default PrivateServiceAccess client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewPrivateServiceAccessClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PrivateServiceAccessClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPrivateServiceAccessClientFromBaseClient(baseClient, configProvider)
}

func newPrivateServiceAccessClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PrivateServiceAccessClient, err error) {
	// PrivateServiceAccess service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("PrivateServiceAccess"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PrivateServiceAccessClient{BaseClient: baseClient}
	client.BasePath = "20240301"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PrivateServiceAccessClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("psa", "https://psasvc.{region}.oci.{secondLevelDomain}", "psasvc")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PrivateServiceAccessClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PrivateServiceAccessClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelPsaWorkRequest Cancels a PrivateServiceAccess work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/CancelPsaWorkRequest.go.html to see an example of how to use CancelPsaWorkRequest API.
// A default retry strategy applies to this operation CancelPsaWorkRequest()
func (client PrivateServiceAccessClient) CancelPsaWorkRequest(ctx context.Context, request CancelPsaWorkRequestRequest) (response CancelPsaWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelPsaWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelPsaWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelPsaWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelPsaWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelPsaWorkRequestResponse")
	}
	return
}

// cancelPsaWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) cancelPsaWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/psaWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelPsaWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "CancelPsaWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePrivateServiceAccessCompartment Moves a PrivateServiceAccess into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ChangePrivateServiceAccessCompartment.go.html to see an example of how to use ChangePrivateServiceAccessCompartment API.
// A default retry strategy applies to this operation ChangePrivateServiceAccessCompartment()
func (client PrivateServiceAccessClient) ChangePrivateServiceAccessCompartment(ctx context.Context, request ChangePrivateServiceAccessCompartmentRequest) (response ChangePrivateServiceAccessCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePrivateServiceAccessCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePrivateServiceAccessCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePrivateServiceAccessCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePrivateServiceAccessCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePrivateServiceAccessCompartmentResponse")
	}
	return
}

// changePrivateServiceAccessCompartment implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) changePrivateServiceAccessCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privateServiceAccess/{privateServiceAccessId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePrivateServiceAccessCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "ChangePrivateServiceAccessCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePrivateServiceAccess Creates a private service access in the specified subnet (in the consumer's VCN) and the specified
// compartment for a particular service.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/CreatePrivateServiceAccess.go.html to see an example of how to use CreatePrivateServiceAccess API.
// A default retry strategy applies to this operation CreatePrivateServiceAccess()
func (client PrivateServiceAccessClient) CreatePrivateServiceAccess(ctx context.Context, request CreatePrivateServiceAccessRequest) (response CreatePrivateServiceAccessResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPrivateServiceAccess, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePrivateServiceAccessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePrivateServiceAccessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePrivateServiceAccessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePrivateServiceAccessResponse")
	}
	return
}

// createPrivateServiceAccess implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) createPrivateServiceAccess(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privateServiceAccess", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePrivateServiceAccessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "CreatePrivateServiceAccess", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePrivateServiceAccess Deletes a PrivateServiceAccess.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/DeletePrivateServiceAccess.go.html to see an example of how to use DeletePrivateServiceAccess API.
// A default retry strategy applies to this operation DeletePrivateServiceAccess()
func (client PrivateServiceAccessClient) DeletePrivateServiceAccess(ctx context.Context, request DeletePrivateServiceAccessRequest) (response DeletePrivateServiceAccessResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePrivateServiceAccess, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePrivateServiceAccessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePrivateServiceAccessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePrivateServiceAccessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePrivateServiceAccessResponse")
	}
	return
}

// deletePrivateServiceAccess implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) deletePrivateServiceAccess(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/privateServiceAccess/{privateServiceAccessId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePrivateServiceAccessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "DeletePrivateServiceAccess", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivateServiceAccess Gets information about a PrivateServiceAccess.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/GetPrivateServiceAccess.go.html to see an example of how to use GetPrivateServiceAccess API.
// A default retry strategy applies to this operation GetPrivateServiceAccess()
func (client PrivateServiceAccessClient) GetPrivateServiceAccess(ctx context.Context, request GetPrivateServiceAccessRequest) (response GetPrivateServiceAccessResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateServiceAccess, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateServiceAccessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateServiceAccessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateServiceAccessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateServiceAccessResponse")
	}
	return
}

// getPrivateServiceAccess implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) getPrivateServiceAccess(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateServiceAccess/{privateServiceAccessId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivateServiceAccessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "GetPrivateServiceAccess", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPsaWorkRequest Gets the details of a PrivateServiceAccess work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/GetPsaWorkRequest.go.html to see an example of how to use GetPsaWorkRequest API.
// A default retry strategy applies to this operation GetPsaWorkRequest()
func (client PrivateServiceAccessClient) GetPsaWorkRequest(ctx context.Context, request GetPsaWorkRequestRequest) (response GetPsaWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPsaWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPsaWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPsaWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPsaWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPsaWorkRequestResponse")
	}
	return
}

// getPsaWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) getPsaWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPsaWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "GetPsaWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPrivateServiceAccesses List the private service accesses in the specified compartment. You can optionally filter the list by
// specifying the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a subnet in the cunsumer's VCN.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPrivateServiceAccesses.go.html to see an example of how to use ListPrivateServiceAccesses API.
// A default retry strategy applies to this operation ListPrivateServiceAccesses()
func (client PrivateServiceAccessClient) ListPrivateServiceAccesses(ctx context.Context, request ListPrivateServiceAccessesRequest) (response ListPrivateServiceAccessesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPrivateServiceAccesses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPrivateServiceAccessesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPrivateServiceAccessesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPrivateServiceAccessesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPrivateServiceAccessesResponse")
	}
	return
}

// listPrivateServiceAccesses implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) listPrivateServiceAccesses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateServiceAccess", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPrivateServiceAccessesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "ListPrivateServiceAccesses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPsaServices List the OCI services available for Private Service Access catalog in the region, sorted by service name.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPsaServices.go.html to see an example of how to use ListPsaServices API.
// A default retry strategy applies to this operation ListPsaServices()
func (client PrivateServiceAccessClient) ListPsaServices(ctx context.Context, request ListPsaServicesRequest) (response ListPsaServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPsaServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPsaServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPsaServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPsaServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPsaServicesResponse")
	}
	return
}

// listPsaServices implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) listPsaServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPsaServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "ListPsaServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPsaWorkRequestErrors Lists the errors for a PrivateServiceAccess work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPsaWorkRequestErrors.go.html to see an example of how to use ListPsaWorkRequestErrors API.
// A default retry strategy applies to this operation ListPsaWorkRequestErrors()
func (client PrivateServiceAccessClient) ListPsaWorkRequestErrors(ctx context.Context, request ListPsaWorkRequestErrorsRequest) (response ListPsaWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPsaWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPsaWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPsaWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPsaWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPsaWorkRequestErrorsResponse")
	}
	return
}

// listPsaWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) listPsaWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaWorkRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPsaWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "ListPsaWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPsaWorkRequestLogs Lists the logs for a PrivateServiceAccess work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPsaWorkRequestLogs.go.html to see an example of how to use ListPsaWorkRequestLogs API.
// A default retry strategy applies to this operation ListPsaWorkRequestLogs()
func (client PrivateServiceAccessClient) ListPsaWorkRequestLogs(ctx context.Context, request ListPsaWorkRequestLogsRequest) (response ListPsaWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPsaWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPsaWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPsaWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPsaWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPsaWorkRequestLogsResponse")
	}
	return
}

// listPsaWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) listPsaWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaWorkRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPsaWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "ListPsaWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPsaWorkRequests Lists the PrivateServiceAccess work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/ListPsaWorkRequests.go.html to see an example of how to use ListPsaWorkRequests API.
// A default retry strategy applies to this operation ListPsaWorkRequests()
func (client PrivateServiceAccessClient) ListPsaWorkRequests(ctx context.Context, request ListPsaWorkRequestsRequest) (response ListPsaWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPsaWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPsaWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPsaWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPsaWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPsaWorkRequestsResponse")
	}
	return
}

// listPsaWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) listPsaWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/psaWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPsaWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "ListPsaWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePrivateServiceAccess Updates a PrivateServiceAccess.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psa/UpdatePrivateServiceAccess.go.html to see an example of how to use UpdatePrivateServiceAccess API.
// A default retry strategy applies to this operation UpdatePrivateServiceAccess()
func (client PrivateServiceAccessClient) UpdatePrivateServiceAccess(ctx context.Context, request UpdatePrivateServiceAccessRequest) (response UpdatePrivateServiceAccessResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePrivateServiceAccess, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePrivateServiceAccessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePrivateServiceAccessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePrivateServiceAccessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePrivateServiceAccessResponse")
	}
	return
}

// updatePrivateServiceAccess implements the OCIOperation interface (enables retrying operations)
func (client PrivateServiceAccessClient) updatePrivateServiceAccess(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/privateServiceAccess/{privateServiceAccessId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePrivateServiceAccessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "PrivateServiceAccess", "UpdatePrivateServiceAccess", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
