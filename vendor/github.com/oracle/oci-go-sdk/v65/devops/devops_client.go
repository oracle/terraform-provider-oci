// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DevopsClient a client for Devops
type DevopsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDevopsClientWithConfigurationProvider Creates a new default Devops client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDevopsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DevopsClient, err error) {
	if enabled := common.CheckForEnabledServices("devops"); !enabled {
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
	return newDevopsClientFromBaseClient(baseClient, provider)
}

// NewDevopsClientWithOboToken Creates a new default Devops client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDevopsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DevopsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDevopsClientFromBaseClient(baseClient, configProvider)
}

func newDevopsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DevopsClient, err error) {
	// Devops service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Devops"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DevopsClient{BaseClient: baseClient}
	client.BasePath = "20210630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DevopsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("devops", "https://devops.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DevopsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DevopsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ApproveDeployment Submit stage approval.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ApproveDeployment.go.html to see an example of how to use ApproveDeployment API.
// A default retry strategy applies to this operation ApproveDeployment()
func (client DevopsClient) ApproveDeployment(ctx context.Context, request ApproveDeploymentRequest) (response ApproveDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.approveDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApproveDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApproveDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApproveDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApproveDeploymentResponse")
	}
	return
}

// approveDeployment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) approveDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/approve", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApproveDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Deployment/ApproveDeployment"
		err = common.PostProcessServiceError(err, "Devops", "ApproveDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// CancelBuildRun Cancels the build run based on the build run ID provided in the request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CancelBuildRun.go.html to see an example of how to use CancelBuildRun API.
// A default retry strategy applies to this operation CancelBuildRun()
func (client DevopsClient) CancelBuildRun(ctx context.Context, request CancelBuildRunRequest) (response CancelBuildRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelBuildRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelBuildRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelBuildRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelBuildRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelBuildRunResponse")
	}
	return
}

// cancelBuildRun implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) cancelBuildRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/buildRuns/{buildRunId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelBuildRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildRun/CancelBuildRun"
		err = common.PostProcessServiceError(err, "Devops", "CancelBuildRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelDeployment Cancels a deployment resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CancelDeployment.go.html to see an example of how to use CancelDeployment API.
// A default retry strategy applies to this operation CancelDeployment()
func (client DevopsClient) CancelDeployment(ctx context.Context, request CancelDeploymentRequest) (response CancelDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelDeploymentResponse")
	}
	return
}

// cancelDeployment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) cancelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Deployment/CancelDeployment"
		err = common.PostProcessServiceError(err, "Devops", "CancelDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// CancelScheduledCascadingProjectDeletion Cascading operation that restores Project and child resources from a DELETING state to an active state
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CancelScheduledCascadingProjectDeletion.go.html to see an example of how to use CancelScheduledCascadingProjectDeletion API.
// A default retry strategy applies to this operation CancelScheduledCascadingProjectDeletion()
func (client DevopsClient) CancelScheduledCascadingProjectDeletion(ctx context.Context, request CancelScheduledCascadingProjectDeletionRequest) (response CancelScheduledCascadingProjectDeletionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelScheduledCascadingProjectDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelScheduledCascadingProjectDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelScheduledCascadingProjectDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelScheduledCascadingProjectDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelScheduledCascadingProjectDeletionResponse")
	}
	return
}

// cancelScheduledCascadingProjectDeletion implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) cancelScheduledCascadingProjectDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects/{projectId}/actions/cancelScheduledCascadingProjectDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelScheduledCascadingProjectDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Project/CancelScheduledCascadingProjectDeletion"
		err = common.PostProcessServiceError(err, "Devops", "CancelScheduledCascadingProjectDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeProjectCompartment Moves a project resource from one compartment OCID to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ChangeProjectCompartment.go.html to see an example of how to use ChangeProjectCompartment API.
// A default retry strategy applies to this operation ChangeProjectCompartment()
func (client DevopsClient) ChangeProjectCompartment(ctx context.Context, request ChangeProjectCompartmentRequest) (response ChangeProjectCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeProjectCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeProjectCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeProjectCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeProjectCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeProjectCompartmentResponse")
	}
	return
}

// changeProjectCompartment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) changeProjectCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects/{projectId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeProjectCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Project/ChangeProjectCompartment"
		err = common.PostProcessServiceError(err, "Devops", "ChangeProjectCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBuildPipeline Creates a new build pipeline.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateBuildPipeline.go.html to see an example of how to use CreateBuildPipeline API.
// A default retry strategy applies to this operation CreateBuildPipeline()
func (client DevopsClient) CreateBuildPipeline(ctx context.Context, request CreateBuildPipelineRequest) (response CreateBuildPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBuildPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBuildPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBuildPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBuildPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBuildPipelineResponse")
	}
	return
}

// createBuildPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createBuildPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/buildPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBuildPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipeline/CreateBuildPipeline"
		err = common.PostProcessServiceError(err, "Devops", "CreateBuildPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBuildPipelineStage Creates a new stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateBuildPipelineStage.go.html to see an example of how to use CreateBuildPipelineStage API.
// A default retry strategy applies to this operation CreateBuildPipelineStage()
func (client DevopsClient) CreateBuildPipelineStage(ctx context.Context, request CreateBuildPipelineStageRequest) (response CreateBuildPipelineStageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBuildPipelineStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBuildPipelineStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBuildPipelineStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBuildPipelineStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBuildPipelineStageResponse")
	}
	return
}

// createBuildPipelineStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createBuildPipelineStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/buildPipelineStages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBuildPipelineStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipelineStage/CreateBuildPipelineStage"
		err = common.PostProcessServiceError(err, "Devops", "CreateBuildPipelineStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &buildpipelinestage{})
	return response, err
}

// CreateBuildRun Starts a build pipeline run for a predefined build pipeline. Please ensure the completion of any work request for creation/updation of Build Pipeline before starting a Build Run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateBuildRun.go.html to see an example of how to use CreateBuildRun API.
// A default retry strategy applies to this operation CreateBuildRun()
func (client DevopsClient) CreateBuildRun(ctx context.Context, request CreateBuildRunRequest) (response CreateBuildRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBuildRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBuildRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBuildRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBuildRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBuildRunResponse")
	}
	return
}

// createBuildRun implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createBuildRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/buildRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBuildRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildRun/CreateBuildRun"
		err = common.PostProcessServiceError(err, "Devops", "CreateBuildRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConnection Creates a new connection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateConnection.go.html to see an example of how to use CreateConnection API.
// A default retry strategy applies to this operation CreateConnection()
func (client DevopsClient) CreateConnection(ctx context.Context, request CreateConnectionRequest) (response CreateConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionResponse")
	}
	return
}

// createConnection implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Connection/CreateConnection"
		err = common.PostProcessServiceError(err, "Devops", "CreateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// CreateDeployArtifact Creates a new deployment artifact.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployArtifact.go.html to see an example of how to use CreateDeployArtifact API.
// A default retry strategy applies to this operation CreateDeployArtifact()
func (client DevopsClient) CreateDeployArtifact(ctx context.Context, request CreateDeployArtifactRequest) (response CreateDeployArtifactResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDeployArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeployArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeployArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeployArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeployArtifactResponse")
	}
	return
}

// createDeployArtifact implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployArtifacts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDeployArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployArtifact/CreateDeployArtifact"
		err = common.PostProcessServiceError(err, "Devops", "CreateDeployArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDeployEnvironment Creates a new deployment environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployEnvironment.go.html to see an example of how to use CreateDeployEnvironment API.
// A default retry strategy applies to this operation CreateDeployEnvironment()
func (client DevopsClient) CreateDeployEnvironment(ctx context.Context, request CreateDeployEnvironmentRequest) (response CreateDeployEnvironmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDeployEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeployEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeployEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeployEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeployEnvironmentResponse")
	}
	return
}

// createDeployEnvironment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDeployEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployEnvironment/CreateDeployEnvironment"
		err = common.PostProcessServiceError(err, "Devops", "CreateDeployEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployenvironment{})
	return response, err
}

// CreateDeployPipeline Creates a new deployment pipeline.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployPipeline.go.html to see an example of how to use CreateDeployPipeline API.
// A default retry strategy applies to this operation CreateDeployPipeline()
func (client DevopsClient) CreateDeployPipeline(ctx context.Context, request CreateDeployPipelineRequest) (response CreateDeployPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDeployPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeployPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeployPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeployPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeployPipelineResponse")
	}
	return
}

// createDeployPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDeployPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployPipeline/CreateDeployPipeline"
		err = common.PostProcessServiceError(err, "Devops", "CreateDeployPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDeployStage Creates a new deployment stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployStage.go.html to see an example of how to use CreateDeployStage API.
// A default retry strategy applies to this operation CreateDeployStage()
func (client DevopsClient) CreateDeployStage(ctx context.Context, request CreateDeployStageRequest) (response CreateDeployStageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDeployStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeployStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeployStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeployStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeployStageResponse")
	}
	return
}

// createDeployStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployStages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDeployStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployStage/CreateDeployStage"
		err = common.PostProcessServiceError(err, "Devops", "CreateDeployStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deploystage{})
	return response, err
}

// CreateDeployment Creates a new deployment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployment.go.html to see an example of how to use CreateDeployment API.
// A default retry strategy applies to this operation CreateDeployment()
func (client DevopsClient) CreateDeployment(ctx context.Context, request CreateDeploymentRequest) (response CreateDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeploymentResponse")
	}
	return
}

// createDeployment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Deployment/CreateDeployment"
		err = common.PostProcessServiceError(err, "Devops", "CreateDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// CreateProject Creates a new project.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateProject.go.html to see an example of how to use CreateProject API.
// A default retry strategy applies to this operation CreateProject()
func (client DevopsClient) CreateProject(ctx context.Context, request CreateProjectRequest) (response CreateProjectResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProjectResponse")
	}
	return
}

// createProject implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Project/CreateProject"
		err = common.PostProcessServiceError(err, "Devops", "CreateProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRepository Creates a new repository.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateRepository.go.html to see an example of how to use CreateRepository API.
// A default retry strategy applies to this operation CreateRepository()
func (client DevopsClient) CreateRepository(ctx context.Context, request CreateRepositoryRequest) (response CreateRepositoryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRepositoryResponse")
	}
	return
}

// createRepository implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/repositories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRepositoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/CreateRepository"
		err = common.PostProcessServiceError(err, "Devops", "CreateRepository", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTrigger Creates a new trigger.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateTrigger.go.html to see an example of how to use CreateTrigger API.
// A default retry strategy applies to this operation CreateTrigger()
func (client DevopsClient) CreateTrigger(ctx context.Context, request CreateTriggerRequest) (response CreateTriggerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTrigger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTriggerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTriggerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTriggerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTriggerResponse")
	}
	return
}

// createTrigger implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) createTrigger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/triggers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTriggerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Trigger/CreateTrigger"
		err = common.PostProcessServiceError(err, "Devops", "CreateTrigger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &triggercreateresult{})
	return response, err
}

// DeleteBuildPipeline Deletes a build pipeline resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteBuildPipeline.go.html to see an example of how to use DeleteBuildPipeline API.
// A default retry strategy applies to this operation DeleteBuildPipeline()
func (client DevopsClient) DeleteBuildPipeline(ctx context.Context, request DeleteBuildPipelineRequest) (response DeleteBuildPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBuildPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBuildPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBuildPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBuildPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBuildPipelineResponse")
	}
	return
}

// deleteBuildPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteBuildPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/buildPipelines/{buildPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBuildPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipeline/DeleteBuildPipeline"
		err = common.PostProcessServiceError(err, "Devops", "DeleteBuildPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBuildPipelineStage Deletes a stage based on the stage ID provided in the request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteBuildPipelineStage.go.html to see an example of how to use DeleteBuildPipelineStage API.
// A default retry strategy applies to this operation DeleteBuildPipelineStage()
func (client DevopsClient) DeleteBuildPipelineStage(ctx context.Context, request DeleteBuildPipelineStageRequest) (response DeleteBuildPipelineStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBuildPipelineStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBuildPipelineStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBuildPipelineStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBuildPipelineStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBuildPipelineStageResponse")
	}
	return
}

// deleteBuildPipelineStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteBuildPipelineStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/buildPipelineStages/{buildPipelineStageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBuildPipelineStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipelineStage/DeleteBuildPipelineStage"
		err = common.PostProcessServiceError(err, "Devops", "DeleteBuildPipelineStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnection Deletes a connection resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteConnection.go.html to see an example of how to use DeleteConnection API.
// A default retry strategy applies to this operation DeleteConnection()
func (client DevopsClient) DeleteConnection(ctx context.Context, request DeleteConnectionRequest) (response DeleteConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionResponse")
	}
	return
}

// deleteConnection implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Connection/DeleteConnection"
		err = common.PostProcessServiceError(err, "Devops", "DeleteConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDeployArtifact Deletes a deployment artifact resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployArtifact.go.html to see an example of how to use DeleteDeployArtifact API.
// A default retry strategy applies to this operation DeleteDeployArtifact()
func (client DevopsClient) DeleteDeployArtifact(ctx context.Context, request DeleteDeployArtifactRequest) (response DeleteDeployArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDeployArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDeployArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDeployArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDeployArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDeployArtifactResponse")
	}
	return
}

// deleteDeployArtifact implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployArtifacts/{deployArtifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployArtifact/DeleteDeployArtifact"
		err = common.PostProcessServiceError(err, "Devops", "DeleteDeployArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDeployEnvironment Deletes a deployment environment resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployEnvironment.go.html to see an example of how to use DeleteDeployEnvironment API.
// A default retry strategy applies to this operation DeleteDeployEnvironment()
func (client DevopsClient) DeleteDeployEnvironment(ctx context.Context, request DeleteDeployEnvironmentRequest) (response DeleteDeployEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDeployEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDeployEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDeployEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDeployEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDeployEnvironmentResponse")
	}
	return
}

// deleteDeployEnvironment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployEnvironments/{deployEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployEnvironment/DeleteDeployEnvironment"
		err = common.PostProcessServiceError(err, "Devops", "DeleteDeployEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDeployPipeline Deletes a deployment pipeline resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployPipeline.go.html to see an example of how to use DeleteDeployPipeline API.
// A default retry strategy applies to this operation DeleteDeployPipeline()
func (client DevopsClient) DeleteDeployPipeline(ctx context.Context, request DeleteDeployPipelineRequest) (response DeleteDeployPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDeployPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDeployPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDeployPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDeployPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDeployPipelineResponse")
	}
	return
}

// deleteDeployPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployPipelines/{deployPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployPipeline/DeleteDeployPipeline"
		err = common.PostProcessServiceError(err, "Devops", "DeleteDeployPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDeployStage Deletes a deployment stage resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployStage.go.html to see an example of how to use DeleteDeployStage API.
// A default retry strategy applies to this operation DeleteDeployStage()
func (client DevopsClient) DeleteDeployStage(ctx context.Context, request DeleteDeployStageRequest) (response DeleteDeployStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDeployStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDeployStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDeployStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDeployStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDeployStageResponse")
	}
	return
}

// deleteDeployStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployStages/{deployStageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployStage/DeleteDeployStage"
		err = common.PostProcessServiceError(err, "Devops", "DeleteDeployStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProject Deletes a project resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteProject.go.html to see an example of how to use DeleteProject API.
// A default retry strategy applies to this operation DeleteProject()
func (client DevopsClient) DeleteProject(ctx context.Context, request DeleteProjectRequest) (response DeleteProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProjectResponse")
	}
	return
}

// deleteProject implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/projects/{projectId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Project/DeleteProject"
		err = common.PostProcessServiceError(err, "Devops", "DeleteProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRef Deletes a Repository's Ref by its name. Returns an error if the name is ambiguous. Can be disambiguated by using full names like "heads/<name>" or "tags/<name>".
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteRef.go.html to see an example of how to use DeleteRef API.
// A default retry strategy applies to this operation DeleteRef()
func (client DevopsClient) DeleteRef(ctx context.Context, request DeleteRefRequest) (response DeleteRefResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteRef, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRefResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRefResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRefResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRefResponse")
	}
	return
}

// deleteRef implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteRef(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/repositories/{repositoryId}/refs/{refName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRefResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/DeleteRef"
		err = common.PostProcessServiceError(err, "Devops", "DeleteRef", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRepository Deletes a repository resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteRepository.go.html to see an example of how to use DeleteRepository API.
// A default retry strategy applies to this operation DeleteRepository()
func (client DevopsClient) DeleteRepository(ctx context.Context, request DeleteRepositoryRequest) (response DeleteRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRepositoryResponse")
	}
	return
}

// deleteRepository implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/repositories/{repositoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRepositoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/DeleteRepository"
		err = common.PostProcessServiceError(err, "Devops", "DeleteRepository", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTrigger Deletes a trigger resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteTrigger.go.html to see an example of how to use DeleteTrigger API.
// A default retry strategy applies to this operation DeleteTrigger()
func (client DevopsClient) DeleteTrigger(ctx context.Context, request DeleteTriggerRequest) (response DeleteTriggerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTrigger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTriggerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTriggerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTriggerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTriggerResponse")
	}
	return
}

// deleteTrigger implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) deleteTrigger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/triggers/{triggerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTriggerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Trigger/DeleteTrigger"
		err = common.PostProcessServiceError(err, "Devops", "DeleteTrigger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBuildPipeline Retrieves a build pipeline by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetBuildPipeline.go.html to see an example of how to use GetBuildPipeline API.
// A default retry strategy applies to this operation GetBuildPipeline()
func (client DevopsClient) GetBuildPipeline(ctx context.Context, request GetBuildPipelineRequest) (response GetBuildPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBuildPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBuildPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBuildPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBuildPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBuildPipelineResponse")
	}
	return
}

// getBuildPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getBuildPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/buildPipelines/{buildPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBuildPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipeline/GetBuildPipeline"
		err = common.PostProcessServiceError(err, "Devops", "GetBuildPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBuildPipelineStage Retrieves a stage based on the stage ID provided in the request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetBuildPipelineStage.go.html to see an example of how to use GetBuildPipelineStage API.
// A default retry strategy applies to this operation GetBuildPipelineStage()
func (client DevopsClient) GetBuildPipelineStage(ctx context.Context, request GetBuildPipelineStageRequest) (response GetBuildPipelineStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBuildPipelineStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBuildPipelineStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBuildPipelineStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBuildPipelineStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBuildPipelineStageResponse")
	}
	return
}

// getBuildPipelineStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getBuildPipelineStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/buildPipelineStages/{buildPipelineStageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBuildPipelineStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipelineStage/GetBuildPipelineStage"
		err = common.PostProcessServiceError(err, "Devops", "GetBuildPipelineStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &buildpipelinestage{})
	return response, err
}

// GetBuildRun Returns the details of a build run for a given build run ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetBuildRun.go.html to see an example of how to use GetBuildRun API.
// A default retry strategy applies to this operation GetBuildRun()
func (client DevopsClient) GetBuildRun(ctx context.Context, request GetBuildRunRequest) (response GetBuildRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBuildRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBuildRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBuildRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBuildRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBuildRunResponse")
	}
	return
}

// getBuildRun implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getBuildRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/buildRuns/{buildRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBuildRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildRun/GetBuildRun"
		err = common.PostProcessServiceError(err, "Devops", "GetBuildRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCommit Retrieves a repository's commit by commit ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetCommit.go.html to see an example of how to use GetCommit API.
// A default retry strategy applies to this operation GetCommit()
func (client DevopsClient) GetCommit(ctx context.Context, request GetCommitRequest) (response GetCommitResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCommit, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCommitResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCommitResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCommitResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCommitResponse")
	}
	return
}

// getCommit implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getCommit(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/commits/{commitId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCommitResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetCommit"
		err = common.PostProcessServiceError(err, "Devops", "GetCommit", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCommitDiff Compares two revisions for their differences. Supports comparison between two references or commits.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetCommitDiff.go.html to see an example of how to use GetCommitDiff API.
// A default retry strategy applies to this operation GetCommitDiff()
func (client DevopsClient) GetCommitDiff(ctx context.Context, request GetCommitDiffRequest) (response GetCommitDiffResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCommitDiff, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCommitDiffResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCommitDiffResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCommitDiffResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCommitDiffResponse")
	}
	return
}

// getCommitDiff implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getCommitDiff(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/diff", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCommitDiffResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetCommitDiff"
		err = common.PostProcessServiceError(err, "Devops", "GetCommitDiff", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConnection Retrieves a connection by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetConnection.go.html to see an example of how to use GetConnection API.
// A default retry strategy applies to this operation GetConnection()
func (client DevopsClient) GetConnection(ctx context.Context, request GetConnectionRequest) (response GetConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionResponse")
	}
	return
}

// getConnection implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Connection/GetConnection"
		err = common.PostProcessServiceError(err, "Devops", "GetConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// GetDeployArtifact Retrieves a deployment artifact by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployArtifact.go.html to see an example of how to use GetDeployArtifact API.
// A default retry strategy applies to this operation GetDeployArtifact()
func (client DevopsClient) GetDeployArtifact(ctx context.Context, request GetDeployArtifactRequest) (response GetDeployArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeployArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeployArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeployArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeployArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeployArtifactResponse")
	}
	return
}

// getDeployArtifact implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployArtifacts/{deployArtifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeployArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployArtifact/GetDeployArtifact"
		err = common.PostProcessServiceError(err, "Devops", "GetDeployArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDeployEnvironment Retrieves a deployment environment by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployEnvironment.go.html to see an example of how to use GetDeployEnvironment API.
// A default retry strategy applies to this operation GetDeployEnvironment()
func (client DevopsClient) GetDeployEnvironment(ctx context.Context, request GetDeployEnvironmentRequest) (response GetDeployEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeployEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeployEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeployEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeployEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeployEnvironmentResponse")
	}
	return
}

// getDeployEnvironment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployEnvironments/{deployEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeployEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployEnvironment/GetDeployEnvironment"
		err = common.PostProcessServiceError(err, "Devops", "GetDeployEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployenvironment{})
	return response, err
}

// GetDeployPipeline Retrieves a deployment pipeline by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployPipeline.go.html to see an example of how to use GetDeployPipeline API.
// A default retry strategy applies to this operation GetDeployPipeline()
func (client DevopsClient) GetDeployPipeline(ctx context.Context, request GetDeployPipelineRequest) (response GetDeployPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeployPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeployPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeployPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeployPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeployPipelineResponse")
	}
	return
}

// getDeployPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployPipelines/{deployPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeployPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployPipeline/GetDeployPipeline"
		err = common.PostProcessServiceError(err, "Devops", "GetDeployPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDeployStage Retrieves a deployment stage by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployStage.go.html to see an example of how to use GetDeployStage API.
// A default retry strategy applies to this operation GetDeployStage()
func (client DevopsClient) GetDeployStage(ctx context.Context, request GetDeployStageRequest) (response GetDeployStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeployStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeployStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeployStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeployStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeployStageResponse")
	}
	return
}

// getDeployStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployStages/{deployStageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeployStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployStage/GetDeployStage"
		err = common.PostProcessServiceError(err, "Devops", "GetDeployStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deploystage{})
	return response, err
}

// GetDeployment Retrieves a deployment by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployment.go.html to see an example of how to use GetDeployment API.
// A default retry strategy applies to this operation GetDeployment()
func (client DevopsClient) GetDeployment(ctx context.Context, request GetDeploymentRequest) (response GetDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeploymentResponse")
	}
	return
}

// getDeployment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Deployment/GetDeployment"
		err = common.PostProcessServiceError(err, "Devops", "GetDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// GetFileDiff Gets the line-by-line difference between file on different commits. This API will be deprecated on Wed, 29 Mar 2023 01:00:00 GMT as it does not get recognized when filePath has '/'. This will be replaced by "/repositories/{repositoryId}/file/diffs"
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetFileDiff.go.html to see an example of how to use GetFileDiff API.
// A default retry strategy applies to this operation GetFileDiff()
func (client DevopsClient) GetFileDiff(ctx context.Context, request GetFileDiffRequest) (response GetFileDiffResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFileDiff, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFileDiffResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFileDiffResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFileDiffResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFileDiffResponse")
	}
	return
}

// getFileDiff implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getFileDiff(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/diffs/{filePath}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFileDiffResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetFileDiff"
		err = common.PostProcessServiceError(err, "Devops", "GetFileDiff", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMirrorRecord Returns either current mirror record or last successful mirror record for a specific mirror repository.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetMirrorRecord.go.html to see an example of how to use GetMirrorRecord API.
// A default retry strategy applies to this operation GetMirrorRecord()
func (client DevopsClient) GetMirrorRecord(ctx context.Context, request GetMirrorRecordRequest) (response GetMirrorRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMirrorRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMirrorRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMirrorRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMirrorRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMirrorRecordResponse")
	}
	return
}

// getMirrorRecord implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getMirrorRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/mirrorRecords/{mirrorRecordType}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMirrorRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetMirrorRecord"
		err = common.PostProcessServiceError(err, "Devops", "GetMirrorRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetObject Retrieves blob of specific branch name/commit ID and file path.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetObject.go.html to see an example of how to use GetObject API.
// A default retry strategy applies to this operation GetObject()
func (client DevopsClient) GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getObject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetObjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetObjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetObjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetObjectResponse")
	}
	return
}

// getObject implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getObject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/object", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/RepositoryObject/GetObject"
		err = common.PostProcessServiceError(err, "Devops", "GetObject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetObjectContent Retrieve contents of a specified object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetObjectContent.go.html to see an example of how to use GetObjectContent API.
// A default retry strategy applies to this operation GetObjectContent()
func (client DevopsClient) GetObjectContent(ctx context.Context, request GetObjectContentRequest) (response GetObjectContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getObjectContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetObjectContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetObjectContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetObjectContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetObjectContentResponse")
	}
	return
}

// getObjectContent implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getObjectContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/objects/{sha}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetObjectContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetObjectContent"
		err = common.PostProcessServiceError(err, "Devops", "GetObjectContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProject Retrieves a project by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetProject.go.html to see an example of how to use GetProject API.
// A default retry strategy applies to this operation GetProject()
func (client DevopsClient) GetProject(ctx context.Context, request GetProjectRequest) (response GetProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProjectResponse")
	}
	return
}

// getProject implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/projects/{projectId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Project/GetProject"
		err = common.PostProcessServiceError(err, "Devops", "GetProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRef Retrieves a repository's reference by its name with preference for branches over tags if the name is ambiguous. This can be disambiguated by using full names like "heads/<name>" or "tags/<name>".
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetRef.go.html to see an example of how to use GetRef API.
// A default retry strategy applies to this operation GetRef()
func (client DevopsClient) GetRef(ctx context.Context, request GetRefRequest) (response GetRefResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRef, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRefResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRefResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRefResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRefResponse")
	}
	return
}

// getRef implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getRef(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/refs/{refName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRefResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetRef"
		err = common.PostProcessServiceError(err, "Devops", "GetRef", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &repositoryref{})
	return response, err
}

// GetRepoFileDiff Gets the line-by-line difference between file on different commits.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetRepoFileDiff.go.html to see an example of how to use GetRepoFileDiff API.
// A default retry strategy applies to this operation GetRepoFileDiff()
func (client DevopsClient) GetRepoFileDiff(ctx context.Context, request GetRepoFileDiffRequest) (response GetRepoFileDiffResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRepoFileDiff, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRepoFileDiffResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRepoFileDiffResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRepoFileDiffResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRepoFileDiffResponse")
	}
	return
}

// getRepoFileDiff implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getRepoFileDiff(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/file/diffs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRepoFileDiffResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetRepoFileDiff"
		err = common.PostProcessServiceError(err, "Devops", "GetRepoFileDiff", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRepoFileLines Retrieve lines of a specified file. Supports starting line number and limit.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetRepoFileLines.go.html to see an example of how to use GetRepoFileLines API.
// A default retry strategy applies to this operation GetRepoFileLines()
func (client DevopsClient) GetRepoFileLines(ctx context.Context, request GetRepoFileLinesRequest) (response GetRepoFileLinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRepoFileLines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRepoFileLinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRepoFileLinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRepoFileLinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRepoFileLinesResponse")
	}
	return
}

// getRepoFileLines implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getRepoFileLines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/file/lines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRepoFileLinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetRepoFileLines"
		err = common.PostProcessServiceError(err, "Devops", "GetRepoFileLines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRepository Retrieves a repository by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetRepository.go.html to see an example of how to use GetRepository API.
// A default retry strategy applies to this operation GetRepository()
func (client DevopsClient) GetRepository(ctx context.Context, request GetRepositoryRequest) (response GetRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRepositoryResponse")
	}
	return
}

// getRepository implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRepositoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetRepository"
		err = common.PostProcessServiceError(err, "Devops", "GetRepository", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRepositoryArchiveContent Returns the archived repository information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetRepositoryArchiveContent.go.html to see an example of how to use GetRepositoryArchiveContent API.
// A default retry strategy applies to this operation GetRepositoryArchiveContent()
func (client DevopsClient) GetRepositoryArchiveContent(ctx context.Context, request GetRepositoryArchiveContentRequest) (response GetRepositoryArchiveContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRepositoryArchiveContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRepositoryArchiveContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRepositoryArchiveContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRepositoryArchiveContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRepositoryArchiveContentResponse")
	}
	return
}

// getRepositoryArchiveContent implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getRepositoryArchiveContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/archive/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRepositoryArchiveContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetRepositoryArchiveContent"
		err = common.PostProcessServiceError(err, "Devops", "GetRepositoryArchiveContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRepositoryFileLines Retrieve lines of a specified file. Supports starting line number and limit. This API will be deprecated on Wed, 29 Mar 2023 01:00:00 GMT as it does not get recognized when filePath has '/'. This will be replaced by "/repositories/{repositoryId}/file/lines"
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetRepositoryFileLines.go.html to see an example of how to use GetRepositoryFileLines API.
// A default retry strategy applies to this operation GetRepositoryFileLines()
func (client DevopsClient) GetRepositoryFileLines(ctx context.Context, request GetRepositoryFileLinesRequest) (response GetRepositoryFileLinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRepositoryFileLines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRepositoryFileLinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRepositoryFileLinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRepositoryFileLinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRepositoryFileLinesResponse")
	}
	return
}

// getRepositoryFileLines implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getRepositoryFileLines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/files/{filePath}/lines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRepositoryFileLinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/GetRepositoryFileLines"
		err = common.PostProcessServiceError(err, "Devops", "GetRepositoryFileLines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTrigger Retrieves a trigger by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetTrigger.go.html to see an example of how to use GetTrigger API.
// A default retry strategy applies to this operation GetTrigger()
func (client DevopsClient) GetTrigger(ctx context.Context, request GetTriggerRequest) (response GetTriggerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTrigger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTriggerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTriggerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTriggerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTriggerResponse")
	}
	return
}

// getTrigger implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) getTrigger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/triggers/{triggerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTriggerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Trigger/GetTrigger"
		err = common.PostProcessServiceError(err, "Devops", "GetTrigger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &trigger{})
	return response, err
}

// GetWorkRequest Retrieves the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DevopsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DevopsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "Devops", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAuthors Retrieve a list of all the authors.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListAuthors.go.html to see an example of how to use ListAuthors API.
// A default retry strategy applies to this operation ListAuthors()
func (client DevopsClient) ListAuthors(ctx context.Context, request ListAuthorsRequest) (response ListAuthorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuthors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuthorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuthorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuthorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuthorsResponse")
	}
	return
}

// listAuthors implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listAuthors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/authors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuthorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/ListAuthors"
		err = common.PostProcessServiceError(err, "Devops", "ListAuthors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBuildPipelineStages Returns a list of all stages in a compartment or build pipeline.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListBuildPipelineStages.go.html to see an example of how to use ListBuildPipelineStages API.
// A default retry strategy applies to this operation ListBuildPipelineStages()
func (client DevopsClient) ListBuildPipelineStages(ctx context.Context, request ListBuildPipelineStagesRequest) (response ListBuildPipelineStagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBuildPipelineStages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBuildPipelineStagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBuildPipelineStagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBuildPipelineStagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBuildPipelineStagesResponse")
	}
	return
}

// listBuildPipelineStages implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listBuildPipelineStages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/buildPipelineStages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBuildPipelineStagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipelineStageSummary/ListBuildPipelineStages"
		err = common.PostProcessServiceError(err, "Devops", "ListBuildPipelineStages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBuildPipelines Returns a list of build pipelines.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListBuildPipelines.go.html to see an example of how to use ListBuildPipelines API.
// A default retry strategy applies to this operation ListBuildPipelines()
func (client DevopsClient) ListBuildPipelines(ctx context.Context, request ListBuildPipelinesRequest) (response ListBuildPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBuildPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBuildPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBuildPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBuildPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBuildPipelinesResponse")
	}
	return
}

// listBuildPipelines implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listBuildPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/buildPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBuildPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipelineCollection/ListBuildPipelines"
		err = common.PostProcessServiceError(err, "Devops", "ListBuildPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBuildRuns Returns a list of build run summary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListBuildRuns.go.html to see an example of how to use ListBuildRuns API.
// A default retry strategy applies to this operation ListBuildRuns()
func (client DevopsClient) ListBuildRuns(ctx context.Context, request ListBuildRunsRequest) (response ListBuildRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBuildRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBuildRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBuildRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBuildRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBuildRunsResponse")
	}
	return
}

// listBuildRuns implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listBuildRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/buildRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBuildRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildRunSummary/ListBuildRuns"
		err = common.PostProcessServiceError(err, "Devops", "ListBuildRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCommitDiffs Compares two revisions and lists the differences. Supports comparison between two references or commits.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListCommitDiffs.go.html to see an example of how to use ListCommitDiffs API.
// A default retry strategy applies to this operation ListCommitDiffs()
func (client DevopsClient) ListCommitDiffs(ctx context.Context, request ListCommitDiffsRequest) (response ListCommitDiffsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCommitDiffs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCommitDiffsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCommitDiffsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCommitDiffsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCommitDiffsResponse")
	}
	return
}

// listCommitDiffs implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listCommitDiffs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/diffs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCommitDiffsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/ListCommitDiffs"
		err = common.PostProcessServiceError(err, "Devops", "ListCommitDiffs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCommits Returns a list of commits.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListCommits.go.html to see an example of how to use ListCommits API.
// A default retry strategy applies to this operation ListCommits()
func (client DevopsClient) ListCommits(ctx context.Context, request ListCommitsRequest) (response ListCommitsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCommits, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCommitsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCommitsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCommitsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCommitsResponse")
	}
	return
}

// listCommits implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listCommits(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/commits", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCommitsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/RepositoryCommit/ListCommits"
		err = common.PostProcessServiceError(err, "Devops", "ListCommits", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnections Returns a list of connections.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListConnections.go.html to see an example of how to use ListConnections API.
// A default retry strategy applies to this operation ListConnections()
func (client DevopsClient) ListConnections(ctx context.Context, request ListConnectionsRequest) (response ListConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionsResponse")
	}
	return
}

// listConnections implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/ConnectionCollection/ListConnections"
		err = common.PostProcessServiceError(err, "Devops", "ListConnections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeployArtifacts Returns a list of deployment artifacts.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployArtifacts.go.html to see an example of how to use ListDeployArtifacts API.
// A default retry strategy applies to this operation ListDeployArtifacts()
func (client DevopsClient) ListDeployArtifacts(ctx context.Context, request ListDeployArtifactsRequest) (response ListDeployArtifactsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeployArtifacts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeployArtifactsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeployArtifactsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeployArtifactsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeployArtifactsResponse")
	}
	return
}

// listDeployArtifacts implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listDeployArtifacts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployArtifacts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeployArtifactsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployArtifactSummary/ListDeployArtifacts"
		err = common.PostProcessServiceError(err, "Devops", "ListDeployArtifacts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeployEnvironments Returns a list of deployment environments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployEnvironments.go.html to see an example of how to use ListDeployEnvironments API.
// A default retry strategy applies to this operation ListDeployEnvironments()
func (client DevopsClient) ListDeployEnvironments(ctx context.Context, request ListDeployEnvironmentsRequest) (response ListDeployEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeployEnvironments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeployEnvironmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeployEnvironmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeployEnvironmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeployEnvironmentsResponse")
	}
	return
}

// listDeployEnvironments implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listDeployEnvironments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeployEnvironmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployEnvironmentSummary/ListDeployEnvironments"
		err = common.PostProcessServiceError(err, "Devops", "ListDeployEnvironments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeployPipelines Returns a list of deployment pipelines.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployPipelines.go.html to see an example of how to use ListDeployPipelines API.
// A default retry strategy applies to this operation ListDeployPipelines()
func (client DevopsClient) ListDeployPipelines(ctx context.Context, request ListDeployPipelinesRequest) (response ListDeployPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeployPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeployPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeployPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeployPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeployPipelinesResponse")
	}
	return
}

// listDeployPipelines implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listDeployPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeployPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployPipelineSummary/ListDeployPipelines"
		err = common.PostProcessServiceError(err, "Devops", "ListDeployPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeployStages Retrieves a list of deployment stages.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployStages.go.html to see an example of how to use ListDeployStages API.
// A default retry strategy applies to this operation ListDeployStages()
func (client DevopsClient) ListDeployStages(ctx context.Context, request ListDeployStagesRequest) (response ListDeployStagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeployStages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeployStagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeployStagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeployStagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeployStagesResponse")
	}
	return
}

// listDeployStages implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listDeployStages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployStages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeployStagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployStageSummary/ListDeployStages"
		err = common.PostProcessServiceError(err, "Devops", "ListDeployStages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeployments Returns a list of deployments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployments.go.html to see an example of how to use ListDeployments API.
// A default retry strategy applies to this operation ListDeployments()
func (client DevopsClient) ListDeployments(ctx context.Context, request ListDeploymentsRequest) (response ListDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeployments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentsResponse")
	}
	return
}

// listDeployments implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listDeployments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeploymentSummary/ListDeployments"
		err = common.PostProcessServiceError(err, "Devops", "ListDeployments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMirrorRecords Returns a list of mirror entry in history within 30 days.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListMirrorRecords.go.html to see an example of how to use ListMirrorRecords API.
// A default retry strategy applies to this operation ListMirrorRecords()
func (client DevopsClient) ListMirrorRecords(ctx context.Context, request ListMirrorRecordsRequest) (response ListMirrorRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMirrorRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMirrorRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMirrorRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMirrorRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMirrorRecordsResponse")
	}
	return
}

// listMirrorRecords implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listMirrorRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/mirrorRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMirrorRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/ListMirrorRecords"
		err = common.PostProcessServiceError(err, "Devops", "ListMirrorRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPaths Retrieves a list of files and directories in a repository.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListPaths.go.html to see an example of how to use ListPaths API.
// A default retry strategy applies to this operation ListPaths()
func (client DevopsClient) ListPaths(ctx context.Context, request ListPathsRequest) (response ListPathsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPaths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPathsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPathsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPathsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPathsResponse")
	}
	return
}

// listPaths implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listPaths(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/paths", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPathsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/RepositoryPathSummary/ListPaths"
		err = common.PostProcessServiceError(err, "Devops", "ListPaths", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProjects Returns a list of projects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListProjects.go.html to see an example of how to use ListProjects API.
// A default retry strategy applies to this operation ListProjects()
func (client DevopsClient) ListProjects(ctx context.Context, request ListProjectsRequest) (response ListProjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProjects, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProjectsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProjectsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProjectsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProjectsResponse")
	}
	return
}

// listProjects implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listProjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/projects", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/ProjectSummary/ListProjects"
		err = common.PostProcessServiceError(err, "Devops", "ListProjects", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRefs Returns a list of references.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListRefs.go.html to see an example of how to use ListRefs API.
// A default retry strategy applies to this operation ListRefs()
func (client DevopsClient) ListRefs(ctx context.Context, request ListRefsRequest) (response ListRefsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRefs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRefsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRefsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRefsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRefsResponse")
	}
	return
}

// listRefs implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listRefs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories/{repositoryId}/refs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRefsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/RepositoryRef/ListRefs"
		err = common.PostProcessServiceError(err, "Devops", "ListRefs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRepositories Returns a list of repositories given a compartment ID or a project ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListRepositories.go.html to see an example of how to use ListRepositories API.
// A default retry strategy applies to this operation ListRepositories()
func (client DevopsClient) ListRepositories(ctx context.Context, request ListRepositoriesRequest) (response ListRepositoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRepositories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRepositoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRepositoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRepositoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRepositoriesResponse")
	}
	return
}

// listRepositories implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listRepositories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/repositories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRepositoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/ListRepositories"
		err = common.PostProcessServiceError(err, "Devops", "ListRepositories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTriggers Returns a list of triggers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListTriggers.go.html to see an example of how to use ListTriggers API.
// A default retry strategy applies to this operation ListTriggers()
func (client DevopsClient) ListTriggers(ctx context.Context, request ListTriggersRequest) (response ListTriggersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTriggers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTriggersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTriggersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTriggersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTriggersResponse")
	}
	return
}

// listTriggers implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) listTriggers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/triggers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTriggersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/TriggerCollection/ListTriggers"
		err = common.PostProcessServiceError(err, "Devops", "ListTriggers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DevopsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DevopsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Devops", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DevopsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DevopsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Devops", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DevopsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DevopsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "Devops", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MirrorRepository Synchronize a mirrored repository to the latest version from external providers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/MirrorRepository.go.html to see an example of how to use MirrorRepository API.
// A default retry strategy applies to this operation MirrorRepository()
func (client DevopsClient) MirrorRepository(ctx context.Context, request MirrorRepositoryRequest) (response MirrorRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.mirrorRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MirrorRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MirrorRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MirrorRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MirrorRepositoryResponse")
	}
	return
}

// mirrorRepository implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) mirrorRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/repositories/{repositoryId}/actions/mirror", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MirrorRepositoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/MirrorRepository"
		err = common.PostProcessServiceError(err, "Devops", "MirrorRepository", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutRepositoryRef Creates a new reference or updates an existing one.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/PutRepositoryRef.go.html to see an example of how to use PutRepositoryRef API.
// A default retry strategy applies to this operation PutRepositoryRef()
func (client DevopsClient) PutRepositoryRef(ctx context.Context, request PutRepositoryRefRequest) (response PutRepositoryRefResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.putRepositoryRef, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutRepositoryRefResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutRepositoryRefResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutRepositoryRefResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutRepositoryRefResponse")
	}
	return
}

// putRepositoryRef implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) putRepositoryRef(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/repositories/{repositoryId}/refs/{refName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutRepositoryRefResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/PutRepositoryRef"
		err = common.PostProcessServiceError(err, "Devops", "PutRepositoryRef", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &repositoryref{})
	return response, err
}

// ScheduleCascadingProjectDeletion Cascading operation that marks Project and child DevOps resources in a DELETING state for a retention period
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ScheduleCascadingProjectDeletion.go.html to see an example of how to use ScheduleCascadingProjectDeletion API.
// A default retry strategy applies to this operation ScheduleCascadingProjectDeletion()
func (client DevopsClient) ScheduleCascadingProjectDeletion(ctx context.Context, request ScheduleCascadingProjectDeletionRequest) (response ScheduleCascadingProjectDeletionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scheduleCascadingProjectDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleCascadingProjectDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleCascadingProjectDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleCascadingProjectDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleCascadingProjectDeletionResponse")
	}
	return
}

// scheduleCascadingProjectDeletion implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) scheduleCascadingProjectDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects/{projectId}/actions/scheduleCascadingProjectDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleCascadingProjectDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Project/ScheduleCascadingProjectDeletion"
		err = common.PostProcessServiceError(err, "Devops", "ScheduleCascadingProjectDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBuildPipeline Updates the build pipeline.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateBuildPipeline.go.html to see an example of how to use UpdateBuildPipeline API.
// A default retry strategy applies to this operation UpdateBuildPipeline()
func (client DevopsClient) UpdateBuildPipeline(ctx context.Context, request UpdateBuildPipelineRequest) (response UpdateBuildPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBuildPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBuildPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBuildPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBuildPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBuildPipelineResponse")
	}
	return
}

// updateBuildPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateBuildPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/buildPipelines/{buildPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBuildPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipeline/UpdateBuildPipeline"
		err = common.PostProcessServiceError(err, "Devops", "UpdateBuildPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBuildPipelineStage Updates the stage based on the stage ID provided in the request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateBuildPipelineStage.go.html to see an example of how to use UpdateBuildPipelineStage API.
// A default retry strategy applies to this operation UpdateBuildPipelineStage()
func (client DevopsClient) UpdateBuildPipelineStage(ctx context.Context, request UpdateBuildPipelineStageRequest) (response UpdateBuildPipelineStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBuildPipelineStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBuildPipelineStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBuildPipelineStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBuildPipelineStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBuildPipelineStageResponse")
	}
	return
}

// updateBuildPipelineStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateBuildPipelineStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/buildPipelineStages/{buildPipelineStageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBuildPipelineStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildPipelineStage/UpdateBuildPipelineStage"
		err = common.PostProcessServiceError(err, "Devops", "UpdateBuildPipelineStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &buildpipelinestage{})
	return response, err
}

// UpdateBuildRun Updates the build run.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateBuildRun.go.html to see an example of how to use UpdateBuildRun API.
// A default retry strategy applies to this operation UpdateBuildRun()
func (client DevopsClient) UpdateBuildRun(ctx context.Context, request UpdateBuildRunRequest) (response UpdateBuildRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBuildRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBuildRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBuildRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBuildRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBuildRunResponse")
	}
	return
}

// updateBuildRun implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateBuildRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/buildRuns/{buildRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBuildRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/BuildRun/UpdateBuildRun"
		err = common.PostProcessServiceError(err, "Devops", "UpdateBuildRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConnection Updates the connection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateConnection.go.html to see an example of how to use UpdateConnection API.
// A default retry strategy applies to this operation UpdateConnection()
func (client DevopsClient) UpdateConnection(ctx context.Context, request UpdateConnectionRequest) (response UpdateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConnectionResponse")
	}
	return
}

// updateConnection implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Connection/UpdateConnection"
		err = common.PostProcessServiceError(err, "Devops", "UpdateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// UpdateDeployArtifact Updates the deployment artifact.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployArtifact.go.html to see an example of how to use UpdateDeployArtifact API.
// A default retry strategy applies to this operation UpdateDeployArtifact()
func (client DevopsClient) UpdateDeployArtifact(ctx context.Context, request UpdateDeployArtifactRequest) (response UpdateDeployArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeployArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeployArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeployArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeployArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeployArtifactResponse")
	}
	return
}

// updateDeployArtifact implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployArtifacts/{deployArtifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployArtifactResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployArtifact/UpdateDeployArtifact"
		err = common.PostProcessServiceError(err, "Devops", "UpdateDeployArtifact", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDeployEnvironment Updates the deployment environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployEnvironment.go.html to see an example of how to use UpdateDeployEnvironment API.
// A default retry strategy applies to this operation UpdateDeployEnvironment()
func (client DevopsClient) UpdateDeployEnvironment(ctx context.Context, request UpdateDeployEnvironmentRequest) (response UpdateDeployEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeployEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeployEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeployEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeployEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeployEnvironmentResponse")
	}
	return
}

// updateDeployEnvironment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployEnvironments/{deployEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployEnvironment/UpdateDeployEnvironment"
		err = common.PostProcessServiceError(err, "Devops", "UpdateDeployEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployenvironment{})
	return response, err
}

// UpdateDeployPipeline Updates the deployment pipeline.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployPipeline.go.html to see an example of how to use UpdateDeployPipeline API.
// A default retry strategy applies to this operation UpdateDeployPipeline()
func (client DevopsClient) UpdateDeployPipeline(ctx context.Context, request UpdateDeployPipelineRequest) (response UpdateDeployPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeployPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeployPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeployPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeployPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeployPipelineResponse")
	}
	return
}

// updateDeployPipeline implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployPipelines/{deployPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployPipeline/UpdateDeployPipeline"
		err = common.PostProcessServiceError(err, "Devops", "UpdateDeployPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDeployStage Updates the deployment stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployStage.go.html to see an example of how to use UpdateDeployStage API.
// A default retry strategy applies to this operation UpdateDeployStage()
func (client DevopsClient) UpdateDeployStage(ctx context.Context, request UpdateDeployStageRequest) (response UpdateDeployStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeployStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeployStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeployStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeployStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeployStageResponse")
	}
	return
}

// updateDeployStage implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployStages/{deployStageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/DeployStage/UpdateDeployStage"
		err = common.PostProcessServiceError(err, "Devops", "UpdateDeployStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deploystage{})
	return response, err
}

// UpdateDeployment Updates the deployment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployment.go.html to see an example of how to use UpdateDeployment API.
// A default retry strategy applies to this operation UpdateDeployment()
func (client DevopsClient) UpdateDeployment(ctx context.Context, request UpdateDeploymentRequest) (response UpdateDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeploymentResponse")
	}
	return
}

// updateDeployment implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployments/{deploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Deployment/UpdateDeployment"
		err = common.PostProcessServiceError(err, "Devops", "UpdateDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// UpdateProject Updates the project.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateProject.go.html to see an example of how to use UpdateProject API.
// A default retry strategy applies to this operation UpdateProject()
func (client DevopsClient) UpdateProject(ctx context.Context, request UpdateProjectRequest) (response UpdateProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProject, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProjectResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProjectResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProjectResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProjectResponse")
	}
	return
}

// updateProject implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/projects/{projectId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Project/UpdateProject"
		err = common.PostProcessServiceError(err, "Devops", "UpdateProject", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRepository Updates the repository.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateRepository.go.html to see an example of how to use UpdateRepository API.
// A default retry strategy applies to this operation UpdateRepository()
func (client DevopsClient) UpdateRepository(ctx context.Context, request UpdateRepositoryRequest) (response UpdateRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRepositoryResponse")
	}
	return
}

// updateRepository implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/repositories/{repositoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRepositoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Repository/UpdateRepository"
		err = common.PostProcessServiceError(err, "Devops", "UpdateRepository", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTrigger Updates the trigger.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateTrigger.go.html to see an example of how to use UpdateTrigger API.
// A default retry strategy applies to this operation UpdateTrigger()
func (client DevopsClient) UpdateTrigger(ctx context.Context, request UpdateTriggerRequest) (response UpdateTriggerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTrigger, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTriggerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTriggerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTriggerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTriggerResponse")
	}
	return
}

// updateTrigger implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) updateTrigger(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/triggers/{triggerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTriggerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Trigger/UpdateTrigger"
		err = common.PostProcessServiceError(err, "Devops", "UpdateTrigger", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &trigger{})
	return response, err
}

// ValidateConnection Return whether the credentials of the connection are valid.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ValidateConnection.go.html to see an example of how to use ValidateConnection API.
// A default retry strategy applies to this operation ValidateConnection()
func (client DevopsClient) ValidateConnection(ctx context.Context, request ValidateConnectionRequest) (response ValidateConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateConnectionResponse")
	}
	return
}

// validateConnection implements the OCIOperation interface (enables retrying operations)
func (client DevopsClient) validateConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections/{connectionId}/actions/validate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/devops/20210630/Connection/ValidateConnection"
		err = common.PostProcessServiceError(err, "Devops", "ValidateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}
