// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// PrivilegedApiControlClient a client for PrivilegedApiControl
type PrivilegedApiControlClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPrivilegedApiControlClientWithConfigurationProvider Creates a new default PrivilegedApiControl client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPrivilegedApiControlClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PrivilegedApiControlClient, err error) {
	if enabled := common.CheckForEnabledServices("apiaccesscontrol"); !enabled {
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
	return newPrivilegedApiControlClientFromBaseClient(baseClient, provider)
}

// NewPrivilegedApiControlClientWithOboToken Creates a new default PrivilegedApiControl client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewPrivilegedApiControlClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PrivilegedApiControlClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPrivilegedApiControlClientFromBaseClient(baseClient, configProvider)
}

func newPrivilegedApiControlClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PrivilegedApiControlClient, err error) {
	// PrivilegedApiControl service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("PrivilegedApiControl"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PrivilegedApiControlClient{BaseClient: baseClient}
	client.BasePath = "20241130"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PrivilegedApiControlClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apiaccesscontrol", "https://pactl.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PrivilegedApiControlClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PrivilegedApiControlClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangePrivilegedApiControlCompartment Moves a PrivilegedApiControl into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ChangePrivilegedApiControlCompartment.go.html to see an example of how to use ChangePrivilegedApiControlCompartment API.
// A default retry strategy applies to this operation ChangePrivilegedApiControlCompartment()
func (client PrivilegedApiControlClient) ChangePrivilegedApiControlCompartment(ctx context.Context, request ChangePrivilegedApiControlCompartmentRequest) (response ChangePrivilegedApiControlCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePrivilegedApiControlCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePrivilegedApiControlCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePrivilegedApiControlCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePrivilegedApiControlCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePrivilegedApiControlCompartmentResponse")
	}
	return
}

// changePrivilegedApiControlCompartment implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiControlClient) changePrivilegedApiControlCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privilegedApiControls/{privilegedApiControlId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePrivilegedApiControlCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/20241130/PrivilegedApiControl/ChangePrivilegedApiControlCompartment"
		err = common.PostProcessServiceError(err, "PrivilegedApiControl", "ChangePrivilegedApiControlCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePrivilegedApiControl Creates a PrivilegedApiControl.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/CreatePrivilegedApiControl.go.html to see an example of how to use CreatePrivilegedApiControl API.
// A default retry strategy applies to this operation CreatePrivilegedApiControl()
func (client PrivilegedApiControlClient) CreatePrivilegedApiControl(ctx context.Context, request CreatePrivilegedApiControlRequest) (response CreatePrivilegedApiControlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPrivilegedApiControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePrivilegedApiControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePrivilegedApiControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePrivilegedApiControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePrivilegedApiControlResponse")
	}
	return
}

// createPrivilegedApiControl implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiControlClient) createPrivilegedApiControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privilegedApiControls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePrivilegedApiControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/20241130/PrivilegedApiControl/CreatePrivilegedApiControl"
		err = common.PostProcessServiceError(err, "PrivilegedApiControl", "CreatePrivilegedApiControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePrivilegedApiControl Deletes a PrivilegedApiControl.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/DeletePrivilegedApiControl.go.html to see an example of how to use DeletePrivilegedApiControl API.
// A default retry strategy applies to this operation DeletePrivilegedApiControl()
func (client PrivilegedApiControlClient) DeletePrivilegedApiControl(ctx context.Context, request DeletePrivilegedApiControlRequest) (response DeletePrivilegedApiControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePrivilegedApiControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePrivilegedApiControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePrivilegedApiControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePrivilegedApiControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePrivilegedApiControlResponse")
	}
	return
}

// deletePrivilegedApiControl implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiControlClient) deletePrivilegedApiControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/privilegedApiControls/{privilegedApiControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePrivilegedApiControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/20241130/PrivilegedApiControl/DeletePrivilegedApiControl"
		err = common.PostProcessServiceError(err, "PrivilegedApiControl", "DeletePrivilegedApiControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivilegedApiControl Gets information about a PrivilegedApiControl.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/GetPrivilegedApiControl.go.html to see an example of how to use GetPrivilegedApiControl API.
// A default retry strategy applies to this operation GetPrivilegedApiControl()
func (client PrivilegedApiControlClient) GetPrivilegedApiControl(ctx context.Context, request GetPrivilegedApiControlRequest) (response GetPrivilegedApiControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivilegedApiControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivilegedApiControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivilegedApiControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivilegedApiControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivilegedApiControlResponse")
	}
	return
}

// getPrivilegedApiControl implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiControlClient) getPrivilegedApiControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privilegedApiControls/{privilegedApiControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivilegedApiControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/20241130/PrivilegedApiControl/GetPrivilegedApiControl"
		err = common.PostProcessServiceError(err, "PrivilegedApiControl", "GetPrivilegedApiControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPrivilegedApiControls Gets a list of PrivilegedApiControls.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/ListPrivilegedApiControls.go.html to see an example of how to use ListPrivilegedApiControls API.
// A default retry strategy applies to this operation ListPrivilegedApiControls()
func (client PrivilegedApiControlClient) ListPrivilegedApiControls(ctx context.Context, request ListPrivilegedApiControlsRequest) (response ListPrivilegedApiControlsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPrivilegedApiControls, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPrivilegedApiControlsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPrivilegedApiControlsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPrivilegedApiControlsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPrivilegedApiControlsResponse")
	}
	return
}

// listPrivilegedApiControls implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiControlClient) listPrivilegedApiControls(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privilegedApiControls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPrivilegedApiControlsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/20241130/PrivilegedApiControlCollection/ListPrivilegedApiControls"
		err = common.PostProcessServiceError(err, "PrivilegedApiControl", "ListPrivilegedApiControls", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePrivilegedApiControl Updates a PrivilegedApiControl.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apiaccesscontrol/UpdatePrivilegedApiControl.go.html to see an example of how to use UpdatePrivilegedApiControl API.
// A default retry strategy applies to this operation UpdatePrivilegedApiControl()
func (client PrivilegedApiControlClient) UpdatePrivilegedApiControl(ctx context.Context, request UpdatePrivilegedApiControlRequest) (response UpdatePrivilegedApiControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePrivilegedApiControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePrivilegedApiControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePrivilegedApiControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePrivilegedApiControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePrivilegedApiControlResponse")
	}
	return
}

// updatePrivilegedApiControl implements the OCIOperation interface (enables retrying operations)
func (client PrivilegedApiControlClient) updatePrivilegedApiControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/privilegedApiControls/{privilegedApiControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePrivilegedApiControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/oracle-api-access-control/20241130/PrivilegedApiControl/UpdatePrivilegedApiControl"
		err = common.PostProcessServiceError(err, "PrivilegedApiControl", "UpdatePrivilegedApiControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
