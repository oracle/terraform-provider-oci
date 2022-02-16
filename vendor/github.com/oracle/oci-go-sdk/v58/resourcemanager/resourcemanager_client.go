// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//ResourceManagerClient a client for ResourceManager
type ResourceManagerClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewResourceManagerClientWithConfigurationProvider Creates a new default ResourceManager client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewResourceManagerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ResourceManagerClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newResourceManagerClientFromBaseClient(baseClient, provider)
}

// NewResourceManagerClientWithOboToken Creates a new default ResourceManager client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewResourceManagerClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ResourceManagerClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newResourceManagerClientFromBaseClient(baseClient, configProvider)
}

func newResourceManagerClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ResourceManagerClient, err error) {
	// ResourceManager service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ResourceManagerClient{BaseClient: baseClient}
	client.BasePath = "20180917"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ResourceManagerClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("resourcemanager", "https://resourcemanager.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ResourceManagerClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ResourceManagerClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelJob Indicates the intention to cancel the specified job.
// Cancellation of the job is not immediate, and may be delayed,
// or may not happen at all.
// You can optionally choose forced cancellation by setting `isForced` to true.
// A forced cancellation can result in an incorrect state file.
// For example, the state file might not reflect the exact state of the provisioned resources.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/CancelJob.go.html to see an example of how to use CancelJob API.
func (client ResourceManagerClient) CancelJob(ctx context.Context, request CancelJobRequest) (response CancelJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelJobResponse")
	}
	return
}

// cancelJob implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) cancelJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/jobs/{jobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelJobResponse
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

// ChangeConfigurationSourceProviderCompartment Moves a configuration source provider into a different compartment within the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ChangeConfigurationSourceProviderCompartment.go.html to see an example of how to use ChangeConfigurationSourceProviderCompartment API.
func (client ResourceManagerClient) ChangeConfigurationSourceProviderCompartment(ctx context.Context, request ChangeConfigurationSourceProviderCompartmentRequest) (response ChangeConfigurationSourceProviderCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeConfigurationSourceProviderCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeConfigurationSourceProviderCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeConfigurationSourceProviderCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeConfigurationSourceProviderCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeConfigurationSourceProviderCompartmentResponse")
	}
	return
}

// changeConfigurationSourceProviderCompartment implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) changeConfigurationSourceProviderCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configurationSourceProviders/{configurationSourceProviderId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeConfigurationSourceProviderCompartmentResponse
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

// ChangeStackCompartment Moves a Stack and it's associated Jobs into a different compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ChangeStackCompartment.go.html to see an example of how to use ChangeStackCompartment API.
func (client ResourceManagerClient) ChangeStackCompartment(ctx context.Context, request ChangeStackCompartmentRequest) (response ChangeStackCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeStackCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeStackCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeStackCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeStackCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeStackCompartmentResponse")
	}
	return
}

// changeStackCompartment implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) changeStackCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/stacks/{stackId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeStackCompartmentResponse
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

// ChangeTemplateCompartment Moves a template into a different compartment within the same tenancy.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ChangeTemplateCompartment.go.html to see an example of how to use ChangeTemplateCompartment API.
func (client ResourceManagerClient) ChangeTemplateCompartment(ctx context.Context, request ChangeTemplateCompartmentRequest) (response ChangeTemplateCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTemplateCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTemplateCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTemplateCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTemplateCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTemplateCompartmentResponse")
	}
	return
}

// changeTemplateCompartment implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) changeTemplateCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/templates/{templateId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTemplateCompartmentResponse
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

// CreateConfigurationSourceProvider Creates a configuration source provider in the specified compartment.
// For more information, see
// To create a configuration source provider (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/managingconfigurationsourceproviders.htm#CreateConfigurationSourceProvider).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/CreateConfigurationSourceProvider.go.html to see an example of how to use CreateConfigurationSourceProvider API.
func (client ResourceManagerClient) CreateConfigurationSourceProvider(ctx context.Context, request CreateConfigurationSourceProviderRequest) (response CreateConfigurationSourceProviderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConfigurationSourceProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConfigurationSourceProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConfigurationSourceProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConfigurationSourceProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConfigurationSourceProviderResponse")
	}
	return
}

// createConfigurationSourceProvider implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) createConfigurationSourceProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/configurationSourceProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConfigurationSourceProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &configurationsourceprovider{})
	return response, err
}

// CreateJob Creates a job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/CreateJob.go.html to see an example of how to use CreateJob API.
func (client ResourceManagerClient) CreateJob(ctx context.Context, request CreateJobRequest) (response CreateJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobResponse")
	}
	return
}

// createJob implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) createJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/jobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJobResponse
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

// CreateStack Creates a stack in the specified compartment.
// You can create a stack from a Terraform configuration.
// The Terraform configuration can be directly uploaded or referenced from a source code control system.
// You can also create a stack from an existing compartment.
// You can also upload the Terraform configuration from an Object Storage bucket.
// For more information, see
// To create a stack (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/managingstacksandjobs.htm#createstack-all).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/CreateStack.go.html to see an example of how to use CreateStack API.
func (client ResourceManagerClient) CreateStack(ctx context.Context, request CreateStackRequest) (response CreateStackResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createStack, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateStackResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateStackResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStackResponse")
	}
	return
}

// createStack implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) createStack(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/stacks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateStackResponse
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

// CreateTemplate Creates a private template in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/CreateTemplate.go.html to see an example of how to use CreateTemplate API.
func (client ResourceManagerClient) CreateTemplate(ctx context.Context, request CreateTemplateRequest) (response CreateTemplateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTemplate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTemplateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTemplateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTemplateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTemplateResponse")
	}
	return
}

// createTemplate implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) createTemplate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/templates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTemplateResponse
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

// DeleteConfigurationSourceProvider Deletes the specified configuration source provider.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/DeleteConfigurationSourceProvider.go.html to see an example of how to use DeleteConfigurationSourceProvider API.
func (client ResourceManagerClient) DeleteConfigurationSourceProvider(ctx context.Context, request DeleteConfigurationSourceProviderRequest) (response DeleteConfigurationSourceProviderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConfigurationSourceProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConfigurationSourceProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConfigurationSourceProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConfigurationSourceProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConfigurationSourceProviderResponse")
	}
	return
}

// deleteConfigurationSourceProvider implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) deleteConfigurationSourceProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/configurationSourceProviders/{configurationSourceProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConfigurationSourceProviderResponse
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

// DeleteStack Deletes the specified stack object.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/DeleteStack.go.html to see an example of how to use DeleteStack API.
func (client ResourceManagerClient) DeleteStack(ctx context.Context, request DeleteStackRequest) (response DeleteStackResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStack, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteStackResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteStackResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteStackResponse")
	}
	return
}

// deleteStack implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) deleteStack(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/stacks/{stackId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteStackResponse
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

// DeleteTemplate Deletes the specified template.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/DeleteTemplate.go.html to see an example of how to use DeleteTemplate API.
func (client ResourceManagerClient) DeleteTemplate(ctx context.Context, request DeleteTemplateRequest) (response DeleteTemplateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTemplate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTemplateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTemplateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTemplateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTemplateResponse")
	}
	return
}

// deleteTemplate implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) deleteTemplate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/templates/{templateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTemplateResponse
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

// DetectStackDrift Checks drift status for the specified stack.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/DetectStackDrift.go.html to see an example of how to use DetectStackDrift API.
func (client ResourceManagerClient) DetectStackDrift(ctx context.Context, request DetectStackDriftRequest) (response DetectStackDriftResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detectStackDrift, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetectStackDriftResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetectStackDriftResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetectStackDriftResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetectStackDriftResponse")
	}
	return
}

// detectStackDrift implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) detectStackDrift(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/stacks/{stackId}/actions/detectDrift", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetectStackDriftResponse
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

// GetConfigurationSourceProvider Gets the properties of the specified configuration source provider.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetConfigurationSourceProvider.go.html to see an example of how to use GetConfigurationSourceProvider API.
func (client ResourceManagerClient) GetConfigurationSourceProvider(ctx context.Context, request GetConfigurationSourceProviderRequest) (response GetConfigurationSourceProviderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfigurationSourceProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigurationSourceProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigurationSourceProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationSourceProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationSourceProviderResponse")
	}
	return
}

// getConfigurationSourceProvider implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getConfigurationSourceProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configurationSourceProviders/{configurationSourceProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigurationSourceProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &configurationsourceprovider{})
	return response, err
}

// GetJob Returns the specified job along with the job details.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJob.go.html to see an example of how to use GetJob API.
func (client ResourceManagerClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobResponse")
	}
	return
}

// getJob implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobResponse
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

// GetJobDetailedLogContent Returns the Terraform detailed log content for the specified job in plain text. Learn about Terraform detailed log. (https://www.terraform.io/docs/internals/debugging.html)
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJobDetailedLogContent.go.html to see an example of how to use GetJobDetailedLogContent API.
func (client ResourceManagerClient) GetJobDetailedLogContent(ctx context.Context, request GetJobDetailedLogContentRequest) (response GetJobDetailedLogContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobDetailedLogContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobDetailedLogContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobDetailedLogContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobDetailedLogContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobDetailedLogContentResponse")
	}
	return
}

// getJobDetailedLogContent implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getJobDetailedLogContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/detailedLogContent", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobDetailedLogContentResponse
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

// GetJobLogs Returns console log entries for the specified job in JSON format.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJobLogs.go.html to see an example of how to use GetJobLogs API.
func (client ResourceManagerClient) GetJobLogs(ctx context.Context, request GetJobLogsRequest) (response GetJobLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobLogsResponse")
	}
	return
}

// getJobLogs implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getJobLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobLogsResponse
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

// GetJobLogsContent Returns a raw log file for the specified job. The raw log file contains console log entries in text format. The maximum number of entries in a file is 100,000.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJobLogsContent.go.html to see an example of how to use GetJobLogsContent API.
func (client ResourceManagerClient) GetJobLogsContent(ctx context.Context, request GetJobLogsContentRequest) (response GetJobLogsContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobLogsContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobLogsContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobLogsContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobLogsContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobLogsContentResponse")
	}
	return
}

// getJobLogsContent implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getJobLogsContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/logs/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobLogsContentResponse
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

// GetJobTfConfig Returns the Terraform configuration file for the specified job in .zip format.
// Returns an error if no zip file is found.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJobTfConfig.go.html to see an example of how to use GetJobTfConfig API.
func (client ResourceManagerClient) GetJobTfConfig(ctx context.Context, request GetJobTfConfigRequest) (response GetJobTfConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobTfConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobTfConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobTfConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobTfConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobTfConfigResponse")
	}
	return
}

// getJobTfConfig implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getJobTfConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/tfConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobTfConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobTfState Returns the Terraform state for the specified job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJobTfState.go.html to see an example of how to use GetJobTfState API.
func (client ResourceManagerClient) GetJobTfState(ctx context.Context, request GetJobTfStateRequest) (response GetJobTfStateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobTfState, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobTfStateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobTfStateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobTfStateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobTfStateResponse")
	}
	return
}

// getJobTfState implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getJobTfState(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs/{jobId}/tfState", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobTfStateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStack Gets a stack using the stack ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetStack.go.html to see an example of how to use GetStack API.
func (client ResourceManagerClient) GetStack(ctx context.Context, request GetStackRequest) (response GetStackResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStack, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStackResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStackResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStackResponse")
	}
	return
}

// getStack implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getStack(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/stacks/{stackId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStackResponse
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

// GetStackTfConfig Returns the Terraform configuration file in .zip format for the specified stack.
// Returns an error if no zip file is found.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetStackTfConfig.go.html to see an example of how to use GetStackTfConfig API.
func (client ResourceManagerClient) GetStackTfConfig(ctx context.Context, request GetStackTfConfigRequest) (response GetStackTfConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStackTfConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStackTfConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStackTfConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStackTfConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStackTfConfigResponse")
	}
	return
}

// getStackTfConfig implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getStackTfConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/stacks/{stackId}/tfConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStackTfConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStackTfState Returns the Terraform state for the specified stack.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetStackTfState.go.html to see an example of how to use GetStackTfState API.
func (client ResourceManagerClient) GetStackTfState(ctx context.Context, request GetStackTfStateRequest) (response GetStackTfStateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStackTfState, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStackTfStateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStackTfStateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStackTfStateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStackTfStateResponse")
	}
	return
}

// getStackTfState implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getStackTfState(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/stacks/{stackId}/tfState", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStackTfStateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTemplate Gets the specified template.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetTemplate.go.html to see an example of how to use GetTemplate API.
func (client ResourceManagerClient) GetTemplate(ctx context.Context, request GetTemplateRequest) (response GetTemplateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTemplate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTemplateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTemplateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTemplateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTemplateResponse")
	}
	return
}

// getTemplate implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getTemplate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/templates/{templateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTemplateResponse
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

// GetTemplateLogo Returns the Terraform logo file in .logo format for the specified template.
// Returns an error if no logo file is found.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetTemplateLogo.go.html to see an example of how to use GetTemplateLogo API.
func (client ResourceManagerClient) GetTemplateLogo(ctx context.Context, request GetTemplateLogoRequest) (response GetTemplateLogoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTemplateLogo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTemplateLogoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTemplateLogoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTemplateLogoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTemplateLogoResponse")
	}
	return
}

// getTemplateLogo implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getTemplateLogo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/templates/{templateId}/logo", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTemplateLogoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTemplateTfConfig Returns the Terraform configuration file in .zip format for the specified template.
// Returns an error if no zip file is found.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetTemplateTfConfig.go.html to see an example of how to use GetTemplateTfConfig API.
func (client ResourceManagerClient) GetTemplateTfConfig(ctx context.Context, request GetTemplateTfConfigRequest) (response GetTemplateTfConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTemplateTfConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTemplateTfConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTemplateTfConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTemplateTfConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTemplateTfConfigResponse")
	}
	return
}

// getTemplateTfConfig implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) getTemplateTfConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/templates/{templateId}/tfConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTemplateTfConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Return the given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client ResourceManagerClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ResourceManagerClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConfigurationSourceProviders Lists configuration source providers according to the specified filter.
// - For `compartmentId`, lists all configuration source providers in the matching compartment.
// - For `configurationSourceProviderId`, lists the matching configuration source provider.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListConfigurationSourceProviders.go.html to see an example of how to use ListConfigurationSourceProviders API.
func (client ResourceManagerClient) ListConfigurationSourceProviders(ctx context.Context, request ListConfigurationSourceProvidersRequest) (response ListConfigurationSourceProvidersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConfigurationSourceProviders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConfigurationSourceProvidersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConfigurationSourceProvidersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConfigurationSourceProvidersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConfigurationSourceProvidersResponse")
	}
	return
}

// listConfigurationSourceProviders implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listConfigurationSourceProviders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configurationSourceProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConfigurationSourceProvidersResponse
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

// ListJobs Returns a list of jobs in a stack or compartment, ordered by time created.
// - To list all jobs in a stack, provide the stack OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// - To list all jobs in a compartment, provide the compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
// - To return a specific job, provide the job OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListJobs.go.html to see an example of how to use ListJobs API.
func (client ResourceManagerClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobsResponse")
	}
	return
}

// listJobs implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/jobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobsResponse
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

// ListResourceDiscoveryServices Returns a list of supported services for Resource Discovery. For reference on service names, see the Terraform provider documentation (https://www.terraform.io/docs/providers/oci/guides/resource_discovery.html#services).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListResourceDiscoveryServices.go.html to see an example of how to use ListResourceDiscoveryServices API.
func (client ResourceManagerClient) ListResourceDiscoveryServices(ctx context.Context, request ListResourceDiscoveryServicesRequest) (response ListResourceDiscoveryServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceDiscoveryServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceDiscoveryServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceDiscoveryServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceDiscoveryServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceDiscoveryServicesResponse")
	}
	return
}

// listResourceDiscoveryServices implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listResourceDiscoveryServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceDiscoveryServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceDiscoveryServicesResponse
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

// ListStackResourceDriftDetails Lists drift status details for each resource defined in the specified stack.
// The drift status details for a given resource indicate differences, if any, between the actual state
// and the expected (defined) state for that resource.
// The drift status details correspond to the specified work request (`workRequestId`).
// If no work request is specified, then the drift status details correspond to the latest completed work request for the stack.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListStackResourceDriftDetails.go.html to see an example of how to use ListStackResourceDriftDetails API.
func (client ResourceManagerClient) ListStackResourceDriftDetails(ctx context.Context, request ListStackResourceDriftDetailsRequest) (response ListStackResourceDriftDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStackResourceDriftDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStackResourceDriftDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStackResourceDriftDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStackResourceDriftDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStackResourceDriftDetailsResponse")
	}
	return
}

// listStackResourceDriftDetails implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listStackResourceDriftDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/stacks/{stackId}/actions/listResourceDriftDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStackResourceDriftDetailsResponse
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

// ListStacks Returns a list of stacks.
// - If called using the compartment ID, returns all stacks in the specified compartment.
// - If called using the stack ID, returns the specified stack.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListStacks.go.html to see an example of how to use ListStacks API.
func (client ResourceManagerClient) ListStacks(ctx context.Context, request ListStacksRequest) (response ListStacksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStacks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStacksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStacksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStacksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStacksResponse")
	}
	return
}

// listStacks implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listStacks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/stacks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStacksResponse
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

// ListTemplateCategories Lists template categories.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListTemplateCategories.go.html to see an example of how to use ListTemplateCategories API.
func (client ResourceManagerClient) ListTemplateCategories(ctx context.Context, request ListTemplateCategoriesRequest) (response ListTemplateCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTemplateCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTemplateCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTemplateCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTemplateCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTemplateCategoriesResponse")
	}
	return
}

// listTemplateCategories implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listTemplateCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/templateCategories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTemplateCategoriesResponse
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

// ListTemplates Lists templates according to the specified filter.
// The attributes `compartmentId` and `templateCategoryId` are required unless `templateId` is specified.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListTemplates.go.html to see an example of how to use ListTemplates API.
func (client ResourceManagerClient) ListTemplates(ctx context.Context, request ListTemplatesRequest) (response ListTemplatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTemplates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTemplatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTemplatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTemplatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTemplatesResponse")
	}
	return
}

// listTemplates implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listTemplates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/templates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTemplatesResponse
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

// ListTerraformVersions Returns a list of supported Terraform versions for use with stacks.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListTerraformVersions.go.html to see an example of how to use ListTerraformVersions API.
func (client ResourceManagerClient) ListTerraformVersions(ctx context.Context, request ListTerraformVersionsRequest) (response ListTerraformVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTerraformVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTerraformVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTerraformVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTerraformVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTerraformVersionsResponse")
	}
	return
}

// listTerraformVersions implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) listTerraformVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/terraformVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTerraformVersionsResponse
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client ResourceManagerClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ResourceManagerClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client ResourceManagerClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ResourceManagerClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a given compartment or for a given resource.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client ResourceManagerClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ResourceManagerClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConfigurationSourceProvider Updates the properties of the specified configuration source provider.
// For more information, see
// To edit a configuration source provider (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/managingconfigurationsourceproviders.htm#EditConfigurationSourceProvider).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/UpdateConfigurationSourceProvider.go.html to see an example of how to use UpdateConfigurationSourceProvider API.
func (client ResourceManagerClient) UpdateConfigurationSourceProvider(ctx context.Context, request UpdateConfigurationSourceProviderRequest) (response UpdateConfigurationSourceProviderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConfigurationSourceProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigurationSourceProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigurationSourceProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigurationSourceProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigurationSourceProviderResponse")
	}
	return
}

// updateConfigurationSourceProvider implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) updateConfigurationSourceProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configurationSourceProviders/{configurationSourceProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationSourceProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &configurationsourceprovider{})
	return response, err
}

// UpdateJob Updates the specified job.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/UpdateJob.go.html to see an example of how to use UpdateJob API.
func (client ResourceManagerClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobResponse")
	}
	return
}

// updateJob implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) updateJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/jobs/{jobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateJobResponse
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

// UpdateStack Updates the specified stack object.
// Use `UpdateStack` when you update your Terraform configuration
// and want your changes to be reflected in the execution plan.
// For more information, see
// To update a stack (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/managingstacksandjobs.htm#UpdateStack) and
// To edit a stack (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/managingstacksandjobs.htm#EditStack).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/UpdateStack.go.html to see an example of how to use UpdateStack API.
func (client ResourceManagerClient) UpdateStack(ctx context.Context, request UpdateStackRequest) (response UpdateStackResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStack, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStackResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStackResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStackResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStackResponse")
	}
	return
}

// updateStack implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) updateStack(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/stacks/{stackId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateStackResponse
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

// UpdateTemplate Updates the specified template.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/UpdateTemplate.go.html to see an example of how to use UpdateTemplate API.
func (client ResourceManagerClient) UpdateTemplate(ctx context.Context, request UpdateTemplateRequest) (response UpdateTemplateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTemplate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTemplateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTemplateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTemplateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTemplateResponse")
	}
	return
}

// updateTemplate implements the OCIOperation interface (enables retrying operations)
func (client ResourceManagerClient) updateTemplate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/templates/{templateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTemplateResponse
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
