// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// IotClient a client for Iot
type IotClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIotClientWithConfigurationProvider Creates a new default Iot client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewIotClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client IotClient, err error) {
	if enabled := common.CheckForEnabledServices("iot"); !enabled {
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
	return newIotClientFromBaseClient(baseClient, provider)
}

// NewIotClientWithOboToken Creates a new default Iot client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewIotClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client IotClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newIotClientFromBaseClient(baseClient, configProvider)
}

func newIotClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client IotClient, err error) {
	// Iot service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Iot"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = IotClient{BaseClient: baseClient}
	client.BasePath = "20250531"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *IotClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("iot", "https://iot.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IotClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *IotClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeIotDomainCompartment Moves an IoT domain to a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ChangeIotDomainCompartment.go.html to see an example of how to use ChangeIotDomainCompartment API.
// A default retry strategy applies to this operation ChangeIotDomainCompartment()
func (client IotClient) ChangeIotDomainCompartment(ctx context.Context, request ChangeIotDomainCompartmentRequest) (response ChangeIotDomainCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeIotDomainCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeIotDomainCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeIotDomainCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeIotDomainCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeIotDomainCompartmentResponse")
	}
	return
}

// changeIotDomainCompartment implements the OCIOperation interface (enables retrying operations)
func (client IotClient) changeIotDomainCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/iotDomains/{iotDomainId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeIotDomainCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ChangeIotDomainCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeIotDomainDataRetentionPeriod Updates Data Retention Period of the IoT Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ChangeIotDomainDataRetentionPeriod.go.html to see an example of how to use ChangeIotDomainDataRetentionPeriod API.
// A default retry strategy applies to this operation ChangeIotDomainDataRetentionPeriod()
func (client IotClient) ChangeIotDomainDataRetentionPeriod(ctx context.Context, request ChangeIotDomainDataRetentionPeriodRequest) (response ChangeIotDomainDataRetentionPeriodResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeIotDomainDataRetentionPeriod, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeIotDomainDataRetentionPeriodResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeIotDomainDataRetentionPeriodResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeIotDomainDataRetentionPeriodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeIotDomainDataRetentionPeriodResponse")
	}
	return
}

// changeIotDomainDataRetentionPeriod implements the OCIOperation interface (enables retrying operations)
func (client IotClient) changeIotDomainDataRetentionPeriod(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/iotDomains/{iotDomainId}/actions/changeDataRetentionPeriod", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeIotDomainDataRetentionPeriodResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ChangeIotDomainDataRetentionPeriod", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeIotDomainGroupCompartment Moves an IoT domain group to a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ChangeIotDomainGroupCompartment.go.html to see an example of how to use ChangeIotDomainGroupCompartment API.
// A default retry strategy applies to this operation ChangeIotDomainGroupCompartment()
func (client IotClient) ChangeIotDomainGroupCompartment(ctx context.Context, request ChangeIotDomainGroupCompartmentRequest) (response ChangeIotDomainGroupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeIotDomainGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeIotDomainGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeIotDomainGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeIotDomainGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeIotDomainGroupCompartmentResponse")
	}
	return
}

// changeIotDomainGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client IotClient) changeIotDomainGroupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/iotDomainGroups/{iotDomainGroupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeIotDomainGroupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ChangeIotDomainGroupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureIotDomainDataAccess Updates an IoT domain Data Access.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ConfigureIotDomainDataAccess.go.html to see an example of how to use ConfigureIotDomainDataAccess API.
// A default retry strategy applies to this operation ConfigureIotDomainDataAccess()
func (client IotClient) ConfigureIotDomainDataAccess(ctx context.Context, request ConfigureIotDomainDataAccessRequest) (response ConfigureIotDomainDataAccessResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureIotDomainDataAccess, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureIotDomainDataAccessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureIotDomainDataAccessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureIotDomainDataAccessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureIotDomainDataAccessResponse")
	}
	return
}

// configureIotDomainDataAccess implements the OCIOperation interface (enables retrying operations)
func (client IotClient) configureIotDomainDataAccess(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/iotDomains/{iotDomainId}/actions/configureDataAccess", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureIotDomainDataAccessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ConfigureIotDomainDataAccess", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ConfigureIotDomainGroupDataAccess Updates an IoT domain Group Data Access.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ConfigureIotDomainGroupDataAccess.go.html to see an example of how to use ConfigureIotDomainGroupDataAccess API.
// A default retry strategy applies to this operation ConfigureIotDomainGroupDataAccess()
func (client IotClient) ConfigureIotDomainGroupDataAccess(ctx context.Context, request ConfigureIotDomainGroupDataAccessRequest) (response ConfigureIotDomainGroupDataAccessResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.configureIotDomainGroupDataAccess, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureIotDomainGroupDataAccessResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureIotDomainGroupDataAccessResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureIotDomainGroupDataAccessResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureIotDomainGroupDataAccessResponse")
	}
	return
}

// configureIotDomainGroupDataAccess implements the OCIOperation interface (enables retrying operations)
func (client IotClient) configureIotDomainGroupDataAccess(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/iotDomainGroups/{iotDomainGroupId}/actions/configureDataAccess", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureIotDomainGroupDataAccessResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ConfigureIotDomainGroupDataAccess", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDigitalTwinAdapter Creates a new digital twin adapter.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/CreateDigitalTwinAdapter.go.html to see an example of how to use CreateDigitalTwinAdapter API.
// A default retry strategy applies to this operation CreateDigitalTwinAdapter()
func (client IotClient) CreateDigitalTwinAdapter(ctx context.Context, request CreateDigitalTwinAdapterRequest) (response CreateDigitalTwinAdapterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDigitalTwinAdapter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDigitalTwinAdapterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDigitalTwinAdapterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDigitalTwinAdapterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDigitalTwinAdapterResponse")
	}
	return
}

// createDigitalTwinAdapter implements the OCIOperation interface (enables retrying operations)
func (client IotClient) createDigitalTwinAdapter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/digitalTwinAdapters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDigitalTwinAdapterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "CreateDigitalTwinAdapter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDigitalTwinInstance Creates a new digital twin instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/CreateDigitalTwinInstance.go.html to see an example of how to use CreateDigitalTwinInstance API.
// A default retry strategy applies to this operation CreateDigitalTwinInstance()
func (client IotClient) CreateDigitalTwinInstance(ctx context.Context, request CreateDigitalTwinInstanceRequest) (response CreateDigitalTwinInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDigitalTwinInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDigitalTwinInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDigitalTwinInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDigitalTwinInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDigitalTwinInstanceResponse")
	}
	return
}

// createDigitalTwinInstance implements the OCIOperation interface (enables retrying operations)
func (client IotClient) createDigitalTwinInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/digitalTwinInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDigitalTwinInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "CreateDigitalTwinInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDigitalTwinModel Creates a new digital twin model.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/CreateDigitalTwinModel.go.html to see an example of how to use CreateDigitalTwinModel API.
// A default retry strategy applies to this operation CreateDigitalTwinModel()
func (client IotClient) CreateDigitalTwinModel(ctx context.Context, request CreateDigitalTwinModelRequest) (response CreateDigitalTwinModelResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDigitalTwinModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDigitalTwinModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDigitalTwinModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDigitalTwinModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDigitalTwinModelResponse")
	}
	return
}

// createDigitalTwinModel implements the OCIOperation interface (enables retrying operations)
func (client IotClient) createDigitalTwinModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/digitalTwinModels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDigitalTwinModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "CreateDigitalTwinModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDigitalTwinRelationship Creates a new digital twin relationship.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/CreateDigitalTwinRelationship.go.html to see an example of how to use CreateDigitalTwinRelationship API.
// A default retry strategy applies to this operation CreateDigitalTwinRelationship()
func (client IotClient) CreateDigitalTwinRelationship(ctx context.Context, request CreateDigitalTwinRelationshipRequest) (response CreateDigitalTwinRelationshipResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDigitalTwinRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDigitalTwinRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDigitalTwinRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDigitalTwinRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDigitalTwinRelationshipResponse")
	}
	return
}

// createDigitalTwinRelationship implements the OCIOperation interface (enables retrying operations)
func (client IotClient) createDigitalTwinRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/digitalTwinRelationships", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDigitalTwinRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "CreateDigitalTwinRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateIotDomain Creates a new IoT domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/CreateIotDomain.go.html to see an example of how to use CreateIotDomain API.
// A default retry strategy applies to this operation CreateIotDomain()
func (client IotClient) CreateIotDomain(ctx context.Context, request CreateIotDomainRequest) (response CreateIotDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createIotDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIotDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIotDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIotDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIotDomainResponse")
	}
	return
}

// createIotDomain implements the OCIOperation interface (enables retrying operations)
func (client IotClient) createIotDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/iotDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIotDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "CreateIotDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateIotDomainGroup Creates a new IoT domain group.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/CreateIotDomainGroup.go.html to see an example of how to use CreateIotDomainGroup API.
// A default retry strategy applies to this operation CreateIotDomainGroup()
func (client IotClient) CreateIotDomainGroup(ctx context.Context, request CreateIotDomainGroupRequest) (response CreateIotDomainGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createIotDomainGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIotDomainGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIotDomainGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIotDomainGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIotDomainGroupResponse")
	}
	return
}

// createIotDomainGroup implements the OCIOperation interface (enables retrying operations)
func (client IotClient) createIotDomainGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/iotDomainGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIotDomainGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "CreateIotDomainGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDigitalTwinAdapter Deletes the digital twin adapter identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/DeleteDigitalTwinAdapter.go.html to see an example of how to use DeleteDigitalTwinAdapter API.
// A default retry strategy applies to this operation DeleteDigitalTwinAdapter()
func (client IotClient) DeleteDigitalTwinAdapter(ctx context.Context, request DeleteDigitalTwinAdapterRequest) (response DeleteDigitalTwinAdapterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDigitalTwinAdapter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDigitalTwinAdapterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDigitalTwinAdapterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDigitalTwinAdapterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDigitalTwinAdapterResponse")
	}
	return
}

// deleteDigitalTwinAdapter implements the OCIOperation interface (enables retrying operations)
func (client IotClient) deleteDigitalTwinAdapter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/digitalTwinAdapters/{digitalTwinAdapterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDigitalTwinAdapterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "DeleteDigitalTwinAdapter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDigitalTwinInstance Deletes the digital twin instance identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/DeleteDigitalTwinInstance.go.html to see an example of how to use DeleteDigitalTwinInstance API.
// A default retry strategy applies to this operation DeleteDigitalTwinInstance()
func (client IotClient) DeleteDigitalTwinInstance(ctx context.Context, request DeleteDigitalTwinInstanceRequest) (response DeleteDigitalTwinInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDigitalTwinInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDigitalTwinInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDigitalTwinInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDigitalTwinInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDigitalTwinInstanceResponse")
	}
	return
}

// deleteDigitalTwinInstance implements the OCIOperation interface (enables retrying operations)
func (client IotClient) deleteDigitalTwinInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/digitalTwinInstances/{digitalTwinInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDigitalTwinInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "DeleteDigitalTwinInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDigitalTwinModel Deletes the digital twin model identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/DeleteDigitalTwinModel.go.html to see an example of how to use DeleteDigitalTwinModel API.
// A default retry strategy applies to this operation DeleteDigitalTwinModel()
func (client IotClient) DeleteDigitalTwinModel(ctx context.Context, request DeleteDigitalTwinModelRequest) (response DeleteDigitalTwinModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDigitalTwinModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDigitalTwinModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDigitalTwinModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDigitalTwinModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDigitalTwinModelResponse")
	}
	return
}

// deleteDigitalTwinModel implements the OCIOperation interface (enables retrying operations)
func (client IotClient) deleteDigitalTwinModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/digitalTwinModels/{digitalTwinModelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDigitalTwinModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "DeleteDigitalTwinModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDigitalTwinRelationship Deletes the digital twin relationship identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/DeleteDigitalTwinRelationship.go.html to see an example of how to use DeleteDigitalTwinRelationship API.
// A default retry strategy applies to this operation DeleteDigitalTwinRelationship()
func (client IotClient) DeleteDigitalTwinRelationship(ctx context.Context, request DeleteDigitalTwinRelationshipRequest) (response DeleteDigitalTwinRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDigitalTwinRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDigitalTwinRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDigitalTwinRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDigitalTwinRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDigitalTwinRelationshipResponse")
	}
	return
}

// deleteDigitalTwinRelationship implements the OCIOperation interface (enables retrying operations)
func (client IotClient) deleteDigitalTwinRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/digitalTwinRelationships/{digitalTwinRelationshipId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDigitalTwinRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "DeleteDigitalTwinRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteIotDomain Deletes the IoT domain identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/DeleteIotDomain.go.html to see an example of how to use DeleteIotDomain API.
// A default retry strategy applies to this operation DeleteIotDomain()
func (client IotClient) DeleteIotDomain(ctx context.Context, request DeleteIotDomainRequest) (response DeleteIotDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteIotDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteIotDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteIotDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteIotDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteIotDomainResponse")
	}
	return
}

// deleteIotDomain implements the OCIOperation interface (enables retrying operations)
func (client IotClient) deleteIotDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/iotDomains/{iotDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteIotDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "DeleteIotDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteIotDomainGroup Deletes the IoT domain group identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/DeleteIotDomainGroup.go.html to see an example of how to use DeleteIotDomainGroup API.
// A default retry strategy applies to this operation DeleteIotDomainGroup()
func (client IotClient) DeleteIotDomainGroup(ctx context.Context, request DeleteIotDomainGroupRequest) (response DeleteIotDomainGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteIotDomainGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteIotDomainGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteIotDomainGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteIotDomainGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteIotDomainGroupResponse")
	}
	return
}

// deleteIotDomainGroup implements the OCIOperation interface (enables retrying operations)
func (client IotClient) deleteIotDomainGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/iotDomainGroups/{iotDomainGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteIotDomainGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "DeleteIotDomainGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDigitalTwinAdapter Retrieves the digital twin adapter identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetDigitalTwinAdapter.go.html to see an example of how to use GetDigitalTwinAdapter API.
// A default retry strategy applies to this operation GetDigitalTwinAdapter()
func (client IotClient) GetDigitalTwinAdapter(ctx context.Context, request GetDigitalTwinAdapterRequest) (response GetDigitalTwinAdapterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalTwinAdapter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalTwinAdapterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalTwinAdapterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalTwinAdapterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalTwinAdapterResponse")
	}
	return
}

// getDigitalTwinAdapter implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getDigitalTwinAdapter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinAdapters/{digitalTwinAdapterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalTwinAdapterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetDigitalTwinAdapter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDigitalTwinInstance Retrieves the digital twin instance identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetDigitalTwinInstance.go.html to see an example of how to use GetDigitalTwinInstance API.
// A default retry strategy applies to this operation GetDigitalTwinInstance()
func (client IotClient) GetDigitalTwinInstance(ctx context.Context, request GetDigitalTwinInstanceRequest) (response GetDigitalTwinInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalTwinInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalTwinInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalTwinInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalTwinInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalTwinInstanceResponse")
	}
	return
}

// getDigitalTwinInstance implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getDigitalTwinInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinInstances/{digitalTwinInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalTwinInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetDigitalTwinInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDigitalTwinInstanceContent Retrieves the latest snapshot data of digital twin instance identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetDigitalTwinInstanceContent.go.html to see an example of how to use GetDigitalTwinInstanceContent API.
// A default retry strategy applies to this operation GetDigitalTwinInstanceContent()
func (client IotClient) GetDigitalTwinInstanceContent(ctx context.Context, request GetDigitalTwinInstanceContentRequest) (response GetDigitalTwinInstanceContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalTwinInstanceContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalTwinInstanceContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalTwinInstanceContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalTwinInstanceContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalTwinInstanceContentResponse")
	}
	return
}

// getDigitalTwinInstanceContent implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getDigitalTwinInstanceContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinInstances/{digitalTwinInstanceId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalTwinInstanceContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetDigitalTwinInstanceContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDigitalTwinModel Retrieves the digital twin model identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetDigitalTwinModel.go.html to see an example of how to use GetDigitalTwinModel API.
// A default retry strategy applies to this operation GetDigitalTwinModel()
func (client IotClient) GetDigitalTwinModel(ctx context.Context, request GetDigitalTwinModelRequest) (response GetDigitalTwinModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalTwinModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalTwinModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalTwinModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalTwinModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalTwinModelResponse")
	}
	return
}

// getDigitalTwinModel implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getDigitalTwinModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinModels/{digitalTwinModelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalTwinModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetDigitalTwinModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDigitalTwinModelSpec Retrieves the spec of digital twin model identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetDigitalTwinModelSpec.go.html to see an example of how to use GetDigitalTwinModelSpec API.
// A default retry strategy applies to this operation GetDigitalTwinModelSpec()
func (client IotClient) GetDigitalTwinModelSpec(ctx context.Context, request GetDigitalTwinModelSpecRequest) (response GetDigitalTwinModelSpecResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalTwinModelSpec, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalTwinModelSpecResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalTwinModelSpecResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalTwinModelSpecResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalTwinModelSpecResponse")
	}
	return
}

// getDigitalTwinModelSpec implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getDigitalTwinModelSpec(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinModels/{digitalTwinModelId}/spec", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalTwinModelSpecResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetDigitalTwinModelSpec", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDigitalTwinRelationship Retrieves the digital twin relationship identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetDigitalTwinRelationship.go.html to see an example of how to use GetDigitalTwinRelationship API.
// A default retry strategy applies to this operation GetDigitalTwinRelationship()
func (client IotClient) GetDigitalTwinRelationship(ctx context.Context, request GetDigitalTwinRelationshipRequest) (response GetDigitalTwinRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalTwinRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalTwinRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalTwinRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalTwinRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalTwinRelationshipResponse")
	}
	return
}

// getDigitalTwinRelationship implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getDigitalTwinRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinRelationships/{digitalTwinRelationshipId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalTwinRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetDigitalTwinRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIotDomain Retrieves the IoT domain identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetIotDomain.go.html to see an example of how to use GetIotDomain API.
// A default retry strategy applies to this operation GetIotDomain()
func (client IotClient) GetIotDomain(ctx context.Context, request GetIotDomainRequest) (response GetIotDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIotDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIotDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIotDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIotDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIotDomainResponse")
	}
	return
}

// getIotDomain implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getIotDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iotDomains/{iotDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIotDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetIotDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetIotDomainGroup Retrieves the IoT domain group identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetIotDomainGroup.go.html to see an example of how to use GetIotDomainGroup API.
// A default retry strategy applies to this operation GetIotDomainGroup()
func (client IotClient) GetIotDomainGroup(ctx context.Context, request GetIotDomainGroupRequest) (response GetIotDomainGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIotDomainGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIotDomainGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIotDomainGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIotDomainGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIotDomainGroupResponse")
	}
	return
}

// getIotDomainGroup implements the OCIOperation interface (enables retrying operations)
func (client IotClient) getIotDomainGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iotDomainGroups/{iotDomainGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIotDomainGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetIotDomainGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Retrieves the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client IotClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client IotClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InvokeRawCommand Invokes the raw command on the specified digital twin instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/InvokeRawCommand.go.html to see an example of how to use InvokeRawCommand API.
// A default retry strategy applies to this operation InvokeRawCommand()
func (client IotClient) InvokeRawCommand(ctx context.Context, request InvokeRawCommandRequest) (response InvokeRawCommandResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.invokeRawCommand, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InvokeRawCommandResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InvokeRawCommandResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InvokeRawCommandResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InvokeRawCommandResponse")
	}
	return
}

// invokeRawCommand implements the OCIOperation interface (enables retrying operations)
func (client IotClient) invokeRawCommand(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/digitalTwinInstances/{digitalTwinInstanceId}/actions/invokeRawCommand", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InvokeRawCommandResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "InvokeRawCommand", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDigitalTwinAdapters Retrieves a list of digital twin adapters within the specified IoT domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinAdapters.go.html to see an example of how to use ListDigitalTwinAdapters API.
// A default retry strategy applies to this operation ListDigitalTwinAdapters()
func (client IotClient) ListDigitalTwinAdapters(ctx context.Context, request ListDigitalTwinAdaptersRequest) (response ListDigitalTwinAdaptersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDigitalTwinAdapters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDigitalTwinAdaptersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDigitalTwinAdaptersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDigitalTwinAdaptersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDigitalTwinAdaptersResponse")
	}
	return
}

// listDigitalTwinAdapters implements the OCIOperation interface (enables retrying operations)
func (client IotClient) listDigitalTwinAdapters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinAdapters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDigitalTwinAdaptersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListDigitalTwinAdapters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDigitalTwinInstances Retrieves a list of digital twin instances within the specified IoT domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinInstances.go.html to see an example of how to use ListDigitalTwinInstances API.
// A default retry strategy applies to this operation ListDigitalTwinInstances()
func (client IotClient) ListDigitalTwinInstances(ctx context.Context, request ListDigitalTwinInstancesRequest) (response ListDigitalTwinInstancesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDigitalTwinInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDigitalTwinInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDigitalTwinInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDigitalTwinInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDigitalTwinInstancesResponse")
	}
	return
}

// listDigitalTwinInstances implements the OCIOperation interface (enables retrying operations)
func (client IotClient) listDigitalTwinInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDigitalTwinInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListDigitalTwinInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDigitalTwinModels Retrieves a list of digital twin models within the specified IoT domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinModels.go.html to see an example of how to use ListDigitalTwinModels API.
// A default retry strategy applies to this operation ListDigitalTwinModels()
func (client IotClient) ListDigitalTwinModels(ctx context.Context, request ListDigitalTwinModelsRequest) (response ListDigitalTwinModelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDigitalTwinModels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDigitalTwinModelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDigitalTwinModelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDigitalTwinModelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDigitalTwinModelsResponse")
	}
	return
}

// listDigitalTwinModels implements the OCIOperation interface (enables retrying operations)
func (client IotClient) listDigitalTwinModels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinModels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDigitalTwinModelsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListDigitalTwinModels", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDigitalTwinRelationships Retrieves a list of digital twin relationships within the specified IoT domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListDigitalTwinRelationships.go.html to see an example of how to use ListDigitalTwinRelationships API.
// A default retry strategy applies to this operation ListDigitalTwinRelationships()
func (client IotClient) ListDigitalTwinRelationships(ctx context.Context, request ListDigitalTwinRelationshipsRequest) (response ListDigitalTwinRelationshipsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDigitalTwinRelationships, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDigitalTwinRelationshipsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDigitalTwinRelationshipsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDigitalTwinRelationshipsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDigitalTwinRelationshipsResponse")
	}
	return
}

// listDigitalTwinRelationships implements the OCIOperation interface (enables retrying operations)
func (client IotClient) listDigitalTwinRelationships(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/digitalTwinRelationships", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDigitalTwinRelationshipsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListDigitalTwinRelationships", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIotDomainGroups Retrieves a list of IoT domain groups within the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListIotDomainGroups.go.html to see an example of how to use ListIotDomainGroups API.
// A default retry strategy applies to this operation ListIotDomainGroups()
func (client IotClient) ListIotDomainGroups(ctx context.Context, request ListIotDomainGroupsRequest) (response ListIotDomainGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIotDomainGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIotDomainGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIotDomainGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIotDomainGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIotDomainGroupsResponse")
	}
	return
}

// listIotDomainGroups implements the OCIOperation interface (enables retrying operations)
func (client IotClient) listIotDomainGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iotDomainGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIotDomainGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListIotDomainGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListIotDomains Retrieves a list of IoT domains within the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListIotDomains.go.html to see an example of how to use ListIotDomains API.
// A default retry strategy applies to this operation ListIotDomains()
func (client IotClient) ListIotDomains(ctx context.Context, request ListIotDomainsRequest) (response ListIotDomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIotDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIotDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIotDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIotDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIotDomainsResponse")
	}
	return
}

// listIotDomains implements the OCIOperation interface (enables retrying operations)
func (client IotClient) listIotDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iotDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIotDomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListIotDomains", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists errors associated with the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client IotClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client IotClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs associated with the specified work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client IotClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client IotClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists work requests in the specified compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client IotClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client IotClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDigitalTwinAdapter Updates the details of digital twin adapter identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/UpdateDigitalTwinAdapter.go.html to see an example of how to use UpdateDigitalTwinAdapter API.
// A default retry strategy applies to this operation UpdateDigitalTwinAdapter()
func (client IotClient) UpdateDigitalTwinAdapter(ctx context.Context, request UpdateDigitalTwinAdapterRequest) (response UpdateDigitalTwinAdapterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDigitalTwinAdapter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDigitalTwinAdapterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDigitalTwinAdapterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDigitalTwinAdapterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDigitalTwinAdapterResponse")
	}
	return
}

// updateDigitalTwinAdapter implements the OCIOperation interface (enables retrying operations)
func (client IotClient) updateDigitalTwinAdapter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/digitalTwinAdapters/{digitalTwinAdapterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDigitalTwinAdapterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "UpdateDigitalTwinAdapter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDigitalTwinInstance Updates the details of digital twin instance identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/UpdateDigitalTwinInstance.go.html to see an example of how to use UpdateDigitalTwinInstance API.
// A default retry strategy applies to this operation UpdateDigitalTwinInstance()
func (client IotClient) UpdateDigitalTwinInstance(ctx context.Context, request UpdateDigitalTwinInstanceRequest) (response UpdateDigitalTwinInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDigitalTwinInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDigitalTwinInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDigitalTwinInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDigitalTwinInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDigitalTwinInstanceResponse")
	}
	return
}

// updateDigitalTwinInstance implements the OCIOperation interface (enables retrying operations)
func (client IotClient) updateDigitalTwinInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/digitalTwinInstances/{digitalTwinInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDigitalTwinInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "UpdateDigitalTwinInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDigitalTwinModel Updates the details of the digital twin model identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/UpdateDigitalTwinModel.go.html to see an example of how to use UpdateDigitalTwinModel API.
// A default retry strategy applies to this operation UpdateDigitalTwinModel()
func (client IotClient) UpdateDigitalTwinModel(ctx context.Context, request UpdateDigitalTwinModelRequest) (response UpdateDigitalTwinModelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDigitalTwinModel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDigitalTwinModelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDigitalTwinModelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDigitalTwinModelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDigitalTwinModelResponse")
	}
	return
}

// updateDigitalTwinModel implements the OCIOperation interface (enables retrying operations)
func (client IotClient) updateDigitalTwinModel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/digitalTwinModels/{digitalTwinModelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDigitalTwinModelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "UpdateDigitalTwinModel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDigitalTwinRelationship Updates the details of digital twin relationship identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/UpdateDigitalTwinRelationship.go.html to see an example of how to use UpdateDigitalTwinRelationship API.
// A default retry strategy applies to this operation UpdateDigitalTwinRelationship()
func (client IotClient) UpdateDigitalTwinRelationship(ctx context.Context, request UpdateDigitalTwinRelationshipRequest) (response UpdateDigitalTwinRelationshipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDigitalTwinRelationship, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDigitalTwinRelationshipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDigitalTwinRelationshipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDigitalTwinRelationshipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDigitalTwinRelationshipResponse")
	}
	return
}

// updateDigitalTwinRelationship implements the OCIOperation interface (enables retrying operations)
func (client IotClient) updateDigitalTwinRelationship(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/digitalTwinRelationships/{digitalTwinRelationshipId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDigitalTwinRelationshipResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "UpdateDigitalTwinRelationship", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateIotDomain Updates the details of IoT domain identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/UpdateIotDomain.go.html to see an example of how to use UpdateIotDomain API.
// A default retry strategy applies to this operation UpdateIotDomain()
func (client IotClient) UpdateIotDomain(ctx context.Context, request UpdateIotDomainRequest) (response UpdateIotDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateIotDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateIotDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateIotDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIotDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIotDomainResponse")
	}
	return
}

// updateIotDomain implements the OCIOperation interface (enables retrying operations)
func (client IotClient) updateIotDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/iotDomains/{iotDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateIotDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "UpdateIotDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateIotDomainGroup Updates the details of IoT domain group identified by the specified OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/iot/UpdateIotDomainGroup.go.html to see an example of how to use UpdateIotDomainGroup API.
// A default retry strategy applies to this operation UpdateIotDomainGroup()
func (client IotClient) UpdateIotDomainGroup(ctx context.Context, request UpdateIotDomainGroupRequest) (response UpdateIotDomainGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateIotDomainGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateIotDomainGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateIotDomainGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIotDomainGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIotDomainGroupResponse")
	}
	return
}

// updateIotDomainGroup implements the OCIOperation interface (enables retrying operations)
func (client IotClient) updateIotDomainGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/iotDomainGroups/{iotDomainGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateIotDomainGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Iot", "UpdateIotDomainGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
