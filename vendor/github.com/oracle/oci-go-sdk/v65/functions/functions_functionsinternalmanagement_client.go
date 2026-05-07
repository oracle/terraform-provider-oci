// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"

	"regexp"
)

// FunctionsInternalManagementClient a client for FunctionsInternalManagement
type FunctionsInternalManagementClient struct {
	common.BaseClient
	config                   *common.ConfigurationProvider
	requiredParamsInEndpoint map[string][]common.TemplateParamForPerRealmEndpoint
}

// NewFunctionsInternalManagementClientWithConfigurationProvider Creates a new default FunctionsInternalManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFunctionsInternalManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FunctionsInternalManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("functions"); !enabled {
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
	return newFunctionsInternalManagementClientFromBaseClient(baseClient, provider)
}

// NewFunctionsInternalManagementClientWithOboToken Creates a new default FunctionsInternalManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFunctionsInternalManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FunctionsInternalManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFunctionsInternalManagementClientFromBaseClient(baseClient, configProvider)
}

func newFunctionsInternalManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FunctionsInternalManagementClient, err error) {
	// FunctionsInternalManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FunctionsInternalManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FunctionsInternalManagementClient{BaseClient: baseClient}
	client.BasePath = "20260325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FunctionsInternalManagementClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("functions", client.getEndpointTemplatePerRealm(region), "functions")
	client.parseEndpointTemplatePerRealm()
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FunctionsInternalManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FunctionsInternalManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// EnableDualStackEndpoints Determines whether dual stack endpoint should be used or not.
// Default value is false
func (client *FunctionsInternalManagementClient) EnableDualStackEndpoints(enableDualStack bool) {
	client.BaseClient.EnableDualStackEndpoints(enableDualStack)
}

// getEndpointTemplatePerRealm returns the endpoint template for the given region, if not found, returns the default endpoint template
func (client *FunctionsInternalManagementClient) getEndpointTemplatePerRealm(region string) string {
	if client.IsOciRealmSpecificServiceEndpointTemplateEnabled() {
		realm, _ := common.StringToRegion(region).RealmID()
		templatePerRealmDict := map[string]string{
			"oc1":  "https://{dualStack?ds.:}functions.{region}.oci.{secondLevelDomain}",
			"oc16": "https://{dualStack?ds.:}functions.{region}.oci.{secondLevelDomain}",
		}
		if template, ok := templatePerRealmDict[realm]; ok {
			return template
		}
	}
	return "https://functions.{region}.oci.{secondLevelDomain}"
}

// parseEndpointTemplatePerRealm parses the endpoint template per realm from the service endpoint template
// This function will build a map of template params to their values, this map is used when building the API endpoint
func (client *FunctionsInternalManagementClient) parseEndpointTemplatePerRealm() {
	client.requiredParamsInEndpoint = make(map[string][]common.TemplateParamForPerRealmEndpoint)
	templateRegex := regexp.MustCompile(`{.*?}`)
	templateSubRegex := regexp.MustCompile(`{(.+)\+Dot}`)
	templates := templateRegex.FindAllString(client.Host, -1)
	for _, template := range templates {
		templateParam := templateSubRegex.FindStringSubmatch(template)
		if len(templateParam) > 1 {
			client.requiredParamsInEndpoint[templateParam[1]] = append(client.requiredParamsInEndpoint[templateParam[1]], common.TemplateParamForPerRealmEndpoint{
				Template:    templateParam[0],
				EndsWithDot: true,
			})
		} else {
			templateParam := template[1 : len(template)-1]
			client.requiredParamsInEndpoint[templateParam] = append(client.requiredParamsInEndpoint[templateParam], common.TemplateParamForPerRealmEndpoint{
				Template:    template,
				EndsWithDot: false,
			})
		}
	}
}

// SetCustomClientConfiguration sets client with retry and other custom configurations
func (client *FunctionsInternalManagementClient) SetCustomClientConfiguration(config common.CustomClientConfiguration) {
	client.Configuration = config
	client.refreshRegion()
}

// refreshRegion will refresh the region of this client, this function will be called after setting the CustomClientConfiguration
func (client *FunctionsInternalManagementClient) refreshRegion() {
	configProvider := *client.config
	region, _ := configProvider.Region()
	client.SetRegion(region)
}

// DeleteApplicationInternalLogging Deletes the internal Lumberjack logging configuration for an application.
// A default retry strategy applies to this operation DeleteApplicationInternalLogging()
func (client FunctionsInternalManagementClient) DeleteApplicationInternalLogging(ctx context.Context, request DeleteApplicationInternalLoggingRequest) (response DeleteApplicationInternalLoggingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteApplicationInternalLogging, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApplicationInternalLoggingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApplicationInternalLoggingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApplicationInternalLoggingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApplicationInternalLoggingResponse")
	}
	return
}

// deleteApplicationInternalLogging implements the OCIOperation interface (enables retrying operations)
func (client FunctionsInternalManagementClient) deleteApplicationInternalLogging(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/applications/{applicationId}/internalLogging", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(DeleteApplicationInternalLoggingRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response DeleteApplicationInternalLoggingResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "functionsInternalManagement", "DeleteApplicationInternalLogging")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/ApplicationInternalLogging/DeleteApplicationInternalLogging"
		err = common.PostProcessServiceError(err, "FunctionsInternalManagement", "DeleteApplicationInternalLogging", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApplicationInternalLogging Retrieves the internal Lumberjack logging configuration for an application.
// A default retry strategy applies to this operation GetApplicationInternalLogging()
func (client FunctionsInternalManagementClient) GetApplicationInternalLogging(ctx context.Context, request GetApplicationInternalLoggingRequest) (response GetApplicationInternalLoggingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApplicationInternalLogging, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApplicationInternalLoggingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApplicationInternalLoggingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApplicationInternalLoggingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApplicationInternalLoggingResponse")
	}
	return
}

// getApplicationInternalLogging implements the OCIOperation interface (enables retrying operations)
func (client FunctionsInternalManagementClient) getApplicationInternalLogging(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications/{applicationId}/internalLogging", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetApplicationInternalLoggingRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetApplicationInternalLoggingResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "functionsInternalManagement", "GetApplicationInternalLogging")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/ApplicationInternalLogging/GetApplicationInternalLogging"
		err = common.PostProcessServiceError(err, "FunctionsInternalManagement", "GetApplicationInternalLogging", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateApplicationInternalLogging Creates or replaces the internal Lumberjack logging configuration for an application.
// A default retry strategy applies to this operation UpdateApplicationInternalLogging()
func (client FunctionsInternalManagementClient) UpdateApplicationInternalLogging(ctx context.Context, request UpdateApplicationInternalLoggingRequest) (response UpdateApplicationInternalLoggingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateApplicationInternalLogging, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateApplicationInternalLoggingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateApplicationInternalLoggingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateApplicationInternalLoggingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateApplicationInternalLoggingResponse")
	}
	return
}

// updateApplicationInternalLogging implements the OCIOperation interface (enables retrying operations)
func (client FunctionsInternalManagementClient) updateApplicationInternalLogging(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/applications/{applicationId}/internalLogging", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(UpdateApplicationInternalLoggingRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response UpdateApplicationInternalLoggingResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "functionsInternalManagement", "UpdateApplicationInternalLogging")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/ApplicationInternalLogging/UpdateApplicationInternalLogging"
		err = common.PostProcessServiceError(err, "FunctionsInternalManagement", "UpdateApplicationInternalLogging", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
