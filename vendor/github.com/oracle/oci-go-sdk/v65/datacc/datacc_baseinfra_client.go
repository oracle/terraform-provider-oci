// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// BaseinfraClient a client for Baseinfra
type BaseinfraClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBaseinfraClientWithConfigurationProvider Creates a new default Baseinfra client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBaseinfraClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BaseinfraClient, err error) {
	if enabled := common.CheckForEnabledServices("datacc"); !enabled {
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
	return newBaseinfraClientFromBaseClient(baseClient, provider)
}

// NewBaseinfraClientWithOboToken Creates a new default Baseinfra client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewBaseinfraClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BaseinfraClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newBaseinfraClientFromBaseClient(baseClient, configProvider)
}

func newBaseinfraClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BaseinfraClient, err error) {
	// Baseinfra service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Baseinfra"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = BaseinfraClient{BaseClient: baseClient}
	client.BasePath = "20251101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BaseinfraClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("datacc", "https://datacc.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BaseinfraClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *BaseinfraClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateInfrastructure Activate an Database Infrastructure using the specified details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ActivateInfrastructure.go.html to see an example of how to use ActivateInfrastructure API.
// A default retry strategy applies to this operation ActivateInfrastructure()
func (client BaseinfraClient) ActivateInfrastructure(ctx context.Context, request ActivateInfrastructureRequest) (response ActivateInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.activateInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateInfrastructureResponse")
	}
	return
}

// activateInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) activateInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/infrastructures/{infrastructureId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ActivateInfrastructure")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ActivateInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeInfrastructureCompartment Change the Database Infrastructure Cloud@Customer Compartment to the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ChangeInfrastructureCompartment.go.html to see an example of how to use ChangeInfrastructureCompartment API.
// A default retry strategy applies to this operation ChangeInfrastructureCompartment()
func (client BaseinfraClient) ChangeInfrastructureCompartment(ctx context.Context, request ChangeInfrastructureCompartmentRequest) (response ChangeInfrastructureCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeInfrastructureCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeInfrastructureCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeInfrastructureCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeInfrastructureCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeInfrastructureCompartmentResponse")
	}
	return
}

// changeInfrastructureCompartment implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) changeInfrastructureCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/infrastructures/{infrastructureId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeInfrastructureCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ChangeInfrastructureCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ChangeInfrastructureCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVmClusterNetworkCompartment Change the VM cluster network Compartment to the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ChangeVmClusterNetworkCompartment.go.html to see an example of how to use ChangeVmClusterNetworkCompartment API.
// A default retry strategy applies to this operation ChangeVmClusterNetworkCompartment()
func (client BaseinfraClient) ChangeVmClusterNetworkCompartment(ctx context.Context, request ChangeVmClusterNetworkCompartmentRequest) (response ChangeVmClusterNetworkCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVmClusterNetworkCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVmClusterNetworkCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVmClusterNetworkCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVmClusterNetworkCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVmClusterNetworkCompartmentResponse")
	}
	return
}

// changeVmClusterNetworkCompartment implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) changeVmClusterNetworkCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusterNetworks/{vmClusterNetworkId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVmClusterNetworkCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ChangeVmClusterNetworkCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ChangeVmClusterNetworkCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeVmInstanceCompartment Change the VM instance on Database Infrastructure Compartment to the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ChangeVmInstanceCompartment.go.html to see an example of how to use ChangeVmInstanceCompartment API.
// A default retry strategy applies to this operation ChangeVmInstanceCompartment()
func (client BaseinfraClient) ChangeVmInstanceCompartment(ctx context.Context, request ChangeVmInstanceCompartmentRequest) (response ChangeVmInstanceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeVmInstanceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeVmInstanceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeVmInstanceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeVmInstanceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeVmInstanceCompartmentResponse")
	}
	return
}

// changeVmInstanceCompartment implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) changeVmInstanceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmInstances/{vmInstanceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeVmInstanceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ChangeVmInstanceCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ChangeVmInstanceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateInfrastructure Create a Database Infrastructure using the specified details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/CreateInfrastructure.go.html to see an example of how to use CreateInfrastructure API.
// A default retry strategy applies to this operation CreateInfrastructure()
func (client BaseinfraClient) CreateInfrastructure(ctx context.Context, request CreateInfrastructureRequest) (response CreateInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateInfrastructureResponse")
	}
	return
}

// createInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) createInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/infrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "CreateInfrastructure")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "CreateInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVmClusterNetwork Create an VM cluster on Database Infrastructure network using the specified details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/CreateVmClusterNetwork.go.html to see an example of how to use CreateVmClusterNetwork API.
// A default retry strategy applies to this operation CreateVmClusterNetwork()
func (client BaseinfraClient) CreateVmClusterNetwork(ctx context.Context, request CreateVmClusterNetworkRequest) (response CreateVmClusterNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVmClusterNetworkResponse")
	}
	return
}

// createVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) createVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusterNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "CreateVmClusterNetwork")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "CreateVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateVmInstance Create an VM instance on Database Infrastructure using the specified details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/CreateVmInstance.go.html to see an example of how to use CreateVmInstance API.
// A default retry strategy applies to this operation CreateVmInstance()
func (client BaseinfraClient) CreateVmInstance(ctx context.Context, request CreateVmInstanceRequest) (response CreateVmInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateVmInstanceResponse")
	}
	return
}

// createVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) createVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "CreateVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "CreateVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteInfrastructure Delete the Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). It must be called once the resource usage is complete.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/DeleteInfrastructure.go.html to see an example of how to use DeleteInfrastructure API.
// A default retry strategy applies to this operation DeleteInfrastructure()
func (client BaseinfraClient) DeleteInfrastructure(ctx context.Context, request DeleteInfrastructureRequest) (response DeleteInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteInfrastructureResponse")
	}
	return
}

// deleteInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) deleteInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/infrastructures/{infrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "DeleteInfrastructure")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "DeleteInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVmClusterNetwork Delete the VM cluster network on Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). It must be called once the resource usage is complete.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/DeleteVmClusterNetwork.go.html to see an example of how to use DeleteVmClusterNetwork API.
// A default retry strategy applies to this operation DeleteVmClusterNetwork()
func (client BaseinfraClient) DeleteVmClusterNetwork(ctx context.Context, request DeleteVmClusterNetworkRequest) (response DeleteVmClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVmClusterNetworkResponse")
	}
	return
}

// deleteVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) deleteVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "DeleteVmClusterNetwork")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "DeleteVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteVmInstance Delete the VM instance on Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). It must be called once the resource usage is complete.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/DeleteVmInstance.go.html to see an example of how to use DeleteVmInstance API.
// A default retry strategy applies to this operation DeleteVmInstance()
func (client BaseinfraClient) DeleteVmInstance(ctx context.Context, request DeleteVmInstanceRequest) (response DeleteVmInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteVmInstanceResponse")
	}
	return
}

// deleteVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) deleteVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/vmInstances/{vmInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "DeleteVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "DeleteVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadInfrastructureConfigFile Downloads the configuration file for the specified Database Infrastructure Cloud@Customer.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/DownloadInfrastructureConfigFile.go.html to see an example of how to use DownloadInfrastructureConfigFile API.
// A default retry strategy applies to this operation DownloadInfrastructureConfigFile()
func (client BaseinfraClient) DownloadInfrastructureConfigFile(ctx context.Context, request DownloadInfrastructureConfigFileRequest) (response DownloadInfrastructureConfigFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadInfrastructureConfigFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadInfrastructureConfigFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadInfrastructureConfigFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadInfrastructureConfigFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadInfrastructureConfigFileResponse")
	}
	return
}

// downloadInfrastructureConfigFile implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) downloadInfrastructureConfigFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/infrastructures/{infrastructureId}/actions/downloadConfigFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadInfrastructureConfigFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "DownloadInfrastructureConfigFile")
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "DownloadInfrastructureConfigFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadInfrastructureValidationScript Downloads the validation script for the specified Database Infrastructure Cloud@Customer.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/DownloadInfrastructureValidationScript.go.html to see an example of how to use DownloadInfrastructureValidationScript API.
// A default retry strategy applies to this operation DownloadInfrastructureValidationScript()
func (client BaseinfraClient) DownloadInfrastructureValidationScript(ctx context.Context, request DownloadInfrastructureValidationScriptRequest) (response DownloadInfrastructureValidationScriptResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.downloadInfrastructureValidationScript, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadInfrastructureValidationScriptResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadInfrastructureValidationScriptResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadInfrastructureValidationScriptResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadInfrastructureValidationScriptResponse")
	}
	return
}

// downloadInfrastructureValidationScript implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) downloadInfrastructureValidationScript(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/infrastructures/{infrastructureId}/actions/downloadValidationScript", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadInfrastructureValidationScriptResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "DownloadInfrastructureValidationScript")
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "DownloadInfrastructureValidationScript", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateRecommendedVmClusterNetwork Generates an VM cluster network configuration using the specified details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GenerateRecommendedVmClusterNetwork.go.html to see an example of how to use GenerateRecommendedVmClusterNetwork API.
// A default retry strategy applies to this operation GenerateRecommendedVmClusterNetwork()
func (client BaseinfraClient) GenerateRecommendedVmClusterNetwork(ctx context.Context, request GenerateRecommendedVmClusterNetworkRequest) (response GenerateRecommendedVmClusterNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateRecommendedVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateRecommendedVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateRecommendedVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateRecommendedVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateRecommendedVmClusterNetworkResponse")
	}
	return
}

// generateRecommendedVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) generateRecommendedVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusterNetworks/actions/generateRecommendedVmClusterNetwork", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateRecommendedVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GenerateRecommendedVmClusterNetwork")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GenerateRecommendedVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInfrastructure Obtain the Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GetInfrastructure.go.html to see an example of how to use GetInfrastructure API.
// A default retry strategy applies to this operation GetInfrastructure()
func (client BaseinfraClient) GetInfrastructure(ctx context.Context, request GetInfrastructureRequest) (response GetInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInfrastructureResponse")
	}
	return
}

// getInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) getInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/infrastructures/{infrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GetInfrastructure")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GetInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetInfrastructureScaleOption Get the available scale options for the infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GetInfrastructureScaleOption.go.html to see an example of how to use GetInfrastructureScaleOption API.
// A default retry strategy applies to this operation GetInfrastructureScaleOption()
func (client BaseinfraClient) GetInfrastructureScaleOption(ctx context.Context, request GetInfrastructureScaleOptionRequest) (response GetInfrastructureScaleOptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInfrastructureScaleOption, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInfrastructureScaleOptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInfrastructureScaleOptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInfrastructureScaleOptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInfrastructureScaleOptionResponse")
	}
	return
}

// getInfrastructureScaleOption implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) getInfrastructureScaleOption(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/infrastructures/{infrastructureId}/scaleOptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInfrastructureScaleOptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GetInfrastructureScaleOption")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GetInfrastructureScaleOption", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaintenanceExecution Gets information about the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GetMaintenanceExecution.go.html to see an example of how to use GetMaintenanceExecution API.
// A default retry strategy applies to this operation GetMaintenanceExecution()
func (client BaseinfraClient) GetMaintenanceExecution(ctx context.Context, request GetMaintenanceExecutionRequest) (response GetMaintenanceExecutionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaintenanceExecution, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaintenanceExecutionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaintenanceExecutionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaintenanceExecutionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaintenanceExecutionResponse")
	}
	return
}

// getMaintenanceExecution implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) getMaintenanceExecution(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceExecutions/{maintenanceExecutionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceExecutionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GetMaintenanceExecution")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GetMaintenanceExecution", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaintenanceRun Gets information about the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GetMaintenanceRun.go.html to see an example of how to use GetMaintenanceRun API.
// A default retry strategy applies to this operation GetMaintenanceRun()
func (client BaseinfraClient) GetMaintenanceRun(ctx context.Context, request GetMaintenanceRunRequest) (response GetMaintenanceRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaintenanceRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaintenanceRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaintenanceRunResponse")
	}
	return
}

// getMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) getMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns/{maintenanceRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GetMaintenanceRun")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GetMaintenanceRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmClusterNetwork Obtain the VM cluster network on Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GetVmClusterNetwork.go.html to see an example of how to use GetVmClusterNetwork API.
// A default retry strategy applies to this operation GetVmClusterNetwork()
func (client BaseinfraClient) GetVmClusterNetwork(ctx context.Context, request GetVmClusterNetworkRequest) (response GetVmClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmClusterNetworkResponse")
	}
	return
}

// getVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) getVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GetVmClusterNetwork")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GetVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetVmInstance Obtain the VM instance on Database Infrastructure that has the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GetVmInstance.go.html to see an example of how to use GetVmInstance API.
// A default retry strategy applies to this operation GetVmInstance()
func (client BaseinfraClient) GetVmInstance(ctx context.Context, request GetVmInstanceRequest) (response GetVmInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetVmInstanceResponse")
	}
	return
}

// getVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) getVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmInstances/{vmInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GetVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GetVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Get the status of the work request having the specified work request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client BaseinfraClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client BaseinfraClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "GetWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInfrastructures Obtain a list of Database Infrastructures.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListInfrastructures.go.html to see an example of how to use ListInfrastructures API.
// A default retry strategy applies to this operation ListInfrastructures()
func (client BaseinfraClient) ListInfrastructures(ctx context.Context, request ListInfrastructuresRequest) (response ListInfrastructuresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInfrastructures, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInfrastructuresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInfrastructuresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInfrastructuresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInfrastructuresResponse")
	}
	return
}

// listInfrastructures implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) listInfrastructures(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/infrastructures", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInfrastructuresResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListInfrastructures")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListInfrastructures", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaintenanceExecutions Gets a list of the maintenance executions in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListMaintenanceExecutions.go.html to see an example of how to use ListMaintenanceExecutions API.
// A default retry strategy applies to this operation ListMaintenanceExecutions()
func (client BaseinfraClient) ListMaintenanceExecutions(ctx context.Context, request ListMaintenanceExecutionsRequest) (response ListMaintenanceExecutionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaintenanceExecutions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaintenanceExecutionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaintenanceExecutionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaintenanceExecutionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaintenanceExecutionsResponse")
	}
	return
}

// listMaintenanceExecutions implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) listMaintenanceExecutions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceExecutions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceExecutionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListMaintenanceExecutions")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListMaintenanceExecutions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaintenanceRuns Gets a list of the maintenance runs in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListMaintenanceRuns.go.html to see an example of how to use ListMaintenanceRuns API.
// A default retry strategy applies to this operation ListMaintenanceRuns()
func (client BaseinfraClient) ListMaintenanceRuns(ctx context.Context, request ListMaintenanceRunsRequest) (response ListMaintenanceRunsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaintenanceRuns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaintenanceRunsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaintenanceRunsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaintenanceRunsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaintenanceRunsResponse")
	}
	return
}

// listMaintenanceRuns implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) listMaintenanceRuns(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceRuns", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceRunsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListMaintenanceRuns")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListMaintenanceRuns", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmClusterNetworks Obtain a list of VM cluster networks on Database Infrastructure.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListVmClusterNetworks.go.html to see an example of how to use ListVmClusterNetworks API.
// A default retry strategy applies to this operation ListVmClusterNetworks()
func (client BaseinfraClient) ListVmClusterNetworks(ctx context.Context, request ListVmClusterNetworksRequest) (response ListVmClusterNetworksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmClusterNetworks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmClusterNetworksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmClusterNetworksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmClusterNetworksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmClusterNetworksResponse")
	}
	return
}

// listVmClusterNetworks implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) listVmClusterNetworks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmClusterNetworks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmClusterNetworksResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListVmClusterNetworks")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListVmClusterNetworks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListVmInstances Obtain a list of VM instances.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListVmInstances.go.html to see an example of how to use ListVmInstances API.
// A default retry strategy applies to this operation ListVmInstances()
func (client BaseinfraClient) ListVmInstances(ctx context.Context, request ListVmInstancesRequest) (response ListVmInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listVmInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListVmInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListVmInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListVmInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListVmInstancesResponse")
	}
	return
}

// listVmInstances implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) listVmInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/vmInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListVmInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListVmInstances")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListVmInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors View the list of errors for the work request that have the specified work request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client BaseinfraClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client BaseinfraClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListWorkRequestErrors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs View the log entries for the work request that have the specified work request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client BaseinfraClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client BaseinfraClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListWorkRequestLogs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests View a list of work request details in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client BaseinfraClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client BaseinfraClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ListWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MigrateVmInstance Migrate the VM instance on Database Infrastructure that has the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/MigrateVmInstance.go.html to see an example of how to use MigrateVmInstance API.
// A default retry strategy applies to this operation MigrateVmInstance()
func (client BaseinfraClient) MigrateVmInstance(ctx context.Context, request MigrateVmInstanceRequest) (response MigrateVmInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.migrateVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MigrateVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MigrateVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MigrateVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MigrateVmInstanceResponse")
	}
	return
}

// migrateVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) migrateVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmInstances/{vmInstanceId}/actions/migrate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MigrateVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "MigrateVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "MigrateVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RescheduleMaintenanceRun Reschedules the specified maintenance run by providing a new `timeScheduled` value.
// Only maintenance runs in eligible states may be rescheduled.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/RescheduleMaintenanceRun.go.html to see an example of how to use RescheduleMaintenanceRun API.
// A default retry strategy applies to this operation RescheduleMaintenanceRun()
func (client BaseinfraClient) RescheduleMaintenanceRun(ctx context.Context, request RescheduleMaintenanceRunRequest) (response RescheduleMaintenanceRunResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rescheduleMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RescheduleMaintenanceRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RescheduleMaintenanceRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RescheduleMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RescheduleMaintenanceRunResponse")
	}
	return
}

// rescheduleMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) rescheduleMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maintenanceRuns/{maintenanceRunId}/actions/reschedule", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RescheduleMaintenanceRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "RescheduleMaintenanceRun")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "RescheduleMaintenanceRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestartVmInstance Restart the VM instance on Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/RestartVmInstance.go.html to see an example of how to use RestartVmInstance API.
// A default retry strategy applies to this operation RestartVmInstance()
func (client BaseinfraClient) RestartVmInstance(ctx context.Context, request RestartVmInstanceRequest) (response RestartVmInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restartVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartVmInstanceResponse")
	}
	return
}

// restartVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) restartVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmInstances/{vmInstanceId}/actions/restart", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestartVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "RestartVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "RestartVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScaleInfrastructureStorage Scale storage for the specified infrastructure using the specified details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ScaleInfrastructureStorage.go.html to see an example of how to use ScaleInfrastructureStorage API.
// A default retry strategy applies to this operation ScaleInfrastructureStorage()
func (client BaseinfraClient) ScaleInfrastructureStorage(ctx context.Context, request ScaleInfrastructureStorageRequest) (response ScaleInfrastructureStorageResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scaleInfrastructureStorage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScaleInfrastructureStorageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScaleInfrastructureStorageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScaleInfrastructureStorageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScaleInfrastructureStorageResponse")
	}
	return
}

// scaleInfrastructureStorage implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) scaleInfrastructureStorage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/infrastructures/{infrastructureId}/actions/scaleStorage", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScaleInfrastructureStorageResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ScaleInfrastructureStorage")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ScaleInfrastructureStorage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScaleVmInstance Scale the VM instance on Database Infrastructure that has the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ScaleVmInstance.go.html to see an example of how to use ScaleVmInstance API.
// A default retry strategy applies to this operation ScaleVmInstance()
func (client BaseinfraClient) ScaleVmInstance(ctx context.Context, request ScaleVmInstanceRequest) (response ScaleVmInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scaleVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScaleVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScaleVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScaleVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScaleVmInstanceResponse")
	}
	return
}

// scaleVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) scaleVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmInstances/{vmInstanceId}/actions/scale", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScaleVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ScaleVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ScaleVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartVmInstance Start the VM instance on Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/StartVmInstance.go.html to see an example of how to use StartVmInstance API.
// A default retry strategy applies to this operation StartVmInstance()
func (client BaseinfraClient) StartVmInstance(ctx context.Context, request StartVmInstanceRequest) (response StartVmInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartVmInstanceResponse")
	}
	return
}

// startVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) startVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmInstances/{vmInstanceId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "StartVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "StartVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopVmInstance Stop the VM instance on Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/StopVmInstance.go.html to see an example of how to use StopVmInstance API.
// A default retry strategy applies to this operation StopVmInstance()
func (client BaseinfraClient) StopVmInstance(ctx context.Context, request StopVmInstanceRequest) (response StopVmInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopVmInstanceResponse")
	}
	return
}

// stopVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) stopVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmInstances/{vmInstanceId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "StopVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "StopVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateInfrastructure Update the Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/UpdateInfrastructure.go.html to see an example of how to use UpdateInfrastructure API.
// A default retry strategy applies to this operation UpdateInfrastructure()
func (client BaseinfraClient) UpdateInfrastructure(ctx context.Context, request UpdateInfrastructureRequest) (response UpdateInfrastructureResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateInfrastructureResponse")
	}
	return
}

// updateInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) updateInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/infrastructures/{infrastructureId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "UpdateInfrastructure")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "UpdateInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMaintenanceRun Update maintenance run attributes.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/UpdateMaintenanceRun.go.html to see an example of how to use UpdateMaintenanceRun API.
// A default retry strategy applies to this operation UpdateMaintenanceRun()
func (client BaseinfraClient) UpdateMaintenanceRun(ctx context.Context, request UpdateMaintenanceRunRequest) (response UpdateMaintenanceRunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMaintenanceRun, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMaintenanceRunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMaintenanceRunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMaintenanceRunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMaintenanceRunResponse")
	}
	return
}

// updateMaintenanceRun implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) updateMaintenanceRun(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maintenanceRuns/{maintenanceRunId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMaintenanceRunResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "UpdateMaintenanceRun")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "UpdateMaintenanceRun", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVmClusterNetwork Update the VM cluster network on Database Infrastructure that has the specified
// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/UpdateVmClusterNetwork.go.html to see an example of how to use UpdateVmClusterNetwork API.
// A default retry strategy applies to this operation UpdateVmClusterNetwork()
func (client BaseinfraClient) UpdateVmClusterNetwork(ctx context.Context, request UpdateVmClusterNetworkRequest) (response UpdateVmClusterNetworkResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVmClusterNetworkResponse")
	}
	return
}

// updateVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) updateVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/vmClusterNetworks/{vmClusterNetworkId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "UpdateVmClusterNetwork")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "UpdateVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateVmInstance Update the VM instance on Database Infrastructure that has the specified OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/UpdateVmInstance.go.html to see an example of how to use UpdateVmInstance API.
// A default retry strategy applies to this operation UpdateVmInstance()
func (client BaseinfraClient) UpdateVmInstance(ctx context.Context, request UpdateVmInstanceRequest) (response UpdateVmInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateVmInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateVmInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateVmInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateVmInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateVmInstanceResponse")
	}
	return
}

// updateVmInstance implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) updateVmInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/vmInstances/{vmInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateVmInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "UpdateVmInstance")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "UpdateVmInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateInfrastructure Validate an Database Infrastructure using the specified details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ValidateInfrastructure.go.html to see an example of how to use ValidateInfrastructure API.
// A default retry strategy applies to this operation ValidateInfrastructure()
func (client BaseinfraClient) ValidateInfrastructure(ctx context.Context, request ValidateInfrastructureRequest) (response ValidateInfrastructureResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateInfrastructure, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateInfrastructureResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateInfrastructureResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateInfrastructureResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateInfrastructureResponse")
	}
	return
}

// validateInfrastructure implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) validateInfrastructure(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/infrastructures/{infrastructureId}/actions/validate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateInfrastructureResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ValidateInfrastructure")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ValidateInfrastructure", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateVmClusterNetwork Validates an VM cluster network on Database Infrastructure.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ValidateVmClusterNetwork.go.html to see an example of how to use ValidateVmClusterNetwork API.
// A default retry strategy applies to this operation ValidateVmClusterNetwork()
func (client BaseinfraClient) ValidateVmClusterNetwork(ctx context.Context, request ValidateVmClusterNetworkRequest) (response ValidateVmClusterNetworkResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.validateVmClusterNetwork, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateVmClusterNetworkResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateVmClusterNetworkResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateVmClusterNetworkResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateVmClusterNetworkResponse")
	}
	return
}

// validateVmClusterNetwork implements the OCIOperation interface (enables retrying operations)
func (client BaseinfraClient) validateVmClusterNetwork(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/vmClusterNetworks/{vmClusterNetworkId}/actions/validateVmClusterNetwork", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ValidateVmClusterNetworkResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "baseinfra", "ValidateVmClusterNetwork")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Baseinfra", "ValidateVmClusterNetwork", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
