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
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
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

// ChangeMonitoredResourceCompartment Moves a monitored resource from one compartment to another.
// When provided, If-Match is checked against ETag values of the resource.
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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/stack-monitoring/20210330/monitoredResourceTask/ChangeMonitoredResourceTaskCompartment"
		err = common.PostProcessServiceError(err, "StackMonitoring", "ChangeMonitoredResourceTaskCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDiscoveryJob API to create discovery Job and submit discovery Details to agent.
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

// DeleteDiscoveryJob Deletes a DiscoveryJob by identifier
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

// DisableExternalDatabase Disable external database resource monitoring. All the references in DBaaS,
// DBM and resource service will be deleted as part of this operation.
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

// GetDiscoveryJob API to get the details of discovery Job by identifier.
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

// GetWorkRequest Gets the status of the work request with the given ID.
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

// ListDiscoveryJobLogs API to get all the logs of a Discovery Job.
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

// ListMonitoredResourceTasks Returns a list of stack monitoring resource tasks in the compartment.
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
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

// RequestMonitoredResourcesSummarizedCount Gets resource count based on the aggregation criteria specified using "groupBy" parameter.
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

// SearchAssociatedResources List all associated resources recursively up-to a specified level,
// for the monitored resources of type specified.
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

// UpdateMonitoredResource Update monitored resource by the given identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// Note that "properties" object, if specified, will entirely replace the existing object,
// as part this operation.
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
