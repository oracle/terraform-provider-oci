// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// LicenseManagerClient a client for LicenseManager
type LicenseManagerClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLicenseManagerClientWithConfigurationProvider Creates a new default LicenseManager client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLicenseManagerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LicenseManagerClient, err error) {
	if enabled := common.CheckForEnabledServices("licensemanager"); !enabled {
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
	return newLicenseManagerClientFromBaseClient(baseClient, provider)
}

// NewLicenseManagerClientWithOboToken Creates a new default LicenseManager client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewLicenseManagerClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LicenseManagerClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newLicenseManagerClientFromBaseClient(baseClient, configProvider)
}

func newLicenseManagerClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LicenseManagerClient, err error) {
	// LicenseManager service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("LicenseManager"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = LicenseManagerClient{BaseClient: baseClient}
	client.BasePath = "20220430"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LicenseManagerClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("licensemanager", "https://licensemanager.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LicenseManagerClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LicenseManagerClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkUploadLicenseRecords Bulk upload the product licenses and license records for a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/BulkUploadLicenseRecords.go.html to see an example of how to use BulkUploadLicenseRecords API.
// A default retry strategy applies to this operation BulkUploadLicenseRecords()
func (client LicenseManagerClient) BulkUploadLicenseRecords(ctx context.Context, request BulkUploadLicenseRecordsRequest) (response BulkUploadLicenseRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkUploadLicenseRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadLicenseRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadLicenseRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadLicenseRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadLicenseRecordsResponse")
	}
	return
}

// bulkUploadLicenseRecords implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) bulkUploadLicenseRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/licenses/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadLicenseRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/BulkUploadLicenseRecordsDetails/BulkUploadLicenseRecords"
		err = common.PostProcessServiceError(err, "LicenseManager", "BulkUploadLicenseRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLicenseRecord Creates a new license record for the given product license ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/CreateLicenseRecord.go.html to see an example of how to use CreateLicenseRecord API.
// A default retry strategy applies to this operation CreateLicenseRecord()
func (client LicenseManagerClient) CreateLicenseRecord(ctx context.Context, request CreateLicenseRecordRequest) (response CreateLicenseRecordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createLicenseRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLicenseRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLicenseRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLicenseRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLicenseRecordResponse")
	}
	return
}

// createLicenseRecord implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) createLicenseRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/licenseRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateLicenseRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/LicenseRecord/CreateLicenseRecord"
		err = common.PostProcessServiceError(err, "LicenseManager", "CreateLicenseRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProductLicense Creates a new product license.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/CreateProductLicense.go.html to see an example of how to use CreateProductLicense API.
// A default retry strategy applies to this operation CreateProductLicense()
func (client LicenseManagerClient) CreateProductLicense(ctx context.Context, request CreateProductLicenseRequest) (response CreateProductLicenseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProductLicense, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProductLicenseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProductLicenseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProductLicenseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProductLicenseResponse")
	}
	return
}

// createProductLicense implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) createProductLicense(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/productLicenses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProductLicenseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/ProductLicense/CreateProductLicense"
		err = common.PostProcessServiceError(err, "LicenseManager", "CreateProductLicense", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLicenseRecord Removes a license record.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/DeleteLicenseRecord.go.html to see an example of how to use DeleteLicenseRecord API.
// A default retry strategy applies to this operation DeleteLicenseRecord()
func (client LicenseManagerClient) DeleteLicenseRecord(ctx context.Context, request DeleteLicenseRecordRequest) (response DeleteLicenseRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLicenseRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLicenseRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLicenseRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLicenseRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLicenseRecordResponse")
	}
	return
}

// deleteLicenseRecord implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) deleteLicenseRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/licenseRecords/{licenseRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteLicenseRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/LicenseRecord/DeleteLicenseRecord"
		err = common.PostProcessServiceError(err, "LicenseManager", "DeleteLicenseRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProductLicense Removes a product license.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/DeleteProductLicense.go.html to see an example of how to use DeleteProductLicense API.
// A default retry strategy applies to this operation DeleteProductLicense()
func (client LicenseManagerClient) DeleteProductLicense(ctx context.Context, request DeleteProductLicenseRequest) (response DeleteProductLicenseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProductLicense, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProductLicenseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProductLicenseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProductLicenseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProductLicenseResponse")
	}
	return
}

// deleteProductLicense implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) deleteProductLicense(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/productLicenses/{productLicenseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProductLicenseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/ProductLicense/DeleteProductLicense"
		err = common.PostProcessServiceError(err, "LicenseManager", "DeleteProductLicense", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBulkUploadTemplate Provides the bulk upload file template.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/GetBulkUploadTemplate.go.html to see an example of how to use GetBulkUploadTemplate API.
// A default retry strategy applies to this operation GetBulkUploadTemplate()
func (client LicenseManagerClient) GetBulkUploadTemplate(ctx context.Context, request GetBulkUploadTemplateRequest) (response GetBulkUploadTemplateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBulkUploadTemplate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBulkUploadTemplateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBulkUploadTemplateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBulkUploadTemplateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBulkUploadTemplateResponse")
	}
	return
}

// getBulkUploadTemplate implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) getBulkUploadTemplate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/licenses/actions/bulkUploadTemplate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBulkUploadTemplateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/BulkUploadTemplate/GetBulkUploadTemplate"
		err = common.PostProcessServiceError(err, "LicenseManager", "GetBulkUploadTemplate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfiguration Retrieves configuration for a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/GetConfiguration.go.html to see an example of how to use GetConfiguration API.
// A default retry strategy applies to this operation GetConfiguration()
func (client LicenseManagerClient) GetConfiguration(ctx context.Context, request GetConfigurationRequest) (response GetConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigurationResponse")
	}
	return
}

// getConfiguration implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) getConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/Configuration/GetConfiguration"
		err = common.PostProcessServiceError(err, "LicenseManager", "GetConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLicenseMetric Retrieves the license metrics for a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/GetLicenseMetric.go.html to see an example of how to use GetLicenseMetric API.
// A default retry strategy applies to this operation GetLicenseMetric()
func (client LicenseManagerClient) GetLicenseMetric(ctx context.Context, request GetLicenseMetricRequest) (response GetLicenseMetricResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLicenseMetric, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLicenseMetricResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLicenseMetricResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLicenseMetricResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLicenseMetricResponse")
	}
	return
}

// getLicenseMetric implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) getLicenseMetric(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/licenseMetrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLicenseMetricResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/LicenseMetric/GetLicenseMetric"
		err = common.PostProcessServiceError(err, "LicenseManager", "GetLicenseMetric", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLicenseRecord Retrieves license record details by the license record ID in a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/GetLicenseRecord.go.html to see an example of how to use GetLicenseRecord API.
// A default retry strategy applies to this operation GetLicenseRecord()
func (client LicenseManagerClient) GetLicenseRecord(ctx context.Context, request GetLicenseRecordRequest) (response GetLicenseRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLicenseRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLicenseRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLicenseRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLicenseRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLicenseRecordResponse")
	}
	return
}

// getLicenseRecord implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) getLicenseRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/licenseRecords/{licenseRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLicenseRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/LicenseRecord/GetLicenseRecord"
		err = common.PostProcessServiceError(err, "LicenseManager", "GetLicenseRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProductLicense Retrieves product license details by product license ID in a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/GetProductLicense.go.html to see an example of how to use GetProductLicense API.
// A default retry strategy applies to this operation GetProductLicense()
func (client LicenseManagerClient) GetProductLicense(ctx context.Context, request GetProductLicenseRequest) (response GetProductLicenseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProductLicense, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProductLicenseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProductLicenseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProductLicenseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProductLicenseResponse")
	}
	return
}

// getProductLicense implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) getProductLicense(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/productLicenses/{productLicenseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProductLicenseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/ProductLicense/GetProductLicense"
		err = common.PostProcessServiceError(err, "LicenseManager", "GetProductLicense", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLicenseRecords Retrieves all license records for a given product license ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListLicenseRecords.go.html to see an example of how to use ListLicenseRecords API.
// A default retry strategy applies to this operation ListLicenseRecords()
func (client LicenseManagerClient) ListLicenseRecords(ctx context.Context, request ListLicenseRecordsRequest) (response ListLicenseRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLicenseRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLicenseRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLicenseRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLicenseRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLicenseRecordsResponse")
	}
	return
}

// listLicenseRecords implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) listLicenseRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/licenseRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLicenseRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/LicenseRecordCollection/ListLicenseRecords"
		err = common.PostProcessServiceError(err, "LicenseManager", "ListLicenseRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProductLicenseConsumers Retrieves the product license consumers for a particular product license ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListProductLicenseConsumers.go.html to see an example of how to use ListProductLicenseConsumers API.
// A default retry strategy applies to this operation ListProductLicenseConsumers()
func (client LicenseManagerClient) ListProductLicenseConsumers(ctx context.Context, request ListProductLicenseConsumersRequest) (response ListProductLicenseConsumersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProductLicenseConsumers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProductLicenseConsumersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProductLicenseConsumersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProductLicenseConsumersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProductLicenseConsumersResponse")
	}
	return
}

// listProductLicenseConsumers implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) listProductLicenseConsumers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/productLicenseConsumers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProductLicenseConsumersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/ProductLicenseConsumerCollection/ListProductLicenseConsumers"
		err = common.PostProcessServiceError(err, "LicenseManager", "ListProductLicenseConsumers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProductLicenses Retrieves all the product licenses from a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListProductLicenses.go.html to see an example of how to use ListProductLicenses API.
// A default retry strategy applies to this operation ListProductLicenses()
func (client LicenseManagerClient) ListProductLicenses(ctx context.Context, request ListProductLicensesRequest) (response ListProductLicensesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProductLicenses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProductLicensesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProductLicensesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProductLicensesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProductLicensesResponse")
	}
	return
}

// listProductLicenses implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) listProductLicenses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/productLicenses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProductLicensesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/ProductLicenseCollection/ListProductLicenses"
		err = common.PostProcessServiceError(err, "LicenseManager", "ListProductLicenses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTopUtilizedProductLicenses Retrieves the top utilized product licenses for a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListTopUtilizedProductLicenses.go.html to see an example of how to use ListTopUtilizedProductLicenses API.
// A default retry strategy applies to this operation ListTopUtilizedProductLicenses()
func (client LicenseManagerClient) ListTopUtilizedProductLicenses(ctx context.Context, request ListTopUtilizedProductLicensesRequest) (response ListTopUtilizedProductLicensesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTopUtilizedProductLicenses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTopUtilizedProductLicensesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTopUtilizedProductLicensesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTopUtilizedProductLicensesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTopUtilizedProductLicensesResponse")
	}
	return
}

// listTopUtilizedProductLicenses implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) listTopUtilizedProductLicenses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/topUtilizedProductLicenses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTopUtilizedProductLicensesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/TopUtilizedProductLicenseCollection/ListTopUtilizedProductLicenses"
		err = common.PostProcessServiceError(err, "LicenseManager", "ListTopUtilizedProductLicenses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTopUtilizedResources Retrieves the top utilized resources for a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/ListTopUtilizedResources.go.html to see an example of how to use ListTopUtilizedResources API.
// A default retry strategy applies to this operation ListTopUtilizedResources()
func (client LicenseManagerClient) ListTopUtilizedResources(ctx context.Context, request ListTopUtilizedResourcesRequest) (response ListTopUtilizedResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTopUtilizedResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTopUtilizedResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTopUtilizedResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTopUtilizedResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTopUtilizedResourcesResponse")
	}
	return
}

// listTopUtilizedResources implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) listTopUtilizedResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/topUtilizedResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTopUtilizedResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/TopUtilizedResourceCollection/ListTopUtilizedResources"
		err = common.PostProcessServiceError(err, "LicenseManager", "ListTopUtilizedResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConfiguration Updates the configuration for the compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/UpdateConfiguration.go.html to see an example of how to use UpdateConfiguration API.
// A default retry strategy applies to this operation UpdateConfiguration()
func (client LicenseManagerClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigurationResponse")
	}
	return
}

// updateConfiguration implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) updateConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/Configuration/UpdateConfiguration"
		err = common.PostProcessServiceError(err, "LicenseManager", "UpdateConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLicenseRecord Updates license record entity details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/UpdateLicenseRecord.go.html to see an example of how to use UpdateLicenseRecord API.
// A default retry strategy applies to this operation UpdateLicenseRecord()
func (client LicenseManagerClient) UpdateLicenseRecord(ctx context.Context, request UpdateLicenseRecordRequest) (response UpdateLicenseRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLicenseRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLicenseRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLicenseRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLicenseRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLicenseRecordResponse")
	}
	return
}

// updateLicenseRecord implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) updateLicenseRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/licenseRecords/{licenseRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateLicenseRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/LicenseRecord/UpdateLicenseRecord"
		err = common.PostProcessServiceError(err, "LicenseManager", "UpdateLicenseRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProductLicense Updates the list of images for a product license.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/licensemanager/UpdateProductLicense.go.html to see an example of how to use UpdateProductLicense API.
// A default retry strategy applies to this operation UpdateProductLicense()
func (client LicenseManagerClient) UpdateProductLicense(ctx context.Context, request UpdateProductLicenseRequest) (response UpdateProductLicenseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProductLicense, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProductLicenseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProductLicenseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProductLicenseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProductLicenseResponse")
	}
	return
}

// updateProductLicense implements the OCIOperation interface (enables retrying operations)
func (client LicenseManagerClient) updateProductLicense(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/productLicenses/{productLicenseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProductLicenseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/licensemanager/20220430/ProductLicense/UpdateProductLicense"
		err = common.PostProcessServiceError(err, "LicenseManager", "UpdateProductLicense", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
