// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ByolClient a client for Byol
type ByolClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewByolClientWithConfigurationProvider Creates a new default Byol client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewByolClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ByolClient, err error) {
	if enabled := common.CheckForEnabledServices("ocvp"); !enabled {
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
	return newByolClientFromBaseClient(baseClient, provider)
}

// NewByolClientWithOboToken Creates a new default Byol client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewByolClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ByolClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newByolClientFromBaseClient(baseClient, configProvider)
}

func newByolClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ByolClient, err error) {
	// Byol service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Byol"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ByolClient{BaseClient: baseClient}
	client.BasePath = "20230701"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ByolClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ByolClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ByolClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeByolCompartment Moves an BYOL into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ChangeByolCompartment.go.html to see an example of how to use ChangeByolCompartment API.
// A default retry strategy applies to this operation ChangeByolCompartment()
func (client ByolClient) ChangeByolCompartment(ctx context.Context, request ChangeByolCompartmentRequest) (response ChangeByolCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeByolCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeByolCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeByolCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeByolCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeByolCompartmentResponse")
	}
	return
}

// changeByolCompartment implements the OCIOperation interface (enables retrying operations)
func (client ByolClient) changeByolCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/byols/{byolId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeByolCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Byol/ChangeByolCompartment"
		err = common.PostProcessServiceError(err, "Byol", "ChangeByolCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateByol Creates an Oracle Cloud VMware Solution Bring-You-Own-License (BYOL).
// Use the WorkRequest operations to track the
// creation of the BYOL.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/CreateByol.go.html to see an example of how to use CreateByol API.
// A default retry strategy applies to this operation CreateByol()
func (client ByolClient) CreateByol(ctx context.Context, request CreateByolRequest) (response CreateByolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createByol, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateByolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateByolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateByolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateByolResponse")
	}
	return
}

// createByol implements the OCIOperation interface (enables retrying operations)
func (client ByolClient) createByol(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/byols", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateByolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Byol/CreateByol"
		err = common.PostProcessServiceError(err, "Byol", "CreateByol", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteByol Deletes the specified BYOL.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/DeleteByol.go.html to see an example of how to use DeleteByol API.
// A default retry strategy applies to this operation DeleteByol()
func (client ByolClient) DeleteByol(ctx context.Context, request DeleteByolRequest) (response DeleteByolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteByol, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteByolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteByolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteByolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteByolResponse")
	}
	return
}

// deleteByol implements the OCIOperation interface (enables retrying operations)
func (client ByolClient) deleteByol(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/byols/{byolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteByolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/BYOL/DeleteByol"
		err = common.PostProcessServiceError(err, "Byol", "DeleteByol", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetByol Gets the specified BYOL's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/GetByol.go.html to see an example of how to use GetByol API.
// A default retry strategy applies to this operation GetByol()
func (client ByolClient) GetByol(ctx context.Context, request GetByolRequest) (response GetByolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getByol, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetByolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetByolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetByolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetByolResponse")
	}
	return
}

// getByol implements the OCIOperation interface (enables retrying operations)
func (client ByolClient) getByol(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/byols/{byolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetByolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Byol/GetByol"
		err = common.PostProcessServiceError(err, "Byol", "GetByol", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListByols Lists the BYOLs in the specified compartment. The list can be
// filtered by display name or availability domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListByols.go.html to see an example of how to use ListByols API.
// A default retry strategy applies to this operation ListByols()
func (client ByolClient) ListByols(ctx context.Context, request ListByolsRequest) (response ListByolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listByols, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListByolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListByolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListByolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListByolsResponse")
	}
	return
}

// listByols implements the OCIOperation interface (enables retrying operations)
func (client ByolClient) listByols(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/byols", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListByolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ByolSummary/ListByols"
		err = common.PostProcessServiceError(err, "Byol", "ListByols", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetrieveByolRealmAllocations Retrieve realm-wide allocation status for a specified BYOL.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/RetrieveByolRealmAllocations.go.html to see an example of how to use RetrieveByolRealmAllocations API.
// A default retry strategy applies to this operation RetrieveByolRealmAllocations()
func (client ByolClient) RetrieveByolRealmAllocations(ctx context.Context, request RetrieveByolRealmAllocationsRequest) (response RetrieveByolRealmAllocationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.retrieveByolRealmAllocations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetrieveByolRealmAllocationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetrieveByolRealmAllocationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetrieveByolRealmAllocationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetrieveByolRealmAllocationsResponse")
	}
	return
}

// retrieveByolRealmAllocations implements the OCIOperation interface (enables retrying operations)
func (client ByolClient) retrieveByolRealmAllocations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/byols/{byolId}/actions/retrieveByolRealmAllocations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetrieveByolRealmAllocationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Byol/RetrieveByolRealmAllocations"
		err = common.PostProcessServiceError(err, "Byol", "RetrieveByolRealmAllocations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateByol Updates the specified BYOL.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/UpdateByol.go.html to see an example of how to use UpdateByol API.
// A default retry strategy applies to this operation UpdateByol()
func (client ByolClient) UpdateByol(ctx context.Context, request UpdateByolRequest) (response UpdateByolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateByol, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateByolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateByolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateByolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateByolResponse")
	}
	return
}

// updateByol implements the OCIOperation interface (enables retrying operations)
func (client ByolClient) updateByol(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/byols/{byolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateByolResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/Byol/UpdateByol"
		err = common.PostProcessServiceError(err, "Byol", "UpdateByol", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
