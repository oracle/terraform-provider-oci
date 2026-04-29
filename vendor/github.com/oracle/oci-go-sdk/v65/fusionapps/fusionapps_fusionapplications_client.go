// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FusionApplicationsClient a client for FusionApplications
type FusionApplicationsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFusionApplicationsClientWithConfigurationProvider Creates a new default FusionApplications client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFusionApplicationsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FusionApplicationsClient, err error) {
	if enabled := common.CheckForEnabledServices("fusionapps"); !enabled {
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
	return newFusionApplicationsClientFromBaseClient(baseClient, provider)
}

// NewFusionApplicationsClientWithOboToken Creates a new default FusionApplications client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFusionApplicationsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FusionApplicationsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFusionApplicationsClientFromBaseClient(baseClient, configProvider)
}

func newFusionApplicationsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FusionApplicationsClient, err error) {
	// FusionApplications service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FusionApplications"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FusionApplicationsClient{BaseClient: baseClient}
	client.BasePath = "20211201"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FusionApplicationsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fusionapps", "https://fusionapps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FusionApplicationsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FusionApplicationsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeFusionEnvironmentCompartment Moves a FusionEnvironment into a different compartment. When provided, If-Match is checked against ETag
// values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ChangeFusionEnvironmentCompartment.go.html to see an example of how to use ChangeFusionEnvironmentCompartment API.
// A default retry strategy applies to this operation ChangeFusionEnvironmentCompartment()
func (client FusionApplicationsClient) ChangeFusionEnvironmentCompartment(ctx context.Context, request ChangeFusionEnvironmentCompartmentRequest) (response ChangeFusionEnvironmentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFusionEnvironmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFusionEnvironmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFusionEnvironmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFusionEnvironmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFusionEnvironmentCompartmentResponse")
	}
	return
}

// changeFusionEnvironmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) changeFusionEnvironmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFusionEnvironmentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ChangeFusionEnvironmentCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/ChangeFusionEnvironmentCompartment"
		err = common.PostProcessServiceError(err, "FusionApplications", "ChangeFusionEnvironmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFusionEnvironmentFamilyCompartment Moves a FusionEnvironmentFamily into a different compartment. When provided, If-Match is checked against ETag
// values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ChangeFusionEnvironmentFamilyCompartment.go.html to see an example of how to use ChangeFusionEnvironmentFamilyCompartment API.
// A default retry strategy applies to this operation ChangeFusionEnvironmentFamilyCompartment()
func (client FusionApplicationsClient) ChangeFusionEnvironmentFamilyCompartment(ctx context.Context, request ChangeFusionEnvironmentFamilyCompartmentRequest) (response ChangeFusionEnvironmentFamilyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFusionEnvironmentFamilyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFusionEnvironmentFamilyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFusionEnvironmentFamilyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFusionEnvironmentFamilyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFusionEnvironmentFamilyCompartmentResponse")
	}
	return
}

// changeFusionEnvironmentFamilyCompartment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) changeFusionEnvironmentFamilyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFusionEnvironmentFamilyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ChangeFusionEnvironmentFamilyCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamily/ChangeFusionEnvironmentFamilyCompartment"
		err = common.PostProcessServiceError(err, "FusionApplications", "ChangeFusionEnvironmentFamilyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataMaskingActivity Creates a new DataMaskingActivity.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateDataMaskingActivity.go.html to see an example of how to use CreateDataMaskingActivity API.
// A default retry strategy applies to this operation CreateDataMaskingActivity()
func (client FusionApplicationsClient) CreateDataMaskingActivity(ctx context.Context, request CreateDataMaskingActivityRequest) (response CreateDataMaskingActivityResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDataMaskingActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataMaskingActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataMaskingActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataMaskingActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataMaskingActivityResponse")
	}
	return
}

// createDataMaskingActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createDataMaskingActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/dataMaskingActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataMaskingActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateDataMaskingActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/DataMaskingActivity/CreateDataMaskingActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateDataMaskingActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailSubdomain Creates an email Subdomain for a brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateEmailSubdomain.go.html to see an example of how to use CreateEmailSubdomain API.
// A default retry strategy applies to this operation CreateEmailSubdomain()
func (client FusionApplicationsClient) CreateEmailSubdomain(ctx context.Context, request CreateEmailSubdomainRequest) (response CreateEmailSubdomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailSubdomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailSubdomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailSubdomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailSubdomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailSubdomainResponse")
	}
	return
}

// createEmailSubdomain implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createEmailSubdomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailSubdomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateEmailSubdomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/CreateEmailSubdomain"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateEmailSubdomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFusionEnvironment Creates a new FusionEnvironment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateFusionEnvironment.go.html to see an example of how to use CreateFusionEnvironment API.
// A default retry strategy applies to this operation CreateFusionEnvironment()
func (client FusionApplicationsClient) CreateFusionEnvironment(ctx context.Context, request CreateFusionEnvironmentRequest) (response CreateFusionEnvironmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFusionEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFusionEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFusionEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFusionEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFusionEnvironmentResponse")
	}
	return
}

// createFusionEnvironment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createFusionEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFusionEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateFusionEnvironment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/CreateFusionEnvironment"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateFusionEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFusionEnvironmentAdminUser Create a FusionEnvironment admin user
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateFusionEnvironmentAdminUser.go.html to see an example of how to use CreateFusionEnvironmentAdminUser API.
// A default retry strategy applies to this operation CreateFusionEnvironmentAdminUser()
func (client FusionApplicationsClient) CreateFusionEnvironmentAdminUser(ctx context.Context, request CreateFusionEnvironmentAdminUserRequest) (response CreateFusionEnvironmentAdminUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFusionEnvironmentAdminUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFusionEnvironmentAdminUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFusionEnvironmentAdminUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFusionEnvironmentAdminUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFusionEnvironmentAdminUserResponse")
	}
	return
}

// createFusionEnvironmentAdminUser implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createFusionEnvironmentAdminUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/adminUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFusionEnvironmentAdminUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateFusionEnvironmentAdminUser")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/CreateFusionEnvironmentAdminUser"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateFusionEnvironmentAdminUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFusionEnvironmentFamily Creates a new FusionEnvironmentFamily.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateFusionEnvironmentFamily.go.html to see an example of how to use CreateFusionEnvironmentFamily API.
// A default retry strategy applies to this operation CreateFusionEnvironmentFamily()
func (client FusionApplicationsClient) CreateFusionEnvironmentFamily(ctx context.Context, request CreateFusionEnvironmentFamilyRequest) (response CreateFusionEnvironmentFamilyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFusionEnvironmentFamilyResponse")
	}
	return
}

// createFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironmentFamilies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateFusionEnvironmentFamily")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamily/CreateFusionEnvironmentFamily"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMarketingBrand Creates a marketing brand for fusion environment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateMarketingBrand.go.html to see an example of how to use CreateMarketingBrand API.
// A default retry strategy applies to this operation CreateMarketingBrand()
func (client FusionApplicationsClient) CreateMarketingBrand(ctx context.Context, request CreateMarketingBrandRequest) (response CreateMarketingBrandResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMarketingBrand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMarketingBrandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMarketingBrandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMarketingBrandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMarketingBrandResponse")
	}
	return
}

// createMarketingBrand implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createMarketingBrand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMarketingBrandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateMarketingBrand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/MarketingBrand/CreateMarketingBrand"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateMarketingBrand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMicrosite Creates a microsite for brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateMicrosite.go.html to see an example of how to use CreateMicrosite API.
// A default retry strategy applies to this operation CreateMicrosite()
func (client FusionApplicationsClient) CreateMicrosite(ctx context.Context, request CreateMicrositeRequest) (response CreateMicrositeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMicrosite, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMicrositeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMicrositeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMicrositeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMicrositeResponse")
	}
	return
}

// createMicrosite implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createMicrosite(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/microsites", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMicrositeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateMicrosite")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/Microsite/CreateMicrosite"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateMicrosite", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRefreshActivity Creates a new RefreshActivity.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateRefreshActivity.go.html to see an example of how to use CreateRefreshActivity API.
// A default retry strategy applies to this operation CreateRefreshActivity()
func (client FusionApplicationsClient) CreateRefreshActivity(ctx context.Context, request CreateRefreshActivityRequest) (response CreateRefreshActivityResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRefreshActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRefreshActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRefreshActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRefreshActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRefreshActivityResponse")
	}
	return
}

// createRefreshActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createRefreshActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/refreshActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRefreshActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateRefreshActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/CreateRefreshActivityDetails/CreateRefreshActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateRefreshActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateServiceAttachment Attaches a service instance to the fusion pod.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateServiceAttachment.go.html to see an example of how to use CreateServiceAttachment API.
// A default retry strategy applies to this operation CreateServiceAttachment()
func (client FusionApplicationsClient) CreateServiceAttachment(ctx context.Context, request CreateServiceAttachmentRequest) (response CreateServiceAttachmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createServiceAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateServiceAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateServiceAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateServiceAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateServiceAttachmentResponse")
	}
	return
}

// createServiceAttachment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createServiceAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/serviceAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateServiceAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateServiceAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/ServiceAttachment/CreateServiceAttachment"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateServiceAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVanityDomain Create a VanityDomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateVanityDomain.go.html to see an example of how to use CreateVanityDomain API.
// A default retry strategy applies to this operation CreateVanityDomain()
func (client FusionApplicationsClient) CreateVanityDomain(ctx context.Context, request CreateVanityDomainRequest) (response CreateVanityDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVanityDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVanityDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVanityDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVanityDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVanityDomainResponse")
	}
	return
}

// createVanityDomain implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createVanityDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVanityDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateVanityDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateVanityDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVanityDomainActivity Create a VanityDomainActivity
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateVanityDomainActivity.go.html to see an example of how to use CreateVanityDomainActivity API.
// A default retry strategy applies to this operation CreateVanityDomainActivity()
func (client FusionApplicationsClient) CreateVanityDomainActivity(ctx context.Context, request CreateVanityDomainActivityRequest) (response CreateVanityDomainActivityResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVanityDomainActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVanityDomainActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVanityDomainActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVanityDomainActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVanityDomainActivityResponse")
	}
	return
}

// createVanityDomainActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) createVanityDomainActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomainActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVanityDomainActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "CreateVanityDomainActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomainActivity/CreateVanityDomainActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "CreateVanityDomainActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailSubdomain Delete an email subdomain for a brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteEmailSubdomain.go.html to see an example of how to use DeleteEmailSubdomain API.
// A default retry strategy applies to this operation DeleteEmailSubdomain()
func (client FusionApplicationsClient) DeleteEmailSubdomain(ctx context.Context, request DeleteEmailSubdomainRequest) (response DeleteEmailSubdomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailSubdomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailSubdomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailSubdomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailSubdomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailSubdomainResponse")
	}
	return
}

// deleteEmailSubdomain implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteEmailSubdomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailSubdomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteEmailSubdomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/DeleteEmailSubdomain"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteEmailSubdomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFusionEnvironment Deletes the Fusion environment identified by it's OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteFusionEnvironment.go.html to see an example of how to use DeleteFusionEnvironment API.
// A default retry strategy applies to this operation DeleteFusionEnvironment()
func (client FusionApplicationsClient) DeleteFusionEnvironment(ctx context.Context, request DeleteFusionEnvironmentRequest) (response DeleteFusionEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFusionEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFusionEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFusionEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFusionEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFusionEnvironmentResponse")
	}
	return
}

// deleteFusionEnvironment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteFusionEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFusionEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteFusionEnvironment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/DeleteFusionEnvironment"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteFusionEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFusionEnvironmentAdminUser Deletes the FusionEnvironment administrator user identified by the username.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteFusionEnvironmentAdminUser.go.html to see an example of how to use DeleteFusionEnvironmentAdminUser API.
// A default retry strategy applies to this operation DeleteFusionEnvironmentAdminUser()
func (client FusionApplicationsClient) DeleteFusionEnvironmentAdminUser(ctx context.Context, request DeleteFusionEnvironmentAdminUserRequest) (response DeleteFusionEnvironmentAdminUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFusionEnvironmentAdminUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFusionEnvironmentAdminUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFusionEnvironmentAdminUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFusionEnvironmentAdminUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFusionEnvironmentAdminUserResponse")
	}
	return
}

// deleteFusionEnvironmentAdminUser implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteFusionEnvironmentAdminUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}/adminUsers/{adminUsername}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFusionEnvironmentAdminUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteFusionEnvironmentAdminUser")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/DeleteFusionEnvironmentAdminUser"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteFusionEnvironmentAdminUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFusionEnvironmentFamily Deletes a FusionEnvironmentFamily resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteFusionEnvironmentFamily.go.html to see an example of how to use DeleteFusionEnvironmentFamily API.
// A default retry strategy applies to this operation DeleteFusionEnvironmentFamily()
func (client FusionApplicationsClient) DeleteFusionEnvironmentFamily(ctx context.Context, request DeleteFusionEnvironmentFamilyRequest) (response DeleteFusionEnvironmentFamilyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFusionEnvironmentFamilyResponse")
	}
	return
}

// deleteFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteFusionEnvironmentFamily")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamily/DeleteFusionEnvironmentFamily"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMarketingBrand Deletes a Marketing brand for fusion Environment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteMarketingBrand.go.html to see an example of how to use DeleteMarketingBrand API.
// A default retry strategy applies to this operation DeleteMarketingBrand()
func (client FusionApplicationsClient) DeleteMarketingBrand(ctx context.Context, request DeleteMarketingBrandRequest) (response DeleteMarketingBrandResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMarketingBrand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMarketingBrandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMarketingBrandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMarketingBrandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMarketingBrandResponse")
	}
	return
}

// deleteMarketingBrand implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteMarketingBrand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMarketingBrandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteMarketingBrand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/MarketingBrand/DeleteMarketingBrand"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteMarketingBrand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMicrosite Delete microsite for a brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteMicrosite.go.html to see an example of how to use DeleteMicrosite API.
// A default retry strategy applies to this operation DeleteMicrosite()
func (client FusionApplicationsClient) DeleteMicrosite(ctx context.Context, request DeleteMicrositeRequest) (response DeleteMicrositeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMicrosite, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMicrositeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMicrositeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMicrositeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMicrositeResponse")
	}
	return
}

// deleteMicrosite implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteMicrosite(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/microsites/{micrositeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMicrositeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteMicrosite")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/Microsite/DeleteMicrosite"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteMicrosite", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRefreshActivity Deletes a scheduled RefreshActivity resource by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteRefreshActivity.go.html to see an example of how to use DeleteRefreshActivity API.
// A default retry strategy applies to this operation DeleteRefreshActivity()
func (client FusionApplicationsClient) DeleteRefreshActivity(ctx context.Context, request DeleteRefreshActivityRequest) (response DeleteRefreshActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRefreshActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRefreshActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRefreshActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRefreshActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRefreshActivityResponse")
	}
	return
}

// deleteRefreshActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteRefreshActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}/refreshActivities/{refreshActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRefreshActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteRefreshActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/RefreshActivity/DeleteRefreshActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteRefreshActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteServiceAttachment Delete a service attachment by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteServiceAttachment.go.html to see an example of how to use DeleteServiceAttachment API.
// A default retry strategy applies to this operation DeleteServiceAttachment()
func (client FusionApplicationsClient) DeleteServiceAttachment(ctx context.Context, request DeleteServiceAttachmentRequest) (response DeleteServiceAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteServiceAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteServiceAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteServiceAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteServiceAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteServiceAttachmentResponse")
	}
	return
}

// deleteServiceAttachment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteServiceAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}/serviceAttachments/{serviceAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteServiceAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteServiceAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/ServiceAttachment/DeleteServiceAttachment"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteServiceAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVanityDomainActivity Deletes a VanityDomainActivity
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteVanityDomainActivity.go.html to see an example of how to use DeleteVanityDomainActivity API.
// A default retry strategy applies to this operation DeleteVanityDomainActivity()
func (client FusionApplicationsClient) DeleteVanityDomainActivity(ctx context.Context, request DeleteVanityDomainActivityRequest) (response DeleteVanityDomainActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVanityDomainActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVanityDomainActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVanityDomainActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVanityDomainActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVanityDomainActivityResponse")
	}
	return
}

// deleteVanityDomainActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) deleteVanityDomainActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomainActivities/{vanityDomainActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVanityDomainActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "DeleteVanityDomainActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomainActivity/DeleteVanityDomainActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "DeleteVanityDomainActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateExtractDetails Begin the process of showing the details about where to retrieve data extract for a Fusion environment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GenerateExtractDetails.go.html to see an example of how to use GenerateExtractDetails API.
// A default retry strategy applies to this operation GenerateExtractDetails()
func (client FusionApplicationsClient) GenerateExtractDetails(ctx context.Context, request GenerateExtractDetailsRequest) (response GenerateExtractDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateExtractDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateExtractDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateExtractDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateExtractDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateExtractDetailsResponse")
	}
	return
}

// generateExtractDetails implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) generateExtractDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/actions/generateExtractDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateExtractDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GenerateExtractDetails")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/GenerateExtractDetails"
		err = common.PostProcessServiceError(err, "FusionApplications", "GenerateExtractDetails", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataMaskingActivity Gets a DataMaskingActivity by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetDataMaskingActivity.go.html to see an example of how to use GetDataMaskingActivity API.
// A default retry strategy applies to this operation GetDataMaskingActivity()
func (client FusionApplicationsClient) GetDataMaskingActivity(ctx context.Context, request GetDataMaskingActivityRequest) (response GetDataMaskingActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataMaskingActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataMaskingActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataMaskingActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataMaskingActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataMaskingActivityResponse")
	}
	return
}

// getDataMaskingActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getDataMaskingActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/dataMaskingActivities/{dataMaskingActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataMaskingActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetDataMaskingActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/DataMaskingActivity/GetDataMaskingActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetDataMaskingActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailSubdomain Gets an email subdomain for the brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetEmailSubdomain.go.html to see an example of how to use GetEmailSubdomain API.
// A default retry strategy applies to this operation GetEmailSubdomain()
func (client FusionApplicationsClient) GetEmailSubdomain(ctx context.Context, request GetEmailSubdomainRequest) (response GetEmailSubdomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailSubdomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailSubdomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailSubdomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailSubdomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailSubdomainResponse")
	}
	return
}

// getEmailSubdomain implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getEmailSubdomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailSubdomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetEmailSubdomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/GetEmailSubdomain"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetEmailSubdomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailSubdomainCsr Gets a CSR for email subdomain for a brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetEmailSubdomainCsr.go.html to see an example of how to use GetEmailSubdomainCsr API.
// A default retry strategy applies to this operation GetEmailSubdomainCsr()
func (client FusionApplicationsClient) GetEmailSubdomainCsr(ctx context.Context, request GetEmailSubdomainCsrRequest) (response GetEmailSubdomainCsrResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailSubdomainCsr, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailSubdomainCsrResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailSubdomainCsrResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailSubdomainCsrResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailSubdomainCsrResponse")
	}
	return
}

// getEmailSubdomainCsr implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getEmailSubdomainCsr(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}/csr", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailSubdomainCsrResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetEmailSubdomainCsr")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/GetEmailSubdomainCsr"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetEmailSubdomainCsr", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailSubdomainDnsConfig Get all DNS records for emailSubdomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetEmailSubdomainDnsConfig.go.html to see an example of how to use GetEmailSubdomainDnsConfig API.
// A default retry strategy applies to this operation GetEmailSubdomainDnsConfig()
func (client FusionApplicationsClient) GetEmailSubdomainDnsConfig(ctx context.Context, request GetEmailSubdomainDnsConfigRequest) (response GetEmailSubdomainDnsConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailSubdomainDnsConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailSubdomainDnsConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailSubdomainDnsConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailSubdomainDnsConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailSubdomainDnsConfigResponse")
	}
	return
}

// getEmailSubdomainDnsConfig implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getEmailSubdomainDnsConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}/dnsConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailSubdomainDnsConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetEmailSubdomainDnsConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/GetEmailSubdomainDnsConfig"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetEmailSubdomainDnsConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironment Gets a FusionEnvironment by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironment.go.html to see an example of how to use GetFusionEnvironment API.
// A default retry strategy applies to this operation GetFusionEnvironment()
func (client FusionApplicationsClient) GetFusionEnvironment(ctx context.Context, request GetFusionEnvironmentRequest) (response GetFusionEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentResponse")
	}
	return
}

// getFusionEnvironment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getFusionEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetFusionEnvironment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/GetFusionEnvironment"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetFusionEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironmentFamily Retrieves a fusion environment family identified by its OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamily.go.html to see an example of how to use GetFusionEnvironmentFamily API.
// A default retry strategy applies to this operation GetFusionEnvironmentFamily()
func (client FusionApplicationsClient) GetFusionEnvironmentFamily(ctx context.Context, request GetFusionEnvironmentFamilyRequest) (response GetFusionEnvironmentFamilyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentFamilyResponse")
	}
	return
}

// getFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetFusionEnvironmentFamily")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamily/GetFusionEnvironmentFamily"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironmentFamilyLimitsAndUsage Gets the number of environments (usage) of each type in the fusion environment family, as well as the limit that's allowed to be created based on the group's associated subscriptions.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamilyLimitsAndUsage.go.html to see an example of how to use GetFusionEnvironmentFamilyLimitsAndUsage API.
// A default retry strategy applies to this operation GetFusionEnvironmentFamilyLimitsAndUsage()
func (client FusionApplicationsClient) GetFusionEnvironmentFamilyLimitsAndUsage(ctx context.Context, request GetFusionEnvironmentFamilyLimitsAndUsageRequest) (response GetFusionEnvironmentFamilyLimitsAndUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironmentFamilyLimitsAndUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentFamilyLimitsAndUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentFamilyLimitsAndUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentFamilyLimitsAndUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentFamilyLimitsAndUsageResponse")
	}
	return
}

// getFusionEnvironmentFamilyLimitsAndUsage implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getFusionEnvironmentFamilyLimitsAndUsage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}/limitsAndUsage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentFamilyLimitsAndUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetFusionEnvironmentFamilyLimitsAndUsage")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamilyLimitsAndUsage/GetFusionEnvironmentFamilyLimitsAndUsage"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetFusionEnvironmentFamilyLimitsAndUsage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironmentFamilySubscriptionDetail Gets the subscription details of an fusion environment family.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamilySubscriptionDetail.go.html to see an example of how to use GetFusionEnvironmentFamilySubscriptionDetail API.
// A default retry strategy applies to this operation GetFusionEnvironmentFamilySubscriptionDetail()
func (client FusionApplicationsClient) GetFusionEnvironmentFamilySubscriptionDetail(ctx context.Context, request GetFusionEnvironmentFamilySubscriptionDetailRequest) (response GetFusionEnvironmentFamilySubscriptionDetailResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironmentFamilySubscriptionDetail, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentFamilySubscriptionDetailResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentFamilySubscriptionDetailResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentFamilySubscriptionDetailResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentFamilySubscriptionDetailResponse")
	}
	return
}

// getFusionEnvironmentFamilySubscriptionDetail implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getFusionEnvironmentFamilySubscriptionDetail(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}/subscriptionDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentFamilySubscriptionDetailResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetFusionEnvironmentFamilySubscriptionDetail")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamily/GetFusionEnvironmentFamilySubscriptionDetail"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetFusionEnvironmentFamilySubscriptionDetail", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFusionEnvironmentStatus Gets the status of a Fusion environment identified by its OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentStatus.go.html to see an example of how to use GetFusionEnvironmentStatus API.
// A default retry strategy applies to this operation GetFusionEnvironmentStatus()
func (client FusionApplicationsClient) GetFusionEnvironmentStatus(ctx context.Context, request GetFusionEnvironmentStatusRequest) (response GetFusionEnvironmentStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFusionEnvironmentStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFusionEnvironmentStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFusionEnvironmentStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFusionEnvironmentStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFusionEnvironmentStatusResponse")
	}
	return
}

// getFusionEnvironmentStatus implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getFusionEnvironmentStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/status", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFusionEnvironmentStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetFusionEnvironmentStatus")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentStatus/GetFusionEnvironmentStatus"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetFusionEnvironmentStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMarketingBrand Gets a Marketing Brand by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetMarketingBrand.go.html to see an example of how to use GetMarketingBrand API.
// A default retry strategy applies to this operation GetMarketingBrand()
func (client FusionApplicationsClient) GetMarketingBrand(ctx context.Context, request GetMarketingBrandRequest) (response GetMarketingBrandResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMarketingBrand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMarketingBrandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMarketingBrandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMarketingBrandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMarketingBrandResponse")
	}
	return
}

// getMarketingBrand implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getMarketingBrand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMarketingBrandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetMarketingBrand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/MarketingBrand/GetMarketingBrand"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetMarketingBrand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMicrosite Get the microsite for the brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetMicrosite.go.html to see an example of how to use GetMicrosite API.
// A default retry strategy applies to this operation GetMicrosite()
func (client FusionApplicationsClient) GetMicrosite(ctx context.Context, request GetMicrositeRequest) (response GetMicrositeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMicrosite, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMicrositeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMicrositeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMicrositeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMicrositeResponse")
	}
	return
}

// getMicrosite implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getMicrosite(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/microsites/{micrositeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMicrositeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetMicrosite")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/Microsite/GetMicrosite"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetMicrosite", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMicrositeDnsConfig Get DNS records for microsite
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetMicrositeDnsConfig.go.html to see an example of how to use GetMicrositeDnsConfig API.
// A default retry strategy applies to this operation GetMicrositeDnsConfig()
func (client FusionApplicationsClient) GetMicrositeDnsConfig(ctx context.Context, request GetMicrositeDnsConfigRequest) (response GetMicrositeDnsConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMicrositeDnsConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMicrositeDnsConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMicrositeDnsConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMicrositeDnsConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMicrositeDnsConfigResponse")
	}
	return
}

// getMicrositeDnsConfig implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getMicrositeDnsConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/microsites/{micrositeId}/dnsConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMicrositeDnsConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetMicrositeDnsConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/Microsite/GetMicrositeDnsConfig"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetMicrositeDnsConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRefreshActivity Gets a RefreshActivity by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetRefreshActivity.go.html to see an example of how to use GetRefreshActivity API.
// A default retry strategy applies to this operation GetRefreshActivity()
func (client FusionApplicationsClient) GetRefreshActivity(ctx context.Context, request GetRefreshActivityRequest) (response GetRefreshActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRefreshActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRefreshActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRefreshActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRefreshActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRefreshActivityResponse")
	}
	return
}

// getRefreshActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getRefreshActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/refreshActivities/{refreshActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRefreshActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetRefreshActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/RefreshActivity/GetRefreshActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetRefreshActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScheduledActivity Gets a ScheduledActivity by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetScheduledActivity.go.html to see an example of how to use GetScheduledActivity API.
// A default retry strategy applies to this operation GetScheduledActivity()
func (client FusionApplicationsClient) GetScheduledActivity(ctx context.Context, request GetScheduledActivityRequest) (response GetScheduledActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScheduledActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledActivityResponse")
	}
	return
}

// getScheduledActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getScheduledActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/scheduledActivities/{scheduledActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetScheduledActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetScheduledActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/ScheduledActivity/GetScheduledActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetScheduledActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetServiceAttachment Gets a Service Attachment by identifier
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetServiceAttachment.go.html to see an example of how to use GetServiceAttachment API.
// A default retry strategy applies to this operation GetServiceAttachment()
func (client FusionApplicationsClient) GetServiceAttachment(ctx context.Context, request GetServiceAttachmentRequest) (response GetServiceAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceAttachmentResponse")
	}
	return
}

// getServiceAttachment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getServiceAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/serviceAttachments/{serviceAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetServiceAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/ServiceAttachment/GetServiceAttachment"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetServiceAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVanityDomain Gets a VanityDomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetVanityDomain.go.html to see an example of how to use GetVanityDomain API.
// A default retry strategy applies to this operation GetVanityDomain()
func (client FusionApplicationsClient) GetVanityDomain(ctx context.Context, request GetVanityDomainRequest) (response GetVanityDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVanityDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVanityDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVanityDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVanityDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVanityDomainResponse")
	}
	return
}

// getVanityDomain implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getVanityDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomains/{vanityDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVanityDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetVanityDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomain/GetVanityDomain"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetVanityDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVanityDomainActivity Gets a VanityDomainActivity
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetVanityDomainActivity.go.html to see an example of how to use GetVanityDomainActivity API.
// A default retry strategy applies to this operation GetVanityDomainActivity()
func (client FusionApplicationsClient) GetVanityDomainActivity(ctx context.Context, request GetVanityDomainActivityRequest) (response GetVanityDomainActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVanityDomainActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVanityDomainActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVanityDomainActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVanityDomainActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVanityDomainActivityResponse")
	}
	return
}

// getVanityDomainActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) getVanityDomainActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomainActivities/{vanityDomainActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVanityDomainActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetVanityDomainActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomainActivity/GetVanityDomainActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetVanityDomainActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client FusionApplicationsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client FusionApplicationsClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "GetWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "FusionApplications", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InitiateExtract Begin the process of generating the data extract for a Fusion environment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/InitiateExtract.go.html to see an example of how to use InitiateExtract API.
// A default retry strategy applies to this operation InitiateExtract()
func (client FusionApplicationsClient) InitiateExtract(ctx context.Context, request InitiateExtractRequest) (response InitiateExtractResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.initiateExtract, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InitiateExtractResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InitiateExtractResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InitiateExtractResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InitiateExtractResponse")
	}
	return
}

// initiateExtract implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) initiateExtract(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/actions/initiateExtract", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InitiateExtractResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "InitiateExtract")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/InitiateExtract"
		err = common.PostProcessServiceError(err, "FusionApplications", "InitiateExtract", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAdminUsers List all FusionEnvironment admin users
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListAdminUsers.go.html to see an example of how to use ListAdminUsers API.
// A default retry strategy applies to this operation ListAdminUsers()
func (client FusionApplicationsClient) ListAdminUsers(ctx context.Context, request ListAdminUsersRequest) (response ListAdminUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAdminUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAdminUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAdminUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAdminUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAdminUsersResponse")
	}
	return
}

// listAdminUsers implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listAdminUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/adminUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAdminUsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListAdminUsers")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/ListAdminUsers"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListAdminUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataMaskingActivities Returns a list of DataMaskingActivities.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListDataMaskingActivities.go.html to see an example of how to use ListDataMaskingActivities API.
// A default retry strategy applies to this operation ListDataMaskingActivities()
func (client FusionApplicationsClient) ListDataMaskingActivities(ctx context.Context, request ListDataMaskingActivitiesRequest) (response ListDataMaskingActivitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataMaskingActivities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataMaskingActivitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataMaskingActivitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataMaskingActivitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataMaskingActivitiesResponse")
	}
	return
}

// listDataMaskingActivities implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listDataMaskingActivities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/dataMaskingActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataMaskingActivitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListDataMaskingActivities")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/DataMaskingActivity/ListDataMaskingActivities"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListDataMaskingActivities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailSubdomains Returns a list of email subdomains for a brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListEmailSubdomains.go.html to see an example of how to use ListEmailSubdomains API.
// A default retry strategy applies to this operation ListEmailSubdomains()
func (client FusionApplicationsClient) ListEmailSubdomains(ctx context.Context, request ListEmailSubdomainsRequest) (response ListEmailSubdomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailSubdomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailSubdomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailSubdomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailSubdomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailSubdomainsResponse")
	}
	return
}

// listEmailSubdomains implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listEmailSubdomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailSubdomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListEmailSubdomains")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/ListEmailSubdomains"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListEmailSubdomains", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFusionEnvironmentFamilies Returns a list of FusionEnvironmentFamilies.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListFusionEnvironmentFamilies.go.html to see an example of how to use ListFusionEnvironmentFamilies API.
// A default retry strategy applies to this operation ListFusionEnvironmentFamilies()
func (client FusionApplicationsClient) ListFusionEnvironmentFamilies(ctx context.Context, request ListFusionEnvironmentFamiliesRequest) (response ListFusionEnvironmentFamiliesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFusionEnvironmentFamilies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFusionEnvironmentFamiliesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFusionEnvironmentFamiliesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFusionEnvironmentFamiliesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFusionEnvironmentFamiliesResponse")
	}
	return
}

// listFusionEnvironmentFamilies implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listFusionEnvironmentFamilies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironmentFamilies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFusionEnvironmentFamiliesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListFusionEnvironmentFamilies")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamily/ListFusionEnvironmentFamilies"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListFusionEnvironmentFamilies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFusionEnvironments Returns a list of FusionEnvironments.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListFusionEnvironments.go.html to see an example of how to use ListFusionEnvironments API.
// A default retry strategy applies to this operation ListFusionEnvironments()
func (client FusionApplicationsClient) ListFusionEnvironments(ctx context.Context, request ListFusionEnvironmentsRequest) (response ListFusionEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFusionEnvironments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFusionEnvironmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFusionEnvironmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFusionEnvironmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFusionEnvironmentsResponse")
	}
	return
}

// listFusionEnvironments implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listFusionEnvironments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFusionEnvironmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListFusionEnvironments")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/ListFusionEnvironments"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListFusionEnvironments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMarketingBrands Returns a list of marketing brands
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListMarketingBrands.go.html to see an example of how to use ListMarketingBrands API.
// A default retry strategy applies to this operation ListMarketingBrands()
func (client FusionApplicationsClient) ListMarketingBrands(ctx context.Context, request ListMarketingBrandsRequest) (response ListMarketingBrandsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMarketingBrands, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMarketingBrandsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMarketingBrandsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMarketingBrandsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMarketingBrandsResponse")
	}
	return
}

// listMarketingBrands implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listMarketingBrands(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMarketingBrandsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListMarketingBrands")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/MarketingBrand/ListMarketingBrands"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListMarketingBrands", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMicrosites Returns a list of microsites
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListMicrosites.go.html to see an example of how to use ListMicrosites API.
// A default retry strategy applies to this operation ListMicrosites()
func (client FusionApplicationsClient) ListMicrosites(ctx context.Context, request ListMicrositesRequest) (response ListMicrositesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMicrosites, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMicrositesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMicrositesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMicrositesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMicrositesResponse")
	}
	return
}

// listMicrosites implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listMicrosites(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/microsites", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMicrositesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListMicrosites")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/Microsite/ListMicrosites"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListMicrosites", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRefreshActivities Returns a list of RefreshActivities.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListRefreshActivities.go.html to see an example of how to use ListRefreshActivities API.
// A default retry strategy applies to this operation ListRefreshActivities()
func (client FusionApplicationsClient) ListRefreshActivities(ctx context.Context, request ListRefreshActivitiesRequest) (response ListRefreshActivitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRefreshActivities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRefreshActivitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRefreshActivitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRefreshActivitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRefreshActivitiesResponse")
	}
	return
}

// listRefreshActivities implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listRefreshActivities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/refreshActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRefreshActivitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListRefreshActivities")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/RefreshActivity/ListRefreshActivities"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListRefreshActivities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledActivities Returns a list of ScheduledActivities.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListScheduledActivities.go.html to see an example of how to use ListScheduledActivities API.
// A default retry strategy applies to this operation ListScheduledActivities()
func (client FusionApplicationsClient) ListScheduledActivities(ctx context.Context, request ListScheduledActivitiesRequest) (response ListScheduledActivitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledActivities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledActivitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledActivitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledActivitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledActivitiesResponse")
	}
	return
}

// listScheduledActivities implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listScheduledActivities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/scheduledActivities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListScheduledActivitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListScheduledActivities")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/ScheduledActivity/ListScheduledActivities"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListScheduledActivities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceAttachments Returns a list of service attachments.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListServiceAttachments.go.html to see an example of how to use ListServiceAttachments API.
// A default retry strategy applies to this operation ListServiceAttachments()
func (client FusionApplicationsClient) ListServiceAttachments(ctx context.Context, request ListServiceAttachmentsRequest) (response ListServiceAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceAttachmentsResponse")
	}
	return
}

// listServiceAttachments implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listServiceAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/serviceAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListServiceAttachments")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/ServiceAttachment/ListServiceAttachments"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListServiceAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTimeAvailableForRefreshes Gets available refresh time for this fusion environment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListTimeAvailableForRefreshes.go.html to see an example of how to use ListTimeAvailableForRefreshes API.
// A default retry strategy applies to this operation ListTimeAvailableForRefreshes()
func (client FusionApplicationsClient) ListTimeAvailableForRefreshes(ctx context.Context, request ListTimeAvailableForRefreshesRequest) (response ListTimeAvailableForRefreshesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTimeAvailableForRefreshes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTimeAvailableForRefreshesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTimeAvailableForRefreshesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTimeAvailableForRefreshesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTimeAvailableForRefreshesResponse")
	}
	return
}

// listTimeAvailableForRefreshes implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listTimeAvailableForRefreshes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/timeAvailableForRefresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTimeAvailableForRefreshesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListTimeAvailableForRefreshes")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/TimeAvailableForRefresh/ListTimeAvailableForRefreshes"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListTimeAvailableForRefreshes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVanityDomains Lists all VanityDomains.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListVanityDomains.go.html to see an example of how to use ListVanityDomains API.
// A default retry strategy applies to this operation ListVanityDomains()
func (client FusionApplicationsClient) ListVanityDomains(ctx context.Context, request ListVanityDomainsRequest) (response ListVanityDomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVanityDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVanityDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVanityDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVanityDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVanityDomainsResponse")
	}
	return
}

// listVanityDomains implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) listVanityDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVanityDomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListVanityDomains")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomain/ListVanityDomains"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListVanityDomains", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client FusionApplicationsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client FusionApplicationsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListWorkRequestErrors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client FusionApplicationsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client FusionApplicationsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListWorkRequestLogs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client FusionApplicationsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client FusionApplicationsClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ListWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "FusionApplications", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestEmailSubdomainCsr Request Email Subdomain CSR
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/RequestEmailSubdomainCsr.go.html to see an example of how to use RequestEmailSubdomainCsr API.
// A default retry strategy applies to this operation RequestEmailSubdomainCsr()
func (client FusionApplicationsClient) RequestEmailSubdomainCsr(ctx context.Context, request RequestEmailSubdomainCsrRequest) (response RequestEmailSubdomainCsrResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestEmailSubdomainCsr, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestEmailSubdomainCsrResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestEmailSubdomainCsrResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestEmailSubdomainCsrResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestEmailSubdomainCsrResponse")
	}
	return
}

// requestEmailSubdomainCsr implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) requestEmailSubdomainCsr(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}/actions/requestEmailSubdomainCsr", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestEmailSubdomainCsrResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "RequestEmailSubdomainCsr")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/RequestEmailSubdomainCsr"
		err = common.PostProcessServiceError(err, "FusionApplications", "RequestEmailSubdomainCsr", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResetFusionEnvironmentPassword Reset FusionEnvironment admin password. This API will be deprecated on Mon, 15 Jan 2024 01:00:00 GMT. Users can reset password themselves, FAaaS will no longer provide an API for this.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ResetFusionEnvironmentPassword.go.html to see an example of how to use ResetFusionEnvironmentPassword API.
// A default retry strategy applies to this operation ResetFusionEnvironmentPassword()
func (client FusionApplicationsClient) ResetFusionEnvironmentPassword(ctx context.Context, request ResetFusionEnvironmentPasswordRequest) (response ResetFusionEnvironmentPasswordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resetFusionEnvironmentPassword, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResetFusionEnvironmentPasswordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResetFusionEnvironmentPasswordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResetFusionEnvironmentPasswordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResetFusionEnvironmentPasswordResponse")
	}
	return
}

// resetFusionEnvironmentPassword implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) resetFusionEnvironmentPassword(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/adminUsers/{adminUsername}/actions/resetPassword", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResetFusionEnvironmentPasswordResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ResetFusionEnvironmentPassword")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/ResetFusionEnvironmentPassword"
		err = common.PostProcessServiceError(err, "FusionApplications", "ResetFusionEnvironmentPassword", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SubmitVanityDomainValidation Submit Vanity Domain Validation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/SubmitVanityDomainValidation.go.html to see an example of how to use SubmitVanityDomainValidation API.
// A default retry strategy applies to this operation SubmitVanityDomainValidation()
func (client FusionApplicationsClient) SubmitVanityDomainValidation(ctx context.Context, request SubmitVanityDomainValidationRequest) (response SubmitVanityDomainValidationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.submitVanityDomainValidation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SubmitVanityDomainValidationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SubmitVanityDomainValidationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SubmitVanityDomainValidationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SubmitVanityDomainValidationResponse")
	}
	return
}

// submitVanityDomainValidation implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) submitVanityDomainValidation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomains/{vanityDomainId}/actions/submitValidation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SubmitVanityDomainValidationResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "SubmitVanityDomainValidation")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomain/SubmitVanityDomainValidation"
		err = common.PostProcessServiceError(err, "FusionApplications", "SubmitVanityDomainValidation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailSubdomain Updates an email subdomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateEmailSubdomain.go.html to see an example of how to use UpdateEmailSubdomain API.
// A default retry strategy applies to this operation UpdateEmailSubdomain()
func (client FusionApplicationsClient) UpdateEmailSubdomain(ctx context.Context, request UpdateEmailSubdomainRequest) (response UpdateEmailSubdomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailSubdomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailSubdomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailSubdomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailSubdomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailSubdomainResponse")
	}
	return
}

// updateEmailSubdomain implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateEmailSubdomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailSubdomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateEmailSubdomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/UpdateEmailSubdomain"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateEmailSubdomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFusionEnvironment Updates the FusionEnvironment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateFusionEnvironment.go.html to see an example of how to use UpdateFusionEnvironment API.
// A default retry strategy applies to this operation UpdateFusionEnvironment()
func (client FusionApplicationsClient) UpdateFusionEnvironment(ctx context.Context, request UpdateFusionEnvironmentRequest) (response UpdateFusionEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFusionEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFusionEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFusionEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFusionEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFusionEnvironmentResponse")
	}
	return
}

// updateFusionEnvironment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateFusionEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironments/{fusionEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFusionEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateFusionEnvironment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironment/UpdateFusionEnvironment"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateFusionEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFusionEnvironmentFamily Updates the FusionEnvironmentFamily
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateFusionEnvironmentFamily.go.html to see an example of how to use UpdateFusionEnvironmentFamily API.
// A default retry strategy applies to this operation UpdateFusionEnvironmentFamily()
func (client FusionApplicationsClient) UpdateFusionEnvironmentFamily(ctx context.Context, request UpdateFusionEnvironmentFamilyRequest) (response UpdateFusionEnvironmentFamilyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFusionEnvironmentFamily, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFusionEnvironmentFamilyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFusionEnvironmentFamilyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFusionEnvironmentFamilyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFusionEnvironmentFamilyResponse")
	}
	return
}

// updateFusionEnvironmentFamily implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateFusionEnvironmentFamily(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironmentFamilies/{fusionEnvironmentFamilyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFusionEnvironmentFamilyResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateFusionEnvironmentFamily")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/FusionEnvironmentFamily/UpdateFusionEnvironmentFamily"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateFusionEnvironmentFamily", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMarketingBrand Updates a Marketing Brand
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateMarketingBrand.go.html to see an example of how to use UpdateMarketingBrand API.
// A default retry strategy applies to this operation UpdateMarketingBrand()
func (client FusionApplicationsClient) UpdateMarketingBrand(ctx context.Context, request UpdateMarketingBrandRequest) (response UpdateMarketingBrandResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMarketingBrand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMarketingBrandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMarketingBrandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMarketingBrandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMarketingBrandResponse")
	}
	return
}

// updateMarketingBrand implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateMarketingBrand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMarketingBrandResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateMarketingBrand")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/MarketingBrand/UpdateMarketingBrand"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateMarketingBrand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMicrosite Updates an microsite
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateMicrosite.go.html to see an example of how to use UpdateMicrosite API.
// A default retry strategy applies to this operation UpdateMicrosite()
func (client FusionApplicationsClient) UpdateMicrosite(ctx context.Context, request UpdateMicrositeRequest) (response UpdateMicrositeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMicrosite, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMicrositeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMicrositeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMicrositeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMicrositeResponse")
	}
	return
}

// updateMicrosite implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateMicrosite(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/microsites/{micrositeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMicrositeResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateMicrosite")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/Microsite/UpdateMicrosite"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateMicrosite", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRefreshActivity Updates a scheduled RefreshActivity.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateRefreshActivity.go.html to see an example of how to use UpdateRefreshActivity API.
// A default retry strategy applies to this operation UpdateRefreshActivity()
func (client FusionApplicationsClient) UpdateRefreshActivity(ctx context.Context, request UpdateRefreshActivityRequest) (response UpdateRefreshActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRefreshActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRefreshActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRefreshActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRefreshActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRefreshActivityResponse")
	}
	return
}

// updateRefreshActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateRefreshActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironments/{fusionEnvironmentId}/refreshActivities/{refreshActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRefreshActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateRefreshActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/UpdateRefreshActivityDetails/UpdateRefreshActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateRefreshActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVanityDomain Updates a VanityDomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateVanityDomain.go.html to see an example of how to use UpdateVanityDomain API.
// A default retry strategy applies to this operation UpdateVanityDomain()
func (client FusionApplicationsClient) UpdateVanityDomain(ctx context.Context, request UpdateVanityDomainRequest) (response UpdateVanityDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVanityDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVanityDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVanityDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVanityDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVanityDomainResponse")
	}
	return
}

// updateVanityDomain implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateVanityDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomains/{vanityDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVanityDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateVanityDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomain/UpdateVanityDomain"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateVanityDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVanityDomainActivity Updates a VanityDomainActivity
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateVanityDomainActivity.go.html to see an example of how to use UpdateVanityDomainActivity API.
// A default retry strategy applies to this operation UpdateVanityDomainActivity()
func (client FusionApplicationsClient) UpdateVanityDomainActivity(ctx context.Context, request UpdateVanityDomainActivityRequest) (response UpdateVanityDomainActivityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVanityDomainActivity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVanityDomainActivityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVanityDomainActivityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVanityDomainActivityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVanityDomainActivityResponse")
	}
	return
}

// updateVanityDomainActivity implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) updateVanityDomainActivity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomainActivities/{vanityDomainActivityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVanityDomainActivityResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UpdateVanityDomainActivity")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomainActivity/UpdateVanityDomainActivity"
		err = common.PostProcessServiceError(err, "FusionApplications", "UpdateVanityDomainActivity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadEmailSubdomainCertificate upload certificate for emailSubdomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UploadEmailSubdomainCertificate.go.html to see an example of how to use UploadEmailSubdomainCertificate API.
// A default retry strategy applies to this operation UploadEmailSubdomainCertificate()
func (client FusionApplicationsClient) UploadEmailSubdomainCertificate(ctx context.Context, request UploadEmailSubdomainCertificateRequest) (response UploadEmailSubdomainCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.uploadEmailSubdomainCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadEmailSubdomainCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadEmailSubdomainCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadEmailSubdomainCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadEmailSubdomainCertificateResponse")
	}
	return
}

// uploadEmailSubdomainCertificate implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) uploadEmailSubdomainCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}/actions/uploadEmailSubdomainCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadEmailSubdomainCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UploadEmailSubdomainCertificate")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/UploadEmailSubdomainCertificate"
		err = common.PostProcessServiceError(err, "FusionApplications", "UploadEmailSubdomainCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadVanityDomainCertificate Upload Vanity Domain certificate
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UploadVanityDomainCertificate.go.html to see an example of how to use UploadVanityDomainCertificate API.
// A default retry strategy applies to this operation UploadVanityDomainCertificate()
func (client FusionApplicationsClient) UploadVanityDomainCertificate(ctx context.Context, request UploadVanityDomainCertificateRequest) (response UploadVanityDomainCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.uploadVanityDomainCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadVanityDomainCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadVanityDomainCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadVanityDomainCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadVanityDomainCertificateResponse")
	}
	return
}

// uploadVanityDomainCertificate implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) uploadVanityDomainCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/vanityDomains/{vanityDomainId}/actions/uploadCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadVanityDomainCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "UploadVanityDomainCertificate")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/VanityDomain/UploadVanityDomainCertificate"
		err = common.PostProcessServiceError(err, "FusionApplications", "UploadVanityDomainCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateAndConfigureEmailSubdomainCertificate Validate and configure certificate for emailSubdomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ValidateAndConfigureEmailSubdomainCertificate.go.html to see an example of how to use ValidateAndConfigureEmailSubdomainCertificate API.
// A default retry strategy applies to this operation ValidateAndConfigureEmailSubdomainCertificate()
func (client FusionApplicationsClient) ValidateAndConfigureEmailSubdomainCertificate(ctx context.Context, request ValidateAndConfigureEmailSubdomainCertificateRequest) (response ValidateAndConfigureEmailSubdomainCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateAndConfigureEmailSubdomainCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateAndConfigureEmailSubdomainCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateAndConfigureEmailSubdomainCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateAndConfigureEmailSubdomainCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateAndConfigureEmailSubdomainCertificateResponse")
	}
	return
}

// validateAndConfigureEmailSubdomainCertificate implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) validateAndConfigureEmailSubdomainCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}/actions/validateAndConfigureCertificate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateAndConfigureEmailSubdomainCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ValidateAndConfigureEmailSubdomainCertificate")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/ValidateAndConfigureEmailSubdomainCertificate"
		err = common.PostProcessServiceError(err, "FusionApplications", "ValidateAndConfigureEmailSubdomainCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateAndConfigureEmailSubdomainDns Validate and configure DNS records for emailSubdomain
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ValidateAndConfigureEmailSubdomainDns.go.html to see an example of how to use ValidateAndConfigureEmailSubdomainDns API.
// A default retry strategy applies to this operation ValidateAndConfigureEmailSubdomainDns()
func (client FusionApplicationsClient) ValidateAndConfigureEmailSubdomainDns(ctx context.Context, request ValidateAndConfigureEmailSubdomainDnsRequest) (response ValidateAndConfigureEmailSubdomainDnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateAndConfigureEmailSubdomainDns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateAndConfigureEmailSubdomainDnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateAndConfigureEmailSubdomainDnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateAndConfigureEmailSubdomainDnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateAndConfigureEmailSubdomainDnsResponse")
	}
	return
}

// validateAndConfigureEmailSubdomainDns implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) validateAndConfigureEmailSubdomainDns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/emailSubdomains/{emailSubdomainId}/actions/validateAndConfigureDns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateAndConfigureEmailSubdomainDnsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ValidateAndConfigureEmailSubdomainDns")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/EmailSubdomain/ValidateAndConfigureEmailSubdomainDns"
		err = common.PostProcessServiceError(err, "FusionApplications", "ValidateAndConfigureEmailSubdomainDns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateAndConfigureMicrositeDns Validate and configure DNS records for microsite
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ValidateAndConfigureMicrositeDns.go.html to see an example of how to use ValidateAndConfigureMicrositeDns API.
// A default retry strategy applies to this operation ValidateAndConfigureMicrositeDns()
func (client FusionApplicationsClient) ValidateAndConfigureMicrositeDns(ctx context.Context, request ValidateAndConfigureMicrositeDnsRequest) (response ValidateAndConfigureMicrositeDnsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateAndConfigureMicrositeDns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateAndConfigureMicrositeDnsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateAndConfigureMicrositeDnsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateAndConfigureMicrositeDnsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateAndConfigureMicrositeDnsResponse")
	}
	return
}

// validateAndConfigureMicrositeDns implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) validateAndConfigureMicrositeDns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/marketingBrands/{marketingBrandId}/microsites/{micrositeId}/actions/validateAndConfigureDns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateAndConfigureMicrositeDnsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "ValidateAndConfigureMicrositeDns")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/Microsite/ValidateAndConfigureMicrositeDns"
		err = common.PostProcessServiceError(err, "FusionApplications", "ValidateAndConfigureMicrositeDns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// VerifyServiceAttachment Verify whether a service instance can be attached to the fusion pod
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/VerifyServiceAttachment.go.html to see an example of how to use VerifyServiceAttachment API.
// A default retry strategy applies to this operation VerifyServiceAttachment()
func (client FusionApplicationsClient) VerifyServiceAttachment(ctx context.Context, request VerifyServiceAttachmentRequest) (response VerifyServiceAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.verifyServiceAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = VerifyServiceAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = VerifyServiceAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(VerifyServiceAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into VerifyServiceAttachmentResponse")
	}
	return
}

// verifyServiceAttachment implements the OCIOperation interface (enables retrying operations)
func (client FusionApplicationsClient) verifyServiceAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fusionEnvironments/{fusionEnvironmentId}/serviceAttachments/actions/verify", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response VerifyServiceAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "fusionApplications", "VerifyServiceAttachment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fusion-applications/20211201/ServiceAttachment/VerifyServiceAttachment"
		err = common.PostProcessServiceError(err, "FusionApplications", "VerifyServiceAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
