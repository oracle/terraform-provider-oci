// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// OccDemandSignalClient a client for OccDemandSignal
type OccDemandSignalClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOccDemandSignalClientWithConfigurationProvider Creates a new default OccDemandSignal client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOccDemandSignalClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OccDemandSignalClient, err error) {
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
	return newOccDemandSignalClientFromBaseClient(baseClient, provider)
}

// NewOccDemandSignalClientWithOboToken Creates a new default OccDemandSignal client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOccDemandSignalClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OccDemandSignalClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOccDemandSignalClientFromBaseClient(baseClient, configProvider)
}

func newOccDemandSignalClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OccDemandSignalClient, err error) {
	// OccDemandSignal service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OccDemandSignal"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OccDemandSignalClient{BaseClient: baseClient}
	client.BasePath = "20240430"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OccDemandSignalClient) SetRegion(region string) {
	client.Host, _ = common.StringToRegion(region).EndpointForTemplateDottedRegion("demandsignal", "https://control-center-ds.{region}.oci.{secondLevelDomain}", "control-center-ds")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OccDemandSignalClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OccDemandSignalClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOccDemandSignalCompartment Moves a OccDemandSignal into a different compartment within the same tenancy. For information about moving resources between
// compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/ChangeOccDemandSignalCompartment.go.html to see an example of how to use ChangeOccDemandSignalCompartment API.
// A default retry strategy applies to this operation ChangeOccDemandSignalCompartment()
func (client OccDemandSignalClient) ChangeOccDemandSignalCompartment(ctx context.Context, request ChangeOccDemandSignalCompartmentRequest) (response ChangeOccDemandSignalCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeOccDemandSignalCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOccDemandSignalCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOccDemandSignalCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOccDemandSignalCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOccDemandSignalCompartmentResponse")
	}
	return
}

// changeOccDemandSignalCompartment implements the OCIOperation interface (enables retrying operations)
func (client OccDemandSignalClient) changeOccDemandSignalCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occDemandSignals/{occDemandSignalId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOccDemandSignalCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccDemandSignal/ChangeOccDemandSignalCompartment"
		err = common.PostProcessServiceError(err, "OccDemandSignal", "ChangeOccDemandSignalCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOccDemandSignal Creates a OccDemandSignal.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/CreateOccDemandSignal.go.html to see an example of how to use CreateOccDemandSignal API.
// A default retry strategy applies to this operation CreateOccDemandSignal()
func (client OccDemandSignalClient) CreateOccDemandSignal(ctx context.Context, request CreateOccDemandSignalRequest) (response CreateOccDemandSignalResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOccDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOccDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOccDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOccDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOccDemandSignalResponse")
	}
	return
}

// createOccDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client OccDemandSignalClient) createOccDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/occDemandSignals", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOccDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccDemandSignal/CreateOccDemandSignal"
		err = common.PostProcessServiceError(err, "OccDemandSignal", "CreateOccDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOccDemandSignal Deletes a OccDemandSignal.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/DeleteOccDemandSignal.go.html to see an example of how to use DeleteOccDemandSignal API.
// A default retry strategy applies to this operation DeleteOccDemandSignal()
func (client OccDemandSignalClient) DeleteOccDemandSignal(ctx context.Context, request DeleteOccDemandSignalRequest) (response DeleteOccDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOccDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOccDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOccDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOccDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOccDemandSignalResponse")
	}
	return
}

// deleteOccDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client OccDemandSignalClient) deleteOccDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/occDemandSignals/{occDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOccDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccDemandSignal/DeleteOccDemandSignal"
		err = common.PostProcessServiceError(err, "OccDemandSignal", "DeleteOccDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOccDemandSignal Gets information about a OccDemandSignal.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/GetOccDemandSignal.go.html to see an example of how to use GetOccDemandSignal API.
// A default retry strategy applies to this operation GetOccDemandSignal()
func (client OccDemandSignalClient) GetOccDemandSignal(ctx context.Context, request GetOccDemandSignalRequest) (response GetOccDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOccDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOccDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOccDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOccDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOccDemandSignalResponse")
	}
	return
}

// getOccDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client OccDemandSignalClient) getOccDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occDemandSignals/{occDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOccDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccDemandSignal/GetOccDemandSignal"
		err = common.PostProcessServiceError(err, "OccDemandSignal", "GetOccDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOccDemandSignals Gets a list of OccDemandSignals.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/ListOccDemandSignals.go.html to see an example of how to use ListOccDemandSignals API.
// A default retry strategy applies to this operation ListOccDemandSignals()
func (client OccDemandSignalClient) ListOccDemandSignals(ctx context.Context, request ListOccDemandSignalsRequest) (response ListOccDemandSignalsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOccDemandSignals, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOccDemandSignalsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOccDemandSignalsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOccDemandSignalsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOccDemandSignalsResponse")
	}
	return
}

// listOccDemandSignals implements the OCIOperation interface (enables retrying operations)
func (client OccDemandSignalClient) listOccDemandSignals(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/occDemandSignals", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOccDemandSignalsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccDemandSignalCollection/ListOccDemandSignals"
		err = common.PostProcessServiceError(err, "OccDemandSignal", "ListOccDemandSignals", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchOccDemandSignal Updates the data of an OccDemandSignal.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/PatchOccDemandSignal.go.html to see an example of how to use PatchOccDemandSignal API.
// A default retry strategy applies to this operation PatchOccDemandSignal()
func (client OccDemandSignalClient) PatchOccDemandSignal(ctx context.Context, request PatchOccDemandSignalRequest) (response PatchOccDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchOccDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchOccDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchOccDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchOccDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchOccDemandSignalResponse")
	}
	return
}

// patchOccDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client OccDemandSignalClient) patchOccDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/occDemandSignals/{occDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchOccDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccDemandSignal/PatchOccDemandSignal"
		err = common.PostProcessServiceError(err, "OccDemandSignal", "PatchOccDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOccDemandSignal Updates a OccDemandSignal.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/demandsignal/UpdateOccDemandSignal.go.html to see an example of how to use UpdateOccDemandSignal API.
// A default retry strategy applies to this operation UpdateOccDemandSignal()
func (client OccDemandSignalClient) UpdateOccDemandSignal(ctx context.Context, request UpdateOccDemandSignalRequest) (response UpdateOccDemandSignalResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOccDemandSignal, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOccDemandSignalResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOccDemandSignalResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOccDemandSignalResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOccDemandSignalResponse")
	}
	return
}

// updateOccDemandSignal implements the OCIOperation interface (enables retrying operations)
func (client OccDemandSignalClient) updateOccDemandSignal(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/occDemandSignals/{occDemandSignalId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOccDemandSignalResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/occds/20240430/OccDemandSignal/UpdateOccDemandSignal"
		err = common.PostProcessServiceError(err, "OccDemandSignal", "UpdateOccDemandSignal", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
