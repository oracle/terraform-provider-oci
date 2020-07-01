// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage the Oracle Cloud VMware Solution.
//

package ocvp

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//SddcClient a client for Sddc
type SddcClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSddcClientWithConfigurationProvider Creates a new default Sddc client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSddcClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SddcClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newSddcClientFromBaseClient(baseClient, configProvider)
}

// NewSddcClientWithOboToken Creates a new default Sddc client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewSddcClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SddcClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newSddcClientFromBaseClient(baseClient, configProvider)
}

func newSddcClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SddcClient, err error) {
	client = SddcClient{BaseClient: baseClient}
	client.BasePath = "20200501"
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *SddcClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeSddcCompartment Moves an SDDC into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client SddcClient) ChangeSddcCompartment(ctx context.Context, request ChangeSddcCompartmentRequest) (response ChangeSddcCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client SddcClient) changeSddcCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs/{sddcId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeSddcCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSddc Creates a software-defined data center (SDDC).
// Use the WorkRequest operations to track the
// creation of the SDDC.
func (client SddcClient) CreateSddc(ctx context.Context, request CreateSddcRequest) (response CreateSddcResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client SddcClient) createSddc(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/sddcs")
	if err != nil {
		return nil, err
	}

	var response CreateSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
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
func (client SddcClient) DeleteSddc(ctx context.Context, request DeleteSddcRequest) (response DeleteSddcResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client SddcClient) deleteSddc(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/sddcs/{sddcId}")
	if err != nil {
		return nil, err
	}

	var response DeleteSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSddc Gets the specified SDDC's information.
func (client SddcClient) GetSddc(ctx context.Context, request GetSddcRequest) (response GetSddcResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client SddcClient) getSddc(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sddcs/{sddcId}")
	if err != nil {
		return nil, err
	}

	var response GetSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSddcs Lists the SDDCs in the specified compartment. The list can be
// filtered by display name or availability domain.
func (client SddcClient) ListSddcs(ctx context.Context, request ListSddcsRequest) (response ListSddcsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client SddcClient) listSddcs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sddcs")
	if err != nil {
		return nil, err
	}

	var response ListSddcsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSupportedVmwareSoftwareVersions Lists the versions of bundled VMware software supported by the Oracle Cloud
// VMware Solution.
func (client SddcClient) ListSupportedVmwareSoftwareVersions(ctx context.Context, request ListSupportedVmwareSoftwareVersionsRequest) (response ListSupportedVmwareSoftwareVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client SddcClient) listSupportedVmwareSoftwareVersions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/supportedVmwareSoftwareVersions")
	if err != nil {
		return nil, err
	}

	var response ListSupportedVmwareSoftwareVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
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
func (client SddcClient) UpdateSddc(ctx context.Context, request UpdateSddcRequest) (response UpdateSddcResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client SddcClient) updateSddc(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/sddcs/{sddcId}")
	if err != nil {
		return nil, err
	}

	var response UpdateSddcResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
