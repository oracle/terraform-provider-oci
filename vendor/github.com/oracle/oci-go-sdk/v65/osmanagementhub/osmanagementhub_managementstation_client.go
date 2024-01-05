// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ManagementStationClient a client for ManagementStation
type ManagementStationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagementStationClientWithConfigurationProvider Creates a new default ManagementStation client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagementStationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagementStationClient, err error) {
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
	return newManagementStationClientFromBaseClient(baseClient, provider)
}

// NewManagementStationClientWithOboToken Creates a new default ManagementStation client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewManagementStationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ManagementStationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newManagementStationClientFromBaseClient(baseClient, configProvider)
}

func newManagementStationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ManagementStationClient, err error) {
	// ManagementStation service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ManagementStation"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ManagementStationClient{BaseClient: baseClient}
	client.BasePath = "20220901"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ManagementStationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("osmanagementhub", "https://osmh.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ManagementStationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ManagementStationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateManagementStation Creates a management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/CreateManagementStation.go.html to see an example of how to use CreateManagementStation API.
// A default retry strategy applies to this operation CreateManagementStation()
func (client ManagementStationClient) CreateManagementStation(ctx context.Context, request CreateManagementStationRequest) (response CreateManagementStationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createManagementStation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateManagementStationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateManagementStationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateManagementStationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateManagementStationResponse")
	}
	return
}

// createManagementStation implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) createManagementStation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementStations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateManagementStationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagementStation/CreateManagementStation"
		err = common.PostProcessServiceError(err, "ManagementStation", "CreateManagementStation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteManagementStation Deletes a management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/DeleteManagementStation.go.html to see an example of how to use DeleteManagementStation API.
// A default retry strategy applies to this operation DeleteManagementStation()
func (client ManagementStationClient) DeleteManagementStation(ctx context.Context, request DeleteManagementStationRequest) (response DeleteManagementStationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteManagementStation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteManagementStationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteManagementStationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteManagementStationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteManagementStationResponse")
	}
	return
}

// deleteManagementStation implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) deleteManagementStation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/managementStations/{managementStationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteManagementStationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagementStation/DeleteManagementStation"
		err = common.PostProcessServiceError(err, "ManagementStation", "DeleteManagementStation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagementStation Gets information about the specified management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/GetManagementStation.go.html to see an example of how to use GetManagementStation API.
// A default retry strategy applies to this operation GetManagementStation()
func (client ManagementStationClient) GetManagementStation(ctx context.Context, request GetManagementStationRequest) (response GetManagementStationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagementStation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagementStationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagementStationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagementStationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagementStationResponse")
	}
	return
}

// getManagementStation implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) getManagementStation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementStations/{managementStationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagementStationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagementStation/GetManagementStation"
		err = common.PostProcessServiceError(err, "ManagementStation", "GetManagementStation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagementStations Lists management stations in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListManagementStations.go.html to see an example of how to use ListManagementStations API.
// A default retry strategy applies to this operation ListManagementStations()
func (client ManagementStationClient) ListManagementStations(ctx context.Context, request ListManagementStationsRequest) (response ListManagementStationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagementStations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagementStationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagementStationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagementStationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagementStationsResponse")
	}
	return
}

// listManagementStations implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) listManagementStations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementStations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagementStationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagementStation/ListManagementStations"
		err = common.PostProcessServiceError(err, "ManagementStation", "ListManagementStations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMirrors Lists all software source mirrors associated with a specified management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/ListMirrors.go.html to see an example of how to use ListMirrors API.
// A default retry strategy applies to this operation ListMirrors()
func (client ManagementStationClient) ListMirrors(ctx context.Context, request ListMirrorsRequest) (response ListMirrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMirrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMirrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMirrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMirrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMirrorsResponse")
	}
	return
}

// listMirrors implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) listMirrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managementStations/{managementStationId}/mirrors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMirrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/MirrorsCollection/ListMirrors"
		err = common.PostProcessServiceError(err, "ManagementStation", "ListMirrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SynchronizeMirrors Synchronizes the specified mirrors associated with the management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/SynchronizeMirrors.go.html to see an example of how to use SynchronizeMirrors API.
// A default retry strategy applies to this operation SynchronizeMirrors()
func (client ManagementStationClient) SynchronizeMirrors(ctx context.Context, request SynchronizeMirrorsRequest) (response SynchronizeMirrorsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.synchronizeMirrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SynchronizeMirrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SynchronizeMirrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SynchronizeMirrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SynchronizeMirrorsResponse")
	}
	return
}

// synchronizeMirrors implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) synchronizeMirrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementStations/{managementStationId}/actions/synchronizeMirrors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SynchronizeMirrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagementStation/SynchronizeMirrors"
		err = common.PostProcessServiceError(err, "ManagementStation", "SynchronizeMirrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SynchronizeSingleMirrors Synchronize the specified mirror associated with a management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/SynchronizeSingleMirrors.go.html to see an example of how to use SynchronizeSingleMirrors API.
// A default retry strategy applies to this operation SynchronizeSingleMirrors()
func (client ManagementStationClient) SynchronizeSingleMirrors(ctx context.Context, request SynchronizeSingleMirrorsRequest) (response SynchronizeSingleMirrorsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.synchronizeSingleMirrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SynchronizeSingleMirrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SynchronizeSingleMirrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SynchronizeSingleMirrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SynchronizeSingleMirrorsResponse")
	}
	return
}

// synchronizeSingleMirrors implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) synchronizeSingleMirrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managementStations/{managementStationId}/mirrors/{mirrorId}/actions/synchronize", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SynchronizeSingleMirrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagementStation/SynchronizeSingleMirrors"
		err = common.PostProcessServiceError(err, "ManagementStation", "SynchronizeSingleMirrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagementStation Updates the configuration of the specified management station.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagementhub/UpdateManagementStation.go.html to see an example of how to use UpdateManagementStation API.
// A default retry strategy applies to this operation UpdateManagementStation()
func (client ManagementStationClient) UpdateManagementStation(ctx context.Context, request UpdateManagementStationRequest) (response UpdateManagementStationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateManagementStation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagementStationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagementStationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagementStationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagementStationResponse")
	}
	return
}

// updateManagementStation implements the OCIOperation interface (enables retrying operations)
func (client ManagementStationClient) updateManagementStation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managementStations/{managementStationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagementStationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/osmh/20220901/ManagementStation/UpdateManagementStation"
		err = common.PostProcessServiceError(err, "ManagementStation", "UpdateManagementStation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
