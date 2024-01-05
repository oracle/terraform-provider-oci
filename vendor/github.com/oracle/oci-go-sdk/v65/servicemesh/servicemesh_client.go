// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ServiceMeshClient a client for ServiceMesh
type ServiceMeshClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewServiceMeshClientWithConfigurationProvider Creates a new default ServiceMesh client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewServiceMeshClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ServiceMeshClient, err error) {
	if enabled := common.CheckForEnabledServices("servicemesh"); !enabled {
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
	return newServiceMeshClientFromBaseClient(baseClient, provider)
}

// NewServiceMeshClientWithOboToken Creates a new default ServiceMesh client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewServiceMeshClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ServiceMeshClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newServiceMeshClientFromBaseClient(baseClient, configProvider)
}

func newServiceMeshClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ServiceMeshClient, err error) {
	// ServiceMesh service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ServiceMesh"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ServiceMeshClient{BaseClient: baseClient}
	client.BasePath = "20220615"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ServiceMeshClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("servicemesh", "https://servicemesh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ServiceMeshClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ServiceMeshClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelWorkRequest Cancels the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client ServiceMeshClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
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
func (client ServiceMeshClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAccessPolicyCompartment Moves an AccessPolicy resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ChangeAccessPolicyCompartment.go.html to see an example of how to use ChangeAccessPolicyCompartment API.
// A default retry strategy applies to this operation ChangeAccessPolicyCompartment()
func (client ServiceMeshClient) ChangeAccessPolicyCompartment(ctx context.Context, request ChangeAccessPolicyCompartmentRequest) (response ChangeAccessPolicyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAccessPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAccessPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAccessPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAccessPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAccessPolicyCompartmentResponse")
	}
	return
}

// changeAccessPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) changeAccessPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/accessPolicies/{accessPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAccessPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/AccessPolicy/ChangeAccessPolicyCompartment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ChangeAccessPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeIngressGatewayCompartment Moves a IngressGateway resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ChangeIngressGatewayCompartment.go.html to see an example of how to use ChangeIngressGatewayCompartment API.
// A default retry strategy applies to this operation ChangeIngressGatewayCompartment()
func (client ServiceMeshClient) ChangeIngressGatewayCompartment(ctx context.Context, request ChangeIngressGatewayCompartmentRequest) (response ChangeIngressGatewayCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeIngressGatewayCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeIngressGatewayCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeIngressGatewayCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeIngressGatewayCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeIngressGatewayCompartmentResponse")
	}
	return
}

// changeIngressGatewayCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) changeIngressGatewayCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ingressGateways/{ingressGatewayId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeIngressGatewayCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGateway/ChangeIngressGatewayCompartment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ChangeIngressGatewayCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeIngressGatewayRouteTableCompartment Moves a IngressGatewayRouteTable resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ChangeIngressGatewayRouteTableCompartment.go.html to see an example of how to use ChangeIngressGatewayRouteTableCompartment API.
// A default retry strategy applies to this operation ChangeIngressGatewayRouteTableCompartment()
func (client ServiceMeshClient) ChangeIngressGatewayRouteTableCompartment(ctx context.Context, request ChangeIngressGatewayRouteTableCompartmentRequest) (response ChangeIngressGatewayRouteTableCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeIngressGatewayRouteTableCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeIngressGatewayRouteTableCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeIngressGatewayRouteTableCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeIngressGatewayRouteTableCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeIngressGatewayRouteTableCompartmentResponse")
	}
	return
}

// changeIngressGatewayRouteTableCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) changeIngressGatewayRouteTableCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ingressGatewayRouteTables/{ingressGatewayRouteTableId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeIngressGatewayRouteTableCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGatewayRouteTable/ChangeIngressGatewayRouteTableCompartment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ChangeIngressGatewayRouteTableCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeMeshCompartment Moves a Mesh resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ChangeMeshCompartment.go.html to see an example of how to use ChangeMeshCompartment API.
// A default retry strategy applies to this operation ChangeMeshCompartment()
func (client ServiceMeshClient) ChangeMeshCompartment(ctx context.Context, request ChangeMeshCompartmentRequest) (response ChangeMeshCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeMeshCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeMeshCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeMeshCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeMeshCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeMeshCompartmentResponse")
	}
	return
}

// changeMeshCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) changeMeshCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/meshes/{meshId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeMeshCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/Mesh/ChangeMeshCompartment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ChangeMeshCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVirtualDeploymentCompartment Moves a VirtualDeployment resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ChangeVirtualDeploymentCompartment.go.html to see an example of how to use ChangeVirtualDeploymentCompartment API.
// A default retry strategy applies to this operation ChangeVirtualDeploymentCompartment()
func (client ServiceMeshClient) ChangeVirtualDeploymentCompartment(ctx context.Context, request ChangeVirtualDeploymentCompartmentRequest) (response ChangeVirtualDeploymentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVirtualDeploymentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVirtualDeploymentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVirtualDeploymentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVirtualDeploymentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVirtualDeploymentCompartmentResponse")
	}
	return
}

// changeVirtualDeploymentCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) changeVirtualDeploymentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualDeployments/{virtualDeploymentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVirtualDeploymentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualDeployment/ChangeVirtualDeploymentCompartment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ChangeVirtualDeploymentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVirtualServiceCompartment Moves a VirtualService resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ChangeVirtualServiceCompartment.go.html to see an example of how to use ChangeVirtualServiceCompartment API.
// A default retry strategy applies to this operation ChangeVirtualServiceCompartment()
func (client ServiceMeshClient) ChangeVirtualServiceCompartment(ctx context.Context, request ChangeVirtualServiceCompartmentRequest) (response ChangeVirtualServiceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVirtualServiceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVirtualServiceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVirtualServiceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVirtualServiceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVirtualServiceCompartmentResponse")
	}
	return
}

// changeVirtualServiceCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) changeVirtualServiceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualServices/{virtualServiceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVirtualServiceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualService/ChangeVirtualServiceCompartment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ChangeVirtualServiceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVirtualServiceRouteTableCompartment Moves a VirtualServiceRouteTable resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ChangeVirtualServiceRouteTableCompartment.go.html to see an example of how to use ChangeVirtualServiceRouteTableCompartment API.
// A default retry strategy applies to this operation ChangeVirtualServiceRouteTableCompartment()
func (client ServiceMeshClient) ChangeVirtualServiceRouteTableCompartment(ctx context.Context, request ChangeVirtualServiceRouteTableCompartmentRequest) (response ChangeVirtualServiceRouteTableCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVirtualServiceRouteTableCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVirtualServiceRouteTableCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVirtualServiceRouteTableCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVirtualServiceRouteTableCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVirtualServiceRouteTableCompartmentResponse")
	}
	return
}

// changeVirtualServiceRouteTableCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) changeVirtualServiceRouteTableCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualServiceRouteTables/{virtualServiceRouteTableId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVirtualServiceRouteTableCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualServiceRouteTable/ChangeVirtualServiceRouteTableCompartment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ChangeVirtualServiceRouteTableCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAccessPolicy Creates a new AccessPolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CreateAccessPolicy.go.html to see an example of how to use CreateAccessPolicy API.
// A default retry strategy applies to this operation CreateAccessPolicy()
func (client ServiceMeshClient) CreateAccessPolicy(ctx context.Context, request CreateAccessPolicyRequest) (response CreateAccessPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAccessPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAccessPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAccessPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAccessPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAccessPolicyResponse")
	}
	return
}

// createAccessPolicy implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) createAccessPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/accessPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAccessPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/AccessPolicy/CreateAccessPolicy"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CreateAccessPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateIngressGateway Creates a new IngressGateway.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CreateIngressGateway.go.html to see an example of how to use CreateIngressGateway API.
// A default retry strategy applies to this operation CreateIngressGateway()
func (client ServiceMeshClient) CreateIngressGateway(ctx context.Context, request CreateIngressGatewayRequest) (response CreateIngressGatewayResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createIngressGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIngressGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIngressGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIngressGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIngressGatewayResponse")
	}
	return
}

// createIngressGateway implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) createIngressGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ingressGateways", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIngressGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGateway/CreateIngressGateway"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CreateIngressGateway", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateIngressGatewayRouteTable Creates a new IngressGatewayRouteTable.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CreateIngressGatewayRouteTable.go.html to see an example of how to use CreateIngressGatewayRouteTable API.
// A default retry strategy applies to this operation CreateIngressGatewayRouteTable()
func (client ServiceMeshClient) CreateIngressGatewayRouteTable(ctx context.Context, request CreateIngressGatewayRouteTableRequest) (response CreateIngressGatewayRouteTableResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createIngressGatewayRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIngressGatewayRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIngressGatewayRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIngressGatewayRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIngressGatewayRouteTableResponse")
	}
	return
}

// createIngressGatewayRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) createIngressGatewayRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ingressGatewayRouteTables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIngressGatewayRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGatewayRouteTable/CreateIngressGatewayRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CreateIngressGatewayRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMesh Creates a new Mesh.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CreateMesh.go.html to see an example of how to use CreateMesh API.
// A default retry strategy applies to this operation CreateMesh()
func (client ServiceMeshClient) CreateMesh(ctx context.Context, request CreateMeshRequest) (response CreateMeshResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMesh, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMeshResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMeshResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMeshResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMeshResponse")
	}
	return
}

// createMesh implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) createMesh(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/meshes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMeshResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/Mesh/CreateMesh"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CreateMesh", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVirtualDeployment Creates a new VirtualDeployment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CreateVirtualDeployment.go.html to see an example of how to use CreateVirtualDeployment API.
// A default retry strategy applies to this operation CreateVirtualDeployment()
func (client ServiceMeshClient) CreateVirtualDeployment(ctx context.Context, request CreateVirtualDeploymentRequest) (response CreateVirtualDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVirtualDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVirtualDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVirtualDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVirtualDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVirtualDeploymentResponse")
	}
	return
}

// createVirtualDeployment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) createVirtualDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualDeployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVirtualDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualDeployment/CreateVirtualDeployment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CreateVirtualDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVirtualService Creates a new VirtualService.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CreateVirtualService.go.html to see an example of how to use CreateVirtualService API.
// A default retry strategy applies to this operation CreateVirtualService()
func (client ServiceMeshClient) CreateVirtualService(ctx context.Context, request CreateVirtualServiceRequest) (response CreateVirtualServiceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVirtualService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVirtualServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVirtualServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVirtualServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVirtualServiceResponse")
	}
	return
}

// createVirtualService implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) createVirtualService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVirtualServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualService/CreateVirtualService"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CreateVirtualService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVirtualServiceRouteTable Creates a new VirtualServiceRouteTable.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/CreateVirtualServiceRouteTable.go.html to see an example of how to use CreateVirtualServiceRouteTable API.
// A default retry strategy applies to this operation CreateVirtualServiceRouteTable()
func (client ServiceMeshClient) CreateVirtualServiceRouteTable(ctx context.Context, request CreateVirtualServiceRouteTableRequest) (response CreateVirtualServiceRouteTableResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVirtualServiceRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVirtualServiceRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVirtualServiceRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVirtualServiceRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVirtualServiceRouteTableResponse")
	}
	return
}

// createVirtualServiceRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) createVirtualServiceRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/virtualServiceRouteTables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVirtualServiceRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualServiceRouteTable/CreateVirtualServiceRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "CreateVirtualServiceRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAccessPolicy Deletes an AccessPolicy resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/DeleteAccessPolicy.go.html to see an example of how to use DeleteAccessPolicy API.
// A default retry strategy applies to this operation DeleteAccessPolicy()
func (client ServiceMeshClient) DeleteAccessPolicy(ctx context.Context, request DeleteAccessPolicyRequest) (response DeleteAccessPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAccessPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAccessPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAccessPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAccessPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAccessPolicyResponse")
	}
	return
}

// deleteAccessPolicy implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) deleteAccessPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/accessPolicies/{accessPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAccessPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/AccessPolicy/DeleteAccessPolicy"
		err = common.PostProcessServiceError(err, "ServiceMesh", "DeleteAccessPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteIngressGateway Deletes an IngressGateway resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/DeleteIngressGateway.go.html to see an example of how to use DeleteIngressGateway API.
// A default retry strategy applies to this operation DeleteIngressGateway()
func (client ServiceMeshClient) DeleteIngressGateway(ctx context.Context, request DeleteIngressGatewayRequest) (response DeleteIngressGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteIngressGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteIngressGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteIngressGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteIngressGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteIngressGatewayResponse")
	}
	return
}

// deleteIngressGateway implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) deleteIngressGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/ingressGateways/{ingressGatewayId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteIngressGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGateway/DeleteIngressGateway"
		err = common.PostProcessServiceError(err, "ServiceMesh", "DeleteIngressGateway", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteIngressGatewayRouteTable Deletes a IngressGatewayRouteTable resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/DeleteIngressGatewayRouteTable.go.html to see an example of how to use DeleteIngressGatewayRouteTable API.
// A default retry strategy applies to this operation DeleteIngressGatewayRouteTable()
func (client ServiceMeshClient) DeleteIngressGatewayRouteTable(ctx context.Context, request DeleteIngressGatewayRouteTableRequest) (response DeleteIngressGatewayRouteTableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteIngressGatewayRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteIngressGatewayRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteIngressGatewayRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteIngressGatewayRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteIngressGatewayRouteTableResponse")
	}
	return
}

// deleteIngressGatewayRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) deleteIngressGatewayRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/ingressGatewayRouteTables/{ingressGatewayRouteTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteIngressGatewayRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGatewayRouteTable/DeleteIngressGatewayRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "DeleteIngressGatewayRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMesh Deletes a Mesh resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/DeleteMesh.go.html to see an example of how to use DeleteMesh API.
// A default retry strategy applies to this operation DeleteMesh()
func (client ServiceMeshClient) DeleteMesh(ctx context.Context, request DeleteMeshRequest) (response DeleteMeshResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMesh, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMeshResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMeshResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMeshResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMeshResponse")
	}
	return
}

// deleteMesh implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) deleteMesh(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/meshes/{meshId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMeshResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/Mesh/DeleteMesh"
		err = common.PostProcessServiceError(err, "ServiceMesh", "DeleteMesh", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVirtualDeployment Deletes a VirtualDeployment resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/DeleteVirtualDeployment.go.html to see an example of how to use DeleteVirtualDeployment API.
// A default retry strategy applies to this operation DeleteVirtualDeployment()
func (client ServiceMeshClient) DeleteVirtualDeployment(ctx context.Context, request DeleteVirtualDeploymentRequest) (response DeleteVirtualDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVirtualDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVirtualDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVirtualDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVirtualDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVirtualDeploymentResponse")
	}
	return
}

// deleteVirtualDeployment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) deleteVirtualDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/virtualDeployments/{virtualDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVirtualDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualDeployment/DeleteVirtualDeployment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "DeleteVirtualDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVirtualService Deletes a VirtualService resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/DeleteVirtualService.go.html to see an example of how to use DeleteVirtualService API.
// A default retry strategy applies to this operation DeleteVirtualService()
func (client ServiceMeshClient) DeleteVirtualService(ctx context.Context, request DeleteVirtualServiceRequest) (response DeleteVirtualServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVirtualService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVirtualServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVirtualServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVirtualServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVirtualServiceResponse")
	}
	return
}

// deleteVirtualService implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) deleteVirtualService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/virtualServices/{virtualServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVirtualServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualService/DeleteVirtualService"
		err = common.PostProcessServiceError(err, "ServiceMesh", "DeleteVirtualService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVirtualServiceRouteTable Deletes a VirtualServiceRouteTable resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/DeleteVirtualServiceRouteTable.go.html to see an example of how to use DeleteVirtualServiceRouteTable API.
// A default retry strategy applies to this operation DeleteVirtualServiceRouteTable()
func (client ServiceMeshClient) DeleteVirtualServiceRouteTable(ctx context.Context, request DeleteVirtualServiceRouteTableRequest) (response DeleteVirtualServiceRouteTableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVirtualServiceRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVirtualServiceRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVirtualServiceRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVirtualServiceRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVirtualServiceRouteTableResponse")
	}
	return
}

// deleteVirtualServiceRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) deleteVirtualServiceRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/virtualServiceRouteTables/{virtualServiceRouteTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVirtualServiceRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualServiceRouteTable/DeleteVirtualServiceRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "DeleteVirtualServiceRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAccessPolicy Get an AccessPolicy by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetAccessPolicy.go.html to see an example of how to use GetAccessPolicy API.
// A default retry strategy applies to this operation GetAccessPolicy()
func (client ServiceMeshClient) GetAccessPolicy(ctx context.Context, request GetAccessPolicyRequest) (response GetAccessPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAccessPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAccessPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAccessPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAccessPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAccessPolicyResponse")
	}
	return
}

// getAccessPolicy implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getAccessPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/accessPolicies/{accessPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAccessPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/AccessPolicy/GetAccessPolicy"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetAccessPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIngressGateway Gets an IngressGateway by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetIngressGateway.go.html to see an example of how to use GetIngressGateway API.
// A default retry strategy applies to this operation GetIngressGateway()
func (client ServiceMeshClient) GetIngressGateway(ctx context.Context, request GetIngressGatewayRequest) (response GetIngressGatewayResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIngressGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIngressGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIngressGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIngressGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIngressGatewayResponse")
	}
	return
}

// getIngressGateway implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getIngressGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ingressGateways/{ingressGatewayId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIngressGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGateway/GetIngressGateway"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetIngressGateway", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIngressGatewayRouteTable Gets a IngressGatewayRouteTable by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetIngressGatewayRouteTable.go.html to see an example of how to use GetIngressGatewayRouteTable API.
// A default retry strategy applies to this operation GetIngressGatewayRouteTable()
func (client ServiceMeshClient) GetIngressGatewayRouteTable(ctx context.Context, request GetIngressGatewayRouteTableRequest) (response GetIngressGatewayRouteTableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIngressGatewayRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIngressGatewayRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIngressGatewayRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIngressGatewayRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIngressGatewayRouteTableResponse")
	}
	return
}

// getIngressGatewayRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getIngressGatewayRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ingressGatewayRouteTables/{ingressGatewayRouteTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIngressGatewayRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGatewayRouteTable/GetIngressGatewayRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetIngressGatewayRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMesh Gets a Mesh by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetMesh.go.html to see an example of how to use GetMesh API.
// A default retry strategy applies to this operation GetMesh()
func (client ServiceMeshClient) GetMesh(ctx context.Context, request GetMeshRequest) (response GetMeshResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMesh, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMeshResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMeshResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMeshResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMeshResponse")
	}
	return
}

// getMesh implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getMesh(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/meshes/{meshId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMeshResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/Mesh/GetMesh"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetMesh", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProxyDetails Returns the attributes of the Proxy such as proxy image version.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetProxyDetails.go.html to see an example of how to use GetProxyDetails API.
// A default retry strategy applies to this operation GetProxyDetails()
func (client ServiceMeshClient) GetProxyDetails(ctx context.Context, request GetProxyDetailsRequest) (response GetProxyDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProxyDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProxyDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProxyDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProxyDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProxyDetailsResponse")
	}
	return
}

// getProxyDetails implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getProxyDetails(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/proxyDetails", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProxyDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/ProxyDetails/GetProxyDetails"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetProxyDetails", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVirtualDeployment Gets a VirtualDeployment by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetVirtualDeployment.go.html to see an example of how to use GetVirtualDeployment API.
// A default retry strategy applies to this operation GetVirtualDeployment()
func (client ServiceMeshClient) GetVirtualDeployment(ctx context.Context, request GetVirtualDeploymentRequest) (response GetVirtualDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVirtualDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVirtualDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVirtualDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVirtualDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVirtualDeploymentResponse")
	}
	return
}

// getVirtualDeployment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getVirtualDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualDeployments/{virtualDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVirtualDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualDeployment/GetVirtualDeployment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetVirtualDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVirtualService Gets a VirtualService by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetVirtualService.go.html to see an example of how to use GetVirtualService API.
// A default retry strategy applies to this operation GetVirtualService()
func (client ServiceMeshClient) GetVirtualService(ctx context.Context, request GetVirtualServiceRequest) (response GetVirtualServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVirtualService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVirtualServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVirtualServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVirtualServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVirtualServiceResponse")
	}
	return
}

// getVirtualService implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getVirtualService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualServices/{virtualServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVirtualServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualService/GetVirtualService"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetVirtualService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVirtualServiceRouteTable Gets a VirtualServiceRouteTable by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetVirtualServiceRouteTable.go.html to see an example of how to use GetVirtualServiceRouteTable API.
// A default retry strategy applies to this operation GetVirtualServiceRouteTable()
func (client ServiceMeshClient) GetVirtualServiceRouteTable(ctx context.Context, request GetVirtualServiceRouteTableRequest) (response GetVirtualServiceRouteTableResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVirtualServiceRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVirtualServiceRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVirtualServiceRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVirtualServiceRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVirtualServiceRouteTableResponse")
	}
	return
}

// getVirtualServiceRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) getVirtualServiceRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualServiceRouteTables/{virtualServiceRouteTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVirtualServiceRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualServiceRouteTable/GetVirtualServiceRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetVirtualServiceRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client ServiceMeshClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client ServiceMeshClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ServiceMesh", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAccessPolicies Returns a list of AccessPolicy objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListAccessPolicies.go.html to see an example of how to use ListAccessPolicies API.
// A default retry strategy applies to this operation ListAccessPolicies()
func (client ServiceMeshClient) ListAccessPolicies(ctx context.Context, request ListAccessPoliciesRequest) (response ListAccessPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAccessPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAccessPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAccessPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAccessPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAccessPoliciesResponse")
	}
	return
}

// listAccessPolicies implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) listAccessPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/accessPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAccessPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/AccessPolicy/ListAccessPolicies"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListAccessPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIngressGatewayRouteTables Returns a list of IngressGatewayRouteTable objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListIngressGatewayRouteTables.go.html to see an example of how to use ListIngressGatewayRouteTables API.
// A default retry strategy applies to this operation ListIngressGatewayRouteTables()
func (client ServiceMeshClient) ListIngressGatewayRouteTables(ctx context.Context, request ListIngressGatewayRouteTablesRequest) (response ListIngressGatewayRouteTablesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIngressGatewayRouteTables, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIngressGatewayRouteTablesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIngressGatewayRouteTablesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIngressGatewayRouteTablesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIngressGatewayRouteTablesResponse")
	}
	return
}

// listIngressGatewayRouteTables implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) listIngressGatewayRouteTables(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ingressGatewayRouteTables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIngressGatewayRouteTablesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGatewayRouteTable/ListIngressGatewayRouteTables"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListIngressGatewayRouteTables", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIngressGateways Returns a list of IngressGateway objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListIngressGateways.go.html to see an example of how to use ListIngressGateways API.
// A default retry strategy applies to this operation ListIngressGateways()
func (client ServiceMeshClient) ListIngressGateways(ctx context.Context, request ListIngressGatewaysRequest) (response ListIngressGatewaysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIngressGateways, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIngressGatewaysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIngressGatewaysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIngressGatewaysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIngressGatewaysResponse")
	}
	return
}

// listIngressGateways implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) listIngressGateways(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ingressGateways", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIngressGatewaysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGateway/ListIngressGateways"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListIngressGateways", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMeshes Returns a list of Mesh objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListMeshes.go.html to see an example of how to use ListMeshes API.
// A default retry strategy applies to this operation ListMeshes()
func (client ServiceMeshClient) ListMeshes(ctx context.Context, request ListMeshesRequest) (response ListMeshesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMeshes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMeshesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMeshesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMeshesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMeshesResponse")
	}
	return
}

// listMeshes implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) listMeshes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/meshes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMeshesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/Mesh/ListMeshes"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListMeshes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVirtualDeployments Returns a list of VirtualDeployments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListVirtualDeployments.go.html to see an example of how to use ListVirtualDeployments API.
// A default retry strategy applies to this operation ListVirtualDeployments()
func (client ServiceMeshClient) ListVirtualDeployments(ctx context.Context, request ListVirtualDeploymentsRequest) (response ListVirtualDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVirtualDeployments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVirtualDeploymentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVirtualDeploymentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVirtualDeploymentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVirtualDeploymentsResponse")
	}
	return
}

// listVirtualDeployments implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) listVirtualDeployments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualDeployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVirtualDeploymentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualDeployment/ListVirtualDeployments"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListVirtualDeployments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVirtualServiceRouteTables Returns a list of VirtualServiceRouteTable objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListVirtualServiceRouteTables.go.html to see an example of how to use ListVirtualServiceRouteTables API.
// A default retry strategy applies to this operation ListVirtualServiceRouteTables()
func (client ServiceMeshClient) ListVirtualServiceRouteTables(ctx context.Context, request ListVirtualServiceRouteTablesRequest) (response ListVirtualServiceRouteTablesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVirtualServiceRouteTables, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVirtualServiceRouteTablesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVirtualServiceRouteTablesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVirtualServiceRouteTablesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVirtualServiceRouteTablesResponse")
	}
	return
}

// listVirtualServiceRouteTables implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) listVirtualServiceRouteTables(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualServiceRouteTables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVirtualServiceRouteTablesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualServiceRouteTable/ListVirtualServiceRouteTables"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListVirtualServiceRouteTables", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVirtualServices Returns a list of VirtualService objects.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListVirtualServices.go.html to see an example of how to use ListVirtualServices API.
// A default retry strategy applies to this operation ListVirtualServices()
func (client ServiceMeshClient) ListVirtualServices(ctx context.Context, request ListVirtualServicesRequest) (response ListVirtualServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVirtualServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVirtualServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVirtualServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVirtualServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVirtualServicesResponse")
	}
	return
}

// listVirtualServices implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) listVirtualServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/virtualServices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVirtualServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualService/ListVirtualServices"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListVirtualServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client ServiceMeshClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client ServiceMeshClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/WorkRequest/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client ServiceMeshClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client ServiceMeshClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/WorkRequest/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client ServiceMeshClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client ServiceMeshClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ServiceMesh", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAccessPolicy Updates the AccessPolicy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/UpdateAccessPolicy.go.html to see an example of how to use UpdateAccessPolicy API.
// A default retry strategy applies to this operation UpdateAccessPolicy()
func (client ServiceMeshClient) UpdateAccessPolicy(ctx context.Context, request UpdateAccessPolicyRequest) (response UpdateAccessPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateAccessPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAccessPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAccessPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAccessPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAccessPolicyResponse")
	}
	return
}

// updateAccessPolicy implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) updateAccessPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/accessPolicies/{accessPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAccessPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/AccessPolicy/UpdateAccessPolicy"
		err = common.PostProcessServiceError(err, "ServiceMesh", "UpdateAccessPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateIngressGateway Updates the IngressGateway.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/UpdateIngressGateway.go.html to see an example of how to use UpdateIngressGateway API.
// A default retry strategy applies to this operation UpdateIngressGateway()
func (client ServiceMeshClient) UpdateIngressGateway(ctx context.Context, request UpdateIngressGatewayRequest) (response UpdateIngressGatewayResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateIngressGateway, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateIngressGatewayResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateIngressGatewayResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIngressGatewayResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIngressGatewayResponse")
	}
	return
}

// updateIngressGateway implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) updateIngressGateway(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/ingressGateways/{ingressGatewayId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateIngressGatewayResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGateway/UpdateIngressGateway"
		err = common.PostProcessServiceError(err, "ServiceMesh", "UpdateIngressGateway", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateIngressGatewayRouteTable Updates the IngressGatewayRouteTable.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/UpdateIngressGatewayRouteTable.go.html to see an example of how to use UpdateIngressGatewayRouteTable API.
// A default retry strategy applies to this operation UpdateIngressGatewayRouteTable()
func (client ServiceMeshClient) UpdateIngressGatewayRouteTable(ctx context.Context, request UpdateIngressGatewayRouteTableRequest) (response UpdateIngressGatewayRouteTableResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateIngressGatewayRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateIngressGatewayRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateIngressGatewayRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIngressGatewayRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIngressGatewayRouteTableResponse")
	}
	return
}

// updateIngressGatewayRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) updateIngressGatewayRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/ingressGatewayRouteTables/{ingressGatewayRouteTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateIngressGatewayRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/IngressGatewayRouteTable/UpdateIngressGatewayRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "UpdateIngressGatewayRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMesh Updates the Mesh.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/UpdateMesh.go.html to see an example of how to use UpdateMesh API.
// A default retry strategy applies to this operation UpdateMesh()
func (client ServiceMeshClient) UpdateMesh(ctx context.Context, request UpdateMeshRequest) (response UpdateMeshResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateMesh, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMeshResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMeshResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMeshResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMeshResponse")
	}
	return
}

// updateMesh implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) updateMesh(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/meshes/{meshId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMeshResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/Mesh/UpdateMesh"
		err = common.PostProcessServiceError(err, "ServiceMesh", "UpdateMesh", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVirtualDeployment Updates the VirtualDeployment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/UpdateVirtualDeployment.go.html to see an example of how to use UpdateVirtualDeployment API.
// A default retry strategy applies to this operation UpdateVirtualDeployment()
func (client ServiceMeshClient) UpdateVirtualDeployment(ctx context.Context, request UpdateVirtualDeploymentRequest) (response UpdateVirtualDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateVirtualDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVirtualDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVirtualDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVirtualDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVirtualDeploymentResponse")
	}
	return
}

// updateVirtualDeployment implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) updateVirtualDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/virtualDeployments/{virtualDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVirtualDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualDeployment/UpdateVirtualDeployment"
		err = common.PostProcessServiceError(err, "ServiceMesh", "UpdateVirtualDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVirtualService Updates the VirtualService.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/UpdateVirtualService.go.html to see an example of how to use UpdateVirtualService API.
// A default retry strategy applies to this operation UpdateVirtualService()
func (client ServiceMeshClient) UpdateVirtualService(ctx context.Context, request UpdateVirtualServiceRequest) (response UpdateVirtualServiceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateVirtualService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVirtualServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVirtualServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVirtualServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVirtualServiceResponse")
	}
	return
}

// updateVirtualService implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) updateVirtualService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/virtualServices/{virtualServiceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVirtualServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualService/UpdateVirtualService"
		err = common.PostProcessServiceError(err, "ServiceMesh", "UpdateVirtualService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVirtualServiceRouteTable Updates the VirtualServiceRouteTable.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicemesh/UpdateVirtualServiceRouteTable.go.html to see an example of how to use UpdateVirtualServiceRouteTable API.
// A default retry strategy applies to this operation UpdateVirtualServiceRouteTable()
func (client ServiceMeshClient) UpdateVirtualServiceRouteTable(ctx context.Context, request UpdateVirtualServiceRouteTableRequest) (response UpdateVirtualServiceRouteTableResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateVirtualServiceRouteTable, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVirtualServiceRouteTableResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVirtualServiceRouteTableResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVirtualServiceRouteTableResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVirtualServiceRouteTableResponse")
	}
	return
}

// updateVirtualServiceRouteTable implements the OCIOperation interface (enables retrying operations)
func (client ServiceMeshClient) updateVirtualServiceRouteTable(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/virtualServiceRouteTables/{virtualServiceRouteTableId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVirtualServiceRouteTableResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-mesh/20220615/VirtualServiceRouteTable/UpdateVirtualServiceRouteTable"
		err = common.PostProcessServiceError(err, "ServiceMesh", "UpdateVirtualServiceRouteTable", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
