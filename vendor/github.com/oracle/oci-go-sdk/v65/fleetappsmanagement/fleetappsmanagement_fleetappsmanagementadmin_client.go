// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FleetAppsManagementAdminClient a client for FleetAppsManagementAdmin
type FleetAppsManagementAdminClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFleetAppsManagementAdminClientWithConfigurationProvider Creates a new default FleetAppsManagementAdmin client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFleetAppsManagementAdminClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FleetAppsManagementAdminClient, err error) {
	if enabled := common.CheckForEnabledServices("fleetappsmanagement"); !enabled {
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
	return newFleetAppsManagementAdminClientFromBaseClient(baseClient, provider)
}

// NewFleetAppsManagementAdminClientWithOboToken Creates a new default FleetAppsManagementAdmin client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFleetAppsManagementAdminClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FleetAppsManagementAdminClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFleetAppsManagementAdminClientFromBaseClient(baseClient, configProvider)
}

func newFleetAppsManagementAdminClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FleetAppsManagementAdminClient, err error) {
	// FleetAppsManagementAdmin service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FleetAppsManagementAdmin"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FleetAppsManagementAdminClient{BaseClient: baseClient}
	client.BasePath = "20230831"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FleetAppsManagementAdminClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fleetappsmanagement", "https://fams.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FleetAppsManagementAdminClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FleetAppsManagementAdminClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateCompliancePolicyRule Creates a CompliancePolicyRule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateCompliancePolicyRule.go.html to see an example of how to use CreateCompliancePolicyRule API.
// A default retry strategy applies to this operation CreateCompliancePolicyRule()
func (client FleetAppsManagementAdminClient) CreateCompliancePolicyRule(ctx context.Context, request CreateCompliancePolicyRuleRequest) (response CreateCompliancePolicyRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCompliancePolicyRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCompliancePolicyRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCompliancePolicyRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCompliancePolicyRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCompliancePolicyRuleResponse")
	}
	return
}

// createCompliancePolicyRule implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) createCompliancePolicyRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/compliancePolicyRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCompliancePolicyRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/CompliancePolicyRule/CreateCompliancePolicyRule"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "CreateCompliancePolicyRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOnboarding Onboard a tenant to Fleet Application Management.
// The onboarding process lets Fleet Application Management create a few required policies that you need to start using it and its features.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateOnboarding.go.html to see an example of how to use CreateOnboarding API.
// A default retry strategy applies to this operation CreateOnboarding()
func (client FleetAppsManagementAdminClient) CreateOnboarding(ctx context.Context, request CreateOnboardingRequest) (response CreateOnboardingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOnboarding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOnboardingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOnboardingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOnboardingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOnboardingResponse")
	}
	return
}

// createOnboarding implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) createOnboarding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/Onboardings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOnboardingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Onboarding/CreateOnboarding"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "CreateOnboarding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePlatformConfiguration Creates a new PlatformConfiguration.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreatePlatformConfiguration.go.html to see an example of how to use CreatePlatformConfiguration API.
// A default retry strategy applies to this operation CreatePlatformConfiguration()
func (client FleetAppsManagementAdminClient) CreatePlatformConfiguration(ctx context.Context, request CreatePlatformConfigurationRequest) (response CreatePlatformConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPlatformConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePlatformConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePlatformConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePlatformConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePlatformConfigurationResponse")
	}
	return
}

// createPlatformConfiguration implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) createPlatformConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/platformConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePlatformConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/PlatformConfiguration/CreatePlatformConfiguration"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "CreatePlatformConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProperty Create a business-specific metadata property in Fleet Application Management and capture the business metadata classifications.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateProperty.go.html to see an example of how to use CreateProperty API.
// A default retry strategy applies to this operation CreateProperty()
func (client FleetAppsManagementAdminClient) CreateProperty(ctx context.Context, request CreatePropertyRequest) (response CreatePropertyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePropertyResponse")
	}
	return
}

// createProperty implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) createProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/properties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Property/CreateProperty"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "CreateProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCompliancePolicyRule Deletes a CompliancePolicyRule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteCompliancePolicyRule.go.html to see an example of how to use DeleteCompliancePolicyRule API.
// A default retry strategy applies to this operation DeleteCompliancePolicyRule()
func (client FleetAppsManagementAdminClient) DeleteCompliancePolicyRule(ctx context.Context, request DeleteCompliancePolicyRuleRequest) (response DeleteCompliancePolicyRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCompliancePolicyRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCompliancePolicyRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCompliancePolicyRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCompliancePolicyRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCompliancePolicyRuleResponse")
	}
	return
}

// deleteCompliancePolicyRule implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) deleteCompliancePolicyRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/compliancePolicyRules/{compliancePolicyRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCompliancePolicyRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/CompliancePolicyRule/DeleteCompliancePolicyRule"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "DeleteCompliancePolicyRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOnboarding Deletes Fleet Application Management onboarding resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteOnboarding.go.html to see an example of how to use DeleteOnboarding API.
// A default retry strategy applies to this operation DeleteOnboarding()
func (client FleetAppsManagementAdminClient) DeleteOnboarding(ctx context.Context, request DeleteOnboardingRequest) (response DeleteOnboardingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOnboarding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOnboardingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOnboardingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOnboardingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOnboardingResponse")
	}
	return
}

// deleteOnboarding implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) deleteOnboarding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/Onboardings/{onboardingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOnboardingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Onboarding/DeleteOnboarding"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "DeleteOnboarding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePlatformConfiguration Deletes a PlatformConfiguration resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeletePlatformConfiguration.go.html to see an example of how to use DeletePlatformConfiguration API.
// A default retry strategy applies to this operation DeletePlatformConfiguration()
func (client FleetAppsManagementAdminClient) DeletePlatformConfiguration(ctx context.Context, request DeletePlatformConfigurationRequest) (response DeletePlatformConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePlatformConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePlatformConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePlatformConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePlatformConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePlatformConfigurationResponse")
	}
	return
}

// deletePlatformConfiguration implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) deletePlatformConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/platformConfigurations/{platformConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePlatformConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/PlatformConfiguration/DeletePlatformConfiguration"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "DeletePlatformConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProperty Delete a property in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteProperty.go.html to see an example of how to use DeleteProperty API.
// A default retry strategy applies to this operation DeleteProperty()
func (client FleetAppsManagementAdminClient) DeleteProperty(ctx context.Context, request DeletePropertyRequest) (response DeletePropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePropertyResponse")
	}
	return
}

// deleteProperty implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) deleteProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/properties/{propertyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Property/DeleteProperty"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "DeleteProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableLatestPolicy Enable Policies for a newer version of Fleet Application Management
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/EnableLatestPolicy.go.html to see an example of how to use EnableLatestPolicy API.
// A default retry strategy applies to this operation EnableLatestPolicy()
func (client FleetAppsManagementAdminClient) EnableLatestPolicy(ctx context.Context, request EnableLatestPolicyRequest) (response EnableLatestPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableLatestPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableLatestPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableLatestPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableLatestPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableLatestPolicyResponse")
	}
	return
}

// enableLatestPolicy implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) enableLatestPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/Onboardings/{onboardingId}/actions/enableLatestPolicy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableLatestPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Onboarding/EnableLatestPolicy"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "EnableLatestPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCompliancePolicy Gets information about a CompliancePolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetCompliancePolicy.go.html to see an example of how to use GetCompliancePolicy API.
// A default retry strategy applies to this operation GetCompliancePolicy()
func (client FleetAppsManagementAdminClient) GetCompliancePolicy(ctx context.Context, request GetCompliancePolicyRequest) (response GetCompliancePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCompliancePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCompliancePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCompliancePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCompliancePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCompliancePolicyResponse")
	}
	return
}

// getCompliancePolicy implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) getCompliancePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compliancePolicies/{compliancePolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCompliancePolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/CompliancePolicy/GetCompliancePolicy"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "GetCompliancePolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCompliancePolicyRule Gets information about a CompliancePolicyRule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetCompliancePolicyRule.go.html to see an example of how to use GetCompliancePolicyRule API.
// A default retry strategy applies to this operation GetCompliancePolicyRule()
func (client FleetAppsManagementAdminClient) GetCompliancePolicyRule(ctx context.Context, request GetCompliancePolicyRuleRequest) (response GetCompliancePolicyRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCompliancePolicyRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCompliancePolicyRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCompliancePolicyRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCompliancePolicyRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCompliancePolicyRuleResponse")
	}
	return
}

// getCompliancePolicyRule implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) getCompliancePolicyRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compliancePolicyRules/{compliancePolicyRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCompliancePolicyRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/CompliancePolicyRule/GetCompliancePolicyRule"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "GetCompliancePolicyRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOnboarding Gets a Fleet Application Management Onboarding by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetOnboarding.go.html to see an example of how to use GetOnboarding API.
// A default retry strategy applies to this operation GetOnboarding()
func (client FleetAppsManagementAdminClient) GetOnboarding(ctx context.Context, request GetOnboardingRequest) (response GetOnboardingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOnboarding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOnboardingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOnboardingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOnboardingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOnboardingResponse")
	}
	return
}

// getOnboarding implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) getOnboarding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/Onboardings/{onboardingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOnboardingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Onboarding/GetOnboarding"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "GetOnboarding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPlatformConfiguration Gets a PlatformConfiguration by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetPlatformConfiguration.go.html to see an example of how to use GetPlatformConfiguration API.
// A default retry strategy applies to this operation GetPlatformConfiguration()
func (client FleetAppsManagementAdminClient) GetPlatformConfiguration(ctx context.Context, request GetPlatformConfigurationRequest) (response GetPlatformConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPlatformConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPlatformConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPlatformConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPlatformConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPlatformConfigurationResponse")
	}
	return
}

// getPlatformConfiguration implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) getPlatformConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/platformConfigurations/{platformConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPlatformConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/PlatformConfiguration/GetPlatformConfiguration"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "GetPlatformConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProperty Gets a Property by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetProperty.go.html to see an example of how to use GetProperty API.
// A default retry strategy applies to this operation GetProperty()
func (client FleetAppsManagementAdminClient) GetProperty(ctx context.Context, request GetPropertyRequest) (response GetPropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPropertyResponse")
	}
	return
}

// getProperty implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) getProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/properties/{propertyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Property/GetProperty"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "GetProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCompliancePolicies Gets a list of compliancePolicies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListCompliancePolicies.go.html to see an example of how to use ListCompliancePolicies API.
// A default retry strategy applies to this operation ListCompliancePolicies()
func (client FleetAppsManagementAdminClient) ListCompliancePolicies(ctx context.Context, request ListCompliancePoliciesRequest) (response ListCompliancePoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCompliancePolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCompliancePoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCompliancePoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCompliancePoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCompliancePoliciesResponse")
	}
	return
}

// listCompliancePolicies implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) listCompliancePolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compliancePolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCompliancePoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/CompliancePolicyCollection/ListCompliancePolicies"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "ListCompliancePolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCompliancePolicyRules Gets a list of CompliancePolicyRules.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListCompliancePolicyRules.go.html to see an example of how to use ListCompliancePolicyRules API.
// A default retry strategy applies to this operation ListCompliancePolicyRules()
func (client FleetAppsManagementAdminClient) ListCompliancePolicyRules(ctx context.Context, request ListCompliancePolicyRulesRequest) (response ListCompliancePolicyRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCompliancePolicyRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCompliancePolicyRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCompliancePolicyRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCompliancePolicyRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCompliancePolicyRulesResponse")
	}
	return
}

// listCompliancePolicyRules implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) listCompliancePolicyRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compliancePolicyRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCompliancePolicyRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/CompliancePolicyRuleCollection/ListCompliancePolicyRules"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "ListCompliancePolicyRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOnboardingPolicies Returns a list of onboarding policy information for Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListOnboardingPolicies.go.html to see an example of how to use ListOnboardingPolicies API.
// A default retry strategy applies to this operation ListOnboardingPolicies()
func (client FleetAppsManagementAdminClient) ListOnboardingPolicies(ctx context.Context, request ListOnboardingPoliciesRequest) (response ListOnboardingPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOnboardingPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOnboardingPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOnboardingPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOnboardingPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOnboardingPoliciesResponse")
	}
	return
}

// listOnboardingPolicies implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) listOnboardingPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/OnboardingPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOnboardingPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/OnboardingPolicyCollection/ListOnboardingPolicies"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "ListOnboardingPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOnboardings Returns a list of onboarding information for the Tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListOnboardings.go.html to see an example of how to use ListOnboardings API.
// A default retry strategy applies to this operation ListOnboardings()
func (client FleetAppsManagementAdminClient) ListOnboardings(ctx context.Context, request ListOnboardingsRequest) (response ListOnboardingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOnboardings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOnboardingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOnboardingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOnboardingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOnboardingsResponse")
	}
	return
}

// listOnboardings implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) listOnboardings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/Onboardings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOnboardingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/OnboardingCollection/ListOnboardings"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "ListOnboardings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPlatformConfigurations Returns a list of PlatformConfiguration for Tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListPlatformConfigurations.go.html to see an example of how to use ListPlatformConfigurations API.
// A default retry strategy applies to this operation ListPlatformConfigurations()
func (client FleetAppsManagementAdminClient) ListPlatformConfigurations(ctx context.Context, request ListPlatformConfigurationsRequest) (response ListPlatformConfigurationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPlatformConfigurations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPlatformConfigurationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPlatformConfigurationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPlatformConfigurationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPlatformConfigurationsResponse")
	}
	return
}

// listPlatformConfigurations implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) listPlatformConfigurations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/platformConfigurations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPlatformConfigurationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/PlatformConfigurationCollection/ListPlatformConfigurations"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "ListPlatformConfigurations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProperties List properties and their values for a tenancy in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListProperties.go.html to see an example of how to use ListProperties API.
// A default retry strategy applies to this operation ListProperties()
func (client FleetAppsManagementAdminClient) ListProperties(ctx context.Context, request ListPropertiesRequest) (response ListPropertiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProperties, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPropertiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPropertiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPropertiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPropertiesResponse")
	}
	return
}

// listProperties implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) listProperties(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/properties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPropertiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/PropertyCollection/ListProperties"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "ListProperties", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ManageSettings Updates the Onboarding setting
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ManageSettings.go.html to see an example of how to use ManageSettings API.
// A default retry strategy applies to this operation ManageSettings()
func (client FleetAppsManagementAdminClient) ManageSettings(ctx context.Context, request ManageSettingsRequest) (response ManageSettingsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.manageSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ManageSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ManageSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ManageSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ManageSettingsResponse")
	}
	return
}

// manageSettings implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) manageSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/Onboardings/{onboardingId}/actions/manageSettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ManageSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Onboarding/ManageSettings"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "ManageSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCompliancePolicyRule Updates a CompliancePolicyRule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateCompliancePolicyRule.go.html to see an example of how to use UpdateCompliancePolicyRule API.
// A default retry strategy applies to this operation UpdateCompliancePolicyRule()
func (client FleetAppsManagementAdminClient) UpdateCompliancePolicyRule(ctx context.Context, request UpdateCompliancePolicyRuleRequest) (response UpdateCompliancePolicyRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCompliancePolicyRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCompliancePolicyRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCompliancePolicyRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCompliancePolicyRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCompliancePolicyRuleResponse")
	}
	return
}

// updateCompliancePolicyRule implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) updateCompliancePolicyRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/compliancePolicyRules/{compliancePolicyRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCompliancePolicyRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/CompliancePolicyRule/UpdateCompliancePolicyRule"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "UpdateCompliancePolicyRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOnboarding Updates the Onboarding
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateOnboarding.go.html to see an example of how to use UpdateOnboarding API.
// A default retry strategy applies to this operation UpdateOnboarding()
func (client FleetAppsManagementAdminClient) UpdateOnboarding(ctx context.Context, request UpdateOnboardingRequest) (response UpdateOnboardingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOnboarding, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOnboardingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOnboardingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOnboardingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOnboardingResponse")
	}
	return
}

// updateOnboarding implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) updateOnboarding(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/Onboardings/{onboardingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOnboardingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Onboarding/UpdateOnboarding"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "UpdateOnboarding", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePlatformConfiguration Updates the PlatformConfiguration
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdatePlatformConfiguration.go.html to see an example of how to use UpdatePlatformConfiguration API.
// A default retry strategy applies to this operation UpdatePlatformConfiguration()
func (client FleetAppsManagementAdminClient) UpdatePlatformConfiguration(ctx context.Context, request UpdatePlatformConfigurationRequest) (response UpdatePlatformConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePlatformConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePlatformConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePlatformConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePlatformConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePlatformConfigurationResponse")
	}
	return
}

// updatePlatformConfiguration implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) updatePlatformConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/platformConfigurations/{platformConfigurationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePlatformConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/PlatformConfiguration/UpdatePlatformConfiguration"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "UpdatePlatformConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProperty Update a property in Fleet Application Management.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateProperty.go.html to see an example of how to use UpdateProperty API.
// A default retry strategy applies to this operation UpdateProperty()
func (client FleetAppsManagementAdminClient) UpdateProperty(ctx context.Context, request UpdatePropertyRequest) (response UpdatePropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePropertyResponse")
	}
	return
}

// updateProperty implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementAdminClient) updateProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/properties/{propertyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Property/UpdateProperty"
		err = common.PostProcessServiceError(err, "FleetAppsManagementAdmin", "UpdateProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
