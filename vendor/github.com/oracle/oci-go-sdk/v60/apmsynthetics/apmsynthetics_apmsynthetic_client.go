// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"github.com/oracle/oci-go-sdk/v60/common/auth"
	"net/http"
)

//ApmSyntheticClient a client for ApmSynthetic
type ApmSyntheticClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewApmSyntheticClientWithConfigurationProvider Creates a new default ApmSynthetic client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewApmSyntheticClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ApmSyntheticClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newApmSyntheticClientFromBaseClient(baseClient, provider)
}

// NewApmSyntheticClientWithOboToken Creates a new default ApmSynthetic client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewApmSyntheticClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ApmSyntheticClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newApmSyntheticClientFromBaseClient(baseClient, configProvider)
}

func newApmSyntheticClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ApmSyntheticClient, err error) {
	// ApmSynthetic service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ApmSyntheticClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ApmSyntheticClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmsynthetics", "https://apm-synthetic.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ApmSyntheticClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ApmSyntheticClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateDedicatedVantagePoint Registers a new dedicated vantage point.
// A default retry strategy applies to this operation CreateDedicatedVantagePoint()
func (client ApmSyntheticClient) CreateDedicatedVantagePoint(ctx context.Context, request CreateDedicatedVantagePointRequest) (response CreateDedicatedVantagePointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDedicatedVantagePoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDedicatedVantagePointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDedicatedVantagePointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDedicatedVantagePointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDedicatedVantagePointResponse")
	}
	return
}

// createDedicatedVantagePoint implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) createDedicatedVantagePoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dedicatedVantagePoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDedicatedVantagePointResponse
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

// CreateMonitor Creates a new monitor.
// A default retry strategy applies to this operation CreateMonitor()
func (client ApmSyntheticClient) CreateMonitor(ctx context.Context, request CreateMonitorRequest) (response CreateMonitorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMonitorResponse")
	}
	return
}

// createMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) createMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMonitorResponse
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

// CreateScript Creates a new script.
// A default retry strategy applies to this operation CreateScript()
func (client ApmSyntheticClient) CreateScript(ctx context.Context, request CreateScriptRequest) (response CreateScriptResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScriptResponse")
	}
	return
}

// createScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) createScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/scripts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateScriptResponse
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

// DeleteDedicatedVantagePoint Deregisters the specified dedicated vantage point.
// A default retry strategy applies to this operation DeleteDedicatedVantagePoint()
func (client ApmSyntheticClient) DeleteDedicatedVantagePoint(ctx context.Context, request DeleteDedicatedVantagePointRequest) (response DeleteDedicatedVantagePointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDedicatedVantagePoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDedicatedVantagePointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDedicatedVantagePointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDedicatedVantagePointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDedicatedVantagePointResponse")
	}
	return
}

// deleteDedicatedVantagePoint implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) deleteDedicatedVantagePoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dedicatedVantagePoints/{dedicatedVantagePointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDedicatedVantagePointResponse
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

// DeleteMonitor Deletes the specified monitor.
// A default retry strategy applies to this operation DeleteMonitor()
func (client ApmSyntheticClient) DeleteMonitor(ctx context.Context, request DeleteMonitorRequest) (response DeleteMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMonitorResponse")
	}
	return
}

// deleteMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) deleteMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/monitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMonitorResponse
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

// DeleteScript Deletes the specified script.
// A default retry strategy applies to this operation DeleteScript()
func (client ApmSyntheticClient) DeleteScript(ctx context.Context, request DeleteScriptRequest) (response DeleteScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScriptResponse")
	}
	return
}

// deleteScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) deleteScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/scripts/{scriptId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteScriptResponse
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

// GetDedicatedVantagePoint Gets the details of the dedicated vantage point identified by the OCID.
// A default retry strategy applies to this operation GetDedicatedVantagePoint()
func (client ApmSyntheticClient) GetDedicatedVantagePoint(ctx context.Context, request GetDedicatedVantagePointRequest) (response GetDedicatedVantagePointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDedicatedVantagePoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDedicatedVantagePointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDedicatedVantagePointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDedicatedVantagePointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDedicatedVantagePointResponse")
	}
	return
}

// getDedicatedVantagePoint implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) getDedicatedVantagePoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dedicatedVantagePoints/{dedicatedVantagePointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDedicatedVantagePointResponse
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

// GetMonitor Gets the configuration of the monitor identified by the OCID.
// A default retry strategy applies to this operation GetMonitor()
func (client ApmSyntheticClient) GetMonitor(ctx context.Context, request GetMonitorRequest) (response GetMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitorResponse")
	}
	return
}

// getMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) getMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMonitorResponse
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

// GetMonitorResult Gets the results for a specific execution of a monitor identified by OCID. The results are in a HAR file, Screenshot, Console Log or Network details.
// A default retry strategy applies to this operation GetMonitorResult()
func (client ApmSyntheticClient) GetMonitorResult(ctx context.Context, request GetMonitorResultRequest) (response GetMonitorResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitorResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitorResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitorResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitorResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitorResultResponse")
	}
	return
}

// getMonitorResult implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) getMonitorResult(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitors/{monitorId}/results/{executionTime}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMonitorResultResponse
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

// GetScript Gets the configuration of the script identified by the OCID.
// A default retry strategy applies to this operation GetScript()
func (client ApmSyntheticClient) GetScript(ctx context.Context, request GetScriptRequest) (response GetScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScriptResponse")
	}
	return
}

// getScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) getScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scripts/{scriptId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScriptResponse
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

// ListDedicatedVantagePoints Returns a list of dedicated vantage points.
// A default retry strategy applies to this operation ListDedicatedVantagePoints()
func (client ApmSyntheticClient) ListDedicatedVantagePoints(ctx context.Context, request ListDedicatedVantagePointsRequest) (response ListDedicatedVantagePointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDedicatedVantagePoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDedicatedVantagePointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDedicatedVantagePointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDedicatedVantagePointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDedicatedVantagePointsResponse")
	}
	return
}

// listDedicatedVantagePoints implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) listDedicatedVantagePoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dedicatedVantagePoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDedicatedVantagePointsResponse
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

// ListMonitors Returns a list of monitors.
// A default retry strategy applies to this operation ListMonitors()
func (client ApmSyntheticClient) ListMonitors(ctx context.Context, request ListMonitorsRequest) (response ListMonitorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMonitors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMonitorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMonitorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMonitorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMonitorsResponse")
	}
	return
}

// listMonitors implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) listMonitors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMonitorsResponse
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

// ListPublicVantagePoints Returns a list of public vantage points.
// A default retry strategy applies to this operation ListPublicVantagePoints()
func (client ApmSyntheticClient) ListPublicVantagePoints(ctx context.Context, request ListPublicVantagePointsRequest) (response ListPublicVantagePointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublicVantagePoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPublicVantagePointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPublicVantagePointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublicVantagePointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublicVantagePointsResponse")
	}
	return
}

// listPublicVantagePoints implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) listPublicVantagePoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publicVantagePoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPublicVantagePointsResponse
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

// ListScripts Returns a list of scripts.
// A default retry strategy applies to this operation ListScripts()
func (client ApmSyntheticClient) ListScripts(ctx context.Context, request ListScriptsRequest) (response ListScriptsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScripts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScriptsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScriptsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScriptsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScriptsResponse")
	}
	return
}

// listScripts implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) listScripts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/scripts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScriptsResponse
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

// UpdateDedicatedVantagePoint Updates the dedicated vantage point.
// A default retry strategy applies to this operation UpdateDedicatedVantagePoint()
func (client ApmSyntheticClient) UpdateDedicatedVantagePoint(ctx context.Context, request UpdateDedicatedVantagePointRequest) (response UpdateDedicatedVantagePointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDedicatedVantagePoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDedicatedVantagePointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDedicatedVantagePointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDedicatedVantagePointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDedicatedVantagePointResponse")
	}
	return
}

// updateDedicatedVantagePoint implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) updateDedicatedVantagePoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dedicatedVantagePoints/{dedicatedVantagePointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDedicatedVantagePointResponse
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

// UpdateMonitor Updates the monitor.
// A default retry strategy applies to this operation UpdateMonitor()
func (client ApmSyntheticClient) UpdateMonitor(ctx context.Context, request UpdateMonitorRequest) (response UpdateMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMonitorResponse")
	}
	return
}

// updateMonitor implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) updateMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/monitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMonitorResponse
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

// UpdateScript Updates the script.
// A default retry strategy applies to this operation UpdateScript()
func (client ApmSyntheticClient) UpdateScript(ctx context.Context, request UpdateScriptRequest) (response UpdateScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScriptResponse")
	}
	return
}

// updateScript implements the OCIOperation interface (enables retrying operations)
func (client ApmSyntheticClient) updateScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/scripts/{scriptId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateScriptResponse
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
