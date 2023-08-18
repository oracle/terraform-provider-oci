// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ContainerEngineClient a client for ContainerEngine
type ContainerEngineClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewContainerEngineClientWithConfigurationProvider Creates a new default ContainerEngine client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewContainerEngineClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ContainerEngineClient, err error) {
	if enabled := common.CheckForEnabledServices("containerengine"); !enabled {
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
	return newContainerEngineClientFromBaseClient(baseClient, provider)
}

// NewContainerEngineClientWithOboToken Creates a new default ContainerEngine client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewContainerEngineClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ContainerEngineClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newContainerEngineClientFromBaseClient(baseClient, configProvider)
}

func newContainerEngineClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ContainerEngineClient, err error) {
	// ContainerEngine service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ContainerEngine"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ContainerEngineClient{BaseClient: baseClient}
	client.BasePath = "20180222"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ContainerEngineClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("containerengine", "https://containerengine.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ContainerEngineClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ContainerEngineClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeClusterAttachmentCompartment Moves a ClusterAttachment resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
// A default retry strategy applies to this operation ChangeClusterAttachmentCompartment()
func (client ContainerEngineClient) ChangeClusterAttachmentCompartment(ctx context.Context, request ChangeClusterAttachmentCompartmentRequest) (response ChangeClusterAttachmentCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeClusterAttachmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeClusterAttachmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeClusterAttachmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeClusterAttachmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeClusterAttachmentCompartmentResponse")
	}
	return
}

// changeClusterAttachmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) changeClusterAttachmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterAttachments/{clusterAttachmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeClusterAttachmentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterAttachment/ChangeClusterAttachmentCompartment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ChangeClusterAttachmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeClusterNamespaceCompartment Moves a ClusterNamespace resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
// A default retry strategy applies to this operation ChangeClusterNamespaceCompartment()
func (client ContainerEngineClient) ChangeClusterNamespaceCompartment(ctx context.Context, request ChangeClusterNamespaceCompartmentRequest) (response ChangeClusterNamespaceCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeClusterNamespaceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeClusterNamespaceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeClusterNamespaceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeClusterNamespaceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeClusterNamespaceCompartmentResponse")
	}
	return
}

// changeClusterNamespaceCompartment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) changeClusterNamespaceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNamespaces/{clusterNamespaceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeClusterNamespaceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespace/ChangeClusterNamespaceCompartment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ChangeClusterNamespaceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeClusterNamespaceProfileCompartment Moves a ClusterNamespaceProfile resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
// A default retry strategy applies to this operation ChangeClusterNamespaceProfileCompartment()
func (client ContainerEngineClient) ChangeClusterNamespaceProfileCompartment(ctx context.Context, request ChangeClusterNamespaceProfileCompartmentRequest) (response ChangeClusterNamespaceProfileCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeClusterNamespaceProfileCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeClusterNamespaceProfileCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeClusterNamespaceProfileCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeClusterNamespaceProfileCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeClusterNamespaceProfileCompartmentResponse")
	}
	return
}

// changeClusterNamespaceProfileCompartment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) changeClusterNamespaceProfileCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNamespaceProfiles/{clusterNamespaceProfileId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeClusterNamespaceProfileCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfile/ChangeClusterNamespaceProfileCompartment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ChangeClusterNamespaceProfileCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeClusterNamespaceProfileVersionCompartment Moves a ClusterNamespaceProfileVersion resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
// A default retry strategy applies to this operation ChangeClusterNamespaceProfileVersionCompartment()
func (client ContainerEngineClient) ChangeClusterNamespaceProfileVersionCompartment(ctx context.Context, request ChangeClusterNamespaceProfileVersionCompartmentRequest) (response ChangeClusterNamespaceProfileVersionCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeClusterNamespaceProfileVersionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeClusterNamespaceProfileVersionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeClusterNamespaceProfileVersionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeClusterNamespaceProfileVersionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeClusterNamespaceProfileVersionCompartmentResponse")
	}
	return
}

// changeClusterNamespaceProfileVersionCompartment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) changeClusterNamespaceProfileVersionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNamespaceProfileVersions/{clusterNamespaceProfileVersionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeClusterNamespaceProfileVersionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfileVersion/ChangeClusterNamespaceProfileVersionCompartment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ChangeClusterNamespaceProfileVersionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ClusterMigrateToNativeVcn Initiates cluster migration to use native VCN.
// A default retry strategy applies to this operation ClusterMigrateToNativeVcn()
func (client ContainerEngineClient) ClusterMigrateToNativeVcn(ctx context.Context, request ClusterMigrateToNativeVcnRequest) (response ClusterMigrateToNativeVcnResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.clusterMigrateToNativeVcn, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ClusterMigrateToNativeVcnResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ClusterMigrateToNativeVcnResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ClusterMigrateToNativeVcnResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ClusterMigrateToNativeVcnResponse")
	}
	return
}

// clusterMigrateToNativeVcn implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) clusterMigrateToNativeVcn(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/actions/migrateToNativeVcn", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ClusterMigrateToNativeVcnResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/ClusterMigrateToNativeVcn"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ClusterMigrateToNativeVcn", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CompleteCredentialRotation Complete cluster credential rotation. Retire old credentials from kubernetes components.
// A default retry strategy applies to this operation CompleteCredentialRotation()
func (client ContainerEngineClient) CompleteCredentialRotation(ctx context.Context, request CompleteCredentialRotationRequest) (response CompleteCredentialRotationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.completeCredentialRotation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CompleteCredentialRotationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CompleteCredentialRotationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CompleteCredentialRotationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CompleteCredentialRotationResponse")
	}
	return
}

// completeCredentialRotation implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) completeCredentialRotation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/actions/completeCredentialRotation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CompleteCredentialRotationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/CompleteCredentialRotation"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CompleteCredentialRotation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCluster Create a new cluster.
// A default retry strategy applies to this operation CreateCluster()
func (client ContainerEngineClient) CreateCluster(ctx context.Context, request CreateClusterRequest) (response CreateClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateClusterResponse")
	}
	return
}

// createCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/CreateCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateClusterAttachment Creates a new ClusterAttachment.
// A default retry strategy applies to this operation CreateClusterAttachment()
func (client ContainerEngineClient) CreateClusterAttachment(ctx context.Context, request CreateClusterAttachmentRequest) (response CreateClusterAttachmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createClusterAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateClusterAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateClusterAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateClusterAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateClusterAttachmentResponse")
	}
	return
}

// createClusterAttachment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createClusterAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateClusterAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterAttachment/CreateClusterAttachment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateClusterAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateClusterNamespace Creates a new ClusterNamespace.
// A default retry strategy applies to this operation CreateClusterNamespace()
func (client ContainerEngineClient) CreateClusterNamespace(ctx context.Context, request CreateClusterNamespaceRequest) (response CreateClusterNamespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createClusterNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateClusterNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateClusterNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateClusterNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateClusterNamespaceResponse")
	}
	return
}

// createClusterNamespace implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createClusterNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNamespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateClusterNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespace/CreateClusterNamespace"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateClusterNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateClusterNamespaceProfile Creates a new ClusterNamespaceProfile.
// A default retry strategy applies to this operation CreateClusterNamespaceProfile()
func (client ContainerEngineClient) CreateClusterNamespaceProfile(ctx context.Context, request CreateClusterNamespaceProfileRequest) (response CreateClusterNamespaceProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createClusterNamespaceProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateClusterNamespaceProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateClusterNamespaceProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateClusterNamespaceProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateClusterNamespaceProfileResponse")
	}
	return
}

// createClusterNamespaceProfile implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createClusterNamespaceProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNamespaceProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateClusterNamespaceProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfile/CreateClusterNamespaceProfile"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateClusterNamespaceProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateClusterNamespaceProfileVersion Creates a new ClusterNamespaceProfileVersion.
// A default retry strategy applies to this operation CreateClusterNamespaceProfileVersion()
func (client ContainerEngineClient) CreateClusterNamespaceProfileVersion(ctx context.Context, request CreateClusterNamespaceProfileVersionRequest) (response CreateClusterNamespaceProfileVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createClusterNamespaceProfileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateClusterNamespaceProfileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateClusterNamespaceProfileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateClusterNamespaceProfileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateClusterNamespaceProfileVersionResponse")
	}
	return
}

// createClusterNamespaceProfileVersion implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createClusterNamespaceProfileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusterNamespaceProfileVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateClusterNamespaceProfileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfileVersion/CreateClusterNamespaceProfileVersion"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateClusterNamespaceProfileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateKubeconfig Create the Kubeconfig YAML for a cluster.
// A default retry strategy applies to this operation CreateKubeconfig()
func (client ContainerEngineClient) CreateKubeconfig(ctx context.Context, request CreateKubeconfigRequest) (response CreateKubeconfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createKubeconfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateKubeconfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateKubeconfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateKubeconfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateKubeconfigResponse")
	}
	return
}

// createKubeconfig implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createKubeconfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/kubeconfig/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateKubeconfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/CreateKubeconfig"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateKubeconfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateNodePool Create a new node pool.
// A default retry strategy applies to this operation CreateNodePool()
func (client ContainerEngineClient) CreateNodePool(ctx context.Context, request CreateNodePoolRequest) (response CreateNodePoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNodePoolResponse")
	}
	return
}

// createNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/nodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/CreateNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVirtualNodePool Create a new virtual node pool.
// A default retry strategy applies to this operation CreateVirtualNodePool()
func (client ContainerEngineClient) CreateVirtualNodePool(ctx context.Context, request CreateVirtualNodePoolRequest) (response CreateVirtualNodePoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVirtualNodePoolResponse")
	}
	return
}

// createVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualNodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/CreateVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateWorkloadMapping Create the specified workloadMapping for a cluster.
// A default retry strategy applies to this operation CreateWorkloadMapping()
func (client ContainerEngineClient) CreateWorkloadMapping(ctx context.Context, request CreateWorkloadMappingRequest) (response CreateWorkloadMappingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createWorkloadMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateWorkloadMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateWorkloadMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateWorkloadMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateWorkloadMappingResponse")
	}
	return
}

// createWorkloadMapping implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) createWorkloadMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/workloadMappings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateWorkloadMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkloadMapping/CreateWorkloadMapping"
		err = common.PostProcessServiceError(err, "ContainerEngine", "CreateWorkloadMapping", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCluster Delete a cluster.
// A default retry strategy applies to this operation DeleteCluster()
func (client ContainerEngineClient) DeleteCluster(ctx context.Context, request DeleteClusterRequest) (response DeleteClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteClusterResponse")
	}
	return
}

// deleteCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusters/{clusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/DeleteCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteClusterAttachment Deletes a ClusterAttachment resource by identifier
// A default retry strategy applies to this operation DeleteClusterAttachment()
func (client ContainerEngineClient) DeleteClusterAttachment(ctx context.Context, request DeleteClusterAttachmentRequest) (response DeleteClusterAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteClusterAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteClusterAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteClusterAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteClusterAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteClusterAttachmentResponse")
	}
	return
}

// deleteClusterAttachment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteClusterAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusterAttachments/{clusterAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteClusterAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterAttachment/DeleteClusterAttachment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteClusterAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteClusterNamespace Deletes a ClusterNamespace resource by identifier
// A default retry strategy applies to this operation DeleteClusterNamespace()
func (client ContainerEngineClient) DeleteClusterNamespace(ctx context.Context, request DeleteClusterNamespaceRequest) (response DeleteClusterNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteClusterNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteClusterNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteClusterNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteClusterNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteClusterNamespaceResponse")
	}
	return
}

// deleteClusterNamespace implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteClusterNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusterNamespaces/{clusterNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteClusterNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespace/DeleteClusterNamespace"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteClusterNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteClusterNamespaceProfile Deletes a ClusterNamespaceProfile resource by identifier
// A default retry strategy applies to this operation DeleteClusterNamespaceProfile()
func (client ContainerEngineClient) DeleteClusterNamespaceProfile(ctx context.Context, request DeleteClusterNamespaceProfileRequest) (response DeleteClusterNamespaceProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteClusterNamespaceProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteClusterNamespaceProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteClusterNamespaceProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteClusterNamespaceProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteClusterNamespaceProfileResponse")
	}
	return
}

// deleteClusterNamespaceProfile implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteClusterNamespaceProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusterNamespaceProfiles/{clusterNamespaceProfileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteClusterNamespaceProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfile/DeleteClusterNamespaceProfile"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteClusterNamespaceProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteClusterNamespaceProfileVersion Deletes a ClusterNamespaceProfileVersion resource by identifier
// A default retry strategy applies to this operation DeleteClusterNamespaceProfileVersion()
func (client ContainerEngineClient) DeleteClusterNamespaceProfileVersion(ctx context.Context, request DeleteClusterNamespaceProfileVersionRequest) (response DeleteClusterNamespaceProfileVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteClusterNamespaceProfileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteClusterNamespaceProfileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteClusterNamespaceProfileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteClusterNamespaceProfileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteClusterNamespaceProfileVersionResponse")
	}
	return
}

// deleteClusterNamespaceProfileVersion implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteClusterNamespaceProfileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusterNamespaceProfileVersions/{clusterNamespaceProfileVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteClusterNamespaceProfileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfileVersion/DeleteClusterNamespaceProfileVersion"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteClusterNamespaceProfileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNode Delete node.
// A default retry strategy applies to this operation DeleteNode()
func (client ContainerEngineClient) DeleteNode(ctx context.Context, request DeleteNodeRequest) (response DeleteNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNodeResponse")
	}
	return
}

// deleteNode implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/nodePools/{nodePoolId}/node/{nodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/DeleteNode"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNodePool Delete a node pool.
// A default retry strategy applies to this operation DeleteNodePool()
func (client ContainerEngineClient) DeleteNodePool(ctx context.Context, request DeleteNodePoolRequest) (response DeleteNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNodePoolResponse")
	}
	return
}

// deleteNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/nodePools/{nodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/DeleteNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVirtualNodePool Delete a virtual node pool.
// A default retry strategy applies to this operation DeleteVirtualNodePool()
func (client ContainerEngineClient) DeleteVirtualNodePool(ctx context.Context, request DeleteVirtualNodePoolRequest) (response DeleteVirtualNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVirtualNodePoolResponse")
	}
	return
}

// deleteVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/virtualNodePools/{virtualNodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/DeleteVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWorkRequest Cancel a work request that has not started.
// A default retry strategy applies to this operation DeleteWorkRequest()
func (client ContainerEngineClient) DeleteWorkRequest(ctx context.Context, request DeleteWorkRequestRequest) (response DeleteWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWorkRequestResponse")
	}
	return
}

// deleteWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequest/DeleteWorkRequest"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWorkloadMapping Delete workloadMapping for a provisioned cluster.
// A default retry strategy applies to this operation DeleteWorkloadMapping()
func (client ContainerEngineClient) DeleteWorkloadMapping(ctx context.Context, request DeleteWorkloadMappingRequest) (response DeleteWorkloadMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteWorkloadMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWorkloadMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWorkloadMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWorkloadMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWorkloadMappingResponse")
	}
	return
}

// deleteWorkloadMapping implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) deleteWorkloadMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusters/{clusterId}/workloadMappings/{workloadMappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWorkloadMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkloadMapping/DeleteWorkloadMapping"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DeleteWorkloadMapping", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableAddon Disable addon for a provisioned cluster.
// A default retry strategy applies to this operation DisableAddon()
func (client ContainerEngineClient) DisableAddon(ctx context.Context, request DisableAddonRequest) (response DisableAddonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableAddon, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableAddonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableAddonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableAddonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableAddonResponse")
	}
	return
}

// disableAddon implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) disableAddon(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/clusters/{clusterId}/addons/{addonName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisableAddonResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/DisableAddon"
		err = common.PostProcessServiceError(err, "ContainerEngine", "DisableAddon", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAddon Get the specified addon for a cluster.
// A default retry strategy applies to this operation GetAddon()
func (client ContainerEngineClient) GetAddon(ctx context.Context, request GetAddonRequest) (response GetAddonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAddon, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAddonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAddonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAddonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAddonResponse")
	}
	return
}

// getAddon implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getAddon(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}/addons/{addonName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAddonResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/GetAddon"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetAddon", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCluster Get the details of a cluster.
// A default retry strategy applies to this operation GetCluster()
func (client ContainerEngineClient) GetCluster(ctx context.Context, request GetClusterRequest) (response GetClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterResponse")
	}
	return
}

// getCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/GetCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterAttachment Gets a ClusterAttachment by identifier
// A default retry strategy applies to this operation GetClusterAttachment()
func (client ContainerEngineClient) GetClusterAttachment(ctx context.Context, request GetClusterAttachmentRequest) (response GetClusterAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterAttachmentResponse")
	}
	return
}

// getClusterAttachment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterAttachments/{clusterAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterAttachment/GetClusterAttachment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterMigrateToNativeVcnStatus Get details on a cluster's migration to native VCN.
// A default retry strategy applies to this operation GetClusterMigrateToNativeVcnStatus()
func (client ContainerEngineClient) GetClusterMigrateToNativeVcnStatus(ctx context.Context, request GetClusterMigrateToNativeVcnStatusRequest) (response GetClusterMigrateToNativeVcnStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterMigrateToNativeVcnStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterMigrateToNativeVcnStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterMigrateToNativeVcnStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterMigrateToNativeVcnStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterMigrateToNativeVcnStatusResponse")
	}
	return
}

// getClusterMigrateToNativeVcnStatus implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterMigrateToNativeVcnStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}/migrateToNativeVcnStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterMigrateToNativeVcnStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterMigrateToNativeVcnStatus/GetClusterMigrateToNativeVcnStatus"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterMigrateToNativeVcnStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterNamespace Gets a ClusterNamespace by identifier
// A default retry strategy applies to this operation GetClusterNamespace()
func (client ContainerEngineClient) GetClusterNamespace(ctx context.Context, request GetClusterNamespaceRequest) (response GetClusterNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterNamespaceResponse")
	}
	return
}

// getClusterNamespace implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNamespaces/{clusterNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespace/GetClusterNamespace"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterNamespaceProfile Gets a ClusterNamespaceProfile by identifier
// A default retry strategy applies to this operation GetClusterNamespaceProfile()
func (client ContainerEngineClient) GetClusterNamespaceProfile(ctx context.Context, request GetClusterNamespaceProfileRequest) (response GetClusterNamespaceProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterNamespaceProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterNamespaceProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterNamespaceProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterNamespaceProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterNamespaceProfileResponse")
	}
	return
}

// getClusterNamespaceProfile implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterNamespaceProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNamespaceProfiles/{clusterNamespaceProfileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterNamespaceProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfile/GetClusterNamespaceProfile"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterNamespaceProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterNamespaceProfileVersion Gets a ClusterNamespaceProfileVersion by identifier
// A default retry strategy applies to this operation GetClusterNamespaceProfileVersion()
func (client ContainerEngineClient) GetClusterNamespaceProfileVersion(ctx context.Context, request GetClusterNamespaceProfileVersionRequest) (response GetClusterNamespaceProfileVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterNamespaceProfileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterNamespaceProfileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterNamespaceProfileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterNamespaceProfileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterNamespaceProfileVersionResponse")
	}
	return
}

// getClusterNamespaceProfileVersion implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterNamespaceProfileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNamespaceProfileVersions/{clusterNamespaceProfileVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterNamespaceProfileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfileVersion/GetClusterNamespaceProfileVersion"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterNamespaceProfileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetClusterOptions Get options available for clusters.
// A default retry strategy applies to this operation GetClusterOptions()
func (client ContainerEngineClient) GetClusterOptions(ctx context.Context, request GetClusterOptionsRequest) (response GetClusterOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getClusterOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetClusterOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetClusterOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetClusterOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetClusterOptionsResponse")
	}
	return
}

// getClusterOptions implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getClusterOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterOptions/{clusterOptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetClusterOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterOptions/GetClusterOptions"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetClusterOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCredentialRotationStatus Get cluster credential rotation status.
// A default retry strategy applies to this operation GetCredentialRotationStatus()
func (client ContainerEngineClient) GetCredentialRotationStatus(ctx context.Context, request GetCredentialRotationStatusRequest) (response GetCredentialRotationStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCredentialRotationStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCredentialRotationStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCredentialRotationStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCredentialRotationStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCredentialRotationStatusResponse")
	}
	return
}

// getCredentialRotationStatus implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getCredentialRotationStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}/credentialRotationStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCredentialRotationStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/CredentialRotationStatus/GetCredentialRotationStatus"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetCredentialRotationStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNodePool Get the details of a node pool.
// A default retry strategy applies to this operation GetNodePool()
func (client ContainerEngineClient) GetNodePool(ctx context.Context, request GetNodePoolRequest) (response GetNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNodePoolResponse")
	}
	return
}

// getNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/nodePools/{nodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/GetNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNodePoolOptions Get options available for node pools.
// A default retry strategy applies to this operation GetNodePoolOptions()
func (client ContainerEngineClient) GetNodePoolOptions(ctx context.Context, request GetNodePoolOptionsRequest) (response GetNodePoolOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNodePoolOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNodePoolOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNodePoolOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNodePoolOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNodePoolOptionsResponse")
	}
	return
}

// getNodePoolOptions implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getNodePoolOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/nodePoolOptions/{nodePoolOptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNodePoolOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePoolOptions/GetNodePoolOptions"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetNodePoolOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVirtualNode Get the details of a virtual node.
// A default retry strategy applies to this operation GetVirtualNode()
func (client ContainerEngineClient) GetVirtualNode(ctx context.Context, request GetVirtualNodeRequest) (response GetVirtualNodeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVirtualNode, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVirtualNodeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVirtualNodeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVirtualNodeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVirtualNodeResponse")
	}
	return
}

// getVirtualNode implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getVirtualNode(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools/{virtualNodePoolId}/virtualNodes/{virtualNodeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVirtualNodeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/GetVirtualNode"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetVirtualNode", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVirtualNodePool Get the details of a virtual node pool.
// A default retry strategy applies to this operation GetVirtualNodePool()
func (client ContainerEngineClient) GetVirtualNodePool(ctx context.Context, request GetVirtualNodePoolRequest) (response GetVirtualNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVirtualNodePoolResponse")
	}
	return
}

// getVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools/{virtualNodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/GetVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Get the details of a work request.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ContainerEngineClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client ContainerEngineClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkloadMapping Get the specified workloadMapping for a cluster.
// A default retry strategy applies to this operation GetWorkloadMapping()
func (client ContainerEngineClient) GetWorkloadMapping(ctx context.Context, request GetWorkloadMappingRequest) (response GetWorkloadMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkloadMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkloadMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkloadMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkloadMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkloadMappingResponse")
	}
	return
}

// getWorkloadMapping implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) getWorkloadMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}/workloadMappings/{workloadMappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkloadMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkloadMapping/GetWorkloadMapping"
		err = common.PostProcessServiceError(err, "ContainerEngine", "GetWorkloadMapping", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallAddon Install the specified addon for a cluster.
// A default retry strategy applies to this operation InstallAddon()
func (client ContainerEngineClient) InstallAddon(ctx context.Context, request InstallAddonRequest) (response InstallAddonResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installAddon, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallAddonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallAddonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallAddonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallAddonResponse")
	}
	return
}

// installAddon implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) installAddon(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/addons", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallAddonResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/InstallAddon"
		err = common.PostProcessServiceError(err, "ContainerEngine", "InstallAddon", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddonOptions Get list of supported addons for a specific kubernetes version.
// A default retry strategy applies to this operation ListAddonOptions()
func (client ContainerEngineClient) ListAddonOptions(ctx context.Context, request ListAddonOptionsRequest) (response ListAddonOptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddonOptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddonOptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddonOptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddonOptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddonOptionsResponse")
	}
	return
}

// listAddonOptions implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listAddonOptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/addonOptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddonOptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/AddonOptionSummary/ListAddonOptions"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListAddonOptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddons List addon for a provisioned cluster.
// A default retry strategy applies to this operation ListAddons()
func (client ContainerEngineClient) ListAddons(ctx context.Context, request ListAddonsRequest) (response ListAddonsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddons, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddonsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddonsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddonsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddonsResponse")
	}
	return
}

// listAddons implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listAddons(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}/addons", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddonsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/ListAddons"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListAddons", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusterAttachments Returns a list of ClusterAttachments.
// A default retry strategy applies to this operation ListClusterAttachments()
func (client ContainerEngineClient) ListClusterAttachments(ctx context.Context, request ListClusterAttachmentsRequest) (response ListClusterAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusterAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClusterAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClusterAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClusterAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClusterAttachmentsResponse")
	}
	return
}

// listClusterAttachments implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listClusterAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClusterAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterAttachment/ListClusterAttachments"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListClusterAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusterNamespaceProfileVersions Returns a list of ClusterNamespaceProfileVersions.
// A default retry strategy applies to this operation ListClusterNamespaceProfileVersions()
func (client ContainerEngineClient) ListClusterNamespaceProfileVersions(ctx context.Context, request ListClusterNamespaceProfileVersionsRequest) (response ListClusterNamespaceProfileVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusterNamespaceProfileVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClusterNamespaceProfileVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClusterNamespaceProfileVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClusterNamespaceProfileVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClusterNamespaceProfileVersionsResponse")
	}
	return
}

// listClusterNamespaceProfileVersions implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listClusterNamespaceProfileVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNamespaceProfileVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClusterNamespaceProfileVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfileVersion/ListClusterNamespaceProfileVersions"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListClusterNamespaceProfileVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusterNamespaceProfiles Returns a list of ClusterNamespaceProfiles.
// A default retry strategy applies to this operation ListClusterNamespaceProfiles()
func (client ContainerEngineClient) ListClusterNamespaceProfiles(ctx context.Context, request ListClusterNamespaceProfilesRequest) (response ListClusterNamespaceProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusterNamespaceProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClusterNamespaceProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClusterNamespaceProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClusterNamespaceProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClusterNamespaceProfilesResponse")
	}
	return
}

// listClusterNamespaceProfiles implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listClusterNamespaceProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNamespaceProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClusterNamespaceProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfile/ListClusterNamespaceProfiles"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListClusterNamespaceProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusterNamespaces Returns a list of ClusterNamespaces.
// A default retry strategy applies to this operation ListClusterNamespaces()
func (client ContainerEngineClient) ListClusterNamespaces(ctx context.Context, request ListClusterNamespacesRequest) (response ListClusterNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusterNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClusterNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClusterNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClusterNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClusterNamespacesResponse")
	}
	return
}

// listClusterNamespaces implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listClusterNamespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusterNamespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClusterNamespacesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespace/ListClusterNamespaces"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListClusterNamespaces", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListClusters List all the cluster objects in a compartment.
// A default retry strategy applies to this operation ListClusters()
func (client ContainerEngineClient) ListClusters(ctx context.Context, request ListClustersRequest) (response ListClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListClustersResponse")
	}
	return
}

// listClusters implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterSummary/ListClusters"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNodePools List all the node pools in a compartment, and optionally filter by cluster.
// A default retry strategy applies to this operation ListNodePools()
func (client ContainerEngineClient) ListNodePools(ctx context.Context, request ListNodePoolsRequest) (response ListNodePoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNodePools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNodePoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNodePoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNodePoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNodePoolsResponse")
	}
	return
}

// listNodePools implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listNodePools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/nodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNodePoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePoolSummary/ListNodePools"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListNodePools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPodShapes List all the Pod Shapes in a compartment.
// A default retry strategy applies to this operation ListPodShapes()
func (client ContainerEngineClient) ListPodShapes(ctx context.Context, request ListPodShapesRequest) (response ListPodShapesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPodShapes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPodShapesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPodShapesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPodShapesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPodShapesResponse")
	}
	return
}

// listPodShapes implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listPodShapes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/podShapes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPodShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/PodShapeSummary/ListPodShapes"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListPodShapes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVirtualNodePools List all the virtual node pools in a compartment, and optionally filter by cluster.
// A default retry strategy applies to this operation ListVirtualNodePools()
func (client ContainerEngineClient) ListVirtualNodePools(ctx context.Context, request ListVirtualNodePoolsRequest) (response ListVirtualNodePoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVirtualNodePools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVirtualNodePoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVirtualNodePoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVirtualNodePoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVirtualNodePoolsResponse")
	}
	return
}

// listVirtualNodePools implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listVirtualNodePools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVirtualNodePoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePoolSummary/ListVirtualNodePools"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListVirtualNodePools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVirtualNodes List virtual nodes in a virtual node pool.
// A default retry strategy applies to this operation ListVirtualNodes()
func (client ContainerEngineClient) ListVirtualNodes(ctx context.Context, request ListVirtualNodesRequest) (response ListVirtualNodesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVirtualNodes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVirtualNodesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVirtualNodesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVirtualNodesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVirtualNodesResponse")
	}
	return
}

// listVirtualNodes implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listVirtualNodes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualNodePools/{virtualNodePoolId}/virtualNodes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVirtualNodesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/ListVirtualNodes"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListVirtualNodes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Get the errors of a work request.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ContainerEngineClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client ContainerEngineClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Get the logs of a work request.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ContainerEngineClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client ContainerEngineClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests List all work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ContainerEngineClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client ContainerEngineClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkRequestSummary/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkloadMappings List workloadMappings for a provisioned cluster.
// A default retry strategy applies to this operation ListWorkloadMappings()
func (client ContainerEngineClient) ListWorkloadMappings(ctx context.Context, request ListWorkloadMappingsRequest) (response ListWorkloadMappingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkloadMappings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkloadMappingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkloadMappingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkloadMappingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkloadMappingsResponse")
	}
	return
}

// listWorkloadMappings implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) listWorkloadMappings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/clusters/{clusterId}/workloadMappings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkloadMappingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkloadMapping/ListWorkloadMappings"
		err = common.PostProcessServiceError(err, "ContainerEngine", "ListWorkloadMappings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartCredentialRotation Start cluster credential rotation by adding new credentials, old credentials will still work after this operation.
// A default retry strategy applies to this operation StartCredentialRotation()
func (client ContainerEngineClient) StartCredentialRotation(ctx context.Context, request StartCredentialRotationRequest) (response StartCredentialRotationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startCredentialRotation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartCredentialRotationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartCredentialRotationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartCredentialRotationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartCredentialRotationResponse")
	}
	return
}

// startCredentialRotation implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) startCredentialRotation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/actions/startCredentialRotation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartCredentialRotationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/StartCredentialRotation"
		err = common.PostProcessServiceError(err, "ContainerEngine", "StartCredentialRotation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAddon Update addon details for a cluster.
// A default retry strategy applies to this operation UpdateAddon()
func (client ContainerEngineClient) UpdateAddon(ctx context.Context, request UpdateAddonRequest) (response UpdateAddonResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAddon, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAddonResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAddonResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAddonResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAddonResponse")
	}
	return
}

// updateAddon implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateAddon(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusters/{clusterId}/addons/{addonName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAddonResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/UpdateAddon"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateAddon", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCluster Update the details of a cluster.
// A default retry strategy applies to this operation UpdateCluster()
func (client ContainerEngineClient) UpdateCluster(ctx context.Context, request UpdateClusterRequest) (response UpdateClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterResponse")
	}
	return
}

// updateCluster implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusters/{clusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/UpdateCluster"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateClusterAttachment Updates the ClusterAttachment
// A default retry strategy applies to this operation UpdateClusterAttachment()
func (client ContainerEngineClient) UpdateClusterAttachment(ctx context.Context, request UpdateClusterAttachmentRequest) (response UpdateClusterAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateClusterAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterAttachmentResponse")
	}
	return
}

// updateClusterAttachment implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateClusterAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusterAttachments/{clusterAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterAttachment/UpdateClusterAttachment"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateClusterAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateClusterEndpointConfig Update the details of the cluster endpoint configuration.
// A default retry strategy applies to this operation UpdateClusterEndpointConfig()
func (client ContainerEngineClient) UpdateClusterEndpointConfig(ctx context.Context, request UpdateClusterEndpointConfigRequest) (response UpdateClusterEndpointConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateClusterEndpointConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterEndpointConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterEndpointConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterEndpointConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterEndpointConfigResponse")
	}
	return
}

// updateClusterEndpointConfig implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateClusterEndpointConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/clusters/{clusterId}/actions/updateEndpointConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterEndpointConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/Cluster/UpdateClusterEndpointConfig"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateClusterEndpointConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateClusterNamespace Updates the ClusterNamespace
// A default retry strategy applies to this operation UpdateClusterNamespace()
func (client ContainerEngineClient) UpdateClusterNamespace(ctx context.Context, request UpdateClusterNamespaceRequest) (response UpdateClusterNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateClusterNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterNamespaceResponse")
	}
	return
}

// updateClusterNamespace implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateClusterNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusterNamespaces/{clusterNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespace/UpdateClusterNamespace"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateClusterNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateClusterNamespaceProfile Updates the ClusterNamespaceProfile
// A default retry strategy applies to this operation UpdateClusterNamespaceProfile()
func (client ContainerEngineClient) UpdateClusterNamespaceProfile(ctx context.Context, request UpdateClusterNamespaceProfileRequest) (response UpdateClusterNamespaceProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateClusterNamespaceProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterNamespaceProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterNamespaceProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterNamespaceProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterNamespaceProfileResponse")
	}
	return
}

// updateClusterNamespaceProfile implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateClusterNamespaceProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusterNamespaceProfiles/{clusterNamespaceProfileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterNamespaceProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfile/UpdateClusterNamespaceProfile"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateClusterNamespaceProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateClusterNamespaceProfileVersion Updates the ClusterNamespaceProfileVersion
// A default retry strategy applies to this operation UpdateClusterNamespaceProfileVersion()
func (client ContainerEngineClient) UpdateClusterNamespaceProfileVersion(ctx context.Context, request UpdateClusterNamespaceProfileVersionRequest) (response UpdateClusterNamespaceProfileVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateClusterNamespaceProfileVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateClusterNamespaceProfileVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateClusterNamespaceProfileVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateClusterNamespaceProfileVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateClusterNamespaceProfileVersionResponse")
	}
	return
}

// updateClusterNamespaceProfileVersion implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateClusterNamespaceProfileVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusterNamespaceProfileVersions/{clusterNamespaceProfileVersionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateClusterNamespaceProfileVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/ClusterNamespaceProfileVersion/UpdateClusterNamespaceProfileVersion"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateClusterNamespaceProfileVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNodePool Update the details of a node pool.
// A default retry strategy applies to this operation UpdateNodePool()
func (client ContainerEngineClient) UpdateNodePool(ctx context.Context, request UpdateNodePoolRequest) (response UpdateNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNodePoolResponse")
	}
	return
}

// updateNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/nodePools/{nodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/NodePool/UpdateNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVirtualNodePool Update the details of a virtual node pool.
// A default retry strategy applies to this operation UpdateVirtualNodePool()
func (client ContainerEngineClient) UpdateVirtualNodePool(ctx context.Context, request UpdateVirtualNodePoolRequest) (response UpdateVirtualNodePoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVirtualNodePool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVirtualNodePoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVirtualNodePoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVirtualNodePoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVirtualNodePoolResponse")
	}
	return
}

// updateVirtualNodePool implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateVirtualNodePool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/virtualNodePools/{virtualNodePoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVirtualNodePoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/VirtualNodePool/UpdateVirtualNodePool"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateVirtualNodePool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWorkloadMapping Update workloadMapping details for a cluster.
// A default retry strategy applies to this operation UpdateWorkloadMapping()
func (client ContainerEngineClient) UpdateWorkloadMapping(ctx context.Context, request UpdateWorkloadMappingRequest) (response UpdateWorkloadMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWorkloadMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWorkloadMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWorkloadMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWorkloadMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWorkloadMappingResponse")
	}
	return
}

// updateWorkloadMapping implements the OCIOperation interface (enables retrying operations)
func (client ContainerEngineClient) updateWorkloadMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/clusters/{clusterId}/workloadMappings/{workloadMappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWorkloadMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/containerengine/20180222/WorkloadMapping/UpdateWorkloadMapping"
		err = common.PostProcessServiceError(err, "ContainerEngine", "UpdateWorkloadMapping", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
