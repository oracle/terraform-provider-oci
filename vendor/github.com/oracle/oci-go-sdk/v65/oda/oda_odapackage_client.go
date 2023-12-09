// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OdapackageClient a client for Odapackage
type OdapackageClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOdapackageClientWithConfigurationProvider Creates a new default Odapackage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOdapackageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OdapackageClient, err error) {
	if enabled := common.CheckForEnabledServices("oda"); !enabled {
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
	return newOdapackageClientFromBaseClient(baseClient, provider)
}

// NewOdapackageClientWithOboToken Creates a new default Odapackage client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOdapackageClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OdapackageClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOdapackageClientFromBaseClient(baseClient, configProvider)
}

func newOdapackageClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OdapackageClient, err error) {
	// Odapackage service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Odapackage"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OdapackageClient{BaseClient: baseClient}
	client.BasePath = "20190506"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OdapackageClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("oda", "https://digitalassistant-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OdapackageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OdapackageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateImportedPackage Starts an asynchronous job to import a package into a Digital Assistant instance.
// To monitor the status of the job, take the `opc-work-request-id` response
// header value and use it to call `GET /workRequests/{workRequestId}`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/CreateImportedPackage.go.html to see an example of how to use CreateImportedPackage API.
// A default retry strategy applies to this operation CreateImportedPackage()
func (client OdapackageClient) CreateImportedPackage(ctx context.Context, request CreateImportedPackageRequest) (response CreateImportedPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createImportedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateImportedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateImportedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateImportedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateImportedPackageResponse")
	}
	return
}

// createImportedPackage implements the OCIOperation interface (enables retrying operations)
func (client OdapackageClient) createImportedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/importedPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateImportedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Odapackage", "CreateImportedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteImportedPackage Starts an asynchronous job to delete a package from a Digital Assistant instance.
// To monitor the status of the job, take the `opc-work-request-id` response
// header value and use it to call `GET /workRequests/{workRequestId}`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/DeleteImportedPackage.go.html to see an example of how to use DeleteImportedPackage API.
// A default retry strategy applies to this operation DeleteImportedPackage()
func (client OdapackageClient) DeleteImportedPackage(ctx context.Context, request DeleteImportedPackageRequest) (response DeleteImportedPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteImportedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteImportedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteImportedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteImportedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteImportedPackageResponse")
	}
	return
}

// deleteImportedPackage implements the OCIOperation interface (enables retrying operations)
func (client OdapackageClient) deleteImportedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}/importedPackages/{packageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteImportedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/ImportedPackage/DeleteImportedPackage"
		err = common.PostProcessServiceError(err, "Odapackage", "DeleteImportedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetImportedPackage Returns a list of summaries for imported packages in the instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/GetImportedPackage.go.html to see an example of how to use GetImportedPackage API.
// A default retry strategy applies to this operation GetImportedPackage()
func (client OdapackageClient) GetImportedPackage(ctx context.Context, request GetImportedPackageRequest) (response GetImportedPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getImportedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetImportedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetImportedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetImportedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetImportedPackageResponse")
	}
	return
}

// getImportedPackage implements the OCIOperation interface (enables retrying operations)
func (client OdapackageClient) getImportedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/importedPackages/{packageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetImportedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/ImportedPackage/GetImportedPackage"
		err = common.PostProcessServiceError(err, "Odapackage", "GetImportedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPackage Returns details about a package, and how to import it.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/GetPackage.go.html to see an example of how to use GetPackage API.
// A default retry strategy applies to this operation GetPackage()
func (client OdapackageClient) GetPackage(ctx context.Context, request GetPackageRequest) (response GetPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPackageResponse")
	}
	return
}

// getPackage implements the OCIOperation interface (enables retrying operations)
func (client OdapackageClient) getPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/packages/{packageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Package/GetPackage"
		err = common.PostProcessServiceError(err, "Odapackage", "GetPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListImportedPackages Returns a list of summaries for imported packages in the instance.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListImportedPackages.go.html to see an example of how to use ListImportedPackages API.
// A default retry strategy applies to this operation ListImportedPackages()
func (client OdapackageClient) ListImportedPackages(ctx context.Context, request ListImportedPackagesRequest) (response ListImportedPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listImportedPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListImportedPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListImportedPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListImportedPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListImportedPackagesResponse")
	}
	return
}

// listImportedPackages implements the OCIOperation interface (enables retrying operations)
func (client OdapackageClient) listImportedPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/importedPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListImportedPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/ImportedPackageSummary/ListImportedPackages"
		err = common.PostProcessServiceError(err, "Odapackage", "ListImportedPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPackages Returns a page of summaries for packages that are available for import. The optional odaInstanceId query
// parameter can be used to filter packages that are available for import by a specific instance. If odaInstanceId
// query parameter is not provided, the returned list will
// include packages available within the region indicated by the request URL. The optional resourceType query
// param may be specified to filter packages that contain the indicated resource type. If no resourceType query
// param is given, packages containing all resource types will be returned. The optional name query parameter can
// be used to limit the list to packages whose name matches the given name. The optional displayName query
// parameter can be used to limit the list to packages whose displayName matches the given name. The optional
// isLatestVersionOnly query parameter can be used to limit the returned list to include only the latest version
// of any given package. If not specified, all versions of any otherwise matching package will be returned.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListPackages.go.html to see an example of how to use ListPackages API.
// A default retry strategy applies to this operation ListPackages()
func (client OdapackageClient) ListPackages(ctx context.Context, request ListPackagesRequest) (response ListPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPackagesResponse")
	}
	return
}

// listPackages implements the OCIOperation interface (enables retrying operations)
func (client OdapackageClient) listPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/packages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/PackageSummary/ListPackages"
		err = common.PostProcessServiceError(err, "Odapackage", "ListPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateImportedPackage Starts an asynchronous job to update a package within a Digital Assistant instance.
// To monitor the status of the job, take the `opc-work-request-id` response
// header value and use it to call `GET /workRequests/{workRequestId}`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/UpdateImportedPackage.go.html to see an example of how to use UpdateImportedPackage API.
// A default retry strategy applies to this operation UpdateImportedPackage()
func (client OdapackageClient) UpdateImportedPackage(ctx context.Context, request UpdateImportedPackageRequest) (response UpdateImportedPackageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateImportedPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateImportedPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateImportedPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateImportedPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateImportedPackageResponse")
	}
	return
}

// updateImportedPackage implements the OCIOperation interface (enables retrying operations)
func (client OdapackageClient) updateImportedPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/importedPackages/{packageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateImportedPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Odapackage", "UpdateImportedPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
