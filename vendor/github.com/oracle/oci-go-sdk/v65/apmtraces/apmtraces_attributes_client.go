// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// AttributesClient a client for Attributes
type AttributesClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAttributesClientWithConfigurationProvider Creates a new default Attributes client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAttributesClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AttributesClient, err error) {
	if enabled := common.CheckForEnabledServices("apmtraces"); !enabled {
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
	return newAttributesClientFromBaseClient(baseClient, provider)
}

// NewAttributesClientWithOboToken Creates a new default Attributes client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAttributesClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AttributesClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAttributesClientFromBaseClient(baseClient, configProvider)
}

func newAttributesClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AttributesClient, err error) {
	// Attributes service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Attributes"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AttributesClient{BaseClient: baseClient}
	client.BasePath = "20200630"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AttributesClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("apmtraces", "https://apm-trace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AttributesClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AttributesClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkActivateAttribute Activates a set of attributes for the given APM Domain.  The API is not case-sensitive.  Any duplicates present in the bulk activation
// request are deduplicated and only unique attributes are activated.  A maximum number of 700 string attributes and 100 numeric attributes
// can be activated in an APM Domain subject to the available string and numeric slots.  Once an attribute has been activated, it may take sometime
// for it to be appear in searches as ingest might not have picked up the changes or any associated caches might not have refreshed.  The
// bulk activation operation is atomic, and the operation succeeds only if all the attributes in the request have been processed successfully and they
// get a success status back.  If the processing of any attribute results in a processing or validation error, then none of the attributes in the bulk
// request are activated.  Attributes that are activated are unpinned by default if they are pinned.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/BulkActivateAttribute.go.html to see an example of how to use BulkActivateAttribute API.
func (client AttributesClient) BulkActivateAttribute(ctx context.Context, request BulkActivateAttributeRequest) (response BulkActivateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkActivateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkActivateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkActivateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkActivateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkActivateAttributeResponse")
	}
	return
}

// bulkActivateAttribute implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) bulkActivateAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/activateAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkActivateAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/BulkActivationStatus/BulkActivateAttribute"
		err = common.PostProcessServiceError(err, "Attributes", "BulkActivateAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkDeActivateAttribute Deactivates a set of attributes for the given APM Domain.  The API is case in-sensitive.  Any duplicates present in the bulk deactivation
// request are deduplicated and only unique attributes are deactivated.  A maximum number of 700 string attributes and 100 numeric attributes
// can be deactivated in an APM Domain subject to the available string and numeric slots.  Out of box attributes (Trace and Span) cannot be
// deactivated, and will result in a processing error.  Once an attribute has been deactivated, it may take sometime for it to disappear in
// searches as ingest might not have picked up the changes or any associated caches might not have refreshed.  The bulk deactivation
// operation is atomic, and the operation succeeds only if all the attributes in the request have been processed successfully and they get a success
// status back.  If the processing of any attribute results in a processing or validation error, then none of the attributes in the bulk request
// are deactivated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/BulkDeActivateAttribute.go.html to see an example of how to use BulkDeActivateAttribute API.
func (client AttributesClient) BulkDeActivateAttribute(ctx context.Context, request BulkDeActivateAttributeRequest) (response BulkDeActivateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkDeActivateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkDeActivateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkDeActivateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkDeActivateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkDeActivateAttributeResponse")
	}
	return
}

// bulkDeActivateAttribute implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) bulkDeActivateAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/deActivateAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkDeActivateAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/BulkDeActivationStatus/BulkDeActivateAttribute"
		err = common.PostProcessServiceError(err, "Attributes", "BulkDeActivateAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkPinAttribute Pin a set of attributes in the APM Domain.  Attributes that are marked pinned are not autoactivated by ingest.
// Attributes that are deactivated are pinned by default.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/BulkPinAttribute.go.html to see an example of how to use BulkPinAttribute API.
func (client AttributesClient) BulkPinAttribute(ctx context.Context, request BulkPinAttributeRequest) (response BulkPinAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkPinAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkPinAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkPinAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkPinAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkPinAttributeResponse")
	}
	return
}

// bulkPinAttribute implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) bulkPinAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/pinAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkPinAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/BulkPinStatus/BulkPinAttribute"
		err = common.PostProcessServiceError(err, "Attributes", "BulkPinAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUnpinAttribute Unpin a set of attributes in the APM Domain.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/BulkUnpinAttribute.go.html to see an example of how to use BulkUnpinAttribute API.
func (client AttributesClient) BulkUnpinAttribute(ctx context.Context, request BulkUnpinAttributeRequest) (response BulkUnpinAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkUnpinAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUnpinAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUnpinAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUnpinAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUnpinAttributeResponse")
	}
	return
}

// bulkUnpinAttribute implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) bulkUnpinAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/unPinAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUnpinAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/BulkUnpinStatus/BulkUnpinAttribute"
		err = common.PostProcessServiceError(err, "Attributes", "BulkUnpinAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUpdateAttribute Update a set of attribute properties in the APM Domain.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/BulkUpdateAttribute.go.html to see an example of how to use BulkUpdateAttribute API.
func (client AttributesClient) BulkUpdateAttribute(ctx context.Context, request BulkUpdateAttributeRequest) (response BulkUpdateAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkUpdateAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUpdateAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUpdateAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUpdateAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUpdateAttributeResponse")
	}
	return
}

// bulkUpdateAttribute implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) bulkUpdateAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/updateAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUpdateAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/BulkUpdateAttributeStatus/BulkUpdateAttribute"
		err = common.PostProcessServiceError(err, "Attributes", "BulkUpdateAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUpdateAttributeNotes Add or edit notes to a set of attributes in the APM Domain.  Notes can be added to either an active or an inactive attribute.  The
// notes will be preserved even if the attribute changes state (when an active attribute is deactivated or when an inactive attribute
// is activated).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/BulkUpdateAttributeNotes.go.html to see an example of how to use BulkUpdateAttributeNotes API.
func (client AttributesClient) BulkUpdateAttributeNotes(ctx context.Context, request BulkUpdateAttributeNotesRequest) (response BulkUpdateAttributeNotesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkUpdateAttributeNotes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUpdateAttributeNotesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUpdateAttributeNotesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUpdateAttributeNotesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUpdateAttributeNotesResponse")
	}
	return
}

// bulkUpdateAttributeNotes implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) bulkUpdateAttributeNotes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/updateNotes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUpdateAttributeNotesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/BulkUpdateNotesStatus/BulkUpdateAttributeNotes"
		err = common.PostProcessServiceError(err, "Attributes", "BulkUpdateAttributeNotes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStatusAutoActivate Get autoactivation status for a private data key or public data key in the APM Domain.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetStatusAutoActivate.go.html to see an example of how to use GetStatusAutoActivate API.
func (client AttributesClient) GetStatusAutoActivate(ctx context.Context, request GetStatusAutoActivateRequest) (response GetStatusAutoActivateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStatusAutoActivate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStatusAutoActivateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStatusAutoActivateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStatusAutoActivateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStatusAutoActivateResponse")
	}
	return
}

// getStatusAutoActivate implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) getStatusAutoActivate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/attributes/autoActivateStatus", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetStatusAutoActivateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/AutoActivateStatus/GetStatusAutoActivate"
		err = common.PostProcessServiceError(err, "Attributes", "GetStatusAutoActivate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutToggleAutoActivate Turn on or off autoactivate for private data key or public data key traffic a given APM Domain.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/PutToggleAutoActivate.go.html to see an example of how to use PutToggleAutoActivate API.
func (client AttributesClient) PutToggleAutoActivate(ctx context.Context, request PutToggleAutoActivateRequest) (response PutToggleAutoActivateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putToggleAutoActivate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutToggleAutoActivateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutToggleAutoActivateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutToggleAutoActivateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutToggleAutoActivateResponse")
	}
	return
}

// putToggleAutoActivate implements the OCIOperation interface (enables retrying operations)
func (client AttributesClient) putToggleAutoActivate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/attributes/actions/autoActivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PutToggleAutoActivateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/apm-trace-explorer/20200630/AutoActivateToggleStatus/PutToggleAutoActivate"
		err = common.PostProcessServiceError(err, "Attributes", "PutToggleAutoActivate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
