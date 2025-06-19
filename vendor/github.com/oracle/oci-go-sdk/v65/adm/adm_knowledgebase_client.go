// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// KnowledgeBaseClient a client for KnowledgeBase
type KnowledgeBaseClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewKnowledgeBaseClientWithConfigurationProvider Creates a new default KnowledgeBase client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewKnowledgeBaseClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client KnowledgeBaseClient, err error) {
	if enabled := common.CheckForEnabledServices("adm"); !enabled {
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
	return newKnowledgeBaseClientFromBaseClient(baseClient, provider)
}

// NewKnowledgeBaseClientWithOboToken Creates a new default KnowledgeBase client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewKnowledgeBaseClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client KnowledgeBaseClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newKnowledgeBaseClientFromBaseClient(baseClient, configProvider)
}

func newKnowledgeBaseClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client KnowledgeBaseClient, err error) {
	// KnowledgeBase service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("KnowledgeBase"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = KnowledgeBaseClient{BaseClient: baseClient}
	client.BasePath = "20220421"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *KnowledgeBaseClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("adm", "https://adm.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *KnowledgeBaseClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *KnowledgeBaseClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetArtifactVersion Gets an Artifact Version by identifier.
// A default retry strategy applies to this operation GetArtifactVersion()
func (client KnowledgeBaseClient) GetArtifactVersion(ctx context.Context, request GetArtifactVersionRequest) (response GetArtifactVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getArtifactVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetArtifactVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetArtifactVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetArtifactVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetArtifactVersionResponse")
	}
	return
}

// getArtifactVersion implements the OCIOperation interface (enables retrying operations)
func (client KnowledgeBaseClient) getArtifactVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/knowledgeBases/{knowledgeBaseId}/artifactVersions/{artifactVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetArtifactVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/ArtifactVersion/GetArtifactVersion"
		err = common.PostProcessServiceError(err, "KnowledgeBase", "GetArtifactVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListArtifactVersions Returns a list of Artifact Versions in the knowledge base.
// A default retry strategy applies to this operation ListArtifactVersions()
func (client KnowledgeBaseClient) ListArtifactVersions(ctx context.Context, request ListArtifactVersionsRequest) (response ListArtifactVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listArtifactVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListArtifactVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListArtifactVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListArtifactVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListArtifactVersionsResponse")
	}
	return
}

// listArtifactVersions implements the OCIOperation interface (enables retrying operations)
func (client KnowledgeBaseClient) listArtifactVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/knowledgeBases/{knowledgeBaseId}/artifactVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListArtifactVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/ArtifactVersionSummaryCollection/ListArtifactVersions"
		err = common.PostProcessServiceError(err, "KnowledgeBase", "ListArtifactVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCopyrights Gets the Copyrights of an Artifact Version by identifier.
// A default retry strategy applies to this operation ListCopyrights()
func (client KnowledgeBaseClient) ListCopyrights(ctx context.Context, request ListCopyrightsRequest) (response ListCopyrightsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCopyrights, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCopyrightsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCopyrightsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCopyrightsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCopyrightsResponse")
	}
	return
}

// listCopyrights implements the OCIOperation interface (enables retrying operations)
func (client KnowledgeBaseClient) listCopyrights(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/knowledgeBases/{knowledgeBaseId}/artifactVersions/{artifactVersionId}/copyrights", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCopyrightsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/CopyrightCollection/ListCopyrights"
		err = common.PostProcessServiceError(err, "KnowledgeBase", "ListCopyrights", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNoticeFiles Gets the notices of an Artifact Version by identifier.
// A default retry strategy applies to this operation ListNoticeFiles()
func (client KnowledgeBaseClient) ListNoticeFiles(ctx context.Context, request ListNoticeFilesRequest) (response ListNoticeFilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNoticeFiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNoticeFilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNoticeFilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNoticeFilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNoticeFilesResponse")
	}
	return
}

// listNoticeFiles implements the OCIOperation interface (enables retrying operations)
func (client KnowledgeBaseClient) listNoticeFiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/knowledgeBases/{knowledgeBaseId}/artifactVersions/{artifactVersionId}/notices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNoticeFilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/NoticeFileCollection/ListNoticeFiles"
		err = common.PostProcessServiceError(err, "KnowledgeBase", "ListNoticeFiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
