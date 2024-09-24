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

// FleetAppsManagementRunbooksClient a client for FleetAppsManagementRunbooks
type FleetAppsManagementRunbooksClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFleetAppsManagementRunbooksClientWithConfigurationProvider Creates a new default FleetAppsManagementRunbooks client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFleetAppsManagementRunbooksClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FleetAppsManagementRunbooksClient, err error) {
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
	return newFleetAppsManagementRunbooksClientFromBaseClient(baseClient, provider)
}

// NewFleetAppsManagementRunbooksClientWithOboToken Creates a new default FleetAppsManagementRunbooks client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFleetAppsManagementRunbooksClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FleetAppsManagementRunbooksClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFleetAppsManagementRunbooksClientFromBaseClient(baseClient, configProvider)
}

func newFleetAppsManagementRunbooksClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FleetAppsManagementRunbooksClient, err error) {
	// FleetAppsManagementRunbooks service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FleetAppsManagementRunbooks"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FleetAppsManagementRunbooksClient{BaseClient: baseClient}
	client.BasePath = "20230831"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FleetAppsManagementRunbooksClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fleetappsmanagement", "https://fams.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FleetAppsManagementRunbooksClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FleetAppsManagementRunbooksClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetRunbook Gets a Runbook by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetRunbook.go.html to see an example of how to use GetRunbook API.
// A default retry strategy applies to this operation GetRunbook()
func (client FleetAppsManagementRunbooksClient) GetRunbook(ctx context.Context, request GetRunbookRequest) (response GetRunbookResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRunbook, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRunbookResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRunbookResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRunbookResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRunbookResponse")
	}
	return
}

// getRunbook implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) getRunbook(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks/{runbookId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRunbookResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/Runbook/GetRunbook"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetRunbook", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTaskRecord Gets a TaskRecord by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/GetTaskRecord.go.html to see an example of how to use GetTaskRecord API.
// A default retry strategy applies to this operation GetTaskRecord()
func (client FleetAppsManagementRunbooksClient) GetTaskRecord(ctx context.Context, request GetTaskRecordRequest) (response GetTaskRecordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTaskRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTaskRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTaskRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTaskRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTaskRecordResponse")
	}
	return
}

// getTaskRecord implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) getTaskRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taskRecords/{taskRecordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTaskRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/TaskRecord/GetTaskRecord"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "GetTaskRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRunbooks Returns a list of Runbooks.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListRunbooks.go.html to see an example of how to use ListRunbooks API.
// A default retry strategy applies to this operation ListRunbooks()
func (client FleetAppsManagementRunbooksClient) ListRunbooks(ctx context.Context, request ListRunbooksRequest) (response ListRunbooksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRunbooks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRunbooksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRunbooksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRunbooksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRunbooksResponse")
	}
	return
}

// listRunbooks implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) listRunbooks(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/runbooks", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRunbooksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/RunbookCollection/ListRunbooks"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListRunbooks", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaskRecords Returns a list of TaskRecords.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListTaskRecords.go.html to see an example of how to use ListTaskRecords API.
// A default retry strategy applies to this operation ListTaskRecords()
func (client FleetAppsManagementRunbooksClient) ListTaskRecords(ctx context.Context, request ListTaskRecordsRequest) (response ListTaskRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaskRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaskRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaskRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaskRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaskRecordsResponse")
	}
	return
}

// listTaskRecords implements the OCIOperation interface (enables retrying operations)
func (client FleetAppsManagementRunbooksClient) listTaskRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taskRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaskRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/fleet-management/20230831/TaskRecordCollection/ListTaskRecords"
		err = common.PostProcessServiceError(err, "FleetAppsManagementRunbooks", "ListTaskRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
