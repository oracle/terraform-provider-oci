// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// LinkFeaturesClient a client for LinkFeatures
type LinkFeaturesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLinkFeaturesClientWithConfigurationProvider Creates a new default LinkFeatures client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLinkFeaturesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LinkFeaturesClient, err error) {
	if enabled := common.CheckForEnabledServices("tenantmanagercontrolplane"); !enabled {
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
	return newLinkFeaturesClientFromBaseClient(baseClient, provider)
}

// NewLinkFeaturesClientWithOboToken Creates a new default LinkFeatures client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLinkFeaturesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LinkFeaturesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLinkFeaturesClientFromBaseClient(baseClient, configProvider)
}

func newLinkFeaturesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LinkFeaturesClient, err error) {
	// LinkFeatures service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LinkFeatures"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LinkFeaturesClient{BaseClient: baseClient}
	client.BasePath = "20230401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LinkFeaturesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LinkFeaturesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LinkFeaturesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListLinkFeatures Return a list of link features.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListLinkFeatures.go.html to see an example of how to use ListLinkFeatures API.
// A default retry strategy applies to this operation ListLinkFeatures()
func (client LinkFeaturesClient) ListLinkFeatures(ctx context.Context, request ListLinkFeaturesRequest) (response ListLinkFeaturesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLinkFeatures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLinkFeaturesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLinkFeaturesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLinkFeaturesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLinkFeaturesResponse")
	}
	return
}

// listLinkFeatures implements the OCIOperation interface (enables retrying operations)
func (client LinkFeaturesClient) listLinkFeatures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/linkFeatures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLinkFeaturesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "linkFeatures", "ListLinkFeatures")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/LinkFeaturesCollection/ListLinkFeatures"
		err = common.PostProcessServiceError(err, "LinkFeatures", "ListLinkFeatures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
