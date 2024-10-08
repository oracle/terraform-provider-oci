// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// StackMonitoringClient a client for StackMonitoring
type StackMonitoringClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewStackMonitoringClientWithConfigurationProvider Creates a new default StackMonitoring client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewStackMonitoringClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client StackMonitoringClient, err error) {
	if enabled := common.CheckForEnabledServices("stackmonitoring"); !enabled {
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
	return newStackMonitoringClientFromBaseClient(baseClient, provider)
}

// NewStackMonitoringClientWithOboToken Creates a new default StackMonitoring client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewStackMonitoringClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client StackMonitoringClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newStackMonitoringClientFromBaseClient(baseClient, configProvider)
}

func newStackMonitoringClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client StackMonitoringClient, err error) {
	// StackMonitoring service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("StackMonitoring"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = StackMonitoringClient{BaseClient: baseClient}
	client.BasePath = "20210330"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *StackMonitoringClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("stackmonitoring", "https://stack-monitoring.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *StackMonitoringClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *StackMonitoringClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AssociateMonitoredResources Create an association between two monitored resources. Associations can be created
// between resources from different compartments as long they are in same tenancy.
// User should have required access in both the compartments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/AssociateMonitoredResources.go.html to see an example of how to use AssociateMonitoredResources API.
func (client StackMonitoringClient) AssociateMonitoredResources(ctx context.Context, request AssociateMonitoredResourcesRequest) (response AssociateMonitoredResourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.associateMonitoredResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AssociateMonitoredResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AssociateMonitoredResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AssociateMonitoredResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AssociateMonitoredResourcesResponse")
	}
	return
}

// associateMonitoredResources implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) associateMonitoredResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/actions/associateMonitoredResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AssociateMonitoredResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/AssociateMonitoredResources"
		err = common.PostProcessServiceError(err, "StackMonitoring", "AssociateMonitoredResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeConfigCompartment Moves the configuration item to another compartment.
// Basically, this will disable any configuration for this configuration type in thie compartment,
// and will enable it in the new one.
// For example, if for a HOST resource type, the configuration with AUTO_PROMOTE in the configuration type
// and TRUE as value is moved, automatic discovery will not take place in this compartment any more, but in the new one.
// So this operation will have the same effect as deleting the configuration item in the old compartment and
// recreating it in another compartment.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ChangeConfigCompartment.go.html to see an example of how to use ChangeConfigCompartment API.
// A default retry strategy applies to this operation ChangeConfigCompartment()
func (client StackMonitoringClient) ChangeConfigCompartment(ctx context.Context, request ChangeConfigCompartmentRequest) (response ChangeConfigCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeConfigCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeConfigCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeConfigCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeConfigCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeConfigCompartmentResponse")
	}
	return
}

// changeConfigCompartment implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) changeConfigCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configs/{configId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeConfigCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/Config/ChangeConfigCompartment"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ChangeConfigCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMetricExtensionCompartment Moves a Metric Extension resource from one compartment identifier to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ChangeMetricExtensionCompartment.go.html to see an example of how to use ChangeMetricExtensionCompartment API.
func (client StackMonitoringClient) ChangeMetricExtensionCompartment(ctx context.Context, request ChangeMetricExtensionCompartmentRequest) (response ChangeMetricExtensionCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMetricExtensionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMetricExtensionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMetricExtensionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMetricExtensionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMetricExtensionCompartmentResponse")
	}
	return
}

// changeMetricExtensionCompartment implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) changeMetricExtensionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metricExtensions/{metricExtensionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMetricExtensionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/ChangeMetricExtensionCompartment"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ChangeMetricExtensionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMonitoredResourceCompartment Moves a monitored resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ChangeMonitoredResourceCompartment.go.html to see an example of how to use ChangeMonitoredResourceCompartment API.
func (client StackMonitoringClient) ChangeMonitoredResourceCompartment(ctx context.Context, request ChangeMonitoredResourceCompartmentRequest) (response ChangeMonitoredResourceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMonitoredResourceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMonitoredResourceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMonitoredResourceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMonitoredResourceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMonitoredResourceCompartmentResponse")
	}
	return
}

// changeMonitoredResourceCompartment implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) changeMonitoredResourceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/{monitoredResourceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMonitoredResourceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/ChangeMonitoredResourceCompartment"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ChangeMonitoredResourceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMonitoredResourceTaskCompartment Moves a stack monitoring resource task from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ChangeMonitoredResourceTaskCompartment.go.html to see an example of how to use ChangeMonitoredResourceTaskCompartment API.
// A default retry strategy applies to this operation ChangeMonitoredResourceTaskCompartment()
func (client StackMonitoringClient) ChangeMonitoredResourceTaskCompartment(ctx context.Context, request ChangeMonitoredResourceTaskCompartmentRequest) (response ChangeMonitoredResourceTaskCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMonitoredResourceTaskCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMonitoredResourceTaskCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMonitoredResourceTaskCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMonitoredResourceTaskCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMonitoredResourceTaskCompartmentResponse")
	}
	return
}

// changeMonitoredResourceTaskCompartment implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) changeMonitoredResourceTaskCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResourceTasks/{monitoredResourceTaskId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMonitoredResourceTaskCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceTask/ChangeMonitoredResourceTaskCompartment"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ChangeMonitoredResourceTaskCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeProcessSetCompartment Moves a ProcessSet resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ChangeProcessSetCompartment.go.html to see an example of how to use ChangeProcessSetCompartment API.
// A default retry strategy applies to this operation ChangeProcessSetCompartment()
func (client StackMonitoringClient) ChangeProcessSetCompartment(ctx context.Context, request ChangeProcessSetCompartmentRequest) (response ChangeProcessSetCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeProcessSetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeProcessSetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeProcessSetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeProcessSetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeProcessSetCompartmentResponse")
	}
	return
}

// changeProcessSetCompartment implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) changeProcessSetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/processSets/{processSetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeProcessSetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/ProcessSet/ChangeProcessSetCompartment"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ChangeProcessSetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBaselineableMetric Creates the specified Baseline-able metric
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateBaselineableMetric.go.html to see an example of how to use CreateBaselineableMetric API.
// A default retry strategy applies to this operation CreateBaselineableMetric()
func (client StackMonitoringClient) CreateBaselineableMetric(ctx context.Context, request CreateBaselineableMetricRequest) (response CreateBaselineableMetricResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBaselineableMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBaselineableMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBaselineableMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBaselineableMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBaselineableMetricResponse")
	}
	return
}

// createBaselineableMetric implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createBaselineableMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/baselineableMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBaselineableMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/BaselineableMetric/CreateBaselineableMetric"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateBaselineableMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConfig Creates a configuration item, for example to define
// whether resources of a specific type should be discovered automatically.
// For example, when a new Management Agent gets registered in a certain compartment,
// this Management Agent can potentially get promoted to a HOST resource.
// The configuration item will determine if HOST resources in the selected compartment will be
// discovered automatically.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateConfig.go.html to see an example of how to use CreateConfig API.
// A default retry strategy applies to this operation CreateConfig()
func (client StackMonitoringClient) CreateConfig(ctx context.Context, request CreateConfigRequest) (response CreateConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConfigResponse")
	}
	return
}

// createConfig implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/Config/CreateConfig"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &config{})
	return response, err
}

// CreateDiscoveryJob API to create discovery Job and submit discovery Details to agent.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateDiscoveryJob.go.html to see an example of how to use CreateDiscoveryJob API.
func (client StackMonitoringClient) CreateDiscoveryJob(ctx context.Context, request CreateDiscoveryJobRequest) (response CreateDiscoveryJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDiscoveryJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDiscoveryJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDiscoveryJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDiscoveryJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDiscoveryJobResponse")
	}
	return
}

// createDiscoveryJob implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createDiscoveryJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/discoveryJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDiscoveryJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/DiscoveryJob/CreateDiscoveryJob"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateDiscoveryJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMaintenanceWindow Creates a new Maintenance Window for the given resources. It will create also the
// Alarms Suppression for each alarm that the resource migth trigger.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateMaintenanceWindow.go.html to see an example of how to use CreateMaintenanceWindow API.
// A default retry strategy applies to this operation CreateMaintenanceWindow()
func (client StackMonitoringClient) CreateMaintenanceWindow(ctx context.Context, request CreateMaintenanceWindowRequest) (response CreateMaintenanceWindowResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMaintenanceWindowResponse")
	}
	return
}

// createMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maintenanceWindows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MaintenanceWindow/CreateMaintenanceWindow"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMetricExtension Creates a new metric extension resource for a given compartment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateMetricExtension.go.html to see an example of how to use CreateMetricExtension API.
func (client StackMonitoringClient) CreateMetricExtension(ctx context.Context, request CreateMetricExtensionRequest) (response CreateMetricExtensionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMetricExtensionResponse")
	}
	return
}

// createMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metricExtensions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/CreateMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMonitoredResource Creates a new monitored resource for the given resource type with the details and submits
// a work request for promoting the resource to agent. Once the resource is successfully
// added to agent, resource state will be marked active.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateMonitoredResource.go.html to see an example of how to use CreateMonitoredResource API.
func (client StackMonitoringClient) CreateMonitoredResource(ctx context.Context, request CreateMonitoredResourceRequest) (response CreateMonitoredResourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMonitoredResource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMonitoredResourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMonitoredResourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMonitoredResourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMonitoredResourceResponse")
	}
	return
}

// createMonitoredResource implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createMonitoredResource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMonitoredResourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/CreateMonitoredResource"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateMonitoredResource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMonitoredResourceTask Create a new stack monitoring resource task.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateMonitoredResourceTask.go.html to see an example of how to use CreateMonitoredResourceTask API.
// A default retry strategy applies to this operation CreateMonitoredResourceTask()
func (client StackMonitoringClient) CreateMonitoredResourceTask(ctx context.Context, request CreateMonitoredResourceTaskRequest) (response CreateMonitoredResourceTaskResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMonitoredResourceTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMonitoredResourceTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMonitoredResourceTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMonitoredResourceTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMonitoredResourceTaskResponse")
	}
	return
}

// createMonitoredResourceTask implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createMonitoredResourceTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResourceTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMonitoredResourceTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceTask/CreateMonitoredResourceTask"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateMonitoredResourceTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMonitoredResourceType Creates a new monitored resource type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateMonitoredResourceType.go.html to see an example of how to use CreateMonitoredResourceType API.
// A default retry strategy applies to this operation CreateMonitoredResourceType()
func (client StackMonitoringClient) CreateMonitoredResourceType(ctx context.Context, request CreateMonitoredResourceTypeRequest) (response CreateMonitoredResourceTypeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMonitoredResourceType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMonitoredResourceTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMonitoredResourceTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMonitoredResourceTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMonitoredResourceTypeResponse")
	}
	return
}

// createMonitoredResourceType implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createMonitoredResourceType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResourceTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMonitoredResourceTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceType/CreateMonitoredResourceType"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateMonitoredResourceType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProcessSet API to create Process Set.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/CreateProcessSet.go.html to see an example of how to use CreateProcessSet API.
// A default retry strategy applies to this operation CreateProcessSet()
func (client StackMonitoringClient) CreateProcessSet(ctx context.Context, request CreateProcessSetRequest) (response CreateProcessSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProcessSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProcessSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProcessSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProcessSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProcessSetResponse")
	}
	return
}

// createProcessSet implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) createProcessSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/processSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProcessSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/ProcessSet/CreateProcessSet"
		err = common.PostProcessServiceError(err, "StackMonitoring", "CreateProcessSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBaselineableMetric Deletes the Baseline-able metric for the given id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteBaselineableMetric.go.html to see an example of how to use DeleteBaselineableMetric API.
// A default retry strategy applies to this operation DeleteBaselineableMetric()
func (client StackMonitoringClient) DeleteBaselineableMetric(ctx context.Context, request DeleteBaselineableMetricRequest) (response DeleteBaselineableMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBaselineableMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBaselineableMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBaselineableMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBaselineableMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBaselineableMetricResponse")
	}
	return
}

// deleteBaselineableMetric implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteBaselineableMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/baselineableMetrics/{baselineableMetricId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBaselineableMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/BaselineableMetric/DeleteBaselineableMetric"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteBaselineableMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConfig Deletes a configuration identified by the id.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteConfig.go.html to see an example of how to use DeleteConfig API.
// A default retry strategy applies to this operation DeleteConfig()
func (client StackMonitoringClient) DeleteConfig(ctx context.Context, request DeleteConfigRequest) (response DeleteConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConfigResponse")
	}
	return
}

// deleteConfig implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/configs/{configId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/Config/DeleteConfig"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDiscoveryJob Deletes a DiscoveryJob by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteDiscoveryJob.go.html to see an example of how to use DeleteDiscoveryJob API.
// A default retry strategy applies to this operation DeleteDiscoveryJob()
func (client StackMonitoringClient) DeleteDiscoveryJob(ctx context.Context, request DeleteDiscoveryJobRequest) (response DeleteDiscoveryJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDiscoveryJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDiscoveryJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDiscoveryJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDiscoveryJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDiscoveryJobResponse")
	}
	return
}

// deleteDiscoveryJob implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteDiscoveryJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/discoveryJobs/{discoveryJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDiscoveryJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/DiscoveryJob/DeleteDiscoveryJob"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteDiscoveryJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMaintenanceWindow Deletes a maintenance window by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteMaintenanceWindow.go.html to see an example of how to use DeleteMaintenanceWindow API.
func (client StackMonitoringClient) DeleteMaintenanceWindow(ctx context.Context, request DeleteMaintenanceWindowRequest) (response DeleteMaintenanceWindowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMaintenanceWindowResponse")
	}
	return
}

// deleteMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/maintenanceWindows/{maintenanceWindowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MaintenanceWindow/DeleteMaintenanceWindow"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMetricExtension Deletes a metric extension by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteMetricExtension.go.html to see an example of how to use DeleteMetricExtension API.
func (client StackMonitoringClient) DeleteMetricExtension(ctx context.Context, request DeleteMetricExtensionRequest) (response DeleteMetricExtensionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMetricExtensionResponse")
	}
	return
}

// deleteMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/metricExtensions/{metricExtensionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/DeleteMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMonitoredResource Delete monitored resource by the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// By default, only the specified resource is deleted. If the parameter 'isDeleteMembers' is set to true,
// then the member resources will be deleted too. If the operation fails partially, the deleted entries
// will not be rolled back.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteMonitoredResource.go.html to see an example of how to use DeleteMonitoredResource API.
func (client StackMonitoringClient) DeleteMonitoredResource(ctx context.Context, request DeleteMonitoredResourceRequest) (response DeleteMonitoredResourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMonitoredResource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMonitoredResourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMonitoredResourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMonitoredResourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMonitoredResourceResponse")
	}
	return
}

// deleteMonitoredResource implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteMonitoredResource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/monitoredResources/{monitoredResourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMonitoredResourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/DeleteMonitoredResource"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteMonitoredResource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMonitoredResourceType Deletes a monitored resource type by identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteMonitoredResourceType.go.html to see an example of how to use DeleteMonitoredResourceType API.
func (client StackMonitoringClient) DeleteMonitoredResourceType(ctx context.Context, request DeleteMonitoredResourceTypeRequest) (response DeleteMonitoredResourceTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMonitoredResourceType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMonitoredResourceTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMonitoredResourceTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMonitoredResourceTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMonitoredResourceTypeResponse")
	}
	return
}

// deleteMonitoredResourceType implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteMonitoredResourceType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/monitoredResourceTypes/{monitoredResourceTypeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMonitoredResourceTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceType/DeleteMonitoredResourceType"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteMonitoredResourceType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProcessSet Deletes a Process Set
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DeleteProcessSet.go.html to see an example of how to use DeleteProcessSet API.
// A default retry strategy applies to this operation DeleteProcessSet()
func (client StackMonitoringClient) DeleteProcessSet(ctx context.Context, request DeleteProcessSetRequest) (response DeleteProcessSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProcessSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProcessSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProcessSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProcessSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProcessSetResponse")
	}
	return
}

// deleteProcessSet implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) deleteProcessSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/processSets/{processSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProcessSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/ProcessSet/DeleteProcessSet"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DeleteProcessSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableExternalDatabase Disable external database resource monitoring. All the references in DBaaS,
// DBM and resource service will be deleted as part of this operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DisableExternalDatabase.go.html to see an example of how to use DisableExternalDatabase API.
func (client StackMonitoringClient) DisableExternalDatabase(ctx context.Context, request DisableExternalDatabaseRequest) (response DisableExternalDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableExternalDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableExternalDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableExternalDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableExternalDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableExternalDatabaseResponse")
	}
	return
}

// disableExternalDatabase implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) disableExternalDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/{monitoredResourceId}/actions/disableExternalDatabase", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableExternalDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/DisableExternalDatabase"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DisableExternalDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableMetricExtension Submits a request to disable matching metric extension Id for the given Resource IDs
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DisableMetricExtension.go.html to see an example of how to use DisableMetricExtension API.
func (client StackMonitoringClient) DisableMetricExtension(ctx context.Context, request DisableMetricExtensionRequest) (response DisableMetricExtensionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disableMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableMetricExtensionResponse")
	}
	return
}

// disableMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) disableMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metricExtensions/{metricExtensionId}/actions/disable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/DisableMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DisableMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisassociateMonitoredResources Removes associations between two monitored resources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/DisassociateMonitoredResources.go.html to see an example of how to use DisassociateMonitoredResources API.
func (client StackMonitoringClient) DisassociateMonitoredResources(ctx context.Context, request DisassociateMonitoredResourcesRequest) (response DisassociateMonitoredResourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.disassociateMonitoredResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisassociateMonitoredResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisassociateMonitoredResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisassociateMonitoredResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisassociateMonitoredResourcesResponse")
	}
	return
}

// disassociateMonitoredResources implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) disassociateMonitoredResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/actions/disassociateMonitoredResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisassociateMonitoredResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/DisassociateMonitoredResources"
		err = common.PostProcessServiceError(err, "StackMonitoring", "DisassociateMonitoredResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableMetricExtension Submits a request to enable matching metric extension Id for the given Resource IDs
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/EnableMetricExtension.go.html to see an example of how to use EnableMetricExtension API.
func (client StackMonitoringClient) EnableMetricExtension(ctx context.Context, request EnableMetricExtensionRequest) (response EnableMetricExtensionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableMetricExtensionResponse")
	}
	return
}

// enableMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) enableMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metricExtensions/{metricExtensionId}/actions/enable", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/EnableMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "EnableMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EvaluateBaselineableMetric Evaluates metric for anomalies for the given data points
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/EvaluateBaselineableMetric.go.html to see an example of how to use EvaluateBaselineableMetric API.
// A default retry strategy applies to this operation EvaluateBaselineableMetric()
func (client StackMonitoringClient) EvaluateBaselineableMetric(ctx context.Context, request EvaluateBaselineableMetricRequest) (response EvaluateBaselineableMetricResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.evaluateBaselineableMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EvaluateBaselineableMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EvaluateBaselineableMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EvaluateBaselineableMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EvaluateBaselineableMetricResponse")
	}
	return
}

// evaluateBaselineableMetric implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) evaluateBaselineableMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/baselineableMetrics/{baselineableMetricId}/actions/evaluate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EvaluateBaselineableMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/BaselineableMetric/EvaluateBaselineableMetric"
		err = common.PostProcessServiceError(err, "StackMonitoring", "EvaluateBaselineableMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportMetricExtension Export generates a template used to create new metric extension resources similar to matching metric extension id.
// Response is a file that contains metric extension definition with placeholders for fields to be changed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ExportMetricExtension.go.html to see an example of how to use ExportMetricExtension API.
// A default retry strategy applies to this operation ExportMetricExtension()
func (client StackMonitoringClient) ExportMetricExtension(ctx context.Context, request ExportMetricExtensionRequest) (response ExportMetricExtensionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportMetricExtensionResponse")
	}
	return
}

// exportMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) exportMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metricExtensions/{metricExtensionId}/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/ExportMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ExportMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBaselineableMetric Get the Baseline-able metric for the given id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetBaselineableMetric.go.html to see an example of how to use GetBaselineableMetric API.
// A default retry strategy applies to this operation GetBaselineableMetric()
func (client StackMonitoringClient) GetBaselineableMetric(ctx context.Context, request GetBaselineableMetricRequest) (response GetBaselineableMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBaselineableMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBaselineableMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBaselineableMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBaselineableMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBaselineableMetricResponse")
	}
	return
}

// getBaselineableMetric implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getBaselineableMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/baselineableMetrics/{baselineableMetricId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBaselineableMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/BaselineableMetric/GetBaselineableMetric"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetBaselineableMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfig Gets the details of a configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetConfig.go.html to see an example of how to use GetConfig API.
// A default retry strategy applies to this operation GetConfig()
func (client StackMonitoringClient) GetConfig(ctx context.Context, request GetConfigRequest) (response GetConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigResponse")
	}
	return
}

// getConfig implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configs/{configId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/Config/GetConfig"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &config{})
	return response, err
}

// GetDiscoveryJob API to get the details of discovery Job by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetDiscoveryJob.go.html to see an example of how to use GetDiscoveryJob API.
// A default retry strategy applies to this operation GetDiscoveryJob()
func (client StackMonitoringClient) GetDiscoveryJob(ctx context.Context, request GetDiscoveryJobRequest) (response GetDiscoveryJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDiscoveryJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDiscoveryJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDiscoveryJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDiscoveryJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDiscoveryJobResponse")
	}
	return
}

// getDiscoveryJob implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getDiscoveryJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryJobs/{discoveryJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDiscoveryJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/DiscoveryJob/GetDiscoveryJob"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetDiscoveryJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaintenanceWindow Get maintenance window for the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetMaintenanceWindow.go.html to see an example of how to use GetMaintenanceWindow API.
// A default retry strategy applies to this operation GetMaintenanceWindow()
func (client StackMonitoringClient) GetMaintenanceWindow(ctx context.Context, request GetMaintenanceWindowRequest) (response GetMaintenanceWindowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaintenanceWindowResponse")
	}
	return
}

// getMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceWindows/{maintenanceWindowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MaintenanceWindow/GetMaintenanceWindow"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMetricExtension Gets a Metric Extension by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetMetricExtension.go.html to see an example of how to use GetMetricExtension API.
// A default retry strategy applies to this operation GetMetricExtension()
func (client StackMonitoringClient) GetMetricExtension(ctx context.Context, request GetMetricExtensionRequest) (response GetMetricExtensionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMetricExtensionResponse")
	}
	return
}

// getMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/metricExtensions/{metricExtensionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/GetMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMonitoredResource Get monitored resource for the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetMonitoredResource.go.html to see an example of how to use GetMonitoredResource API.
// A default retry strategy applies to this operation GetMonitoredResource()
func (client StackMonitoringClient) GetMonitoredResource(ctx context.Context, request GetMonitoredResourceRequest) (response GetMonitoredResourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitoredResource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitoredResourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitoredResourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitoredResourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitoredResourceResponse")
	}
	return
}

// getMonitoredResource implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getMonitoredResource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredResources/{monitoredResourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMonitoredResourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/GetMonitoredResource"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetMonitoredResource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMonitoredResourceTask Gets stack monitoring resource task details by identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetMonitoredResourceTask.go.html to see an example of how to use GetMonitoredResourceTask API.
// A default retry strategy applies to this operation GetMonitoredResourceTask()
func (client StackMonitoringClient) GetMonitoredResourceTask(ctx context.Context, request GetMonitoredResourceTaskRequest) (response GetMonitoredResourceTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitoredResourceTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitoredResourceTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitoredResourceTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitoredResourceTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitoredResourceTaskResponse")
	}
	return
}

// getMonitoredResourceTask implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getMonitoredResourceTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredResourceTasks/{monitoredResourceTaskId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMonitoredResourceTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceTask/GetMonitoredResourceTask"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetMonitoredResourceTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMonitoredResourceType Gets a monitored resource type by identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetMonitoredResourceType.go.html to see an example of how to use GetMonitoredResourceType API.
// A default retry strategy applies to this operation GetMonitoredResourceType()
func (client StackMonitoringClient) GetMonitoredResourceType(ctx context.Context, request GetMonitoredResourceTypeRequest) (response GetMonitoredResourceTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMonitoredResourceType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMonitoredResourceTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMonitoredResourceTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMonitoredResourceTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMonitoredResourceTypeResponse")
	}
	return
}

// getMonitoredResourceType implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getMonitoredResourceType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredResourceTypes/{monitoredResourceTypeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMonitoredResourceTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceType/GetMonitoredResourceType"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetMonitoredResourceType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProcessSet API to get the details of a Process Set by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetProcessSet.go.html to see an example of how to use GetProcessSet API.
// A default retry strategy applies to this operation GetProcessSet()
func (client StackMonitoringClient) GetProcessSet(ctx context.Context, request GetProcessSetRequest) (response GetProcessSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProcessSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProcessSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProcessSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProcessSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProcessSetResponse")
	}
	return
}

// getProcessSet implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) getProcessSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/processSets/{processSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProcessSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/ProcessSet/GetProcessSet"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetProcessSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client StackMonitoringClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client StackMonitoringClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "StackMonitoring", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBaselineableMetrics List of summary of baseline-able metrics for a given resource group if specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListBaselineableMetrics.go.html to see an example of how to use ListBaselineableMetrics API.
// A default retry strategy applies to this operation ListBaselineableMetrics()
func (client StackMonitoringClient) ListBaselineableMetrics(ctx context.Context, request ListBaselineableMetricsRequest) (response ListBaselineableMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBaselineableMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBaselineableMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBaselineableMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBaselineableMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBaselineableMetricsResponse")
	}
	return
}

// listBaselineableMetrics implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listBaselineableMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/baselineableMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBaselineableMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/BaselineableMetric/ListBaselineableMetrics"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListBaselineableMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConfigs Get a list of configurations in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListConfigs.go.html to see an example of how to use ListConfigs API.
// A default retry strategy applies to this operation ListConfigs()
func (client StackMonitoringClient) ListConfigs(ctx context.Context, request ListConfigsRequest) (response ListConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConfigsResponse")
	}
	return
}

// listConfigs implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/ConfigCollection/ListConfigs"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDiscoveryJobLogs API to get all the logs of a Discovery Job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListDiscoveryJobLogs.go.html to see an example of how to use ListDiscoveryJobLogs API.
// A default retry strategy applies to this operation ListDiscoveryJobLogs()
func (client StackMonitoringClient) ListDiscoveryJobLogs(ctx context.Context, request ListDiscoveryJobLogsRequest) (response ListDiscoveryJobLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDiscoveryJobLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDiscoveryJobLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDiscoveryJobLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDiscoveryJobLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDiscoveryJobLogsResponse")
	}
	return
}

// listDiscoveryJobLogs implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listDiscoveryJobLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryJobs/{discoveryJobId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDiscoveryJobLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/DiscoveryJobLogCollection/ListDiscoveryJobLogs"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListDiscoveryJobLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDiscoveryJobs API to get the details of all Discovery Jobs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListDiscoveryJobs.go.html to see an example of how to use ListDiscoveryJobs API.
// A default retry strategy applies to this operation ListDiscoveryJobs()
func (client StackMonitoringClient) ListDiscoveryJobs(ctx context.Context, request ListDiscoveryJobsRequest) (response ListDiscoveryJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDiscoveryJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDiscoveryJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDiscoveryJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDiscoveryJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDiscoveryJobsResponse")
	}
	return
}

// listDiscoveryJobs implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listDiscoveryJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/discoveryJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDiscoveryJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/DiscoveryJobCollection/ListDiscoveryJobs"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListDiscoveryJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaintenanceWindows Returns a list of maintenance windows.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMaintenanceWindows.go.html to see an example of how to use ListMaintenanceWindows API.
// A default retry strategy applies to this operation ListMaintenanceWindows()
func (client StackMonitoringClient) ListMaintenanceWindows(ctx context.Context, request ListMaintenanceWindowsRequest) (response ListMaintenanceWindowsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaintenanceWindows, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaintenanceWindowsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaintenanceWindowsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaintenanceWindowsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaintenanceWindowsResponse")
	}
	return
}

// listMaintenanceWindows implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listMaintenanceWindows(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceWindows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceWindowsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MaintenanceWindow/ListMaintenanceWindows"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListMaintenanceWindows", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMetricExtensions Returns a list of metric extensions
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMetricExtensions.go.html to see an example of how to use ListMetricExtensions API.
// A default retry strategy applies to this operation ListMetricExtensions()
func (client StackMonitoringClient) ListMetricExtensions(ctx context.Context, request ListMetricExtensionsRequest) (response ListMetricExtensionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMetricExtensions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMetricExtensionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMetricExtensionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMetricExtensionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMetricExtensionsResponse")
	}
	return
}

// listMetricExtensions implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listMetricExtensions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/metricExtensions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMetricExtensionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/ListMetricExtensions"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListMetricExtensions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMonitoredResourceTasks Returns a list of stack monitoring resource tasks in the compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMonitoredResourceTasks.go.html to see an example of how to use ListMonitoredResourceTasks API.
// A default retry strategy applies to this operation ListMonitoredResourceTasks()
func (client StackMonitoringClient) ListMonitoredResourceTasks(ctx context.Context, request ListMonitoredResourceTasksRequest) (response ListMonitoredResourceTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMonitoredResourceTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMonitoredResourceTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMonitoredResourceTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMonitoredResourceTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMonitoredResourceTasksResponse")
	}
	return
}

// listMonitoredResourceTasks implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listMonitoredResourceTasks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredResourceTasks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMonitoredResourceTasksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceTask/ListMonitoredResourceTasks"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListMonitoredResourceTasks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMonitoredResourceTypes Returns list of resource types accessible to the customer.
// There are two types of resource types - System resource types and User resource types.
// System resource types are available out of the box in the stack monitoring resource service
// and are accessible to all the tenant users. User resource types are created in the context
// of a tenancy and are visible only for the tenancy. By default, both System resource types
// and User resource types are returned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMonitoredResourceTypes.go.html to see an example of how to use ListMonitoredResourceTypes API.
// A default retry strategy applies to this operation ListMonitoredResourceTypes()
func (client StackMonitoringClient) ListMonitoredResourceTypes(ctx context.Context, request ListMonitoredResourceTypesRequest) (response ListMonitoredResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMonitoredResourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMonitoredResourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMonitoredResourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMonitoredResourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMonitoredResourceTypesResponse")
	}
	return
}

// listMonitoredResourceTypes implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listMonitoredResourceTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredResourceTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMonitoredResourceTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceType/ListMonitoredResourceTypes"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListMonitoredResourceTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMonitoredResources Returns a list of monitored resources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListMonitoredResources.go.html to see an example of how to use ListMonitoredResources API.
// A default retry strategy applies to this operation ListMonitoredResources()
func (client StackMonitoringClient) ListMonitoredResources(ctx context.Context, request ListMonitoredResourcesRequest) (response ListMonitoredResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMonitoredResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMonitoredResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMonitoredResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMonitoredResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMonitoredResourcesResponse")
	}
	return
}

// listMonitoredResources implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listMonitoredResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/monitoredResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMonitoredResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/ListMonitoredResources"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListMonitoredResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProcessSets API to get the details of all Process Sets.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListProcessSets.go.html to see an example of how to use ListProcessSets API.
// A default retry strategy applies to this operation ListProcessSets()
func (client StackMonitoringClient) ListProcessSets(ctx context.Context, request ListProcessSetsRequest) (response ListProcessSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProcessSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProcessSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProcessSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProcessSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProcessSetsResponse")
	}
	return
}

// listProcessSets implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) listProcessSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/processSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProcessSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/ProcessSetCollection/ListProcessSets"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListProcessSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client StackMonitoringClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client StackMonitoringClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/WorkRequestErrorCollection/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client StackMonitoringClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client StackMonitoringClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/WorkRequestLogEntryCollection/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client StackMonitoringClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client StackMonitoringClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/WorkRequestSummaryCollection/ListWorkRequests"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ManageLicense Each resource is assigned a license based on which features are enabled for it.
// User is charged differently based on license.
// Specify the license type to be updated for the parent resource in the topology.
// The license type value is propagated to the member resources as well.
// Member resource is a resource which has "contains" association with the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ManageLicense.go.html to see an example of how to use ManageLicense API.
// A default retry strategy applies to this operation ManageLicense()
func (client StackMonitoringClient) ManageLicense(ctx context.Context, request ManageLicenseRequest) (response ManageLicenseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.manageLicense, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ManageLicenseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ManageLicenseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ManageLicenseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ManageLicenseResponse")
	}
	return
}

// manageLicense implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) manageLicense(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/{monitoredResourceId}/actions/manageLicense", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ManageLicenseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/ManageLicense"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ManageLicense", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishMetricExtension Publish the Metric Extension identified by the id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/PublishMetricExtension.go.html to see an example of how to use PublishMetricExtension API.
func (client StackMonitoringClient) PublishMetricExtension(ctx context.Context, request PublishMetricExtensionRequest) (response PublishMetricExtensionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.publishMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishMetricExtensionResponse")
	}
	return
}

// publishMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) publishMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metricExtensions/{metricExtensionId}/actions/publish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/PublishMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "PublishMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestMonitoredResourcesSummarizedCount Gets resource count based on the aggregation criteria specified using "groupBy" parameter.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/RequestMonitoredResourcesSummarizedCount.go.html to see an example of how to use RequestMonitoredResourcesSummarizedCount API.
// A default retry strategy applies to this operation RequestMonitoredResourcesSummarizedCount()
func (client StackMonitoringClient) RequestMonitoredResourcesSummarizedCount(ctx context.Context, request RequestMonitoredResourcesSummarizedCountRequest) (response RequestMonitoredResourcesSummarizedCountResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestMonitoredResourcesSummarizedCount, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestMonitoredResourcesSummarizedCountResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestMonitoredResourcesSummarizedCountResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestMonitoredResourcesSummarizedCountResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestMonitoredResourcesSummarizedCountResponse")
	}
	return
}

// requestMonitoredResourcesSummarizedCount implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) requestMonitoredResourcesSummarizedCount(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/actions/summarizeCount", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestMonitoredResourcesSummarizedCountResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/RequestMonitoredResourcesSummarizedCount"
		err = common.PostProcessServiceError(err, "StackMonitoring", "RequestMonitoredResourcesSummarizedCount", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetryFailedMaintenanceWindowOperation Retry the last failed operation. The operation failed will be the most recent one. It won't apply for previous failed operations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/RetryFailedMaintenanceWindowOperation.go.html to see an example of how to use RetryFailedMaintenanceWindowOperation API.
func (client StackMonitoringClient) RetryFailedMaintenanceWindowOperation(ctx context.Context, request RetryFailedMaintenanceWindowOperationRequest) (response RetryFailedMaintenanceWindowOperationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retryFailedMaintenanceWindowOperation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetryFailedMaintenanceWindowOperationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetryFailedMaintenanceWindowOperationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetryFailedMaintenanceWindowOperationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetryFailedMaintenanceWindowOperationResponse")
	}
	return
}

// retryFailedMaintenanceWindowOperation implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) retryFailedMaintenanceWindowOperation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maintenanceWindows/{maintenanceWindowId}/actions/retryFailedOperation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetryFailedMaintenanceWindowOperationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MaintenanceWindow/RetryFailedMaintenanceWindowOperation"
		err = common.PostProcessServiceError(err, "StackMonitoring", "RetryFailedMaintenanceWindowOperation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchAssociatedResources List all associated resources recursively up-to a specified level,
// for the monitored resources of type specified.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/SearchAssociatedResources.go.html to see an example of how to use SearchAssociatedResources API.
// A default retry strategy applies to this operation SearchAssociatedResources()
func (client StackMonitoringClient) SearchAssociatedResources(ctx context.Context, request SearchAssociatedResourcesRequest) (response SearchAssociatedResourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchAssociatedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchAssociatedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchAssociatedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchAssociatedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchAssociatedResourcesResponse")
	}
	return
}

// searchAssociatedResources implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) searchAssociatedResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/actions/searchAssociatedResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchAssociatedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/SearchAssociatedResources"
		err = common.PostProcessServiceError(err, "StackMonitoring", "SearchAssociatedResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchMonitoredResourceAssociations Search associations in the given compartment based on the search criteria.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/SearchMonitoredResourceAssociations.go.html to see an example of how to use SearchMonitoredResourceAssociations API.
// A default retry strategy applies to this operation SearchMonitoredResourceAssociations()
func (client StackMonitoringClient) SearchMonitoredResourceAssociations(ctx context.Context, request SearchMonitoredResourceAssociationsRequest) (response SearchMonitoredResourceAssociationsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchMonitoredResourceAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchMonitoredResourceAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchMonitoredResourceAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchMonitoredResourceAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchMonitoredResourceAssociationsResponse")
	}
	return
}

// searchMonitoredResourceAssociations implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) searchMonitoredResourceAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/actions/searchAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchMonitoredResourceAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/SearchMonitoredResourceAssociations"
		err = common.PostProcessServiceError(err, "StackMonitoring", "SearchMonitoredResourceAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchMonitoredResourceMembers List the member resources for the given monitored resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/SearchMonitoredResourceMembers.go.html to see an example of how to use SearchMonitoredResourceMembers API.
// A default retry strategy applies to this operation SearchMonitoredResourceMembers()
func (client StackMonitoringClient) SearchMonitoredResourceMembers(ctx context.Context, request SearchMonitoredResourceMembersRequest) (response SearchMonitoredResourceMembersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchMonitoredResourceMembers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchMonitoredResourceMembersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchMonitoredResourceMembersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchMonitoredResourceMembersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchMonitoredResourceMembersResponse")
	}
	return
}

// searchMonitoredResourceMembers implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) searchMonitoredResourceMembers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/{monitoredResourceId}/actions/listMembers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchMonitoredResourceMembersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/SearchMonitoredResourceMembers"
		err = common.PostProcessServiceError(err, "StackMonitoring", "SearchMonitoredResourceMembers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchMonitoredResources Gets a list of all monitored resources in a compartment for the given search criteria.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/SearchMonitoredResources.go.html to see an example of how to use SearchMonitoredResources API.
// A default retry strategy applies to this operation SearchMonitoredResources()
func (client StackMonitoringClient) SearchMonitoredResources(ctx context.Context, request SearchMonitoredResourcesRequest) (response SearchMonitoredResourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.searchMonitoredResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchMonitoredResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchMonitoredResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchMonitoredResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchMonitoredResourcesResponse")
	}
	return
}

// searchMonitoredResources implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) searchMonitoredResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/actions/search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchMonitoredResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/SearchMonitoredResources"
		err = common.PostProcessServiceError(err, "StackMonitoring", "SearchMonitoredResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopMaintenanceWindow Stop a maintenance window before the end time is reached.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/StopMaintenanceWindow.go.html to see an example of how to use StopMaintenanceWindow API.
func (client StackMonitoringClient) StopMaintenanceWindow(ctx context.Context, request StopMaintenanceWindowRequest) (response StopMaintenanceWindowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.stopMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopMaintenanceWindowResponse")
	}
	return
}

// stopMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) stopMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maintenanceWindows/{maintenanceWindowId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MaintenanceWindow/StopMaintenanceWindow"
		err = common.PostProcessServiceError(err, "StackMonitoring", "StopMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestMetricExtension Performs test of Metric Extension on a specific resource Id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/TestMetricExtension.go.html to see an example of how to use TestMetricExtension API.
func (client StackMonitoringClient) TestMetricExtension(ctx context.Context, request TestMetricExtensionRequest) (response TestMetricExtensionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.testMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestMetricExtensionResponse")
	}
	return
}

// testMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) testMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metricExtensions/{metricExtensionId}/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/TestMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "TestMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAndPropagateTags Provided tags will be added or updated in the existing list of tags for the affected resources.
// Resources to be updated are identified based on association types specified.
// If association types not specified, then tags will be updated only for the resource identified by
// the given monitored resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateAndPropagateTags.go.html to see an example of how to use UpdateAndPropagateTags API.
// A default retry strategy applies to this operation UpdateAndPropagateTags()
func (client StackMonitoringClient) UpdateAndPropagateTags(ctx context.Context, request UpdateAndPropagateTagsRequest) (response UpdateAndPropagateTagsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateAndPropagateTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAndPropagateTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAndPropagateTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAndPropagateTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAndPropagateTagsResponse")
	}
	return
}

// updateAndPropagateTags implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateAndPropagateTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/monitoredResources/{monitoredResourceId}/actions/updateAndPropagateTags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAndPropagateTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/UpdateAndPropagateTags"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateAndPropagateTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBaselineableMetric Updates the Baseline-able metric for the given id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateBaselineableMetric.go.html to see an example of how to use UpdateBaselineableMetric API.
// A default retry strategy applies to this operation UpdateBaselineableMetric()
func (client StackMonitoringClient) UpdateBaselineableMetric(ctx context.Context, request UpdateBaselineableMetricRequest) (response UpdateBaselineableMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBaselineableMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBaselineableMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBaselineableMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBaselineableMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBaselineableMetricResponse")
	}
	return
}

// updateBaselineableMetric implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateBaselineableMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/baselineableMetrics/{baselineableMetricId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBaselineableMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/BaselineableMetric/UpdateBaselineableMetric"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateBaselineableMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConfig Updates the configuration identified by the id given.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateConfig.go.html to see an example of how to use UpdateConfig API.
// A default retry strategy applies to this operation UpdateConfig()
func (client StackMonitoringClient) UpdateConfig(ctx context.Context, request UpdateConfigRequest) (response UpdateConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigResponse")
	}
	return
}

// updateConfig implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configs/{configId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/Config/UpdateConfig"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &config{})
	return response, err
}

// UpdateMaintenanceWindow Update maintenance window by the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateMaintenanceWindow.go.html to see an example of how to use UpdateMaintenanceWindow API.
func (client StackMonitoringClient) UpdateMaintenanceWindow(ctx context.Context, request UpdateMaintenanceWindowRequest) (response UpdateMaintenanceWindowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMaintenanceWindowResponse")
	}
	return
}

// updateMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maintenanceWindows/{maintenanceWindowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MaintenanceWindow/UpdateMaintenanceWindow"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMetricExtension Updates the Metric Extension
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateMetricExtension.go.html to see an example of how to use UpdateMetricExtension API.
func (client StackMonitoringClient) UpdateMetricExtension(ctx context.Context, request UpdateMetricExtensionRequest) (response UpdateMetricExtensionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMetricExtension, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMetricExtensionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMetricExtensionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMetricExtensionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMetricExtensionResponse")
	}
	return
}

// updateMetricExtension implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateMetricExtension(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/metricExtensions/{metricExtensionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMetricExtensionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MetricExtension/UpdateMetricExtension"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateMetricExtension", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMonitoredResource Update monitored resource by the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// Note that "properties" object, if specified, will entirely replace the existing object,
// as part this operation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateMonitoredResource.go.html to see an example of how to use UpdateMonitoredResource API.
func (client StackMonitoringClient) UpdateMonitoredResource(ctx context.Context, request UpdateMonitoredResourceRequest) (response UpdateMonitoredResourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMonitoredResource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMonitoredResourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMonitoredResourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMonitoredResourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMonitoredResourceResponse")
	}
	return
}

// updateMonitoredResource implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateMonitoredResource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/monitoredResources/{monitoredResourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMonitoredResourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResource/UpdateMonitoredResource"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateMonitoredResource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMonitoredResourceTask Update stack monitoring resource task by the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateMonitoredResourceTask.go.html to see an example of how to use UpdateMonitoredResourceTask API.
func (client StackMonitoringClient) UpdateMonitoredResourceTask(ctx context.Context, request UpdateMonitoredResourceTaskRequest) (response UpdateMonitoredResourceTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMonitoredResourceTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMonitoredResourceTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMonitoredResourceTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMonitoredResourceTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMonitoredResourceTaskResponse")
	}
	return
}

// updateMonitoredResourceTask implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateMonitoredResourceTask(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/monitoredResourceTasks/{monitoredResourceTaskId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMonitoredResourceTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceTask/UpdateMonitoredResourceTask"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateMonitoredResourceTask", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMonitoredResourceType Update the Monitored Resource Type identified by the identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateMonitoredResourceType.go.html to see an example of how to use UpdateMonitoredResourceType API.
func (client StackMonitoringClient) UpdateMonitoredResourceType(ctx context.Context, request UpdateMonitoredResourceTypeRequest) (response UpdateMonitoredResourceTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMonitoredResourceType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMonitoredResourceTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMonitoredResourceTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMonitoredResourceTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMonitoredResourceTypeResponse")
	}
	return
}

// updateMonitoredResourceType implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateMonitoredResourceType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/monitoredResourceTypes/{monitoredResourceTypeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMonitoredResourceTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/MonitoredResourceType/UpdateMonitoredResourceType"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateMonitoredResourceType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProcessSet API to update a Process Set identified by a given ocid.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/UpdateProcessSet.go.html to see an example of how to use UpdateProcessSet API.
// A default retry strategy applies to this operation UpdateProcessSet()
func (client StackMonitoringClient) UpdateProcessSet(ctx context.Context, request UpdateProcessSetRequest) (response UpdateProcessSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProcessSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProcessSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProcessSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProcessSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProcessSetResponse")
	}
	return
}

// updateProcessSet implements the OCIOperation interface (enables retrying operations)
func (client StackMonitoringClient) updateProcessSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/processSets/{processSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProcessSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/ProcessSet/UpdateProcessSet"
		err = common.PostProcessServiceError(err, "StackMonitoring", "UpdateProcessSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
