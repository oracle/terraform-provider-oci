// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FleetAppsManagementMaintenanceWindowClient a client for FleetAppsManagementMaintenanceWindow
type FleetAppsManagementMaintenanceWindowClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFleetAppsManagementMaintenanceWindowClientWithConfigurationProvider Creates a new default FleetAppsManagementMaintenanceWindow client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFleetAppsManagementMaintenanceWindowClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FleetAppsManagementMaintenanceWindowClient, err error) {
	if enabled := common.CheckForEnabledServices("fleetappsmanagement"); !enabled {
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
	return newFleetAppsManagementMaintenanceWindowClientFromBaseClient(baseClient, provider)
}

// NewFleetAppsManagementMaintenanceWindowClientWithOboToken Creates a new default FleetAppsManagementMaintenanceWindow client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFleetAppsManagementMaintenanceWindowClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FleetAppsManagementMaintenanceWindowClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFleetAppsManagementMaintenanceWindowClientFromBaseClient(baseClient, configProvider)
}

func newFleetAppsManagementMaintenanceWindowClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FleetAppsManagementMaintenanceWindowClient, err error) {
	// FleetAppsManagementMaintenanceWindow service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FleetAppsManagementMaintenanceWindow"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FleetAppsManagementMaintenanceWindowClient{BaseClient: baseClient}
	client.BasePath = "20230831"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FleetAppsManagementMaintenanceWindowClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fleetappsmanagement", "https://fams.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FleetAppsManagementMaintenanceWindowClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FleetAppsManagementMaintenanceWindowClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateMaintenanceWindow Creates a new MaintenanceWindow.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/CreateMaintenanceWindow.go.html to see an example of how to use CreateMaintenanceWindow API.
// A default retry strategy applies to this operation CreateMaintenanceWindow()
func (client FleetAppsManagementMaintenanceWindowClient) CreateMaintenanceWindow(ctx context.Context, request CreateMaintenanceWindowRequest) (response CreateMaintenanceWindowResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMaintenanceWindowResponse")
	}
	return
}

// createMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementMaintenanceWindowClient) createMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/maintenanceWindows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/MaintenanceWindow/CreateMaintenanceWindow"
		err = common.PostProcessServiceError(err, "FleetAppsManagementMaintenanceWindow", "CreateMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMaintenanceWindow Deletes a MaintenanceWindow resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/DeleteMaintenanceWindow.go.html to see an example of how to use DeleteMaintenanceWindow API.
// A default retry strategy applies to this operation DeleteMaintenanceWindow()
func (client FleetAppsManagementMaintenanceWindowClient) DeleteMaintenanceWindow(ctx context.Context, request DeleteMaintenanceWindowRequest) (response DeleteMaintenanceWindowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMaintenanceWindowResponse")
	}
	return
}

// deleteMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementMaintenanceWindowClient) deleteMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/maintenanceWindows/{maintenanceWindowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/MaintenanceWindow/DeleteMaintenanceWindow"
		err = common.PostProcessServiceError(err, "FleetAppsManagementMaintenanceWindow", "DeleteMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMaintenanceWindow Gets a MaintenanceWindow by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetMaintenanceWindow.go.html to see an example of how to use GetMaintenanceWindow API.
// A default retry strategy applies to this operation GetMaintenanceWindow()
func (client FleetAppsManagementMaintenanceWindowClient) GetMaintenanceWindow(ctx context.Context, request GetMaintenanceWindowRequest) (response GetMaintenanceWindowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMaintenanceWindowResponse")
	}
	return
}

// getMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementMaintenanceWindowClient) getMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceWindows/{maintenanceWindowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/MaintenanceWindow/GetMaintenanceWindow"
		err = common.PostProcessServiceError(err, "FleetAppsManagementMaintenanceWindow", "GetMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMaintenanceWindows Returns a list of MaintenanceWindows in the specified Tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListMaintenanceWindows.go.html to see an example of how to use ListMaintenanceWindows API.
// A default retry strategy applies to this operation ListMaintenanceWindows()
func (client FleetAppsManagementMaintenanceWindowClient) ListMaintenanceWindows(ctx context.Context, request ListMaintenanceWindowsRequest) (response ListMaintenanceWindowsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMaintenanceWindows, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMaintenanceWindowsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMaintenanceWindowsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMaintenanceWindowsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMaintenanceWindowsResponse")
	}
	return
}

// listMaintenanceWindows implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementMaintenanceWindowClient) listMaintenanceWindows(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/maintenanceWindows", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMaintenanceWindowsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/MaintenanceWindowCollection/ListMaintenanceWindows"
		err = common.PostProcessServiceError(err, "FleetAppsManagementMaintenanceWindow", "ListMaintenanceWindows", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMaintenanceWindow Updates the MaintenanceWindow
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/UpdateMaintenanceWindow.go.html to see an example of how to use UpdateMaintenanceWindow API.
// A default retry strategy applies to this operation UpdateMaintenanceWindow()
func (client FleetAppsManagementMaintenanceWindowClient) UpdateMaintenanceWindow(ctx context.Context, request UpdateMaintenanceWindowRequest) (response UpdateMaintenanceWindowResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMaintenanceWindow, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMaintenanceWindowResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMaintenanceWindowResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMaintenanceWindowResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMaintenanceWindowResponse")
	}
	return
}

// updateMaintenanceWindow implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementMaintenanceWindowClient) updateMaintenanceWindow(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/maintenanceWindows/{maintenanceWindowId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMaintenanceWindowResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/MaintenanceWindow/UpdateMaintenanceWindow"
		err = common.PostProcessServiceError(err, "FleetAppsManagementMaintenanceWindow", "UpdateMaintenanceWindow", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
