// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GovernanceClient a client for Governance
type GovernanceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGovernanceClientWithConfigurationProvider Creates a new default Governance client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGovernanceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GovernanceClient, err error) {
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
	return newGovernanceClientFromBaseClient(baseClient, provider)
}

// NewGovernanceClientWithOboToken Creates a new default Governance client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGovernanceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GovernanceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGovernanceClientFromBaseClient(baseClient, configProvider)
}

func newGovernanceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GovernanceClient, err error) {
	// Governance service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Governance"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GovernanceClient{BaseClient: baseClient}
	client.BasePath = "20230401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GovernanceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GovernanceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GovernanceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddGovernance Starts a work request to opt the tenancy in to governance rules.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/AddGovernance.go.html to see an example of how to use AddGovernance API.
func (client GovernanceClient) AddGovernance(ctx context.Context, request AddGovernanceRequest) (response AddGovernanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addGovernance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddGovernanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddGovernanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddGovernanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddGovernanceResponse")
	}
	return
}

// addGovernance implements the OCIOperation interface (enables retrying operations)
func (client GovernanceClient) addGovernance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/organizations/{organizationId}/tenancies/{organizationTenancyId}/actions/addGovernance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddGovernanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/OrganizationTenancy/AddGovernance"
		err = common.PostProcessServiceError(err, "Governance", "AddGovernance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveGovernance Starts a work request to opt the tenancy out of governance rules.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/RemoveGovernance.go.html to see an example of how to use RemoveGovernance API.
func (client GovernanceClient) RemoveGovernance(ctx context.Context, request RemoveGovernanceRequest) (response RemoveGovernanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeGovernance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveGovernanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveGovernanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveGovernanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveGovernanceResponse")
	}
	return
}

// removeGovernance implements the OCIOperation interface (enables retrying operations)
func (client GovernanceClient) removeGovernance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/organizations/{organizationId}/tenancies/{organizationTenancyId}/actions/removeGovernance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveGovernanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/OrganizationTenancy/RemoveGovernance"
		err = common.PostProcessServiceError(err, "Governance", "RemoveGovernance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
