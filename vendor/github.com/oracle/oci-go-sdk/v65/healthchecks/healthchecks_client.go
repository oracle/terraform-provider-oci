// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// HealthChecksClient a client for HealthChecks
type HealthChecksClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewHealthChecksClientWithConfigurationProvider Creates a new default HealthChecks client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewHealthChecksClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client HealthChecksClient, err error) {
	if enabled := common.CheckForEnabledServices("healthchecks"); !enabled {
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
	return newHealthChecksClientFromBaseClient(baseClient, provider)
}

// NewHealthChecksClientWithOboToken Creates a new default HealthChecks client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewHealthChecksClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client HealthChecksClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newHealthChecksClientFromBaseClient(baseClient, configProvider)
}

func newHealthChecksClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client HealthChecksClient, err error) {
	// HealthChecks service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("HealthChecks"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = HealthChecksClient{BaseClient: baseClient}
	client.BasePath = "20180501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *HealthChecksClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("healthchecks", "https://healthchecks.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *HealthChecksClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *HealthChecksClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeHttpMonitorCompartment Moves a monitor into a different compartment. When provided, `If-Match` is checked
// against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ChangeHttpMonitorCompartment.go.html to see an example of how to use ChangeHttpMonitorCompartment API.
func (client HealthChecksClient) ChangeHttpMonitorCompartment(ctx context.Context, request ChangeHttpMonitorCompartmentRequest) (response ChangeHttpMonitorCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeHttpMonitorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeHttpMonitorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeHttpMonitorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeHttpMonitorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeHttpMonitorCompartmentResponse")
	}
	return
}

// changeHttpMonitorCompartment implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) changeHttpMonitorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/httpMonitors/{monitorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeHttpMonitorCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpMonitor/ChangeHttpMonitorCompartment"
		err = common.PostProcessServiceError(err, "HealthChecks", "ChangeHttpMonitorCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePingMonitorCompartment Moves a monitor into a different compartment. When provided, `If-Match` is checked
// against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ChangePingMonitorCompartment.go.html to see an example of how to use ChangePingMonitorCompartment API.
func (client HealthChecksClient) ChangePingMonitorCompartment(ctx context.Context, request ChangePingMonitorCompartmentRequest) (response ChangePingMonitorCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePingMonitorCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePingMonitorCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePingMonitorCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePingMonitorCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePingMonitorCompartmentResponse")
	}
	return
}

// changePingMonitorCompartment implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) changePingMonitorCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pingMonitors/{monitorId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePingMonitorCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingMonitor/ChangePingMonitorCompartment"
		err = common.PostProcessServiceError(err, "HealthChecks", "ChangePingMonitorCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateHttpMonitor Creates an HTTP monitor. Vantage points will be automatically selected if not specified,
// and probes will be initiated from each vantage point to each of the targets at the frequency
// specified by `intervalInSeconds`.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/CreateHttpMonitor.go.html to see an example of how to use CreateHttpMonitor API.
func (client HealthChecksClient) CreateHttpMonitor(ctx context.Context, request CreateHttpMonitorRequest) (response CreateHttpMonitorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createHttpMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateHttpMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateHttpMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateHttpMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateHttpMonitorResponse")
	}
	return
}

// createHttpMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) createHttpMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/httpMonitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateHttpMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpMonitor/CreateHttpMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "CreateHttpMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOnDemandHttpProbe Creates an on-demand HTTP probe. The location response header contains the URL for
// fetching the probe results.
// *Note:* On-demand probe configurations are not saved.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/CreateOnDemandHttpProbe.go.html to see an example of how to use CreateOnDemandHttpProbe API.
func (client HealthChecksClient) CreateOnDemandHttpProbe(ctx context.Context, request CreateOnDemandHttpProbeRequest) (response CreateOnDemandHttpProbeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createOnDemandHttpProbe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOnDemandHttpProbeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOnDemandHttpProbeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOnDemandHttpProbeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOnDemandHttpProbeResponse")
	}
	return
}

// createOnDemandHttpProbe implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) createOnDemandHttpProbe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/httpProbeResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOnDemandHttpProbeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpProbe/CreateOnDemandHttpProbe"
		err = common.PostProcessServiceError(err, "HealthChecks", "CreateOnDemandHttpProbe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOnDemandPingProbe Creates an on-demand ping probe. The location response header contains the URL for
// fetching probe results.
// *Note:* The on-demand probe configuration is not saved.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/CreateOnDemandPingProbe.go.html to see an example of how to use CreateOnDemandPingProbe API.
func (client HealthChecksClient) CreateOnDemandPingProbe(ctx context.Context, request CreateOnDemandPingProbeRequest) (response CreateOnDemandPingProbeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createOnDemandPingProbe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOnDemandPingProbeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOnDemandPingProbeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOnDemandPingProbeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOnDemandPingProbeResponse")
	}
	return
}

// createOnDemandPingProbe implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) createOnDemandPingProbe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pingProbeResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOnDemandPingProbeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingProbe/CreateOnDemandPingProbe"
		err = common.PostProcessServiceError(err, "HealthChecks", "CreateOnDemandPingProbe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePingMonitor Creates a ping monitor. Vantage points will be automatically selected if not specified,
// and probes will be initiated from each vantage point to each of the targets at the frequency
// specified by `intervalInSeconds`.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/CreatePingMonitor.go.html to see an example of how to use CreatePingMonitor API.
func (client HealthChecksClient) CreatePingMonitor(ctx context.Context, request CreatePingMonitorRequest) (response CreatePingMonitorResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPingMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePingMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePingMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePingMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePingMonitorResponse")
	}
	return
}

// createPingMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) createPingMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pingMonitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePingMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingMonitor/CreatePingMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "CreatePingMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteHttpMonitor Deletes the HTTP monitor and its configuration. All future probes of this
// monitor are stopped. Results associated with the monitor are not deleted.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/DeleteHttpMonitor.go.html to see an example of how to use DeleteHttpMonitor API.
func (client HealthChecksClient) DeleteHttpMonitor(ctx context.Context, request DeleteHttpMonitorRequest) (response DeleteHttpMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteHttpMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteHttpMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteHttpMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteHttpMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteHttpMonitorResponse")
	}
	return
}

// deleteHttpMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) deleteHttpMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/httpMonitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteHttpMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpMonitor/DeleteHttpMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "DeleteHttpMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePingMonitor Deletes the ping monitor and its configuration. All future probes of this
// monitor are stopped. Results associated with the monitor are not deleted.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/DeletePingMonitor.go.html to see an example of how to use DeletePingMonitor API.
func (client HealthChecksClient) DeletePingMonitor(ctx context.Context, request DeletePingMonitorRequest) (response DeletePingMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePingMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePingMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePingMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePingMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePingMonitorResponse")
	}
	return
}

// deletePingMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) deletePingMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/pingMonitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePingMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingMonitor/DeletePingMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "DeletePingMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetHttpMonitor Gets the configuration for the specified monitor.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/GetHttpMonitor.go.html to see an example of how to use GetHttpMonitor API.
func (client HealthChecksClient) GetHttpMonitor(ctx context.Context, request GetHttpMonitorRequest) (response GetHttpMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHttpMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHttpMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHttpMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHttpMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHttpMonitorResponse")
	}
	return
}

// getHttpMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) getHttpMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/httpMonitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetHttpMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpMonitor/GetHttpMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "GetHttpMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPingMonitor Gets the configuration for the specified ping monitor.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/GetPingMonitor.go.html to see an example of how to use GetPingMonitor API.
func (client HealthChecksClient) GetPingMonitor(ctx context.Context, request GetPingMonitorRequest) (response GetPingMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPingMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPingMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPingMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPingMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPingMonitorResponse")
	}
	return
}

// getPingMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) getPingMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pingMonitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPingMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingMonitor/GetPingMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "GetPingMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHealthChecksVantagePoints Gets information about all vantage points available to the user.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListHealthChecksVantagePoints.go.html to see an example of how to use ListHealthChecksVantagePoints API.
func (client HealthChecksClient) ListHealthChecksVantagePoints(ctx context.Context, request ListHealthChecksVantagePointsRequest) (response ListHealthChecksVantagePointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHealthChecksVantagePoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHealthChecksVantagePointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHealthChecksVantagePointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHealthChecksVantagePointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHealthChecksVantagePointsResponse")
	}
	return
}

// listHealthChecksVantagePoints implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) listHealthChecksVantagePoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vantagePoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHealthChecksVantagePointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HealthChecksVantagePointSummary/ListHealthChecksVantagePoints"
		err = common.PostProcessServiceError(err, "HealthChecks", "ListHealthChecksVantagePoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHttpMonitors Gets a list of HTTP monitors.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListHttpMonitors.go.html to see an example of how to use ListHttpMonitors API.
func (client HealthChecksClient) ListHttpMonitors(ctx context.Context, request ListHttpMonitorsRequest) (response ListHttpMonitorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHttpMonitors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHttpMonitorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHttpMonitorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHttpMonitorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHttpMonitorsResponse")
	}
	return
}

// listHttpMonitors implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) listHttpMonitors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/httpMonitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHttpMonitorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpMonitorSummary/ListHttpMonitors"
		err = common.PostProcessServiceError(err, "HealthChecks", "ListHttpMonitors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHttpProbeResults Gets the HTTP probe results for the specified probe or monitor, where
// the `probeConfigurationId` is the OCID of either a monitor or an
// on-demand probe.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListHttpProbeResults.go.html to see an example of how to use ListHttpProbeResults API.
func (client HealthChecksClient) ListHttpProbeResults(ctx context.Context, request ListHttpProbeResultsRequest) (response ListHttpProbeResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHttpProbeResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHttpProbeResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHttpProbeResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHttpProbeResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHttpProbeResultsResponse")
	}
	return
}

// listHttpProbeResults implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) listHttpProbeResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/httpProbeResults/{probeConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHttpProbeResultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpProbeResultSummary/ListHttpProbeResults"
		err = common.PostProcessServiceError(err, "HealthChecks", "ListHttpProbeResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPingMonitors Gets a list of configured ping monitors.
// Results are paginated based on `page` and `limit`.  The `opc-next-page` header provides
// a URL for fetching the next page.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListPingMonitors.go.html to see an example of how to use ListPingMonitors API.
func (client HealthChecksClient) ListPingMonitors(ctx context.Context, request ListPingMonitorsRequest) (response ListPingMonitorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPingMonitors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPingMonitorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPingMonitorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPingMonitorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPingMonitorsResponse")
	}
	return
}

// listPingMonitors implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) listPingMonitors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pingMonitors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPingMonitorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingMonitorSummary/ListPingMonitors"
		err = common.PostProcessServiceError(err, "HealthChecks", "ListPingMonitors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPingProbeResults Returns the results for the specified probe, where the `probeConfigurationId`
// is the OCID of either a monitor or an on-demand probe.
// Results are paginated based on `page` and `limit`.  The `opc-next-page` header provides
// a URL for fetching the next page.  Use `sortOrder` to set the order of the
// results.  If `sortOrder` is unspecified, results are sorted in ascending order by
// `startTime`.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListPingProbeResults.go.html to see an example of how to use ListPingProbeResults API.
func (client HealthChecksClient) ListPingProbeResults(ctx context.Context, request ListPingProbeResultsRequest) (response ListPingProbeResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPingProbeResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPingProbeResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPingProbeResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPingProbeResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPingProbeResultsResponse")
	}
	return
}

// listPingProbeResults implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) listPingProbeResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pingProbeResults/{probeConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPingProbeResultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingProbeResultSummary/ListPingProbeResults"
		err = common.PostProcessServiceError(err, "HealthChecks", "ListPingProbeResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateHttpMonitor Updates the configuration of the specified HTTP monitor. Only the fields
// specified in the request body will be updated; all other configuration
// properties will remain unchanged.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/UpdateHttpMonitor.go.html to see an example of how to use UpdateHttpMonitor API.
func (client HealthChecksClient) UpdateHttpMonitor(ctx context.Context, request UpdateHttpMonitorRequest) (response UpdateHttpMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateHttpMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateHttpMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateHttpMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateHttpMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateHttpMonitorResponse")
	}
	return
}

// updateHttpMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) updateHttpMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/httpMonitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateHttpMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/HttpMonitor/UpdateHttpMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "UpdateHttpMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePingMonitor Updates the configuration of the specified ping monitor. Only the fields
// specified in the request body will be updated; all other configuration properties
// will remain unchanged.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/UpdatePingMonitor.go.html to see an example of how to use UpdatePingMonitor API.
func (client HealthChecksClient) UpdatePingMonitor(ctx context.Context, request UpdatePingMonitorRequest) (response UpdatePingMonitorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePingMonitor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePingMonitorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePingMonitorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePingMonitorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePingMonitorResponse")
	}
	return
}

// updatePingMonitor implements the OCIOperation interface (enables retrying operations)
func (client HealthChecksClient) updatePingMonitor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/pingMonitors/{monitorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePingMonitorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/healthchecks/20180501/PingMonitor/UpdatePingMonitor"
		err = common.PostProcessServiceError(err, "HealthChecks", "UpdatePingMonitor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
