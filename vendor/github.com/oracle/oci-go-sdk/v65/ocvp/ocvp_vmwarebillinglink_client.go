// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// VmwareBillingLinkClient a client for VmwareBillingLink
type VmwareBillingLinkClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewVmwareBillingLinkClientWithConfigurationProvider Creates a new default VmwareBillingLink client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewVmwareBillingLinkClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client VmwareBillingLinkClient, err error) {
	if enabled := common.CheckForEnabledServices("ocvp"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newVmwareBillingLinkClientFromBaseClient(baseClient, provider)
}

// NewVmwareBillingLinkClientWithOboToken Creates a new default VmwareBillingLink client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewVmwareBillingLinkClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client VmwareBillingLinkClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newVmwareBillingLinkClientFromBaseClient(baseClient, configProvider)
}

func newVmwareBillingLinkClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client VmwareBillingLinkClient, err error) {
	// VmwareBillingLink service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("VmwareBillingLink"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = VmwareBillingLinkClient{BaseClient: baseClient}
	client.BasePath = "20230701"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *VmwareBillingLinkClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *VmwareBillingLinkClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *VmwareBillingLinkClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddVmwareSubscriptions Add a subscription to a VMware billing link.
// A default retry strategy applies to this operation AddVmwareSubscriptions()
func (client VmwareBillingLinkClient) AddVmwareSubscriptions(ctx context.Context, request AddVmwareSubscriptionsRequest) (response AddVmwareSubscriptionsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addVmwareSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddVmwareSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddVmwareSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddVmwareSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddVmwareSubscriptionsResponse")
	}
	return
}

// addVmwareSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) addVmwareSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmwareBillingLinks/{vmwareBillingLinkId}/actions/addVmwareSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddVmwareSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLink/AddVmwareSubscriptions"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "AddVmwareSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVmwareBillingLinkCompartment Moves a VMware billing link into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation ChangeVmwareBillingLinkCompartment()
func (client VmwareBillingLinkClient) ChangeVmwareBillingLinkCompartment(ctx context.Context, request ChangeVmwareBillingLinkCompartmentRequest) (response ChangeVmwareBillingLinkCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVmwareBillingLinkCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVmwareBillingLinkCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVmwareBillingLinkCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVmwareBillingLinkCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVmwareBillingLinkCompartmentResponse")
	}
	return
}

// changeVmwareBillingLinkCompartment implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) changeVmwareBillingLinkCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmwareBillingLinks/{vmwareBillingLinkId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVmwareBillingLinkCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLink/ChangeVmwareBillingLinkCompartment"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "ChangeVmwareBillingLinkCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVmwareBillingLink Creates a billing link between a VMware cloud account and OCI tenancy.
// A default retry strategy applies to this operation CreateVmwareBillingLink()
func (client VmwareBillingLinkClient) CreateVmwareBillingLink(ctx context.Context, request CreateVmwareBillingLinkRequest) (response CreateVmwareBillingLinkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVmwareBillingLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVmwareBillingLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVmwareBillingLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVmwareBillingLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVmwareBillingLinkResponse")
	}
	return
}

// createVmwareBillingLink implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) createVmwareBillingLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmwareBillingLinks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVmwareBillingLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLink/CreateVmwareBillingLink"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "CreateVmwareBillingLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVmwareBillingLink Deletes the specified VMware billing link.
// A default retry strategy applies to this operation DeleteVmwareBillingLink()
func (client VmwareBillingLinkClient) DeleteVmwareBillingLink(ctx context.Context, request DeleteVmwareBillingLinkRequest) (response DeleteVmwareBillingLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVmwareBillingLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVmwareBillingLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVmwareBillingLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVmwareBillingLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVmwareBillingLinkResponse")
	}
	return
}

// deleteVmwareBillingLink implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) deleteVmwareBillingLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/vmwareBillingLinks/{vmwareBillingLinkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVmwareBillingLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLink/DeleteVmwareBillingLink"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "DeleteVmwareBillingLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmwareBillingLink Gets the specified VmwareBillingLink's information.
// A default retry strategy applies to this operation GetVmwareBillingLink()
func (client VmwareBillingLinkClient) GetVmwareBillingLink(ctx context.Context, request GetVmwareBillingLinkRequest) (response GetVmwareBillingLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmwareBillingLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmwareBillingLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmwareBillingLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmwareBillingLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmwareBillingLinkResponse")
	}
	return
}

// getVmwareBillingLink implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) getVmwareBillingLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmwareBillingLinks/{vmwareBillingLinkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmwareBillingLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLink/GetVmwareBillingLink"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "GetVmwareBillingLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmwareBillingLinks Lists the VMware billing links in the specified compartment. The list can be
// filtered by display name or availability domain.
// A default retry strategy applies to this operation ListVmwareBillingLinks()
func (client VmwareBillingLinkClient) ListVmwareBillingLinks(ctx context.Context, request ListVmwareBillingLinksRequest) (response ListVmwareBillingLinksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmwareBillingLinks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmwareBillingLinksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmwareBillingLinksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmwareBillingLinksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmwareBillingLinksResponse")
	}
	return
}

// listVmwareBillingLinks implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) listVmwareBillingLinks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmwareBillingLinks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmwareBillingLinksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLinkSummary/ListVmwareBillingLinks"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "ListVmwareBillingLinks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RenewVmwareSubscription Renew a subscription
// A default retry strategy applies to this operation RenewVmwareSubscription()
func (client VmwareBillingLinkClient) RenewVmwareSubscription(ctx context.Context, request RenewVmwareSubscriptionRequest) (response RenewVmwareSubscriptionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.renewVmwareSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RenewVmwareSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RenewVmwareSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RenewVmwareSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RenewVmwareSubscriptionResponse")
	}
	return
}

// renewVmwareSubscription implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) renewVmwareSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmwareBillingLinks/{vmwareBillingLinkId}/actions/renewVmwareSubscription", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RenewVmwareSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLink/RenewVmwareSubscription"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "RenewVmwareSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVmwareBillingLink Updates the specified VMware billing link.
// UpdateVmwareBillingLinkDetails.
// A default retry strategy applies to this operation UpdateVmwareBillingLink()
func (client VmwareBillingLinkClient) UpdateVmwareBillingLink(ctx context.Context, request UpdateVmwareBillingLinkRequest) (response UpdateVmwareBillingLinkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVmwareBillingLink, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVmwareBillingLinkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVmwareBillingLinkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVmwareBillingLinkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVmwareBillingLinkResponse")
	}
	return
}

// updateVmwareBillingLink implements the OCIOperation interface (enables retrying operations)
func (client VmwareBillingLinkClient) updateVmwareBillingLink(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/vmwareBillingLinks/{vmwareBillingLinkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVmwareBillingLinkResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/VmwareBillingLink/UpdateVmwareBillingLink"
		err = common.PostProcessServiceError(err, "VmwareBillingLink", "UpdateVmwareBillingLink", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
