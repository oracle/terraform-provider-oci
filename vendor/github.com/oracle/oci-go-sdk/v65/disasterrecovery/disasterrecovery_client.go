// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DisasterRecoveryClient a client for DisasterRecovery
type DisasterRecoveryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDisasterRecoveryClientWithConfigurationProvider Creates a new default DisasterRecovery client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDisasterRecoveryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DisasterRecoveryClient, err error) {
	if enabled := common.CheckForEnabledServices("disasterrecovery"); !enabled {
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
	return newDisasterRecoveryClientFromBaseClient(baseClient, provider)
}

// NewDisasterRecoveryClientWithOboToken Creates a new default DisasterRecovery client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDisasterRecoveryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DisasterRecoveryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDisasterRecoveryClientFromBaseClient(baseClient, configProvider)
}

func newDisasterRecoveryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DisasterRecoveryClient, err error) {
	// DisasterRecovery service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DisasterRecovery"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DisasterRecoveryClient{BaseClient: baseClient}
	client.BasePath = "20220125"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DisasterRecoveryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("disasterrecovery", "https://disaster-recovery.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DisasterRecoveryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DisasterRecoveryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AssociateDrProtectionGroup Create an association between the DR protection group identified by *drProtectionGroupId* and
// another DR protection group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/AssociateDrProtectionGroup.go.html to see an example of how to use AssociateDrProtectionGroup API.
// A default retry strategy applies to this operation AssociateDrProtectionGroup()
func (client DisasterRecoveryClient) AssociateDrProtectionGroup(ctx context.Context, request AssociateDrProtectionGroupRequest) (response AssociateDrProtectionGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.associateDrProtectionGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AssociateDrProtectionGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AssociateDrProtectionGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AssociateDrProtectionGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AssociateDrProtectionGroupResponse")
	}
	return
}

// associateDrProtectionGroup implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) associateDrProtectionGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drProtectionGroups/{drProtectionGroupId}/actions/associate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AssociateDrProtectionGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/AssociateDrProtectionGroup"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "AssociateDrProtectionGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelDrPlanExecution Cancel the DR plan execution identified by *drPlanExecutionId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/CancelDrPlanExecution.go.html to see an example of how to use CancelDrPlanExecution API.
// A default retry strategy applies to this operation CancelDrPlanExecution()
func (client DisasterRecoveryClient) CancelDrPlanExecution(ctx context.Context, request CancelDrPlanExecutionRequest) (response CancelDrPlanExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelDrPlanExecutionResponse")
	}
	return
}

// cancelDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) cancelDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drPlanExecutions/{drPlanExecutionId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/CancelDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "CancelDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancel the work request identified by *workRequestId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client DisasterRecoveryClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDrProtectionGroupCompartment Move the DR protection group identified by *drProtectionGroupId* to a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ChangeDrProtectionGroupCompartment.go.html to see an example of how to use ChangeDrProtectionGroupCompartment API.
// A default retry strategy applies to this operation ChangeDrProtectionGroupCompartment()
func (client DisasterRecoveryClient) ChangeDrProtectionGroupCompartment(ctx context.Context, request ChangeDrProtectionGroupCompartmentRequest) (response ChangeDrProtectionGroupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDrProtectionGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDrProtectionGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDrProtectionGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDrProtectionGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDrProtectionGroupCompartmentResponse")
	}
	return
}

// changeDrProtectionGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) changeDrProtectionGroupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drProtectionGroups/{drProtectionGroupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDrProtectionGroupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/ChangeDrProtectionGroupCompartment"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ChangeDrProtectionGroupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDrPlan Create a DR plan of the specified DR plan type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/CreateDrPlan.go.html to see an example of how to use CreateDrPlan API.
// A default retry strategy applies to this operation CreateDrPlan()
func (client DisasterRecoveryClient) CreateDrPlan(ctx context.Context, request CreateDrPlanRequest) (response CreateDrPlanResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDrPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDrPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDrPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDrPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDrPlanResponse")
	}
	return
}

// createDrPlan implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) createDrPlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drPlans", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDrPlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlan/CreateDrPlan"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "CreateDrPlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDrPlanExecution Execute a DR plan for a DR protection group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/CreateDrPlanExecution.go.html to see an example of how to use CreateDrPlanExecution API.
// A default retry strategy applies to this operation CreateDrPlanExecution()
func (client DisasterRecoveryClient) CreateDrPlanExecution(ctx context.Context, request CreateDrPlanExecutionRequest) (response CreateDrPlanExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDrPlanExecutionResponse")
	}
	return
}

// createDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) createDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drPlanExecutions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/CreateDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "CreateDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDrProtectionGroup Create a DR protection group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/CreateDrProtectionGroup.go.html to see an example of how to use CreateDrProtectionGroup API.
// A default retry strategy applies to this operation CreateDrProtectionGroup()
func (client DisasterRecoveryClient) CreateDrProtectionGroup(ctx context.Context, request CreateDrProtectionGroupRequest) (response CreateDrProtectionGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDrProtectionGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDrProtectionGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDrProtectionGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDrProtectionGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDrProtectionGroupResponse")
	}
	return
}

// createDrProtectionGroup implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) createDrProtectionGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drProtectionGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDrProtectionGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/CreateDrProtectionGroup"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "CreateDrProtectionGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDrPlan Delete the DR plan identified by *drPlanId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/DeleteDrPlan.go.html to see an example of how to use DeleteDrPlan API.
// A default retry strategy applies to this operation DeleteDrPlan()
func (client DisasterRecoveryClient) DeleteDrPlan(ctx context.Context, request DeleteDrPlanRequest) (response DeleteDrPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDrPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDrPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDrPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDrPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDrPlanResponse")
	}
	return
}

// deleteDrPlan implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) deleteDrPlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/drPlans/{drPlanId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDrPlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlan/DeleteDrPlan"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "DeleteDrPlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDrPlanExecution Delete the DR plan execution identified by *drPlanExecutionId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/DeleteDrPlanExecution.go.html to see an example of how to use DeleteDrPlanExecution API.
// A default retry strategy applies to this operation DeleteDrPlanExecution()
func (client DisasterRecoveryClient) DeleteDrPlanExecution(ctx context.Context, request DeleteDrPlanExecutionRequest) (response DeleteDrPlanExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDrPlanExecutionResponse")
	}
	return
}

// deleteDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) deleteDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/drPlanExecutions/{drPlanExecutionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/DeleteDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "DeleteDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDrProtectionGroup Delete the DR protection group identified by *drProtectionGroupId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/DeleteDrProtectionGroup.go.html to see an example of how to use DeleteDrProtectionGroup API.
// A default retry strategy applies to this operation DeleteDrProtectionGroup()
func (client DisasterRecoveryClient) DeleteDrProtectionGroup(ctx context.Context, request DeleteDrProtectionGroupRequest) (response DeleteDrProtectionGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDrProtectionGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDrProtectionGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDrProtectionGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDrProtectionGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDrProtectionGroupResponse")
	}
	return
}

// deleteDrProtectionGroup implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) deleteDrProtectionGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/drProtectionGroups/{drProtectionGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDrProtectionGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/DeleteDrProtectionGroup"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "DeleteDrProtectionGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisassociateDrProtectionGroup Delete the association between the DR protection group identified by *drProtectionGroupId*.
// and its peer DR protection group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/DisassociateDrProtectionGroup.go.html to see an example of how to use DisassociateDrProtectionGroup API.
// A default retry strategy applies to this operation DisassociateDrProtectionGroup()
func (client DisasterRecoveryClient) DisassociateDrProtectionGroup(ctx context.Context, request DisassociateDrProtectionGroupRequest) (response DisassociateDrProtectionGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disassociateDrProtectionGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisassociateDrProtectionGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisassociateDrProtectionGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisassociateDrProtectionGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisassociateDrProtectionGroupResponse")
	}
	return
}

// disassociateDrProtectionGroup implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) disassociateDrProtectionGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drProtectionGroups/{drProtectionGroupId}/actions/disassociate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisassociateDrProtectionGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/DisassociateDrProtectionGroup"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "DisassociateDrProtectionGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDrPlan Get details for the DR plan identified by *drPlanId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/GetDrPlan.go.html to see an example of how to use GetDrPlan API.
// A default retry strategy applies to this operation GetDrPlan()
func (client DisasterRecoveryClient) GetDrPlan(ctx context.Context, request GetDrPlanRequest) (response GetDrPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDrPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDrPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDrPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDrPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDrPlanResponse")
	}
	return
}

// getDrPlan implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) getDrPlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/drPlans/{drPlanId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDrPlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlan/GetDrPlan"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "GetDrPlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDrPlanExecution Get details for the DR plan execution identified by *drPlanExecutionId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/GetDrPlanExecution.go.html to see an example of how to use GetDrPlanExecution API.
// A default retry strategy applies to this operation GetDrPlanExecution()
func (client DisasterRecoveryClient) GetDrPlanExecution(ctx context.Context, request GetDrPlanExecutionRequest) (response GetDrPlanExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDrPlanExecutionResponse")
	}
	return
}

// getDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) getDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/drPlanExecutions/{drPlanExecutionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/GetDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "GetDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDrProtectionGroup Get the DR protection group identified by *drProtectionGroupId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/GetDrProtectionGroup.go.html to see an example of how to use GetDrProtectionGroup API.
// A default retry strategy applies to this operation GetDrProtectionGroup()
func (client DisasterRecoveryClient) GetDrProtectionGroup(ctx context.Context, request GetDrProtectionGroupRequest) (response GetDrProtectionGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDrProtectionGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDrProtectionGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDrProtectionGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDrProtectionGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDrProtectionGroupResponse")
	}
	return
}

// getDrProtectionGroup implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) getDrProtectionGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/drProtectionGroups/{drProtectionGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDrProtectionGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/GetDrProtectionGroup"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "GetDrProtectionGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Get the status of the work request identified by *workRequestId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DisasterRecoveryClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DisasterRecoveryClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IgnoreDrPlanExecution Ignore the failed group or step in DR plan execution identified by *drPlanExecutionId* and resume execution.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/IgnoreDrPlanExecution.go.html to see an example of how to use IgnoreDrPlanExecution API.
// A default retry strategy applies to this operation IgnoreDrPlanExecution()
func (client DisasterRecoveryClient) IgnoreDrPlanExecution(ctx context.Context, request IgnoreDrPlanExecutionRequest) (response IgnoreDrPlanExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ignoreDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IgnoreDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IgnoreDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IgnoreDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IgnoreDrPlanExecutionResponse")
	}
	return
}

// ignoreDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) ignoreDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drPlanExecutions/{drPlanExecutionId}/actions/ignore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IgnoreDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/IgnoreDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "IgnoreDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDrPlanExecutions Get a summary list of all DR plan executions for a DR protection group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListDrPlanExecutions.go.html to see an example of how to use ListDrPlanExecutions API.
// A default retry strategy applies to this operation ListDrPlanExecutions()
func (client DisasterRecoveryClient) ListDrPlanExecutions(ctx context.Context, request ListDrPlanExecutionsRequest) (response ListDrPlanExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDrPlanExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDrPlanExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDrPlanExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDrPlanExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDrPlanExecutionsResponse")
	}
	return
}

// listDrPlanExecutions implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) listDrPlanExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/drPlanExecutions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDrPlanExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/ListDrPlanExecutions"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ListDrPlanExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDrPlans Get a summary list of all DR plans for a DR protection group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListDrPlans.go.html to see an example of how to use ListDrPlans API.
// A default retry strategy applies to this operation ListDrPlans()
func (client DisasterRecoveryClient) ListDrPlans(ctx context.Context, request ListDrPlansRequest) (response ListDrPlansResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDrPlans, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDrPlansResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDrPlansResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDrPlansResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDrPlansResponse")
	}
	return
}

// listDrPlans implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) listDrPlans(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/drPlans", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDrPlansResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlan/ListDrPlans"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ListDrPlans", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDrProtectionGroups Get a summary list of all DR protection groups in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListDrProtectionGroups.go.html to see an example of how to use ListDrProtectionGroups API.
// A default retry strategy applies to this operation ListDrProtectionGroups()
func (client DisasterRecoveryClient) ListDrProtectionGroups(ctx context.Context, request ListDrProtectionGroupsRequest) (response ListDrProtectionGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDrProtectionGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDrProtectionGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDrProtectionGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDrProtectionGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDrProtectionGroupsResponse")
	}
	return
}

// listDrProtectionGroups implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) listDrProtectionGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/drProtectionGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDrProtectionGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/ListDrProtectionGroups"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ListDrProtectionGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Get a list of work request errors for the work request identified by *workRequestId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DisasterRecoveryClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DisasterRecoveryClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Get a list of logs for the work request identified by *workRequestId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DisasterRecoveryClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DisasterRecoveryClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DisasterRecoveryClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DisasterRecoveryClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PauseDrPlanExecution Pause the DR plan execution identified by *drPlanExecutionId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/PauseDrPlanExecution.go.html to see an example of how to use PauseDrPlanExecution API.
// A default retry strategy applies to this operation PauseDrPlanExecution()
func (client DisasterRecoveryClient) PauseDrPlanExecution(ctx context.Context, request PauseDrPlanExecutionRequest) (response PauseDrPlanExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.pauseDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PauseDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PauseDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PauseDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PauseDrPlanExecutionResponse")
	}
	return
}

// pauseDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) pauseDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drPlanExecutions/{drPlanExecutionId}/actions/pause", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PauseDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/PauseDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "PauseDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResumeDrPlanExecution Resume the DR plan execution identified by *drPlanExecutionId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ResumeDrPlanExecution.go.html to see an example of how to use ResumeDrPlanExecution API.
// A default retry strategy applies to this operation ResumeDrPlanExecution()
func (client DisasterRecoveryClient) ResumeDrPlanExecution(ctx context.Context, request ResumeDrPlanExecutionRequest) (response ResumeDrPlanExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resumeDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeDrPlanExecutionResponse")
	}
	return
}

// resumeDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) resumeDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drPlanExecutions/{drPlanExecutionId}/actions/resume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResumeDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/ResumeDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "ResumeDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetryDrPlanExecution Retry the failed group or step in DR plan execution identified by *drPlanExecutionId* and resume execution.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/RetryDrPlanExecution.go.html to see an example of how to use RetryDrPlanExecution API.
// A default retry strategy applies to this operation RetryDrPlanExecution()
func (client DisasterRecoveryClient) RetryDrPlanExecution(ctx context.Context, request RetryDrPlanExecutionRequest) (response RetryDrPlanExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.retryDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetryDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetryDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetryDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetryDrPlanExecutionResponse")
	}
	return
}

// retryDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) retryDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drPlanExecutions/{drPlanExecutionId}/actions/retry", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetryDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/RetryDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "RetryDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDrPlan Update the DR plan identified by *drPlanId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/UpdateDrPlan.go.html to see an example of how to use UpdateDrPlan API.
// A default retry strategy applies to this operation UpdateDrPlan()
func (client DisasterRecoveryClient) UpdateDrPlan(ctx context.Context, request UpdateDrPlanRequest) (response UpdateDrPlanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDrPlan, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDrPlanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDrPlanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDrPlanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDrPlanResponse")
	}
	return
}

// updateDrPlan implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) updateDrPlan(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/drPlans/{drPlanId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDrPlanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlan/UpdateDrPlan"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "UpdateDrPlan", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDrPlanExecution Update the DR plan execution identified by *drPlanExecutionId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/UpdateDrPlanExecution.go.html to see an example of how to use UpdateDrPlanExecution API.
// A default retry strategy applies to this operation UpdateDrPlanExecution()
func (client DisasterRecoveryClient) UpdateDrPlanExecution(ctx context.Context, request UpdateDrPlanExecutionRequest) (response UpdateDrPlanExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDrPlanExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDrPlanExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDrPlanExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDrPlanExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDrPlanExecutionResponse")
	}
	return
}

// updateDrPlanExecution implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) updateDrPlanExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/drPlanExecutions/{drPlanExecutionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDrPlanExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrPlanExecution/UpdateDrPlanExecution"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "UpdateDrPlanExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDrProtectionGroup Update the DR protection group identified by *drProtectionGroupId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/UpdateDrProtectionGroup.go.html to see an example of how to use UpdateDrProtectionGroup API.
// A default retry strategy applies to this operation UpdateDrProtectionGroup()
func (client DisasterRecoveryClient) UpdateDrProtectionGroup(ctx context.Context, request UpdateDrProtectionGroupRequest) (response UpdateDrProtectionGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDrProtectionGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDrProtectionGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDrProtectionGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDrProtectionGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDrProtectionGroupResponse")
	}
	return
}

// updateDrProtectionGroup implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) updateDrProtectionGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/drProtectionGroups/{drProtectionGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDrProtectionGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/UpdateDrProtectionGroup"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "UpdateDrProtectionGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDrProtectionGroupRole Update the role of the DR protection group identified by *drProtectionGroupId*.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/UpdateDrProtectionGroupRole.go.html to see an example of how to use UpdateDrProtectionGroupRole API.
// A default retry strategy applies to this operation UpdateDrProtectionGroupRole()
func (client DisasterRecoveryClient) UpdateDrProtectionGroupRole(ctx context.Context, request UpdateDrProtectionGroupRoleRequest) (response UpdateDrProtectionGroupRoleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateDrProtectionGroupRole, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDrProtectionGroupRoleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDrProtectionGroupRoleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDrProtectionGroupRoleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDrProtectionGroupRoleResponse")
	}
	return
}

// updateDrProtectionGroupRole implements the OCIOperation interface (enables retrying operations)
func (client DisasterRecoveryClient) updateDrProtectionGroupRole(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/drProtectionGroups/{drProtectionGroupId}/actions/updateRole", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDrProtectionGroupRoleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/disaster-recovery/20220125/DrProtectionGroup/UpdateDrProtectionGroupRole"
		err = common.PostProcessServiceError(err, "DisasterRecovery", "UpdateDrProtectionGroupRole", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
