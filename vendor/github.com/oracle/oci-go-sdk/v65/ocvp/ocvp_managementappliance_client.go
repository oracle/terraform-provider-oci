// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ManagementApplianceClient a client for ManagementAppliance
type ManagementApplianceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagementApplianceClientWithConfigurationProvider Creates a new default ManagementAppliance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagementApplianceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagementApplianceClient, err error) {
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
	return newManagementApplianceClientFromBaseClient(baseClient, provider)
}

// NewManagementApplianceClientWithOboToken Creates a new default ManagementAppliance client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewManagementApplianceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ManagementApplianceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newManagementApplianceClientFromBaseClient(baseClient, configProvider)
}

func newManagementApplianceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ManagementApplianceClient, err error) {
	// ManagementAppliance service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ManagementAppliance"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ManagementApplianceClient{BaseClient: baseClient}
	client.BasePath = "20230701"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ManagementApplianceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ManagementApplianceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ManagementApplianceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateManagementAppliance Creates a management appliance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/CreateManagementAppliance.go.html to see an example of how to use CreateManagementAppliance API.
// A default retry strategy applies to this operation CreateManagementAppliance()
func (client ManagementApplianceClient) CreateManagementAppliance(ctx context.Context, request CreateManagementApplianceRequest) (response CreateManagementApplianceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createManagementAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagementApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagementApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagementApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagementApplianceResponse")
	}
	return
}

// createManagementAppliance implements the OCIOperation interface (enables retrying operations)
func (client ManagementApplianceClient) createManagementAppliance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementAppliances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateManagementApplianceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ManagementAppliance/CreateManagementAppliance"
		err = common.PostProcessServiceError(err, "ManagementAppliance", "CreateManagementAppliance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagementAppliance Deletes management appliance specified.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/DeleteManagementAppliance.go.html to see an example of how to use DeleteManagementAppliance API.
// A default retry strategy applies to this operation DeleteManagementAppliance()
func (client ManagementApplianceClient) DeleteManagementAppliance(ctx context.Context, request DeleteManagementApplianceRequest) (response DeleteManagementApplianceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementApplianceResponse")
	}
	return
}

// deleteManagementAppliance implements the OCIOperation interface (enables retrying operations)
func (client ManagementApplianceClient) deleteManagementAppliance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementAppliances/{managementApplianceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteManagementApplianceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ManagementAppliance/DeleteManagementAppliance"
		err = common.PostProcessServiceError(err, "ManagementAppliance", "DeleteManagementAppliance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagementAppliance Get the specified management appliance information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/GetManagementAppliance.go.html to see an example of how to use GetManagementAppliance API.
// A default retry strategy applies to this operation GetManagementAppliance()
func (client ManagementApplianceClient) GetManagementAppliance(ctx context.Context, request GetManagementApplianceRequest) (response GetManagementApplianceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagementAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementApplianceResponse")
	}
	return
}

// getManagementAppliance implements the OCIOperation interface (enables retrying operations)
func (client ManagementApplianceClient) getManagementAppliance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAppliances/{managementApplianceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagementApplianceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ManagementAppliance/GetManagementAppliance"
		err = common.PostProcessServiceError(err, "ManagementAppliance", "GetManagementAppliance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagementAppliances Lists management appliances in compartment specified. List can be filtered by management appliance, compartment, name and lifecycle state.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListManagementAppliances.go.html to see an example of how to use ListManagementAppliances API.
// A default retry strategy applies to this operation ListManagementAppliances()
func (client ManagementApplianceClient) ListManagementAppliances(ctx context.Context, request ListManagementAppliancesRequest) (response ListManagementAppliancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementAppliances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementAppliancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementAppliancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementAppliancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementAppliancesResponse")
	}
	return
}

// listManagementAppliances implements the OCIOperation interface (enables retrying operations)
func (client ManagementApplianceClient) listManagementAppliances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementAppliances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagementAppliancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ManagementAppliance/ListManagementAppliances"
		err = common.PostProcessServiceError(err, "ManagementAppliance", "ListManagementAppliances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagementAppliance Updates management appliance specified.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/UpdateManagementAppliance.go.html to see an example of how to use UpdateManagementAppliance API.
// A default retry strategy applies to this operation UpdateManagementAppliance()
func (client ManagementApplianceClient) UpdateManagementAppliance(ctx context.Context, request UpdateManagementApplianceRequest) (response UpdateManagementApplianceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateManagementAppliance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementApplianceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementApplianceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementApplianceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementApplianceResponse")
	}
	return
}

// updateManagementAppliance implements the OCIOperation interface (enables retrying operations)
func (client ManagementApplianceClient) updateManagementAppliance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementAppliances/{managementApplianceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagementApplianceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ManagementAppliance/UpdateManagementAppliance"
		err = common.PostProcessServiceError(err, "ManagementAppliance", "UpdateManagementAppliance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
