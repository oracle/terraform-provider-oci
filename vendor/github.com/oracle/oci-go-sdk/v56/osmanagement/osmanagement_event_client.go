// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//EventClient a client for Event
type EventClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEventClientWithConfigurationProvider Creates a new default Event client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEventClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EventClient, err error) {
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
//  as well as reading the region
func NewEventClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EventClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEventClientFromBaseClient(baseClient, configProvider)
}

func newEventClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EventClient, err error) {
	// Event service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EventClient{BaseClient: baseClient}
	client.BasePath = "20190801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EventClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagement", "https://osms.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EventClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *EventClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DeleteEventContent Delete an event content ZIP archive from the service
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/DeleteEventContent.go.html to see an example of how to use DeleteEventContent API.
func (client EventClient) DeleteEventContent(ctx context.Context, request DeleteEventContentRequest) (response DeleteEventContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managedInstances/{managedInstanceId}/events/{eventId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEventContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEvent Gets an Event by identifier
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetEvent.go.html to see an example of how to use GetEvent API.
func (client EventClient) GetEvent(ctx context.Context, request GetEventRequest) (response GetEventResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/events/{eventId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEventResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &event{})
	return response, err
}

// GetEventContent Get additional data about a event as a ZIP archive. The archive content depends on the event eventType.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetEventContent.go.html to see an example of how to use GetEventContent API.
func (client EventClient) GetEventContent(ctx context.Context, request GetEventContentRequest) (response GetEventContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/events/{eventId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEventContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEventReport Get summary information about events on this instance.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetEventReport.go.html to see an example of how to use GetEventReport API.
func (client EventClient) GetEventReport(ctx context.Context, request GetEventReportRequest) (response GetEventReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEventReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEventReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEventReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEventReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEventReportResponse")
	}
	return
}

// getEventReport implements the OCIOperation interface (enables retrying operations)
func (client EventClient) getEventReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/events/report", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEventReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEvents Returns a list of Events.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListEvents.go.html to see an example of how to use ListEvents API.
func (client EventClient) ListEvents(ctx context.Context, request ListEventsRequest) (response ListEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/events", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRelatedEvents Returns a list of related events. For now pagination is not implemented.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListRelatedEvents.go.html to see an example of how to use ListRelatedEvents API.
func (client EventClient) ListRelatedEvents(ctx context.Context, request ListRelatedEventsRequest) (response ListRelatedEventsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRelatedEvents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRelatedEventsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRelatedEventsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRelatedEventsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRelatedEventsResponse")
	}
	return
}

// listRelatedEvents implements the OCIOperation interface (enables retrying operations)
func (client EventClient) listRelatedEvents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/relatedEvents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRelatedEventsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEvent Updates an existing event associated to a managed instance
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/UpdateEvent.go.html to see an example of how to use UpdateEvent API.
func (client EventClient) UpdateEvent(ctx context.Context, request UpdateEventRequest) (response UpdateEventResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedInstances/{managedInstanceId}/events/{eventId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEventResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &event{})
	return response, err
}

// UploadEventContent Upload the event content as a ZIP archive from the managed instance to the service
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/UploadEventContent.go.html to see an example of how to use UploadEventContent API.
func (client EventClient) UploadEventContent(ctx context.Context, request UploadEventContentRequest) (response UploadEventContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.uploadEventContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadEventContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadEventContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadEventContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadEventContentResponse")
	}
	return
}

// uploadEventContent implements the OCIOperation interface (enables retrying operations)
func (client EventClient) uploadEventContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/events/{eventId}/content/actions/upload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadEventContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
