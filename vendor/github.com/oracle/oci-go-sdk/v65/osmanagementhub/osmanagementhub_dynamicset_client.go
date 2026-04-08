// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DynamicSetClient a client for DynamicSet
type DynamicSetClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDynamicSetClientWithConfigurationProvider Creates a new default DynamicSet client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDynamicSetClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DynamicSetClient, err error) {
	if enabled := common.CheckForEnabledServices("osmanagementhub"); !enabled {
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
	return newDynamicSetClientFromBaseClient(baseClient, provider)
}

// NewDynamicSetClientWithOboToken Creates a new default DynamicSet client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDynamicSetClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DynamicSetClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDynamicSetClientFromBaseClient(baseClient, configProvider)
}

func newDynamicSetClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DynamicSetClient, err error) {
	// DynamicSet service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DynamicSet"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DynamicSetClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DynamicSetClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DynamicSetClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DynamicSetClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeDynamicSetCompartment Move the specified Dynamic Set to a different compartment
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ChangeDynamicSetCompartment.go.html to see an example of how to use ChangeDynamicSetCompartment API.
// A default retry strategy applies to this operation ChangeDynamicSetCompartment()
func (client DynamicSetClient) ChangeDynamicSetCompartment(ctx context.Context, request ChangeDynamicSetCompartmentRequest) (response ChangeDynamicSetCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDynamicSetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDynamicSetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDynamicSetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDynamicSetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDynamicSetCompartmentResponse")
	}
	return
}

// changeDynamicSetCompartment implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) changeDynamicSetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicSets/{dynamicSetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDynamicSetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "ChangeDynamicSetCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/ChangeDynamicSetCompartment"
		err = common.PostProcessServiceError(err, "DynamicSet", "ChangeDynamicSetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDynamicSet Creates a new dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateDynamicSet.go.html to see an example of how to use CreateDynamicSet API.
// A default retry strategy applies to this operation CreateDynamicSet()
func (client DynamicSetClient) CreateDynamicSet(ctx context.Context, request CreateDynamicSetRequest) (response CreateDynamicSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDynamicSetResponse")
	}
	return
}

// createDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) createDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "CreateDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/CreateDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "CreateDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDynamicSet Deletes the specific dynamic set
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteDynamicSet.go.html to see an example of how to use DeleteDynamicSet API.
// A default retry strategy applies to this operation DeleteDynamicSet()
func (client DynamicSetClient) DeleteDynamicSet(ctx context.Context, request DeleteDynamicSetRequest) (response DeleteDynamicSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDynamicSetResponse")
	}
	return
}

// deleteDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) deleteDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dynamicSets/{dynamicSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "DeleteDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/DeleteDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "DeleteDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDynamicSet Gets information about the specified dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetDynamicSet.go.html to see an example of how to use GetDynamicSet API.
// A default retry strategy applies to this operation GetDynamicSet()
func (client DynamicSetClient) GetDynamicSet(ctx context.Context, request GetDynamicSetRequest) (response GetDynamicSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDynamicSetResponse")
	}
	return
}

// getDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) getDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dynamicSets/{dynamicSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "GetDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/GetDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "GetDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallPackagesOnDynamicSet Installs specified software packages on all managed instances in the dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/InstallPackagesOnDynamicSet.go.html to see an example of how to use InstallPackagesOnDynamicSet API.
// A default retry strategy applies to this operation InstallPackagesOnDynamicSet()
func (client DynamicSetClient) InstallPackagesOnDynamicSet(ctx context.Context, request InstallPackagesOnDynamicSetRequest) (response InstallPackagesOnDynamicSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installPackagesOnDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallPackagesOnDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallPackagesOnDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallPackagesOnDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallPackagesOnDynamicSetResponse")
	}
	return
}

// installPackagesOnDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) installPackagesOnDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicSets/{dynamicSetId}/actions/installPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallPackagesOnDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "InstallPackagesOnDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/InstallPackagesOnDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "InstallPackagesOnDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDynamicSets Lists dynamic sets that match the specified compartment or dynamic set OCID. Filter the list against a variety of criteria including but not limited to its name, status, architecture, and OS version.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListDynamicSets.go.html to see an example of how to use ListDynamicSets API.
// A default retry strategy applies to this operation ListDynamicSets()
func (client DynamicSetClient) ListDynamicSets(ctx context.Context, request ListDynamicSetsRequest) (response ListDynamicSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDynamicSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDynamicSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDynamicSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDynamicSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDynamicSetsResponse")
	}
	return
}

// listDynamicSets implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) listDynamicSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dynamicSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDynamicSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "ListDynamicSets")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/ListDynamicSets"
		err = common.PostProcessServiceError(err, "DynamicSet", "ListDynamicSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstancesInDynamicSet Retrieves a list of managed instances associated with a specified dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagedInstancesInDynamicSet.go.html to see an example of how to use ListManagedInstancesInDynamicSet API.
// A default retry strategy applies to this operation ListManagedInstancesInDynamicSet()
func (client DynamicSetClient) ListManagedInstancesInDynamicSet(ctx context.Context, request ListManagedInstancesInDynamicSetRequest) (response ListManagedInstancesInDynamicSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstancesInDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstancesInDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstancesInDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstancesInDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstancesInDynamicSetResponse")
	}
	return
}

// listManagedInstancesInDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) listManagedInstancesInDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dynamicSets/{dynamicSetId}/managedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstancesInDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "ListManagedInstancesInDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/ListManagedInstancesInDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "ListManagedInstancesInDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PreviewManagedInstances Preview a dynamic set
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/PreviewManagedInstances.go.html to see an example of how to use PreviewManagedInstances API.
// A default retry strategy applies to this operation PreviewManagedInstances()
func (client DynamicSetClient) PreviewManagedInstances(ctx context.Context, request PreviewManagedInstancesRequest) (response PreviewManagedInstancesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.previewManagedInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PreviewManagedInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PreviewManagedInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PreviewManagedInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PreviewManagedInstancesResponse")
	}
	return
}

// previewManagedInstances implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) previewManagedInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicSets/actions/previewManagedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PreviewManagedInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "PreviewManagedInstances")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/PreviewManagedInstances"
		err = common.PostProcessServiceError(err, "DynamicSet", "PreviewManagedInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RebootDynamicSet Initiates a reboot of all managed instances within the specified dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RebootDynamicSet.go.html to see an example of how to use RebootDynamicSet API.
// A default retry strategy applies to this operation RebootDynamicSet()
func (client DynamicSetClient) RebootDynamicSet(ctx context.Context, request RebootDynamicSetRequest) (response RebootDynamicSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rebootDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RebootDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RebootDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RebootDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RebootDynamicSetResponse")
	}
	return
}

// rebootDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) rebootDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicSets/{dynamicSetId}/actions/reboot", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RebootDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "RebootDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/RebootDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "RebootDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePackagesFromDynamicSet Removes specified software packages from all managed instances in the dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/RemovePackagesFromDynamicSet.go.html to see an example of how to use RemovePackagesFromDynamicSet API.
// A default retry strategy applies to this operation RemovePackagesFromDynamicSet()
func (client DynamicSetClient) RemovePackagesFromDynamicSet(ctx context.Context, request RemovePackagesFromDynamicSetRequest) (response RemovePackagesFromDynamicSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removePackagesFromDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemovePackagesFromDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemovePackagesFromDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemovePackagesFromDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemovePackagesFromDynamicSetResponse")
	}
	return
}

// removePackagesFromDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) removePackagesFromDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicSets/{dynamicSetId}/actions/removePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemovePackagesFromDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "RemovePackagesFromDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/RemovePackagesFromDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "RemovePackagesFromDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDynamicSet Updates the specified dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateDynamicSet.go.html to see an example of how to use UpdateDynamicSet API.
// A default retry strategy applies to this operation UpdateDynamicSet()
func (client DynamicSetClient) UpdateDynamicSet(ctx context.Context, request UpdateDynamicSetRequest) (response UpdateDynamicSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDynamicSetResponse")
	}
	return
}

// updateDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) updateDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dynamicSets/{dynamicSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "UpdateDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/UpdateDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "UpdateDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePackagesOnDynamicSet Updates all installed software packages on managed instances in the dynamic set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdatePackagesOnDynamicSet.go.html to see an example of how to use UpdatePackagesOnDynamicSet API.
// A default retry strategy applies to this operation UpdatePackagesOnDynamicSet()
func (client DynamicSetClient) UpdatePackagesOnDynamicSet(ctx context.Context, request UpdatePackagesOnDynamicSetRequest) (response UpdatePackagesOnDynamicSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updatePackagesOnDynamicSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePackagesOnDynamicSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePackagesOnDynamicSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePackagesOnDynamicSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePackagesOnDynamicSetResponse")
	}
	return
}

// updatePackagesOnDynamicSet implements the OCIOperation interface (enables retrying operations)
func (client DynamicSetClient) updatePackagesOnDynamicSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicSets/{dynamicSetId}/actions/updatePackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePackagesOnDynamicSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "dynamicSet", "UpdatePackagesOnDynamicSet")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/DynamicSet/UpdatePackagesOnDynamicSet"
		err = common.PostProcessServiceError(err, "DynamicSet", "UpdatePackagesOnDynamicSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
