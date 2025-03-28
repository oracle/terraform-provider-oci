// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MonitoringClient a client for Monitoring
type MonitoringClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMonitoringClientWithConfigurationProvider Creates a new default Monitoring client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMonitoringClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MonitoringClient, err error) {
	if enabled := common.CheckForEnabledServices("monitoring"); !enabled {
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
	return newMonitoringClientFromBaseClient(baseClient, provider)
}

// NewMonitoringClientWithOboToken Creates a new default Monitoring client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMonitoringClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MonitoringClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMonitoringClientFromBaseClient(baseClient, configProvider)
}

func newMonitoringClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MonitoringClient, err error) {
	// Monitoring service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Monitoring"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MonitoringClient{BaseClient: baseClient}
	client.BasePath = "20180401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MonitoringClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("telemetry", "https://telemetry.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MonitoringClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MonitoringClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAlarmCompartment Moves an alarm into a different compartment within the same tenancy.
// For more information, see
// Moving an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/change-compartment-alarm.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/ChangeAlarmCompartment.go.html to see an example of how to use ChangeAlarmCompartment API.
func (client MonitoringClient) ChangeAlarmCompartment(ctx context.Context, request ChangeAlarmCompartmentRequest) (response ChangeAlarmCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeAlarmCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAlarmCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAlarmCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAlarmCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAlarmCompartmentResponse")
	}
	return
}

// changeAlarmCompartment implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) changeAlarmCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms/{alarmId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAlarmCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/Alarm/ChangeAlarmCompartment"
		err = common.PostProcessServiceError(err, "Monitoring", "ChangeAlarmCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAlarm Creates a new alarm in the specified compartment.
// For more information, see
// Creating an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/create-alarm.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/CreateAlarm.go.html to see an example of how to use CreateAlarm API.
func (client MonitoringClient) CreateAlarm(ctx context.Context, request CreateAlarmRequest) (response CreateAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAlarmResponse")
	}
	return
}

// createAlarm implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) createAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/Alarm/CreateAlarm"
		err = common.PostProcessServiceError(err, "Monitoring", "CreateAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAlarmSuppression Creates a new alarm suppression at the specified level (alarm-wide or dimension-specific).
// For more information, see
// Adding an Alarm-wide Suppression (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/add-alarm-suppression.htm) and
// Adding a Dimension-Specific Alarm Suppression (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/create-alarm-suppression.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/CreateAlarmSuppression.go.html to see an example of how to use CreateAlarmSuppression API.
func (client MonitoringClient) CreateAlarmSuppression(ctx context.Context, request CreateAlarmSuppressionRequest) (response CreateAlarmSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAlarmSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAlarmSuppressionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAlarmSuppressionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAlarmSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAlarmSuppressionResponse")
	}
	return
}

// createAlarmSuppression implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) createAlarmSuppression(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarmSuppressions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAlarmSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmSuppression/CreateAlarmSuppression"
		err = common.PostProcessServiceError(err, "Monitoring", "CreateAlarmSuppression", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAlarm Deletes the specified alarm.
// For more information, see
// Deleting an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/delete-alarm.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/DeleteAlarm.go.html to see an example of how to use DeleteAlarm API.
func (client MonitoringClient) DeleteAlarm(ctx context.Context, request DeleteAlarmRequest) (response DeleteAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAlarmResponse")
	}
	return
}

// deleteAlarm implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) deleteAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/alarms/{alarmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/Alarm/DeleteAlarm"
		err = common.PostProcessServiceError(err, "Monitoring", "DeleteAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAlarmSuppression Deletes the specified alarm suppression. For more information, see
// Removing an Alarm-wide Suppression (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/remove-alarm-suppression.htm) and
// Removing a Dimension-Specific Alarm Suppression (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/delete-alarm-suppression.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/DeleteAlarmSuppression.go.html to see an example of how to use DeleteAlarmSuppression API.
func (client MonitoringClient) DeleteAlarmSuppression(ctx context.Context, request DeleteAlarmSuppressionRequest) (response DeleteAlarmSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAlarmSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAlarmSuppressionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAlarmSuppressionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAlarmSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAlarmSuppressionResponse")
	}
	return
}

// deleteAlarmSuppression implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) deleteAlarmSuppression(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/alarmSuppressions/{alarmSuppressionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAlarmSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmSuppression/DeleteAlarmSuppression"
		err = common.PostProcessServiceError(err, "Monitoring", "DeleteAlarmSuppression", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlarm Gets the specified alarm.
// For more information, see
// Getting an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/get-alarm.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/GetAlarm.go.html to see an example of how to use GetAlarm API.
func (client MonitoringClient) GetAlarm(ctx context.Context, request GetAlarmRequest) (response GetAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlarmResponse")
	}
	return
}

// getAlarm implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) getAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/{alarmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/Alarm/GetAlarm"
		err = common.PostProcessServiceError(err, "Monitoring", "GetAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlarmHistory Get the history of the specified alarm.
// For more information, see
// Getting History of an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/get-alarm-history.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/GetAlarmHistory.go.html to see an example of how to use GetAlarmHistory API.
func (client MonitoringClient) GetAlarmHistory(ctx context.Context, request GetAlarmHistoryRequest) (response GetAlarmHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlarmHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAlarmHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAlarmHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlarmHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlarmHistoryResponse")
	}
	return
}

// getAlarmHistory implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) getAlarmHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/{alarmId}/history", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAlarmHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmHistoryCollection/GetAlarmHistory"
		err = common.PostProcessServiceError(err, "Monitoring", "GetAlarmHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAlarmSuppression Gets the specified alarm suppression. For more information, see
// Getting an Alarm-wide Suppression (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/get-alarm-suppression.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/GetAlarmSuppression.go.html to see an example of how to use GetAlarmSuppression API.
func (client MonitoringClient) GetAlarmSuppression(ctx context.Context, request GetAlarmSuppressionRequest) (response GetAlarmSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAlarmSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAlarmSuppressionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAlarmSuppressionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAlarmSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAlarmSuppressionResponse")
	}
	return
}

// getAlarmSuppression implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) getAlarmSuppression(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarmSuppressions/{alarmSuppressionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAlarmSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmSuppression/GetAlarmSuppression"
		err = common.PostProcessServiceError(err, "Monitoring", "GetAlarmSuppression", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlarmSuppressions Lists alarm suppressions for the specified alarm. For more information, see
// Listing Alarm Suppressions (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/list-alarm-suppression.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/ListAlarmSuppressions.go.html to see an example of how to use ListAlarmSuppressions API.
func (client MonitoringClient) ListAlarmSuppressions(ctx context.Context, request ListAlarmSuppressionsRequest) (response ListAlarmSuppressionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlarmSuppressions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlarmSuppressionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlarmSuppressionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlarmSuppressionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlarmSuppressionsResponse")
	}
	return
}

// listAlarmSuppressions implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) listAlarmSuppressions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarmSuppressions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlarmSuppressionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmSuppressionCollection/ListAlarmSuppressions"
		err = common.PostProcessServiceError(err, "Monitoring", "ListAlarmSuppressions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlarms Lists the alarms for the specified compartment.
// For more information, see
// Listing Alarms (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/list-alarm.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/ListAlarms.go.html to see an example of how to use ListAlarms API.
func (client MonitoringClient) ListAlarms(ctx context.Context, request ListAlarmsRequest) (response ListAlarmsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlarms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlarmsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlarmsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlarmsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlarmsResponse")
	}
	return
}

// listAlarms implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) listAlarms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlarmsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmSummary/ListAlarms"
		err = common.PostProcessServiceError(err, "Monitoring", "ListAlarms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAlarmsStatus List the status of each alarm in the specified compartment.
// Status is collective, across all metric streams in the alarm.
// To list alarm status for each metric stream, use RetrieveDimensionStates.
// Optionally filter by resource or status value.
// For more information, see
// Listing Alarm Statuses (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/list-alarm-status.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/ListAlarmsStatus.go.html to see an example of how to use ListAlarmsStatus API.
// A default retry strategy applies to this operation ListAlarmsStatus()
func (client MonitoringClient) ListAlarmsStatus(ctx context.Context, request ListAlarmsStatusRequest) (response ListAlarmsStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAlarmsStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAlarmsStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAlarmsStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAlarmsStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAlarmsStatusResponse")
	}
	return
}

// listAlarmsStatus implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) listAlarmsStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/status", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAlarmsStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmStatusSummary/ListAlarmsStatus"
		err = common.PostProcessServiceError(err, "Monitoring", "ListAlarmsStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMetrics Returns metric definitions that match the criteria specified in the request. Compartment OCID required.
// For more information, see
// Listing Metric Definitions (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/list-metric.htm).
// For information about metrics, see
// Metrics Overview (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#MetricsOverview).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// Transactions Per Second (TPS) per-tenancy limit for this operation: 10.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/ListMetrics.go.html to see an example of how to use ListMetrics API.
func (client MonitoringClient) ListMetrics(ctx context.Context, request ListMetricsRequest) (response ListMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMetricsResponse")
	}
	return
}

// listMetrics implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) listMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics/actions/listMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/Metric/ListMetrics"
		err = common.PostProcessServiceError(err, "Monitoring", "ListMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PostMetricData Publishes raw metric data points to the Monitoring service.
// For a data point to be posted, its timestamp must be near current time (less than two hours in the past and less than 10 minutes in the future).
// For more information about publishing metrics, see
// Publishing Custom Metrics (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/publishingcustommetrics.htm)
// and
// Custom Metrics Walkthrough (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/custom-metrics-walkthrough.htm).
// For information about developing a metric-posting client, see
// Developer Guide (https://docs.oracle.com/iaas/Content/API/Concepts/devtoolslanding.htm).
// For an example client, see
// MonitoringMetricPostExample.java (https://github.com/oracle/oci-java-sdk/blob/master/bmc-examples/src/main/java/MonitoringMetricPostExample.java).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// Per-call limits information follows.
// * Dimensions per metric group*. Maximum: 20. Minimum: 1.
// * Unique metric streams*. Maximum: 50.
// * Transactions Per Second (TPS) per-tenancy limit for this operation: 50.
// *A metric group is the combination of a given metric, metric namespace, and tenancy for the purpose of determining limits.
// A dimension is a qualifier provided in a metric definition.
// A metric stream is an individual set of aggregated data for a metric with zero or more dimension values.
// For more information about metric-related concepts, see
// Monitoring Concepts (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#concepts).
// **Note:** The endpoints for this operation differ from other Monitoring operations. Replace the string `telemetry` with `telemetry-ingestion` in the endpoint, as in the following example:
// https://telemetry-ingestion.eu-frankfurt-1.oraclecloud.com
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/PostMetricData.go.html to see an example of how to use PostMetricData API.
func (client MonitoringClient) PostMetricData(ctx context.Context, request PostMetricDataRequest) (response PostMetricDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.postMetricData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PostMetricDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PostMetricDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PostMetricDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PostMetricDataResponse")
	}
	return
}

// postMetricData implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) postMetricData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PostMetricDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/MetricData/PostMetricData"
		err = common.PostProcessServiceError(err, "Monitoring", "PostMetricData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAlarmSuppression Removes any existing suppression for the specified alarm.
// For more information, see
// Removing Suppression from an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/remove-alarm-suppression.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/RemoveAlarmSuppression.go.html to see an example of how to use RemoveAlarmSuppression API.
func (client MonitoringClient) RemoveAlarmSuppression(ctx context.Context, request RemoveAlarmSuppressionRequest) (response RemoveAlarmSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeAlarmSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveAlarmSuppressionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveAlarmSuppressionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveAlarmSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveAlarmSuppressionResponse")
	}
	return
}

// removeAlarmSuppression implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) removeAlarmSuppression(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms/{alarmId}/actions/removeSuppression", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveAlarmSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/Suppression/RemoveAlarmSuppression"
		err = common.PostProcessServiceError(err, "Monitoring", "RemoveAlarmSuppression", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveDimensionStates Lists the current alarm status of each metric stream, where status is derived from the metric stream's last associated transition.
// Optionally filter by status value and one or more dimension key-value pairs.
// For more information, see
// Listing Metric Stream Status in an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/list-alarm-status-metric-stream.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/RetrieveDimensionStates.go.html to see an example of how to use RetrieveDimensionStates API.
func (client MonitoringClient) RetrieveDimensionStates(ctx context.Context, request RetrieveDimensionStatesRequest) (response RetrieveDimensionStatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveDimensionStates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveDimensionStatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveDimensionStatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveDimensionStatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveDimensionStatesResponse")
	}
	return
}

// retrieveDimensionStates implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) retrieveDimensionStates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms/{alarmId}/actions/retrieveDimensionStates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveDimensionStatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmDimensionStatesCollection/RetrieveDimensionStates"
		err = common.PostProcessServiceError(err, "Monitoring", "RetrieveDimensionStates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeAlarmSuppressionHistory Returns history of suppressions for the specified alarm, including both dimension-specific and and alarm-wide suppressions. For more information, see
// Getting Suppression History for an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/summarize-alarm-suppression-history.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/SummarizeAlarmSuppressionHistory.go.html to see an example of how to use SummarizeAlarmSuppressionHistory API.
func (client MonitoringClient) SummarizeAlarmSuppressionHistory(ctx context.Context, request SummarizeAlarmSuppressionHistoryRequest) (response SummarizeAlarmSuppressionHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeAlarmSuppressionHistory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeAlarmSuppressionHistoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeAlarmSuppressionHistoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeAlarmSuppressionHistoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeAlarmSuppressionHistoryResponse")
	}
	return
}

// summarizeAlarmSuppressionHistory implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) summarizeAlarmSuppressionHistory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms/{alarmId}/actions/summarizeAlarmSuppressionHistory", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeAlarmSuppressionHistoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/AlarmSuppression/SummarizeAlarmSuppressionHistory"
		err = common.PostProcessServiceError(err, "Monitoring", "SummarizeAlarmSuppressionHistory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeMetricsData Returns aggregated data that match the criteria specified in the request. Compartment OCID required.
// For more information, see
// Querying Metric Data (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/query-metric-landing.htm)
// and
// Creating a Query (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/query-metric.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// Transactions Per Second (TPS) per-tenancy limit for this operation: 10.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/SummarizeMetricsData.go.html to see an example of how to use SummarizeMetricsData API.
func (client MonitoringClient) SummarizeMetricsData(ctx context.Context, request SummarizeMetricsDataRequest) (response SummarizeMetricsDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeMetricsData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeMetricsDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeMetricsDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeMetricsDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeMetricsDataResponse")
	}
	return
}

// summarizeMetricsData implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) summarizeMetricsData(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics/actions/summarizeMetricsData", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeMetricsDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/MetricData/SummarizeMetricsData"
		err = common.PostProcessServiceError(err, "Monitoring", "SummarizeMetricsData", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAlarm Updates the specified alarm.
// For more information, see
// Updating an Alarm (https://docs.oracle.com/iaas/Content/Monitoring/Tasks/update-alarm.htm).
// For important limits information, see
// Limits on Monitoring (https://docs.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/UpdateAlarm.go.html to see an example of how to use UpdateAlarm API.
func (client MonitoringClient) UpdateAlarm(ctx context.Context, request UpdateAlarmRequest) (response UpdateAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAlarmResponse")
	}
	return
}

// updateAlarm implements the OCIOperation interface (enables retrying operations)
func (client MonitoringClient) updateAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/alarms/{alarmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/monitoring/20180401/Alarm/UpdateAlarm"
		err = common.PostProcessServiceError(err, "Monitoring", "UpdateAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
