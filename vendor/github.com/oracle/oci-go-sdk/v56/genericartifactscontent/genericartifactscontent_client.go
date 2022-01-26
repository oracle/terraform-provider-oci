// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generic Artifacts Content API
//
// API covering the Generic Artifacts Service content
// Use this API to put and get generic artifact content.
//

package genericartifactscontent

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//GenericArtifactsContentClient a client for GenericArtifactsContent
type GenericArtifactsContentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGenericArtifactsContentClientWithConfigurationProvider Creates a new default GenericArtifactsContent client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGenericArtifactsContentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GenericArtifactsContentClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newGenericArtifactsContentClientFromBaseClient(baseClient, provider)
}

// NewGenericArtifactsContentClientWithOboToken Creates a new default GenericArtifactsContent client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewGenericArtifactsContentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GenericArtifactsContentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGenericArtifactsContentClientFromBaseClient(baseClient, configProvider)
}

func newGenericArtifactsContentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GenericArtifactsContentClient, err error) {
	// GenericArtifactsContent service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GenericArtifactsContentClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GenericArtifactsContentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("genericartifactscontent", "https://generic.artifacts.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GenericArtifactsContentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GenericArtifactsContentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetGenericArtifactContent Gets the specified artifact's content.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/genericartifactscontent/GetGenericArtifactContent.go.html to see an example of how to use GetGenericArtifactContent API.
func (client GenericArtifactsContentClient) GetGenericArtifactContent(ctx context.Context, request GetGenericArtifactContentRequest) (response GetGenericArtifactContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGenericArtifactContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGenericArtifactContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGenericArtifactContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGenericArtifactContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGenericArtifactContentResponse")
	}
	return
}

// getGenericArtifactContent implements the OCIOperation interface (enables retrying operations)
func (client GenericArtifactsContentClient) getGenericArtifactContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/generic/artifacts/{artifactId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGenericArtifactContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGenericArtifactContentByPath Gets the content of an artifact with a specified `artifactPath` and `version`.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/genericartifactscontent/GetGenericArtifactContentByPath.go.html to see an example of how to use GetGenericArtifactContentByPath API.
func (client GenericArtifactsContentClient) GetGenericArtifactContentByPath(ctx context.Context, request GetGenericArtifactContentByPathRequest) (response GetGenericArtifactContentByPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGenericArtifactContentByPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGenericArtifactContentByPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGenericArtifactContentByPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGenericArtifactContentByPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGenericArtifactContentByPathResponse")
	}
	return
}

// getGenericArtifactContentByPath implements the OCIOperation interface (enables retrying operations)
func (client GenericArtifactsContentClient) getGenericArtifactContentByPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/generic/repositories/{repositoryId}/artifactPaths/{artifactPath}/versions/{version}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGenericArtifactContentByPathResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutGenericArtifactContentByPath Uploads an artifact. Provide `artifactPath`, `version` and content. Avoid entering confidential information when you define the path and version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/genericartifactscontent/PutGenericArtifactContentByPath.go.html to see an example of how to use PutGenericArtifactContentByPath API.
func (client GenericArtifactsContentClient) PutGenericArtifactContentByPath(ctx context.Context, request PutGenericArtifactContentByPathRequest) (response PutGenericArtifactContentByPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putGenericArtifactContentByPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutGenericArtifactContentByPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutGenericArtifactContentByPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutGenericArtifactContentByPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutGenericArtifactContentByPathResponse")
	}
	return
}

// putGenericArtifactContentByPath implements the OCIOperation interface (enables retrying operations)
func (client GenericArtifactsContentClient) putGenericArtifactContentByPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/generic/repositories/{repositoryId}/artifactPaths/{artifactPath}/versions/{version}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutGenericArtifactContentByPathResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
