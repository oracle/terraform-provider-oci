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

// WorkRequestManagementClient a client for WorkRequestManagement
type WorkRequestManagementClient struct {
	common.BaseClient
	config                   *common.ConfigurationProvider
	requiredParamsInEndpoint map[string][]common.TemplateParamForPerRealmEndpoint
}

// NewWorkRequestManagementClientWithConfigurationProvider Creates a new default WorkRequestManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewWorkRequestManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client WorkRequestManagementClient, err error) {
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
	return newWorkRequestManagementClientFromBaseClient(baseClient, provider)
}

// NewWorkRequestManagementClientWithOboToken Creates a new default WorkRequestManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewWorkRequestManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client WorkRequestManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newWorkRequestManagementClientFromBaseClient(baseClient, configProvider)
}

func newWorkRequestManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client WorkRequestManagementClient, err error) {
	// WorkRequestManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("WorkRequestManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = WorkRequestManagementClient{BaseClient: baseClient}
	client.BasePath = "20260325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *WorkRequestManagementClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("functions", client.getEndpointTemplatePerRealm(region), "functions")
	client.parseEndpointTemplatePerRealm()
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *WorkRequestManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *WorkRequestManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// EnableDualStackEndpoints Determines whether dual stack endpoint should be used or not.
// Default value is false
func (client *WorkRequestManagementClient) EnableDualStackEndpoints(enableDualStack bool) {
	client.BaseClient.EnableDualStackEndpoints(enableDualStack)
}

// getEndpointTemplatePerRealm returns the endpoint template for the given region, if not found, returns the default endpoint template
func (client *WorkRequestManagementClient) getEndpointTemplatePerRealm(region string) string {
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
func (client *WorkRequestManagementClient) parseEndpointTemplatePerRealm() {
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
func (client *WorkRequestManagementClient) SetCustomClientConfiguration(config common.CustomClientConfiguration) {
	client.Configuration = config
	client.refreshRegion()
}

// refreshRegion will refresh the region of this client, this function will be called after setting the CustomClientConfiguration
func (client *WorkRequestManagementClient) refreshRegion() {
	configProvider := *client.config
	region, _ := configProvider.Region()
	client.SetRegion(region)
}

// CancelWorkRequest Cancels a work request.
func (client WorkRequestManagementClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
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
func (client WorkRequestManagementClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(CancelWorkRequestRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "workRequestManagement", "CancelWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "WorkRequestManagement", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the details of a work request.
// A default retry strategy applies to this operation GetWorkRequest()
func (client WorkRequestManagementClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client WorkRequestManagementClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(GetWorkRequestRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "workRequestManagement", "GetWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "WorkRequestManagement", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists the errors for a work request.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client WorkRequestManagementClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client WorkRequestManagementClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListWorkRequestErrorsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "workRequestManagement", "ListWorkRequestErrors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "WorkRequestManagement", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs for a work request.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client WorkRequestManagementClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client WorkRequestManagementClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListWorkRequestLogsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "workRequestManagement", "ListWorkRequestLogs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "WorkRequestManagement", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
// A default retry strategy applies to this operation ListWorkRequests()
func (client WorkRequestManagementClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client WorkRequestManagementClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	host := client.Host
	request.(ListWorkRequestsRequest).ReplaceMandatoryParamInPath(&client.BaseClient, client.requiredParamsInEndpoint)
	common.UpdateEndpointTemplateForOptions(&client.BaseClient)
	common.SetMissingTemplateParams(&client.BaseClient)
	defer func() {
		client.Host = host
	}()

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "workRequestManagement", "ListWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/functions/20260325/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "WorkRequestManagement", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
