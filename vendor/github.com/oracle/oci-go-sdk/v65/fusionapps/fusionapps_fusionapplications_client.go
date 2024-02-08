// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ChangeFusionEnvironmentCompartment.go.html to see an example of how to use ChangeFusionEnvironmentCompartment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ChangeFusionEnvironmentFamilyCompartment.go.html to see an example of how to use ChangeFusionEnvironmentFamilyCompartment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateDataMaskingActivity.go.html to see an example of how to use CreateDataMaskingActivity API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// CreateFusionEnvironment Creates a new FusionEnvironment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateFusionEnvironment.go.html to see an example of how to use CreateFusionEnvironment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateFusionEnvironmentAdminUser.go.html to see an example of how to use CreateFusionEnvironmentAdminUser API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateFusionEnvironmentFamily.go.html to see an example of how to use CreateFusionEnvironmentFamily API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// CreateRefreshActivity Creates a new RefreshActivity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateRefreshActivity.go.html to see an example of how to use CreateRefreshActivity API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/CreateServiceAttachment.go.html to see an example of how to use CreateServiceAttachment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// DeleteFusionEnvironment Deletes the Fusion environment identified by it's OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteFusionEnvironment.go.html to see an example of how to use DeleteFusionEnvironment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteFusionEnvironmentAdminUser.go.html to see an example of how to use DeleteFusionEnvironmentAdminUser API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteFusionEnvironmentFamily.go.html to see an example of how to use DeleteFusionEnvironmentFamily API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// DeleteRefreshActivity Deletes a scheduled RefreshActivity resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteRefreshActivity.go.html to see an example of how to use DeleteRefreshActivity API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/DeleteServiceAttachment.go.html to see an example of how to use DeleteServiceAttachment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// GetDataMaskingActivity Gets a DataMaskingActivity by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetDataMaskingActivity.go.html to see an example of how to use GetDataMaskingActivity API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// GetFusionEnvironment Gets a FusionEnvironment by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironment.go.html to see an example of how to use GetFusionEnvironment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamily.go.html to see an example of how to use GetFusionEnvironmentFamily API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamilyLimitsAndUsage.go.html to see an example of how to use GetFusionEnvironmentFamilyLimitsAndUsage API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentFamilySubscriptionDetail.go.html to see an example of how to use GetFusionEnvironmentFamilySubscriptionDetail API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetFusionEnvironmentStatus.go.html to see an example of how to use GetFusionEnvironmentStatus API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// GetRefreshActivity Gets a RefreshActivity by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetRefreshActivity.go.html to see an example of how to use GetRefreshActivity API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetScheduledActivity.go.html to see an example of how to use GetScheduledActivity API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetServiceAttachment.go.html to see an example of how to use GetServiceAttachment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// ListAdminUsers List all FusionEnvironment admin users
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListAdminUsers.go.html to see an example of how to use ListAdminUsers API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListDataMaskingActivities.go.html to see an example of how to use ListDataMaskingActivities API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// ListFusionEnvironmentFamilies Returns a list of FusionEnvironmentFamilies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListFusionEnvironmentFamilies.go.html to see an example of how to use ListFusionEnvironmentFamilies API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListFusionEnvironments.go.html to see an example of how to use ListFusionEnvironments API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// ListRefreshActivities Returns a list of RefreshActivities.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListRefreshActivities.go.html to see an example of how to use ListRefreshActivities API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListScheduledActivities.go.html to see an example of how to use ListScheduledActivities API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListServiceAttachments.go.html to see an example of how to use ListServiceAttachments API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListTimeAvailableForRefreshes.go.html to see an example of how to use ListTimeAvailableForRefreshes API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// ResetFusionEnvironmentPassword Resets the password of the Fusion Environment Administrator.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ResetFusionEnvironmentPassword.go.html to see an example of how to use ResetFusionEnvironmentPassword API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// UpdateFusionEnvironment Updates the FusionEnvironment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateFusionEnvironment.go.html to see an example of how to use UpdateFusionEnvironment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateFusionEnvironmentFamily.go.html to see an example of how to use UpdateFusionEnvironmentFamily API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// UpdateRefreshActivity Updates a scheduled RefreshActivity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/UpdateRefreshActivity.go.html to see an example of how to use UpdateRefreshActivity API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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

// VerifyServiceAttachment Verify whether a service instance can be attached to the fusion pod
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/VerifyServiceAttachment.go.html to see an example of how to use VerifyServiceAttachment API.
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
	httpResponse, err = client.Call(ctx, &httpRequest)
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
