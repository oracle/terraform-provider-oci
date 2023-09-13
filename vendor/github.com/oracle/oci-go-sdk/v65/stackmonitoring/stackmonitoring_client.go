// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
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
// See also
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
// See also
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

// ChangeMonitoredResourceCompartment Moves a monitored resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// See also
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

// CreateConfig Creates a configuration item, for example to define
// whether resources of a specific type should be discovered automatically.
// For example, when a new Management Agent gets registered in a certain compartment,
// this Management Agent can potentially get promoted to a HOST resource.
// The configuration item will determine if HOST resources in the selected compartment will be
// discovered automatically.
//
// See also
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
// See also
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

// CreateMonitoredResource Creates a new monitored resource for the given resource type with the details and submits
// a work request for promoting the resource to agent. Once the resource is successfully
// added to agent, resource state will be marked active.
//
// See also
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

// DeleteConfig Deletes a configuration identified by the id.
//
// See also
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
// See also
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

// DeleteMonitoredResource Delete monitored resource by the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// By default, only the specified resource is deleted. If the parameter 'isDeleteMembers' is set to true,
// then the member resources will be deleted too. If the operation fails partially, the deleted entries
// will not be rolled back.
//
// See also
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

// DisableExternalDatabase Disable external database resource monitoring. All the references in DBaaS,
// DBM and resource service will be deleted as part of this operation.
//
// See also
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

// DisassociateMonitoredResources Removes associations between two monitored resources.
//
// See also
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

// GetConfig Gets the details of a configuration.
//
// See also
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
// See also
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

// GetMonitoredResource Get monitored resource for the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// See also
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

// GetWorkRequest Gets the status of the work request with the given ID.
//
// See also
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

// ListConfigs Get a list of configurations in a compartment.
//
// See also
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
// See also
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
// See also
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// See also
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
// See also
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
// See also
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

// SearchAssociatedResources List all associated resources recursively up-to a specified level,
// for the monitored resources of type specified.
//
// See also
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
// See also
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
// See also
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
// See also
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

// UpdateAndPropagateTags Provided tags will be added or updated in the existing list of tags for the affected resources.
// Resources to be updated are identified based on association types specified.
// If association types not specified, then tags will be updated only for the resource identified by
// the given monitored resource identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
//
// See also
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

// UpdateConfig Updates the configuration identified by the id given.
//
// See also
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

// UpdateMonitoredResource Update monitored resource by the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// Note that "properties" object, if specified, will entirely replace the existing object,
// as part this operation.
//
// See also
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
