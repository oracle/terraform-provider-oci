// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OciCacheConfigSetClient a client for OciCacheConfigSet
type OciCacheConfigSetClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOciCacheConfigSetClientWithConfigurationProvider Creates a new default OciCacheConfigSet client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOciCacheConfigSetClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OciCacheConfigSetClient, err error) {
	if enabled := common.CheckForEnabledServices("redis"); !enabled {
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
	return newOciCacheConfigSetClientFromBaseClient(baseClient, provider)
}

// NewOciCacheConfigSetClientWithOboToken Creates a new default OciCacheConfigSet client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOciCacheConfigSetClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OciCacheConfigSetClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOciCacheConfigSetClientFromBaseClient(baseClient, configProvider)
}

func newOciCacheConfigSetClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OciCacheConfigSetClient, err error) {
	// OciCacheConfigSet service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OciCacheConfigSet"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OciCacheConfigSetClient{BaseClient: baseClient}
	client.BasePath = "20220315"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OciCacheConfigSetClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("redis", "https://redis.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OciCacheConfigSetClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OciCacheConfigSetClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOciCacheConfigSetCompartment Moves an OCI Cache Config Set into a different compartment within the same tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ChangeOciCacheConfigSetCompartment.go.html to see an example of how to use ChangeOciCacheConfigSetCompartment API.
// A default retry strategy applies to this operation ChangeOciCacheConfigSetCompartment()
func (client OciCacheConfigSetClient) ChangeOciCacheConfigSetCompartment(ctx context.Context, request ChangeOciCacheConfigSetCompartmentRequest) (response ChangeOciCacheConfigSetCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOciCacheConfigSetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOciCacheConfigSetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOciCacheConfigSetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOciCacheConfigSetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOciCacheConfigSetCompartmentResponse")
	}
	return
}

// changeOciCacheConfigSetCompartment implements the OCIOperation interface (enables retrying operations)
func (client OciCacheConfigSetClient) changeOciCacheConfigSetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheConfigSets/{ociCacheConfigSetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOciCacheConfigSetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheConfigSet/ChangeOciCacheConfigSetCompartment"
		err = common.PostProcessServiceError(err, "OciCacheConfigSet", "ChangeOciCacheConfigSetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOciCacheConfigSet Create a new OCI Cache Config Set for the given OCI cache engine version.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/CreateOciCacheConfigSet.go.html to see an example of how to use CreateOciCacheConfigSet API.
// A default retry strategy applies to this operation CreateOciCacheConfigSet()
func (client OciCacheConfigSetClient) CreateOciCacheConfigSet(ctx context.Context, request CreateOciCacheConfigSetRequest) (response CreateOciCacheConfigSetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOciCacheConfigSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOciCacheConfigSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOciCacheConfigSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOciCacheConfigSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOciCacheConfigSetResponse")
	}
	return
}

// createOciCacheConfigSet implements the OCIOperation interface (enables retrying operations)
func (client OciCacheConfigSetClient) createOciCacheConfigSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheConfigSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOciCacheConfigSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheConfigSet/CreateOciCacheConfigSet"
		err = common.PostProcessServiceError(err, "OciCacheConfigSet", "CreateOciCacheConfigSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOciCacheConfigSet Deletes the specified OCI Cache Config Set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/DeleteOciCacheConfigSet.go.html to see an example of how to use DeleteOciCacheConfigSet API.
// A default retry strategy applies to this operation DeleteOciCacheConfigSet()
func (client OciCacheConfigSetClient) DeleteOciCacheConfigSet(ctx context.Context, request DeleteOciCacheConfigSetRequest) (response DeleteOciCacheConfigSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOciCacheConfigSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOciCacheConfigSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOciCacheConfigSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOciCacheConfigSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOciCacheConfigSetResponse")
	}
	return
}

// deleteOciCacheConfigSet implements the OCIOperation interface (enables retrying operations)
func (client OciCacheConfigSetClient) deleteOciCacheConfigSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/ociCacheConfigSets/{ociCacheConfigSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOciCacheConfigSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheConfigSet/DeleteOciCacheConfigSet"
		err = common.PostProcessServiceError(err, "OciCacheConfigSet", "DeleteOciCacheConfigSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOciCacheConfigSet Retrieves the specified OCI Cache Config Set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/GetOciCacheConfigSet.go.html to see an example of how to use GetOciCacheConfigSet API.
// A default retry strategy applies to this operation GetOciCacheConfigSet()
func (client OciCacheConfigSetClient) GetOciCacheConfigSet(ctx context.Context, request GetOciCacheConfigSetRequest) (response GetOciCacheConfigSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOciCacheConfigSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOciCacheConfigSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOciCacheConfigSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOciCacheConfigSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOciCacheConfigSetResponse")
	}
	return
}

// getOciCacheConfigSet implements the OCIOperation interface (enables retrying operations)
func (client OciCacheConfigSetClient) getOciCacheConfigSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheConfigSets/{ociCacheConfigSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOciCacheConfigSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheConfigSet/GetOciCacheConfigSet"
		err = common.PostProcessServiceError(err, "OciCacheConfigSet", "GetOciCacheConfigSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssociatedOciCacheClusters Gets a list of associated OCI Cache clusters for an OCI Cache Config Set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListAssociatedOciCacheClusters.go.html to see an example of how to use ListAssociatedOciCacheClusters API.
// A default retry strategy applies to this operation ListAssociatedOciCacheClusters()
func (client OciCacheConfigSetClient) ListAssociatedOciCacheClusters(ctx context.Context, request ListAssociatedOciCacheClustersRequest) (response ListAssociatedOciCacheClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssociatedOciCacheClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssociatedOciCacheClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssociatedOciCacheClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssociatedOciCacheClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssociatedOciCacheClustersResponse")
	}
	return
}

// listAssociatedOciCacheClusters implements the OCIOperation interface (enables retrying operations)
func (client OciCacheConfigSetClient) listAssociatedOciCacheClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheConfigSets/{ociCacheConfigSetId}/actions/listAssociatedOciCacheClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssociatedOciCacheClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/AssociatedOciCacheClusterSummary/ListAssociatedOciCacheClusters"
		err = common.PostProcessServiceError(err, "OciCacheConfigSet", "ListAssociatedOciCacheClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOciCacheConfigSets Lists the OCI Cache Config Sets in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheConfigSets.go.html to see an example of how to use ListOciCacheConfigSets API.
// A default retry strategy applies to this operation ListOciCacheConfigSets()
func (client OciCacheConfigSetClient) ListOciCacheConfigSets(ctx context.Context, request ListOciCacheConfigSetsRequest) (response ListOciCacheConfigSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOciCacheConfigSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOciCacheConfigSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOciCacheConfigSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOciCacheConfigSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOciCacheConfigSetsResponse")
	}
	return
}

// listOciCacheConfigSets implements the OCIOperation interface (enables retrying operations)
func (client OciCacheConfigSetClient) listOciCacheConfigSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheConfigSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOciCacheConfigSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheConfigSetSummary/ListOciCacheConfigSets"
		err = common.PostProcessServiceError(err, "OciCacheConfigSet", "ListOciCacheConfigSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOciCacheConfigSet Updates the specified OCI Cache Config Set.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/UpdateOciCacheConfigSet.go.html to see an example of how to use UpdateOciCacheConfigSet API.
// A default retry strategy applies to this operation UpdateOciCacheConfigSet()
func (client OciCacheConfigSetClient) UpdateOciCacheConfigSet(ctx context.Context, request UpdateOciCacheConfigSetRequest) (response UpdateOciCacheConfigSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOciCacheConfigSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOciCacheConfigSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOciCacheConfigSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOciCacheConfigSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOciCacheConfigSetResponse")
	}
	return
}

// updateOciCacheConfigSet implements the OCIOperation interface (enables retrying operations)
func (client OciCacheConfigSetClient) updateOciCacheConfigSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/ociCacheConfigSets/{ociCacheConfigSetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOciCacheConfigSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheConfigSet/UpdateOciCacheConfigSet"
		err = common.PostProcessServiceError(err, "OciCacheConfigSet", "UpdateOciCacheConfigSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
