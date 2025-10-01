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

// MultiCloudsMetadataClient a client for MultiCloudsMetadata
type MultiCloudsMetadataClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMultiCloudsMetadataClientWithConfigurationProvider Creates a new default MultiCloudsMetadata client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMultiCloudsMetadataClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MultiCloudsMetadataClient, err error) {
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
	return newMultiCloudsMetadataClientFromBaseClient(baseClient, provider)
}

// NewMultiCloudsMetadataClientWithOboToken Creates a new default MultiCloudsMetadata client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMultiCloudsMetadataClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MultiCloudsMetadataClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMultiCloudsMetadataClientFromBaseClient(baseClient, configProvider)
}

func newMultiCloudsMetadataClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MultiCloudsMetadataClient, err error) {
	// MultiCloudsMetadata service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MultiCloudsMetadata"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MultiCloudsMetadataClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MultiCloudsMetadataClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("multicloud", "https://multicloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MultiCloudsMetadataClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MultiCloudsMetadataClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetMultiCloudMetadata Gets information about multicloud base compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/GetMultiCloudMetadata.go.html to see an example of how to use GetMultiCloudMetadata API.
// A default retry strategy applies to this operation GetMultiCloudMetadata()
func (client MultiCloudsMetadataClient) GetMultiCloudMetadata(ctx context.Context, request GetMultiCloudMetadataRequest) (response GetMultiCloudMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMultiCloudMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMultiCloudMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMultiCloudMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMultiCloudMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMultiCloudMetadataResponse")
	}
	return
}

// getMultiCloudMetadata implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudsMetadataClient) getMultiCloudMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/omHub/multiCloudsMetadata/{subscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMultiCloudMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/MultiCloudMetadata/GetMultiCloudMetadata"
		err = common.PostProcessServiceError(err, "MultiCloudsMetadata", "GetMultiCloudMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMultiCloudMetadata Gets a list of multicloud metadata with multicloud base compartment and subscription across Cloud Service Providers.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListMultiCloudMetadata.go.html to see an example of how to use ListMultiCloudMetadata API.
// A default retry strategy applies to this operation ListMultiCloudMetadata()
func (client MultiCloudsMetadataClient) ListMultiCloudMetadata(ctx context.Context, request ListMultiCloudMetadataRequest) (response ListMultiCloudMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMultiCloudMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMultiCloudMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMultiCloudMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMultiCloudMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMultiCloudMetadataResponse")
	}
	return
}

// listMultiCloudMetadata implements the OCIOperation interface (enables retrying operations)
func (client MultiCloudsMetadataClient) listMultiCloudMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/omHub/multiCloudsMetadata", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMultiCloudMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/MultiCloudMetadataCollection/ListMultiCloudMetadata"
		err = common.PostProcessServiceError(err, "MultiCloudsMetadata", "ListMultiCloudMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
