// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
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

// CreateOnboarding Onboard a tenant to Fleet Application Management Service
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

// CreateProperty Creates a new Property.
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

// DeleteProperty Deletes a Property resource by identifier
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

// ListOnboardingPolicies Returns a list of onboarding policy information for FAMS.
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

// ListProperties Returns a list of Properties for Tenancy.
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

// UpdateProperty Updates the Property
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
