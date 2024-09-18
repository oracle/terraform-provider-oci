// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// CapacityManagementClient a client for CapacityManagement
type CapacityManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCapacityManagementClientWithConfigurationProvider Creates a new default CapacityManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCapacityManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CapacityManagementClient, err error) {
	if enabled := common.CheckForEnabledServices("capacitymanagement"); !enabled {
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
	return newCapacityManagementClientFromBaseClient(baseClient, provider)
}

// NewCapacityManagementClientWithOboToken Creates a new default CapacityManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCapacityManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CapacityManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCapacityManagementClientFromBaseClient(baseClient, configProvider)
}

func newCapacityManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CapacityManagementClient, err error) {
	// CapacityManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CapacityManagement"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CapacityManagementClient{BaseClient: baseClient}
	client.BasePath = "20231107"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CapacityManagementClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("capacitymanagement", "https://control-center-cp.{region}.oci.{secondLevelDomain}", "control-center-cp")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CapacityManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CapacityManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateOccAvailabilityCatalog Create availability catalog
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/CreateOccAvailabilityCatalog.go.html to see an example of how to use CreateOccAvailabilityCatalog API.
// A default retry strategy applies to this operation CreateOccAvailabilityCatalog()
func (client CapacityManagementClient) CreateOccAvailabilityCatalog(ctx context.Context, request CreateOccAvailabilityCatalogRequest) (response CreateOccAvailabilityCatalogResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccAvailabilityCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccAvailabilityCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccAvailabilityCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccAvailabilityCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccAvailabilityCatalogResponse")
	}
	return
}

// createOccAvailabilityCatalog implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) createOccAvailabilityCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occAvailabilityCatalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccAvailabilityCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalog/CreateOccAvailabilityCatalog"
		err = common.PostProcessServiceError(err, "CapacityManagement", "CreateOccAvailabilityCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOccCapacityRequest Create Capacity Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/CreateOccCapacityRequest.go.html to see an example of how to use CreateOccCapacityRequest API.
// A default retry strategy applies to this operation CreateOccCapacityRequest()
func (client CapacityManagementClient) CreateOccCapacityRequest(ctx context.Context, request CreateOccCapacityRequestRequest) (response CreateOccCapacityRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccCapacityRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccCapacityRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccCapacityRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccCapacityRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccCapacityRequestResponse")
	}
	return
}

// createOccCapacityRequest implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) createOccCapacityRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occCapacityRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccCapacityRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequest/CreateOccCapacityRequest"
		err = common.PostProcessServiceError(err, "CapacityManagement", "CreateOccCapacityRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOccCustomer Create customer.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/CreateOccCustomer.go.html to see an example of how to use CreateOccCustomer API.
// A default retry strategy applies to this operation CreateOccCustomer()
func (client CapacityManagementClient) CreateOccCustomer(ctx context.Context, request CreateOccCustomerRequest) (response CreateOccCustomerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccCustomer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccCustomerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccCustomerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccCustomerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccCustomerResponse")
	}
	return
}

// createOccCustomer implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) createOccCustomer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occCustomerGroups/{occCustomerGroupId}/occCustomers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccCustomerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomer/CreateOccCustomer"
		err = common.PostProcessServiceError(err, "CapacityManagement", "CreateOccCustomer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOccCustomerGroup Create customer group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/CreateOccCustomerGroup.go.html to see an example of how to use CreateOccCustomerGroup API.
// A default retry strategy applies to this operation CreateOccCustomerGroup()
func (client CapacityManagementClient) CreateOccCustomerGroup(ctx context.Context, request CreateOccCustomerGroupRequest) (response CreateOccCustomerGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccCustomerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccCustomerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccCustomerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccCustomerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccCustomerGroupResponse")
	}
	return
}

// createOccCustomerGroup implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) createOccCustomerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occCustomerGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccCustomerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomerGroup/CreateOccCustomerGroup"
		err = common.PostProcessServiceError(err, "CapacityManagement", "CreateOccCustomerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccAvailabilityCatalog Deletes the availability catalog resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/DeleteOccAvailabilityCatalog.go.html to see an example of how to use DeleteOccAvailabilityCatalog API.
// A default retry strategy applies to this operation DeleteOccAvailabilityCatalog()
func (client CapacityManagementClient) DeleteOccAvailabilityCatalog(ctx context.Context, request DeleteOccAvailabilityCatalogRequest) (response DeleteOccAvailabilityCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccAvailabilityCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccAvailabilityCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccAvailabilityCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccAvailabilityCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccAvailabilityCatalogResponse")
	}
	return
}

// deleteOccAvailabilityCatalog implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) deleteOccAvailabilityCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occAvailabilityCatalogs/{occAvailabilityCatalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccAvailabilityCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalog/DeleteOccAvailabilityCatalog"
		err = common.PostProcessServiceError(err, "CapacityManagement", "DeleteOccAvailabilityCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccCapacityRequest Deletes the capacity request resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/DeleteOccCapacityRequest.go.html to see an example of how to use DeleteOccCapacityRequest API.
// A default retry strategy applies to this operation DeleteOccCapacityRequest()
func (client CapacityManagementClient) DeleteOccCapacityRequest(ctx context.Context, request DeleteOccCapacityRequestRequest) (response DeleteOccCapacityRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccCapacityRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccCapacityRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccCapacityRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccCapacityRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccCapacityRequestResponse")
	}
	return
}

// deleteOccCapacityRequest implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) deleteOccCapacityRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occCapacityRequests/{occCapacityRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccCapacityRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequest/DeleteOccCapacityRequest"
		err = common.PostProcessServiceError(err, "CapacityManagement", "DeleteOccCapacityRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccCustomer Deletes the customer resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/DeleteOccCustomer.go.html to see an example of how to use DeleteOccCustomer API.
// A default retry strategy applies to this operation DeleteOccCustomer()
func (client CapacityManagementClient) DeleteOccCustomer(ctx context.Context, request DeleteOccCustomerRequest) (response DeleteOccCustomerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccCustomer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccCustomerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccCustomerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccCustomerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccCustomerResponse")
	}
	return
}

// deleteOccCustomer implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) deleteOccCustomer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occCustomerGroups/{occCustomerGroupId}/occCustomers/{occCustomerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccCustomerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomer/DeleteOccCustomer"
		err = common.PostProcessServiceError(err, "CapacityManagement", "DeleteOccCustomer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccCustomerGroup Deletes the customer group resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/DeleteOccCustomerGroup.go.html to see an example of how to use DeleteOccCustomerGroup API.
// A default retry strategy applies to this operation DeleteOccCustomerGroup()
func (client CapacityManagementClient) DeleteOccCustomerGroup(ctx context.Context, request DeleteOccCustomerGroupRequest) (response DeleteOccCustomerGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccCustomerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccCustomerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccCustomerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccCustomerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccCustomerGroupResponse")
	}
	return
}

// deleteOccCustomerGroup implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) deleteOccCustomerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occCustomerGroups/{occCustomerGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccCustomerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomerGroup/DeleteOccCustomerGroup"
		err = common.PostProcessServiceError(err, "CapacityManagement", "DeleteOccCustomerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccAvailabilityCatalog Get details about availability catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetOccAvailabilityCatalog.go.html to see an example of how to use GetOccAvailabilityCatalog API.
// A default retry strategy applies to this operation GetOccAvailabilityCatalog()
func (client CapacityManagementClient) GetOccAvailabilityCatalog(ctx context.Context, request GetOccAvailabilityCatalogRequest) (response GetOccAvailabilityCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOccAvailabilityCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccAvailabilityCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccAvailabilityCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccAvailabilityCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccAvailabilityCatalogResponse")
	}
	return
}

// getOccAvailabilityCatalog implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) getOccAvailabilityCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occAvailabilityCatalogs/{occAvailabilityCatalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccAvailabilityCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalog/GetOccAvailabilityCatalog"
		err = common.PostProcessServiceError(err, "CapacityManagement", "GetOccAvailabilityCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccAvailabilityCatalogContent Returns the binary contents of the availability catalog. Can be saved as a csv file.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetOccAvailabilityCatalogContent.go.html to see an example of how to use GetOccAvailabilityCatalogContent API.
// A default retry strategy applies to this operation GetOccAvailabilityCatalogContent()
func (client CapacityManagementClient) GetOccAvailabilityCatalogContent(ctx context.Context, request GetOccAvailabilityCatalogContentRequest) (response GetOccAvailabilityCatalogContentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getOccAvailabilityCatalogContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccAvailabilityCatalogContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccAvailabilityCatalogContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccAvailabilityCatalogContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccAvailabilityCatalogContentResponse")
	}
	return
}

// getOccAvailabilityCatalogContent implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) getOccAvailabilityCatalogContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occAvailabilityCatalogs/{occAvailabilityCatalogId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccAvailabilityCatalogContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalog/GetOccAvailabilityCatalogContent"
		err = common.PostProcessServiceError(err, "CapacityManagement", "GetOccAvailabilityCatalogContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccCapacityRequest Get details about the capacity request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetOccCapacityRequest.go.html to see an example of how to use GetOccCapacityRequest API.
// A default retry strategy applies to this operation GetOccCapacityRequest()
func (client CapacityManagementClient) GetOccCapacityRequest(ctx context.Context, request GetOccCapacityRequestRequest) (response GetOccCapacityRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOccCapacityRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccCapacityRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccCapacityRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccCapacityRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccCapacityRequestResponse")
	}
	return
}

// getOccCapacityRequest implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) getOccCapacityRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occCapacityRequests/{occCapacityRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccCapacityRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequest/GetOccCapacityRequest"
		err = common.PostProcessServiceError(err, "CapacityManagement", "GetOccCapacityRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccCustomerGroup Gets information about the specified customer group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/GetOccCustomerGroup.go.html to see an example of how to use GetOccCustomerGroup API.
// A default retry strategy applies to this operation GetOccCustomerGroup()
func (client CapacityManagementClient) GetOccCustomerGroup(ctx context.Context, request GetOccCustomerGroupRequest) (response GetOccCustomerGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOccCustomerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccCustomerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccCustomerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccCustomerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccCustomerGroupResponse")
	}
	return
}

// getOccCustomerGroup implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) getOccCustomerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occCustomerGroups/{occCustomerGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccCustomerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomerGroup/GetOccCustomerGroup"
		err = common.PostProcessServiceError(err, "CapacityManagement", "GetOccCustomerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalNamespaceOccOverviews Lists an overview of all resources in that namespace in a given time interval.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalNamespaceOccOverviews.go.html to see an example of how to use ListInternalNamespaceOccOverviews API.
// A default retry strategy applies to this operation ListInternalNamespaceOccOverviews()
func (client CapacityManagementClient) ListInternalNamespaceOccOverviews(ctx context.Context, request ListInternalNamespaceOccOverviewsRequest) (response ListInternalNamespaceOccOverviewsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalNamespaceOccOverviews, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalNamespaceOccOverviewsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalNamespaceOccOverviewsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalNamespaceOccOverviewsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalNamespaceOccOverviewsResponse")
	}
	return
}

// listInternalNamespaceOccOverviews implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listInternalNamespaceOccOverviews(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/namespace/{namespace}/occOverview", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalNamespaceOccOverviewsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccOverviewCollection/ListInternalNamespaceOccOverviews"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListInternalNamespaceOccOverviews", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalOccHandoverResourceBlockDetails List details about a given occHandoverResourceBlock.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccHandoverResourceBlockDetails.go.html to see an example of how to use ListInternalOccHandoverResourceBlockDetails API.
// A default retry strategy applies to this operation ListInternalOccHandoverResourceBlockDetails()
func (client CapacityManagementClient) ListInternalOccHandoverResourceBlockDetails(ctx context.Context, request ListInternalOccHandoverResourceBlockDetailsRequest) (response ListInternalOccHandoverResourceBlockDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalOccHandoverResourceBlockDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalOccHandoverResourceBlockDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalOccHandoverResourceBlockDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalOccHandoverResourceBlockDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalOccHandoverResourceBlockDetailsResponse")
	}
	return
}

// listInternalOccHandoverResourceBlockDetails implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listInternalOccHandoverResourceBlockDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occHandoverResourceBlockDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalOccHandoverResourceBlockDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccHandoverResourceBlockDetailCollection/ListInternalOccHandoverResourceBlockDetails"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListInternalOccHandoverResourceBlockDetails", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInternalOccHandoverResourceBlocks List Occ Handover Resource blocks.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccHandoverResourceBlocks.go.html to see an example of how to use ListInternalOccHandoverResourceBlocks API.
// A default retry strategy applies to this operation ListInternalOccHandoverResourceBlocks()
func (client CapacityManagementClient) ListInternalOccHandoverResourceBlocks(ctx context.Context, request ListInternalOccHandoverResourceBlocksRequest) (response ListInternalOccHandoverResourceBlocksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInternalOccHandoverResourceBlocks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInternalOccHandoverResourceBlocksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInternalOccHandoverResourceBlocksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInternalOccHandoverResourceBlocksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInternalOccHandoverResourceBlocksResponse")
	}
	return
}

// listInternalOccHandoverResourceBlocks implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listInternalOccHandoverResourceBlocks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occHandoverResourceBlocks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInternalOccHandoverResourceBlocksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccHandoverResourceBlockCollection/ListInternalOccHandoverResourceBlocks"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListInternalOccHandoverResourceBlocks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccAvailabilities Lists availabilities for a particular availability catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccAvailabilities.go.html to see an example of how to use ListOccAvailabilities API.
// A default retry strategy applies to this operation ListOccAvailabilities()
func (client CapacityManagementClient) ListOccAvailabilities(ctx context.Context, request ListOccAvailabilitiesRequest) (response ListOccAvailabilitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccAvailabilities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccAvailabilitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccAvailabilitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccAvailabilitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccAvailabilitiesResponse")
	}
	return
}

// listOccAvailabilities implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccAvailabilities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occAvailabilityCatalogs/{occAvailabilityCatalogId}/occAvailabilities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccAvailabilitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCollection/ListOccAvailabilities"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccAvailabilities", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccAvailabilityCatalogs Lists all availability catalogs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccAvailabilityCatalogs.go.html to see an example of how to use ListOccAvailabilityCatalogs API.
// A default retry strategy applies to this operation ListOccAvailabilityCatalogs()
func (client CapacityManagementClient) ListOccAvailabilityCatalogs(ctx context.Context, request ListOccAvailabilityCatalogsRequest) (response ListOccAvailabilityCatalogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccAvailabilityCatalogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccAvailabilityCatalogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccAvailabilityCatalogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccAvailabilityCatalogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccAvailabilityCatalogsResponse")
	}
	return
}

// listOccAvailabilityCatalogs implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccAvailabilityCatalogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occAvailabilityCatalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccAvailabilityCatalogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalogCollection/ListOccAvailabilityCatalogs"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccAvailabilityCatalogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccAvailabilityCatalogsInternal An internal api to list availability catalogs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccAvailabilityCatalogsInternal.go.html to see an example of how to use ListOccAvailabilityCatalogsInternal API.
// A default retry strategy applies to this operation ListOccAvailabilityCatalogsInternal()
func (client CapacityManagementClient) ListOccAvailabilityCatalogsInternal(ctx context.Context, request ListOccAvailabilityCatalogsInternalRequest) (response ListOccAvailabilityCatalogsInternalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccAvailabilityCatalogsInternal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccAvailabilityCatalogsInternalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccAvailabilityCatalogsInternalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccAvailabilityCatalogsInternalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccAvailabilityCatalogsInternalResponse")
	}
	return
}

// listOccAvailabilityCatalogsInternal implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccAvailabilityCatalogsInternal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occAvailabilityCatalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccAvailabilityCatalogsInternalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalogCollection/ListOccAvailabilityCatalogsInternal"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccAvailabilityCatalogsInternal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccCapacityRequests Lists all capacity requests.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccCapacityRequests.go.html to see an example of how to use ListOccCapacityRequests API.
// A default retry strategy applies to this operation ListOccCapacityRequests()
func (client CapacityManagementClient) ListOccCapacityRequests(ctx context.Context, request ListOccCapacityRequestsRequest) (response ListOccCapacityRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccCapacityRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccCapacityRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccCapacityRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccCapacityRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccCapacityRequestsResponse")
	}
	return
}

// listOccCapacityRequests implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccCapacityRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occCapacityRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccCapacityRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequestCollection/ListOccCapacityRequests"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccCapacityRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccCapacityRequestsInternal An internal api to list all capacity requests.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccCapacityRequestsInternal.go.html to see an example of how to use ListOccCapacityRequestsInternal API.
// A default retry strategy applies to this operation ListOccCapacityRequestsInternal()
func (client CapacityManagementClient) ListOccCapacityRequestsInternal(ctx context.Context, request ListOccCapacityRequestsInternalRequest) (response ListOccCapacityRequestsInternalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccCapacityRequestsInternal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccCapacityRequestsInternalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccCapacityRequestsInternalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccCapacityRequestsInternalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccCapacityRequestsInternalResponse")
	}
	return
}

// listOccCapacityRequestsInternal implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccCapacityRequestsInternal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/internal/occCapacityRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccCapacityRequestsInternalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequestCollection/ListOccCapacityRequestsInternal"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccCapacityRequestsInternal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccCustomerGroups Lists all the customer groups.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccCustomerGroups.go.html to see an example of how to use ListOccCustomerGroups API.
// A default retry strategy applies to this operation ListOccCustomerGroups()
func (client CapacityManagementClient) ListOccCustomerGroups(ctx context.Context, request ListOccCustomerGroupsRequest) (response ListOccCustomerGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccCustomerGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccCustomerGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccCustomerGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccCustomerGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccCustomerGroupsResponse")
	}
	return
}

// listOccCustomerGroups implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccCustomerGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occCustomerGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccCustomerGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomerGroupCollection/ListOccCustomerGroups"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccCustomerGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccHandoverResourceBlockDetails List details about a given occHandoverResourceBlock.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccHandoverResourceBlockDetails.go.html to see an example of how to use ListOccHandoverResourceBlockDetails API.
// A default retry strategy applies to this operation ListOccHandoverResourceBlockDetails()
func (client CapacityManagementClient) ListOccHandoverResourceBlockDetails(ctx context.Context, request ListOccHandoverResourceBlockDetailsRequest) (response ListOccHandoverResourceBlockDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccHandoverResourceBlockDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccHandoverResourceBlockDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccHandoverResourceBlockDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccHandoverResourceBlockDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccHandoverResourceBlockDetailsResponse")
	}
	return
}

// listOccHandoverResourceBlockDetails implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccHandoverResourceBlockDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occHandoverResourceBlockDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccHandoverResourceBlockDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccHandoverResourceBlockDetailCollection/ListOccHandoverResourceBlockDetails"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccHandoverResourceBlockDetails", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccHandoverResourceBlocks List Occ Handover Resource blocks.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccHandoverResourceBlocks.go.html to see an example of how to use ListOccHandoverResourceBlocks API.
// A default retry strategy applies to this operation ListOccHandoverResourceBlocks()
func (client CapacityManagementClient) ListOccHandoverResourceBlocks(ctx context.Context, request ListOccHandoverResourceBlocksRequest) (response ListOccHandoverResourceBlocksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccHandoverResourceBlocks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccHandoverResourceBlocksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccHandoverResourceBlocksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccHandoverResourceBlocksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccHandoverResourceBlocksResponse")
	}
	return
}

// listOccHandoverResourceBlocks implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccHandoverResourceBlocks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occHandoverResourceBlocks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccHandoverResourceBlocksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccHandoverResourceBlockCollection/ListOccHandoverResourceBlocks"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccHandoverResourceBlocks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccOverviews Lists an overview of all resources in that namespace in a given time interval.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccOverviews.go.html to see an example of how to use ListOccOverviews API.
// A default retry strategy applies to this operation ListOccOverviews()
func (client CapacityManagementClient) ListOccOverviews(ctx context.Context, request ListOccOverviewsRequest) (response ListOccOverviewsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccOverviews, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccOverviewsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccOverviewsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccOverviewsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccOverviewsResponse")
	}
	return
}

// listOccOverviews implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) listOccOverviews(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespace/{namespace}/occOverview", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccOverviewsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccOverviewCollection/ListOccOverviews"
		err = common.PostProcessServiceError(err, "CapacityManagement", "ListOccOverviews", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchInternalOccCapacityRequest Updates the OccCapacityRequest by evaluating a sequence of instructions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/PatchInternalOccCapacityRequest.go.html to see an example of how to use PatchInternalOccCapacityRequest API.
// A default retry strategy applies to this operation PatchInternalOccCapacityRequest()
func (client CapacityManagementClient) PatchInternalOccCapacityRequest(ctx context.Context, request PatchInternalOccCapacityRequestRequest) (response PatchInternalOccCapacityRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchInternalOccCapacityRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchInternalOccCapacityRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchInternalOccCapacityRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchInternalOccCapacityRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchInternalOccCapacityRequestResponse")
	}
	return
}

// patchInternalOccCapacityRequest implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) patchInternalOccCapacityRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/internal/occCapacityRequests/{occCapacityRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchInternalOccCapacityRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequest/PatchInternalOccCapacityRequest"
		err = common.PostProcessServiceError(err, "CapacityManagement", "PatchInternalOccCapacityRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchOccCapacityRequest Updates the OccCapacityRequest by evaluating a sequence of instructions.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/PatchOccCapacityRequest.go.html to see an example of how to use PatchOccCapacityRequest API.
// A default retry strategy applies to this operation PatchOccCapacityRequest()
func (client CapacityManagementClient) PatchOccCapacityRequest(ctx context.Context, request PatchOccCapacityRequestRequest) (response PatchOccCapacityRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchOccCapacityRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchOccCapacityRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchOccCapacityRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchOccCapacityRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchOccCapacityRequestResponse")
	}
	return
}

// patchOccCapacityRequest implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) patchOccCapacityRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/occCapacityRequests/{occCapacityRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchOccCapacityRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequest/PatchOccCapacityRequest"
		err = common.PostProcessServiceError(err, "CapacityManagement", "PatchOccCapacityRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishOccAvailabilityCatalog Publishes the version of availability catalog specified by the operator. This makes that catalog version visible to customers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/PublishOccAvailabilityCatalog.go.html to see an example of how to use PublishOccAvailabilityCatalog API.
// A default retry strategy applies to this operation PublishOccAvailabilityCatalog()
func (client CapacityManagementClient) PublishOccAvailabilityCatalog(ctx context.Context, request PublishOccAvailabilityCatalogRequest) (response PublishOccAvailabilityCatalogResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.publishOccAvailabilityCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishOccAvailabilityCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishOccAvailabilityCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishOccAvailabilityCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishOccAvailabilityCatalogResponse")
	}
	return
}

// publishOccAvailabilityCatalog implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) publishOccAvailabilityCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occAvailabilityCatalogs/{occAvailabilityCatalogId}/actions/publish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishOccAvailabilityCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalog/PublishOccAvailabilityCatalog"
		err = common.PostProcessServiceError(err, "CapacityManagement", "PublishOccAvailabilityCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateInternalOccCapacityRequest The internal api to update the capacity request. This api will be used by operators for updating the capacity request to either completed, resubmitted or rejected.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateInternalOccCapacityRequest.go.html to see an example of how to use UpdateInternalOccCapacityRequest API.
// A default retry strategy applies to this operation UpdateInternalOccCapacityRequest()
func (client CapacityManagementClient) UpdateInternalOccCapacityRequest(ctx context.Context, request UpdateInternalOccCapacityRequestRequest) (response UpdateInternalOccCapacityRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateInternalOccCapacityRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInternalOccCapacityRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInternalOccCapacityRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInternalOccCapacityRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInternalOccCapacityRequestResponse")
	}
	return
}

// updateInternalOccCapacityRequest implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) updateInternalOccCapacityRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/internal/occCapacityRequests/{occCapacityRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInternalOccCapacityRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequest/UpdateInternalOccCapacityRequest"
		err = common.PostProcessServiceError(err, "CapacityManagement", "UpdateInternalOccCapacityRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccAvailabilityCatalog The request to update the availability catalog. Currently only freeform tags can be updated via this api.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateOccAvailabilityCatalog.go.html to see an example of how to use UpdateOccAvailabilityCatalog API.
// A default retry strategy applies to this operation UpdateOccAvailabilityCatalog()
func (client CapacityManagementClient) UpdateOccAvailabilityCatalog(ctx context.Context, request UpdateOccAvailabilityCatalogRequest) (response UpdateOccAvailabilityCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccAvailabilityCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccAvailabilityCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccAvailabilityCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccAvailabilityCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccAvailabilityCatalogResponse")
	}
	return
}

// updateOccAvailabilityCatalog implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) updateOccAvailabilityCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occAvailabilityCatalogs/{occAvailabilityCatalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccAvailabilityCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccAvailabilityCatalog/UpdateOccAvailabilityCatalog"
		err = common.PostProcessServiceError(err, "CapacityManagement", "UpdateOccAvailabilityCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccCapacityRequest The request to update the capacity request. The user can perform actions like closing a partially completed request so that it doesn't go ahead for full completion.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateOccCapacityRequest.go.html to see an example of how to use UpdateOccCapacityRequest API.
// A default retry strategy applies to this operation UpdateOccCapacityRequest()
func (client CapacityManagementClient) UpdateOccCapacityRequest(ctx context.Context, request UpdateOccCapacityRequestRequest) (response UpdateOccCapacityRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccCapacityRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccCapacityRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccCapacityRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccCapacityRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccCapacityRequestResponse")
	}
	return
}

// updateOccCapacityRequest implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) updateOccCapacityRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occCapacityRequests/{occCapacityRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccCapacityRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCapacityRequest/UpdateOccCapacityRequest"
		err = common.PostProcessServiceError(err, "CapacityManagement", "UpdateOccCapacityRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccCustomer The request to update the customer.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateOccCustomer.go.html to see an example of how to use UpdateOccCustomer API.
// A default retry strategy applies to this operation UpdateOccCustomer()
func (client CapacityManagementClient) UpdateOccCustomer(ctx context.Context, request UpdateOccCustomerRequest) (response UpdateOccCustomerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccCustomer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccCustomerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccCustomerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccCustomerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccCustomerResponse")
	}
	return
}

// updateOccCustomer implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) updateOccCustomer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occCustomerGroups/{occCustomerGroupId}/occCustomers/{occCustomerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccCustomerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomer/UpdateOccCustomer"
		err = common.PostProcessServiceError(err, "CapacityManagement", "UpdateOccCustomer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccCustomerGroup The request to update the customer group.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/UpdateOccCustomerGroup.go.html to see an example of how to use UpdateOccCustomerGroup API.
// A default retry strategy applies to this operation UpdateOccCustomerGroup()
func (client CapacityManagementClient) UpdateOccCustomerGroup(ctx context.Context, request UpdateOccCustomerGroupRequest) (response UpdateOccCustomerGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccCustomerGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccCustomerGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccCustomerGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccCustomerGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccCustomerGroupResponse")
	}
	return
}

// updateOccCustomerGroup implements the OCIOperation interface (enables retrying operations)
func (client CapacityManagementClient) updateOccCustomerGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occCustomerGroups/{occCustomerGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccCustomerGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occcm/20231107/OccCustomerGroup/UpdateOccCustomerGroup"
		err = common.PostProcessServiceError(err, "CapacityManagement", "UpdateOccCustomerGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
