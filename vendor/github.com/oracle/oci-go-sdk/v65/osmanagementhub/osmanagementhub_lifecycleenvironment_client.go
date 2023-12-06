// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// LifecycleEnvironmentClient a client for LifecycleEnvironment
type LifecycleEnvironmentClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLifecycleEnvironmentClientWithConfigurationProvider Creates a new default LifecycleEnvironment client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLifecycleEnvironmentClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LifecycleEnvironmentClient, err error) {
	if enabled := common.CheckForEnabledServices("osmanagementhub"); !enabled {
		return client, fmt.Errorf("the Alloy configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local alloy_config file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newLifecycleEnvironmentClientFromBaseClient(baseClient, provider)
}

// NewLifecycleEnvironmentClientWithOboToken Creates a new default LifecycleEnvironment client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLifecycleEnvironmentClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LifecycleEnvironmentClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLifecycleEnvironmentClientFromBaseClient(baseClient, configProvider)
}

func newLifecycleEnvironmentClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LifecycleEnvironmentClient, err error) {
	// LifecycleEnvironment service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LifecycleEnvironment"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LifecycleEnvironmentClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LifecycleEnvironmentClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LifecycleEnvironmentClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LifecycleEnvironmentClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachManagedInstancesToLifecycleStage Attach(add) managed instances to a lifecycle stage.
// Once added operations can be applied to all managed instances in the lifecycle stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/AttachManagedInstancesToLifecycleStage.go.html to see an example of how to use AttachManagedInstancesToLifecycleStage API.
// A default retry strategy applies to this operation AttachManagedInstancesToLifecycleStage()
func (client LifecycleEnvironmentClient) AttachManagedInstancesToLifecycleStage(ctx context.Context, request AttachManagedInstancesToLifecycleStageRequest) (response AttachManagedInstancesToLifecycleStageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.attachManagedInstancesToLifecycleStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachManagedInstancesToLifecycleStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachManagedInstancesToLifecycleStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachManagedInstancesToLifecycleStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachManagedInstancesToLifecycleStageResponse")
	}
	return
}

// attachManagedInstancesToLifecycleStage implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) attachManagedInstancesToLifecycleStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/lifecycleStages/{lifecycleStageId}/actions/attachManagedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachManagedInstancesToLifecycleStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleStage/AttachManagedInstancesToLifecycleStage"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "AttachManagedInstancesToLifecycleStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLifecycleEnvironment Creates a new lifecycle environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateLifecycleEnvironment.go.html to see an example of how to use CreateLifecycleEnvironment API.
// A default retry strategy applies to this operation CreateLifecycleEnvironment()
func (client LifecycleEnvironmentClient) CreateLifecycleEnvironment(ctx context.Context, request CreateLifecycleEnvironmentRequest) (response CreateLifecycleEnvironmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createLifecycleEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLifecycleEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLifecycleEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLifecycleEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLifecycleEnvironmentResponse")
	}
	return
}

// createLifecycleEnvironment implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) createLifecycleEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/lifecycleEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLifecycleEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleEnvironment/CreateLifecycleEnvironment"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "CreateLifecycleEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLifecycleEnvironment Deletes a lifecycle environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteLifecycleEnvironment.go.html to see an example of how to use DeleteLifecycleEnvironment API.
// A default retry strategy applies to this operation DeleteLifecycleEnvironment()
func (client LifecycleEnvironmentClient) DeleteLifecycleEnvironment(ctx context.Context, request DeleteLifecycleEnvironmentRequest) (response DeleteLifecycleEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLifecycleEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLifecycleEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLifecycleEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLifecycleEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLifecycleEnvironmentResponse")
	}
	return
}

// deleteLifecycleEnvironment implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) deleteLifecycleEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/lifecycleEnvironments/{lifecycleEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLifecycleEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleEnvironment/DeleteLifecycleEnvironment"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "DeleteLifecycleEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachManagedInstancesFromLifecycleStage Detach(remove) managed instance from a lifecycle stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DetachManagedInstancesFromLifecycleStage.go.html to see an example of how to use DetachManagedInstancesFromLifecycleStage API.
// A default retry strategy applies to this operation DetachManagedInstancesFromLifecycleStage()
func (client LifecycleEnvironmentClient) DetachManagedInstancesFromLifecycleStage(ctx context.Context, request DetachManagedInstancesFromLifecycleStageRequest) (response DetachManagedInstancesFromLifecycleStageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.detachManagedInstancesFromLifecycleStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachManagedInstancesFromLifecycleStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachManagedInstancesFromLifecycleStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachManagedInstancesFromLifecycleStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachManagedInstancesFromLifecycleStageResponse")
	}
	return
}

// detachManagedInstancesFromLifecycleStage implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) detachManagedInstancesFromLifecycleStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/lifecycleStages/{lifecycleStageId}/actions/detachManagedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachManagedInstancesFromLifecycleStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleStage/DetachManagedInstancesFromLifecycleStage"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "DetachManagedInstancesFromLifecycleStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLifecycleEnvironment Gets information about the specified lifecycle environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetLifecycleEnvironment.go.html to see an example of how to use GetLifecycleEnvironment API.
// A default retry strategy applies to this operation GetLifecycleEnvironment()
func (client LifecycleEnvironmentClient) GetLifecycleEnvironment(ctx context.Context, request GetLifecycleEnvironmentRequest) (response GetLifecycleEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLifecycleEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLifecycleEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLifecycleEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLifecycleEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLifecycleEnvironmentResponse")
	}
	return
}

// getLifecycleEnvironment implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) getLifecycleEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lifecycleEnvironments/{lifecycleEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLifecycleEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleEnvironment/GetLifecycleEnvironment"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "GetLifecycleEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLifecycleStage Gets information about the specified lifecycle stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetLifecycleStage.go.html to see an example of how to use GetLifecycleStage API.
// A default retry strategy applies to this operation GetLifecycleStage()
func (client LifecycleEnvironmentClient) GetLifecycleStage(ctx context.Context, request GetLifecycleStageRequest) (response GetLifecycleStageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLifecycleStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLifecycleStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLifecycleStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLifecycleStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLifecycleStageResponse")
	}
	return
}

// getLifecycleStage implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) getLifecycleStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lifecycleStages/{lifecycleStageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLifecycleStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleStage/GetLifecycleStage"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "GetLifecycleStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLifecycleEnvironments Lists lifecycle environments that match the specified compartment or lifecycle environment OCID. Filter the list
// against a variety of criteria including but not limited to its name, status, architecture, and OS family.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListLifecycleEnvironments.go.html to see an example of how to use ListLifecycleEnvironments API.
// A default retry strategy applies to this operation ListLifecycleEnvironments()
func (client LifecycleEnvironmentClient) ListLifecycleEnvironments(ctx context.Context, request ListLifecycleEnvironmentsRequest) (response ListLifecycleEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLifecycleEnvironments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLifecycleEnvironmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLifecycleEnvironmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLifecycleEnvironmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLifecycleEnvironmentsResponse")
	}
	return
}

// listLifecycleEnvironments implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) listLifecycleEnvironments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lifecycleEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLifecycleEnvironmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleEnvironment/ListLifecycleEnvironments"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "ListLifecycleEnvironments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLifecycleStageInstalledPackages Lists installed packages on managed instances in a specified lifecycle stage. Filter the list against a variety
// of criteria including but not limited to the package name.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListLifecycleStageInstalledPackages.go.html to see an example of how to use ListLifecycleStageInstalledPackages API.
// A default retry strategy applies to this operation ListLifecycleStageInstalledPackages()
func (client LifecycleEnvironmentClient) ListLifecycleStageInstalledPackages(ctx context.Context, request ListLifecycleStageInstalledPackagesRequest) (response ListLifecycleStageInstalledPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLifecycleStageInstalledPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLifecycleStageInstalledPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLifecycleStageInstalledPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLifecycleStageInstalledPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLifecycleStageInstalledPackagesResponse")
	}
	return
}

// listLifecycleStageInstalledPackages implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) listLifecycleStageInstalledPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lifecycleStages/{lifecycleStageId}/installedPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLifecycleStageInstalledPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleStage/ListLifecycleStageInstalledPackages"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "ListLifecycleStageInstalledPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLifecycleStages Lists lifecycle stages that match the specified compartment or lifecycle stage OCID. Filter the list against
// a variety of criteria including but not limited to its name, status, architecture, and OS family.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListLifecycleStages.go.html to see an example of how to use ListLifecycleStages API.
// A default retry strategy applies to this operation ListLifecycleStages()
func (client LifecycleEnvironmentClient) ListLifecycleStages(ctx context.Context, request ListLifecycleStagesRequest) (response ListLifecycleStagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLifecycleStages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLifecycleStagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLifecycleStagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLifecycleStagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLifecycleStagesResponse")
	}
	return
}

// listLifecycleStages implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) listLifecycleStages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/lifecycleStages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLifecycleStagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleStage/ListLifecycleStages"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "ListLifecycleStages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PromoteSoftwareSourceToLifecycleStage Updates the versioned custom software source content
// for specified lifecycle stage.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/PromoteSoftwareSourceToLifecycleStage.go.html to see an example of how to use PromoteSoftwareSourceToLifecycleStage API.
// A default retry strategy applies to this operation PromoteSoftwareSourceToLifecycleStage()
func (client LifecycleEnvironmentClient) PromoteSoftwareSourceToLifecycleStage(ctx context.Context, request PromoteSoftwareSourceToLifecycleStageRequest) (response PromoteSoftwareSourceToLifecycleStageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.promoteSoftwareSourceToLifecycleStage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PromoteSoftwareSourceToLifecycleStageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PromoteSoftwareSourceToLifecycleStageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PromoteSoftwareSourceToLifecycleStageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PromoteSoftwareSourceToLifecycleStageResponse")
	}
	return
}

// promoteSoftwareSourceToLifecycleStage implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) promoteSoftwareSourceToLifecycleStage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/lifecycleStages/{lifecycleStageId}/actions/promoteSoftwareSource", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PromoteSoftwareSourceToLifecycleStageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleStage/PromoteSoftwareSourceToLifecycleStage"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "PromoteSoftwareSourceToLifecycleStage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLifecycleEnvironment Updates the specified lifecycle environment's name, description, stages, or tags.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateLifecycleEnvironment.go.html to see an example of how to use UpdateLifecycleEnvironment API.
// A default retry strategy applies to this operation UpdateLifecycleEnvironment()
func (client LifecycleEnvironmentClient) UpdateLifecycleEnvironment(ctx context.Context, request UpdateLifecycleEnvironmentRequest) (response UpdateLifecycleEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLifecycleEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLifecycleEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLifecycleEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLifecycleEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLifecycleEnvironmentResponse")
	}
	return
}

// updateLifecycleEnvironment implements the OCIOperation interface (enables retrying operations)
func (client LifecycleEnvironmentClient) updateLifecycleEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/lifecycleEnvironments/{lifecycleEnvironmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLifecycleEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/LifecycleEnvironment/UpdateLifecycleEnvironment"
		err = common.PostProcessServiceError(err, "LifecycleEnvironment", "UpdateLifecycleEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
