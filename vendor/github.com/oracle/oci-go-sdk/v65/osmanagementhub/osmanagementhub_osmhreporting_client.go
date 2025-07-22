// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// OsmhReportingClient a client for OsmhReporting
type OsmhReportingClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOsmhReportingClientWithConfigurationProvider Creates a new default OsmhReporting client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOsmhReportingClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OsmhReportingClient, err error) {
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
	return newOsmhReportingClientFromBaseClient(baseClient, provider)
}

// NewOsmhReportingClientWithOboToken Creates a new default OsmhReporting client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOsmhReportingClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OsmhReportingClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOsmhReportingClientFromBaseClient(baseClient, configProvider)
}

func newOsmhReportingClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OsmhReportingClient, err error) {
	// OsmhReporting service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OsmhReporting"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OsmhReportingClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OsmhReportingClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OsmhReportingClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OsmhReportingClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeReportCompartment Moves a Report into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation ChangeReportCompartment()
func (client OsmhReportingClient) ChangeReportCompartment(ctx context.Context, request ChangeReportCompartmentRequest) (response ChangeReportCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeReportCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeReportCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeReportCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeReportCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeReportCompartmentResponse")
	}
	return
}

// changeReportCompartment implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) changeReportCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reports/{reportId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeReportCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/ChangeReportCompartment"
		err = common.PostProcessServiceError(err, "OsmhReporting", "ChangeReportCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateReport Creates a Report.
// A default retry strategy applies to this operation CreateReport()
func (client OsmhReportingClient) CreateReport(ctx context.Context, request CreateReportRequest) (response CreateReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateReportResponse")
	}
	return
}

// createReport implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) createReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/CreateReport"
		err = common.PostProcessServiceError(err, "OsmhReporting", "CreateReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &report{})
	return response, err
}

// DeleteReport Deletes a Report.
// A default retry strategy applies to this operation DeleteReport()
func (client OsmhReportingClient) DeleteReport(ctx context.Context, request DeleteReportRequest) (response DeleteReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteReportResponse")
	}
	return
}

// deleteReport implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) deleteReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/reports/{reportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/DeleteReport"
		err = common.PostProcessServiceError(err, "OsmhReporting", "DeleteReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReport Gets information about a Report.
// A default retry strategy applies to this operation GetReport()
func (client OsmhReportingClient) GetReport(ctx context.Context, request GetReportRequest) (response GetReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReportResponse")
	}
	return
}

// getReport implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) getReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports/{reportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/GetReport"
		err = common.PostProcessServiceError(err, "OsmhReporting", "GetReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &report{})
	return response, err
}

// GetReportContent Returns report content for the given reportId. You can select CSV, XML, or JSON format.
// A default retry strategy applies to this operation GetReportContent()
func (client OsmhReportingClient) GetReportContent(ctx context.Context, request GetReportContentRequest) (response GetReportContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReportContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReportContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReportContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReportContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReportContentResponse")
	}
	return
}

// getReportContent implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) getReportContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports/{reportId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReportContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/GetReportContent"
		err = common.PostProcessServiceError(err, "OsmhReporting", "GetReportContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetReportVersion Gets information about a Report with specific version.
// A default retry strategy applies to this operation GetReportVersion()
func (client OsmhReportingClient) GetReportVersion(ctx context.Context, request GetReportVersionRequest) (response GetReportVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReportVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReportVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReportVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReportVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReportVersionResponse")
	}
	return
}

// getReportVersion implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) getReportVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports/{reportId}/version/{reportVersion}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReportVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ReportVersion/GetReportVersion"
		err = common.PostProcessServiceError(err, "OsmhReporting", "GetReportVersion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &reportversion{})
	return response, err
}

// GetReportVersionContent Returns report content for the given reportId and version. You can select CSV, XML, or JSON format.
// A default retry strategy applies to this operation GetReportVersionContent()
func (client OsmhReportingClient) GetReportVersionContent(ctx context.Context, request GetReportVersionContentRequest) (response GetReportVersionContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getReportVersionContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetReportVersionContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetReportVersionContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetReportVersionContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetReportVersionContentResponse")
	}
	return
}

// getReportVersionContent implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) getReportVersionContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports/{reportId}/versions/{reportVersion}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetReportVersionContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/GetReportVersionContent"
		err = common.PostProcessServiceError(err, "OsmhReporting", "GetReportVersionContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReportVersions Returns information about the versions of the specified Report.
// A default retry strategy applies to this operation ListReportVersions()
func (client OsmhReportingClient) ListReportVersions(ctx context.Context, request ListReportVersionsRequest) (response ListReportVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReportVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReportVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReportVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReportVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReportVersionsResponse")
	}
	return
}

// listReportVersions implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) listReportVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports/{reportId}/versions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReportVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/VersionCollection/ListReportVersions"
		err = common.PostProcessServiceError(err, "OsmhReporting", "ListReportVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReports Gets a list of Reports.
// A default retry strategy applies to this operation ListReports()
func (client OsmhReportingClient) ListReports(ctx context.Context, request ListReportsRequest) (response ListReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReportsResponse")
	}
	return
}

// listReports implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) listReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/ListReports"
		err = common.PostProcessServiceError(err, "OsmhReporting", "ListReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RerunReport Re-runs a Report.
// A default retry strategy applies to this operation RerunReport()
func (client OsmhReportingClient) RerunReport(ctx context.Context, request RerunReportRequest) (response RerunReportResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rerunReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RerunReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RerunReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RerunReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RerunReportResponse")
	}
	return
}

// rerunReport implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) rerunReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reports/{reportId}/actions/rerunReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RerunReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/RerunReport"
		err = common.PostProcessServiceError(err, "OsmhReporting", "RerunReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &report{})
	return response, err
}

// UpdateReport Updates a Report.
// A default retry strategy applies to this operation UpdateReport()
func (client OsmhReportingClient) UpdateReport(ctx context.Context, request UpdateReportRequest) (response UpdateReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateReportResponse")
	}
	return
}

// updateReport implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) updateReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/reports/{reportId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/UpdateReport"
		err = common.PostProcessServiceError(err, "OsmhReporting", "UpdateReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &report{})
	return response, err
}

// UpdateReportVulnerabilities Updates vulnerabilities for a report
// A default retry strategy applies to this operation UpdateReportVulnerabilities()
func (client OsmhReportingClient) UpdateReportVulnerabilities(ctx context.Context, request UpdateReportVulnerabilitiesRequest) (response UpdateReportVulnerabilitiesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateReportVulnerabilities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateReportVulnerabilitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateReportVulnerabilitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateReportVulnerabilitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateReportVulnerabilitiesResponse")
	}
	return
}

// updateReportVulnerabilities implements the OCIOperation interface (enables retrying operations)
func (client OsmhReportingClient) updateReportVulnerabilities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/reports/{reportId}/actions/updateVulnerabilities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateReportVulnerabilitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Report/UpdateReportVulnerabilities"
		err = common.PostProcessServiceError(err, "OsmhReporting", "UpdateReportVulnerabilities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
