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

// MediaStreamClient a client for MediaStream
type MediaStreamClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMediaStreamClientWithConfigurationProvider Creates a new default MediaStream client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMediaStreamClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MediaStreamClient, err error) {
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
	return newMediaStreamClientFromBaseClient(baseClient, provider)
}

// NewMediaStreamClientWithOboToken Creates a new default MediaStream client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMediaStreamClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MediaStreamClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMediaStreamClientFromBaseClient(baseClient, configProvider)
}

func newMediaStreamClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MediaStreamClient, err error) {
	// MediaStream service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MediaStream"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MediaStreamClient{BaseClient: baseClient}
	client.BasePath = "20211101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MediaStreamClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mediaservices", "https://mediaservices.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MediaStreamClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MediaStreamClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GeneratePlaylist Gets the playlist content for the specified Packaging Configuration and Media Asset combination.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GeneratePlaylist.go.html to see an example of how to use GeneratePlaylist API.
// A default retry strategy applies to this operation GeneratePlaylist()
func (client MediaStreamClient) GeneratePlaylist(ctx context.Context, request GeneratePlaylistRequest) (response GeneratePlaylistResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generatePlaylist, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GeneratePlaylistResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GeneratePlaylistResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GeneratePlaylistResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GeneratePlaylistResponse")
	}
	return
}

// generatePlaylist implements the OCIOperation interface (enables retrying operations)
func (client MediaStreamClient) generatePlaylist(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/actions/generatePlaylist", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GeneratePlaylistResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/GeneratePlaylist"
		err = common.PostProcessServiceError(err, "MediaStream", "GeneratePlaylist", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateSessionToken Generate a new streaming session token.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mediaservices/GenerateSessionToken.go.html to see an example of how to use GenerateSessionToken API.
// A default retry strategy applies to this operation GenerateSessionToken()
func (client MediaStreamClient) GenerateSessionToken(ctx context.Context, request GenerateSessionTokenRequest) (response GenerateSessionTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateSessionToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateSessionTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateSessionTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateSessionTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateSessionTokenResponse")
	}
	return
}

// generateSessionToken implements the OCIOperation interface (enables retrying operations)
func (client MediaStreamClient) generateSessionToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/generateSessionToken", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateSessionTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dms/20211101/StreamDistributionChannel/GenerateSessionToken"
		err = common.PostProcessServiceError(err, "MediaStream", "GenerateSessionToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
