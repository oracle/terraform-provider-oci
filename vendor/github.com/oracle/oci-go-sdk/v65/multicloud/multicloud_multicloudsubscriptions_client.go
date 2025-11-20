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

// MulticloudsubscriptionsClient a client for Multicloudsubscriptions
type MulticloudsubscriptionsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMulticloudsubscriptionsClientWithConfigurationProvider Creates a new default Multicloudsubscriptions client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMulticloudsubscriptionsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MulticloudsubscriptionsClient, err error) {
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
	return newMulticloudsubscriptionsClientFromBaseClient(baseClient, provider)
}

// NewMulticloudsubscriptionsClientWithOboToken Creates a new default Multicloudsubscriptions client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMulticloudsubscriptionsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MulticloudsubscriptionsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMulticloudsubscriptionsClientFromBaseClient(baseClient, configProvider)
}

func newMulticloudsubscriptionsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MulticloudsubscriptionsClient, err error) {
	// Multicloudsubscriptions service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Multicloudsubscriptions"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MulticloudsubscriptionsClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MulticloudsubscriptionsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("multicloud", "https://multicloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MulticloudsubscriptionsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MulticloudsubscriptionsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListMulticloudSubscriptions Gets a list of Multicloud Resources.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListMulticloudSubscriptions.go.html to see an example of how to use ListMulticloudSubscriptions API.
// A default retry strategy applies to this operation ListMulticloudSubscriptions()
func (client MulticloudsubscriptionsClient) ListMulticloudSubscriptions(ctx context.Context, request ListMulticloudSubscriptionsRequest) (response ListMulticloudSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMulticloudSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMulticloudSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMulticloudSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMulticloudSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMulticloudSubscriptionsResponse")
	}
	return
}

// listMulticloudSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client MulticloudsubscriptionsClient) listMulticloudSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/multicloudsubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMulticloudSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/MulticloudSubscriptionCollection/ListMulticloudSubscriptions"
		err = common.PostProcessServiceError(err, "Multicloudsubscriptions", "ListMulticloudSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
