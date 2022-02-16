// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Images API
//
// API covering the Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as container images and repositories.
//

package artifacts

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//ArtifactsClient a client for Artifacts
type ArtifactsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewArtifactsClientWithConfigurationProvider Creates a new default Artifacts client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewArtifactsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ArtifactsClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newArtifactsClientFromBaseClient(baseClient, provider)
}

// NewArtifactsClientWithOboToken Creates a new default Artifacts client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewArtifactsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ArtifactsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newArtifactsClientFromBaseClient(baseClient, configProvider)
}

func newArtifactsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ArtifactsClient, err error) {
	// Artifacts service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ArtifactsClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ArtifactsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("artifacts", "https://artifacts.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ArtifactsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ArtifactsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeContainerRepositoryCompartment Moves a container repository into a different compartment within the same tenancy. For information about moving
// resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ChangeContainerRepositoryCompartment.go.html to see an example of how to use ChangeContainerRepositoryCompartment API.
func (client ArtifactsClient) ChangeContainerRepositoryCompartment(ctx context.Context, request ChangeContainerRepositoryCompartmentRequest) (response ChangeContainerRepositoryCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeContainerRepositoryCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeContainerRepositoryCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeContainerRepositoryCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeContainerRepositoryCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeContainerRepositoryCompartmentResponse")
	}
	return
}

// changeContainerRepositoryCompartment implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) changeContainerRepositoryCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/container/repositories/{repositoryId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeContainerRepositoryCompartmentResponse
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

// ChangeRepositoryCompartment Moves a repository into a different compartment within the same tenancy. For information about moving
// resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ChangeRepositoryCompartment.go.html to see an example of how to use ChangeRepositoryCompartment API.
func (client ArtifactsClient) ChangeRepositoryCompartment(ctx context.Context, request ChangeRepositoryCompartmentRequest) (response ChangeRepositoryCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeRepositoryCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRepositoryCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRepositoryCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRepositoryCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRepositoryCompartmentResponse")
	}
	return
}

// changeRepositoryCompartment implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) changeRepositoryCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/repositories/{repositoryId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRepositoryCompartmentResponse
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

// CreateContainerImageSignature Upload a signature to an image.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/CreateContainerImageSignature.go.html to see an example of how to use CreateContainerImageSignature API.
func (client ArtifactsClient) CreateContainerImageSignature(ctx context.Context, request CreateContainerImageSignatureRequest) (response CreateContainerImageSignatureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createContainerImageSignature, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateContainerImageSignatureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateContainerImageSignatureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateContainerImageSignatureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateContainerImageSignatureResponse")
	}
	return
}

// createContainerImageSignature implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) createContainerImageSignature(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/container/imageSignatures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateContainerImageSignatureResponse
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

// CreateContainerRepository Create a new empty container repository. Avoid entering confidential information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/CreateContainerRepository.go.html to see an example of how to use CreateContainerRepository API.
func (client ArtifactsClient) CreateContainerRepository(ctx context.Context, request CreateContainerRepositoryRequest) (response CreateContainerRepositoryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createContainerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateContainerRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateContainerRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateContainerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateContainerRepositoryResponse")
	}
	return
}

// createContainerRepository implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) createContainerRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/container/repositories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateContainerRepositoryResponse
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

// CreateRepository Creates a new repository for storing artifacts.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/CreateRepository.go.html to see an example of how to use CreateRepository API.
func (client ArtifactsClient) CreateRepository(ctx context.Context, request CreateRepositoryRequest) (response CreateRepositoryResponse, err error) {
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
func (client ArtifactsClient) createRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &repository{})
	return response, err
}

// DeleteContainerImage Delete a container image.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/DeleteContainerImage.go.html to see an example of how to use DeleteContainerImage API.
func (client ArtifactsClient) DeleteContainerImage(ctx context.Context, request DeleteContainerImageRequest) (response DeleteContainerImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteContainerImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteContainerImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteContainerImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteContainerImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteContainerImageResponse")
	}
	return
}

// deleteContainerImage implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) deleteContainerImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/container/images/{imageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteContainerImageResponse
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

// DeleteContainerImageSignature Delete a container image signature.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/DeleteContainerImageSignature.go.html to see an example of how to use DeleteContainerImageSignature API.
func (client ArtifactsClient) DeleteContainerImageSignature(ctx context.Context, request DeleteContainerImageSignatureRequest) (response DeleteContainerImageSignatureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteContainerImageSignature, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteContainerImageSignatureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteContainerImageSignatureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteContainerImageSignatureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteContainerImageSignatureResponse")
	}
	return
}

// deleteContainerImageSignature implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) deleteContainerImageSignature(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/container/imageSignatures/{imageSignatureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteContainerImageSignatureResponse
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

// DeleteContainerRepository Delete container repository.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/DeleteContainerRepository.go.html to see an example of how to use DeleteContainerRepository API.
func (client ArtifactsClient) DeleteContainerRepository(ctx context.Context, request DeleteContainerRepositoryRequest) (response DeleteContainerRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteContainerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteContainerRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteContainerRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteContainerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteContainerRepositoryResponse")
	}
	return
}

// deleteContainerRepository implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) deleteContainerRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/container/repositories/{repositoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteContainerRepositoryResponse
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

// DeleteGenericArtifact Deletes an artifact with a specified OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/DeleteGenericArtifact.go.html to see an example of how to use DeleteGenericArtifact API.
func (client ArtifactsClient) DeleteGenericArtifact(ctx context.Context, request DeleteGenericArtifactRequest) (response DeleteGenericArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGenericArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGenericArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGenericArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGenericArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGenericArtifactResponse")
	}
	return
}

// deleteGenericArtifact implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) deleteGenericArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/generic/artifacts/{artifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGenericArtifactResponse
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

// DeleteGenericArtifactByPath Deletes an artifact with a specified `artifactPath` and `version`.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/DeleteGenericArtifactByPath.go.html to see an example of how to use DeleteGenericArtifactByPath API.
func (client ArtifactsClient) DeleteGenericArtifactByPath(ctx context.Context, request DeleteGenericArtifactByPathRequest) (response DeleteGenericArtifactByPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGenericArtifactByPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGenericArtifactByPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGenericArtifactByPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGenericArtifactByPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGenericArtifactByPathResponse")
	}
	return
}

// deleteGenericArtifactByPath implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) deleteGenericArtifactByPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/generic/repositories/{repositoryId}/artifactPaths/{artifactPath}/versions/{version}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGenericArtifactByPathResponse
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

// DeleteRepository Deletes the specified repository. This operation fails unless all associated artifacts are in a DELETED state. You must delete all associated artifacts before deleting a repository.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/DeleteRepository.go.html to see an example of how to use DeleteRepository API.
func (client ArtifactsClient) DeleteRepository(ctx context.Context, request DeleteRepositoryRequest) (response DeleteRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ArtifactsClient) deleteRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetContainerConfiguration Get container configuration.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetContainerConfiguration.go.html to see an example of how to use GetContainerConfiguration API.
func (client ArtifactsClient) GetContainerConfiguration(ctx context.Context, request GetContainerConfigurationRequest) (response GetContainerConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getContainerConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetContainerConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetContainerConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetContainerConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetContainerConfigurationResponse")
	}
	return
}

// getContainerConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) getContainerConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/container/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetContainerConfigurationResponse
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

// GetContainerImage Get container image metadata.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetContainerImage.go.html to see an example of how to use GetContainerImage API.
func (client ArtifactsClient) GetContainerImage(ctx context.Context, request GetContainerImageRequest) (response GetContainerImageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getContainerImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetContainerImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetContainerImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetContainerImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetContainerImageResponse")
	}
	return
}

// getContainerImage implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) getContainerImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/container/images/{imageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetContainerImageResponse
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

// GetContainerImageSignature Get container image signature metadata.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetContainerImageSignature.go.html to see an example of how to use GetContainerImageSignature API.
func (client ArtifactsClient) GetContainerImageSignature(ctx context.Context, request GetContainerImageSignatureRequest) (response GetContainerImageSignatureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getContainerImageSignature, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetContainerImageSignatureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetContainerImageSignatureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetContainerImageSignatureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetContainerImageSignatureResponse")
	}
	return
}

// getContainerImageSignature implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) getContainerImageSignature(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/container/imageSignatures/{imageSignatureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetContainerImageSignatureResponse
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

// GetContainerRepository Get container repository.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetContainerRepository.go.html to see an example of how to use GetContainerRepository API.
func (client ArtifactsClient) GetContainerRepository(ctx context.Context, request GetContainerRepositoryRequest) (response GetContainerRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getContainerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetContainerRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetContainerRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetContainerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetContainerRepositoryResponse")
	}
	return
}

// getContainerRepository implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) getContainerRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/container/repositories/{repositoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetContainerRepositoryResponse
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

// GetGenericArtifact Gets information about an artifact with a specified OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetGenericArtifact.go.html to see an example of how to use GetGenericArtifact API.
func (client ArtifactsClient) GetGenericArtifact(ctx context.Context, request GetGenericArtifactRequest) (response GetGenericArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGenericArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGenericArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGenericArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGenericArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGenericArtifactResponse")
	}
	return
}

// getGenericArtifact implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) getGenericArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/generic/artifacts/{artifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGenericArtifactResponse
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

// GetGenericArtifactByPath Gets information about an artifact with a specified `artifactPath` and `version`.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetGenericArtifactByPath.go.html to see an example of how to use GetGenericArtifactByPath API.
func (client ArtifactsClient) GetGenericArtifactByPath(ctx context.Context, request GetGenericArtifactByPathRequest) (response GetGenericArtifactByPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGenericArtifactByPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGenericArtifactByPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGenericArtifactByPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGenericArtifactByPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGenericArtifactByPathResponse")
	}
	return
}

// getGenericArtifactByPath implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) getGenericArtifactByPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/generic/repositories/{repositoryId}/artifactPaths/{artifactPath}/versions/{version}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGenericArtifactByPathResponse
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

// GetRepository Gets the specified repository's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/GetRepository.go.html to see an example of how to use GetRepository API.
func (client ArtifactsClient) GetRepository(ctx context.Context, request GetRepositoryRequest) (response GetRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ArtifactsClient) getRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &repository{})
	return response, err
}

// ListContainerImageSignatures List container image signatures in an image.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListContainerImageSignatures.go.html to see an example of how to use ListContainerImageSignatures API.
func (client ArtifactsClient) ListContainerImageSignatures(ctx context.Context, request ListContainerImageSignaturesRequest) (response ListContainerImageSignaturesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listContainerImageSignatures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListContainerImageSignaturesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListContainerImageSignaturesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListContainerImageSignaturesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListContainerImageSignaturesResponse")
	}
	return
}

// listContainerImageSignatures implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) listContainerImageSignatures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/container/imageSignatures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListContainerImageSignaturesResponse
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

// ListContainerImages List container images in a compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListContainerImages.go.html to see an example of how to use ListContainerImages API.
func (client ArtifactsClient) ListContainerImages(ctx context.Context, request ListContainerImagesRequest) (response ListContainerImagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listContainerImages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListContainerImagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListContainerImagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListContainerImagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListContainerImagesResponse")
	}
	return
}

// listContainerImages implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) listContainerImages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/container/images", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListContainerImagesResponse
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

// ListContainerRepositories List container repositories in a compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListContainerRepositories.go.html to see an example of how to use ListContainerRepositories API.
func (client ArtifactsClient) ListContainerRepositories(ctx context.Context, request ListContainerRepositoriesRequest) (response ListContainerRepositoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listContainerRepositories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListContainerRepositoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListContainerRepositoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListContainerRepositoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListContainerRepositoriesResponse")
	}
	return
}

// listContainerRepositories implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) listContainerRepositories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/container/repositories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListContainerRepositoriesResponse
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

// ListGenericArtifacts Lists artifacts in the specified repository.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListGenericArtifacts.go.html to see an example of how to use ListGenericArtifacts API.
func (client ArtifactsClient) ListGenericArtifacts(ctx context.Context, request ListGenericArtifactsRequest) (response ListGenericArtifactsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGenericArtifacts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGenericArtifactsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGenericArtifactsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGenericArtifactsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGenericArtifactsResponse")
	}
	return
}

// listGenericArtifacts implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) listGenericArtifacts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/generic/artifacts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGenericArtifactsResponse
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

// ListRepositories Lists repositories in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/ListRepositories.go.html to see an example of how to use ListRepositories API.
func (client ArtifactsClient) ListRepositories(ctx context.Context, request ListRepositoriesRequest) (response ListRepositoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ArtifactsClient) listRepositories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveContainerVersion Remove version from container image.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/RemoveContainerVersion.go.html to see an example of how to use RemoveContainerVersion API.
func (client ArtifactsClient) RemoveContainerVersion(ctx context.Context, request RemoveContainerVersionRequest) (response RemoveContainerVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeContainerVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveContainerVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveContainerVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveContainerVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveContainerVersionResponse")
	}
	return
}

// removeContainerVersion implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) removeContainerVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/container/images/{imageId}/actions/removeVersion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveContainerVersionResponse
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

// RestoreContainerImage Restore a container image.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/RestoreContainerImage.go.html to see an example of how to use RestoreContainerImage API.
func (client ArtifactsClient) RestoreContainerImage(ctx context.Context, request RestoreContainerImageRequest) (response RestoreContainerImageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restoreContainerImage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreContainerImageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreContainerImageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreContainerImageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreContainerImageResponse")
	}
	return
}

// restoreContainerImage implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) restoreContainerImage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/container/images/{imageId}/actions/restore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestoreContainerImageResponse
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

// UpdateContainerConfiguration Update container configuration.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/UpdateContainerConfiguration.go.html to see an example of how to use UpdateContainerConfiguration API.
func (client ArtifactsClient) UpdateContainerConfiguration(ctx context.Context, request UpdateContainerConfigurationRequest) (response UpdateContainerConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateContainerConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateContainerConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateContainerConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateContainerConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateContainerConfigurationResponse")
	}
	return
}

// updateContainerConfiguration implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) updateContainerConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/container/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateContainerConfigurationResponse
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

// UpdateContainerRepository Modify the properties of a container repository. Avoid entering confidential information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/UpdateContainerRepository.go.html to see an example of how to use UpdateContainerRepository API.
func (client ArtifactsClient) UpdateContainerRepository(ctx context.Context, request UpdateContainerRepositoryRequest) (response UpdateContainerRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateContainerRepository, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateContainerRepositoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateContainerRepositoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateContainerRepositoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateContainerRepositoryResponse")
	}
	return
}

// updateContainerRepository implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) updateContainerRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/container/repositories/{repositoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateContainerRepositoryResponse
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

// UpdateGenericArtifact Updates the artifact with the specified OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). You can only update the tags of an artifact.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/UpdateGenericArtifact.go.html to see an example of how to use UpdateGenericArtifact API.
func (client ArtifactsClient) UpdateGenericArtifact(ctx context.Context, request UpdateGenericArtifactRequest) (response UpdateGenericArtifactResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGenericArtifact, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGenericArtifactResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGenericArtifactResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGenericArtifactResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGenericArtifactResponse")
	}
	return
}

// updateGenericArtifact implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) updateGenericArtifact(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/generic/artifacts/{artifactId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGenericArtifactResponse
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

// UpdateGenericArtifactByPath Updates an artifact with a specified `artifactPath` and `version`. You can only update the tags of an artifact.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/UpdateGenericArtifactByPath.go.html to see an example of how to use UpdateGenericArtifactByPath API.
func (client ArtifactsClient) UpdateGenericArtifactByPath(ctx context.Context, request UpdateGenericArtifactByPathRequest) (response UpdateGenericArtifactByPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGenericArtifactByPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGenericArtifactByPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGenericArtifactByPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGenericArtifactByPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGenericArtifactByPathResponse")
	}
	return
}

// updateGenericArtifactByPath implements the OCIOperation interface (enables retrying operations)
func (client ArtifactsClient) updateGenericArtifactByPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/generic/repositories/{repositoryId}/artifactPaths/{artifactPath}/versions/{version}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGenericArtifactByPathResponse
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

// UpdateRepository Updates the properties of a repository. You can update the `displayName` and  `description` properties.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/artifacts/UpdateRepository.go.html to see an example of how to use UpdateRepository API.
func (client ArtifactsClient) UpdateRepository(ctx context.Context, request UpdateRepositoryRequest) (response UpdateRepositoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client ArtifactsClient) updateRepository(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &repository{})
	return response, err
}
