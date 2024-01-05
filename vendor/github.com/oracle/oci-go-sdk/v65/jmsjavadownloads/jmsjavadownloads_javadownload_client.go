// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// JavaDownloadClient a client for JavaDownload
type JavaDownloadClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewJavaDownloadClientWithConfigurationProvider Creates a new default JavaDownload client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewJavaDownloadClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client JavaDownloadClient, err error) {
	if enabled := common.CheckForEnabledServices("jmsjavadownloads"); !enabled {
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
	return newJavaDownloadClientFromBaseClient(baseClient, provider)
}

// NewJavaDownloadClientWithOboToken Creates a new default JavaDownload client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewJavaDownloadClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client JavaDownloadClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newJavaDownloadClientFromBaseClient(baseClient, configProvider)
}

func newJavaDownloadClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client JavaDownloadClient, err error) {
	// JavaDownload service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("JavaDownload"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = JavaDownloadClient{BaseClient: baseClient}
	client.BasePath = ""
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *JavaDownloadClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("jmsjavadownloads", "https://javamanagementservice-download.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *JavaDownloadClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *JavaDownloadClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Cancels the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client JavaDownloadClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client JavaDownloadClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/20230601/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "JavaDownload", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJavaDownloadReport Create a new report in the specified format containing the download details
// for the tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/CreateJavaDownloadReport.go.html to see an example of how to use CreateJavaDownloadReport API.
// A default retry strategy applies to this operation CreateJavaDownloadReport()
func (client JavaDownloadClient) CreateJavaDownloadReport(ctx context.Context, request CreateJavaDownloadReportRequest) (response CreateJavaDownloadReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createJavaDownloadReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJavaDownloadReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJavaDownloadReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJavaDownloadReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJavaDownloadReportResponse")
	}
	return
}

// createJavaDownloadReport implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) createJavaDownloadReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20230601/javaDownloadReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJavaDownloadReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadReport/CreateJavaDownloadReport"
		err = common.PostProcessServiceError(err, "JavaDownload", "CreateJavaDownloadReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJavaDownloadToken Creates a new JavaDownloadToken in the tenancy with specified attributes.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/CreateJavaDownloadToken.go.html to see an example of how to use CreateJavaDownloadToken API.
// A default retry strategy applies to this operation CreateJavaDownloadToken()
func (client JavaDownloadClient) CreateJavaDownloadToken(ctx context.Context, request CreateJavaDownloadTokenRequest) (response CreateJavaDownloadTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createJavaDownloadToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJavaDownloadTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJavaDownloadTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJavaDownloadTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJavaDownloadTokenResponse")
	}
	return
}

// createJavaDownloadToken implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) createJavaDownloadToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20230601/javaDownloadTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJavaDownloadTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadToken/CreateJavaDownloadToken"
		err = common.PostProcessServiceError(err, "JavaDownload", "CreateJavaDownloadToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJavaLicenseAcceptanceRecord Creates a Java license acceptance record for the specified license type in a tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/CreateJavaLicenseAcceptanceRecord.go.html to see an example of how to use CreateJavaLicenseAcceptanceRecord API.
// A default retry strategy applies to this operation CreateJavaLicenseAcceptanceRecord()
func (client JavaDownloadClient) CreateJavaLicenseAcceptanceRecord(ctx context.Context, request CreateJavaLicenseAcceptanceRecordRequest) (response CreateJavaLicenseAcceptanceRecordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createJavaLicenseAcceptanceRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJavaLicenseAcceptanceRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJavaLicenseAcceptanceRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJavaLicenseAcceptanceRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJavaLicenseAcceptanceRecordResponse")
	}
	return
}

// createJavaLicenseAcceptanceRecord implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) createJavaLicenseAcceptanceRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20230601/javaLicenseAcceptanceRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJavaLicenseAcceptanceRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaLicenseAcceptanceRecord/CreateJavaLicenseAcceptanceRecord"
		err = common.PostProcessServiceError(err, "JavaDownload", "CreateJavaLicenseAcceptanceRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJavaDownloadReport Deletes a JavaDownloadReport resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/DeleteJavaDownloadReport.go.html to see an example of how to use DeleteJavaDownloadReport API.
// A default retry strategy applies to this operation DeleteJavaDownloadReport()
func (client JavaDownloadClient) DeleteJavaDownloadReport(ctx context.Context, request DeleteJavaDownloadReportRequest) (response DeleteJavaDownloadReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJavaDownloadReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJavaDownloadReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJavaDownloadReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJavaDownloadReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJavaDownloadReportResponse")
	}
	return
}

// deleteJavaDownloadReport implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) deleteJavaDownloadReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/20230601/javaDownloadReports/{javaDownloadReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJavaDownloadReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadReport/DeleteJavaDownloadReport"
		err = common.PostProcessServiceError(err, "JavaDownload", "DeleteJavaDownloadReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJavaDownloadToken Deletes a JavaDownloadToken resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/DeleteJavaDownloadToken.go.html to see an example of how to use DeleteJavaDownloadToken API.
// A default retry strategy applies to this operation DeleteJavaDownloadToken()
func (client JavaDownloadClient) DeleteJavaDownloadToken(ctx context.Context, request DeleteJavaDownloadTokenRequest) (response DeleteJavaDownloadTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJavaDownloadToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJavaDownloadTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJavaDownloadTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJavaDownloadTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJavaDownloadTokenResponse")
	}
	return
}

// deleteJavaDownloadToken implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) deleteJavaDownloadToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/20230601/javaDownloadTokens/{javaDownloadTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJavaDownloadTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadToken/DeleteJavaDownloadToken"
		err = common.PostProcessServiceError(err, "JavaDownload", "DeleteJavaDownloadToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJavaLicenseAcceptanceRecord Deletes a Java license acceptance record with the specified identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/DeleteJavaLicenseAcceptanceRecord.go.html to see an example of how to use DeleteJavaLicenseAcceptanceRecord API.
// A default retry strategy applies to this operation DeleteJavaLicenseAcceptanceRecord()
func (client JavaDownloadClient) DeleteJavaLicenseAcceptanceRecord(ctx context.Context, request DeleteJavaLicenseAcceptanceRecordRequest) (response DeleteJavaLicenseAcceptanceRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJavaLicenseAcceptanceRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJavaLicenseAcceptanceRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJavaLicenseAcceptanceRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJavaLicenseAcceptanceRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJavaLicenseAcceptanceRecordResponse")
	}
	return
}

// deleteJavaLicenseAcceptanceRecord implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) deleteJavaLicenseAcceptanceRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/20230601/javaLicenseAcceptanceRecords/{javaLicenseAcceptanceRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJavaLicenseAcceptanceRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaLicenseAcceptanceRecord/DeleteJavaLicenseAcceptanceRecord"
		err = common.PostProcessServiceError(err, "JavaDownload", "DeleteJavaLicenseAcceptanceRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateArtifactDownloadUrl Generates a short-lived download URL and returns it in the response payload.
// The returned URL can then be used for downloading the specific Java runtime artifact.
// Use the GetJavaRelease API
// to get information about available artifacts for a specific release. Each such artifact is uniquely identified by an `artifactId`.
// Refer JavaArtifact for more details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GenerateArtifactDownloadUrl.go.html to see an example of how to use GenerateArtifactDownloadUrl API.
// A default retry strategy applies to this operation GenerateArtifactDownloadUrl()
func (client JavaDownloadClient) GenerateArtifactDownloadUrl(ctx context.Context, request GenerateArtifactDownloadUrlRequest) (response GenerateArtifactDownloadUrlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateArtifactDownloadUrl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateArtifactDownloadUrlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateArtifactDownloadUrlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateArtifactDownloadUrlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateArtifactDownloadUrlResponse")
	}
	return
}

// generateArtifactDownloadUrl implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) generateArtifactDownloadUrl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20230601/actions/generateArtifactDownloadUrl", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateArtifactDownloadUrlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/DownloadUrl/GenerateArtifactDownloadUrl"
		err = common.PostProcessServiceError(err, "JavaDownload", "GenerateArtifactDownloadUrl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaDownloadReport Gets a JavaDownloadReport by the specified identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GetJavaDownloadReport.go.html to see an example of how to use GetJavaDownloadReport API.
// A default retry strategy applies to this operation GetJavaDownloadReport()
func (client JavaDownloadClient) GetJavaDownloadReport(ctx context.Context, request GetJavaDownloadReportRequest) (response GetJavaDownloadReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaDownloadReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaDownloadReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaDownloadReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaDownloadReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaDownloadReportResponse")
	}
	return
}

// getJavaDownloadReport implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) getJavaDownloadReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaDownloadReports/{javaDownloadReportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaDownloadReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadReport/GetJavaDownloadReport"
		err = common.PostProcessServiceError(err, "JavaDownload", "GetJavaDownloadReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaDownloadReportContent Retrieve a Java download report with the specified identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GetJavaDownloadReportContent.go.html to see an example of how to use GetJavaDownloadReportContent API.
// A default retry strategy applies to this operation GetJavaDownloadReportContent()
func (client JavaDownloadClient) GetJavaDownloadReportContent(ctx context.Context, request GetJavaDownloadReportContentRequest) (response GetJavaDownloadReportContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaDownloadReportContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaDownloadReportContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaDownloadReportContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaDownloadReportContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaDownloadReportContentResponse")
	}
	return
}

// getJavaDownloadReportContent implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) getJavaDownloadReportContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaDownloadReports/{javaDownloadReportId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaDownloadReportContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadReport/GetJavaDownloadReportContent"
		err = common.PostProcessServiceError(err, "JavaDownload", "GetJavaDownloadReportContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaDownloadToken Gets a JavaDownloadToken by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GetJavaDownloadToken.go.html to see an example of how to use GetJavaDownloadToken API.
// A default retry strategy applies to this operation GetJavaDownloadToken()
func (client JavaDownloadClient) GetJavaDownloadToken(ctx context.Context, request GetJavaDownloadTokenRequest) (response GetJavaDownloadTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaDownloadToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaDownloadTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaDownloadTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaDownloadTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaDownloadTokenResponse")
	}
	return
}

// getJavaDownloadToken implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) getJavaDownloadToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaDownloadTokens/{javaDownloadTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaDownloadTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadToken/GetJavaDownloadToken"
		err = common.PostProcessServiceError(err, "JavaDownload", "GetJavaDownloadToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaLicense Return details of the specified Java license type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GetJavaLicense.go.html to see an example of how to use GetJavaLicense API.
// A default retry strategy applies to this operation GetJavaLicense()
func (client JavaDownloadClient) GetJavaLicense(ctx context.Context, request GetJavaLicenseRequest) (response GetJavaLicenseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaLicense, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaLicenseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaLicenseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaLicenseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaLicenseResponse")
	}
	return
}

// getJavaLicense implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) getJavaLicense(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaLicenses/{licenseType}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaLicenseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaLicense/GetJavaLicense"
		err = common.PostProcessServiceError(err, "JavaDownload", "GetJavaLicense", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJavaLicenseAcceptanceRecord Returns a specific Java license acceptance record in a tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GetJavaLicenseAcceptanceRecord.go.html to see an example of how to use GetJavaLicenseAcceptanceRecord API.
// A default retry strategy applies to this operation GetJavaLicenseAcceptanceRecord()
func (client JavaDownloadClient) GetJavaLicenseAcceptanceRecord(ctx context.Context, request GetJavaLicenseAcceptanceRecordRequest) (response GetJavaLicenseAcceptanceRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJavaLicenseAcceptanceRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJavaLicenseAcceptanceRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJavaLicenseAcceptanceRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJavaLicenseAcceptanceRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJavaLicenseAcceptanceRecordResponse")
	}
	return
}

// getJavaLicenseAcceptanceRecord implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) getJavaLicenseAcceptanceRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaLicenseAcceptanceRecords/{javaLicenseAcceptanceRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJavaLicenseAcceptanceRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaLicenseAcceptanceRecord/GetJavaLicenseAcceptanceRecord"
		err = common.PostProcessServiceError(err, "JavaDownload", "GetJavaLicenseAcceptanceRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets details of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client JavaDownloadClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client JavaDownloadClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "JavaDownload", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaDownloadRecords Returns a list of Java download records in a tenancy based on specified parameters.
// See ListJavaReleases
// for possible values of `javaFamilyVersion` and `javaReleaseVersion` parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaDownloadRecords.go.html to see an example of how to use ListJavaDownloadRecords API.
// A default retry strategy applies to this operation ListJavaDownloadRecords()
func (client JavaDownloadClient) ListJavaDownloadRecords(ctx context.Context, request ListJavaDownloadRecordsRequest) (response ListJavaDownloadRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaDownloadRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaDownloadRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaDownloadRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaDownloadRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaDownloadRecordsResponse")
	}
	return
}

// listJavaDownloadRecords implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) listJavaDownloadRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaDownloadRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaDownloadRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadRecord/ListJavaDownloadRecords"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListJavaDownloadRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaDownloadReports Returns a list of JavaDownloadReports.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaDownloadReports.go.html to see an example of how to use ListJavaDownloadReports API.
// A default retry strategy applies to this operation ListJavaDownloadReports()
func (client JavaDownloadClient) ListJavaDownloadReports(ctx context.Context, request ListJavaDownloadReportsRequest) (response ListJavaDownloadReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaDownloadReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaDownloadReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaDownloadReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaDownloadReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaDownloadReportsResponse")
	}
	return
}

// listJavaDownloadReports implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) listJavaDownloadReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaDownloadReports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaDownloadReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadReport/ListJavaDownloadReports"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListJavaDownloadReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaDownloadTokens Returns a list of JavaDownloadTokens.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaDownloadTokens.go.html to see an example of how to use ListJavaDownloadTokens API.
// A default retry strategy applies to this operation ListJavaDownloadTokens()
func (client JavaDownloadClient) ListJavaDownloadTokens(ctx context.Context, request ListJavaDownloadTokensRequest) (response ListJavaDownloadTokensResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaDownloadTokens, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaDownloadTokensResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaDownloadTokensResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaDownloadTokensResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaDownloadTokensResponse")
	}
	return
}

// listJavaDownloadTokens implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) listJavaDownloadTokens(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaDownloadTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaDownloadTokensResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadToken/ListJavaDownloadTokens"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListJavaDownloadTokens", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaLicenseAcceptanceRecords Returns a list of all the Java license acceptance records in a tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaLicenseAcceptanceRecords.go.html to see an example of how to use ListJavaLicenseAcceptanceRecords API.
// A default retry strategy applies to this operation ListJavaLicenseAcceptanceRecords()
func (client JavaDownloadClient) ListJavaLicenseAcceptanceRecords(ctx context.Context, request ListJavaLicenseAcceptanceRecordsRequest) (response ListJavaLicenseAcceptanceRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaLicenseAcceptanceRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaLicenseAcceptanceRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaLicenseAcceptanceRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaLicenseAcceptanceRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaLicenseAcceptanceRecordsResponse")
	}
	return
}

// listJavaLicenseAcceptanceRecords implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) listJavaLicenseAcceptanceRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaLicenseAcceptanceRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaLicenseAcceptanceRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaLicenseAcceptanceRecord/ListJavaLicenseAcceptanceRecords"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListJavaLicenseAcceptanceRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJavaLicenses Return a list with details of all Java licenses.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListJavaLicenses.go.html to see an example of how to use ListJavaLicenses API.
// A default retry strategy applies to this operation ListJavaLicenses()
func (client JavaDownloadClient) ListJavaLicenses(ctx context.Context, request ListJavaLicensesRequest) (response ListJavaLicensesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJavaLicenses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJavaLicensesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJavaLicensesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJavaLicensesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJavaLicensesResponse")
	}
	return
}

// listJavaLicenses implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) listJavaLicenses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/javaLicenses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJavaLicensesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaLicense/ListJavaLicenses"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListJavaLicenses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a (paginated) list of errors for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client JavaDownloadClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client JavaDownloadClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a (paginated) list of logs for the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client JavaDownloadClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client JavaDownloadClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client JavaDownloadClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client JavaDownloadClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20230601/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "JavaDownload", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RequestSummarizedJavaDownloadCounts Returns list of download counts grouped by the specified property.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/RequestSummarizedJavaDownloadCounts.go.html to see an example of how to use RequestSummarizedJavaDownloadCounts API.
// A default retry strategy applies to this operation RequestSummarizedJavaDownloadCounts()
func (client JavaDownloadClient) RequestSummarizedJavaDownloadCounts(ctx context.Context, request RequestSummarizedJavaDownloadCountsRequest) (response RequestSummarizedJavaDownloadCountsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.requestSummarizedJavaDownloadCounts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RequestSummarizedJavaDownloadCountsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RequestSummarizedJavaDownloadCountsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RequestSummarizedJavaDownloadCountsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RequestSummarizedJavaDownloadCountsResponse")
	}
	return
}

// requestSummarizedJavaDownloadCounts implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) requestSummarizedJavaDownloadCounts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20230601/actions/requestSummarizedJavaDownloadCounts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RequestSummarizedJavaDownloadCountsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadCountAggregation/RequestSummarizedJavaDownloadCounts"
		err = common.PostProcessServiceError(err, "JavaDownload", "RequestSummarizedJavaDownloadCounts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJavaDownloadToken Updates the JavaDownloadToken specified by the identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/UpdateJavaDownloadToken.go.html to see an example of how to use UpdateJavaDownloadToken API.
// A default retry strategy applies to this operation UpdateJavaDownloadToken()
func (client JavaDownloadClient) UpdateJavaDownloadToken(ctx context.Context, request UpdateJavaDownloadTokenRequest) (response UpdateJavaDownloadTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJavaDownloadToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJavaDownloadTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJavaDownloadTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJavaDownloadTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJavaDownloadTokenResponse")
	}
	return
}

// updateJavaDownloadToken implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) updateJavaDownloadToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/20230601/javaDownloadTokens/{javaDownloadTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateJavaDownloadTokenResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaDownloadToken/UpdateJavaDownloadToken"
		err = common.PostProcessServiceError(err, "JavaDownload", "UpdateJavaDownloadToken", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJavaLicenseAcceptanceRecord Updates a specific Java license acceptance record in a tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/UpdateJavaLicenseAcceptanceRecord.go.html to see an example of how to use UpdateJavaLicenseAcceptanceRecord API.
// A default retry strategy applies to this operation UpdateJavaLicenseAcceptanceRecord()
func (client JavaDownloadClient) UpdateJavaLicenseAcceptanceRecord(ctx context.Context, request UpdateJavaLicenseAcceptanceRecordRequest) (response UpdateJavaLicenseAcceptanceRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJavaLicenseAcceptanceRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJavaLicenseAcceptanceRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJavaLicenseAcceptanceRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJavaLicenseAcceptanceRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJavaLicenseAcceptanceRecordResponse")
	}
	return
}

// updateJavaLicenseAcceptanceRecord implements the OCIOperation interface (enables retrying operations)
func (client JavaDownloadClient) updateJavaLicenseAcceptanceRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/20230601/javaLicenseAcceptanceRecords/{javaLicenseAcceptanceRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateJavaLicenseAcceptanceRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/jms-java-download/20230601/JavaLicenseAcceptanceRecord/UpdateJavaLicenseAcceptanceRecord"
		err = common.PostProcessServiceError(err, "JavaDownload", "UpdateJavaLicenseAcceptanceRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
