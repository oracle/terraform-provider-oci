// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OperatorControlClient a client for OperatorControl
type OperatorControlClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOperatorControlClientWithConfigurationProvider Creates a new default OperatorControl client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOperatorControlClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OperatorControlClient, err error) {
	if enabled := common.CheckForEnabledServices("operatoraccesscontrol"); !enabled {
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
	return newOperatorControlClientFromBaseClient(baseClient, provider)
}

// NewOperatorControlClientWithOboToken Creates a new default OperatorControl client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOperatorControlClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OperatorControlClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOperatorControlClientFromBaseClient(baseClient, configProvider)
}

func newOperatorControlClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OperatorControlClient, err error) {
	// OperatorControl service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OperatorControl"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OperatorControlClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OperatorControlClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("operatoraccesscontrol", "https://operator-access-control.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OperatorControlClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OperatorControlClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOperatorControlCompartment Moves the Operator Control resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ChangeOperatorControlCompartment.go.html to see an example of how to use ChangeOperatorControlCompartment API.
// A default retry strategy applies to this operation ChangeOperatorControlCompartment()
func (client OperatorControlClient) ChangeOperatorControlCompartment(ctx context.Context, request ChangeOperatorControlCompartmentRequest) (response ChangeOperatorControlCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOperatorControlCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOperatorControlCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOperatorControlCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOperatorControlCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOperatorControlCompartmentResponse")
	}
	return
}

// changeOperatorControlCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlClient) changeOperatorControlCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operatorControls/{operatorControlId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOperatorControlCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControl", "ChangeOperatorControlCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOperatorControl Creates an Operator Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/CreateOperatorControl.go.html to see an example of how to use CreateOperatorControl API.
// A default retry strategy applies to this operation CreateOperatorControl()
func (client OperatorControlClient) CreateOperatorControl(ctx context.Context, request CreateOperatorControlRequest) (response CreateOperatorControlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOperatorControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOperatorControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOperatorControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOperatorControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOperatorControlResponse")
	}
	return
}

// createOperatorControl implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlClient) createOperatorControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operatorControls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOperatorControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControl", "CreateOperatorControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOperatorControl Deletes an Operator Control. You cannot delete an Operator Control if it is assigned to govern any target resource currently or in the future.
// In that case, first, delete all of the current and future assignments before deleting the Operator Control. An Operator Control that was previously assigned to a target
// resource is marked as DELETED following a successful deletion. However, it is not completely deleted from the system. This is to ensure auditing information for the accesses
// done under the Operator Control is preserved for future needs. The system purges the deleted Operator Control only when all of the audit data associated with the
// Operator Control are also deleted. Therefore, you cannot reuse the name of the deleted Operator Control until the system purges the Operator Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/DeleteOperatorControl.go.html to see an example of how to use DeleteOperatorControl API.
// A default retry strategy applies to this operation DeleteOperatorControl()
func (client OperatorControlClient) DeleteOperatorControl(ctx context.Context, request DeleteOperatorControlRequest) (response DeleteOperatorControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOperatorControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOperatorControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOperatorControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOperatorControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOperatorControlResponse")
	}
	return
}

// deleteOperatorControl implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlClient) deleteOperatorControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/operatorControls/{operatorControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOperatorControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControl", "DeleteOperatorControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOperatorControl Gets the Operator Control associated with the specified Operator Control ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/GetOperatorControl.go.html to see an example of how to use GetOperatorControl API.
// A default retry strategy applies to this operation GetOperatorControl()
func (client OperatorControlClient) GetOperatorControl(ctx context.Context, request GetOperatorControlRequest) (response GetOperatorControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOperatorControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOperatorControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOperatorControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOperatorControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOperatorControlResponse")
	}
	return
}

// getOperatorControl implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlClient) getOperatorControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operatorControls/{operatorControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOperatorControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControl", "GetOperatorControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOperatorControls Lists the operator controls in the compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListOperatorControls.go.html to see an example of how to use ListOperatorControls API.
// A default retry strategy applies to this operation ListOperatorControls()
func (client OperatorControlClient) ListOperatorControls(ctx context.Context, request ListOperatorControlsRequest) (response ListOperatorControlsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOperatorControls, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOperatorControlsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOperatorControlsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOperatorControlsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOperatorControlsResponse")
	}
	return
}

// listOperatorControls implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlClient) listOperatorControls(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operatorControls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOperatorControlsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControl", "ListOperatorControls", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOperatorControl Modifies the existing OperatorControl for a given operator control id except the operator control id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/UpdateOperatorControl.go.html to see an example of how to use UpdateOperatorControl API.
// A default retry strategy applies to this operation UpdateOperatorControl()
func (client OperatorControlClient) UpdateOperatorControl(ctx context.Context, request UpdateOperatorControlRequest) (response UpdateOperatorControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOperatorControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOperatorControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOperatorControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOperatorControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOperatorControlResponse")
	}
	return
}

// updateOperatorControl implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlClient) updateOperatorControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/operatorControls/{operatorControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOperatorControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControl", "UpdateOperatorControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
