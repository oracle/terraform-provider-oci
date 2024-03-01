// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// AddressRuleServiceClient a client for AddressRuleService
type AddressRuleServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAddressRuleServiceClientWithConfigurationProvider Creates a new default AddressRuleService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAddressRuleServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AddressRuleServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("ospgateway"); !enabled {
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
	return newAddressRuleServiceClientFromBaseClient(baseClient, provider)
}

// NewAddressRuleServiceClientWithOboToken Creates a new default AddressRuleService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAddressRuleServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AddressRuleServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAddressRuleServiceClientFromBaseClient(baseClient, configProvider)
}

func newAddressRuleServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AddressRuleServiceClient, err error) {
	// AddressRuleService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AddressRuleService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AddressRuleServiceClient{BaseClient: baseClient}
	client.BasePath = "20191001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AddressRuleServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ospgateway", "https://osp-oci-integ.osp.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AddressRuleServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AddressRuleServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetAddressRule Get the address rule for the compartment based on the country code
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/GetAddressRule.go.html to see an example of how to use GetAddressRule API.
// A default retry strategy applies to this operation GetAddressRule()
func (client AddressRuleServiceClient) GetAddressRule(ctx context.Context, request GetAddressRuleRequest) (response GetAddressRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAddressRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAddressRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAddressRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAddressRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAddressRuleResponse")
	}
	return
}

// getAddressRule implements the OCIOperation interface (enables retrying operations)
func (client AddressRuleServiceClient) getAddressRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/addressRules/{countryCode}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAddressRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "AddressRuleService", "GetAddressRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
