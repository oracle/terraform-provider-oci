// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For information about monitoring, see Monitoring Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm).
//

package monitoring

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//MonitoringClient a client for Monitoring
type MonitoringClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMonitoringClientWithConfigurationProvider Creates a new default Monitoring client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMonitoringClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MonitoringClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newMonitoringClientFromBaseClient(baseClient, configProvider)
}

// NewMonitoringClientWithOboToken Creates a new default Monitoring client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewMonitoringClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MonitoringClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newMonitoringClientFromBaseClient(baseClient, configProvider)
}

func newMonitoringClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MonitoringClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *MonitoringClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeAlarmCompartment Moves an alarm into a different compartment within the same tenancy.
// For information about moving resources between compartments, see Moving Resources Between Compartments (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
func (client MonitoringClient) ChangeAlarmCompartment(ctx context.Context, request ChangeAlarmCompartmentRequest) (response ChangeAlarmCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) changeAlarmCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms/{alarmId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeAlarmCompartmentResponse
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

// CreateAlarm Creates a new alarm in the specified compartment.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) CreateAlarm(ctx context.Context, request CreateAlarmRequest) (response CreateAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) createAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms")
	if err != nil {
		return nil, err
	}

	var response CreateAlarmResponse
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

// DeleteAlarm Deletes the specified alarm.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) DeleteAlarm(ctx context.Context, request DeleteAlarmRequest) (response DeleteAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) deleteAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/alarms/{alarmId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAlarmResponse
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

// GetAlarm Gets the specified alarm.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) GetAlarm(ctx context.Context, request GetAlarmRequest) (response GetAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) getAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/{alarmId}")
	if err != nil {
		return nil, err
	}

	var response GetAlarmResponse
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

// GetAlarmHistory Get the history of the specified alarm.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) GetAlarmHistory(ctx context.Context, request GetAlarmHistoryRequest) (response GetAlarmHistoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) getAlarmHistory(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/{alarmId}/history")
	if err != nil {
		return nil, err
	}

	var response GetAlarmHistoryResponse
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

// ListAlarms Lists the alarms for the specified compartment.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) ListAlarms(ctx context.Context, request ListAlarmsRequest) (response ListAlarmsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) listAlarms(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms")
	if err != nil {
		return nil, err
	}

	var response ListAlarmsResponse
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

// ListAlarmsStatus List the status of each alarm in the specified compartment.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) ListAlarmsStatus(ctx context.Context, request ListAlarmsStatusRequest) (response ListAlarmsStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) listAlarmsStatus(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/alarms/status")
	if err != nil {
		return nil, err
	}

	var response ListAlarmsStatusResponse
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

// ListMetrics Returns metric definitions that match the criteria specified in the request. Compartment OCID required.
// For information about metrics, see Metrics Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#MetricsOverview).
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// Transactions Per Second (TPS) per-tenancy limit for this operation: 10.
func (client MonitoringClient) ListMetrics(ctx context.Context, request ListMetricsRequest) (response ListMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) listMetrics(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics/actions/listMetrics")
	if err != nil {
		return nil, err
	}

	var response ListMetricsResponse
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

// PostMetricData Publishes raw metric data points to the Monitoring service.
// For more information about publishing metrics, see Publishing Custom Metrics (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/publishingcustommetrics.htm).
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// Per-call limits information follows.
// * Dimensions per metric group*. Maximum: 20. Minimum: 1.
// * Unique metric streams*. Maximum: 50.
// * Transactions Per Second (TPS) per-tenancy limit for this operation: 50.
// *A metric group is the combination of a given metric, metric namespace, and tenancy for the purpose of determining limits.
// A dimension is a qualifier provided in a metric definition.
// A metric stream is an individual set of aggregated data for a metric, typically specific to a resource.
// For more information about metric-related concepts, see Monitoring Concepts (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#concepts).
// The endpoints for this operation differ from other Monitoring operations. Replace the string `telemetry` with `telemetry-ingestion` in the endpoint, as in the following example:
// https://telemetry-ingestion.eu-frankfurt-1.oraclecloud.com
func (client MonitoringClient) PostMetricData(ctx context.Context, request PostMetricDataRequest) (response PostMetricDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) postMetricData(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics")
	if err != nil {
		return nil, err
	}

	var response PostMetricDataResponse
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

// RemoveAlarmSuppression Removes any existing suppression for the specified alarm.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) RemoveAlarmSuppression(ctx context.Context, request RemoveAlarmSuppressionRequest) (response RemoveAlarmSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) removeAlarmSuppression(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/alarms/{alarmId}/actions/removeSuppression")
	if err != nil {
		return nil, err
	}

	var response RemoveAlarmSuppressionResponse
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

// SummarizeMetricsData Returns aggregated data that match the criteria specified in the request. Compartment OCID required.
// For information on metric queries, see Building Metric Queries (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/buildingqueries.htm).
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// Transactions Per Second (TPS) per-tenancy limit for this operation: 10.
func (client MonitoringClient) SummarizeMetricsData(ctx context.Context, request SummarizeMetricsDataRequest) (response SummarizeMetricsDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) summarizeMetricsData(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metrics/actions/summarizeMetricsData")
	if err != nil {
		return nil, err
	}

	var response SummarizeMetricsDataResponse
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

// UpdateAlarm Updates the specified alarm.
// For important limits information, see Limits on Monitoring (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
// or transactions, per second (TPS) for a given tenancy.
func (client MonitoringClient) UpdateAlarm(ctx context.Context, request UpdateAlarmRequest) (response UpdateAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MonitoringClient) updateAlarm(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/alarms/{alarmId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAlarmResponse
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
