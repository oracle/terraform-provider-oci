// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// EventClient a client for Event
type EventClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEventClientWithConfigurationProvider Creates a new default Event client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEventClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EventClient, err error) {
	if enabled := common.CheckForEnabledServices("osmanagementhub"); !enabled {
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
	return newEventClientFromBaseClient(baseClient, provider)
}

// NewEventClientWithOboToken Creates a new default Event client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewEventClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EventClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEventClientFromBaseClient(baseClient, configProvider)
}

func newEventClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EventClient, err error) {
	// Event service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Event"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EventClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EventClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EventClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *EventClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeEventCompartment Moves an event into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ChangeEventCompartment.go.html to see an example of how to use ChangeEventCompartment API.
// A default retry strategy applies to this operation ChangeEventCompartment()
func (client EventClient) ChangeEventCompartment(ctx context.Context, request ChangeEventCompartmentRequest) (response ChangeEventCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEventCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEventCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEventCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEventCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEventCompartmentResponse")
	}
	return
}

// changeEventCompartment implements the OCIOperation interface (enables retrying operations)
func (client EventClient) changeEventCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/events/{eventId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEventCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Event/ChangeEventCompartment"
		err = common.PostProcessServiceError(err, "Event", "ChangeEventCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEvent Deletes the specified event.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteEvent.go.html to see an example of how to use DeleteEvent API.
// A default retry strategy applies to this operation DeleteEvent()
func (client EventClient) DeleteEvent(ctx context.Context, request DeleteEventRequest) (response DeleteEventResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEvent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEventResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEventResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEventResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEventResponse")
	}
	return
}

// deleteEvent implements the OCIOperation interface (enables retrying operations)
func (client EventClient) deleteEvent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/events/{eventId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEventResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Event/DeleteEvent"
		err = common.PostProcessServiceError(err, "Event", "DeleteEvent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEventContent Deletes an event content ZIP archive from the service.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteEventContent.go.html to see an example of how to use DeleteEventContent API.
// A default retry strategy applies to this operation DeleteEventContent()
func (client EventClient) DeleteEventContent(ctx context.Context, request DeleteEventContentRequest) (response DeleteEventContentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteEventContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEventContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEventContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEventContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEventContentResponse")
	}
	return
}

// deleteEventContent implements the OCIOperation interface (enables retrying operations)
func (client EventClient) deleteEventContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/events/{eventId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEventContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Event/DeleteEventContent"
		err = common.PostProcessServiceError(err, "Event", "DeleteEventContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEvent Returns information about the specified event.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetEvent.go.html to see an example of how to use GetEvent API.
// A default retry strategy applies to this operation GetEvent()
func (client EventClient) GetEvent(ctx context.Context, request GetEventRequest) (response GetEventResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEvent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEventResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEventResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEventResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEventResponse")
	}
	return
}

// getEvent implements the OCIOperation interface (enables retrying operations)
func (client EventClient) getEvent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/events/{eventId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEventResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Event/GetEvent"
		err = common.PostProcessServiceError(err, "Event", "GetEvent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &event{})
	return response, err
}

// GetEventContent Returns a ZIP archive with additional information about an event. The archive content depends on the event type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetEventContent.go.html to see an example of how to use GetEventContent API.
// A default retry strategy applies to this operation GetEventContent()
func (client EventClient) GetEventContent(ctx context.Context, request GetEventContentRequest) (response GetEventContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEventContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEventContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEventContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEventContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEventContentResponse")
	}
	return
}

// getEventContent implements the OCIOperation interface (enables retrying operations)
func (client EventClient) getEventContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/events/{eventId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEventContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Event/GetEventContent"
		err = common.PostProcessServiceError(err, "Event", "GetEventContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportEventContent Uploads an event content ZIP archive from an instance to the service.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ImportEventContent.go.html to see an example of how to use ImportEventContent API.
// A default retry strategy applies to this operation ImportEventContent()
func (client EventClient) ImportEventContent(ctx context.Context, request ImportEventContentRequest) (response ImportEventContentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importEventContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportEventContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportEventContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportEventContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportEventContentResponse")
	}
	return
}

// importEventContent implements the OCIOperation interface (enables retrying operations)
func (client EventClient) importEventContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/events/{eventId}/actions/importContent", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportEventContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Event/ImportEventContent"
		err = common.PostProcessServiceError(err, "Event", "ImportEventContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEvents Lists events that match the specified criteria, such as compartment, state, and event type.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListEvents.go.html to see an example of how to use ListEvents API.
// A default retry strategy applies to this operation ListEvents()
func (client EventClient) ListEvents(ctx context.Context, request ListEventsRequest) (response ListEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEventsResponse")
	}
	return
}

// listEvents implements the OCIOperation interface (enables retrying operations)
func (client EventClient) listEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/events", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/EventCollection/ListEvents"
		err = common.PostProcessServiceError(err, "Event", "ListEvents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEvent Updates the tags for an event.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateEvent.go.html to see an example of how to use UpdateEvent API.
// A default retry strategy applies to this operation UpdateEvent()
func (client EventClient) UpdateEvent(ctx context.Context, request UpdateEventRequest) (response UpdateEventResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEvent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEventResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEventResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEventResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEventResponse")
	}
	return
}

// updateEvent implements the OCIOperation interface (enables retrying operations)
func (client EventClient) updateEvent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/events/{eventId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEventResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/Event/UpdateEvent"
		err = common.PostProcessServiceError(err, "Event", "UpdateEvent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &event{})
	return response, err
}
