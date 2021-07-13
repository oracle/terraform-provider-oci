// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/common/auth"
	"net/http"
)

//DevopsClient a client for Devops
type DevopsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDevopsClientWithConfigurationProvider Creates a new default Devops client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDevopsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DevopsClient, err error) {
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
//  as well as reading the region
func NewDevopsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DevopsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDevopsClientFromBaseClient(baseClient, configProvider)
}

func newDevopsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DevopsClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DevopsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ApproveDeployment Submit stage approval.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ApproveDeployment.go.html to see an example of how to use ApproveDeployment API.
func (client DevopsClient) ApproveDeployment(ctx context.Context, request ApproveDeploymentRequest) (response ApproveDeploymentResponse, err error) {
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
func (client DevopsClient) approveDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/approve", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ApproveDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// CancelDeployment Cancels a deployment resource by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CancelDeployment.go.html to see an example of how to use CancelDeployment API.
func (client DevopsClient) CancelDeployment(ctx context.Context, request CancelDeploymentRequest) (response CancelDeploymentResponse, err error) {
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
func (client DevopsClient) cancelDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/cancel", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CancelDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// ChangeProjectCompartment Moves a project resource from one compartment OCID to another.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ChangeProjectCompartment.go.html to see an example of how to use ChangeProjectCompartment API.
func (client DevopsClient) ChangeProjectCompartment(ctx context.Context, request ChangeProjectCompartmentRequest) (response ChangeProjectCompartmentResponse, err error) {
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
func (client DevopsClient) changeProjectCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects/{projectId}/actions/changeCompartment", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ChangeProjectCompartmentResponse
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

// CreateDeployArtifact Creates a new deployment artifact.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployArtifact.go.html to see an example of how to use CreateDeployArtifact API.
func (client DevopsClient) CreateDeployArtifact(ctx context.Context, request CreateDeployArtifactRequest) (response CreateDeployArtifactResponse, err error) {
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
func (client DevopsClient) createDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployArtifacts", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDeployArtifactResponse
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

// CreateDeployEnvironment Creates a new deployment environment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployEnvironment.go.html to see an example of how to use CreateDeployEnvironment API.
func (client DevopsClient) CreateDeployEnvironment(ctx context.Context, request CreateDeployEnvironmentRequest) (response CreateDeployEnvironmentResponse, err error) {
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
func (client DevopsClient) createDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployEnvironments", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDeployEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployenvironment{})
	return response, err
}

// CreateDeployPipeline Creates a new deployment pipeline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployPipeline.go.html to see an example of how to use CreateDeployPipeline API.
func (client DevopsClient) CreateDeployPipeline(ctx context.Context, request CreateDeployPipelineRequest) (response CreateDeployPipelineResponse, err error) {
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
func (client DevopsClient) createDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployPipelines", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDeployPipelineResponse
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

// CreateDeployStage Creates a new deployment stage.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployStage.go.html to see an example of how to use CreateDeployStage API.
func (client DevopsClient) CreateDeployStage(ctx context.Context, request CreateDeployStageRequest) (response CreateDeployStageResponse, err error) {
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
func (client DevopsClient) createDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployStages", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDeployStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deploystage{})
	return response, err
}

// CreateDeployment Creates a new deployment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateDeployment.go.html to see an example of how to use CreateDeployment API.
func (client DevopsClient) CreateDeployment(ctx context.Context, request CreateDeploymentRequest) (response CreateDeploymentResponse, err error) {
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
func (client DevopsClient) createDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// CreateProject Creates a new project.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/CreateProject.go.html to see an example of how to use CreateProject API.
func (client DevopsClient) CreateProject(ctx context.Context, request CreateProjectRequest) (response CreateProjectResponse, err error) {
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
func (client DevopsClient) createProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/projects", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response CreateProjectResponse
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

// DeleteDeployArtifact Deletes a deployment artifact resource by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployArtifact.go.html to see an example of how to use DeleteDeployArtifact API.
func (client DevopsClient) DeleteDeployArtifact(ctx context.Context, request DeleteDeployArtifactRequest) (response DeleteDeployArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) deleteDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployArtifacts/{deployArtifactId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployArtifactResponse
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

// DeleteDeployEnvironment Deletes a deployment environment resource by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployEnvironment.go.html to see an example of how to use DeleteDeployEnvironment API.
func (client DevopsClient) DeleteDeployEnvironment(ctx context.Context, request DeleteDeployEnvironmentRequest) (response DeleteDeployEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) deleteDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployEnvironments/{deployEnvironmentId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployEnvironmentResponse
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

// DeleteDeployPipeline Deletes a deployment pipeline resource by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployPipeline.go.html to see an example of how to use DeleteDeployPipeline API.
func (client DevopsClient) DeleteDeployPipeline(ctx context.Context, request DeleteDeployPipelineRequest) (response DeleteDeployPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) deleteDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployPipelines/{deployPipelineId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployPipelineResponse
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

// DeleteDeployStage Deletes a deployment stage resource by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteDeployStage.go.html to see an example of how to use DeleteDeployStage API.
func (client DevopsClient) DeleteDeployStage(ctx context.Context, request DeleteDeployStageRequest) (response DeleteDeployStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) deleteDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployStages/{deployStageId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteDeployStageResponse
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

// DeleteProject Deletes a project resource by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/DeleteProject.go.html to see an example of how to use DeleteProject API.
func (client DevopsClient) DeleteProject(ctx context.Context, request DeleteProjectRequest) (response DeleteProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) deleteProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/projects/{projectId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response DeleteProjectResponse
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

// GetDeployArtifact Retrieves a deployment artifact by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployArtifact.go.html to see an example of how to use GetDeployArtifact API.
func (client DevopsClient) GetDeployArtifact(ctx context.Context, request GetDeployArtifactRequest) (response GetDeployArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) getDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployArtifacts/{deployArtifactId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDeployArtifactResponse
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

// GetDeployEnvironment Retrieves a deployment environment by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployEnvironment.go.html to see an example of how to use GetDeployEnvironment API.
func (client DevopsClient) GetDeployEnvironment(ctx context.Context, request GetDeployEnvironmentRequest) (response GetDeployEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) getDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployEnvironments/{deployEnvironmentId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDeployEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployenvironment{})
	return response, err
}

// GetDeployPipeline Retrieves a deployment pipeline by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployPipeline.go.html to see an example of how to use GetDeployPipeline API.
func (client DevopsClient) GetDeployPipeline(ctx context.Context, request GetDeployPipelineRequest) (response GetDeployPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) getDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployPipelines/{deployPipelineId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDeployPipelineResponse
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

// GetDeployStage Retrieves a deployment stage by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployStage.go.html to see an example of how to use GetDeployStage API.
func (client DevopsClient) GetDeployStage(ctx context.Context, request GetDeployStageRequest) (response GetDeployStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) getDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployStages/{deployStageId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDeployStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deploystage{})
	return response, err
}

// GetDeployment Retrieves a deployment by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetDeployment.go.html to see an example of how to use GetDeployment API.
func (client DevopsClient) GetDeployment(ctx context.Context, request GetDeploymentRequest) (response GetDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) getDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// GetProject Retrieves a project by identifier.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetProject.go.html to see an example of how to use GetProject API.
func (client DevopsClient) GetProject(ctx context.Context, request GetProjectRequest) (response GetProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) getProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/projects/{projectId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response GetProjectResponse
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

// GetWorkRequest Retrieves the status of the work request with the given ID.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client DevopsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DevopsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody)
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

// ListDeployArtifacts Returns a list of deployment artifacts.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployArtifacts.go.html to see an example of how to use ListDeployArtifacts API.
func (client DevopsClient) ListDeployArtifacts(ctx context.Context, request ListDeployArtifactsRequest) (response ListDeployArtifactsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) listDeployArtifacts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployArtifacts", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDeployArtifactsResponse
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

// ListDeployEnvironments Returns a list of deployment environments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployEnvironments.go.html to see an example of how to use ListDeployEnvironments API.
func (client DevopsClient) ListDeployEnvironments(ctx context.Context, request ListDeployEnvironmentsRequest) (response ListDeployEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) listDeployEnvironments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployEnvironments", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDeployEnvironmentsResponse
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

// ListDeployPipelines Returns a list of deployment pipelines.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployPipelines.go.html to see an example of how to use ListDeployPipelines API.
func (client DevopsClient) ListDeployPipelines(ctx context.Context, request ListDeployPipelinesRequest) (response ListDeployPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) listDeployPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployPipelines", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDeployPipelinesResponse
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

// ListDeployStages Retrieves a list of deployment stages.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployStages.go.html to see an example of how to use ListDeployStages API.
func (client DevopsClient) ListDeployStages(ctx context.Context, request ListDeployStagesRequest) (response ListDeployStagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) listDeployStages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployStages", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDeployStagesResponse
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

// ListDeployments Returns a list of deployments.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListDeployments.go.html to see an example of how to use ListDeployments API.
func (client DevopsClient) ListDeployments(ctx context.Context, request ListDeploymentsRequest) (response ListDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) listDeployments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentsResponse
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

// ListProjects Returns a list of projects.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListProjects.go.html to see an example of how to use ListProjects API.
func (client DevopsClient) ListProjects(ctx context.Context, request ListProjectsRequest) (response ListProjectsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) listProjects(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/projects", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response ListProjectsResponse
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

// ListWorkRequestErrors Returns a list of errors for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client DevopsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DevopsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody)
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

// ListWorkRequestLogs Returns a list of logs for a given work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client DevopsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DevopsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody)
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

// ListWorkRequests Lists the work requests in a compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client DevopsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DevopsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody)
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

// UpdateDeployArtifact Updates the deployment artifact.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployArtifact.go.html to see an example of how to use UpdateDeployArtifact API.
func (client DevopsClient) UpdateDeployArtifact(ctx context.Context, request UpdateDeployArtifactRequest) (response UpdateDeployArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) updateDeployArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployArtifacts/{deployArtifactId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployArtifactResponse
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

// UpdateDeployEnvironment Updates the deployment environment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployEnvironment.go.html to see an example of how to use UpdateDeployEnvironment API.
func (client DevopsClient) UpdateDeployEnvironment(ctx context.Context, request UpdateDeployEnvironmentRequest) (response UpdateDeployEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) updateDeployEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployEnvironments/{deployEnvironmentId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployenvironment{})
	return response, err
}

// UpdateDeployPipeline Updates the deployment pipeline.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployPipeline.go.html to see an example of how to use UpdateDeployPipeline API.
func (client DevopsClient) UpdateDeployPipeline(ctx context.Context, request UpdateDeployPipelineRequest) (response UpdateDeployPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) updateDeployPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployPipelines/{deployPipelineId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployPipelineResponse
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

// UpdateDeployStage Updates the deployment stage.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployStage.go.html to see an example of how to use UpdateDeployStage API.
func (client DevopsClient) UpdateDeployStage(ctx context.Context, request UpdateDeployStageRequest) (response UpdateDeployStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) updateDeployStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployStages/{deployStageId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDeployStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deploystage{})
	return response, err
}

// UpdateDeployment Updates the deployment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateDeployment.go.html to see an example of how to use UpdateDeployment API.
func (client DevopsClient) UpdateDeployment(ctx context.Context, request UpdateDeploymentRequest) (response UpdateDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) updateDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployments/{deploymentId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &deployment{})
	return response, err
}

// UpdateProject Updates the project.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/UpdateProject.go.html to see an example of how to use UpdateProject API.
func (client DevopsClient) UpdateProject(ctx context.Context, request UpdateProjectRequest) (response UpdateProjectResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client DevopsClient) updateProject(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/projects/{projectId}", binaryReqBody)
	if err != nil {
		return nil, err
	}

	var response UpdateProjectResponse
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
