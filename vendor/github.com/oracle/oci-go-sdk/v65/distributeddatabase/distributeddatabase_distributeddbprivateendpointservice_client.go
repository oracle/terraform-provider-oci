// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DistributedDbPrivateEndpointServiceClient a client for DistributedDbPrivateEndpointService
type DistributedDbPrivateEndpointServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDistributedDbPrivateEndpointServiceClientWithConfigurationProvider Creates a new default DistributedDbPrivateEndpointService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDistributedDbPrivateEndpointServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DistributedDbPrivateEndpointServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("distributeddatabase"); !enabled {
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
	return newDistributedDbPrivateEndpointServiceClientFromBaseClient(baseClient, provider)
}

// NewDistributedDbPrivateEndpointServiceClientWithOboToken Creates a new default DistributedDbPrivateEndpointService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDistributedDbPrivateEndpointServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DistributedDbPrivateEndpointServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDistributedDbPrivateEndpointServiceClientFromBaseClient(baseClient, configProvider)
}

func newDistributedDbPrivateEndpointServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DistributedDbPrivateEndpointServiceClient, err error) {
	// DistributedDbPrivateEndpointService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DistributedDbPrivateEndpointService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DistributedDbPrivateEndpointServiceClient{BaseClient: baseClient}
	client.BasePath = "20250101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DistributedDbPrivateEndpointServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("distributeddatabase", "https://globaldb.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DistributedDbPrivateEndpointServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DistributedDbPrivateEndpointServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeDistributedDatabasePrivateEndpointCompartment Moves the DistributedDatabasePrivateEndpoint to the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ChangeDistributedDatabasePrivateEndpointCompartment.go.html to see an example of how to use ChangeDistributedDatabasePrivateEndpointCompartment API.
// A default retry strategy applies to this operation ChangeDistributedDatabasePrivateEndpointCompartment()
func (client DistributedDbPrivateEndpointServiceClient) ChangeDistributedDatabasePrivateEndpointCompartment(ctx context.Context, request ChangeDistributedDatabasePrivateEndpointCompartmentRequest) (response ChangeDistributedDatabasePrivateEndpointCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDistributedDatabasePrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDistributedDatabasePrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDistributedDatabasePrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDistributedDatabasePrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDistributedDatabasePrivateEndpointCompartmentResponse")
	}
	return
}

// changeDistributedDatabasePrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbPrivateEndpointServiceClient) changeDistributedDatabasePrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabasePrivateEndpoints/{distributedDatabasePrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDistributedDatabasePrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabasePrivateEndpoint/ChangeDistributedDatabasePrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "DistributedDbPrivateEndpointService", "ChangeDistributedDatabasePrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDistributedDatabasePrivateEndpoint Creates a DistributedDatabasePrivateEndpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/CreateDistributedDatabasePrivateEndpoint.go.html to see an example of how to use CreateDistributedDatabasePrivateEndpoint API.
// A default retry strategy applies to this operation CreateDistributedDatabasePrivateEndpoint()
func (client DistributedDbPrivateEndpointServiceClient) CreateDistributedDatabasePrivateEndpoint(ctx context.Context, request CreateDistributedDatabasePrivateEndpointRequest) (response CreateDistributedDatabasePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDistributedDatabasePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDistributedDatabasePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDistributedDatabasePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDistributedDatabasePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDistributedDatabasePrivateEndpointResponse")
	}
	return
}

// createDistributedDatabasePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbPrivateEndpointServiceClient) createDistributedDatabasePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabasePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDistributedDatabasePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DistributedDbPrivateEndpointService", "CreateDistributedDatabasePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDistributedDatabasePrivateEndpoint Deletes the given DistributedDatabasePrivateEndpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/DeleteDistributedDatabasePrivateEndpoint.go.html to see an example of how to use DeleteDistributedDatabasePrivateEndpoint API.
// A default retry strategy applies to this operation DeleteDistributedDatabasePrivateEndpoint()
func (client DistributedDbPrivateEndpointServiceClient) DeleteDistributedDatabasePrivateEndpoint(ctx context.Context, request DeleteDistributedDatabasePrivateEndpointRequest) (response DeleteDistributedDatabasePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteDistributedDatabasePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDistributedDatabasePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDistributedDatabasePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDistributedDatabasePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDistributedDatabasePrivateEndpointResponse")
	}
	return
}

// deleteDistributedDatabasePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbPrivateEndpointServiceClient) deleteDistributedDatabasePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/distributedDatabasePrivateEndpoints/{distributedDatabasePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDistributedDatabasePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabasePrivateEndpoint/DeleteDistributedDatabasePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DistributedDbPrivateEndpointService", "DeleteDistributedDatabasePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDistributedDatabasePrivateEndpoint Get the DistributedDatabasePrivateEndpoint resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/GetDistributedDatabasePrivateEndpoint.go.html to see an example of how to use GetDistributedDatabasePrivateEndpoint API.
// A default retry strategy applies to this operation GetDistributedDatabasePrivateEndpoint()
func (client DistributedDbPrivateEndpointServiceClient) GetDistributedDatabasePrivateEndpoint(ctx context.Context, request GetDistributedDatabasePrivateEndpointRequest) (response GetDistributedDatabasePrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDistributedDatabasePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDistributedDatabasePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDistributedDatabasePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDistributedDatabasePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDistributedDatabasePrivateEndpointResponse")
	}
	return
}

// getDistributedDatabasePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbPrivateEndpointServiceClient) getDistributedDatabasePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedDatabasePrivateEndpoints/{distributedDatabasePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDistributedDatabasePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabasePrivateEndpoint/GetDistributedDatabasePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DistributedDbPrivateEndpointService", "GetDistributedDatabasePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDistributedDatabasePrivateEndpoints List of DistributedDatabasePrivateEndpoints.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ListDistributedDatabasePrivateEndpoints.go.html to see an example of how to use ListDistributedDatabasePrivateEndpoints API.
// A default retry strategy applies to this operation ListDistributedDatabasePrivateEndpoints()
func (client DistributedDbPrivateEndpointServiceClient) ListDistributedDatabasePrivateEndpoints(ctx context.Context, request ListDistributedDatabasePrivateEndpointsRequest) (response ListDistributedDatabasePrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDistributedDatabasePrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDistributedDatabasePrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDistributedDatabasePrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDistributedDatabasePrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDistributedDatabasePrivateEndpointsResponse")
	}
	return
}

// listDistributedDatabasePrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbPrivateEndpointServiceClient) listDistributedDatabasePrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/distributedDatabasePrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDistributedDatabasePrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabasePrivateEndpointCollection/ListDistributedDatabasePrivateEndpoints"
		err = common.PostProcessServiceError(err, "DistributedDbPrivateEndpointService", "ListDistributedDatabasePrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReinstateProxyInstance Reinstates the proxy instance associated with the DistributedDatabasePrivateEndpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ReinstateProxyInstance.go.html to see an example of how to use ReinstateProxyInstance API.
// A default retry strategy applies to this operation ReinstateProxyInstance()
func (client DistributedDbPrivateEndpointServiceClient) ReinstateProxyInstance(ctx context.Context, request ReinstateProxyInstanceRequest) (response ReinstateProxyInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.reinstateProxyInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReinstateProxyInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReinstateProxyInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReinstateProxyInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReinstateProxyInstanceResponse")
	}
	return
}

// reinstateProxyInstance implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbPrivateEndpointServiceClient) reinstateProxyInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/distributedDatabasePrivateEndpoints/{distributedDatabasePrivateEndpointId}/actions/reinstateProxyInstance", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ReinstateProxyInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabasePrivateEndpoint/ReinstateProxyInstance"
		err = common.PostProcessServiceError(err, "DistributedDbPrivateEndpointService", "ReinstateProxyInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDistributedDatabasePrivateEndpoint Updates the configuration of DistributedDatabasePrivateEndpoint.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/UpdateDistributedDatabasePrivateEndpoint.go.html to see an example of how to use UpdateDistributedDatabasePrivateEndpoint API.
// A default retry strategy applies to this operation UpdateDistributedDatabasePrivateEndpoint()
func (client DistributedDbPrivateEndpointServiceClient) UpdateDistributedDatabasePrivateEndpoint(ctx context.Context, request UpdateDistributedDatabasePrivateEndpointRequest) (response UpdateDistributedDatabasePrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateDistributedDatabasePrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDistributedDatabasePrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDistributedDatabasePrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDistributedDatabasePrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDistributedDatabasePrivateEndpointResponse")
	}
	return
}

// updateDistributedDatabasePrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DistributedDbPrivateEndpointServiceClient) updateDistributedDatabasePrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/distributedDatabasePrivateEndpoints/{distributedDatabasePrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDistributedDatabasePrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/globally-distributed-database/20250101/DistributedDatabasePrivateEndpoint/UpdateDistributedDatabasePrivateEndpoint"
		err = common.PostProcessServiceError(err, "DistributedDbPrivateEndpointService", "UpdateDistributedDatabasePrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
