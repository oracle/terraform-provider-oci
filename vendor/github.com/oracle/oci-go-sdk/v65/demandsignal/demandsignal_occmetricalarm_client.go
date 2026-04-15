// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OccMetricAlarmClient a client for OccMetricAlarm
type OccMetricAlarmClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOccMetricAlarmClientWithConfigurationProvider Creates a new default OccMetricAlarm client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOccMetricAlarmClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OccMetricAlarmClient, err error) {
	if enabled := common.CheckForEnabledServices("demandsignal"); !enabled {
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
	return newOccMetricAlarmClientFromBaseClient(baseClient, provider)
}

// NewOccMetricAlarmClientWithOboToken Creates a new default OccMetricAlarm client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOccMetricAlarmClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OccMetricAlarmClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOccMetricAlarmClientFromBaseClient(baseClient, configProvider)
}

func newOccMetricAlarmClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OccMetricAlarmClient, err error) {
	// OccMetricAlarm service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OccMetricAlarm"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OccMetricAlarmClient{BaseClient: baseClient}
	client.BasePath = "20240430"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OccMetricAlarmClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("demandsignal", "https://control-center-ds.{region}.oci.{secondLevelDomain}", "control-center-ds")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OccMetricAlarmClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OccMetricAlarmClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateOccMetricAlarm Creates a new OccMetricAlarm resource in the specified compartment with the provided configuration details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/CreateOccMetricAlarm.go.html to see an example of how to use CreateOccMetricAlarm API.
// A default retry strategy applies to this operation CreateOccMetricAlarm()
func (client OccMetricAlarmClient) CreateOccMetricAlarm(ctx context.Context, request CreateOccMetricAlarmRequest) (response CreateOccMetricAlarmResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccMetricAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccMetricAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccMetricAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccMetricAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccMetricAlarmResponse")
	}
	return
}

// createOccMetricAlarm implements the OCIOperation interface (enables retrying operations)
func (client OccMetricAlarmClient) createOccMetricAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occMetricAlarms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccMetricAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "occMetricAlarm", "CreateOccMetricAlarm")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccMetricAlarm/CreateOccMetricAlarm"
		err = common.PostProcessServiceError(err, "OccMetricAlarm", "CreateOccMetricAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccMetricAlarm Deletes the specified OccMetricAlarm resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/DeleteOccMetricAlarm.go.html to see an example of how to use DeleteOccMetricAlarm API.
// A default retry strategy applies to this operation DeleteOccMetricAlarm()
func (client OccMetricAlarmClient) DeleteOccMetricAlarm(ctx context.Context, request DeleteOccMetricAlarmRequest) (response DeleteOccMetricAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccMetricAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccMetricAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccMetricAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccMetricAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccMetricAlarmResponse")
	}
	return
}

// deleteOccMetricAlarm implements the OCIOperation interface (enables retrying operations)
func (client OccMetricAlarmClient) deleteOccMetricAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occMetricAlarms/{occMetricAlarmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccMetricAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "occMetricAlarm", "DeleteOccMetricAlarm")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccMetricAlarm/DeleteOccMetricAlarm"
		err = common.PostProcessServiceError(err, "OccMetricAlarm", "DeleteOccMetricAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccMetricAlarm Retrieves the specified OccMetricAlarm resource based on its identifier.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/GetOccMetricAlarm.go.html to see an example of how to use GetOccMetricAlarm API.
// A default retry strategy applies to this operation GetOccMetricAlarm()
func (client OccMetricAlarmClient) GetOccMetricAlarm(ctx context.Context, request GetOccMetricAlarmRequest) (response GetOccMetricAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOccMetricAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccMetricAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccMetricAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccMetricAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccMetricAlarmResponse")
	}
	return
}

// getOccMetricAlarm implements the OCIOperation interface (enables retrying operations)
func (client OccMetricAlarmClient) getOccMetricAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occMetricAlarms/{occMetricAlarmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccMetricAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "occMetricAlarm", "GetOccMetricAlarm")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccMetricAlarm/GetOccMetricAlarm"
		err = common.PostProcessServiceError(err, "OccMetricAlarm", "GetOccMetricAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccMetricAlarms Gets a list of OccMetricAlarms.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/ListOccMetricAlarms.go.html to see an example of how to use ListOccMetricAlarms API.
// A default retry strategy applies to this operation ListOccMetricAlarms()
func (client OccMetricAlarmClient) ListOccMetricAlarms(ctx context.Context, request ListOccMetricAlarmsRequest) (response ListOccMetricAlarmsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccMetricAlarms, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccMetricAlarmsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccMetricAlarmsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccMetricAlarmsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccMetricAlarmsResponse")
	}
	return
}

// listOccMetricAlarms implements the OCIOperation interface (enables retrying operations)
func (client OccMetricAlarmClient) listOccMetricAlarms(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occMetricAlarms", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccMetricAlarmsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "occMetricAlarm", "ListOccMetricAlarms")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccMetricAlarmCollection/ListOccMetricAlarms"
		err = common.PostProcessServiceError(err, "OccMetricAlarm", "ListOccMetricAlarms", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccMetricAlarm Updates an existing OccMetricAlarm resource with new or modified configuration details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/UpdateOccMetricAlarm.go.html to see an example of how to use UpdateOccMetricAlarm API.
// A default retry strategy applies to this operation UpdateOccMetricAlarm()
func (client OccMetricAlarmClient) UpdateOccMetricAlarm(ctx context.Context, request UpdateOccMetricAlarmRequest) (response UpdateOccMetricAlarmResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccMetricAlarm, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccMetricAlarmResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccMetricAlarmResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccMetricAlarmResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccMetricAlarmResponse")
	}
	return
}

// updateOccMetricAlarm implements the OCIOperation interface (enables retrying operations)
func (client OccMetricAlarmClient) updateOccMetricAlarm(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occMetricAlarms/{occMetricAlarmId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccMetricAlarmResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "occMetricAlarm", "UpdateOccMetricAlarm")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccMetricAlarm/UpdateOccMetricAlarm"
		err = common.PostProcessServiceError(err, "OccMetricAlarm", "UpdateOccMetricAlarm", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
