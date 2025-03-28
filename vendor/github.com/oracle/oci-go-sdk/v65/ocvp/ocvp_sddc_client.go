// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// SddcClient a client for Sddc
type SddcClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSddcClientWithConfigurationProvider Creates a new default Sddc client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSddcClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SddcClient, err error) {
	if enabled := common.CheckForEnabledServices("ocvp"); !enabled {
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
	return newSddcClientFromBaseClient(baseClient, provider)
}

// NewSddcClientWithOboToken Creates a new default Sddc client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewSddcClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SddcClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSddcClientFromBaseClient(baseClient, configProvider)
}

func newSddcClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SddcClient, err error) {
	// Sddc service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Sddc"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SddcClient{BaseClient: baseClient}
	client.BasePath = "20230701"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SddcClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SddcClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SddcClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelDowngradeHcx Cancel the pending SDDC downgrade from HCX Enterprise to HCX Advanced.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/CancelDowngradeHcx.go.html to see an example of how to use CancelDowngradeHcx API.
// A default retry strategy applies to this operation CancelDowngradeHcx()
func (client SddcClient) CancelDowngradeHcx(ctx context.Context, request CancelDowngradeHcxRequest) (response CancelDowngradeHcxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelDowngradeHcx, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelDowngradeHcxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelDowngradeHcxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelDowngradeHcxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelDowngradeHcxResponse")
	}
	return
}

// cancelDowngradeHcx implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) cancelDowngradeHcx(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs/{sddcId}/actions/cancelDowngradeHcx", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelDowngradeHcxResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/CancelDowngradeHcx"
		err = common.PostProcessServiceError(err, "Sddc", "CancelDowngradeHcx", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSddcCompartment Moves an SDDC into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ChangeSddcCompartment.go.html to see an example of how to use ChangeSddcCompartment API.
// A default retry strategy applies to this operation ChangeSddcCompartment()
func (client SddcClient) ChangeSddcCompartment(ctx context.Context, request ChangeSddcCompartmentRequest) (response ChangeSddcCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSddcCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSddcCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSddcCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSddcCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSddcCompartmentResponse")
	}
	return
}

// changeSddcCompartment implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) changeSddcCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs/{sddcId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSddcCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/ChangeSddcCompartment"
		err = common.PostProcessServiceError(err, "Sddc", "ChangeSddcCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSddc Creates an Oracle Cloud VMware Solution software-defined data center (SDDC).
// Use the WorkRequest operations to track the
// creation of the SDDC.
// **Important:** You must configure the SDDC's networking resources with the security rules detailed in Security Rules for Oracle Cloud VMware Solution SDDCs (https://docs.oracle.com/iaas/Content/VMware/Reference/ocvssecurityrules.htm). Otherwise, provisioning the SDDC will fail. The rules are based on the requirements set by VMware.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/CreateSddc.go.html to see an example of how to use CreateSddc API.
// A default retry strategy applies to this operation CreateSddc()
func (client SddcClient) CreateSddc(ctx context.Context, request CreateSddcRequest) (response CreateSddcResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSddc, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSddcResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSddcResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSddcResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSddcResponse")
	}
	return
}

// createSddc implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) createSddc(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/CreateSddc"
		err = common.PostProcessServiceError(err, "Sddc", "CreateSddc", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSddc Deletes the specified SDDC, along with the other resources that were
// created with the SDDC. For example: the Compute instances, DNS records,
// and so on.
// Use the WorkRequest operations to track the
// deletion of the SDDC.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/DeleteSddc.go.html to see an example of how to use DeleteSddc API.
// A default retry strategy applies to this operation DeleteSddc()
func (client SddcClient) DeleteSddc(ctx context.Context, request DeleteSddcRequest) (response DeleteSddcResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSddc, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSddcResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSddcResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSddcResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSddcResponse")
	}
	return
}

// deleteSddc implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) deleteSddc(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sddcs/{sddcId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/DeleteSddc"
		err = common.PostProcessServiceError(err, "Sddc", "DeleteSddc", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DowngradeHcx Downgrade the specified SDDC from HCX Enterprise to HCX Advanced.
// SDDC with standard compute shapes will always use HCX Enterprise if HCX is enabled and cannot be downgraded.
// Downgrading from HCX Enterprise to HCX Advanced reduces the number of provided license keys from 10 to 3.
// Downgrade remains in a `PENDING` state until the end of the current billing cycle. You can use CancelDowngradeHcx
// to cancel the downgrade while it's still in a `PENDING` state.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/DowngradeHcx.go.html to see an example of how to use DowngradeHcx API.
// A default retry strategy applies to this operation DowngradeHcx()
func (client SddcClient) DowngradeHcx(ctx context.Context, request DowngradeHcxRequest) (response DowngradeHcxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downgradeHcx, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DowngradeHcxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DowngradeHcxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DowngradeHcxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DowngradeHcxResponse")
	}
	return
}

// downgradeHcx implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) downgradeHcx(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs/{sddcId}/actions/downgradeHcx", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DowngradeHcxResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/DowngradeHcx"
		err = common.PostProcessServiceError(err, "Sddc", "DowngradeHcx", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSddc Gets the specified SDDC's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/GetSddc.go.html to see an example of how to use GetSddc API.
// A default retry strategy applies to this operation GetSddc()
func (client SddcClient) GetSddc(ctx context.Context, request GetSddcRequest) (response GetSddcResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSddc, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSddcResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSddcResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSddcResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSddcResponse")
	}
	return
}

// getSddc implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) getSddc(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sddcs/{sddcId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/GetSddc"
		err = common.PostProcessServiceError(err, "Sddc", "GetSddc", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSddcs Lists the SDDCs in the specified compartment. The list can be
// filtered by display name or availability domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListSddcs.go.html to see an example of how to use ListSddcs API.
// A default retry strategy applies to this operation ListSddcs()
func (client SddcClient) ListSddcs(ctx context.Context, request ListSddcsRequest) (response ListSddcsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSddcs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSddcsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSddcsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSddcsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSddcsResponse")
	}
	return
}

// listSddcs implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) listSddcs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sddcs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSddcsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/SddcSummary/ListSddcs"
		err = common.PostProcessServiceError(err, "Sddc", "ListSddcs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSupportedCommitments Lists supported Commitments.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListSupportedCommitments.go.html to see an example of how to use ListSupportedCommitments API.
// A default retry strategy applies to this operation ListSupportedCommitments()
func (client SddcClient) ListSupportedCommitments(ctx context.Context, request ListSupportedCommitmentsRequest) (response ListSupportedCommitmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSupportedCommitments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSupportedCommitmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSupportedCommitmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSupportedCommitmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSupportedCommitmentsResponse")
	}
	return
}

// listSupportedCommitments implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) listSupportedCommitments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/supportedCommitments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSupportedCommitmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/SupportedCommitmentSummary/ListSupportedCommitments"
		err = common.PostProcessServiceError(err, "Sddc", "ListSupportedCommitments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSupportedHostShapes Lists supported compute shapes for ESXi hosts.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListSupportedHostShapes.go.html to see an example of how to use ListSupportedHostShapes API.
// A default retry strategy applies to this operation ListSupportedHostShapes()
func (client SddcClient) ListSupportedHostShapes(ctx context.Context, request ListSupportedHostShapesRequest) (response ListSupportedHostShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSupportedHostShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSupportedHostShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSupportedHostShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSupportedHostShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSupportedHostShapesResponse")
	}
	return
}

// listSupportedHostShapes implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) listSupportedHostShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/supportedHostShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSupportedHostShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/SupportedHostShapeSummary/ListSupportedHostShapes"
		err = common.PostProcessServiceError(err, "Sddc", "ListSupportedHostShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSupportedVmwareSoftwareVersions Lists the versions of bundled VMware software supported by the Oracle Cloud
// VMware Solution.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListSupportedVmwareSoftwareVersions.go.html to see an example of how to use ListSupportedVmwareSoftwareVersions API.
// A default retry strategy applies to this operation ListSupportedVmwareSoftwareVersions()
func (client SddcClient) ListSupportedVmwareSoftwareVersions(ctx context.Context, request ListSupportedVmwareSoftwareVersionsRequest) (response ListSupportedVmwareSoftwareVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSupportedVmwareSoftwareVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSupportedVmwareSoftwareVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSupportedVmwareSoftwareVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSupportedVmwareSoftwareVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSupportedVmwareSoftwareVersionsResponse")
	}
	return
}

// listSupportedVmwareSoftwareVersions implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) listSupportedVmwareSoftwareVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/supportedVmwareSoftwareVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSupportedVmwareSoftwareVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions"
		err = common.PostProcessServiceError(err, "Sddc", "ListSupportedVmwareSoftwareVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshHcxLicenseStatus Refresh HCX on-premise licenses status of the specified SDDC.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/RefreshHcxLicenseStatus.go.html to see an example of how to use RefreshHcxLicenseStatus API.
// A default retry strategy applies to this operation RefreshHcxLicenseStatus()
func (client SddcClient) RefreshHcxLicenseStatus(ctx context.Context, request RefreshHcxLicenseStatusRequest) (response RefreshHcxLicenseStatusResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.refreshHcxLicenseStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshHcxLicenseStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshHcxLicenseStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshHcxLicenseStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshHcxLicenseStatusResponse")
	}
	return
}

// refreshHcxLicenseStatus implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) refreshHcxLicenseStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs/{sddcId}/actions/refreshHcxLicenses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshHcxLicenseStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/RefreshHcxLicenseStatus"
		err = common.PostProcessServiceError(err, "Sddc", "RefreshHcxLicenseStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrievePassword Retrieve the SDDC password.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/RetrievePassword.go.html to see an example of how to use RetrievePassword API.
// A default retry strategy applies to this operation RetrievePassword()
func (client SddcClient) RetrievePassword(ctx context.Context, request RetrievePasswordRequest) (response RetrievePasswordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.retrievePassword, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrievePasswordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrievePasswordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrievePasswordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrievePasswordResponse")
	}
	return
}

// retrievePassword implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) retrievePassword(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs/{sddcId}/actions/retrievePassword", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrievePasswordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/RetrievePassword"
		err = common.PostProcessServiceError(err, "Sddc", "RetrievePassword", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSddc Updates the specified SDDC.
// **Important:** Updating an SDDC affects only certain attributes in the `Sddc`
// object and does not affect the VMware environment currently running in
// the SDDC. For more information, see
// UpdateSddcDetails.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/UpdateSddc.go.html to see an example of how to use UpdateSddc API.
// A default retry strategy applies to this operation UpdateSddc()
func (client SddcClient) UpdateSddc(ctx context.Context, request UpdateSddcRequest) (response UpdateSddcResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSddc, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSddcResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSddcResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSddcResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSddcResponse")
	}
	return
}

// updateSddc implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) updateSddc(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sddcs/{sddcId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/UpdateSddc"
		err = common.PostProcessServiceError(err, "Sddc", "UpdateSddc", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpgradeHcx Upgrade the specified SDDC from HCX Advanced to HCX Enterprise.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/UpgradeHcx.go.html to see an example of how to use UpgradeHcx API.
// A default retry strategy applies to this operation UpgradeHcx()
func (client SddcClient) UpgradeHcx(ctx context.Context, request UpgradeHcxRequest) (response UpgradeHcxResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.upgradeHcx, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpgradeHcxResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpgradeHcxResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpgradeHcxResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpgradeHcxResponse")
	}
	return
}

// upgradeHcx implements the OCIOperation interface (enables retrying operations)
func (client SddcClient) upgradeHcx(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs/{sddcId}/actions/upgradeHcx", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpgradeHcxResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/UpgradeHcx"
		err = common.PostProcessServiceError(err, "Sddc", "UpgradeHcx", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
