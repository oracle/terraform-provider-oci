// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DataCatalogClient a client for DataCatalog
type DataCatalogClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDataCatalogClientWithConfigurationProvider Creates a new default DataCatalog client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDataCatalogClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DataCatalogClient, err error) {
	if enabled := common.CheckForEnabledServices("datacatalog"); !enabled {
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
	return newDataCatalogClientFromBaseClient(baseClient, provider)
}

// NewDataCatalogClientWithOboToken Creates a new default DataCatalog client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDataCatalogClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DataCatalogClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDataCatalogClientFromBaseClient(baseClient, configProvider)
}

func newDataCatalogClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DataCatalogClient, err error) {
	// DataCatalog service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DataCatalog"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DataCatalogClient{BaseClient: baseClient}
	client.BasePath = "20190325"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DataCatalogClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datacatalog", "https://datacatalog.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DataCatalogClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DataCatalogClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddCatalogLock Adds a lock to a Catalog resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/AddCatalogLock.go.html to see an example of how to use AddCatalogLock API.
func (client DataCatalogClient) AddCatalogLock(ctx context.Context, request AddCatalogLockRequest) (response AddCatalogLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addCatalogLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddCatalogLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddCatalogLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddCatalogLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddCatalogLockResponse")
	}
	return
}

// addCatalogLock implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) addCatalogLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddCatalogLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/AddCatalogLock"
		err = common.PostProcessServiceError(err, "DataCatalog", "AddCatalogLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddCatalogPrivateEndpointLock Adds a lock to a CatalogPrivateEndpoint resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/AddCatalogPrivateEndpointLock.go.html to see an example of how to use AddCatalogPrivateEndpointLock API.
func (client DataCatalogClient) AddCatalogPrivateEndpointLock(ctx context.Context, request AddCatalogPrivateEndpointLockRequest) (response AddCatalogPrivateEndpointLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addCatalogPrivateEndpointLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddCatalogPrivateEndpointLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddCatalogPrivateEndpointLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddCatalogPrivateEndpointLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddCatalogPrivateEndpointLockResponse")
	}
	return
}

// addCatalogPrivateEndpointLock implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) addCatalogPrivateEndpointLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddCatalogPrivateEndpointLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpoint/AddCatalogPrivateEndpointLock"
		err = common.PostProcessServiceError(err, "DataCatalog", "AddCatalogPrivateEndpointLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddDataSelectorPatterns Add data selector pattern to the data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/AddDataSelectorPatterns.go.html to see an example of how to use AddDataSelectorPatterns API.
func (client DataCatalogClient) AddDataSelectorPatterns(ctx context.Context, request AddDataSelectorPatternsRequest) (response AddDataSelectorPatternsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addDataSelectorPatterns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDataSelectorPatternsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDataSelectorPatternsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDataSelectorPatternsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDataSelectorPatternsResponse")
	}
	return
}

// addDataSelectorPatterns implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) addDataSelectorPatterns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/addDataSelectorPatterns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDataSelectorPatternsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/AddDataSelectorPatterns"
		err = common.PostProcessServiceError(err, "DataCatalog", "AddDataSelectorPatterns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddMetastoreLock Adds a lock to a Metastore resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/AddMetastoreLock.go.html to see an example of how to use AddMetastoreLock API.
func (client DataCatalogClient) AddMetastoreLock(ctx context.Context, request AddMetastoreLockRequest) (response AddMetastoreLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addMetastoreLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddMetastoreLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddMetastoreLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddMetastoreLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddMetastoreLockResponse")
	}
	return
}

// addMetastoreLock implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) addMetastoreLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metastores/{metastoreId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddMetastoreLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Metastore/AddMetastoreLock"
		err = common.PostProcessServiceError(err, "DataCatalog", "AddMetastoreLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AssociateCustomProperty Associate the custom property for the given type
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/AssociateCustomProperty.go.html to see an example of how to use AssociateCustomProperty API.
func (client DataCatalogClient) AssociateCustomProperty(ctx context.Context, request AssociateCustomPropertyRequest) (response AssociateCustomPropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.associateCustomProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AssociateCustomPropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AssociateCustomPropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AssociateCustomPropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AssociateCustomPropertyResponse")
	}
	return
}

// associateCustomProperty implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) associateCustomProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/types/{typeKey}/actions/associateCustomProperties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AssociateCustomPropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Type/AssociateCustomProperty"
		err = common.PostProcessServiceError(err, "DataCatalog", "AssociateCustomProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AsynchronousExportGlossary Exports the contents of a glossary in Excel format. Returns details about the job which actually performs the export.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/AsynchronousExportGlossary.go.html to see an example of how to use AsynchronousExportGlossary API.
func (client DataCatalogClient) AsynchronousExportGlossary(ctx context.Context, request AsynchronousExportGlossaryRequest) (response AsynchronousExportGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.asynchronousExportGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AsynchronousExportGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AsynchronousExportGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AsynchronousExportGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AsynchronousExportGlossaryResponse")
	}
	return
}

// asynchronousExportGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) asynchronousExportGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/actions/asynchronousExport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AsynchronousExportGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/AsynchronousExportGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "AsynchronousExportGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AttachCatalogPrivateEndpoint Attaches a private reverse connection endpoint resource to a data catalog resource. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/AttachCatalogPrivateEndpoint.go.html to see an example of how to use AttachCatalogPrivateEndpoint API.
func (client DataCatalogClient) AttachCatalogPrivateEndpoint(ctx context.Context, request AttachCatalogPrivateEndpointRequest) (response AttachCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.attachCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachCatalogPrivateEndpointResponse")
	}
	return
}

// attachCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) attachCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/attachCatalogPrivateEndpoint", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachCatalogPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/AttachCatalogPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataCatalog", "AttachCatalogPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCatalogCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ChangeCatalogCompartment.go.html to see an example of how to use ChangeCatalogCompartment API.
func (client DataCatalogClient) ChangeCatalogCompartment(ctx context.Context, request ChangeCatalogCompartmentRequest) (response ChangeCatalogCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeCatalogCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCatalogCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCatalogCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCatalogCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCatalogCompartmentResponse")
	}
	return
}

// changeCatalogCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) changeCatalogCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCatalogCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/ChangeCatalogCompartment"
		err = common.PostProcessServiceError(err, "DataCatalog", "ChangeCatalogCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCatalogPrivateEndpointCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ChangeCatalogPrivateEndpointCompartment.go.html to see an example of how to use ChangeCatalogPrivateEndpointCompartment API.
func (client DataCatalogClient) ChangeCatalogPrivateEndpointCompartment(ctx context.Context, request ChangeCatalogPrivateEndpointCompartmentRequest) (response ChangeCatalogPrivateEndpointCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeCatalogPrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCatalogPrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCatalogPrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCatalogPrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCatalogPrivateEndpointCompartmentResponse")
	}
	return
}

// changeCatalogPrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) changeCatalogPrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCatalogPrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpoint/ChangeCatalogPrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "DataCatalog", "ChangeCatalogPrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMetastoreCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ChangeMetastoreCompartment.go.html to see an example of how to use ChangeMetastoreCompartment API.
func (client DataCatalogClient) ChangeMetastoreCompartment(ctx context.Context, request ChangeMetastoreCompartmentRequest) (response ChangeMetastoreCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeMetastoreCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMetastoreCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMetastoreCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMetastoreCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMetastoreCompartmentResponse")
	}
	return
}

// changeMetastoreCompartment implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) changeMetastoreCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metastores/{metastoreId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMetastoreCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Metastore/ChangeMetastoreCompartment"
		err = common.PostProcessServiceError(err, "DataCatalog", "ChangeMetastoreCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAttribute Creates a new entity attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateAttribute.go.html to see an example of how to use CreateAttribute API.
func (client DataCatalogClient) CreateAttribute(ctx context.Context, request CreateAttributeRequest) (response CreateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAttributeResponse")
	}
	return
}

// createAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Attribute/CreateAttribute"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAttributeTag Creates a new entity attribute tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateAttributeTag.go.html to see an example of how to use CreateAttributeTag API.
func (client DataCatalogClient) CreateAttributeTag(ctx context.Context, request CreateAttributeTagRequest) (response CreateAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAttributeTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAttributeTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAttributeTagResponse")
	}
	return
}

// createAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createAttributeTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAttributeTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/AttributeTag/CreateAttributeTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateAttributeTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCatalog Creates a new data catalog instance that includes a console and an API URL for managing metadata operations.
// For more information, please see the documentation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateCatalog.go.html to see an example of how to use CreateCatalog API.
func (client DataCatalogClient) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (response CreateCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCatalogResponse")
	}
	return
}

// createCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/CreateCatalog"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCatalogPrivateEndpoint Create a new private reverse connection endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateCatalogPrivateEndpoint.go.html to see an example of how to use CreateCatalogPrivateEndpoint API.
func (client DataCatalogClient) CreateCatalogPrivateEndpoint(ctx context.Context, request CreateCatalogPrivateEndpointRequest) (response CreateCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCatalogPrivateEndpointResponse")
	}
	return
}

// createCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCatalogPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpoint/CreateCatalogPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateCatalogPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConnection Creates a new connection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateConnection.go.html to see an example of how to use CreateConnection API.
func (client DataCatalogClient) CreateConnection(ctx context.Context, request CreateConnectionRequest) (response CreateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionResponse")
	}
	return
}

// createConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Connection/CreateConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCustomProperty Create a new Custom Property
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateCustomProperty.go.html to see an example of how to use CreateCustomProperty API.
func (client DataCatalogClient) CreateCustomProperty(ctx context.Context, request CreateCustomPropertyRequest) (response CreateCustomPropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createCustomProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCustomPropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCustomPropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCustomPropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCustomPropertyResponse")
	}
	return
}

// createCustomProperty implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createCustomProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/namespaces/{namespaceId}/customProperties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCustomPropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CustomProperty/CreateCustomProperty"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateCustomProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataAsset Create a new data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateDataAsset.go.html to see an example of how to use CreateDataAsset API.
func (client DataCatalogClient) CreateDataAsset(ctx context.Context, request CreateDataAssetRequest) (response CreateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataAssetResponse")
	}
	return
}

// createDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/CreateDataAsset"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateDataAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDataAssetTag Creates a new data asset tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateDataAssetTag.go.html to see an example of how to use CreateDataAssetTag API.
func (client DataCatalogClient) CreateDataAssetTag(ctx context.Context, request CreateDataAssetTagRequest) (response CreateDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDataAssetTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDataAssetTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDataAssetTagResponse")
	}
	return
}

// createDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createDataAssetTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDataAssetTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAssetTag/CreateDataAssetTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateDataAssetTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEntity Creates a new data entity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateEntity.go.html to see an example of how to use CreateEntity API.
func (client DataCatalogClient) CreateEntity(ctx context.Context, request CreateEntityRequest) (response CreateEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntityResponse")
	}
	return
}

// createEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Entity/CreateEntity"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateEntity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEntityTag Creates a new entity tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateEntityTag.go.html to see an example of how to use CreateEntityTag API.
func (client DataCatalogClient) CreateEntityTag(ctx context.Context, request CreateEntityTagRequest) (response CreateEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEntityTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEntityTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEntityTagResponse")
	}
	return
}

// createEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createEntityTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEntityTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/EntityTag/CreateEntityTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateEntityTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFolder Creates a new folder.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateFolder.go.html to see an example of how to use CreateFolder API.
func (client DataCatalogClient) CreateFolder(ctx context.Context, request CreateFolderRequest) (response CreateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFolderResponse")
	}
	return
}

// createFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Folder/CreateFolder"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateFolder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateFolderTag Creates a new folder tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateFolderTag.go.html to see an example of how to use CreateFolderTag API.
func (client DataCatalogClient) CreateFolderTag(ctx context.Context, request CreateFolderTagRequest) (response CreateFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFolderTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFolderTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFolderTagResponse")
	}
	return
}

// createFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createFolderTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFolderTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/FolderTag/CreateFolderTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateFolderTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateGlossary Creates a new glossary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateGlossary.go.html to see an example of how to use CreateGlossary API.
func (client DataCatalogClient) CreateGlossary(ctx context.Context, request CreateGlossaryRequest) (response CreateGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGlossaryResponse")
	}
	return
}

// createGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/CreateGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJob Creates a new job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateJob.go.html to see an example of how to use CreateJob API.
func (client DataCatalogClient) CreateJob(ctx context.Context, request CreateJobRequest) (response CreateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobResponse")
	}
	return
}

// createJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Job/CreateJob"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJobDefinition Creates a new job definition.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateJobDefinition.go.html to see an example of how to use CreateJobDefinition API.
func (client DataCatalogClient) CreateJobDefinition(ctx context.Context, request CreateJobDefinitionRequest) (response CreateJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobDefinitionResponse")
	}
	return
}

// createJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJobDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJobDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobDefinition/CreateJobDefinition"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateJobDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateJobExecution Creates a new job execution.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateJobExecution.go.html to see an example of how to use CreateJobExecution API.
func (client DataCatalogClient) CreateJobExecution(ctx context.Context, request CreateJobExecutionRequest) (response CreateJobExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateJobExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateJobExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateJobExecutionResponse")
	}
	return
}

// createJobExecution implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createJobExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/jobs/{jobKey}/executions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateJobExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobExecution/CreateJobExecution"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateJobExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMetastore Creates a new metastore.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateMetastore.go.html to see an example of how to use CreateMetastore API.
func (client DataCatalogClient) CreateMetastore(ctx context.Context, request CreateMetastoreRequest) (response CreateMetastoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createMetastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMetastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMetastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMetastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMetastoreResponse")
	}
	return
}

// createMetastore implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createMetastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metastores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMetastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Metastore/CreateMetastore"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateMetastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateNamespace Create a new Namespace to be used by a custom property
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateNamespace.go.html to see an example of how to use CreateNamespace API.
func (client DataCatalogClient) CreateNamespace(ctx context.Context, request CreateNamespaceRequest) (response CreateNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNamespaceResponse")
	}
	return
}

// createNamespace implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/namespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Namespace/CreateNamespace"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePattern Create a new pattern.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreatePattern.go.html to see an example of how to use CreatePattern API.
func (client DataCatalogClient) CreatePattern(ctx context.Context, request CreatePatternRequest) (response CreatePatternResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createPattern, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePatternResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePatternResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePatternResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePatternResponse")
	}
	return
}

// createPattern implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createPattern(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/patterns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePatternResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Pattern/CreatePattern"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreatePattern", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTerm Create a new term within a glossary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateTerm.go.html to see an example of how to use CreateTerm API.
func (client DataCatalogClient) CreateTerm(ctx context.Context, request CreateTermRequest) (response CreateTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTermResponse")
	}
	return
}

// createTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Term/CreateTerm"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTermRelationship Creates a new term relationship for this term within a glossary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/CreateTermRelationship.go.html to see an example of how to use CreateTermRelationship API.
func (client DataCatalogClient) CreateTermRelationship(ctx context.Context, request CreateTermRelationshipRequest) (response CreateTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTermRelationshipResponse")
	}
	return
}

// createTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) createTermRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTermRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/TermRelationship/CreateTermRelationship"
		err = common.PostProcessServiceError(err, "DataCatalog", "CreateTermRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAttribute Deletes a specific entity attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteAttribute.go.html to see an example of how to use DeleteAttribute API.
func (client DataCatalogClient) DeleteAttribute(ctx context.Context, request DeleteAttributeRequest) (response DeleteAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAttributeResponse")
	}
	return
}

// deleteAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Attribute/DeleteAttribute"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAttributeTag Deletes a specific entity attribute tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteAttributeTag.go.html to see an example of how to use DeleteAttributeTag API.
func (client DataCatalogClient) DeleteAttributeTag(ctx context.Context, request DeleteAttributeTagRequest) (response DeleteAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAttributeTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAttributeTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAttributeTagResponse")
	}
	return
}

// deleteAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteAttributeTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAttributeTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/AttributeTag/DeleteAttributeTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteAttributeTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCatalog Deletes a data catalog resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteCatalog.go.html to see an example of how to use DeleteCatalog API.
func (client DataCatalogClient) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (response DeleteCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCatalogResponse")
	}
	return
}

// deleteCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/DeleteCatalog"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCatalogPrivateEndpoint Deletes a private reverse connection endpoint by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteCatalogPrivateEndpoint.go.html to see an example of how to use DeleteCatalogPrivateEndpoint API.
func (client DataCatalogClient) DeleteCatalogPrivateEndpoint(ctx context.Context, request DeleteCatalogPrivateEndpointRequest) (response DeleteCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCatalogPrivateEndpointResponse")
	}
	return
}

// deleteCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCatalogPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpoint/DeleteCatalogPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteCatalogPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnection Deletes a specific connection of a data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteConnection.go.html to see an example of how to use DeleteConnection API.
func (client DataCatalogClient) DeleteConnection(ctx context.Context, request DeleteConnectionRequest) (response DeleteConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionResponse")
	}
	return
}

// deleteConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Connection/DeleteConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCustomProperty Deletes a specific custom property identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteCustomProperty.go.html to see an example of how to use DeleteCustomProperty API.
func (client DataCatalogClient) DeleteCustomProperty(ctx context.Context, request DeleteCustomPropertyRequest) (response DeleteCustomPropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCustomProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCustomPropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCustomPropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCustomPropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCustomPropertyResponse")
	}
	return
}

// deleteCustomProperty implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteCustomProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/namespaces/{namespaceId}/customProperties/{customPropertyKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCustomPropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CustomProperty/DeleteCustomProperty"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteCustomProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataAsset Deletes a specific data asset identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteDataAsset.go.html to see an example of how to use DeleteDataAsset API.
func (client DataCatalogClient) DeleteDataAsset(ctx context.Context, request DeleteDataAssetRequest) (response DeleteDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataAssetResponse")
	}
	return
}

// deleteDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/DeleteDataAsset"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteDataAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDataAssetTag Deletes a specific data asset tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteDataAssetTag.go.html to see an example of how to use DeleteDataAssetTag API.
func (client DataCatalogClient) DeleteDataAssetTag(ctx context.Context, request DeleteDataAssetTagRequest) (response DeleteDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDataAssetTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDataAssetTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDataAssetTagResponse")
	}
	return
}

// deleteDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteDataAssetTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDataAssetTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAssetTag/DeleteDataAssetTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteDataAssetTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEntity Deletes a specific data entity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteEntity.go.html to see an example of how to use DeleteEntity API.
func (client DataCatalogClient) DeleteEntity(ctx context.Context, request DeleteEntityRequest) (response DeleteEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEntityResponse")
	}
	return
}

// deleteEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Entity/DeleteEntity"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteEntity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEntityTag Deletes a specific entity tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteEntityTag.go.html to see an example of how to use DeleteEntityTag API.
func (client DataCatalogClient) DeleteEntityTag(ctx context.Context, request DeleteEntityTagRequest) (response DeleteEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEntityTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEntityTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEntityTagResponse")
	}
	return
}

// deleteEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteEntityTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEntityTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/EntityTag/DeleteEntityTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteEntityTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFolder Deletes a specific folder of a data asset identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteFolder.go.html to see an example of how to use DeleteFolder API.
func (client DataCatalogClient) DeleteFolder(ctx context.Context, request DeleteFolderRequest) (response DeleteFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFolderResponse")
	}
	return
}

// deleteFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Folder/DeleteFolder"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteFolder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFolderTag Deletes a specific folder tag.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteFolderTag.go.html to see an example of how to use DeleteFolderTag API.
func (client DataCatalogClient) DeleteFolderTag(ctx context.Context, request DeleteFolderTagRequest) (response DeleteFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFolderTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFolderTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFolderTagResponse")
	}
	return
}

// deleteFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteFolderTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFolderTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/FolderTag/DeleteFolderTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteFolderTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteGlossary Deletes a specific glossary identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteGlossary.go.html to see an example of how to use DeleteGlossary API.
func (client DataCatalogClient) DeleteGlossary(ctx context.Context, request DeleteGlossaryRequest) (response DeleteGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGlossaryResponse")
	}
	return
}

// deleteGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/DeleteGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJob Deletes a specific job identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteJob.go.html to see an example of how to use DeleteJob API.
func (client DataCatalogClient) DeleteJob(ctx context.Context, request DeleteJobRequest) (response DeleteJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobResponse")
	}
	return
}

// deleteJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/jobs/{jobKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Job/DeleteJob"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteJobDefinition Deletes a specific job definition identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteJobDefinition.go.html to see an example of how to use DeleteJobDefinition API.
func (client DataCatalogClient) DeleteJobDefinition(ctx context.Context, request DeleteJobDefinitionRequest) (response DeleteJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteJobDefinitionResponse")
	}
	return
}

// deleteJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteJobDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteJobDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobDefinition/DeleteJobDefinition"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteJobDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMetastore Deletes a metastore resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteMetastore.go.html to see an example of how to use DeleteMetastore API.
func (client DataCatalogClient) DeleteMetastore(ctx context.Context, request DeleteMetastoreRequest) (response DeleteMetastoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMetastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMetastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMetastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMetastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMetastoreResponse")
	}
	return
}

// deleteMetastore implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteMetastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/metastores/{metastoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMetastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Metastore/DeleteMetastore"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteMetastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNamespace Deletes a specific Namespace identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteNamespace.go.html to see an example of how to use DeleteNamespace API.
func (client DataCatalogClient) DeleteNamespace(ctx context.Context, request DeleteNamespaceRequest) (response DeleteNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNamespaceResponse")
	}
	return
}

// deleteNamespace implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/namespaces/{namespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Namespace/DeleteNamespace"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePattern Deletes a specific pattern identified by it's key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeletePattern.go.html to see an example of how to use DeletePattern API.
func (client DataCatalogClient) DeletePattern(ctx context.Context, request DeletePatternRequest) (response DeletePatternResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePattern, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePatternResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePatternResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePatternResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePatternResponse")
	}
	return
}

// deletePattern implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deletePattern(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/patterns/{patternKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePatternResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Pattern/DeletePattern"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeletePattern", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTerm Deletes a specific glossary term.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteTerm.go.html to see an example of how to use DeleteTerm API.
func (client DataCatalogClient) DeleteTerm(ctx context.Context, request DeleteTermRequest) (response DeleteTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTermResponse")
	}
	return
}

// deleteTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Term/DeleteTerm"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTermRelationship Deletes a specific glossary term relationship.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DeleteTermRelationship.go.html to see an example of how to use DeleteTermRelationship API.
func (client DataCatalogClient) DeleteTermRelationship(ctx context.Context, request DeleteTermRelationshipRequest) (response DeleteTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTermRelationshipResponse")
	}
	return
}

// deleteTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) deleteTermRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships/{termRelationshipKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTermRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/TermRelationship/DeleteTermRelationship"
		err = common.PostProcessServiceError(err, "DataCatalog", "DeleteTermRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachCatalogPrivateEndpoint Detaches a private reverse connection endpoint resource to a data catalog resource. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DetachCatalogPrivateEndpoint.go.html to see an example of how to use DetachCatalogPrivateEndpoint API.
func (client DataCatalogClient) DetachCatalogPrivateEndpoint(ctx context.Context, request DetachCatalogPrivateEndpointRequest) (response DetachCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detachCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachCatalogPrivateEndpointResponse")
	}
	return
}

// detachCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) detachCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/detachCatalogPrivateEndpoint", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachCatalogPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/DetachCatalogPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataCatalog", "DetachCatalogPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisassociateCustomProperty Remove the custom property for the given type
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/DisassociateCustomProperty.go.html to see an example of how to use DisassociateCustomProperty API.
func (client DataCatalogClient) DisassociateCustomProperty(ctx context.Context, request DisassociateCustomPropertyRequest) (response DisassociateCustomPropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.disassociateCustomProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisassociateCustomPropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisassociateCustomPropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisassociateCustomPropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisassociateCustomPropertyResponse")
	}
	return
}

// disassociateCustomProperty implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) disassociateCustomProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/types/{typeKey}/actions/disassociateCustomProperties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DisassociateCustomPropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Type/DisassociateCustomProperty"
		err = common.PostProcessServiceError(err, "DataCatalog", "DisassociateCustomProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExpandTreeForGlossary Returns the fully expanded tree hierarchy of parent and child terms in this glossary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ExpandTreeForGlossary.go.html to see an example of how to use ExpandTreeForGlossary API.
// A default retry strategy applies to this operation ExpandTreeForGlossary()
func (client DataCatalogClient) ExpandTreeForGlossary(ctx context.Context, request ExpandTreeForGlossaryRequest) (response ExpandTreeForGlossaryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.expandTreeForGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExpandTreeForGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExpandTreeForGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExpandTreeForGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExpandTreeForGlossaryResponse")
	}
	return
}

// expandTreeForGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) expandTreeForGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/actions/expandTree", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExpandTreeForGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/ExpandTreeForGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "ExpandTreeForGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportGlossary Export the glossary and the terms and return the exported glossary as csv or json.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ExportGlossary.go.html to see an example of how to use ExportGlossary API.
func (client DataCatalogClient) ExportGlossary(ctx context.Context, request ExportGlossaryRequest) (response ExportGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.exportGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportGlossaryResponse")
	}
	return
}

// exportGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) exportGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/ExportGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "ExportGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FetchEntityLineage Returns lineage for a given entity object.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/FetchEntityLineage.go.html to see an example of how to use FetchEntityLineage API.
// A default retry strategy applies to this operation FetchEntityLineage()
func (client DataCatalogClient) FetchEntityLineage(ctx context.Context, request FetchEntityLineageRequest) (response FetchEntityLineageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.fetchEntityLineage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FetchEntityLineageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FetchEntityLineageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FetchEntityLineageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FetchEntityLineageResponse")
	}
	return
}

// fetchEntityLineage implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) fetchEntityLineage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/actions/fetchLineage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FetchEntityLineageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Entity/FetchEntityLineage"
		err = common.PostProcessServiceError(err, "DataCatalog", "FetchEntityLineage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAttribute Gets a specific entity attribute by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetAttribute.go.html to see an example of how to use GetAttribute API.
// A default retry strategy applies to this operation GetAttribute()
func (client DataCatalogClient) GetAttribute(ctx context.Context, request GetAttributeRequest) (response GetAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAttributeResponse")
	}
	return
}

// getAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Attribute/GetAttribute"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAttributeTag Gets a specific entity attribute tag by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetAttributeTag.go.html to see an example of how to use GetAttributeTag API.
// A default retry strategy applies to this operation GetAttributeTag()
func (client DataCatalogClient) GetAttributeTag(ctx context.Context, request GetAttributeTagRequest) (response GetAttributeTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAttributeTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAttributeTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAttributeTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAttributeTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAttributeTagResponse")
	}
	return
}

// getAttributeTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getAttributeTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAttributeTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/AttributeTag/GetAttributeTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetAttributeTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCatalog Gets a data catalog by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetCatalog.go.html to see an example of how to use GetCatalog API.
// A default retry strategy applies to this operation GetCatalog()
func (client DataCatalogClient) GetCatalog(ctx context.Context, request GetCatalogRequest) (response GetCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogResponse")
	}
	return
}

// getCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/GetCatalog"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCatalogPrivateEndpoint Gets a specific private reverse connection by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetCatalogPrivateEndpoint.go.html to see an example of how to use GetCatalogPrivateEndpoint API.
// A default retry strategy applies to this operation GetCatalogPrivateEndpoint()
func (client DataCatalogClient) GetCatalogPrivateEndpoint(ctx context.Context, request GetCatalogPrivateEndpointRequest) (response GetCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCatalogPrivateEndpointResponse")
	}
	return
}

// getCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCatalogPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpoint/GetCatalogPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetCatalogPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConnection Gets a specific data asset connection by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetConnection.go.html to see an example of how to use GetConnection API.
// A default retry strategy applies to this operation GetConnection()
func (client DataCatalogClient) GetConnection(ctx context.Context, request GetConnectionRequest) (response GetConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionResponse")
	}
	return
}

// getConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Connection/GetConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCustomProperty Gets a specific custom property for the given key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetCustomProperty.go.html to see an example of how to use GetCustomProperty API.
// A default retry strategy applies to this operation GetCustomProperty()
func (client DataCatalogClient) GetCustomProperty(ctx context.Context, request GetCustomPropertyRequest) (response GetCustomPropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCustomProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCustomPropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCustomPropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCustomPropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCustomPropertyResponse")
	}
	return
}

// getCustomProperty implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getCustomProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/namespaces/{namespaceId}/customProperties/{customPropertyKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCustomPropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CustomProperty/GetCustomProperty"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetCustomProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataAsset Gets a specific data asset for the given key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetDataAsset.go.html to see an example of how to use GetDataAsset API.
// A default retry strategy applies to this operation GetDataAsset()
func (client DataCatalogClient) GetDataAsset(ctx context.Context, request GetDataAssetRequest) (response GetDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataAssetResponse")
	}
	return
}

// getDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/GetDataAsset"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetDataAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDataAssetTag Gets a specific data asset tag by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetDataAssetTag.go.html to see an example of how to use GetDataAssetTag API.
// A default retry strategy applies to this operation GetDataAssetTag()
func (client DataCatalogClient) GetDataAssetTag(ctx context.Context, request GetDataAssetTagRequest) (response GetDataAssetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDataAssetTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDataAssetTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDataAssetTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDataAssetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDataAssetTagResponse")
	}
	return
}

// getDataAssetTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getDataAssetTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDataAssetTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAssetTag/GetDataAssetTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetDataAssetTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEntity Gets a specific data entity by key for a data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetEntity.go.html to see an example of how to use GetEntity API.
// A default retry strategy applies to this operation GetEntity()
func (client DataCatalogClient) GetEntity(ctx context.Context, request GetEntityRequest) (response GetEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEntityResponse")
	}
	return
}

// getEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Entity/GetEntity"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetEntity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEntityTag Gets a specific entity tag by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetEntityTag.go.html to see an example of how to use GetEntityTag API.
// A default retry strategy applies to this operation GetEntityTag()
func (client DataCatalogClient) GetEntityTag(ctx context.Context, request GetEntityTagRequest) (response GetEntityTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEntityTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEntityTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEntityTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEntityTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEntityTagResponse")
	}
	return
}

// getEntityTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getEntityTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEntityTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/EntityTag/GetEntityTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetEntityTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFolder Gets a specific data asset folder by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetFolder.go.html to see an example of how to use GetFolder API.
// A default retry strategy applies to this operation GetFolder()
func (client DataCatalogClient) GetFolder(ctx context.Context, request GetFolderRequest) (response GetFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFolderResponse")
	}
	return
}

// getFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Folder/GetFolder"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetFolder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFolderTag Gets a specific folder tag by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetFolderTag.go.html to see an example of how to use GetFolderTag API.
// A default retry strategy applies to this operation GetFolderTag()
func (client DataCatalogClient) GetFolderTag(ctx context.Context, request GetFolderTagRequest) (response GetFolderTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFolderTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFolderTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFolderTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFolderTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFolderTagResponse")
	}
	return
}

// getFolderTag implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getFolderTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags/{tagKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFolderTagResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/FolderTag/GetFolderTag"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetFolderTag", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetGlossary Gets a specific glossary by key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetGlossary.go.html to see an example of how to use GetGlossary API.
// A default retry strategy applies to this operation GetGlossary()
func (client DataCatalogClient) GetGlossary(ctx context.Context, request GetGlossaryRequest) (response GetGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGlossaryResponse")
	}
	return
}

// getGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/GetGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJob Gets a specific job by key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJob.go.html to see an example of how to use GetJob API.
// A default retry strategy applies to this operation GetJob()
func (client DataCatalogClient) GetJob(ctx context.Context, request GetJobRequest) (response GetJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobResponse")
	}
	return
}

// getJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Job/GetJob"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobDefinition Gets a specific job definition by key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJobDefinition.go.html to see an example of how to use GetJobDefinition API.
// A default retry strategy applies to this operation GetJobDefinition()
func (client DataCatalogClient) GetJobDefinition(ctx context.Context, request GetJobDefinitionRequest) (response GetJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobDefinitionResponse")
	}
	return
}

// getJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobDefinition/GetJobDefinition"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetJobDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobExecution Gets a specific job execution by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJobExecution.go.html to see an example of how to use GetJobExecution API.
// A default retry strategy applies to this operation GetJobExecution()
func (client DataCatalogClient) GetJobExecution(ctx context.Context, request GetJobExecutionRequest) (response GetJobExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobExecutionResponse")
	}
	return
}

// getJobExecution implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobExecution/GetJobExecution"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetJobExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobLog Gets a specific job log by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJobLog.go.html to see an example of how to use GetJobLog API.
// A default retry strategy applies to this operation GetJobLog()
func (client DataCatalogClient) GetJobLog(ctx context.Context, request GetJobLogRequest) (response GetJobLogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobLog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobLogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobLogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobLogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobLogResponse")
	}
	return
}

// getJobLog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobLog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/logs/{jobLogKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobLogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobLog/GetJobLog"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetJobLog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetJobMetrics Gets a specific job metric by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetJobMetrics.go.html to see an example of how to use GetJobMetrics API.
// A default retry strategy applies to this operation GetJobMetrics()
func (client DataCatalogClient) GetJobMetrics(ctx context.Context, request GetJobMetricsRequest) (response GetJobMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getJobMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetJobMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetJobMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetJobMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetJobMetricsResponse")
	}
	return
}

// getJobMetrics implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getJobMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/metrics/{jobMetricsKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetJobMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobMetric/GetJobMetrics"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetJobMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMetastore Gets a metastore by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetMetastore.go.html to see an example of how to use GetMetastore API.
// A default retry strategy applies to this operation GetMetastore()
func (client DataCatalogClient) GetMetastore(ctx context.Context, request GetMetastoreRequest) (response GetMetastoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMetastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMetastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMetastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMetastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMetastoreResponse")
	}
	return
}

// getMetastore implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getMetastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/metastores/{metastoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMetastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Metastore/GetMetastore"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetMetastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNamespace Gets a specific namespace for the given key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetNamespace.go.html to see an example of how to use GetNamespace API.
// A default retry strategy applies to this operation GetNamespace()
func (client DataCatalogClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceResponse")
	}
	return
}

// getNamespace implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/namespaces/{namespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Namespace/GetNamespace"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPattern Gets a specific pattern for the given key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetPattern.go.html to see an example of how to use GetPattern API.
// A default retry strategy applies to this operation GetPattern()
func (client DataCatalogClient) GetPattern(ctx context.Context, request GetPatternRequest) (response GetPatternResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPattern, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPatternResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPatternResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPatternResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPatternResponse")
	}
	return
}

// getPattern implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getPattern(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/patterns/{patternKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPatternResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Pattern/GetPattern"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetPattern", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTerm Gets a specific glossary term by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetTerm.go.html to see an example of how to use GetTerm API.
// A default retry strategy applies to this operation GetTerm()
func (client DataCatalogClient) GetTerm(ctx context.Context, request GetTermRequest) (response GetTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTermResponse")
	}
	return
}

// getTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Term/GetTerm"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTermRelationship Gets a specific glossary term relationship by key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetTermRelationship.go.html to see an example of how to use GetTermRelationship API.
// A default retry strategy applies to this operation GetTermRelationship()
func (client DataCatalogClient) GetTermRelationship(ctx context.Context, request GetTermRelationshipRequest) (response GetTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTermRelationshipResponse")
	}
	return
}

// getTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getTermRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships/{termRelationshipKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTermRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/TermRelationship/GetTermRelationship"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetTermRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetType Gets a specific type by key within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetType.go.html to see an example of how to use GetType API.
// A default retry strategy applies to this operation GetType()
func (client DataCatalogClient) GetType(ctx context.Context, request GetTypeRequest) (response GetTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTypeResponse")
	}
	return
}

// getType implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) getType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/types/{typeKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Type/GetType"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetType", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given OCID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DataCatalogClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DataCatalogClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "DataCatalog", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportConnection Import new connection for this data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ImportConnection.go.html to see an example of how to use ImportConnection API.
func (client DataCatalogClient) ImportConnection(ctx context.Context, request ImportConnectionRequest) (response ImportConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportConnectionResponse")
	}
	return
}

// importConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) importConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/importConnection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/ImportConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "ImportConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportDataAsset Import technical objects to a Data Asset
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ImportDataAsset.go.html to see an example of how to use ImportDataAsset API.
func (client DataCatalogClient) ImportDataAsset(ctx context.Context, request ImportDataAssetRequest) (response ImportDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportDataAssetResponse")
	}
	return
}

// importDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) importDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/import", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/ImportDataAsset"
		err = common.PostProcessServiceError(err, "DataCatalog", "ImportDataAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportGlossary Import the glossary and the terms from csv or json files and return the imported glossary resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ImportGlossary.go.html to see an example of how to use ImportGlossary API.
func (client DataCatalogClient) ImportGlossary(ctx context.Context, request ImportGlossaryRequest) (response ImportGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportGlossaryResponse")
	}
	return
}

// importGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) importGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/glossaries/{glossaryKey}/actions/import", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/ImportGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "ImportGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAggregatedPhysicalEntities List the physical entities aggregated by this logical entity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListAggregatedPhysicalEntities.go.html to see an example of how to use ListAggregatedPhysicalEntities API.
// A default retry strategy applies to this operation ListAggregatedPhysicalEntities()
func (client DataCatalogClient) ListAggregatedPhysicalEntities(ctx context.Context, request ListAggregatedPhysicalEntitiesRequest) (response ListAggregatedPhysicalEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAggregatedPhysicalEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAggregatedPhysicalEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAggregatedPhysicalEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAggregatedPhysicalEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAggregatedPhysicalEntitiesResponse")
	}
	return
}

// listAggregatedPhysicalEntities implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listAggregatedPhysicalEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/actions/listAggregatedPhysicalEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAggregatedPhysicalEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Entity/ListAggregatedPhysicalEntities"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListAggregatedPhysicalEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAttributeTags Returns a list of all tags for an entity attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListAttributeTags.go.html to see an example of how to use ListAttributeTags API.
// A default retry strategy applies to this operation ListAttributeTags()
func (client DataCatalogClient) ListAttributeTags(ctx context.Context, request ListAttributeTagsRequest) (response ListAttributeTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttributeTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAttributeTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAttributeTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttributeTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttributeTagsResponse")
	}
	return
}

// listAttributeTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listAttributeTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAttributeTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/AttributeTagCollection/ListAttributeTags"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListAttributeTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAttributes Returns a list of all attributes of an data entity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListAttributes.go.html to see an example of how to use ListAttributes API.
// A default retry strategy applies to this operation ListAttributes()
func (client DataCatalogClient) ListAttributes(ctx context.Context, request ListAttributesRequest) (response ListAttributesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttributesResponse")
	}
	return
}

// listAttributes implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listAttributes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAttributesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/AttributeCollection/ListAttributes"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListAttributes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCatalogPrivateEndpoints Returns a list of all the catalog private endpoints in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListCatalogPrivateEndpoints.go.html to see an example of how to use ListCatalogPrivateEndpoints API.
// A default retry strategy applies to this operation ListCatalogPrivateEndpoints()
func (client DataCatalogClient) ListCatalogPrivateEndpoints(ctx context.Context, request ListCatalogPrivateEndpointsRequest) (response ListCatalogPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCatalogPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCatalogPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogPrivateEndpointsResponse")
	}
	return
}

// listCatalogPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listCatalogPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCatalogPrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpointSummary/ListCatalogPrivateEndpoints"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListCatalogPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCatalogs Returns a list of all the data catalogs in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListCatalogs.go.html to see an example of how to use ListCatalogs API.
// A default retry strategy applies to this operation ListCatalogs()
func (client DataCatalogClient) ListCatalogs(ctx context.Context, request ListCatalogsRequest) (response ListCatalogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCatalogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCatalogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCatalogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCatalogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCatalogsResponse")
	}
	return
}

// listCatalogs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listCatalogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCatalogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogSummary/ListCatalogs"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListCatalogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnections Returns a list of all Connections for a data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListConnections.go.html to see an example of how to use ListConnections API.
// A default retry strategy applies to this operation ListConnections()
func (client DataCatalogClient) ListConnections(ctx context.Context, request ListConnectionsRequest) (response ListConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionsResponse")
	}
	return
}

// listConnections implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/ConnectionCollection/ListConnections"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListConnections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCustomProperties Returns a list of custom properties within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListCustomProperties.go.html to see an example of how to use ListCustomProperties API.
// A default retry strategy applies to this operation ListCustomProperties()
func (client DataCatalogClient) ListCustomProperties(ctx context.Context, request ListCustomPropertiesRequest) (response ListCustomPropertiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCustomProperties, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCustomPropertiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCustomPropertiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCustomPropertiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCustomPropertiesResponse")
	}
	return
}

// listCustomProperties implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listCustomProperties(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/namespaces/{namespaceId}/customProperties", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCustomPropertiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CustomProperty/ListCustomProperties"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListCustomProperties", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataAssetTags Returns a list of all tags for a data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListDataAssetTags.go.html to see an example of how to use ListDataAssetTags API.
// A default retry strategy applies to this operation ListDataAssetTags()
func (client DataCatalogClient) ListDataAssetTags(ctx context.Context, request ListDataAssetTagsRequest) (response ListDataAssetTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssetTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataAssetTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataAssetTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetTagsResponse")
	}
	return
}

// listDataAssetTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDataAssetTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataAssetTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAssetTagCollection/ListDataAssetTags"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListDataAssetTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDataAssets Returns a list of data assets within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListDataAssets.go.html to see an example of how to use ListDataAssets API.
// A default retry strategy applies to this operation ListDataAssets()
func (client DataCatalogClient) ListDataAssets(ctx context.Context, request ListDataAssetsRequest) (response ListDataAssetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDataAssets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDataAssetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDataAssetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDataAssetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDataAssetsResponse")
	}
	return
}

// listDataAssets implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDataAssets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDataAssetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAssetCollection/ListDataAssets"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListDataAssets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDerivedLogicalEntities List logical entities derived from this pattern.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListDerivedLogicalEntities.go.html to see an example of how to use ListDerivedLogicalEntities API.
// A default retry strategy applies to this operation ListDerivedLogicalEntities()
func (client DataCatalogClient) ListDerivedLogicalEntities(ctx context.Context, request ListDerivedLogicalEntitiesRequest) (response ListDerivedLogicalEntitiesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listDerivedLogicalEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDerivedLogicalEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDerivedLogicalEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDerivedLogicalEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDerivedLogicalEntitiesResponse")
	}
	return
}

// listDerivedLogicalEntities implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listDerivedLogicalEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/patterns/{patternKey}/actions/listDerivedLogicalEntities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDerivedLogicalEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Pattern/ListDerivedLogicalEntities"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListDerivedLogicalEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEntities Returns a list of all entities of a data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListEntities.go.html to see an example of how to use ListEntities API.
// A default retry strategy applies to this operation ListEntities()
func (client DataCatalogClient) ListEntities(ctx context.Context, request ListEntitiesRequest) (response ListEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntitiesResponse")
	}
	return
}

// listEntities implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listEntities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Entity/ListEntities"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListEntities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEntityTags Returns a list of all tags for a data entity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListEntityTags.go.html to see an example of how to use ListEntityTags API.
// A default retry strategy applies to this operation ListEntityTags()
func (client DataCatalogClient) ListEntityTags(ctx context.Context, request ListEntityTagsRequest) (response ListEntityTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntityTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntityTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntityTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntityTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntityTagsResponse")
	}
	return
}

// listEntityTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listEntityTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEntityTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/EntityTagCollection/ListEntityTags"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListEntityTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFolderTags Returns a list of all tags for a folder.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListFolderTags.go.html to see an example of how to use ListFolderTags API.
// A default retry strategy applies to this operation ListFolderTags()
func (client DataCatalogClient) ListFolderTags(ctx context.Context, request ListFolderTagsRequest) (response ListFolderTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFolderTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFolderTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFolderTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFolderTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFolderTagsResponse")
	}
	return
}

// listFolderTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listFolderTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFolderTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/FolderTagCollection/ListFolderTags"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListFolderTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFolders Returns a list of all folders.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListFolders.go.html to see an example of how to use ListFolders API.
// A default retry strategy applies to this operation ListFolders()
func (client DataCatalogClient) ListFolders(ctx context.Context, request ListFoldersRequest) (response ListFoldersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFolders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFoldersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFoldersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFoldersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFoldersResponse")
	}
	return
}

// listFolders implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listFolders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFoldersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/FolderCollection/ListFolders"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListFolders", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListGlossaries Returns a list of all glossaries within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListGlossaries.go.html to see an example of how to use ListGlossaries API.
// A default retry strategy applies to this operation ListGlossaries()
func (client DataCatalogClient) ListGlossaries(ctx context.Context, request ListGlossariesRequest) (response ListGlossariesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGlossaries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGlossariesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGlossariesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGlossariesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGlossariesResponse")
	}
	return
}

// listGlossaries implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listGlossaries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGlossariesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/ListGlossaries"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListGlossaries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobDefinitions Returns a list of job definitions within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobDefinitions.go.html to see an example of how to use ListJobDefinitions API.
// A default retry strategy applies to this operation ListJobDefinitions()
func (client DataCatalogClient) ListJobDefinitions(ctx context.Context, request ListJobDefinitionsRequest) (response ListJobDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobDefinitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobDefinitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobDefinitionsResponse")
	}
	return
}

// listJobDefinitions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobDefinitions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobDefinitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobDefinitionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobDefinitionCollection/ListJobDefinitions"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListJobDefinitions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobExecutions Returns a list of job executions for a job.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobExecutions.go.html to see an example of how to use ListJobExecutions API.
// A default retry strategy applies to this operation ListJobExecutions()
func (client DataCatalogClient) ListJobExecutions(ctx context.Context, request ListJobExecutionsRequest) (response ListJobExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobExecutionsResponse")
	}
	return
}

// listJobExecutions implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobExecutionCollection/ListJobExecutions"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListJobExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobLogs Returns a list of job logs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobLogs.go.html to see an example of how to use ListJobLogs API.
// A default retry strategy applies to this operation ListJobLogs()
func (client DataCatalogClient) ListJobLogs(ctx context.Context, request ListJobLogsRequest) (response ListJobLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobLogsResponse")
	}
	return
}

// listJobLogs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobLogCollection/ListJobLogs"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListJobLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobMetrics Returns a list of job metrics.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobMetrics.go.html to see an example of how to use ListJobMetrics API.
// A default retry strategy applies to this operation ListJobMetrics()
func (client DataCatalogClient) ListJobMetrics(ctx context.Context, request ListJobMetricsRequest) (response ListJobMetricsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobMetrics, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobMetricsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobMetricsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobMetricsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobMetricsResponse")
	}
	return
}

// listJobMetrics implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobMetrics(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs/{jobKey}/executions/{jobExecutionKey}/metrics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobMetricsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobMetricCollection/ListJobMetrics"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListJobMetrics", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListJobs Returns a list of jobs within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListJobs.go.html to see an example of how to use ListJobs API.
// A default retry strategy applies to this operation ListJobs()
func (client DataCatalogClient) ListJobs(ctx context.Context, request ListJobsRequest) (response ListJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListJobsResponse")
	}
	return
}

// listJobs implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/jobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobCollection/ListJobs"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMetastores Returns a list of all metastores in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListMetastores.go.html to see an example of how to use ListMetastores API.
// A default retry strategy applies to this operation ListMetastores()
func (client DataCatalogClient) ListMetastores(ctx context.Context, request ListMetastoresRequest) (response ListMetastoresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMetastores, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMetastoresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMetastoresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMetastoresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMetastoresResponse")
	}
	return
}

// listMetastores implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listMetastores(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/metastores", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMetastoresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/MetastoreSummary/ListMetastores"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListMetastores", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNamespaces Returns a list of namespaces within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListNamespaces.go.html to see an example of how to use ListNamespaces API.
// A default retry strategy applies to this operation ListNamespaces()
func (client DataCatalogClient) ListNamespaces(ctx context.Context, request ListNamespacesRequest) (response ListNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNamespacesResponse")
	}
	return
}

// listNamespaces implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listNamespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/namespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNamespacesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Namespace/ListNamespaces"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListNamespaces", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPatterns Returns a list of patterns within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListPatterns.go.html to see an example of how to use ListPatterns API.
// A default retry strategy applies to this operation ListPatterns()
func (client DataCatalogClient) ListPatterns(ctx context.Context, request ListPatternsRequest) (response ListPatternsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPatterns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPatternsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPatternsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPatternsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPatternsResponse")
	}
	return
}

// listPatterns implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listPatterns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/patterns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPatternsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Pattern/ListPatterns"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListPatterns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRules Returns a list of all rules of a data entity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListRules.go.html to see an example of how to use ListRules API.
// A default retry strategy applies to this operation ListRules()
func (client DataCatalogClient) ListRules(ctx context.Context, request ListRulesRequest) (response ListRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRulesResponse")
	}
	return
}

// listRules implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/rules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/RuleSummary/ListRules"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTags Returns a list of all user created tags in the system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListTags.go.html to see an example of how to use ListTags API.
// A default retry strategy applies to this operation ListTags()
func (client DataCatalogClient) ListTags(ctx context.Context, request ListTagsRequest) (response ListTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTagsResponse")
	}
	return
}

// listTags implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTagsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Term/ListTags"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListTags", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTermRelationships Returns a list of all term relationships within a glossary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListTermRelationships.go.html to see an example of how to use ListTermRelationships API.
// A default retry strategy applies to this operation ListTermRelationships()
func (client DataCatalogClient) ListTermRelationships(ctx context.Context, request ListTermRelationshipsRequest) (response ListTermRelationshipsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTermRelationships, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTermRelationshipsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTermRelationshipsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTermRelationshipsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTermRelationshipsResponse")
	}
	return
}

// listTermRelationships implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTermRelationships(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTermRelationshipsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/TermRelationship/ListTermRelationships"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListTermRelationships", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTerms Returns a list of all terms within a glossary.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListTerms.go.html to see an example of how to use ListTerms API.
// A default retry strategy applies to this operation ListTerms()
func (client DataCatalogClient) ListTerms(ctx context.Context, request ListTermsRequest) (response ListTermsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTerms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTermsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTermsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTermsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTermsResponse")
	}
	return
}

// listTerms implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTerms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTermsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Term/ListTerms"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListTerms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTypes Returns a list of all types within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListTypes.go.html to see an example of how to use ListTypes API.
// A default retry strategy applies to this operation ListTypes()
func (client DataCatalogClient) ListTypes(ctx context.Context, request ListTypesRequest) (response ListTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTypesResponse")
	}
	return
}

// listTypes implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) listTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/catalogs/{catalogId}/types", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/TypeCollection/ListTypes"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DataCatalogClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DataCatalogClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DataCatalogClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DataCatalogClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/WorkRequestLog/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DataCatalogClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DataCatalogClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "DataCatalog", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ObjectStats Returns stats on objects by type in the repository.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ObjectStats.go.html to see an example of how to use ObjectStats API.
// A default retry strategy applies to this operation ObjectStats()
func (client DataCatalogClient) ObjectStats(ctx context.Context, request ObjectStatsRequest) (response ObjectStatsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.objectStats, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ObjectStatsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ObjectStatsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ObjectStatsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ObjectStatsResponse")
	}
	return
}

// objectStats implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) objectStats(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/objectStats", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ObjectStatsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/ObjectStats"
		err = common.PostProcessServiceError(err, "DataCatalog", "ObjectStats", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ParseConnection Parse data asset references through connections from this data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ParseConnection.go.html to see an example of how to use ParseConnection API.
func (client DataCatalogClient) ParseConnection(ctx context.Context, request ParseConnectionRequest) (response ParseConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.parseConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ParseConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ParseConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ParseConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ParseConnectionResponse")
	}
	return
}

// parseConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) parseConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/parseConnection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ParseConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/ParseConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "ParseConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ProcessRecommendation Act on a recommendation. A recommendation can be accepted or rejected. For example, if a recommendation of type LINK_GLOSSARY_TERM
// is accepted, the system will link the source object (e.g. an attribute) to a target glossary term.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ProcessRecommendation.go.html to see an example of how to use ProcessRecommendation API.
func (client DataCatalogClient) ProcessRecommendation(ctx context.Context, request ProcessRecommendationRequest) (response ProcessRecommendationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.processRecommendation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ProcessRecommendationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ProcessRecommendationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ProcessRecommendationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ProcessRecommendationResponse")
	}
	return
}

// processRecommendation implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) processRecommendation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/processRecommendation", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ProcessRecommendationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/ProcessRecommendation"
		err = common.PostProcessServiceError(err, "DataCatalog", "ProcessRecommendation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Recommendations Returns a list of recommendations for the given object and recommendation type.
// By default, it will return inferred recommendations for review. The optional query param 'RecommendationStatus' can be set,
// to return only recommendations having that status.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/Recommendations.go.html to see an example of how to use Recommendations API.
func (client DataCatalogClient) Recommendations(ctx context.Context, request RecommendationsRequest) (response RecommendationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.recommendations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RecommendationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RecommendationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RecommendationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RecommendationsResponse")
	}
	return
}

// recommendations implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) recommendations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/getRecommendations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RecommendationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/Recommendations"
		err = common.PostProcessServiceError(err, "DataCatalog", "Recommendations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveCatalogLock Removes a lock from a Catalog resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/RemoveCatalogLock.go.html to see an example of how to use RemoveCatalogLock API.
func (client DataCatalogClient) RemoveCatalogLock(ctx context.Context, request RemoveCatalogLockRequest) (response RemoveCatalogLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeCatalogLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveCatalogLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveCatalogLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveCatalogLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveCatalogLockResponse")
	}
	return
}

// removeCatalogLock implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) removeCatalogLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveCatalogLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/RemoveCatalogLock"
		err = common.PostProcessServiceError(err, "DataCatalog", "RemoveCatalogLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveCatalogPrivateEndpointLock Removes a lock from a CatalogPrivateEndpoint resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/RemoveCatalogPrivateEndpointLock.go.html to see an example of how to use RemoveCatalogPrivateEndpointLock API.
func (client DataCatalogClient) RemoveCatalogPrivateEndpointLock(ctx context.Context, request RemoveCatalogPrivateEndpointLockRequest) (response RemoveCatalogPrivateEndpointLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeCatalogPrivateEndpointLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveCatalogPrivateEndpointLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveCatalogPrivateEndpointLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveCatalogPrivateEndpointLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveCatalogPrivateEndpointLockResponse")
	}
	return
}

// removeCatalogPrivateEndpointLock implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) removeCatalogPrivateEndpointLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveCatalogPrivateEndpointLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpoint/RemoveCatalogPrivateEndpointLock"
		err = common.PostProcessServiceError(err, "DataCatalog", "RemoveCatalogPrivateEndpointLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveDataSelectorPatterns Remove data selector pattern from the data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/RemoveDataSelectorPatterns.go.html to see an example of how to use RemoveDataSelectorPatterns API.
func (client DataCatalogClient) RemoveDataSelectorPatterns(ctx context.Context, request RemoveDataSelectorPatternsRequest) (response RemoveDataSelectorPatternsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeDataSelectorPatterns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDataSelectorPatternsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDataSelectorPatternsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDataSelectorPatternsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDataSelectorPatternsResponse")
	}
	return
}

// removeDataSelectorPatterns implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) removeDataSelectorPatterns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/removeDataSelectorPatterns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDataSelectorPatternsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/RemoveDataSelectorPatterns"
		err = common.PostProcessServiceError(err, "DataCatalog", "RemoveDataSelectorPatterns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveMetastoreLock Removes a lock from a Metastore resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/RemoveMetastoreLock.go.html to see an example of how to use RemoveMetastoreLock API.
func (client DataCatalogClient) RemoveMetastoreLock(ctx context.Context, request RemoveMetastoreLockRequest) (response RemoveMetastoreLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeMetastoreLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveMetastoreLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveMetastoreLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveMetastoreLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveMetastoreLockResponse")
	}
	return
}

// removeMetastoreLock implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) removeMetastoreLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/metastores/{metastoreId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveMetastoreLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Metastore/RemoveMetastoreLock"
		err = common.PostProcessServiceError(err, "DataCatalog", "RemoveMetastoreLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchCriteria Returns a list of search results within a data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/SearchCriteria.go.html to see an example of how to use SearchCriteria API.
// A default retry strategy applies to this operation SearchCriteria()
func (client DataCatalogClient) SearchCriteria(ctx context.Context, request SearchCriteriaRequest) (response SearchCriteriaResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchCriteria, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchCriteriaResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchCriteriaResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchCriteriaResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchCriteriaResponse")
	}
	return
}

// searchCriteria implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) searchCriteria(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/search", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchCriteriaResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/SearchResult/SearchCriteria"
		err = common.PostProcessServiceError(err, "DataCatalog", "SearchCriteria", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SuggestMatches Returns a list of potential string matches for a given input string.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/SuggestMatches.go.html to see an example of how to use SuggestMatches API.
// A default retry strategy applies to this operation SuggestMatches()
func (client DataCatalogClient) SuggestMatches(ctx context.Context, request SuggestMatchesRequest) (response SuggestMatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.suggestMatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SuggestMatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SuggestMatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SuggestMatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SuggestMatchesResponse")
	}
	return
}

// suggestMatches implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) suggestMatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/suggest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SuggestMatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/SuggestResults/SuggestMatches"
		err = common.PostProcessServiceError(err, "DataCatalog", "SuggestMatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SynchronousExportDataAsset Export technical objects from a Data Asset
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/SynchronousExportDataAsset.go.html to see an example of how to use SynchronousExportDataAsset API.
func (client DataCatalogClient) SynchronousExportDataAsset(ctx context.Context, request SynchronousExportDataAssetRequest) (response SynchronousExportDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.synchronousExportDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SynchronousExportDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SynchronousExportDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SynchronousExportDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SynchronousExportDataAssetResponse")
	}
	return
}

// synchronousExportDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) synchronousExportDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/synchronousExport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SynchronousExportDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/SynchronousExportDataAsset"
		err = common.PostProcessServiceError(err, "DataCatalog", "SynchronousExportDataAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestConnection Test the connection by connecting to the data asset using credentials in the metadata.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/TestConnection.go.html to see an example of how to use TestConnection API.
// A default retry strategy applies to this operation TestConnection()
func (client DataCatalogClient) TestConnection(ctx context.Context, request TestConnectionRequest) (response TestConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.testConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestConnectionResponse")
	}
	return
}

// testConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) testConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Connection/TestConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "TestConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAttribute Updates a specific data asset attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateAttribute.go.html to see an example of how to use UpdateAttribute API.
func (client DataCatalogClient) UpdateAttribute(ctx context.Context, request UpdateAttributeRequest) (response UpdateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAttributeResponse")
	}
	return
}

// updateAttribute implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}/attributes/{attributeKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Attribute/UpdateAttribute"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCatalog Updates the data catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateCatalog.go.html to see an example of how to use UpdateCatalog API.
func (client DataCatalogClient) UpdateCatalog(ctx context.Context, request UpdateCatalogRequest) (response UpdateCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCatalogResponse")
	}
	return
}

// updateCatalog implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/UpdateCatalog"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCatalogPrivateEndpoint Updates the private reverse connection endpoint.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateCatalogPrivateEndpoint.go.html to see an example of how to use UpdateCatalogPrivateEndpoint API.
func (client DataCatalogClient) UpdateCatalogPrivateEndpoint(ctx context.Context, request UpdateCatalogPrivateEndpointRequest) (response UpdateCatalogPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCatalogPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCatalogPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCatalogPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCatalogPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCatalogPrivateEndpointResponse")
	}
	return
}

// updateCatalogPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateCatalogPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogPrivateEndpoints/{catalogPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCatalogPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CatalogPrivateEndpoint/UpdateCatalogPrivateEndpoint"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateCatalogPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConnection Updates a specific connection of a data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateConnection.go.html to see an example of how to use UpdateConnection API.
func (client DataCatalogClient) UpdateConnection(ctx context.Context, request UpdateConnectionRequest) (response UpdateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConnectionResponse")
	}
	return
}

// updateConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Connection/UpdateConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCustomProperty Updates a specific custom property identified by the given key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateCustomProperty.go.html to see an example of how to use UpdateCustomProperty API.
func (client DataCatalogClient) UpdateCustomProperty(ctx context.Context, request UpdateCustomPropertyRequest) (response UpdateCustomPropertyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCustomProperty, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCustomPropertyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCustomPropertyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCustomPropertyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCustomPropertyResponse")
	}
	return
}

// updateCustomProperty implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateCustomProperty(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/namespaces/{namespaceId}/customProperties/{customPropertyKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCustomPropertyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/CustomProperty/UpdateCustomProperty"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateCustomProperty", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDataAsset Updates a specific data asset identified by the given key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateDataAsset.go.html to see an example of how to use UpdateDataAsset API.
func (client DataCatalogClient) UpdateDataAsset(ctx context.Context, request UpdateDataAssetRequest) (response UpdateDataAssetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDataAsset, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDataAssetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDataAssetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDataAssetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDataAssetResponse")
	}
	return
}

// updateDataAsset implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateDataAsset(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDataAssetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/UpdateDataAsset"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateDataAsset", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEntity Updates a specific data entity.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateEntity.go.html to see an example of how to use UpdateEntity API.
func (client DataCatalogClient) UpdateEntity(ctx context.Context, request UpdateEntityRequest) (response UpdateEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEntityResponse")
	}
	return
}

// updateEntity implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateEntity(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/entities/{entityKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Entity/UpdateEntity"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateEntity", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFolder Updates a specific folder of a data asset.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateFolder.go.html to see an example of how to use UpdateFolder API.
func (client DataCatalogClient) UpdateFolder(ctx context.Context, request UpdateFolderRequest) (response UpdateFolderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFolder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFolderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFolderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFolderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFolderResponse")
	}
	return
}

// updateFolder implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateFolder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/folders/{folderKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFolderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Folder/UpdateFolder"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateFolder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateGlossary Updates a specific glossary identified by the given key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateGlossary.go.html to see an example of how to use UpdateGlossary API.
func (client DataCatalogClient) UpdateGlossary(ctx context.Context, request UpdateGlossaryRequest) (response UpdateGlossaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGlossary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGlossaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGlossaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGlossaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGlossaryResponse")
	}
	return
}

// updateGlossary implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateGlossary(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGlossaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Glossary/UpdateGlossary"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateGlossary", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJob Updates a specific job identified by the given key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateJob.go.html to see an example of how to use UpdateJob API.
func (client DataCatalogClient) UpdateJob(ctx context.Context, request UpdateJobRequest) (response UpdateJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobResponse")
	}
	return
}

// updateJob implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/jobs/{jobKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Job/UpdateJob"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateJobDefinition Update a specific job definition identified by the given key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateJobDefinition.go.html to see an example of how to use UpdateJobDefinition API.
func (client DataCatalogClient) UpdateJobDefinition(ctx context.Context, request UpdateJobDefinitionRequest) (response UpdateJobDefinitionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateJobDefinition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateJobDefinitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateJobDefinitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateJobDefinitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateJobDefinitionResponse")
	}
	return
}

// updateJobDefinition implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateJobDefinition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/jobDefinitions/{jobDefinitionKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateJobDefinitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/JobDefinition/UpdateJobDefinition"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateJobDefinition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMetastore Updates a metastore resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateMetastore.go.html to see an example of how to use UpdateMetastore API.
func (client DataCatalogClient) UpdateMetastore(ctx context.Context, request UpdateMetastoreRequest) (response UpdateMetastoreResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMetastore, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMetastoreResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMetastoreResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMetastoreResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMetastoreResponse")
	}
	return
}

// updateMetastore implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateMetastore(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/metastores/{metastoreId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMetastoreResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Metastore/UpdateMetastore"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateMetastore", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNamespace Updates a specific namespace identified by the given key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateNamespace.go.html to see an example of how to use UpdateNamespace API.
func (client DataCatalogClient) UpdateNamespace(ctx context.Context, request UpdateNamespaceRequest) (response UpdateNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNamespaceResponse")
	}
	return
}

// updateNamespace implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/namespaces/{namespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Namespace/UpdateNamespace"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePattern Updates a specific pattern identified by the given key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdatePattern.go.html to see an example of how to use UpdatePattern API.
func (client DataCatalogClient) UpdatePattern(ctx context.Context, request UpdatePatternRequest) (response UpdatePatternResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePattern, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePatternResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePatternResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePatternResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePatternResponse")
	}
	return
}

// updatePattern implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updatePattern(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/patterns/{patternKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePatternResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Pattern/UpdatePattern"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdatePattern", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTerm Updates a specific glossary term.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateTerm.go.html to see an example of how to use UpdateTerm API.
func (client DataCatalogClient) UpdateTerm(ctx context.Context, request UpdateTermRequest) (response UpdateTermResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTerm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTermResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTermResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTermResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTermResponse")
	}
	return
}

// updateTerm implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateTerm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTermResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Term/UpdateTerm"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateTerm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTermRelationship Updates a specific glossary term relationship.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UpdateTermRelationship.go.html to see an example of how to use UpdateTermRelationship API.
func (client DataCatalogClient) UpdateTermRelationship(ctx context.Context, request UpdateTermRelationshipRequest) (response UpdateTermRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTermRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTermRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTermRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTermRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTermRelationshipResponse")
	}
	return
}

// updateTermRelationship implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) updateTermRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/catalogs/{catalogId}/glossaries/{glossaryKey}/terms/{termKey}/termRelationships/{termRelationshipKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTermRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/TermRelationship/UpdateTermRelationship"
		err = common.PostProcessServiceError(err, "DataCatalog", "UpdateTermRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadCredentials Upload connection credentails and metadata for this connection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/UploadCredentials.go.html to see an example of how to use UploadCredentials API.
func (client DataCatalogClient) UploadCredentials(ctx context.Context, request UploadCredentialsRequest) (response UploadCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.uploadCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadCredentialsResponse")
	}
	return
}

// uploadCredentials implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) uploadCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/connections/{connectionKey}/actions/uploadCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Connection/UploadCredentials"
		err = common.PostProcessServiceError(err, "DataCatalog", "UploadCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Users Returns active users in the system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/Users.go.html to see an example of how to use Users API.
// A default retry strategy applies to this operation Users()
func (client DataCatalogClient) Users(ctx context.Context, request UsersRequest) (response UsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.users, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UsersResponse")
	}
	return
}

// users implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) users(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/actions/users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Catalog/Users"
		err = common.PostProcessServiceError(err, "DataCatalog", "Users", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateConnection Validate connection by connecting to the data asset using credentials in metadata.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ValidateConnection.go.html to see an example of how to use ValidateConnection API.
// A default retry strategy applies to this operation ValidateConnection()
func (client DataCatalogClient) ValidateConnection(ctx context.Context, request ValidateConnectionRequest) (response ValidateConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateConnectionResponse")
	}
	return
}

// validateConnection implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) validateConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/dataAssets/{dataAssetKey}/actions/validateConnection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/DataAsset/ValidateConnection"
		err = common.PostProcessServiceError(err, "DataCatalog", "ValidateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidatePattern Validate pattern by deriving file groups representing logical entities using the expression
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ValidatePattern.go.html to see an example of how to use ValidatePattern API.
// A default retry strategy applies to this operation ValidatePattern()
func (client DataCatalogClient) ValidatePattern(ctx context.Context, request ValidatePatternRequest) (response ValidatePatternResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validatePattern, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidatePatternResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidatePatternResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidatePatternResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidatePatternResponse")
	}
	return
}

// validatePattern implements the OCIOperation interface (enables retrying operations)
func (client DataCatalogClient) validatePattern(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/catalogs/{catalogId}/patterns/{patternKey}/actions/validate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidatePatternResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/data-catalog/20190325/Pattern/ValidatePattern"
		err = common.PostProcessServiceError(err, "DataCatalog", "ValidatePattern", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
