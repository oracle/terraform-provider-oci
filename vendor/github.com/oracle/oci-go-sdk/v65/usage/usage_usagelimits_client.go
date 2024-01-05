// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage Proxy API
//
// Use the Usage Proxy API to list Oracle Support Rewards, view related detailed usage information, and manage users who redeem rewards. For more information, see Oracle Support Rewards Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/supportrewardsoverview.htm).
//

package usage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// UsagelimitsClient a client for Usagelimits
type UsagelimitsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewUsagelimitsClientWithConfigurationProvider Creates a new default Usagelimits client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewUsagelimitsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client UsagelimitsClient, err error) {
	if enabled := common.CheckForEnabledServices("usage"); !enabled {
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
	return newUsagelimitsClientFromBaseClient(baseClient, provider)
}

// NewUsagelimitsClientWithOboToken Creates a new default Usagelimits client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewUsagelimitsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client UsagelimitsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newUsagelimitsClientFromBaseClient(baseClient, configProvider)
}

func newUsagelimitsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client UsagelimitsClient, err error) {
	// Usagelimits service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Usagelimits"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = UsagelimitsClient{BaseClient: baseClient}
	client.BasePath = "20190111"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *UsagelimitsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("usage", "https://identity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *UsagelimitsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *UsagelimitsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListUsageLimits Returns the list of usage limit for the subscription ID and tenant ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListUsageLimits.go.html to see an example of how to use ListUsageLimits API.
// A default retry strategy applies to this operation ListUsageLimits()
func (client UsagelimitsClient) ListUsageLimits(ctx context.Context, request ListUsageLimitsRequest) (response ListUsageLimitsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUsageLimits, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsageLimitsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsageLimitsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsageLimitsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsageLimitsResponse")
	}
	return
}

// listUsageLimits implements the OCIOperation interface (enables retrying operations)
func (client UsagelimitsClient) listUsageLimits(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/usagelimits", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUsageLimitsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/usage-proxy/20190111/UsageLimitSummary/ListUsageLimits"
		err = common.PostProcessServiceError(err, "Usagelimits", "ListUsageLimits", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
