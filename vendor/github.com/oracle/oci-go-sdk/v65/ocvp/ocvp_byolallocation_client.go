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

// ByolAllocationClient a client for ByolAllocation
type ByolAllocationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewByolAllocationClientWithConfigurationProvider Creates a new default ByolAllocation client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewByolAllocationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ByolAllocationClient, err error) {
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
	return newByolAllocationClientFromBaseClient(baseClient, provider)
}

// NewByolAllocationClientWithOboToken Creates a new default ByolAllocation client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewByolAllocationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ByolAllocationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newByolAllocationClientFromBaseClient(baseClient, configProvider)
}

func newByolAllocationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ByolAllocationClient, err error) {
	// ByolAllocation service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ByolAllocation"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ByolAllocationClient{BaseClient: baseClient}
	client.BasePath = "20230701"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ByolAllocationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ocvp", "https://ocvps.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ByolAllocationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ByolAllocationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeByolAllocationCompartment Moves an BYOL Allocation into a different compartment within the same tenancy. For information
// about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ChangeByolAllocationCompartment.go.html to see an example of how to use ChangeByolAllocationCompartment API.
// A default retry strategy applies to this operation ChangeByolAllocationCompartment()
func (client ByolAllocationClient) ChangeByolAllocationCompartment(ctx context.Context, request ChangeByolAllocationCompartmentRequest) (response ChangeByolAllocationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeByolAllocationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeByolAllocationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeByolAllocationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeByolAllocationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeByolAllocationCompartmentResponse")
	}
	return
}

// changeByolAllocationCompartment implements the OCIOperation interface (enables retrying operations)
func (client ByolAllocationClient) changeByolAllocationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/byolAllocations/{byolAllocationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeByolAllocationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ByolAllocation/ChangeByolAllocationCompartment"
		err = common.PostProcessServiceError(err, "ByolAllocation", "ChangeByolAllocationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateByolAllocation Creates an Allocation on an specific Bring-You-Own-License (BYOL).
// Use the WorkRequest operations to track the
// creation of the BYOL.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/CreateByolAllocation.go.html to see an example of how to use CreateByolAllocation API.
// A default retry strategy applies to this operation CreateByolAllocation()
func (client ByolAllocationClient) CreateByolAllocation(ctx context.Context, request CreateByolAllocationRequest) (response CreateByolAllocationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createByolAllocation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateByolAllocationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateByolAllocationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateByolAllocationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateByolAllocationResponse")
	}
	return
}

// createByolAllocation implements the OCIOperation interface (enables retrying operations)
func (client ByolAllocationClient) createByolAllocation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/byolAllocations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateByolAllocationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ByolAllocation/CreateByolAllocation"
		err = common.PostProcessServiceError(err, "ByolAllocation", "CreateByolAllocation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteByolAllocation Deletes the specified BYOL Allocation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/DeleteByolAllocation.go.html to see an example of how to use DeleteByolAllocation API.
// A default retry strategy applies to this operation DeleteByolAllocation()
func (client ByolAllocationClient) DeleteByolAllocation(ctx context.Context, request DeleteByolAllocationRequest) (response DeleteByolAllocationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteByolAllocation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteByolAllocationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteByolAllocationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteByolAllocationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteByolAllocationResponse")
	}
	return
}

// deleteByolAllocation implements the OCIOperation interface (enables retrying operations)
func (client ByolAllocationClient) deleteByolAllocation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/byolAllocations/{byolAllocationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteByolAllocationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ByolAllocation/DeleteByolAllocation"
		err = common.PostProcessServiceError(err, "ByolAllocation", "DeleteByolAllocation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetByolAllocation Gets the specified BYOL Allocation's information.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/GetByolAllocation.go.html to see an example of how to use GetByolAllocation API.
// A default retry strategy applies to this operation GetByolAllocation()
func (client ByolAllocationClient) GetByolAllocation(ctx context.Context, request GetByolAllocationRequest) (response GetByolAllocationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getByolAllocation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetByolAllocationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetByolAllocationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetByolAllocationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetByolAllocationResponse")
	}
	return
}

// getByolAllocation implements the OCIOperation interface (enables retrying operations)
func (client ByolAllocationClient) getByolAllocation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/byolAllocations/{byolAllocationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetByolAllocationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ByolAllocation/GetByolAllocation"
		err = common.PostProcessServiceError(err, "ByolAllocation", "GetByolAllocation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListByolAllocations Lists the BYOL Allocations in the specified compartment. The list can be
// filtered by display name or availability domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/ListByolAllocations.go.html to see an example of how to use ListByolAllocations API.
// A default retry strategy applies to this operation ListByolAllocations()
func (client ByolAllocationClient) ListByolAllocations(ctx context.Context, request ListByolAllocationsRequest) (response ListByolAllocationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listByolAllocations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListByolAllocationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListByolAllocationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListByolAllocationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListByolAllocationsResponse")
	}
	return
}

// listByolAllocations implements the OCIOperation interface (enables retrying operations)
func (client ByolAllocationClient) listByolAllocations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/byolAllocations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListByolAllocationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ByolAllocationSummary/ListByolAllocations"
		err = common.PostProcessServiceError(err, "ByolAllocation", "ListByolAllocations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateByolAllocation Updates the specified BYOL Allocation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ocvp/UpdateByolAllocation.go.html to see an example of how to use UpdateByolAllocation API.
// A default retry strategy applies to this operation UpdateByolAllocation()
func (client ByolAllocationClient) UpdateByolAllocation(ctx context.Context, request UpdateByolAllocationRequest) (response UpdateByolAllocationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateByolAllocation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateByolAllocationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateByolAllocationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateByolAllocationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateByolAllocationResponse")
	}
	return
}

// updateByolAllocation implements the OCIOperation interface (enables retrying operations)
func (client ByolAllocationClient) updateByolAllocation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/byolAllocations/{byolAllocationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateByolAllocationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/vmware/20230701/ByolAllocation/UpdateByolAllocation"
		err = common.PostProcessServiceError(err, "ByolAllocation", "UpdateByolAllocation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
