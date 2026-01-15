// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts. For more information, see Budgets Overview (https://docs.oracle.com/iaas/Content/Billing/Concepts/budgetsoverview.htm).
//

package budget

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// CostAdClient a client for CostAd
type CostAdClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCostAdClientWithConfigurationProvider Creates a new default CostAd client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCostAdClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CostAdClient, err error) {
	if enabled := common.CheckForEnabledServices("budget"); !enabled {
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
	return newCostAdClientFromBaseClient(baseClient, provider)
}

// NewCostAdClientWithOboToken Creates a new default CostAd client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCostAdClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CostAdClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCostAdClientFromBaseClient(baseClient, configProvider)
}

func newCostAdClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CostAdClient, err error) {
	// CostAd service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CostAd"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CostAdClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CostAdClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("budget", "https://usage.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CostAdClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CostAdClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateCostAlertSubscription Creates a new CostAlert Subscription.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/CreateCostAlertSubscription.go.html to see an example of how to use CreateCostAlertSubscription API.
// A default retry strategy applies to this operation CreateCostAlertSubscription()
func (client CostAdClient) CreateCostAlertSubscription(ctx context.Context, request CreateCostAlertSubscriptionRequest) (response CreateCostAlertSubscriptionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCostAlertSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCostAlertSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCostAlertSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCostAlertSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCostAlertSubscriptionResponse")
	}
	return
}

// createCostAlertSubscription implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) createCostAlertSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/costAlertSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCostAlertSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAlertSubscription/CreateCostAlertSubscription"
		err = common.PostProcessServiceError(err, "CostAd", "CreateCostAlertSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCostAnomalyMonitor Creates a new costAnomaly Monitor.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/CreateCostAnomalyMonitor.go.html to see an example of how to use CreateCostAnomalyMonitor API.
// A default retry strategy applies to this operation CreateCostAnomalyMonitor()
func (client CostAdClient) CreateCostAnomalyMonitor(ctx context.Context, request CreateCostAnomalyMonitorRequest) (response CreateCostAnomalyMonitorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCostAnomalyMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCostAnomalyMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCostAnomalyMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCostAnomalyMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCostAnomalyMonitorResponse")
	}
	return
}

// createCostAnomalyMonitor implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) createCostAnomalyMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/costAnomalyMonitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCostAnomalyMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyMonitor/CreateCostAnomalyMonitor"
		err = common.PostProcessServiceError(err, "CostAd", "CreateCostAnomalyMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCostAlertSubscription Deletes a specified CostAlertSubscription resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/DeleteCostAlertSubscription.go.html to see an example of how to use DeleteCostAlertSubscription API.
// A default retry strategy applies to this operation DeleteCostAlertSubscription()
func (client CostAdClient) DeleteCostAlertSubscription(ctx context.Context, request DeleteCostAlertSubscriptionRequest) (response DeleteCostAlertSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCostAlertSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCostAlertSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCostAlertSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCostAlertSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCostAlertSubscriptionResponse")
	}
	return
}

// deleteCostAlertSubscription implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) deleteCostAlertSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/costAlertSubscriptions/{costAlertSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCostAlertSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAlertSubscription/DeleteCostAlertSubscription"
		err = common.PostProcessServiceError(err, "CostAd", "DeleteCostAlertSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCostAnomalyMonitor Deletes a specified CostAnomalyMonitor resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/DeleteCostAnomalyMonitor.go.html to see an example of how to use DeleteCostAnomalyMonitor API.
// A default retry strategy applies to this operation DeleteCostAnomalyMonitor()
func (client CostAdClient) DeleteCostAnomalyMonitor(ctx context.Context, request DeleteCostAnomalyMonitorRequest) (response DeleteCostAnomalyMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCostAnomalyMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCostAnomalyMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCostAnomalyMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCostAnomalyMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCostAnomalyMonitorResponse")
	}
	return
}

// deleteCostAnomalyMonitor implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) deleteCostAnomalyMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/costAnomalyMonitors/{costAnomalyMonitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCostAnomalyMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyMonitor/DeleteCostAnomalyMonitor"
		err = common.PostProcessServiceError(err, "CostAd", "DeleteCostAnomalyMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableCostAnomalyMonitor Disables the cost anomaly monitor. This stops cost anomaly detection for targeted resource(s).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/DisableCostAnomalyMonitor.go.html to see an example of how to use DisableCostAnomalyMonitor API.
// A default retry strategy applies to this operation DisableCostAnomalyMonitor()
func (client CostAdClient) DisableCostAnomalyMonitor(ctx context.Context, request DisableCostAnomalyMonitorRequest) (response DisableCostAnomalyMonitorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableCostAnomalyMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableCostAnomalyMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableCostAnomalyMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableCostAnomalyMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableCostAnomalyMonitorResponse")
	}
	return
}

// disableCostAnomalyMonitor implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) disableCostAnomalyMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/costAnomalyMonitors/{costAnomalyMonitorId}/actions/disable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableCostAnomalyMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyMonitor/DisableCostAnomalyMonitor"
		err = common.PostProcessServiceError(err, "CostAd", "DisableCostAnomalyMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableCostAnomalyMonitor Enables the cost anomaly monitor. This (re)starts the cost anomaly detection for targeted resource(s).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/EnableCostAnomalyMonitor.go.html to see an example of how to use EnableCostAnomalyMonitor API.
// A default retry strategy applies to this operation EnableCostAnomalyMonitor()
func (client CostAdClient) EnableCostAnomalyMonitor(ctx context.Context, request EnableCostAnomalyMonitorRequest) (response EnableCostAnomalyMonitorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableCostAnomalyMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableCostAnomalyMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableCostAnomalyMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableCostAnomalyMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableCostAnomalyMonitorResponse")
	}
	return
}

// enableCostAnomalyMonitor implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) enableCostAnomalyMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/costAnomalyMonitors/{costAnomalyMonitorId}/actions/enable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableCostAnomalyMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyMonitor/EnableCostAnomalyMonitor"
		err = common.PostProcessServiceError(err, "CostAd", "EnableCostAnomalyMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCostAlertSubscription Gets a CostAlertSubscription by the identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/GetCostAlertSubscription.go.html to see an example of how to use GetCostAlertSubscription API.
// A default retry strategy applies to this operation GetCostAlertSubscription()
func (client CostAdClient) GetCostAlertSubscription(ctx context.Context, request GetCostAlertSubscriptionRequest) (response GetCostAlertSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCostAlertSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCostAlertSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCostAlertSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCostAlertSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCostAlertSubscriptionResponse")
	}
	return
}

// getCostAlertSubscription implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) getCostAlertSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/costAlertSubscriptions/{costAlertSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCostAlertSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAlertSubscription/GetCostAlertSubscription"
		err = common.PostProcessServiceError(err, "CostAd", "GetCostAlertSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCostAnomalyEvent Gets a CostAnomalyEvent by the identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/GetCostAnomalyEvent.go.html to see an example of how to use GetCostAnomalyEvent API.
// A default retry strategy applies to this operation GetCostAnomalyEvent()
func (client CostAdClient) GetCostAnomalyEvent(ctx context.Context, request GetCostAnomalyEventRequest) (response GetCostAnomalyEventResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCostAnomalyEvent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCostAnomalyEventResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCostAnomalyEventResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCostAnomalyEventResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCostAnomalyEventResponse")
	}
	return
}

// getCostAnomalyEvent implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) getCostAnomalyEvent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/costAnomalyEvents/{costAnomalyEventId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCostAnomalyEventResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyEvent/GetCostAnomalyEvent"
		err = common.PostProcessServiceError(err, "CostAd", "GetCostAnomalyEvent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCostAnomalyMonitor Gets a CostAnomalyMonitor by the identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/GetCostAnomalyMonitor.go.html to see an example of how to use GetCostAnomalyMonitor API.
// A default retry strategy applies to this operation GetCostAnomalyMonitor()
func (client CostAdClient) GetCostAnomalyMonitor(ctx context.Context, request GetCostAnomalyMonitorRequest) (response GetCostAnomalyMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCostAnomalyMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCostAnomalyMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCostAnomalyMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCostAnomalyMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCostAnomalyMonitorResponse")
	}
	return
}

// getCostAnomalyMonitor implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) getCostAnomalyMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/costAnomalyMonitors/{costAnomalyMonitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCostAnomalyMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyMonitor/GetCostAnomalyMonitor"
		err = common.PostProcessServiceError(err, "CostAd", "GetCostAnomalyMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCostAlertSubscriptions Gets a list of Cost Alert Subscription in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListCostAlertSubscriptions.go.html to see an example of how to use ListCostAlertSubscriptions API.
// A default retry strategy applies to this operation ListCostAlertSubscriptions()
func (client CostAdClient) ListCostAlertSubscriptions(ctx context.Context, request ListCostAlertSubscriptionsRequest) (response ListCostAlertSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCostAlertSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCostAlertSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCostAlertSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCostAlertSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCostAlertSubscriptionsResponse")
	}
	return
}

// listCostAlertSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) listCostAlertSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/costAlertSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCostAlertSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAlertSubscriptionCollection/ListCostAlertSubscriptions"
		err = common.PostProcessServiceError(err, "CostAd", "ListCostAlertSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCostAnomalyEvents Gets a list of Cost Anomaly Event in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListCostAnomalyEvents.go.html to see an example of how to use ListCostAnomalyEvents API.
// A default retry strategy applies to this operation ListCostAnomalyEvents()
func (client CostAdClient) ListCostAnomalyEvents(ctx context.Context, request ListCostAnomalyEventsRequest) (response ListCostAnomalyEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCostAnomalyEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCostAnomalyEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCostAnomalyEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCostAnomalyEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCostAnomalyEventsResponse")
	}
	return
}

// listCostAnomalyEvents implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) listCostAnomalyEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/costAnomalyEvents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCostAnomalyEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyEventCollection/ListCostAnomalyEvents"
		err = common.PostProcessServiceError(err, "CostAd", "ListCostAnomalyEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCostAnomalyMonitors Gets a list of Cost Anomaly Monitors in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListCostAnomalyMonitors.go.html to see an example of how to use ListCostAnomalyMonitors API.
// A default retry strategy applies to this operation ListCostAnomalyMonitors()
func (client CostAdClient) ListCostAnomalyMonitors(ctx context.Context, request ListCostAnomalyMonitorsRequest) (response ListCostAnomalyMonitorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCostAnomalyMonitors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCostAnomalyMonitorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCostAnomalyMonitorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCostAnomalyMonitorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCostAnomalyMonitorsResponse")
	}
	return
}

// listCostAnomalyMonitors implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) listCostAnomalyMonitors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/costAnomalyMonitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCostAnomalyMonitorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyMonitorCollection/ListCostAnomalyMonitors"
		err = common.PostProcessServiceError(err, "CostAd", "ListCostAnomalyMonitors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeCostAnomalyEventAnalytics Gets a list of Cost Anomaly Events analytics summary - aggregated metrics for a given time period.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/SummarizeCostAnomalyEventAnalytics.go.html to see an example of how to use SummarizeCostAnomalyEventAnalytics API.
// A default retry strategy applies to this operation SummarizeCostAnomalyEventAnalytics()
func (client CostAdClient) SummarizeCostAnomalyEventAnalytics(ctx context.Context, request SummarizeCostAnomalyEventAnalyticsRequest) (response SummarizeCostAnomalyEventAnalyticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeCostAnomalyEventAnalytics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeCostAnomalyEventAnalyticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeCostAnomalyEventAnalyticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeCostAnomalyEventAnalyticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeCostAnomalyEventAnalyticsResponse")
	}
	return
}

// summarizeCostAnomalyEventAnalytics implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) summarizeCostAnomalyEventAnalytics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/costAnomalyEventAnalytics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeCostAnomalyEventAnalyticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyEvent/SummarizeCostAnomalyEventAnalytics"
		err = common.PostProcessServiceError(err, "CostAd", "SummarizeCostAnomalyEventAnalytics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCostAlertSubscription Update a CostAlertSubscription identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/UpdateCostAlertSubscription.go.html to see an example of how to use UpdateCostAlertSubscription API.
// A default retry strategy applies to this operation UpdateCostAlertSubscription()
func (client CostAdClient) UpdateCostAlertSubscription(ctx context.Context, request UpdateCostAlertSubscriptionRequest) (response UpdateCostAlertSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCostAlertSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCostAlertSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCostAlertSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCostAlertSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCostAlertSubscriptionResponse")
	}
	return
}

// updateCostAlertSubscription implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) updateCostAlertSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/costAlertSubscriptions/{costAlertSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCostAlertSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAlertSubscription/UpdateCostAlertSubscription"
		err = common.PostProcessServiceError(err, "CostAd", "UpdateCostAlertSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCostAnomalyEvent Update a CostAnomalyEvent identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/UpdateCostAnomalyEvent.go.html to see an example of how to use UpdateCostAnomalyEvent API.
// A default retry strategy applies to this operation UpdateCostAnomalyEvent()
func (client CostAdClient) UpdateCostAnomalyEvent(ctx context.Context, request UpdateCostAnomalyEventRequest) (response UpdateCostAnomalyEventResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCostAnomalyEvent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCostAnomalyEventResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCostAnomalyEventResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCostAnomalyEventResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCostAnomalyEventResponse")
	}
	return
}

// updateCostAnomalyEvent implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) updateCostAnomalyEvent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/costAnomalyEvents/{costAnomalyEventId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCostAnomalyEventResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyEvent/UpdateCostAnomalyEvent"
		err = common.PostProcessServiceError(err, "CostAd", "UpdateCostAnomalyEvent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCostAnomalyMonitor Update a CostAnomalyMonitor identified by the OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/UpdateCostAnomalyMonitor.go.html to see an example of how to use UpdateCostAnomalyMonitor API.
// A default retry strategy applies to this operation UpdateCostAnomalyMonitor()
func (client CostAdClient) UpdateCostAnomalyMonitor(ctx context.Context, request UpdateCostAnomalyMonitorRequest) (response UpdateCostAnomalyMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCostAnomalyMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCostAnomalyMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCostAnomalyMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCostAnomalyMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCostAnomalyMonitorResponse")
	}
	return
}

// updateCostAnomalyMonitor implements the OCIOperation interface (enables retrying operations)
func (client CostAdClient) updateCostAnomalyMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/costAnomalyMonitors/{costAnomalyMonitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCostAnomalyMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/budgets/20190111/CostAnomalyMonitor/UpdateCostAnomalyMonitor"
		err = common.PostProcessServiceError(err, "CostAd", "UpdateCostAnomalyMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
