// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MediaServicesClient a client for MediaServices
type MediaServicesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMediaServicesClientWithConfigurationProvider Creates a new default MediaServices client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMediaServicesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MediaServicesClient, err error) {
	if enabled := common.CheckForEnabledServices("mediaservices"); !enabled {
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
	return newMediaServicesClientFromBaseClient(baseClient, provider)
}

// NewMediaServicesClientWithOboToken Creates a new default MediaServices client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMediaServicesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MediaServicesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMediaServicesClientFromBaseClient(baseClient, configProvider)
}

func newMediaServicesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MediaServicesClient, err error) {
	// MediaServices service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MediaServices"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MediaServicesClient{BaseClient: baseClient}
	client.BasePath = "20211101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MediaServicesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mediaservices", "https://mediaservices.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MediaServicesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MediaServicesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddMediaAssetLock Add a lock to an MediaAsset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/AddMediaAssetLock.go.html to see an example of how to use AddMediaAssetLock API.
// A default retry strategy applies to this operation AddMediaAssetLock()
func (client MediaServicesClient) AddMediaAssetLock(ctx context.Context, request AddMediaAssetLockRequest) (response AddMediaAssetLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addMediaAssetLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMediaAssetLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMediaAssetLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMediaAssetLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMediaAssetLockResponse")
	}
	return
}

// addMediaAssetLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) addMediaAssetLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaAssets/{mediaAssetId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMediaAssetLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/AddMediaAssetLock"
		err = common.PostProcessServiceError(err, "MediaServices", "AddMediaAssetLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddMediaWorkflowConfigurationLock Add a lock to a MediaWorkflowConfiguration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/AddMediaWorkflowConfigurationLock.go.html to see an example of how to use AddMediaWorkflowConfigurationLock API.
// A default retry strategy applies to this operation AddMediaWorkflowConfigurationLock()
func (client MediaServicesClient) AddMediaWorkflowConfigurationLock(ctx context.Context, request AddMediaWorkflowConfigurationLockRequest) (response AddMediaWorkflowConfigurationLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addMediaWorkflowConfigurationLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMediaWorkflowConfigurationLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMediaWorkflowConfigurationLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMediaWorkflowConfigurationLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMediaWorkflowConfigurationLockResponse")
	}
	return
}

// addMediaWorkflowConfigurationLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) addMediaWorkflowConfigurationLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowConfigurations/{mediaWorkflowConfigurationId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMediaWorkflowConfigurationLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfiguration/AddMediaWorkflowConfigurationLock"
		err = common.PostProcessServiceError(err, "MediaServices", "AddMediaWorkflowConfigurationLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddMediaWorkflowJobLock Add a lock to a MediaWorkflowJob.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/AddMediaWorkflowJobLock.go.html to see an example of how to use AddMediaWorkflowJobLock API.
// A default retry strategy applies to this operation AddMediaWorkflowJobLock()
func (client MediaServicesClient) AddMediaWorkflowJobLock(ctx context.Context, request AddMediaWorkflowJobLockRequest) (response AddMediaWorkflowJobLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addMediaWorkflowJobLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMediaWorkflowJobLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMediaWorkflowJobLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMediaWorkflowJobLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMediaWorkflowJobLockResponse")
	}
	return
}

// addMediaWorkflowJobLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) addMediaWorkflowJobLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowJobs/{mediaWorkflowJobId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMediaWorkflowJobLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/AddMediaWorkflowJobLock"
		err = common.PostProcessServiceError(err, "MediaServices", "AddMediaWorkflowJobLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddMediaWorkflowLock Add a lock to a MediaWorkflow.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/AddMediaWorkflowLock.go.html to see an example of how to use AddMediaWorkflowLock API.
// A default retry strategy applies to this operation AddMediaWorkflowLock()
func (client MediaServicesClient) AddMediaWorkflowLock(ctx context.Context, request AddMediaWorkflowLockRequest) (response AddMediaWorkflowLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addMediaWorkflowLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMediaWorkflowLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMediaWorkflowLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMediaWorkflowLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMediaWorkflowLockResponse")
	}
	return
}

// addMediaWorkflowLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) addMediaWorkflowLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflows/{mediaWorkflowId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMediaWorkflowLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/AddMediaWorkflowLock"
		err = common.PostProcessServiceError(err, "MediaServices", "AddMediaWorkflowLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddStreamCdnConfigLock Add a lock to a StreamCdnConfig.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/AddStreamCdnConfigLock.go.html to see an example of how to use AddStreamCdnConfigLock API.
// A default retry strategy applies to this operation AddStreamCdnConfigLock()
func (client MediaServicesClient) AddStreamCdnConfigLock(ctx context.Context, request AddStreamCdnConfigLockRequest) (response AddStreamCdnConfigLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addStreamCdnConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddStreamCdnConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddStreamCdnConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddStreamCdnConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddStreamCdnConfigLockResponse")
	}
	return
}

// addStreamCdnConfigLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) addStreamCdnConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamCdnConfigs/{streamCdnConfigId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddStreamCdnConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamCdnConfig/AddStreamCdnConfigLock"
		err = common.PostProcessServiceError(err, "MediaServices", "AddStreamCdnConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddStreamDistributionChannelLock Add a lock to a StreamDistributionChannel.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/AddStreamDistributionChannelLock.go.html to see an example of how to use AddStreamDistributionChannelLock API.
// A default retry strategy applies to this operation AddStreamDistributionChannelLock()
func (client MediaServicesClient) AddStreamDistributionChannelLock(ctx context.Context, request AddStreamDistributionChannelLockRequest) (response AddStreamDistributionChannelLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addStreamDistributionChannelLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddStreamDistributionChannelLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddStreamDistributionChannelLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddStreamDistributionChannelLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddStreamDistributionChannelLockResponse")
	}
	return
}

// addStreamDistributionChannelLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) addStreamDistributionChannelLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamDistributionChannels/{streamDistributionChannelId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddStreamDistributionChannelLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/AddStreamDistributionChannelLock"
		err = common.PostProcessServiceError(err, "MediaServices", "AddStreamDistributionChannelLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddStreamPackagingConfigLock Add a lock to a StreamPackagingConfig.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/AddStreamPackagingConfigLock.go.html to see an example of how to use AddStreamPackagingConfigLock API.
// A default retry strategy applies to this operation AddStreamPackagingConfigLock()
func (client MediaServicesClient) AddStreamPackagingConfigLock(ctx context.Context, request AddStreamPackagingConfigLockRequest) (response AddStreamPackagingConfigLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addStreamPackagingConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddStreamPackagingConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddStreamPackagingConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddStreamPackagingConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddStreamPackagingConfigLockResponse")
	}
	return
}

// addStreamPackagingConfigLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) addStreamPackagingConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamPackagingConfigs/{streamPackagingConfigId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddStreamPackagingConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamPackagingConfig/AddStreamPackagingConfigLock"
		err = common.PostProcessServiceError(err, "MediaServices", "AddStreamPackagingConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &streampackagingconfig{})
	return response, err
}

// ChangeMediaAssetCompartment Moves a MediaAsset resource from one compartment identifier to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ChangeMediaAssetCompartment.go.html to see an example of how to use ChangeMediaAssetCompartment API.
// A default retry strategy applies to this operation ChangeMediaAssetCompartment()
func (client MediaServicesClient) ChangeMediaAssetCompartment(ctx context.Context, request ChangeMediaAssetCompartmentRequest) (response ChangeMediaAssetCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMediaAssetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMediaAssetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMediaAssetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMediaAssetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMediaAssetCompartmentResponse")
	}
	return
}

// changeMediaAssetCompartment implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) changeMediaAssetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaAssets/{mediaAssetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMediaAssetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/ChangeMediaAssetCompartment"
		err = common.PostProcessServiceError(err, "MediaServices", "ChangeMediaAssetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMediaWorkflowCompartment Moves a MediaWorkflow resource from one compartment identifier to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ChangeMediaWorkflowCompartment.go.html to see an example of how to use ChangeMediaWorkflowCompartment API.
// A default retry strategy applies to this operation ChangeMediaWorkflowCompartment()
func (client MediaServicesClient) ChangeMediaWorkflowCompartment(ctx context.Context, request ChangeMediaWorkflowCompartmentRequest) (response ChangeMediaWorkflowCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMediaWorkflowCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMediaWorkflowCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMediaWorkflowCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMediaWorkflowCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMediaWorkflowCompartmentResponse")
	}
	return
}

// changeMediaWorkflowCompartment implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) changeMediaWorkflowCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflows/{mediaWorkflowId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMediaWorkflowCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/ChangeMediaWorkflowCompartment"
		err = common.PostProcessServiceError(err, "MediaServices", "ChangeMediaWorkflowCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMediaWorkflowConfigurationCompartment Moves a MediaWorkflowConfiguration resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ChangeMediaWorkflowConfigurationCompartment.go.html to see an example of how to use ChangeMediaWorkflowConfigurationCompartment API.
// A default retry strategy applies to this operation ChangeMediaWorkflowConfigurationCompartment()
func (client MediaServicesClient) ChangeMediaWorkflowConfigurationCompartment(ctx context.Context, request ChangeMediaWorkflowConfigurationCompartmentRequest) (response ChangeMediaWorkflowConfigurationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMediaWorkflowConfigurationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMediaWorkflowConfigurationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMediaWorkflowConfigurationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMediaWorkflowConfigurationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMediaWorkflowConfigurationCompartmentResponse")
	}
	return
}

// changeMediaWorkflowConfigurationCompartment implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) changeMediaWorkflowConfigurationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowConfigurations/{mediaWorkflowConfigurationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMediaWorkflowConfigurationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfiguration/ChangeMediaWorkflowConfigurationCompartment"
		err = common.PostProcessServiceError(err, "MediaServices", "ChangeMediaWorkflowConfigurationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMediaWorkflowJobCompartment Moves a MediaWorkflowJob resource from one compartment identifier to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ChangeMediaWorkflowJobCompartment.go.html to see an example of how to use ChangeMediaWorkflowJobCompartment API.
// A default retry strategy applies to this operation ChangeMediaWorkflowJobCompartment()
func (client MediaServicesClient) ChangeMediaWorkflowJobCompartment(ctx context.Context, request ChangeMediaWorkflowJobCompartmentRequest) (response ChangeMediaWorkflowJobCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMediaWorkflowJobCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMediaWorkflowJobCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMediaWorkflowJobCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMediaWorkflowJobCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMediaWorkflowJobCompartmentResponse")
	}
	return
}

// changeMediaWorkflowJobCompartment implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) changeMediaWorkflowJobCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowJobs/{mediaWorkflowJobId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMediaWorkflowJobCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/ChangeMediaWorkflowJobCompartment"
		err = common.PostProcessServiceError(err, "MediaServices", "ChangeMediaWorkflowJobCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeStreamDistributionChannelCompartment Moves a Stream Distribution Channel resource from one compartment identifier to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ChangeStreamDistributionChannelCompartment.go.html to see an example of how to use ChangeStreamDistributionChannelCompartment API.
// A default retry strategy applies to this operation ChangeStreamDistributionChannelCompartment()
func (client MediaServicesClient) ChangeStreamDistributionChannelCompartment(ctx context.Context, request ChangeStreamDistributionChannelCompartmentRequest) (response ChangeStreamDistributionChannelCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeStreamDistributionChannelCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeStreamDistributionChannelCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeStreamDistributionChannelCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeStreamDistributionChannelCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeStreamDistributionChannelCompartmentResponse")
	}
	return
}

// changeStreamDistributionChannelCompartment implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) changeStreamDistributionChannelCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamDistributionChannels/{streamDistributionChannelId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeStreamDistributionChannelCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/ChangeStreamDistributionChannelCompartment"
		err = common.PostProcessServiceError(err, "MediaServices", "ChangeStreamDistributionChannelCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMediaAsset Creates a new MediaAsset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/CreateMediaAsset.go.html to see an example of how to use CreateMediaAsset API.
// A default retry strategy applies to this operation CreateMediaAsset()
func (client MediaServicesClient) CreateMediaAsset(ctx context.Context, request CreateMediaAssetRequest) (response CreateMediaAssetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMediaAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMediaAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMediaAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMediaAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMediaAssetResponse")
	}
	return
}

// createMediaAsset implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) createMediaAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaAssets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMediaAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/CreateMediaAsset"
		err = common.PostProcessServiceError(err, "MediaServices", "CreateMediaAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMediaWorkflow Creates a new MediaWorkflow.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/CreateMediaWorkflow.go.html to see an example of how to use CreateMediaWorkflow API.
// A default retry strategy applies to this operation CreateMediaWorkflow()
func (client MediaServicesClient) CreateMediaWorkflow(ctx context.Context, request CreateMediaWorkflowRequest) (response CreateMediaWorkflowResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMediaWorkflow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMediaWorkflowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMediaWorkflowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMediaWorkflowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMediaWorkflowResponse")
	}
	return
}

// createMediaWorkflow implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) createMediaWorkflow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMediaWorkflowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/CreateMediaWorkflow"
		err = common.PostProcessServiceError(err, "MediaServices", "CreateMediaWorkflow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMediaWorkflowConfiguration Creates a new MediaWorkflowConfiguration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/CreateMediaWorkflowConfiguration.go.html to see an example of how to use CreateMediaWorkflowConfiguration API.
// A default retry strategy applies to this operation CreateMediaWorkflowConfiguration()
func (client MediaServicesClient) CreateMediaWorkflowConfiguration(ctx context.Context, request CreateMediaWorkflowConfigurationRequest) (response CreateMediaWorkflowConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMediaWorkflowConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMediaWorkflowConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMediaWorkflowConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMediaWorkflowConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMediaWorkflowConfigurationResponse")
	}
	return
}

// createMediaWorkflowConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) createMediaWorkflowConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMediaWorkflowConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfiguration/CreateMediaWorkflowConfiguration"
		err = common.PostProcessServiceError(err, "MediaServices", "CreateMediaWorkflowConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMediaWorkflowJob Run the MediaWorkflow according to the given mediaWorkflow definition and configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/CreateMediaWorkflowJob.go.html to see an example of how to use CreateMediaWorkflowJob API.
// A default retry strategy applies to this operation CreateMediaWorkflowJob()
func (client MediaServicesClient) CreateMediaWorkflowJob(ctx context.Context, request CreateMediaWorkflowJobRequest) (response CreateMediaWorkflowJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMediaWorkflowJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMediaWorkflowJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMediaWorkflowJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMediaWorkflowJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMediaWorkflowJobResponse")
	}
	return
}

// createMediaWorkflowJob implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) createMediaWorkflowJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMediaWorkflowJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/CreateMediaWorkflowJob"
		err = common.PostProcessServiceError(err, "MediaServices", "CreateMediaWorkflowJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateStreamCdnConfig Creates a new CDN Configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/CreateStreamCdnConfig.go.html to see an example of how to use CreateStreamCdnConfig API.
// A default retry strategy applies to this operation CreateStreamCdnConfig()
func (client MediaServicesClient) CreateStreamCdnConfig(ctx context.Context, request CreateStreamCdnConfigRequest) (response CreateStreamCdnConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createStreamCdnConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateStreamCdnConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateStreamCdnConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStreamCdnConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStreamCdnConfigResponse")
	}
	return
}

// createStreamCdnConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) createStreamCdnConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamCdnConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateStreamCdnConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamCdnConfig/CreateStreamCdnConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "CreateStreamCdnConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateStreamDistributionChannel Creates a new Stream Distribution Channel.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/CreateStreamDistributionChannel.go.html to see an example of how to use CreateStreamDistributionChannel API.
// A default retry strategy applies to this operation CreateStreamDistributionChannel()
func (client MediaServicesClient) CreateStreamDistributionChannel(ctx context.Context, request CreateStreamDistributionChannelRequest) (response CreateStreamDistributionChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createStreamDistributionChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateStreamDistributionChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateStreamDistributionChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStreamDistributionChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStreamDistributionChannelResponse")
	}
	return
}

// createStreamDistributionChannel implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) createStreamDistributionChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamDistributionChannels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateStreamDistributionChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/CreateStreamDistributionChannel"
		err = common.PostProcessServiceError(err, "MediaServices", "CreateStreamDistributionChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateStreamPackagingConfig Creates a new Packaging Configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/CreateStreamPackagingConfig.go.html to see an example of how to use CreateStreamPackagingConfig API.
// A default retry strategy applies to this operation CreateStreamPackagingConfig()
func (client MediaServicesClient) CreateStreamPackagingConfig(ctx context.Context, request CreateStreamPackagingConfigRequest) (response CreateStreamPackagingConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createStreamPackagingConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateStreamPackagingConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateStreamPackagingConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateStreamPackagingConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateStreamPackagingConfigResponse")
	}
	return
}

// createStreamPackagingConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) createStreamPackagingConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamPackagingConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateStreamPackagingConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamPackagingConfig/CreateStreamPackagingConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "CreateStreamPackagingConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &streampackagingconfig{})
	return response, err
}

// DeleteMediaAsset Deletes a MediaAsset resource by identifier. If DeleteChildren is passed in as the mode, all the assets with the parentMediaAssetId matching the ID will be deleted. If DeleteDerivatives is set as the mode, all the assets with the masterMediaAssetId matching the ID will be deleted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteMediaAsset.go.html to see an example of how to use DeleteMediaAsset API.
// A default retry strategy applies to this operation DeleteMediaAsset()
func (client MediaServicesClient) DeleteMediaAsset(ctx context.Context, request DeleteMediaAssetRequest) (response DeleteMediaAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMediaAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMediaAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMediaAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMediaAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMediaAssetResponse")
	}
	return
}

// deleteMediaAsset implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteMediaAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mediaAssets/{mediaAssetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMediaAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/DeleteMediaAsset"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteMediaAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMediaAssetDistributionChannelAttachment Deletes a MediaAsset from the DistributionChannel by identifiers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteMediaAssetDistributionChannelAttachment.go.html to see an example of how to use DeleteMediaAssetDistributionChannelAttachment API.
// A default retry strategy applies to this operation DeleteMediaAssetDistributionChannelAttachment()
func (client MediaServicesClient) DeleteMediaAssetDistributionChannelAttachment(ctx context.Context, request DeleteMediaAssetDistributionChannelAttachmentRequest) (response DeleteMediaAssetDistributionChannelAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMediaAssetDistributionChannelAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMediaAssetDistributionChannelAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMediaAssetDistributionChannelAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMediaAssetDistributionChannelAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMediaAssetDistributionChannelAttachmentResponse")
	}
	return
}

// deleteMediaAssetDistributionChannelAttachment implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteMediaAssetDistributionChannelAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mediaAssets/{mediaAssetId}/distributionChannelAttachments/{distributionChannelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMediaAssetDistributionChannelAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAssetDistributionChannelAttachment/DeleteMediaAssetDistributionChannelAttachment"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteMediaAssetDistributionChannelAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMediaWorkflow The MediaWorkflow lifecycleState will change to DELETED.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteMediaWorkflow.go.html to see an example of how to use DeleteMediaWorkflow API.
// A default retry strategy applies to this operation DeleteMediaWorkflow()
func (client MediaServicesClient) DeleteMediaWorkflow(ctx context.Context, request DeleteMediaWorkflowRequest) (response DeleteMediaWorkflowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMediaWorkflow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMediaWorkflowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMediaWorkflowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMediaWorkflowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMediaWorkflowResponse")
	}
	return
}

// deleteMediaWorkflow implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteMediaWorkflow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mediaWorkflows/{mediaWorkflowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMediaWorkflowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/DeleteMediaWorkflow"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteMediaWorkflow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMediaWorkflowConfiguration Deletes a MediaWorkflowConfiguration resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteMediaWorkflowConfiguration.go.html to see an example of how to use DeleteMediaWorkflowConfiguration API.
// A default retry strategy applies to this operation DeleteMediaWorkflowConfiguration()
func (client MediaServicesClient) DeleteMediaWorkflowConfiguration(ctx context.Context, request DeleteMediaWorkflowConfigurationRequest) (response DeleteMediaWorkflowConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMediaWorkflowConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMediaWorkflowConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMediaWorkflowConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMediaWorkflowConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMediaWorkflowConfigurationResponse")
	}
	return
}

// deleteMediaWorkflowConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteMediaWorkflowConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mediaWorkflowConfigurations/{mediaWorkflowConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMediaWorkflowConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfiguration/DeleteMediaWorkflowConfiguration"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteMediaWorkflowConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMediaWorkflowJob This is an asynchronous operation. The MediaWorkflowJob lifecycleState will change to CANCELING temporarily until the job is completely CANCELED.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteMediaWorkflowJob.go.html to see an example of how to use DeleteMediaWorkflowJob API.
// A default retry strategy applies to this operation DeleteMediaWorkflowJob()
func (client MediaServicesClient) DeleteMediaWorkflowJob(ctx context.Context, request DeleteMediaWorkflowJobRequest) (response DeleteMediaWorkflowJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMediaWorkflowJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMediaWorkflowJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMediaWorkflowJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMediaWorkflowJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMediaWorkflowJobResponse")
	}
	return
}

// deleteMediaWorkflowJob implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteMediaWorkflowJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/mediaWorkflowJobs/{mediaWorkflowJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMediaWorkflowJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/DeleteMediaWorkflowJob"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteMediaWorkflowJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteStreamCdnConfig The StreamCdnConfig lifecycleState will change to DELETED.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteStreamCdnConfig.go.html to see an example of how to use DeleteStreamCdnConfig API.
// A default retry strategy applies to this operation DeleteStreamCdnConfig()
func (client MediaServicesClient) DeleteStreamCdnConfig(ctx context.Context, request DeleteStreamCdnConfigRequest) (response DeleteStreamCdnConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStreamCdnConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteStreamCdnConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteStreamCdnConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteStreamCdnConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteStreamCdnConfigResponse")
	}
	return
}

// deleteStreamCdnConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteStreamCdnConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/streamCdnConfigs/{streamCdnConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteStreamCdnConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamCdnConfig/DeleteStreamCdnConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteStreamCdnConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteStreamDistributionChannel The Stream Distribution Channel lifecycleState will change to DELETED.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteStreamDistributionChannel.go.html to see an example of how to use DeleteStreamDistributionChannel API.
// A default retry strategy applies to this operation DeleteStreamDistributionChannel()
func (client MediaServicesClient) DeleteStreamDistributionChannel(ctx context.Context, request DeleteStreamDistributionChannelRequest) (response DeleteStreamDistributionChannelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStreamDistributionChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteStreamDistributionChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteStreamDistributionChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteStreamDistributionChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteStreamDistributionChannelResponse")
	}
	return
}

// deleteStreamDistributionChannel implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteStreamDistributionChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/streamDistributionChannels/{streamDistributionChannelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteStreamDistributionChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/DeleteStreamDistributionChannel"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteStreamDistributionChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteStreamPackagingConfig The Stream Packaging Configuration lifecycleState will change to DELETED.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/DeleteStreamPackagingConfig.go.html to see an example of how to use DeleteStreamPackagingConfig API.
// A default retry strategy applies to this operation DeleteStreamPackagingConfig()
func (client MediaServicesClient) DeleteStreamPackagingConfig(ctx context.Context, request DeleteStreamPackagingConfigRequest) (response DeleteStreamPackagingConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteStreamPackagingConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteStreamPackagingConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteStreamPackagingConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteStreamPackagingConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteStreamPackagingConfigResponse")
	}
	return
}

// deleteStreamPackagingConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) deleteStreamPackagingConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/streamPackagingConfigs/{streamPackagingConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteStreamPackagingConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamPackagingConfig/DeleteStreamPackagingConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "DeleteStreamPackagingConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMediaAsset Gets a MediaAsset by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetMediaAsset.go.html to see an example of how to use GetMediaAsset API.
// A default retry strategy applies to this operation GetMediaAsset()
func (client MediaServicesClient) GetMediaAsset(ctx context.Context, request GetMediaAssetRequest) (response GetMediaAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMediaAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMediaAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMediaAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMediaAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMediaAssetResponse")
	}
	return
}

// getMediaAsset implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getMediaAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaAssets/{mediaAssetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMediaAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/GetMediaAsset"
		err = common.PostProcessServiceError(err, "MediaServices", "GetMediaAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMediaAssetDistributionChannelAttachment Gets a MediaAssetDistributionChannelAttachment for a MediaAsset by identifiers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetMediaAssetDistributionChannelAttachment.go.html to see an example of how to use GetMediaAssetDistributionChannelAttachment API.
// A default retry strategy applies to this operation GetMediaAssetDistributionChannelAttachment()
func (client MediaServicesClient) GetMediaAssetDistributionChannelAttachment(ctx context.Context, request GetMediaAssetDistributionChannelAttachmentRequest) (response GetMediaAssetDistributionChannelAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMediaAssetDistributionChannelAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMediaAssetDistributionChannelAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMediaAssetDistributionChannelAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMediaAssetDistributionChannelAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMediaAssetDistributionChannelAttachmentResponse")
	}
	return
}

// getMediaAssetDistributionChannelAttachment implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getMediaAssetDistributionChannelAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaAssets/{mediaAssetId}/distributionChannelAttachments/{distributionChannelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMediaAssetDistributionChannelAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAssetDistributionChannelAttachment/GetMediaAssetDistributionChannelAttachment"
		err = common.PostProcessServiceError(err, "MediaServices", "GetMediaAssetDistributionChannelAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMediaWorkflow Gets a MediaWorkflow by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetMediaWorkflow.go.html to see an example of how to use GetMediaWorkflow API.
// A default retry strategy applies to this operation GetMediaWorkflow()
func (client MediaServicesClient) GetMediaWorkflow(ctx context.Context, request GetMediaWorkflowRequest) (response GetMediaWorkflowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMediaWorkflow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMediaWorkflowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMediaWorkflowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMediaWorkflowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMediaWorkflowResponse")
	}
	return
}

// getMediaWorkflow implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getMediaWorkflow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaWorkflows/{mediaWorkflowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMediaWorkflowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/GetMediaWorkflow"
		err = common.PostProcessServiceError(err, "MediaServices", "GetMediaWorkflow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMediaWorkflowConfiguration Gets a MediaWorkflowConfiguration by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetMediaWorkflowConfiguration.go.html to see an example of how to use GetMediaWorkflowConfiguration API.
// A default retry strategy applies to this operation GetMediaWorkflowConfiguration()
func (client MediaServicesClient) GetMediaWorkflowConfiguration(ctx context.Context, request GetMediaWorkflowConfigurationRequest) (response GetMediaWorkflowConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMediaWorkflowConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMediaWorkflowConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMediaWorkflowConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMediaWorkflowConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMediaWorkflowConfigurationResponse")
	}
	return
}

// getMediaWorkflowConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getMediaWorkflowConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaWorkflowConfigurations/{mediaWorkflowConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMediaWorkflowConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfiguration/GetMediaWorkflowConfiguration"
		err = common.PostProcessServiceError(err, "MediaServices", "GetMediaWorkflowConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMediaWorkflowJob Gets the MediaWorkflowJob.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetMediaWorkflowJob.go.html to see an example of how to use GetMediaWorkflowJob API.
// A default retry strategy applies to this operation GetMediaWorkflowJob()
func (client MediaServicesClient) GetMediaWorkflowJob(ctx context.Context, request GetMediaWorkflowJobRequest) (response GetMediaWorkflowJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMediaWorkflowJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMediaWorkflowJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMediaWorkflowJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMediaWorkflowJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMediaWorkflowJobResponse")
	}
	return
}

// getMediaWorkflowJob implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getMediaWorkflowJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaWorkflowJobs/{mediaWorkflowJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMediaWorkflowJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/GetMediaWorkflowJob"
		err = common.PostProcessServiceError(err, "MediaServices", "GetMediaWorkflowJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStreamCdnConfig Gets a StreamCdnConfig by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetStreamCdnConfig.go.html to see an example of how to use GetStreamCdnConfig API.
// A default retry strategy applies to this operation GetStreamCdnConfig()
func (client MediaServicesClient) GetStreamCdnConfig(ctx context.Context, request GetStreamCdnConfigRequest) (response GetStreamCdnConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStreamCdnConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStreamCdnConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStreamCdnConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStreamCdnConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStreamCdnConfigResponse")
	}
	return
}

// getStreamCdnConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getStreamCdnConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streamCdnConfigs/{streamCdnConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStreamCdnConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamCdnConfig/GetStreamCdnConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "GetStreamCdnConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStreamDistributionChannel Gets a Stream Distribution Channel by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetStreamDistributionChannel.go.html to see an example of how to use GetStreamDistributionChannel API.
// A default retry strategy applies to this operation GetStreamDistributionChannel()
func (client MediaServicesClient) GetStreamDistributionChannel(ctx context.Context, request GetStreamDistributionChannelRequest) (response GetStreamDistributionChannelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStreamDistributionChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStreamDistributionChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStreamDistributionChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStreamDistributionChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStreamDistributionChannelResponse")
	}
	return
}

// getStreamDistributionChannel implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getStreamDistributionChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streamDistributionChannels/{streamDistributionChannelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStreamDistributionChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/GetStreamDistributionChannel"
		err = common.PostProcessServiceError(err, "MediaServices", "GetStreamDistributionChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStreamPackagingConfig Gets a Stream Packaging Configuration by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GetStreamPackagingConfig.go.html to see an example of how to use GetStreamPackagingConfig API.
// A default retry strategy applies to this operation GetStreamPackagingConfig()
func (client MediaServicesClient) GetStreamPackagingConfig(ctx context.Context, request GetStreamPackagingConfigRequest) (response GetStreamPackagingConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStreamPackagingConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStreamPackagingConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStreamPackagingConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStreamPackagingConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStreamPackagingConfigResponse")
	}
	return
}

// getStreamPackagingConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) getStreamPackagingConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streamPackagingConfigs/{streamPackagingConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStreamPackagingConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamPackagingConfig/GetStreamPackagingConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "GetStreamPackagingConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &streampackagingconfig{})
	return response, err
}

// IngestStreamDistributionChannel Ingests an Asset into a Distribution Channel.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/IngestStreamDistributionChannel.go.html to see an example of how to use IngestStreamDistributionChannel API.
// A default retry strategy applies to this operation IngestStreamDistributionChannel()
func (client MediaServicesClient) IngestStreamDistributionChannel(ctx context.Context, request IngestStreamDistributionChannelRequest) (response IngestStreamDistributionChannelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ingestStreamDistributionChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IngestStreamDistributionChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IngestStreamDistributionChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IngestStreamDistributionChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IngestStreamDistributionChannelResponse")
	}
	return
}

// ingestStreamDistributionChannel implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) ingestStreamDistributionChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamDistributionChannels/{streamDistributionChannelId}/actions/ingest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IngestStreamDistributionChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/IngestStreamDistributionChannel"
		err = common.PostProcessServiceError(err, "MediaServices", "IngestStreamDistributionChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMediaAssetDistributionChannelAttachments Lists the MediaAssetDistributionChannelAttachments for a MediaAsset by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaAssetDistributionChannelAttachments.go.html to see an example of how to use ListMediaAssetDistributionChannelAttachments API.
// A default retry strategy applies to this operation ListMediaAssetDistributionChannelAttachments()
func (client MediaServicesClient) ListMediaAssetDistributionChannelAttachments(ctx context.Context, request ListMediaAssetDistributionChannelAttachmentsRequest) (response ListMediaAssetDistributionChannelAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMediaAssetDistributionChannelAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMediaAssetDistributionChannelAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMediaAssetDistributionChannelAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMediaAssetDistributionChannelAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMediaAssetDistributionChannelAttachmentsResponse")
	}
	return
}

// listMediaAssetDistributionChannelAttachments implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listMediaAssetDistributionChannelAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaAssets/{mediaAssetId}/distributionChannelAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMediaAssetDistributionChannelAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAssetDistributionChannelAttachmentCollection/ListMediaAssetDistributionChannelAttachments"
		err = common.PostProcessServiceError(err, "MediaServices", "ListMediaAssetDistributionChannelAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMediaAssets Returns a list of MediaAssetSummary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaAssets.go.html to see an example of how to use ListMediaAssets API.
// A default retry strategy applies to this operation ListMediaAssets()
func (client MediaServicesClient) ListMediaAssets(ctx context.Context, request ListMediaAssetsRequest) (response ListMediaAssetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMediaAssets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMediaAssetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMediaAssetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMediaAssetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMediaAssetsResponse")
	}
	return
}

// listMediaAssets implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listMediaAssets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaAssets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMediaAssetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/ListMediaAssets"
		err = common.PostProcessServiceError(err, "MediaServices", "ListMediaAssets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMediaWorkflowConfigurations Returns a list of MediaWorkflowConfigurations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflowConfigurations.go.html to see an example of how to use ListMediaWorkflowConfigurations API.
// A default retry strategy applies to this operation ListMediaWorkflowConfigurations()
func (client MediaServicesClient) ListMediaWorkflowConfigurations(ctx context.Context, request ListMediaWorkflowConfigurationsRequest) (response ListMediaWorkflowConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMediaWorkflowConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMediaWorkflowConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMediaWorkflowConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMediaWorkflowConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMediaWorkflowConfigurationsResponse")
	}
	return
}

// listMediaWorkflowConfigurations implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listMediaWorkflowConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaWorkflowConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMediaWorkflowConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfigurationCollection/ListMediaWorkflowConfigurations"
		err = common.PostProcessServiceError(err, "MediaServices", "ListMediaWorkflowConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMediaWorkflowJobs Lists the MediaWorkflowJobs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflowJobs.go.html to see an example of how to use ListMediaWorkflowJobs API.
// A default retry strategy applies to this operation ListMediaWorkflowJobs()
func (client MediaServicesClient) ListMediaWorkflowJobs(ctx context.Context, request ListMediaWorkflowJobsRequest) (response ListMediaWorkflowJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMediaWorkflowJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMediaWorkflowJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMediaWorkflowJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMediaWorkflowJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMediaWorkflowJobsResponse")
	}
	return
}

// listMediaWorkflowJobs implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listMediaWorkflowJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaWorkflowJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMediaWorkflowJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/ListMediaWorkflowJobs"
		err = common.PostProcessServiceError(err, "MediaServices", "ListMediaWorkflowJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMediaWorkflowTaskDeclarations Returns a list of MediaWorkflowTaskDeclarations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflowTaskDeclarations.go.html to see an example of how to use ListMediaWorkflowTaskDeclarations API.
// A default retry strategy applies to this operation ListMediaWorkflowTaskDeclarations()
func (client MediaServicesClient) ListMediaWorkflowTaskDeclarations(ctx context.Context, request ListMediaWorkflowTaskDeclarationsRequest) (response ListMediaWorkflowTaskDeclarationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMediaWorkflowTaskDeclarations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMediaWorkflowTaskDeclarationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMediaWorkflowTaskDeclarationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMediaWorkflowTaskDeclarationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMediaWorkflowTaskDeclarationsResponse")
	}
	return
}

// listMediaWorkflowTaskDeclarations implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listMediaWorkflowTaskDeclarations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaWorkflowTaskDeclarations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMediaWorkflowTaskDeclarationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowTaskDeclarationCollection/ListMediaWorkflowTaskDeclarations"
		err = common.PostProcessServiceError(err, "MediaServices", "ListMediaWorkflowTaskDeclarations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMediaWorkflows Lists the MediaWorkflows.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListMediaWorkflows.go.html to see an example of how to use ListMediaWorkflows API.
// A default retry strategy applies to this operation ListMediaWorkflows()
func (client MediaServicesClient) ListMediaWorkflows(ctx context.Context, request ListMediaWorkflowsRequest) (response ListMediaWorkflowsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMediaWorkflows, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMediaWorkflowsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMediaWorkflowsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMediaWorkflowsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMediaWorkflowsResponse")
	}
	return
}

// listMediaWorkflows implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listMediaWorkflows(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/mediaWorkflows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMediaWorkflowsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/ListMediaWorkflows"
		err = common.PostProcessServiceError(err, "MediaServices", "ListMediaWorkflows", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStreamCdnConfigs Lists the StreamCdnConfig.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListStreamCdnConfigs.go.html to see an example of how to use ListStreamCdnConfigs API.
// A default retry strategy applies to this operation ListStreamCdnConfigs()
func (client MediaServicesClient) ListStreamCdnConfigs(ctx context.Context, request ListStreamCdnConfigsRequest) (response ListStreamCdnConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStreamCdnConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStreamCdnConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStreamCdnConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStreamCdnConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStreamCdnConfigsResponse")
	}
	return
}

// listStreamCdnConfigs implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listStreamCdnConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streamCdnConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStreamCdnConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamCdnConfig/ListStreamCdnConfigs"
		err = common.PostProcessServiceError(err, "MediaServices", "ListStreamCdnConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStreamDistributionChannels Lists the Stream Distribution Channels.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListStreamDistributionChannels.go.html to see an example of how to use ListStreamDistributionChannels API.
// A default retry strategy applies to this operation ListStreamDistributionChannels()
func (client MediaServicesClient) ListStreamDistributionChannels(ctx context.Context, request ListStreamDistributionChannelsRequest) (response ListStreamDistributionChannelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStreamDistributionChannels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStreamDistributionChannelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStreamDistributionChannelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStreamDistributionChannelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStreamDistributionChannelsResponse")
	}
	return
}

// listStreamDistributionChannels implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listStreamDistributionChannels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streamDistributionChannels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStreamDistributionChannelsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/ListStreamDistributionChannels"
		err = common.PostProcessServiceError(err, "MediaServices", "ListStreamDistributionChannels", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStreamPackagingConfigs Lists the Stream Packaging Configurations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListStreamPackagingConfigs.go.html to see an example of how to use ListStreamPackagingConfigs API.
// A default retry strategy applies to this operation ListStreamPackagingConfigs()
func (client MediaServicesClient) ListStreamPackagingConfigs(ctx context.Context, request ListStreamPackagingConfigsRequest) (response ListStreamPackagingConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStreamPackagingConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStreamPackagingConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStreamPackagingConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStreamPackagingConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStreamPackagingConfigsResponse")
	}
	return
}

// listStreamPackagingConfigs implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listStreamPackagingConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/streamPackagingConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListStreamPackagingConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamPackagingConfig/ListStreamPackagingConfigs"
		err = common.PostProcessServiceError(err, "MediaServices", "ListStreamPackagingConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSystemMediaWorkflows Lists the SystemMediaWorkflows that can be used to run a job by name or as a template to create a MediaWorkflow.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/ListSystemMediaWorkflows.go.html to see an example of how to use ListSystemMediaWorkflows API.
// A default retry strategy applies to this operation ListSystemMediaWorkflows()
func (client MediaServicesClient) ListSystemMediaWorkflows(ctx context.Context, request ListSystemMediaWorkflowsRequest) (response ListSystemMediaWorkflowsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSystemMediaWorkflows, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSystemMediaWorkflowsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSystemMediaWorkflowsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSystemMediaWorkflowsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSystemMediaWorkflowsResponse")
	}
	return
}

// listSystemMediaWorkflows implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) listSystemMediaWorkflows(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/systemMediaWorkflows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSystemMediaWorkflowsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/ListSystemMediaWorkflows"
		err = common.PostProcessServiceError(err, "MediaServices", "ListSystemMediaWorkflows", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveMediaAssetLock Remove a lock to an MediaAsset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/RemoveMediaAssetLock.go.html to see an example of how to use RemoveMediaAssetLock API.
// A default retry strategy applies to this operation RemoveMediaAssetLock()
func (client MediaServicesClient) RemoveMediaAssetLock(ctx context.Context, request RemoveMediaAssetLockRequest) (response RemoveMediaAssetLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeMediaAssetLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveMediaAssetLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveMediaAssetLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveMediaAssetLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveMediaAssetLockResponse")
	}
	return
}

// removeMediaAssetLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) removeMediaAssetLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaAssets/{mediaAssetId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveMediaAssetLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/RemoveMediaAssetLock"
		err = common.PostProcessServiceError(err, "MediaServices", "RemoveMediaAssetLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveMediaWorkflowConfigurationLock Remove a lock from a MediaWorkflowConfiguration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/RemoveMediaWorkflowConfigurationLock.go.html to see an example of how to use RemoveMediaWorkflowConfigurationLock API.
// A default retry strategy applies to this operation RemoveMediaWorkflowConfigurationLock()
func (client MediaServicesClient) RemoveMediaWorkflowConfigurationLock(ctx context.Context, request RemoveMediaWorkflowConfigurationLockRequest) (response RemoveMediaWorkflowConfigurationLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeMediaWorkflowConfigurationLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveMediaWorkflowConfigurationLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveMediaWorkflowConfigurationLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveMediaWorkflowConfigurationLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveMediaWorkflowConfigurationLockResponse")
	}
	return
}

// removeMediaWorkflowConfigurationLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) removeMediaWorkflowConfigurationLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowConfigurations/{mediaWorkflowConfigurationId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveMediaWorkflowConfigurationLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfiguration/RemoveMediaWorkflowConfigurationLock"
		err = common.PostProcessServiceError(err, "MediaServices", "RemoveMediaWorkflowConfigurationLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveMediaWorkflowJobLock Remove a lock from a MediaWorkflowJob.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/RemoveMediaWorkflowJobLock.go.html to see an example of how to use RemoveMediaWorkflowJobLock API.
// A default retry strategy applies to this operation RemoveMediaWorkflowJobLock()
func (client MediaServicesClient) RemoveMediaWorkflowJobLock(ctx context.Context, request RemoveMediaWorkflowJobLockRequest) (response RemoveMediaWorkflowJobLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeMediaWorkflowJobLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveMediaWorkflowJobLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveMediaWorkflowJobLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveMediaWorkflowJobLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveMediaWorkflowJobLockResponse")
	}
	return
}

// removeMediaWorkflowJobLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) removeMediaWorkflowJobLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflowJobs/{mediaWorkflowJobId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveMediaWorkflowJobLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/RemoveMediaWorkflowJobLock"
		err = common.PostProcessServiceError(err, "MediaServices", "RemoveMediaWorkflowJobLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveMediaWorkflowLock Remove a lock from a MediaWorkflow.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/RemoveMediaWorkflowLock.go.html to see an example of how to use RemoveMediaWorkflowLock API.
// A default retry strategy applies to this operation RemoveMediaWorkflowLock()
func (client MediaServicesClient) RemoveMediaWorkflowLock(ctx context.Context, request RemoveMediaWorkflowLockRequest) (response RemoveMediaWorkflowLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeMediaWorkflowLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveMediaWorkflowLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveMediaWorkflowLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveMediaWorkflowLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveMediaWorkflowLockResponse")
	}
	return
}

// removeMediaWorkflowLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) removeMediaWorkflowLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/mediaWorkflows/{mediaWorkflowId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveMediaWorkflowLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/RemoveMediaWorkflowLock"
		err = common.PostProcessServiceError(err, "MediaServices", "RemoveMediaWorkflowLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveStreamCdnConfigLock Remove a lock from a StreamCdnConfig.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/RemoveStreamCdnConfigLock.go.html to see an example of how to use RemoveStreamCdnConfigLock API.
// A default retry strategy applies to this operation RemoveStreamCdnConfigLock()
func (client MediaServicesClient) RemoveStreamCdnConfigLock(ctx context.Context, request RemoveStreamCdnConfigLockRequest) (response RemoveStreamCdnConfigLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeStreamCdnConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveStreamCdnConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveStreamCdnConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveStreamCdnConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveStreamCdnConfigLockResponse")
	}
	return
}

// removeStreamCdnConfigLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) removeStreamCdnConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamCdnConfigs/{streamCdnConfigId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveStreamCdnConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamCdnConfig/RemoveStreamCdnConfigLock"
		err = common.PostProcessServiceError(err, "MediaServices", "RemoveStreamCdnConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveStreamDistributionChannelLock Remove a lock to a StreamDistributionChannel.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/RemoveStreamDistributionChannelLock.go.html to see an example of how to use RemoveStreamDistributionChannelLock API.
// A default retry strategy applies to this operation RemoveStreamDistributionChannelLock()
func (client MediaServicesClient) RemoveStreamDistributionChannelLock(ctx context.Context, request RemoveStreamDistributionChannelLockRequest) (response RemoveStreamDistributionChannelLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeStreamDistributionChannelLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveStreamDistributionChannelLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveStreamDistributionChannelLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveStreamDistributionChannelLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveStreamDistributionChannelLockResponse")
	}
	return
}

// removeStreamDistributionChannelLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) removeStreamDistributionChannelLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamDistributionChannels/{streamDistributionChannelId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveStreamDistributionChannelLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/RemoveStreamDistributionChannelLock"
		err = common.PostProcessServiceError(err, "MediaServices", "RemoveStreamDistributionChannelLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveStreamPackagingConfigLock Remove a lock from a StreamPackagingConfig.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/RemoveStreamPackagingConfigLock.go.html to see an example of how to use RemoveStreamPackagingConfigLock API.
// A default retry strategy applies to this operation RemoveStreamPackagingConfigLock()
func (client MediaServicesClient) RemoveStreamPackagingConfigLock(ctx context.Context, request RemoveStreamPackagingConfigLockRequest) (response RemoveStreamPackagingConfigLockResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeStreamPackagingConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveStreamPackagingConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveStreamPackagingConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveStreamPackagingConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveStreamPackagingConfigLockResponse")
	}
	return
}

// removeStreamPackagingConfigLock implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) removeStreamPackagingConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/streamPackagingConfigs/{streamPackagingConfigId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveStreamPackagingConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamPackagingConfig/RemoveStreamPackagingConfigLock"
		err = common.PostProcessServiceError(err, "MediaServices", "RemoveStreamPackagingConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &streampackagingconfig{})
	return response, err
}

// UpdateMediaAsset Updates the MediaAsset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/UpdateMediaAsset.go.html to see an example of how to use UpdateMediaAsset API.
// A default retry strategy applies to this operation UpdateMediaAsset()
func (client MediaServicesClient) UpdateMediaAsset(ctx context.Context, request UpdateMediaAssetRequest) (response UpdateMediaAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMediaAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMediaAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMediaAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMediaAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMediaAssetResponse")
	}
	return
}

// updateMediaAsset implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) updateMediaAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mediaAssets/{mediaAssetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMediaAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaAsset/UpdateMediaAsset"
		err = common.PostProcessServiceError(err, "MediaServices", "UpdateMediaAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMediaWorkflow Updates the MediaWorkflow.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/UpdateMediaWorkflow.go.html to see an example of how to use UpdateMediaWorkflow API.
// A default retry strategy applies to this operation UpdateMediaWorkflow()
func (client MediaServicesClient) UpdateMediaWorkflow(ctx context.Context, request UpdateMediaWorkflowRequest) (response UpdateMediaWorkflowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMediaWorkflow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMediaWorkflowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMediaWorkflowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMediaWorkflowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMediaWorkflowResponse")
	}
	return
}

// updateMediaWorkflow implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) updateMediaWorkflow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mediaWorkflows/{mediaWorkflowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMediaWorkflowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflow/UpdateMediaWorkflow"
		err = common.PostProcessServiceError(err, "MediaServices", "UpdateMediaWorkflow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMediaWorkflowConfiguration Updates the MediaWorkflowConfiguration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/UpdateMediaWorkflowConfiguration.go.html to see an example of how to use UpdateMediaWorkflowConfiguration API.
// A default retry strategy applies to this operation UpdateMediaWorkflowConfiguration()
func (client MediaServicesClient) UpdateMediaWorkflowConfiguration(ctx context.Context, request UpdateMediaWorkflowConfigurationRequest) (response UpdateMediaWorkflowConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMediaWorkflowConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMediaWorkflowConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMediaWorkflowConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMediaWorkflowConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMediaWorkflowConfigurationResponse")
	}
	return
}

// updateMediaWorkflowConfiguration implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) updateMediaWorkflowConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mediaWorkflowConfigurations/{mediaWorkflowConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMediaWorkflowConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowConfiguration/UpdateMediaWorkflowConfiguration"
		err = common.PostProcessServiceError(err, "MediaServices", "UpdateMediaWorkflowConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMediaWorkflowJob Updates the MediaWorkflowJob.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/UpdateMediaWorkflowJob.go.html to see an example of how to use UpdateMediaWorkflowJob API.
// A default retry strategy applies to this operation UpdateMediaWorkflowJob()
func (client MediaServicesClient) UpdateMediaWorkflowJob(ctx context.Context, request UpdateMediaWorkflowJobRequest) (response UpdateMediaWorkflowJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMediaWorkflowJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMediaWorkflowJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMediaWorkflowJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMediaWorkflowJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMediaWorkflowJobResponse")
	}
	return
}

// updateMediaWorkflowJob implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) updateMediaWorkflowJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/mediaWorkflowJobs/{mediaWorkflowJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMediaWorkflowJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/MediaWorkflowJob/UpdateMediaWorkflowJob"
		err = common.PostProcessServiceError(err, "MediaServices", "UpdateMediaWorkflowJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateStreamCdnConfig Updates the StreamCdnConfig.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/UpdateStreamCdnConfig.go.html to see an example of how to use UpdateStreamCdnConfig API.
// A default retry strategy applies to this operation UpdateStreamCdnConfig()
func (client MediaServicesClient) UpdateStreamCdnConfig(ctx context.Context, request UpdateStreamCdnConfigRequest) (response UpdateStreamCdnConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStreamCdnConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStreamCdnConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStreamCdnConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStreamCdnConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStreamCdnConfigResponse")
	}
	return
}

// updateStreamCdnConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) updateStreamCdnConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streamCdnConfigs/{streamCdnConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateStreamCdnConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamCdnConfig/UpdateStreamCdnConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "UpdateStreamCdnConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateStreamDistributionChannel Updates the Stream Distribution Channel.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/UpdateStreamDistributionChannel.go.html to see an example of how to use UpdateStreamDistributionChannel API.
// A default retry strategy applies to this operation UpdateStreamDistributionChannel()
func (client MediaServicesClient) UpdateStreamDistributionChannel(ctx context.Context, request UpdateStreamDistributionChannelRequest) (response UpdateStreamDistributionChannelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStreamDistributionChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStreamDistributionChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStreamDistributionChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStreamDistributionChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStreamDistributionChannelResponse")
	}
	return
}

// updateStreamDistributionChannel implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) updateStreamDistributionChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streamDistributionChannels/{streamDistributionChannelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateStreamDistributionChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/UpdateStreamDistributionChannel"
		err = common.PostProcessServiceError(err, "MediaServices", "UpdateStreamDistributionChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateStreamPackagingConfig Updates the Stream Packaging Configuration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/UpdateStreamPackagingConfig.go.html to see an example of how to use UpdateStreamPackagingConfig API.
// A default retry strategy applies to this operation UpdateStreamPackagingConfig()
func (client MediaServicesClient) UpdateStreamPackagingConfig(ctx context.Context, request UpdateStreamPackagingConfigRequest) (response UpdateStreamPackagingConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStreamPackagingConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStreamPackagingConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStreamPackagingConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStreamPackagingConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStreamPackagingConfigResponse")
	}
	return
}

// updateStreamPackagingConfig implements the OCIOperation interface (enables retrying operations)
func (client MediaServicesClient) updateStreamPackagingConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/streamPackagingConfigs/{streamPackagingConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateStreamPackagingConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamPackagingConfig/UpdateStreamPackagingConfig"
		err = common.PostProcessServiceError(err, "MediaServices", "UpdateStreamPackagingConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &streampackagingconfig{})
	return response, err
}
