// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OpensearchClusterPipelineClient a client for OpensearchClusterPipeline
type OpensearchClusterPipelineClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOpensearchClusterPipelineClientWithConfigurationProvider Creates a new default OpensearchClusterPipeline client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOpensearchClusterPipelineClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OpensearchClusterPipelineClient, err error) {
	if enabled := common.CheckForEnabledServices("opensearch"); !enabled {
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
	return newOpensearchClusterPipelineClientFromBaseClient(baseClient, provider)
}

// NewOpensearchClusterPipelineClientWithOboToken Creates a new default OpensearchClusterPipeline client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOpensearchClusterPipelineClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OpensearchClusterPipelineClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOpensearchClusterPipelineClientFromBaseClient(baseClient, configProvider)
}

func newOpensearchClusterPipelineClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OpensearchClusterPipelineClient, err error) {
	// OpensearchClusterPipeline service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OpensearchClusterPipeline"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OpensearchClusterPipelineClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OpensearchClusterPipelineClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("opensearch", "https://search-indexing.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OpensearchClusterPipelineClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OpensearchClusterPipelineClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateOpensearchClusterPipeline Creates a new OpensearchCluster Pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/CreateOpensearchClusterPipeline.go.html to see an example of how to use CreateOpensearchClusterPipeline API.
func (client OpensearchClusterPipelineClient) CreateOpensearchClusterPipeline(ctx context.Context, request CreateOpensearchClusterPipelineRequest) (response CreateOpensearchClusterPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOpensearchClusterPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOpensearchClusterPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOpensearchClusterPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOpensearchClusterPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOpensearchClusterPipelineResponse")
	}
	return
}

// createOpensearchClusterPipeline implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterPipelineClient) createOpensearchClusterPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/opensearchClusterPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOpensearchClusterPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterPipeline/CreateOpensearchClusterPipeline"
		err = common.PostProcessServiceError(err, "OpensearchClusterPipeline", "CreateOpensearchClusterPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOpensearchClusterPipeline Deletes a OpensearchCluster Pipeline resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/DeleteOpensearchClusterPipeline.go.html to see an example of how to use DeleteOpensearchClusterPipeline API.
func (client OpensearchClusterPipelineClient) DeleteOpensearchClusterPipeline(ctx context.Context, request DeleteOpensearchClusterPipelineRequest) (response DeleteOpensearchClusterPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOpensearchClusterPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOpensearchClusterPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOpensearchClusterPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOpensearchClusterPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOpensearchClusterPipelineResponse")
	}
	return
}

// deleteOpensearchClusterPipeline implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterPipelineClient) deleteOpensearchClusterPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/opensearchClusterPipelines/{opensearchClusterPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOpensearchClusterPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterPipeline/DeleteOpensearchClusterPipeline"
		err = common.PostProcessServiceError(err, "OpensearchClusterPipeline", "DeleteOpensearchClusterPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOpensearchClusterPipeline Gets a OpensearchCluster Pipeline by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/GetOpensearchClusterPipeline.go.html to see an example of how to use GetOpensearchClusterPipeline API.
func (client OpensearchClusterPipelineClient) GetOpensearchClusterPipeline(ctx context.Context, request GetOpensearchClusterPipelineRequest) (response GetOpensearchClusterPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOpensearchClusterPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOpensearchClusterPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOpensearchClusterPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOpensearchClusterPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOpensearchClusterPipelineResponse")
	}
	return
}

// getOpensearchClusterPipeline implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterPipelineClient) getOpensearchClusterPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opensearchClusterPipelines/{opensearchClusterPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOpensearchClusterPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterPipeline/GetOpensearchClusterPipeline"
		err = common.PostProcessServiceError(err, "OpensearchClusterPipeline", "GetOpensearchClusterPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOpensearchClusterPipelines Returns a list of OpensearchClusterPipelines.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/ListOpensearchClusterPipelines.go.html to see an example of how to use ListOpensearchClusterPipelines API.
func (client OpensearchClusterPipelineClient) ListOpensearchClusterPipelines(ctx context.Context, request ListOpensearchClusterPipelinesRequest) (response ListOpensearchClusterPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOpensearchClusterPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOpensearchClusterPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOpensearchClusterPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOpensearchClusterPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOpensearchClusterPipelinesResponse")
	}
	return
}

// listOpensearchClusterPipelines implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterPipelineClient) listOpensearchClusterPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/opensearchClusterPipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOpensearchClusterPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterPipelineCollection/ListOpensearchClusterPipelines"
		err = common.PostProcessServiceError(err, "OpensearchClusterPipeline", "ListOpensearchClusterPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOpensearchClusterPipeline Updates the OpensearchCluster Pipeline
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opensearch/UpdateOpensearchClusterPipeline.go.html to see an example of how to use UpdateOpensearchClusterPipeline API.
func (client OpensearchClusterPipelineClient) UpdateOpensearchClusterPipeline(ctx context.Context, request UpdateOpensearchClusterPipelineRequest) (response UpdateOpensearchClusterPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOpensearchClusterPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOpensearchClusterPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOpensearchClusterPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOpensearchClusterPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOpensearchClusterPipelineResponse")
	}
	return
}

// updateOpensearchClusterPipeline implements the OCIOperation interface (enables retrying operations)
func (client OpensearchClusterPipelineClient) updateOpensearchClusterPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/opensearchClusterPipelines/{opensearchClusterPipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOpensearchClusterPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/opensearch/20180828/OpensearchClusterPipeline/UpdateOpensearchClusterPipeline"
		err = common.PostProcessServiceError(err, "OpensearchClusterPipeline", "UpdateOpensearchClusterPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
