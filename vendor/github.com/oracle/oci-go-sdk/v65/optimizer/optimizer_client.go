// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OptimizerClient a client for Optimizer
type OptimizerClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOptimizerClientWithConfigurationProvider Creates a new default Optimizer client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOptimizerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OptimizerClient, err error) {
	if enabled := common.CheckForEnabledServices("optimizer"); !enabled {
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
	return newOptimizerClientFromBaseClient(baseClient, provider)
}

// NewOptimizerClientWithOboToken Creates a new default Optimizer client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOptimizerClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OptimizerClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOptimizerClientFromBaseClient(baseClient, configProvider)
}

func newOptimizerClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OptimizerClient, err error) {
	// Optimizer service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Optimizer"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OptimizerClient{BaseClient: baseClient}
	client.BasePath = "20200606"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OptimizerClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("optimizer", "https://optimizer.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OptimizerClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OptimizerClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkApplyRecommendations Applies the specified recommendations to the resources.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/BulkApplyRecommendations.go.html to see an example of how to use BulkApplyRecommendations API.
// A default retry strategy applies to this operation BulkApplyRecommendations()
func (client OptimizerClient) BulkApplyRecommendations(ctx context.Context, request BulkApplyRecommendationsRequest) (response BulkApplyRecommendationsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkApplyRecommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkApplyRecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkApplyRecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkApplyRecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkApplyRecommendationsResponse")
	}
	return
}

// bulkApplyRecommendations implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) bulkApplyRecommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/recommendations/{recommendationId}/actions/bulkApplyRecommendations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkApplyRecommendationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Recommendation/BulkApplyRecommendations"
		err = common.PostProcessServiceError(err, "Optimizer", "BulkApplyRecommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProfile Creates a new profile.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/CreateProfile.go.html to see an example of how to use CreateProfile API.
// A default retry strategy applies to this operation CreateProfile()
func (client OptimizerClient) CreateProfile(ctx context.Context, request CreateProfileRequest) (response CreateProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProfileResponse")
	}
	return
}

// createProfile implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) createProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/profiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Profile/CreateProfile"
		err = common.PostProcessServiceError(err, "Optimizer", "CreateProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProfile Deletes the specified profile. Uses the profile's OCID to determine which profile to delete.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/DeleteProfile.go.html to see an example of how to use DeleteProfile API.
// A default retry strategy applies to this operation DeleteProfile()
func (client OptimizerClient) DeleteProfile(ctx context.Context, request DeleteProfileRequest) (response DeleteProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProfileResponse")
	}
	return
}

// deleteProfile implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) deleteProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/profiles/{profileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Profile/DeleteProfile"
		err = common.PostProcessServiceError(err, "Optimizer", "DeleteProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FilterResourceActions Queries the Cloud Advisor resource actions that are supported.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/FilterResourceActions.go.html to see an example of how to use FilterResourceActions API.
// A default retry strategy applies to this operation FilterResourceActions()
func (client OptimizerClient) FilterResourceActions(ctx context.Context, request FilterResourceActionsRequest) (response FilterResourceActionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.filterResourceActions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FilterResourceActionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FilterResourceActionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FilterResourceActionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FilterResourceActionsResponse")
	}
	return
}

// filterResourceActions implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) filterResourceActions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/filterResourceActions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FilterResourceActionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/ResourceActionSummary/FilterResourceActions"
		err = common.PostProcessServiceError(err, "Optimizer", "FilterResourceActions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCategory Gets the category that corresponds to the specified OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/GetCategory.go.html to see an example of how to use GetCategory API.
// A default retry strategy applies to this operation GetCategory()
func (client OptimizerClient) GetCategory(ctx context.Context, request GetCategoryRequest) (response GetCategoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCategory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCategoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCategoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCategoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCategoryResponse")
	}
	return
}

// getCategory implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) getCategory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/categories/{categoryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCategoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Category/GetCategory"
		err = common.PostProcessServiceError(err, "Optimizer", "GetCategory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEnrollmentStatus Gets the Cloud Advisor enrollment status.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/GetEnrollmentStatus.go.html to see an example of how to use GetEnrollmentStatus API.
// A default retry strategy applies to this operation GetEnrollmentStatus()
func (client OptimizerClient) GetEnrollmentStatus(ctx context.Context, request GetEnrollmentStatusRequest) (response GetEnrollmentStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEnrollmentStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEnrollmentStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEnrollmentStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEnrollmentStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEnrollmentStatusResponse")
	}
	return
}

// getEnrollmentStatus implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) getEnrollmentStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enrollmentStatus/{enrollmentStatusId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEnrollmentStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/EnrollmentStatus/GetEnrollmentStatus"
		err = common.PostProcessServiceError(err, "Optimizer", "GetEnrollmentStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProfile Gets the specified profile's information. Uses the profile's OCID to determine which profile to retrieve.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/GetProfile.go.html to see an example of how to use GetProfile API.
// A default retry strategy applies to this operation GetProfile()
func (client OptimizerClient) GetProfile(ctx context.Context, request GetProfileRequest) (response GetProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProfileResponse")
	}
	return
}

// getProfile implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) getProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/profiles/{profileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Profile/GetProfile"
		err = common.PostProcessServiceError(err, "Optimizer", "GetProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRecommendation Gets the recommendation for the specified OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/GetRecommendation.go.html to see an example of how to use GetRecommendation API.
// A default retry strategy applies to this operation GetRecommendation()
func (client OptimizerClient) GetRecommendation(ctx context.Context, request GetRecommendationRequest) (response GetRecommendationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRecommendation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRecommendationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRecommendationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRecommendationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRecommendationResponse")
	}
	return
}

// getRecommendation implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) getRecommendation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recommendations/{recommendationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRecommendationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Recommendation/GetRecommendation"
		err = common.PostProcessServiceError(err, "Optimizer", "GetRecommendation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResourceAction Gets the resource action that corresponds to the specified OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/GetResourceAction.go.html to see an example of how to use GetResourceAction API.
// A default retry strategy applies to this operation GetResourceAction()
func (client OptimizerClient) GetResourceAction(ctx context.Context, request GetResourceActionRequest) (response GetResourceActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResourceAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResourceActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResourceActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResourceActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResourceActionResponse")
	}
	return
}

// getResourceAction implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) getResourceAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceActions/{resourceActionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResourceActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/ResourceAction/GetResourceAction"
		err = common.PostProcessServiceError(err, "Optimizer", "GetResourceAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request associated with the specified ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client OptimizerClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client OptimizerClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "Optimizer", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCategories Lists the supported Cloud Advisor categories.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListCategories.go.html to see an example of how to use ListCategories API.
// A default retry strategy applies to this operation ListCategories()
func (client OptimizerClient) ListCategories(ctx context.Context, request ListCategoriesRequest) (response ListCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCategoriesResponse")
	}
	return
}

// listCategories implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/categories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCategoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/CategorySummary/ListCategories"
		err = common.PostProcessServiceError(err, "Optimizer", "ListCategories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEnrollmentStatuses Lists the Cloud Advisor enrollment statuses.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListEnrollmentStatuses.go.html to see an example of how to use ListEnrollmentStatuses API.
// A default retry strategy applies to this operation ListEnrollmentStatuses()
func (client OptimizerClient) ListEnrollmentStatuses(ctx context.Context, request ListEnrollmentStatusesRequest) (response ListEnrollmentStatusesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEnrollmentStatuses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEnrollmentStatusesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEnrollmentStatusesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEnrollmentStatusesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEnrollmentStatusesResponse")
	}
	return
}

// listEnrollmentStatuses implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listEnrollmentStatuses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/enrollmentStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEnrollmentStatusesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/EnrollmentStatusSummary/ListEnrollmentStatuses"
		err = common.PostProcessServiceError(err, "Optimizer", "ListEnrollmentStatuses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHistories Lists changes to the recommendations based on user activity.
// For example, lists when recommendations have been implemented, dismissed, postponed, or reactivated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListHistories.go.html to see an example of how to use ListHistories API.
// A default retry strategy applies to this operation ListHistories()
func (client OptimizerClient) ListHistories(ctx context.Context, request ListHistoriesRequest) (response ListHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHistoriesResponse")
	}
	return
}

// listHistories implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listHistories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/histories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHistoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/HistorySummary/ListHistories"
		err = common.PostProcessServiceError(err, "Optimizer", "ListHistories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProfileLevels Lists the existing profile levels.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListProfileLevels.go.html to see an example of how to use ListProfileLevels API.
// A default retry strategy applies to this operation ListProfileLevels()
func (client OptimizerClient) ListProfileLevels(ctx context.Context, request ListProfileLevelsRequest) (response ListProfileLevelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProfileLevels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProfileLevelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProfileLevelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProfileLevelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProfileLevelsResponse")
	}
	return
}

// listProfileLevels implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listProfileLevels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/profileLevels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProfileLevelsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/ProfileLevelSummary/ListProfileLevels"
		err = common.PostProcessServiceError(err, "Optimizer", "ListProfileLevels", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProfiles Lists the existing profiles.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListProfiles.go.html to see an example of how to use ListProfiles API.
// A default retry strategy applies to this operation ListProfiles()
func (client OptimizerClient) ListProfiles(ctx context.Context, request ListProfilesRequest) (response ListProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProfilesResponse")
	}
	return
}

// listProfiles implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/profiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/ProfileSummary/ListProfiles"
		err = common.PostProcessServiceError(err, "Optimizer", "ListProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecommendationStrategies Lists the existing strategies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListRecommendationStrategies.go.html to see an example of how to use ListRecommendationStrategies API.
// A default retry strategy applies to this operation ListRecommendationStrategies()
func (client OptimizerClient) ListRecommendationStrategies(ctx context.Context, request ListRecommendationStrategiesRequest) (response ListRecommendationStrategiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecommendationStrategies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecommendationStrategiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecommendationStrategiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecommendationStrategiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecommendationStrategiesResponse")
	}
	return
}

// listRecommendationStrategies implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listRecommendationStrategies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recommendationStrategies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRecommendationStrategiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/RecommendationStrategySummary/ListRecommendationStrategies"
		err = common.PostProcessServiceError(err, "Optimizer", "ListRecommendationStrategies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecommendations Lists the Cloud Advisor recommendations that are currently supported.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListRecommendations.go.html to see an example of how to use ListRecommendations API.
// A default retry strategy applies to this operation ListRecommendations()
func (client OptimizerClient) ListRecommendations(ctx context.Context, request ListRecommendationsRequest) (response ListRecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client OptimizerClient) listRecommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/RecommendationSummary/ListRecommendations"
		err = common.PostProcessServiceError(err, "Optimizer", "ListRecommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceActionQueryableFields Lists the fields that are indexed for querying and their associated value types.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListResourceActionQueryableFields.go.html to see an example of how to use ListResourceActionQueryableFields API.
// A default retry strategy applies to this operation ListResourceActionQueryableFields()
func (client OptimizerClient) ListResourceActionQueryableFields(ctx context.Context, request ListResourceActionQueryableFieldsRequest) (response ListResourceActionQueryableFieldsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceActionQueryableFields, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceActionQueryableFieldsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceActionQueryableFieldsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceActionQueryableFieldsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceActionQueryableFieldsResponse")
	}
	return
}

// listResourceActionQueryableFields implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listResourceActionQueryableFields(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceActions/actions/getQueryableFields", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceActionQueryableFieldsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/QueryableFieldSummary/ListResourceActionQueryableFields"
		err = common.PostProcessServiceError(err, "Optimizer", "ListResourceActionQueryableFields", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListResourceActions Lists the Cloud Advisor resource actions that are supported.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListResourceActions.go.html to see an example of how to use ListResourceActions API.
// A default retry strategy applies to this operation ListResourceActions()
func (client OptimizerClient) ListResourceActions(ctx context.Context, request ListResourceActionsRequest) (response ListResourceActionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResourceActions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResourceActionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResourceActionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResourceActionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResourceActionsResponse")
	}
	return
}

// listResourceActions implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) listResourceActions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceActions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResourceActionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/ResourceActionSummary/ListResourceActions"
		err = common.PostProcessServiceError(err, "Optimizer", "ListResourceActions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists errors associated with the specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client OptimizerClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client OptimizerClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Optimizer", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs associated with the specified work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client OptimizerClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client OptimizerClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Optimizer", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in the tenancy. The tenancy is the root compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client OptimizerClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client OptimizerClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "Optimizer", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEnrollmentStatus Updates the enrollment status of the tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/UpdateEnrollmentStatus.go.html to see an example of how to use UpdateEnrollmentStatus API.
// A default retry strategy applies to this operation UpdateEnrollmentStatus()
func (client OptimizerClient) UpdateEnrollmentStatus(ctx context.Context, request UpdateEnrollmentStatusRequest) (response UpdateEnrollmentStatusResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEnrollmentStatus, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEnrollmentStatusResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEnrollmentStatusResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEnrollmentStatusResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEnrollmentStatusResponse")
	}
	return
}

// updateEnrollmentStatus implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) updateEnrollmentStatus(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/enrollmentStatus/{enrollmentStatusId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEnrollmentStatusResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/EnrollmentStatus/UpdateEnrollmentStatus"
		err = common.PostProcessServiceError(err, "Optimizer", "UpdateEnrollmentStatus", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProfile Updates the specified profile. Uses the profile's OCID to determine which profile to update.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/UpdateProfile.go.html to see an example of how to use UpdateProfile API.
// A default retry strategy applies to this operation UpdateProfile()
func (client OptimizerClient) UpdateProfile(ctx context.Context, request UpdateProfileRequest) (response UpdateProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProfileResponse")
	}
	return
}

// updateProfile implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) updateProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/profiles/{profileId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Profile/UpdateProfile"
		err = common.PostProcessServiceError(err, "Optimizer", "UpdateProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRecommendation Updates the recommendation that corresponds to the specified OCID.
// Use this operation to implement the following actions:
//   - Postpone recommendation
//   - Dismiss recommendation
//   - Reactivate recommendation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/UpdateRecommendation.go.html to see an example of how to use UpdateRecommendation API.
func (client OptimizerClient) UpdateRecommendation(ctx context.Context, request UpdateRecommendationRequest) (response UpdateRecommendationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRecommendation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRecommendationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRecommendationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRecommendationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRecommendationResponse")
	}
	return
}

// updateRecommendation implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) updateRecommendation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/recommendations/{recommendationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRecommendationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/Recommendation/UpdateRecommendation"
		err = common.PostProcessServiceError(err, "Optimizer", "UpdateRecommendation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResourceAction Updates the resource action that corresponds to the specified OCID.
// Use this operation to implement the following actions:
//   - Postpone resource action
//   - Ignore resource action
//   - Reactivate resource action
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/UpdateResourceAction.go.html to see an example of how to use UpdateResourceAction API.
// A default retry strategy applies to this operation UpdateResourceAction()
func (client OptimizerClient) UpdateResourceAction(ctx context.Context, request UpdateResourceActionRequest) (response UpdateResourceActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResourceAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResourceActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResourceActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResourceActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResourceActionResponse")
	}
	return
}

// updateResourceAction implements the OCIOperation interface (enables retrying operations)
func (client OptimizerClient) updateResourceAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/resourceActions/{resourceActionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateResourceActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/advisor/20200606/ResourceAction/UpdateResourceAction"
		err = common.PostProcessServiceError(err, "Optimizer", "UpdateResourceAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
