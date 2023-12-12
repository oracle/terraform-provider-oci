// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// CloudGuardClient a client for CloudGuard
type CloudGuardClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCloudGuardClientWithConfigurationProvider Creates a new default CloudGuard client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCloudGuardClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CloudGuardClient, err error) {
	if enabled := common.CheckForEnabledServices("cloudguard"); !enabled {
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
	return newCloudGuardClientFromBaseClient(baseClient, provider)
}

// NewCloudGuardClientWithOboToken Creates a new default CloudGuard client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCloudGuardClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CloudGuardClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCloudGuardClientFromBaseClient(baseClient, configProvider)
}

func newCloudGuardClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CloudGuardClient, err error) {
	// CloudGuard service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CloudGuard"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CloudGuardClient{BaseClient: baseClient}
	client.BasePath = "20200131"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CloudGuardClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cloudguard", "https://cloudguard-cp-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CloudGuardClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CloudGuardClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddCompartment Add an existing compartment to a security zone. If you previously removed a subcompartment from a security zone, you can add it back to the same security zone. The security zone ensures that resources in the subcompartment comply with the security zone's policies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/AddCompartment.go.html to see an example of how to use AddCompartment API.
func (client CloudGuardClient) AddCompartment(ctx context.Context, request AddCompartmentRequest) (response AddCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddCompartmentResponse")
	}
	return
}

// addCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) addCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityZones/{securityZoneId}/actions/addCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZone/AddCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "AddCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancels the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
func (client CloudGuardClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "CloudGuard", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDataSourceCompartment Moves the DataSource from current compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ChangeDataSourceCompartment.go.html to see an example of how to use ChangeDataSourceCompartment API.
func (client CloudGuardClient) ChangeDataSourceCompartment(ctx context.Context, request ChangeDataSourceCompartmentRequest) (response ChangeDataSourceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDataSourceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDataSourceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDataSourceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDataSourceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDataSourceCompartmentResponse")
	}
	return
}

// changeDataSourceCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeDataSourceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSources/{dataSourceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDataSourceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataSource/ChangeDataSourceCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "ChangeDataSourceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDetectorRecipeCompartment Moves the DetectorRecipe from current compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ChangeDetectorRecipeCompartment.go.html to see an example of how to use ChangeDetectorRecipeCompartment API.
func (client CloudGuardClient) ChangeDetectorRecipeCompartment(ctx context.Context, request ChangeDetectorRecipeCompartmentRequest) (response ChangeDetectorRecipeCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDetectorRecipeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDetectorRecipeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDetectorRecipeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDetectorRecipeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDetectorRecipeCompartmentResponse")
	}
	return
}

// changeDetectorRecipeCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeDetectorRecipeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/detectorRecipes/{detectorRecipeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDetectorRecipeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipe/ChangeDetectorRecipeCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "ChangeDetectorRecipeCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeManagedListCompartment Moves the ManagedList from current compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ChangeManagedListCompartment.go.html to see an example of how to use ChangeManagedListCompartment API.
func (client CloudGuardClient) ChangeManagedListCompartment(ctx context.Context, request ChangeManagedListCompartmentRequest) (response ChangeManagedListCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeManagedListCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeManagedListCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeManagedListCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeManagedListCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeManagedListCompartmentResponse")
	}
	return
}

// changeManagedListCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeManagedListCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedLists/{managedListId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeManagedListCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ManagedList/ChangeManagedListCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "ChangeManagedListCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeResponderRecipeCompartment Moves the ResponderRecipe from current compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ChangeResponderRecipeCompartment.go.html to see an example of how to use ChangeResponderRecipeCompartment API.
func (client CloudGuardClient) ChangeResponderRecipeCompartment(ctx context.Context, request ChangeResponderRecipeCompartmentRequest) (response ChangeResponderRecipeCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeResponderRecipeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeResponderRecipeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeResponderRecipeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeResponderRecipeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeResponderRecipeCompartmentResponse")
	}
	return
}

// changeResponderRecipeCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeResponderRecipeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderRecipes/{responderRecipeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeResponderRecipeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipe/ChangeResponderRecipeCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "ChangeResponderRecipeCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecurityRecipeCompartment Moves a security zone recipe to a different compartment. When provided, `If-Match` is checked against `ETag` values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ChangeSecurityRecipeCompartment.go.html to see an example of how to use ChangeSecurityRecipeCompartment API.
func (client CloudGuardClient) ChangeSecurityRecipeCompartment(ctx context.Context, request ChangeSecurityRecipeCompartmentRequest) (response ChangeSecurityRecipeCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSecurityRecipeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecurityRecipeCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecurityRecipeCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecurityRecipeCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecurityRecipeCompartmentResponse")
	}
	return
}

// changeSecurityRecipeCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeSecurityRecipeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityRecipes/{securityRecipeId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecurityRecipeCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityRecipe/ChangeSecurityRecipeCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "ChangeSecurityRecipeCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecurityZoneCompartment Moves a security zone to a different compartment. When provided, `If-Match` is checked against `ETag` values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ChangeSecurityZoneCompartment.go.html to see an example of how to use ChangeSecurityZoneCompartment API.
func (client CloudGuardClient) ChangeSecurityZoneCompartment(ctx context.Context, request ChangeSecurityZoneCompartmentRequest) (response ChangeSecurityZoneCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSecurityZoneCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecurityZoneCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecurityZoneCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecurityZoneCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecurityZoneCompartmentResponse")
	}
	return
}

// changeSecurityZoneCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) changeSecurityZoneCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityZones/{securityZoneId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecurityZoneCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZone/ChangeSecurityZoneCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "ChangeSecurityZoneCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataMaskRule Creates a new Data Mask Rule Definition
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateDataMaskRule.go.html to see an example of how to use CreateDataMaskRule API.
func (client CloudGuardClient) CreateDataMaskRule(ctx context.Context, request CreateDataMaskRuleRequest) (response CreateDataMaskRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDataMaskRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataMaskRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataMaskRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataMaskRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataMaskRuleResponse")
	}
	return
}

// createDataMaskRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createDataMaskRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataMaskRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataMaskRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataMaskRule/CreateDataMaskRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateDataMaskRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataSource Creates a DataSource
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateDataSource.go.html to see an example of how to use CreateDataSource API.
func (client CloudGuardClient) CreateDataSource(ctx context.Context, request CreateDataSourceRequest) (response CreateDataSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDataSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataSourceResponse")
	}
	return
}

// createDataSource implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createDataSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dataSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataSource/CreateDataSource"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateDataSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDetectorRecipe Creates a DetectorRecipe
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateDetectorRecipe.go.html to see an example of how to use CreateDetectorRecipe API.
func (client CloudGuardClient) CreateDetectorRecipe(ctx context.Context, request CreateDetectorRecipeRequest) (response CreateDetectorRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDetectorRecipeResponse")
	}
	return
}

// createDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/detectorRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipe/CreateDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDetectorRecipeDetectorRule Create the DetectorRule
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateDetectorRecipeDetectorRule.go.html to see an example of how to use CreateDetectorRecipeDetectorRule API.
func (client CloudGuardClient) CreateDetectorRecipeDetectorRule(ctx context.Context, request CreateDetectorRecipeDetectorRuleRequest) (response CreateDetectorRecipeDetectorRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDetectorRecipeDetectorRuleResponse")
	}
	return
}

// createDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/detectorRecipes/{detectorRecipeId}/detectorRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipeDetectorRule/CreateDetectorRecipeDetectorRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateDetectorRecipeDetectorRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateManagedList Creates a new ManagedList.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateManagedList.go.html to see an example of how to use CreateManagedList API.
func (client CloudGuardClient) CreateManagedList(ctx context.Context, request CreateManagedListRequest) (response CreateManagedListResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagedListResponse")
	}
	return
}

// createManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createManagedList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ManagedList/CreateManagedList"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateManagedList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateResponderRecipe Create a ResponderRecipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateResponderRecipe.go.html to see an example of how to use CreateResponderRecipe API.
func (client CloudGuardClient) CreateResponderRecipe(ctx context.Context, request CreateResponderRecipeRequest) (response CreateResponderRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateResponderRecipeResponse")
	}
	return
}

// createResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipe/CreateResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityRecipe Creates a security zone recipe. A security zone recipe is a collection of security zone policies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateSecurityRecipe.go.html to see an example of how to use CreateSecurityRecipe API.
func (client CloudGuardClient) CreateSecurityRecipe(ctx context.Context, request CreateSecurityRecipeRequest) (response CreateSecurityRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSecurityRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityRecipeResponse")
	}
	return
}

// createSecurityRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createSecurityRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityRecipe/CreateSecurityRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateSecurityRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityZone Creates a security zone for a compartment. A security zone enforces all security zone policies in a given security zone recipe. Any actions that violate a policy are denied. By default, any subcompartments are also in the same security zone.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateSecurityZone.go.html to see an example of how to use CreateSecurityZone API.
func (client CloudGuardClient) CreateSecurityZone(ctx context.Context, request CreateSecurityZoneRequest) (response CreateSecurityZoneResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSecurityZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityZoneResponse")
	}
	return
}

// createSecurityZone implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createSecurityZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityZones", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZone/CreateSecurityZone"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateSecurityZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTarget Creates a new Target
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateTarget.go.html to see an example of how to use CreateTarget API.
func (client CloudGuardClient) CreateTarget(ctx context.Context, request CreateTargetRequest) (response CreateTargetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetResponse")
	}
	return
}

// createTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Target/CreateTarget"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTargetDetectorRecipe Attach a DetectorRecipe with the Target
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateTargetDetectorRecipe.go.html to see an example of how to use CreateTargetDetectorRecipe API.
func (client CloudGuardClient) CreateTargetDetectorRecipe(ctx context.Context, request CreateTargetDetectorRecipeRequest) (response CreateTargetDetectorRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetDetectorRecipeResponse")
	}
	return
}

// createTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createTargetDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targets/{targetId}/targetDetectorRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipe/CreateTargetDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateTargetDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTargetResponderRecipe Attach a ResponderRecipe with the Target
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/CreateTargetResponderRecipe.go.html to see an example of how to use CreateTargetResponderRecipe API.
func (client CloudGuardClient) CreateTargetResponderRecipe(ctx context.Context, request CreateTargetResponderRecipeRequest) (response CreateTargetResponderRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTargetResponderRecipeResponse")
	}
	return
}

// createTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) createTargetResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/targets/{targetId}/targetResponderRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipe/CreateTargetResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "CreateTargetResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataMaskRule Deletes a DataMaskRule identified by dataMaskRuleId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteDataMaskRule.go.html to see an example of how to use DeleteDataMaskRule API.
func (client CloudGuardClient) DeleteDataMaskRule(ctx context.Context, request DeleteDataMaskRuleRequest) (response DeleteDataMaskRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataMaskRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataMaskRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataMaskRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataMaskRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataMaskRuleResponse")
	}
	return
}

// deleteDataMaskRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteDataMaskRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dataMaskRules/{dataMaskRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataMaskRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataMaskRule/DeleteDataMaskRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteDataMaskRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataSource Deletes a DataSource identified by dataSourceId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteDataSource.go.html to see an example of how to use DeleteDataSource API.
func (client CloudGuardClient) DeleteDataSource(ctx context.Context, request DeleteDataSourceRequest) (response DeleteDataSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteDataSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataSourceResponse")
	}
	return
}

// deleteDataSource implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteDataSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dataSources/{dataSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataSource/DeleteDataSource"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteDataSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDetectorRecipe Deletes a DetectorRecipe identified by detectorRecipeId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteDetectorRecipe.go.html to see an example of how to use DeleteDetectorRecipe API.
func (client CloudGuardClient) DeleteDetectorRecipe(ctx context.Context, request DeleteDetectorRecipeRequest) (response DeleteDetectorRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDetectorRecipeResponse")
	}
	return
}

// deleteDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/detectorRecipes/{detectorRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipe/DeleteDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDetectorRecipeDetectorRule Deletes DetectorRecipeDetectorRule
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteDetectorRecipeDetectorRule.go.html to see an example of how to use DeleteDetectorRecipeDetectorRule API.
func (client CloudGuardClient) DeleteDetectorRecipeDetectorRule(ctx context.Context, request DeleteDetectorRecipeDetectorRuleRequest) (response DeleteDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDetectorRecipeDetectorRuleResponse")
	}
	return
}

// deleteDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/detectorRecipes/{detectorRecipeId}/detectorRules/{detectorRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipeDetectorRule/DeleteDetectorRecipeDetectorRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteDetectorRecipeDetectorRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDetectorRecipeDetectorRuleDataSource Delete the DetectorRecipeDetectorRuleDataSource resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteDetectorRecipeDetectorRuleDataSource.go.html to see an example of how to use DeleteDetectorRecipeDetectorRuleDataSource API.
func (client CloudGuardClient) DeleteDetectorRecipeDetectorRuleDataSource(ctx context.Context, request DeleteDetectorRecipeDetectorRuleDataSourceRequest) (response DeleteDetectorRecipeDetectorRuleDataSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDetectorRecipeDetectorRuleDataSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDetectorRecipeDetectorRuleDataSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDetectorRecipeDetectorRuleDataSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDetectorRecipeDetectorRuleDataSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDetectorRecipeDetectorRuleDataSourceResponse")
	}
	return
}

// deleteDetectorRecipeDetectorRuleDataSource implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteDetectorRecipeDetectorRuleDataSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/detectorRecipes/{detectorRecipeId}/detectorRules/{detectorRuleId}/dataSources/{dataSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDetectorRecipeDetectorRuleDataSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipeDetectorRule/DeleteDetectorRecipeDetectorRuleDataSource"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteDetectorRecipeDetectorRuleDataSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagedList Deletes a managed list identified by managedListId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteManagedList.go.html to see an example of how to use DeleteManagedList API.
func (client CloudGuardClient) DeleteManagedList(ctx context.Context, request DeleteManagedListRequest) (response DeleteManagedListResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagedListResponse")
	}
	return
}

// deleteManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteManagedList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedLists/{managedListId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ManagedList/DeleteManagedList"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteManagedList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteResponderRecipe Delete the ResponderRecipe resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteResponderRecipe.go.html to see an example of how to use DeleteResponderRecipe API.
func (client CloudGuardClient) DeleteResponderRecipe(ctx context.Context, request DeleteResponderRecipeRequest) (response DeleteResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteResponderRecipeResponse")
	}
	return
}

// deleteResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/responderRecipes/{responderRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipe/DeleteResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSecurityRecipe Deletes a security zone recipe. The recipe can't be associated with an existing security zone.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteSecurityRecipe.go.html to see an example of how to use DeleteSecurityRecipe API.
func (client CloudGuardClient) DeleteSecurityRecipe(ctx context.Context, request DeleteSecurityRecipeRequest) (response DeleteSecurityRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityRecipeResponse")
	}
	return
}

// deleteSecurityRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteSecurityRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/securityRecipes/{securityRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityRecipe/DeleteSecurityRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteSecurityRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSecurityZone Deletes an existing security zone with a given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteSecurityZone.go.html to see an example of how to use DeleteSecurityZone API.
func (client CloudGuardClient) DeleteSecurityZone(ctx context.Context, request DeleteSecurityZoneRequest) (response DeleteSecurityZoneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityZoneResponse")
	}
	return
}

// deleteSecurityZone implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteSecurityZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/securityZones/{securityZoneId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZone/DeleteSecurityZone"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteSecurityZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTarget Deletes a Target identified by targetId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteTarget.go.html to see an example of how to use DeleteTarget API.
func (client CloudGuardClient) DeleteTarget(ctx context.Context, request DeleteTargetRequest) (response DeleteTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetResponse")
	}
	return
}

// deleteTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targets/{targetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Target/DeleteTarget"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTargetDetectorRecipe Delete the TargetDetectorRecipe resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteTargetDetectorRecipe.go.html to see an example of how to use DeleteTargetDetectorRecipe API.
func (client CloudGuardClient) DeleteTargetDetectorRecipe(ctx context.Context, request DeleteTargetDetectorRecipeRequest) (response DeleteTargetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetDetectorRecipeResponse")
	}
	return
}

// deleteTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteTargetDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipe/DeleteTargetDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteTargetDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTargetResponderRecipe Delete the TargetResponderRecipe resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/DeleteTargetResponderRecipe.go.html to see an example of how to use DeleteTargetResponderRecipe API.
func (client CloudGuardClient) DeleteTargetResponderRecipe(ctx context.Context, request DeleteTargetResponderRecipeRequest) (response DeleteTargetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTargetResponderRecipeResponse")
	}
	return
}

// deleteTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) deleteTargetResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipe/DeleteTargetResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "DeleteTargetResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExecuteResponderExecution Executes the responder execution. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ExecuteResponderExecution.go.html to see an example of how to use ExecuteResponderExecution API.
func (client CloudGuardClient) ExecuteResponderExecution(ctx context.Context, request ExecuteResponderExecutionRequest) (response ExecuteResponderExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.executeResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExecuteResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExecuteResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExecuteResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExecuteResponderExecutionResponse")
	}
	return
}

// executeResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) executeResponderExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/{responderExecutionId}/actions/execute", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExecuteResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderExecution/ExecuteResponderExecution"
		err = common.PostProcessServiceError(err, "CloudGuard", "ExecuteResponderExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConditionMetadataType Returns ConditionType with its details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetConditionMetadataType.go.html to see an example of how to use GetConditionMetadataType API.
func (client CloudGuardClient) GetConditionMetadataType(ctx context.Context, request GetConditionMetadataTypeRequest) (response GetConditionMetadataTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConditionMetadataType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConditionMetadataTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConditionMetadataTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConditionMetadataTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConditionMetadataTypeResponse")
	}
	return
}

// getConditionMetadataType implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getConditionMetadataType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/conditionMetadataTypes/{conditionMetadataTypeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConditionMetadataTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ConditionMetadataType/GetConditionMetadataType"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetConditionMetadataType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfiguration GET Cloud Guard Configuration Details for a Tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetConfiguration.go.html to see an example of how to use GetConfiguration API.
func (client CloudGuardClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationResponse")
	}
	return
}

// getConfiguration implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Configuration/GetConfiguration"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataMaskRule Returns a DataMaskRule identified by DataMaskRuleId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetDataMaskRule.go.html to see an example of how to use GetDataMaskRule API.
func (client CloudGuardClient) GetDataMaskRule(ctx context.Context, request GetDataMaskRuleRequest) (response GetDataMaskRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataMaskRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataMaskRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataMaskRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataMaskRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataMaskRuleResponse")
	}
	return
}

// getDataMaskRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDataMaskRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataMaskRules/{dataMaskRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataMaskRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataMaskRule/GetDataMaskRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetDataMaskRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataSource Returns a DataSource identified by dataSourceId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetDataSource.go.html to see an example of how to use GetDataSource API.
func (client CloudGuardClient) GetDataSource(ctx context.Context, request GetDataSourceRequest) (response GetDataSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataSourceResponse")
	}
	return
}

// getDataSource implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDataSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSources/{dataSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataSource/GetDataSource"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetDataSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetector Returns a Detector identified by detectorId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetDetector.go.html to see an example of how to use GetDetector API.
func (client CloudGuardClient) GetDetector(ctx context.Context, request GetDetectorRequest) (response GetDetectorResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetector, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorResponse")
	}
	return
}

// getDetector implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetector(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors/{detectorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDetectorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Detector/GetDetector"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetDetector", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetectorRecipe Returns a DetectorRecipe identified by detectorRecipeId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetDetectorRecipe.go.html to see an example of how to use GetDetectorRecipe API.
func (client CloudGuardClient) GetDetectorRecipe(ctx context.Context, request GetDetectorRecipeRequest) (response GetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorRecipeResponse")
	}
	return
}

// getDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes/{detectorRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipe/GetDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetectorRecipeDetectorRule Get DetectorRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetDetectorRecipeDetectorRule.go.html to see an example of how to use GetDetectorRecipeDetectorRule API.
func (client CloudGuardClient) GetDetectorRecipeDetectorRule(ctx context.Context, request GetDetectorRecipeDetectorRuleRequest) (response GetDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorRecipeDetectorRuleResponse")
	}
	return
}

// getDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes/{detectorRecipeId}/detectorRules/{detectorRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipeDetectorRule/GetDetectorRecipeDetectorRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetDetectorRecipeDetectorRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDetectorRule Returns a Detector Rule identified by detectorRuleId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetDetectorRule.go.html to see an example of how to use GetDetectorRule API.
func (client CloudGuardClient) GetDetectorRule(ctx context.Context, request GetDetectorRuleRequest) (response GetDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDetectorRuleResponse")
	}
	return
}

// getDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getDetectorRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors/{detectorId}/detectorRules/{detectorRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRule/GetDetectorRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetDetectorRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedList Returns a managed list identified by managedListId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetManagedList.go.html to see an example of how to use GetManagedList API.
func (client CloudGuardClient) GetManagedList(ctx context.Context, request GetManagedListRequest) (response GetManagedListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedListResponse")
	}
	return
}

// getManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getManagedList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedLists/{managedListId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ManagedList/GetManagedList"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetManagedList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProblem Returns a Problems response
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetProblem.go.html to see an example of how to use GetProblem API.
func (client CloudGuardClient) GetProblem(ctx context.Context, request GetProblemRequest) (response GetProblemResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProblem, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProblemResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProblemResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProblemResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProblemResponse")
	}
	return
}

// getProblem implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getProblem(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProblemResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Problem/GetProblem"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetProblem", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResourceProfile Returns resource profile details
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetResourceProfile.go.html to see an example of how to use GetResourceProfile API.
func (client CloudGuardClient) GetResourceProfile(ctx context.Context, request GetResourceProfileRequest) (response GetResourceProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResourceProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResourceProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResourceProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResourceProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResourceProfileResponse")
	}
	return
}

// getResourceProfile implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResourceProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceProfiles/{resourceProfileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResourceProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResourceProfile/GetResourceProfile"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetResourceProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderExecution Returns a Responder Execution identified by responderExecutionId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetResponderExecution.go.html to see an example of how to use GetResponderExecution API.
func (client CloudGuardClient) GetResponderExecution(ctx context.Context, request GetResponderExecutionRequest) (response GetResponderExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderExecutionResponse")
	}
	return
}

// getResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderExecutions/{responderExecutionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderExecution/GetResponderExecution"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetResponderExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderRecipe Get a ResponderRecipe by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetResponderRecipe.go.html to see an example of how to use GetResponderRecipe API.
func (client CloudGuardClient) GetResponderRecipe(ctx context.Context, request GetResponderRecipeRequest) (response GetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderRecipeResponse")
	}
	return
}

// getResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes/{responderRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipe/GetResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderRecipeResponderRule Get ResponderRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetResponderRecipeResponderRule.go.html to see an example of how to use GetResponderRecipeResponderRule API.
func (client CloudGuardClient) GetResponderRecipeResponderRule(ctx context.Context, request GetResponderRecipeResponderRuleRequest) (response GetResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderRecipeResponderRuleResponse")
	}
	return
}

// getResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes/{responderRecipeId}/responderRules/{responderRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipeResponderRule/GetResponderRecipeResponderRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetResponderRecipeResponderRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResponderRule Get a ResponderRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetResponderRule.go.html to see an example of how to use GetResponderRule API.
func (client CloudGuardClient) GetResponderRule(ctx context.Context, request GetResponderRuleRequest) (response GetResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResponderRuleResponse")
	}
	return
}

// getResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getResponderRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRules/{responderRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRule/GetResponderRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetResponderRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityPolicy Gets a security zone policy using its identifier. When a policy is enabled in a security zone, then any action in the zone that attempts to violate that policy is denied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetSecurityPolicy.go.html to see an example of how to use GetSecurityPolicy API.
func (client CloudGuardClient) GetSecurityPolicy(ctx context.Context, request GetSecurityPolicyRequest) (response GetSecurityPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityPolicyResponse")
	}
	return
}

// getSecurityPolicy implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getSecurityPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicies/{securityPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityPolicy/GetSecurityPolicy"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetSecurityPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityRecipe Gets a security zone recipe by identifier. A security zone recipe is a collection of security zone policies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetSecurityRecipe.go.html to see an example of how to use GetSecurityRecipe API.
func (client CloudGuardClient) GetSecurityRecipe(ctx context.Context, request GetSecurityRecipeRequest) (response GetSecurityRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityRecipeResponse")
	}
	return
}

// getSecurityRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getSecurityRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityRecipes/{securityRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityRecipe/GetSecurityRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetSecurityRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityZone Gets a security zone by its identifier. A security zone is associated with a security zone recipe and enforces all security zone policies in the recipe. Any actions in the zone's compartments that violate a policy are denied.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetSecurityZone.go.html to see an example of how to use GetSecurityZone API.
func (client CloudGuardClient) GetSecurityZone(ctx context.Context, request GetSecurityZoneRequest) (response GetSecurityZoneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityZoneResponse")
	}
	return
}

// getSecurityZone implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getSecurityZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityZones/{securityZoneId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZone/GetSecurityZone"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetSecurityZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSighting Returns Sighting details
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetSighting.go.html to see an example of how to use GetSighting API.
func (client CloudGuardClient) GetSighting(ctx context.Context, request GetSightingRequest) (response GetSightingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSighting, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSightingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSightingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSightingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSightingResponse")
	}
	return
}

// getSighting implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getSighting(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sightings/{sightingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSightingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Sighting/GetSighting"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetSighting", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTarget Returns a Target identified by targetId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetTarget.go.html to see an example of how to use GetTarget API.
func (client CloudGuardClient) GetTarget(ctx context.Context, request GetTargetRequest) (response GetTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetResponse")
	}
	return
}

// getTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Target/GetTarget"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetDetectorRecipe Get a TargetDetectorRecipe by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetTargetDetectorRecipe.go.html to see an example of how to use GetTargetDetectorRecipe API.
func (client CloudGuardClient) GetTargetDetectorRecipe(ctx context.Context, request GetTargetDetectorRecipeRequest) (response GetTargetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetDetectorRecipeResponse")
	}
	return
}

// getTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipe/GetTargetDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetTargetDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetDetectorRecipeDetectorRule Get DetectorRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetTargetDetectorRecipeDetectorRule.go.html to see an example of how to use GetTargetDetectorRecipeDetectorRule API.
func (client CloudGuardClient) GetTargetDetectorRecipeDetectorRule(ctx context.Context, request GetTargetDetectorRecipeDetectorRuleRequest) (response GetTargetDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetDetectorRecipeDetectorRuleResponse")
	}
	return
}

// getTargetDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}/detectorRules/{detectorRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipeDetectorRule/GetTargetDetectorRecipeDetectorRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetTargetDetectorRecipeDetectorRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetResponderRecipe Get a TargetResponderRecipe by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetTargetResponderRecipe.go.html to see an example of how to use GetTargetResponderRecipe API.
func (client CloudGuardClient) GetTargetResponderRecipe(ctx context.Context, request GetTargetResponderRecipeRequest) (response GetTargetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetResponderRecipeResponse")
	}
	return
}

// getTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipe/GetTargetResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetTargetResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTargetResponderRecipeResponderRule Get ResponderRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetTargetResponderRecipeResponderRule.go.html to see an example of how to use GetTargetResponderRecipeResponderRule API.
func (client CloudGuardClient) GetTargetResponderRecipeResponderRule(ctx context.Context, request GetTargetResponderRecipeResponderRuleRequest) (response GetTargetResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTargetResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTargetResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTargetResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTargetResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTargetResponderRecipeResponderRuleResponse")
	}
	return
}

// getTargetResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getTargetResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}/responderRules/{responderRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTargetResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipeResponderRule/GetTargetResponderRecipeResponderRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetTargetResponderRecipeResponderRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets details of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client CloudGuardClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "CloudGuard", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConditionMetadataTypes Returns a list of condition types.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListConditionMetadataTypes.go.html to see an example of how to use ListConditionMetadataTypes API.
func (client CloudGuardClient) ListConditionMetadataTypes(ctx context.Context, request ListConditionMetadataTypesRequest) (response ListConditionMetadataTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConditionMetadataTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConditionMetadataTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConditionMetadataTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConditionMetadataTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConditionMetadataTypesResponse")
	}
	return
}

// listConditionMetadataTypes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listConditionMetadataTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/conditionMetadataTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConditionMetadataTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ConditionMetadataType/ListConditionMetadataTypes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListConditionMetadataTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataMaskRules Returns a list of all Data Mask Rules in the root 'compartmentId' passed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDataMaskRules.go.html to see an example of how to use ListDataMaskRules API.
func (client CloudGuardClient) ListDataMaskRules(ctx context.Context, request ListDataMaskRulesRequest) (response ListDataMaskRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataMaskRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataMaskRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataMaskRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataMaskRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataMaskRulesResponse")
	}
	return
}

// listDataMaskRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDataMaskRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataMaskRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataMaskRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataMaskRule/ListDataMaskRules"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListDataMaskRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataSourceEvents Returns a list of events from CloudGuard DataSource
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDataSourceEvents.go.html to see an example of how to use ListDataSourceEvents API.
func (client CloudGuardClient) ListDataSourceEvents(ctx context.Context, request ListDataSourceEventsRequest) (response ListDataSourceEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataSourceEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataSourceEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataSourceEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataSourceEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataSourceEventsResponse")
	}
	return
}

// listDataSourceEvents implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDataSourceEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSources/{dataSourceId}/events", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataSourceEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataSource/ListDataSourceEvents"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListDataSourceEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataSources Returns a list of all Data Sources in a compartment
// The ListDataSources operation returns only the data Sources in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListdataSources on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDataSources.go.html to see an example of how to use ListDataSources API.
func (client CloudGuardClient) ListDataSources(ctx context.Context, request ListDataSourcesRequest) (response ListDataSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataSourcesResponse")
	}
	return
}

// listDataSources implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDataSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dataSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataSource/ListDataSources"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListDataSources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectorRecipeDetectorRules Returns a list of DetectorRule associated with DetectorRecipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectorRecipeDetectorRules.go.html to see an example of how to use ListDetectorRecipeDetectorRules API.
func (client CloudGuardClient) ListDetectorRecipeDetectorRules(ctx context.Context, request ListDetectorRecipeDetectorRulesRequest) (response ListDetectorRecipeDetectorRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectorRecipeDetectorRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorRecipeDetectorRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorRecipeDetectorRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorRecipeDetectorRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorRecipeDetectorRulesResponse")
	}
	return
}

// listDetectorRecipeDetectorRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectorRecipeDetectorRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes/{detectorRecipeId}/detectorRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDetectorRecipeDetectorRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipeDetectorRule/ListDetectorRecipeDetectorRules"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListDetectorRecipeDetectorRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectorRecipes Returns a list of all Detector Recipes in a compartment
// The ListDetectorRecipes operation returns only the detector recipes in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListDetectorRecipes on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectorRecipes.go.html to see an example of how to use ListDetectorRecipes API.
func (client CloudGuardClient) ListDetectorRecipes(ctx context.Context, request ListDetectorRecipesRequest) (response ListDetectorRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectorRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorRecipesResponse")
	}
	return
}

// listDetectorRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectorRecipes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectorRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDetectorRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipe/ListDetectorRecipes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListDetectorRecipes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectorRules Returns a list of detector rules for the detectorId passed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectorRules.go.html to see an example of how to use ListDetectorRules API.
func (client CloudGuardClient) ListDetectorRules(ctx context.Context, request ListDetectorRulesRequest) (response ListDetectorRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectorRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorRulesResponse")
	}
	return
}

// listDetectorRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectorRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors/{detectorId}/detectorRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDetectorRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRule/ListDetectorRules"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListDetectorRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDetectors Returns detector catalog - list of detectors supported by Cloud Guard
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectors.go.html to see an example of how to use ListDetectors API.
func (client CloudGuardClient) ListDetectors(ctx context.Context, request ListDetectorsRequest) (response ListDetectorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDetectors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDetectorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDetectorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDetectorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDetectorsResponse")
	}
	return
}

// listDetectors implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listDetectors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/detectors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDetectorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Detector/ListDetectors"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListDetectors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListImpactedResources Returns a list of Impacted Resources for a CloudGuard Problem
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListImpactedResources.go.html to see an example of how to use ListImpactedResources API.
func (client CloudGuardClient) ListImpactedResources(ctx context.Context, request ListImpactedResourcesRequest) (response ListImpactedResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listImpactedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListImpactedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListImpactedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListImpactedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListImpactedResourcesResponse")
	}
	return
}

// listImpactedResources implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listImpactedResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/impactedResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListImpactedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ImpactedResourceSummary/ListImpactedResources"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListImpactedResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedListTypes Returns all ManagedList types supported by Cloud Guard
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListManagedListTypes.go.html to see an example of how to use ListManagedListTypes API.
func (client CloudGuardClient) ListManagedListTypes(ctx context.Context, request ListManagedListTypesRequest) (response ListManagedListTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedListTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedListTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedListTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedListTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedListTypesResponse")
	}
	return
}

// listManagedListTypes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listManagedListTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedListTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedListTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ManagedListTypeSummary/ListManagedListTypes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListManagedListTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedLists Returns a list of ListManagedLists.
// The ListManagedLists operation returns only the managed lists in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return ManagedLists in only
// those compartments for which the requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListManagedLists on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListManagedLists.go.html to see an example of how to use ListManagedLists API.
func (client CloudGuardClient) ListManagedLists(ctx context.Context, request ListManagedListsRequest) (response ListManagedListsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedListsResponse")
	}
	return
}

// listManagedLists implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listManagedLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ManagedList/ListManagedLists"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListManagedLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPolicies Returns the list of global policy statements needed by Cloud Guard when enabling
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListPolicies.go.html to see an example of how to use ListPolicies API.
func (client CloudGuardClient) ListPolicies(ctx context.Context, request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPoliciesResponse")
	}
	return
}

// listPolicies implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/policies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/PolicySummary/ListPolicies"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProblemEndpoints Returns a list of endpoints associated with a cloud guard problem
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListProblemEndpoints.go.html to see an example of how to use ListProblemEndpoints API.
func (client CloudGuardClient) ListProblemEndpoints(ctx context.Context, request ListProblemEndpointsRequest) (response ListProblemEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProblemEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProblemEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProblemEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProblemEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProblemEndpointsResponse")
	}
	return
}

// listProblemEndpoints implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listProblemEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/endpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProblemEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ProblemEndpointSummary/ListProblemEndpoints"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListProblemEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProblemEntities Returns a list of entities for a CloudGuard Problem
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListProblemEntities.go.html to see an example of how to use ListProblemEntities API.
func (client CloudGuardClient) ListProblemEntities(ctx context.Context, request ListProblemEntitiesRequest) (response ListProblemEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProblemEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProblemEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProblemEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProblemEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProblemEntitiesResponse")
	}
	return
}

// listProblemEntities implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listProblemEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/entities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProblemEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Problem/ListProblemEntities"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListProblemEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProblemHistories Returns a list of Actions done on CloudGuard Problem
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListProblemHistories.go.html to see an example of how to use ListProblemHistories API.
func (client CloudGuardClient) ListProblemHistories(ctx context.Context, request ListProblemHistoriesRequest) (response ListProblemHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProblemHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProblemHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProblemHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProblemHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProblemHistoriesResponse")
	}
	return
}

// listProblemHistories implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listProblemHistories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/histories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProblemHistoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Problem/ListProblemHistories"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListProblemHistories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProblems Returns a list of all Problems identified by the Cloud Guard
// The ListProblems operation returns only the problems in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListProblems on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListProblems.go.html to see an example of how to use ListProblems API.
func (client CloudGuardClient) ListProblems(ctx context.Context, request ListProblemsRequest) (response ListProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProblemsResponse")
	}
	return
}

// listProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listProblems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Problem/ListProblems"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListProblems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecommendations Returns a list of all Recommendations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListRecommendations.go.html to see an example of how to use ListRecommendations API.
func (client CloudGuardClient) ListRecommendations(ctx context.Context, request ListRecommendationsRequest) (response ListRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecommendationsResponse")
	}
	return
}

// listRecommendations implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listRecommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recommendations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRecommendationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/RecommendationSummary/ListRecommendations"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListRecommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceProfileEndpoints Returns a list of endpoints for Cloud Guard resource profile
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceProfileEndpoints.go.html to see an example of how to use ListResourceProfileEndpoints API.
func (client CloudGuardClient) ListResourceProfileEndpoints(ctx context.Context, request ListResourceProfileEndpointsRequest) (response ListResourceProfileEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceProfileEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceProfileEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceProfileEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceProfileEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceProfileEndpointsResponse")
	}
	return
}

// listResourceProfileEndpoints implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResourceProfileEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceProfiles/{resourceProfileId}/endpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceProfileEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResourceProfileEndpointSummary/ListResourceProfileEndpoints"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResourceProfileEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceProfileImpactedResources Returns a list of impacted resources for Cloud Guard resource profile
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceProfileImpactedResources.go.html to see an example of how to use ListResourceProfileImpactedResources API.
func (client CloudGuardClient) ListResourceProfileImpactedResources(ctx context.Context, request ListResourceProfileImpactedResourcesRequest) (response ListResourceProfileImpactedResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceProfileImpactedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceProfileImpactedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceProfileImpactedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceProfileImpactedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceProfileImpactedResourcesResponse")
	}
	return
}

// listResourceProfileImpactedResources implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResourceProfileImpactedResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceProfiles/{resourceProfileId}/impactedResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceProfileImpactedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResourceProfileImpactedResourceSummary/ListResourceProfileImpactedResources"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResourceProfileImpactedResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceProfiles Returns a list of all resource profiles identified by the Cloud Guard
// The ListResourceProfiles operation returns only resource profiles that match the passed filters.
// The ListResourceProfiles operation returns only the resource profiles in `compartmentId` passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListResourceProfiles on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceProfiles.go.html to see an example of how to use ListResourceProfiles API.
func (client CloudGuardClient) ListResourceProfiles(ctx context.Context, request ListResourceProfilesRequest) (response ListResourceProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceProfilesResponse")
	}
	return
}

// listResourceProfiles implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResourceProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResourceProfileSummary/ListResourceProfiles"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResourceProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceTypes Returns a list of resource types.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceTypes.go.html to see an example of how to use ListResourceTypes API.
func (client CloudGuardClient) ListResourceTypes(ctx context.Context, request ListResourceTypesRequest) (response ListResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceTypesResponse")
	}
	return
}

// listResourceTypes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResourceTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResourceTypeSummary/ListResourceTypes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResourceTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderActivities Returns a list of Responder activities done on CloudGuard Problem
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderActivities.go.html to see an example of how to use ListResponderActivities API.
func (client CloudGuardClient) ListResponderActivities(ctx context.Context, request ListResponderActivitiesRequest) (response ListResponderActivitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderActivities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderActivitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderActivitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderActivitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderActivitiesResponse")
	}
	return
}

// listResponderActivities implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderActivities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/problems/{problemId}/responderActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResponderActivitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderActivitySummary/ListResponderActivities"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResponderActivities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderExecutions Returns a list of Responder Executions. A Responder Execution is an entity that tracks the collective execution of multiple Responder Rule Executions for a given Problem.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderExecutions.go.html to see an example of how to use ListResponderExecutions API.
func (client CloudGuardClient) ListResponderExecutions(ctx context.Context, request ListResponderExecutionsRequest) (response ListResponderExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderExecutionsResponse")
	}
	return
}

// listResponderExecutions implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderExecutions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResponderExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderExecutionSummary/ListResponderExecutions"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResponderExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderRecipeResponderRules Returns a list of ResponderRule associated with ResponderRecipe.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderRecipeResponderRules.go.html to see an example of how to use ListResponderRecipeResponderRules API.
func (client CloudGuardClient) ListResponderRecipeResponderRules(ctx context.Context, request ListResponderRecipeResponderRulesRequest) (response ListResponderRecipeResponderRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderRecipeResponderRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderRecipeResponderRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderRecipeResponderRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderRecipeResponderRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderRecipeResponderRulesResponse")
	}
	return
}

// listResponderRecipeResponderRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderRecipeResponderRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes/{responderRecipeId}/responderRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResponderRecipeResponderRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipeResponderRule/ListResponderRecipeResponderRules"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResponderRecipeResponderRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderRecipes Returns a list of all ResponderRecipes in a compartment
// The ListResponderRecipe operation returns only the targets in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListResponderRecipe on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderRecipes.go.html to see an example of how to use ListResponderRecipes API.
func (client CloudGuardClient) ListResponderRecipes(ctx context.Context, request ListResponderRecipesRequest) (response ListResponderRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderRecipesResponse")
	}
	return
}

// listResponderRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderRecipes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResponderRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipe/ListResponderRecipes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResponderRecipes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResponderRules Returns a list of ResponderRule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderRules.go.html to see an example of how to use ListResponderRules API.
func (client CloudGuardClient) ListResponderRules(ctx context.Context, request ListResponderRulesRequest) (response ListResponderRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResponderRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResponderRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResponderRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResponderRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResponderRulesResponse")
	}
	return
}

// listResponderRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listResponderRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/responderRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResponderRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRule/ListResponderRules"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListResponderRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityPolicies Returns a list of security zone policies. Specify any compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSecurityPolicies.go.html to see an example of how to use ListSecurityPolicies API.
func (client CloudGuardClient) ListSecurityPolicies(ctx context.Context, request ListSecurityPoliciesRequest) (response ListSecurityPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityPoliciesResponse")
	}
	return
}

// listSecurityPolicies implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listSecurityPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityPolicyCollection/ListSecurityPolicies"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListSecurityPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityRecipes Gets a list of all security zone recipes in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSecurityRecipes.go.html to see an example of how to use ListSecurityRecipes API.
func (client CloudGuardClient) ListSecurityRecipes(ctx context.Context, request ListSecurityRecipesRequest) (response ListSecurityRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityRecipesResponse")
	}
	return
}

// listSecurityRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listSecurityRecipes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityRecipeCollection/ListSecurityRecipes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListSecurityRecipes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityZones Gets a list of all security zones in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSecurityZones.go.html to see an example of how to use ListSecurityZones API.
func (client CloudGuardClient) ListSecurityZones(ctx context.Context, request ListSecurityZonesRequest) (response ListSecurityZonesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityZones, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityZonesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityZonesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityZonesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityZonesResponse")
	}
	return
}

// listSecurityZones implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listSecurityZones(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityZones", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityZonesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZoneCollection/ListSecurityZones"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListSecurityZones", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSightingEndpoints Returns Sighting endpoints details
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSightingEndpoints.go.html to see an example of how to use ListSightingEndpoints API.
func (client CloudGuardClient) ListSightingEndpoints(ctx context.Context, request ListSightingEndpointsRequest) (response ListSightingEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSightingEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSightingEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSightingEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSightingEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSightingEndpointsResponse")
	}
	return
}

// listSightingEndpoints implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listSightingEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sightings/{sightingId}/endpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSightingEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SightingEndpointSummary/ListSightingEndpoints"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListSightingEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSightingImpactedResources Return a list of Impacted Resources for a CloudGuard Sighting
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSightingImpactedResources.go.html to see an example of how to use ListSightingImpactedResources API.
func (client CloudGuardClient) ListSightingImpactedResources(ctx context.Context, request ListSightingImpactedResourcesRequest) (response ListSightingImpactedResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSightingImpactedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSightingImpactedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSightingImpactedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSightingImpactedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSightingImpactedResourcesResponse")
	}
	return
}

// listSightingImpactedResources implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listSightingImpactedResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sightings/{sightingId}/impactedResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSightingImpactedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SightingImpactedResourceSummary/ListSightingImpactedResources"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListSightingImpactedResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSightings Returns a list of all Sightings identified by the Cloud Guard
// The ListSightings operation returns only sightings that match the passed filters.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListSightings on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSightings.go.html to see an example of how to use ListSightings API.
func (client CloudGuardClient) ListSightings(ctx context.Context, request ListSightingsRequest) (response ListSightingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSightings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSightingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSightingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSightingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSightingsResponse")
	}
	return
}

// listSightings implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listSightings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/sightings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSightingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SightingSummary/ListSightings"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListSightings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTactics Returns a list of tactics associated with detector rules.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTactics.go.html to see an example of how to use ListTactics API.
func (client CloudGuardClient) ListTactics(ctx context.Context, request ListTacticsRequest) (response ListTacticsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTactics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTacticsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTacticsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTacticsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTacticsResponse")
	}
	return
}

// listTactics implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTactics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tactics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTacticsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TacticSummary/ListTactics"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListTactics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetDetectorRecipeDetectorRules Returns a list of DetectorRule associated with DetectorRecipe within a Target.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetDetectorRecipeDetectorRules.go.html to see an example of how to use ListTargetDetectorRecipeDetectorRules API.
func (client CloudGuardClient) ListTargetDetectorRecipeDetectorRules(ctx context.Context, request ListTargetDetectorRecipeDetectorRulesRequest) (response ListTargetDetectorRecipeDetectorRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetDetectorRecipeDetectorRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetDetectorRecipeDetectorRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetDetectorRecipeDetectorRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetDetectorRecipeDetectorRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetDetectorRecipeDetectorRulesResponse")
	}
	return
}

// listTargetDetectorRecipeDetectorRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetDetectorRecipeDetectorRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}/detectorRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetDetectorRecipeDetectorRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipeDetectorRule/ListTargetDetectorRecipeDetectorRules"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListTargetDetectorRecipeDetectorRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetDetectorRecipes Returns a list of all detector recipes associated with the target identified by targetId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetDetectorRecipes.go.html to see an example of how to use ListTargetDetectorRecipes API.
func (client CloudGuardClient) ListTargetDetectorRecipes(ctx context.Context, request ListTargetDetectorRecipesRequest) (response ListTargetDetectorRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetDetectorRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetDetectorRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetDetectorRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetDetectorRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetDetectorRecipesResponse")
	}
	return
}

// listTargetDetectorRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetDetectorRecipes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetDetectorRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetDetectorRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipe/ListTargetDetectorRecipes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListTargetDetectorRecipes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetResponderRecipeResponderRules Returns a list of ResponderRule associated with ResponderRecipe within a Target.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetResponderRecipeResponderRules.go.html to see an example of how to use ListTargetResponderRecipeResponderRules API.
func (client CloudGuardClient) ListTargetResponderRecipeResponderRules(ctx context.Context, request ListTargetResponderRecipeResponderRulesRequest) (response ListTargetResponderRecipeResponderRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetResponderRecipeResponderRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetResponderRecipeResponderRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetResponderRecipeResponderRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetResponderRecipeResponderRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetResponderRecipeResponderRulesResponse")
	}
	return
}

// listTargetResponderRecipeResponderRules implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetResponderRecipeResponderRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}/responderRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetResponderRecipeResponderRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipeResponderRule/ListTargetResponderRecipeResponderRules"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListTargetResponderRecipeResponderRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargetResponderRecipes Returns a list of all responder recipes associated with the target identified by targetId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetResponderRecipes.go.html to see an example of how to use ListTargetResponderRecipes API.
func (client CloudGuardClient) ListTargetResponderRecipes(ctx context.Context, request ListTargetResponderRecipesRequest) (response ListTargetResponderRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargetResponderRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetResponderRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetResponderRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetResponderRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetResponderRecipesResponse")
	}
	return
}

// listTargetResponderRecipes implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargetResponderRecipes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets/{targetId}/targetResponderRecipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetResponderRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipe/ListTargetResponderRecipes"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListTargetResponderRecipes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTargets Returns a list of all Targets in a compartment
// The ListTargets operation returns only the targets in `compartmentId` passed.
// The list does not include any subcompartments of the compartmentId passed.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform ListTargets on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargets.go.html to see an example of how to use ListTargets API.
func (client CloudGuardClient) ListTargets(ctx context.Context, request ListTargetsRequest) (response ListTargetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTargetsResponse")
	}
	return
}

// listTargets implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/targets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Target/ListTargets"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTechniques Returns a list of techniques associated with detector rules.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTechniques.go.html to see an example of how to use ListTechniques API.
func (client CloudGuardClient) ListTechniques(ctx context.Context, request ListTechniquesRequest) (response ListTechniquesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTechniques, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTechniquesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTechniquesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTechniquesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTechniquesResponse")
	}
	return
}

// listTechniques implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listTechniques(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/techniques", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTechniquesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TechniqueSummary/ListTechniques"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListTechniques", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client CloudGuardClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client CloudGuardClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client CloudGuardClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "CloudGuard", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveCompartment Removes an existing compartment from a security zone. When you remove a subcompartment from a security zone, it no longer enforces security zone policies on the resources in the subcompartment. You can't remove the primary compartment that was used to create the security zone.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RemoveCompartment.go.html to see an example of how to use RemoveCompartment API.
func (client CloudGuardClient) RemoveCompartment(ctx context.Context, request RemoveCompartmentRequest) (response RemoveCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveCompartmentResponse")
	}
	return
}

// removeCompartment implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) removeCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityZones/{securityZoneId}/actions/removeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZone/RemoveCompartment"
		err = common.PostProcessServiceError(err, "CloudGuard", "RemoveCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestRiskScores Examines the number of problems related to the resource and the relative severity of those problems.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestRiskScores.go.html to see an example of how to use RequestRiskScores API.
func (client CloudGuardClient) RequestRiskScores(ctx context.Context, request RequestRiskScoresRequest) (response RequestRiskScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestRiskScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestRiskScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestRiskScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestRiskScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestRiskScoresResponse")
	}
	return
}

// requestRiskScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestRiskScores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/riskScores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestRiskScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/RiskScoreAggregation/RequestRiskScores"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestRiskScores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSecurityScoreSummarizedTrend Measures the number of resources examined across all regions and compares it with the
// number of problems detected, for a given time period.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSecurityScoreSummarizedTrend.go.html to see an example of how to use RequestSecurityScoreSummarizedTrend API.
func (client CloudGuardClient) RequestSecurityScoreSummarizedTrend(ctx context.Context, request RequestSecurityScoreSummarizedTrendRequest) (response RequestSecurityScoreSummarizedTrendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSecurityScoreSummarizedTrend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSecurityScoreSummarizedTrendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSecurityScoreSummarizedTrendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSecurityScoreSummarizedTrendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSecurityScoreSummarizedTrendResponse")
	}
	return
}

// requestSecurityScoreSummarizedTrend implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSecurityScoreSummarizedTrend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityScores/actions/summarizeTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSecurityScoreSummarizedTrendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityScoreTrendAggregation/RequestSecurityScoreSummarizedTrend"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSecurityScoreSummarizedTrend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSecurityScores Measures the number of resources examined across all regions and compares it with the number of problems detected.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSecurityScores.go.html to see an example of how to use RequestSecurityScores API.
func (client CloudGuardClient) RequestSecurityScores(ctx context.Context, request RequestSecurityScoresRequest) (response RequestSecurityScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSecurityScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSecurityScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSecurityScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSecurityScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSecurityScoresResponse")
	}
	return
}

// requestSecurityScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSecurityScores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityScores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSecurityScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityScoreAggregation/RequestSecurityScores"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSecurityScores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedActivityProblems Returns the summary of Activity type problems identified by cloud guard, for a given set of dimensions.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
// The compartmentId to be passed with `accessLevel` and `compartmentIdInSubtree` params has to be the root
// compartment id (tenant-id) only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedActivityProblems.go.html to see an example of how to use RequestSummarizedActivityProblems API.
func (client CloudGuardClient) RequestSummarizedActivityProblems(ctx context.Context, request RequestSummarizedActivityProblemsRequest) (response RequestSummarizedActivityProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedActivityProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedActivityProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedActivityProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedActivityProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedActivityProblemsResponse")
	}
	return
}

// requestSummarizedActivityProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedActivityProblems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeActivityProblems", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedActivityProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ActivityProblemAggregation/RequestSummarizedActivityProblems"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedActivityProblems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedProblems Returns the number of problems identified by cloud guard, for a given set of dimensions.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedProblems.go.html to see an example of how to use RequestSummarizedProblems API.
func (client CloudGuardClient) RequestSummarizedProblems(ctx context.Context, request RequestSummarizedProblemsRequest) (response RequestSummarizedProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedProblemsResponse")
	}
	return
}

// requestSummarizedProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedProblems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ProblemAggregation/RequestSummarizedProblems"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedProblems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedResponderExecutions Returns the number of Responder Executions, for a given set of dimensions.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedResponderExecutions.go.html to see an example of how to use RequestSummarizedResponderExecutions API.
func (client CloudGuardClient) RequestSummarizedResponderExecutions(ctx context.Context, request RequestSummarizedResponderExecutionsRequest) (response RequestSummarizedResponderExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedResponderExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedResponderExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedResponderExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedResponderExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedResponderExecutionsResponse")
	}
	return
}

// requestSummarizedResponderExecutions implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedResponderExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/actions/summarize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedResponderExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderExecutionAggregation/RequestSummarizedResponderExecutions"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedResponderExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedRiskScores DEPRECATED
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedRiskScores.go.html to see an example of how to use RequestSummarizedRiskScores API.
func (client CloudGuardClient) RequestSummarizedRiskScores(ctx context.Context, request RequestSummarizedRiskScoresRequest) (response RequestSummarizedRiskScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedRiskScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedRiskScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedRiskScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedRiskScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedRiskScoresResponse")
	}
	return
}

// requestSummarizedRiskScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedRiskScores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeRiskScore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedRiskScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/RiskScoreAggregation/RequestSummarizedRiskScores"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedRiskScores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedSecurityScores DEPRECATED
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedSecurityScores.go.html to see an example of how to use RequestSummarizedSecurityScores API.
func (client CloudGuardClient) RequestSummarizedSecurityScores(ctx context.Context, request RequestSummarizedSecurityScoresRequest) (response RequestSummarizedSecurityScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedSecurityScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedSecurityScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedSecurityScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedSecurityScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedSecurityScoresResponse")
	}
	return
}

// requestSummarizedSecurityScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedSecurityScores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeSecurityScore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedSecurityScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityScoreAggregation/RequestSummarizedSecurityScores"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedSecurityScores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTopTrendResourceProfileRiskScores Summarizes the resource profile risk score top trends for the given time range based on the search filters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedTopTrendResourceProfileRiskScores.go.html to see an example of how to use RequestSummarizedTopTrendResourceProfileRiskScores API.
func (client CloudGuardClient) RequestSummarizedTopTrendResourceProfileRiskScores(ctx context.Context, request RequestSummarizedTopTrendResourceProfileRiskScoresRequest) (response RequestSummarizedTopTrendResourceProfileRiskScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTopTrendResourceProfileRiskScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTopTrendResourceProfileRiskScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTopTrendResourceProfileRiskScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTopTrendResourceProfileRiskScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTopTrendResourceProfileRiskScoresResponse")
	}
	return
}

// requestSummarizedTopTrendResourceProfileRiskScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTopTrendResourceProfileRiskScores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resourceProfileRiskScores/actions/summarizeTopTrends", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTopTrendResourceProfileRiskScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResourceProfileRiskScoreAggregationSummary/RequestSummarizedTopTrendResourceProfileRiskScores"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedTopTrendResourceProfileRiskScores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTrendProblems Returns the number of problems identified by cloud guard, for a given time period.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedTrendProblems.go.html to see an example of how to use RequestSummarizedTrendProblems API.
func (client CloudGuardClient) RequestSummarizedTrendProblems(ctx context.Context, request RequestSummarizedTrendProblemsRequest) (response RequestSummarizedTrendProblemsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTrendProblems, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTrendProblemsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTrendProblemsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTrendProblemsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTrendProblemsResponse")
	}
	return
}

// requestSummarizedTrendProblems implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTrendProblems(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTrendProblemsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ProblemTrendAggregation/RequestSummarizedTrendProblems"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedTrendProblems", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTrendResourceRiskScores Summarizes the resource risk score trend for the given time range based on the search filters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedTrendResourceRiskScores.go.html to see an example of how to use RequestSummarizedTrendResourceRiskScores API.
func (client CloudGuardClient) RequestSummarizedTrendResourceRiskScores(ctx context.Context, request RequestSummarizedTrendResourceRiskScoresRequest) (response RequestSummarizedTrendResourceRiskScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTrendResourceRiskScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTrendResourceRiskScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTrendResourceRiskScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTrendResourceRiskScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTrendResourceRiskScoresResponse")
	}
	return
}

// requestSummarizedTrendResourceRiskScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTrendResourceRiskScores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resourceRiskScores/actions/summarizeTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTrendResourceRiskScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResourceRiskScoreAggregation/RequestSummarizedTrendResourceRiskScores"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedTrendResourceRiskScores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTrendResponderExecutions Returns the number of remediations performed by Responders, for a given time period.
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (ACCESSIBLE) (the resource can be in a subcompartment) or to return Not Authorized if
// Principal doesn't have access to even one of the child compartments. This is valid only when
// `compartmentIdInSubtree` is set to `true`.
// The parameter `compartmentIdInSubtree` applies when you perform summarize API on the
// `compartmentId` passed and when it is set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ACCESSIBLE.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedTrendResponderExecutions.go.html to see an example of how to use RequestSummarizedTrendResponderExecutions API.
func (client CloudGuardClient) RequestSummarizedTrendResponderExecutions(ctx context.Context, request RequestSummarizedTrendResponderExecutionsRequest) (response RequestSummarizedTrendResponderExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTrendResponderExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTrendResponderExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTrendResponderExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTrendResponderExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTrendResponderExecutionsResponse")
	}
	return
}

// requestSummarizedTrendResponderExecutions implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTrendResponderExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/actions/summarizeTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTrendResponderExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderExecutionTrendAggregation/RequestSummarizedTrendResponderExecutions"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedTrendResponderExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedTrendSecurityScores DEPRECATED
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedTrendSecurityScores.go.html to see an example of how to use RequestSummarizedTrendSecurityScores API.
func (client CloudGuardClient) RequestSummarizedTrendSecurityScores(ctx context.Context, request RequestSummarizedTrendSecurityScoresRequest) (response RequestSummarizedTrendSecurityScoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedTrendSecurityScores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedTrendSecurityScoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedTrendSecurityScoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedTrendSecurityScoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedTrendSecurityScoresResponse")
	}
	return
}

// requestSummarizedTrendSecurityScores implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) requestSummarizedTrendSecurityScores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/summarizeSecurityScoreTrend", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedTrendSecurityScoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityScoreTrendAggregation/RequestSummarizedTrendSecurityScores"
		err = common.PostProcessServiceError(err, "CloudGuard", "RequestSummarizedTrendSecurityScores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SkipBulkResponderExecution Skips the execution for a bulk of responder executions
// The operation is atomic in nature
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/SkipBulkResponderExecution.go.html to see an example of how to use SkipBulkResponderExecution API.
func (client CloudGuardClient) SkipBulkResponderExecution(ctx context.Context, request SkipBulkResponderExecutionRequest) (response SkipBulkResponderExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.skipBulkResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SkipBulkResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SkipBulkResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SkipBulkResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SkipBulkResponderExecutionResponse")
	}
	return
}

// skipBulkResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) skipBulkResponderExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/actions/bulkSkip", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SkipBulkResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderExecution/SkipBulkResponderExecution"
		err = common.PostProcessServiceError(err, "CloudGuard", "SkipBulkResponderExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SkipResponderExecution Skips the execution of the responder execution. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/SkipResponderExecution.go.html to see an example of how to use SkipResponderExecution API.
func (client CloudGuardClient) SkipResponderExecution(ctx context.Context, request SkipResponderExecutionRequest) (response SkipResponderExecutionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.skipResponderExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SkipResponderExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SkipResponderExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SkipResponderExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SkipResponderExecutionResponse")
	}
	return
}

// skipResponderExecution implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) skipResponderExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/responderExecutions/{responderExecutionId}/actions/skip", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SkipResponderExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderExecution/SkipResponderExecution"
		err = common.PostProcessServiceError(err, "CloudGuard", "SkipResponderExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TriggerResponder push the problem to responder
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/TriggerResponder.go.html to see an example of how to use TriggerResponder API.
func (client CloudGuardClient) TriggerResponder(ctx context.Context, request TriggerResponderRequest) (response TriggerResponderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.triggerResponder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TriggerResponderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TriggerResponderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TriggerResponderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TriggerResponderResponse")
	}
	return
}

// triggerResponder implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) triggerResponder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/{problemId}/actions/triggerResponder", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TriggerResponderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Problem/TriggerResponder"
		err = common.PostProcessServiceError(err, "CloudGuard", "TriggerResponder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBulkProblemStatus Updates the statuses in bulk for a list of problems
// The operation is atomic in nature
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateBulkProblemStatus.go.html to see an example of how to use UpdateBulkProblemStatus API.
func (client CloudGuardClient) UpdateBulkProblemStatus(ctx context.Context, request UpdateBulkProblemStatusRequest) (response UpdateBulkProblemStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBulkProblemStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBulkProblemStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBulkProblemStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBulkProblemStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBulkProblemStatusResponse")
	}
	return
}

// updateBulkProblemStatus implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateBulkProblemStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/actions/bulkUpdateStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBulkProblemStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Problem/UpdateBulkProblemStatus"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateBulkProblemStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConfiguration Enable/Disable Cloud Guard. The reporting region cannot be updated once created.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateConfiguration.go.html to see an example of how to use UpdateConfiguration API.
func (client CloudGuardClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigurationResponse")
	}
	return
}

// updateConfiguration implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Configuration/UpdateConfiguration"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDataMaskRule Updates a DataMaskRule identified by dataMaskRuleId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateDataMaskRule.go.html to see an example of how to use UpdateDataMaskRule API.
func (client CloudGuardClient) UpdateDataMaskRule(ctx context.Context, request UpdateDataMaskRuleRequest) (response UpdateDataMaskRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataMaskRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataMaskRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataMaskRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataMaskRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataMaskRuleResponse")
	}
	return
}

// updateDataMaskRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateDataMaskRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dataMaskRules/{dataMaskRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataMaskRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataMaskRule/UpdateDataMaskRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateDataMaskRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDataSource Updates a data source identified by dataSourceId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateDataSource.go.html to see an example of how to use UpdateDataSource API.
func (client CloudGuardClient) UpdateDataSource(ctx context.Context, request UpdateDataSourceRequest) (response UpdateDataSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateDataSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataSourceResponse")
	}
	return
}

// updateDataSource implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateDataSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dataSources/{dataSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DataSource/UpdateDataSource"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateDataSource", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDetectorRecipe Updates a detector recipe identified by detectorRecipeId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateDetectorRecipe.go.html to see an example of how to use UpdateDetectorRecipe API.
func (client CloudGuardClient) UpdateDetectorRecipe(ctx context.Context, request UpdateDetectorRecipeRequest) (response UpdateDetectorRecipeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDetectorRecipeResponse")
	}
	return
}

// updateDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/detectorRecipes/{detectorRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipe/UpdateDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDetectorRecipeDetectorRule Update the DetectorRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateDetectorRecipeDetectorRule.go.html to see an example of how to use UpdateDetectorRecipeDetectorRule API.
func (client CloudGuardClient) UpdateDetectorRecipeDetectorRule(ctx context.Context, request UpdateDetectorRecipeDetectorRuleRequest) (response UpdateDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDetectorRecipeDetectorRuleResponse")
	}
	return
}

// updateDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/detectorRecipes/{detectorRecipeId}/detectorRules/{detectorRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/DetectorRecipeDetectorRule/UpdateDetectorRecipeDetectorRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateDetectorRecipeDetectorRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedList Updates a managed list identified by managedListId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateManagedList.go.html to see an example of how to use UpdateManagedList API.
func (client CloudGuardClient) UpdateManagedList(ctx context.Context, request UpdateManagedListRequest) (response UpdateManagedListResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateManagedList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagedListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagedListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagedListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagedListResponse")
	}
	return
}

// updateManagedList implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateManagedList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedLists/{managedListId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagedListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ManagedList/UpdateManagedList"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateManagedList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProblemStatus updates the problem details
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateProblemStatus.go.html to see an example of how to use UpdateProblemStatus API.
func (client CloudGuardClient) UpdateProblemStatus(ctx context.Context, request UpdateProblemStatusRequest) (response UpdateProblemStatusResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateProblemStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProblemStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProblemStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProblemStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProblemStatusResponse")
	}
	return
}

// updateProblemStatus implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateProblemStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/problems/{problemId}/actions/updateStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProblemStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Problem/UpdateProblemStatus"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateProblemStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResponderRecipe Update the ResponderRecipe resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateResponderRecipe.go.html to see an example of how to use UpdateResponderRecipe API.
func (client CloudGuardClient) UpdateResponderRecipe(ctx context.Context, request UpdateResponderRecipeRequest) (response UpdateResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResponderRecipeResponse")
	}
	return
}

// updateResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/responderRecipes/{responderRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipe/UpdateResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResponderRecipeResponderRule Update the ResponderRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateResponderRecipeResponderRule.go.html to see an example of how to use UpdateResponderRecipeResponderRule API.
func (client CloudGuardClient) UpdateResponderRecipeResponderRule(ctx context.Context, request UpdateResponderRecipeResponderRuleRequest) (response UpdateResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResponderRecipeResponderRuleResponse")
	}
	return
}

// updateResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/responderRecipes/{responderRecipeId}/responderRules/{responderRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/ResponderRecipeResponderRule/UpdateResponderRecipeResponderRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateResponderRecipeResponderRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityRecipe Updates a security zone recipe. A security zone recipe is a collection of security zone policies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateSecurityRecipe.go.html to see an example of how to use UpdateSecurityRecipe API.
func (client CloudGuardClient) UpdateSecurityRecipe(ctx context.Context, request UpdateSecurityRecipeRequest) (response UpdateSecurityRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecurityRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityRecipeResponse")
	}
	return
}

// updateSecurityRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateSecurityRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityRecipes/{securityRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityRecipe/UpdateSecurityRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateSecurityRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityZone Updates the security zone identified by its id
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateSecurityZone.go.html to see an example of how to use UpdateSecurityZone API.
func (client CloudGuardClient) UpdateSecurityZone(ctx context.Context, request UpdateSecurityZoneRequest) (response UpdateSecurityZoneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecurityZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityZoneResponse")
	}
	return
}

// updateSecurityZone implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateSecurityZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityZones/{securityZoneId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/SecurityZone/UpdateSecurityZone"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateSecurityZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTarget Updates a Target identified by targetId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateTarget.go.html to see an example of how to use UpdateTarget API.
func (client CloudGuardClient) UpdateTarget(ctx context.Context, request UpdateTargetRequest) (response UpdateTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetResponse")
	}
	return
}

// updateTarget implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/Target/UpdateTarget"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetDetectorRecipe Update the TargetDetectorRecipe resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateTargetDetectorRecipe.go.html to see an example of how to use UpdateTargetDetectorRecipe API.
func (client CloudGuardClient) UpdateTargetDetectorRecipe(ctx context.Context, request UpdateTargetDetectorRecipeRequest) (response UpdateTargetDetectorRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetDetectorRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetDetectorRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetDetectorRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetDetectorRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetDetectorRecipeResponse")
	}
	return
}

// updateTargetDetectorRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetDetectorRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetDetectorRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipe/UpdateTargetDetectorRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateTargetDetectorRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetDetectorRecipeDetectorRule Update the DetectorRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateTargetDetectorRecipeDetectorRule.go.html to see an example of how to use UpdateTargetDetectorRecipeDetectorRule API.
func (client CloudGuardClient) UpdateTargetDetectorRecipeDetectorRule(ctx context.Context, request UpdateTargetDetectorRecipeDetectorRuleRequest) (response UpdateTargetDetectorRecipeDetectorRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetDetectorRecipeDetectorRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetDetectorRecipeDetectorRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetDetectorRecipeDetectorRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetDetectorRecipeDetectorRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetDetectorRecipeDetectorRuleResponse")
	}
	return
}

// updateTargetDetectorRecipeDetectorRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetDetectorRecipeDetectorRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetDetectorRecipes/{targetDetectorRecipeId}/detectorRules/{detectorRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetDetectorRecipeDetectorRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetDetectorRecipeDetectorRule/UpdateTargetDetectorRecipeDetectorRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateTargetDetectorRecipeDetectorRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetResponderRecipe Update the TargetResponderRecipe resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateTargetResponderRecipe.go.html to see an example of how to use UpdateTargetResponderRecipe API.
func (client CloudGuardClient) UpdateTargetResponderRecipe(ctx context.Context, request UpdateTargetResponderRecipeRequest) (response UpdateTargetResponderRecipeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetResponderRecipe, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetResponderRecipeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetResponderRecipeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetResponderRecipeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetResponderRecipeResponse")
	}
	return
}

// updateTargetResponderRecipe implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetResponderRecipe(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetResponderRecipeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipe/UpdateTargetResponderRecipe"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateTargetResponderRecipe", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTargetResponderRecipeResponderRule Update the ResponderRule by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateTargetResponderRecipeResponderRule.go.html to see an example of how to use UpdateTargetResponderRecipeResponderRule API.
func (client CloudGuardClient) UpdateTargetResponderRecipeResponderRule(ctx context.Context, request UpdateTargetResponderRecipeResponderRuleRequest) (response UpdateTargetResponderRecipeResponderRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTargetResponderRecipeResponderRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTargetResponderRecipeResponderRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTargetResponderRecipeResponderRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTargetResponderRecipeResponderRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTargetResponderRecipeResponderRuleResponse")
	}
	return
}

// updateTargetResponderRecipeResponderRule implements the OCIOperation interface (enables retrying operations)
func (client CloudGuardClient) updateTargetResponderRecipeResponderRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/targets/{targetId}/targetResponderRecipes/{targetResponderRecipeId}/responderRules/{responderRuleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTargetResponderRecipeResponderRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/cloud-guard/20200131/TargetResponderRecipeResponderRule/UpdateTargetResponderRecipeResponderRule"
		err = common.PostProcessServiceError(err, "CloudGuard", "UpdateTargetResponderRecipeResponderRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
