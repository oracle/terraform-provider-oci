// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// OperatorControlAssignmentClient a client for OperatorControlAssignment
type OperatorControlAssignmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOperatorControlAssignmentClientWithConfigurationProvider Creates a new default OperatorControlAssignment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOperatorControlAssignmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OperatorControlAssignmentClient, err error) {
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
	return newOperatorControlAssignmentClientFromBaseClient(baseClient, provider)
}

// NewOperatorControlAssignmentClientWithOboToken Creates a new default OperatorControlAssignment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOperatorControlAssignmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OperatorControlAssignmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOperatorControlAssignmentClientFromBaseClient(baseClient, configProvider)
}

func newOperatorControlAssignmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OperatorControlAssignmentClient, err error) {
	// OperatorControlAssignment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OperatorControlAssignment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OperatorControlAssignmentClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OperatorControlAssignmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("operatoraccesscontrol", "https://operator-access-control.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OperatorControlAssignmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OperatorControlAssignmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOperatorControlAssignmentCompartment Changes the compartment of the specified Operator Control assignment ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ChangeOperatorControlAssignmentCompartment.go.html to see an example of how to use ChangeOperatorControlAssignmentCompartment API.
// A default retry strategy applies to this operation ChangeOperatorControlAssignmentCompartment()
func (client OperatorControlAssignmentClient) ChangeOperatorControlAssignmentCompartment(ctx context.Context, request ChangeOperatorControlAssignmentCompartmentRequest) (response ChangeOperatorControlAssignmentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOperatorControlAssignmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOperatorControlAssignmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOperatorControlAssignmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOperatorControlAssignmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOperatorControlAssignmentCompartmentResponse")
	}
	return
}

// changeOperatorControlAssignmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) changeOperatorControlAssignmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operatorControlAssignments/{operatorControlAssignmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOperatorControlAssignmentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "ChangeOperatorControlAssignmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOperatorControlAssignment Creates an Operator Control Assignment resource. In effect, this brings the target resource under the governance of the Operator Control for specified time duration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/CreateOperatorControlAssignment.go.html to see an example of how to use CreateOperatorControlAssignment API.
// A default retry strategy applies to this operation CreateOperatorControlAssignment()
func (client OperatorControlAssignmentClient) CreateOperatorControlAssignment(ctx context.Context, request CreateOperatorControlAssignmentRequest) (response CreateOperatorControlAssignmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOperatorControlAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOperatorControlAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOperatorControlAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOperatorControlAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOperatorControlAssignmentResponse")
	}
	return
}

// createOperatorControlAssignment implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) createOperatorControlAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operatorControlAssignments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOperatorControlAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "CreateOperatorControlAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOperatorControlAssignment Deletes the specified Operator Control Assignment. This has the effect of unassigning the specific Operator Control from the target resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/DeleteOperatorControlAssignment.go.html to see an example of how to use DeleteOperatorControlAssignment API.
// A default retry strategy applies to this operation DeleteOperatorControlAssignment()
func (client OperatorControlAssignmentClient) DeleteOperatorControlAssignment(ctx context.Context, request DeleteOperatorControlAssignmentRequest) (response DeleteOperatorControlAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOperatorControlAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOperatorControlAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOperatorControlAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOperatorControlAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOperatorControlAssignmentResponse")
	}
	return
}

// deleteOperatorControlAssignment implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) deleteOperatorControlAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/operatorControlAssignments/{operatorControlAssignmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOperatorControlAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "DeleteOperatorControlAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssignmentValidationStatus Returns the status of assignment validation associated with the assignmentId and the executionId.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/GetAssignmentValidationStatus.go.html to see an example of how to use GetAssignmentValidationStatus API.
// A default retry strategy applies to this operation GetAssignmentValidationStatus()
func (client OperatorControlAssignmentClient) GetAssignmentValidationStatus(ctx context.Context, request GetAssignmentValidationStatusRequest) (response GetAssignmentValidationStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssignmentValidationStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssignmentValidationStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssignmentValidationStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssignmentValidationStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssignmentValidationStatusResponse")
	}
	return
}

// getAssignmentValidationStatus implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) getAssignmentValidationStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operatorControlAssignments/{operatorControlAssignmentId}/actions/getAssignmentValidationStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssignmentValidationStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "GetAssignmentValidationStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOperatorControlAssignment Gets the details of an Operator Control Assignment of the specified ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/GetOperatorControlAssignment.go.html to see an example of how to use GetOperatorControlAssignment API.
// A default retry strategy applies to this operation GetOperatorControlAssignment()
func (client OperatorControlAssignmentClient) GetOperatorControlAssignment(ctx context.Context, request GetOperatorControlAssignmentRequest) (response GetOperatorControlAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOperatorControlAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOperatorControlAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOperatorControlAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOperatorControlAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOperatorControlAssignmentResponse")
	}
	return
}

// getOperatorControlAssignment implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) getOperatorControlAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operatorControlAssignments/{operatorControlAssignmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOperatorControlAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "GetOperatorControlAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOperatorControlAssignments Lists all Operator Control Assignments.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListOperatorControlAssignments.go.html to see an example of how to use ListOperatorControlAssignments API.
// A default retry strategy applies to this operation ListOperatorControlAssignments()
func (client OperatorControlAssignmentClient) ListOperatorControlAssignments(ctx context.Context, request ListOperatorControlAssignmentsRequest) (response ListOperatorControlAssignmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOperatorControlAssignments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOperatorControlAssignmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOperatorControlAssignmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOperatorControlAssignmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOperatorControlAssignmentsResponse")
	}
	return
}

// listOperatorControlAssignments implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) listOperatorControlAssignments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/operatorControlAssignments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOperatorControlAssignmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "ListOperatorControlAssignments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOperatorControlAssignment Modifies the existing Operator Control assignment of the specified Operator Control assignment ID. Modifying the assignment does not change the Operator Control assignment ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/UpdateOperatorControlAssignment.go.html to see an example of how to use UpdateOperatorControlAssignment API.
// A default retry strategy applies to this operation UpdateOperatorControlAssignment()
func (client OperatorControlAssignmentClient) UpdateOperatorControlAssignment(ctx context.Context, request UpdateOperatorControlAssignmentRequest) (response UpdateOperatorControlAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOperatorControlAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOperatorControlAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOperatorControlAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOperatorControlAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOperatorControlAssignmentResponse")
	}
	return
}

// updateOperatorControlAssignment implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) updateOperatorControlAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/operatorControlAssignments/{operatorControlAssignmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOperatorControlAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "UpdateOperatorControlAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateOperatorAssignment Validates the assignment for the specified Operator Control Assignment ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ValidateOperatorAssignment.go.html to see an example of how to use ValidateOperatorAssignment API.
// A default retry strategy applies to this operation ValidateOperatorAssignment()
func (client OperatorControlAssignmentClient) ValidateOperatorAssignment(ctx context.Context, request ValidateOperatorAssignmentRequest) (response ValidateOperatorAssignmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateOperatorAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateOperatorAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateOperatorAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateOperatorAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateOperatorAssignmentResponse")
	}
	return
}

// validateOperatorAssignment implements the OCIOperation interface (enables retrying operations)
func (client OperatorControlAssignmentClient) validateOperatorAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/operatorControlAssignments/{operatorControlAssignmentId}/actions/validateAssignment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateOperatorAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "OperatorControlAssignment", "ValidateOperatorAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
