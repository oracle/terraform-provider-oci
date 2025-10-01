// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OmhubResourceAnchorClient a client for OmhubResourceAnchor
type OmhubResourceAnchorClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOmhubResourceAnchorClientWithConfigurationProvider Creates a new default OmhubResourceAnchor client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOmhubResourceAnchorClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OmhubResourceAnchorClient, err error) {
	if enabled := common.CheckForEnabledServices("multicloud"); !enabled {
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
	return newOmhubResourceAnchorClientFromBaseClient(baseClient, provider)
}

// NewOmhubResourceAnchorClientWithOboToken Creates a new default OmhubResourceAnchor client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOmhubResourceAnchorClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OmhubResourceAnchorClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOmhubResourceAnchorClientFromBaseClient(baseClient, configProvider)
}

func newOmhubResourceAnchorClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OmhubResourceAnchorClient, err error) {
	// OmhubResourceAnchor service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OmhubResourceAnchor"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OmhubResourceAnchorClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OmhubResourceAnchorClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("multicloud", "https://multicloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OmhubResourceAnchorClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OmhubResourceAnchorClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetResourceAnchor Gets information about a ResourceAnchor.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/GetResourceAnchor.go.html to see an example of how to use GetResourceAnchor API.
// A default retry strategy applies to this operation GetResourceAnchor()
func (client OmhubResourceAnchorClient) GetResourceAnchor(ctx context.Context, request GetResourceAnchorRequest) (response GetResourceAnchorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResourceAnchor, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResourceAnchorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResourceAnchorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResourceAnchorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResourceAnchorResponse")
	}
	return
}

// getResourceAnchor implements the OCIOperation interface (enables retrying operations)
func (client OmhubResourceAnchorClient) getResourceAnchor(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceAnchors/{resourceAnchorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResourceAnchorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/ResourceAnchor/GetResourceAnchor"
		err = common.PostProcessServiceError(err, "OmhubResourceAnchor", "GetResourceAnchor", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceAnchors Gets a list of ResourceAnchors.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListResourceAnchors.go.html to see an example of how to use ListResourceAnchors API.
// A default retry strategy applies to this operation ListResourceAnchors()
func (client OmhubResourceAnchorClient) ListResourceAnchors(ctx context.Context, request ListResourceAnchorsRequest) (response ListResourceAnchorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceAnchors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceAnchorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceAnchorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceAnchorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceAnchorsResponse")
	}
	return
}

// listResourceAnchors implements the OCIOperation interface (enables retrying operations)
func (client OmhubResourceAnchorClient) listResourceAnchors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceAnchors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceAnchorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/ResourceAnchorCollection/ListResourceAnchors"
		err = common.PostProcessServiceError(err, "OmhubResourceAnchor", "ListResourceAnchors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
